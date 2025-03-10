// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { createLogic } from 'redux-logic'

import ONLINE_STATUS from '@ttn-lw/constants/online-status'
import CONNECTION_STATUS from '@console/constants/connection-status'
import EVENT_TAIL from '@console/constants/event-tail'

import { getCombinedDeviceId } from '@ttn-lw/lib/selectors/id'
import { isUnauthenticatedError, isNetworkError, isTimeoutError } from '@ttn-lw/lib/errors/utils'
import { SET_CONNECTION_STATUS, setStatusChecking } from '@ttn-lw/lib/store/actions/status'
import { selectIsOnlineStatus } from '@ttn-lw/lib/store/selectors/status'

import {
  createStartEventsStreamActionType,
  createStopEventsStreamActionType,
  createStartEventsStreamFailureActionType,
  createStartEventsStreamSuccessActionType,
  createEventStreamClosedActionType,
  createGetEventMessageFailureActionType,
  createGetEventMessageSuccessActionType,
  createSetEventsFilterActionType,
  getEventMessageSuccess,
  getEventMessageFailure,
  startEventsStreamFailure,
  startEventsStreamSuccess,
  eventStreamClosed,
  startEventsStream,
} from '@console/store/actions/events'

import {
  createEventsStatusSelector,
  createEventsInterruptedSelector,
  createInterruptedStreamsSelector,
  createLatestEventSelector,
  createLatestClearedEventSelector,
  createEventsFilterSelector,
} from '@console/store/selectors/events'
import { selectDeviceById } from '@console/store/selectors/devices'

/**
 * Creates `redux-logic` logic from processing entity events.
 *
 * @param {string} reducerName - The name of an entity used to create the events reducer.
 * @param {string} entityName - The name of an entity.
 * @param {Function} onEventsStart - A function to be called to start the events stream.
 * Should accept a list of entity ids.
 * @returns {object} - The `redux-logic` (decorated) logic.
 */
const createEventsConnectLogics = (reducerName, entityName, onEventsStart) => {
  const START_EVENTS = createStartEventsStreamActionType(reducerName)
  const START_EVENTS_SUCCESS = createStartEventsStreamSuccessActionType(reducerName)
  const START_EVENTS_FAILURE = createStartEventsStreamFailureActionType(reducerName)
  const STOP_EVENTS = createStopEventsStreamActionType(reducerName)
  const EVENT_STREAM_CLOSED = createEventStreamClosedActionType(reducerName)
  const GET_EVENT_MESSAGE_FAILURE = createGetEventMessageFailureActionType(reducerName)
  const GET_EVENT_MESSAGE_SUCCESS = createGetEventMessageSuccessActionType(reducerName)
  const SET_EVENT_FILTER = createSetEventsFilterActionType(reducerName)
  const startEventsSuccess = startEventsStreamSuccess(reducerName)
  const startEventsFailure = startEventsStreamFailure(reducerName)
  const closeEvents = eventStreamClosed(reducerName)
  const startEvents = startEventsStream(reducerName)
  const getEventSuccess = getEventMessageSuccess(reducerName)
  const getEventFailure = getEventMessageFailure(reducerName)
  const selectEntityEventsStatus = createEventsStatusSelector(entityName)
  const selectEntityEventsInterrupted = createEventsInterruptedSelector(entityName)
  const selectInterruptedStreams = createInterruptedStreamsSelector(entityName)
  const selectLatestEvent = createLatestEventSelector(entityName)
  const selectLatestClearedEvent = createLatestClearedEventSelector(entityName)
  const selectEventFilter = createEventsFilterSelector(entityName)

  let channel = null

  return [
    createLogic({
      type: START_EVENTS,
      cancelType: [START_EVENTS_FAILURE],
      warnTimeout: 0,
      processOptions: {
        dispatchMultiple: true,
      },
      validate: ({ getState, action = {} }, allow, reject) => {
        if (!action.id) {
          reject()
          return
        }

        const id = typeof action.id === 'object' ? getCombinedDeviceId(action.id) : action.id

        // Only proceed if not already connected and online.
        const state = getState()
        const isOnline = selectIsOnlineStatus(state)
        const status = selectEntityEventsStatus(state, id)
        const connected = status === CONNECTION_STATUS.CONNECTED
        const connecting = status === CONNECTION_STATUS.CONNECTING
        if (connected || connecting || !isOnline) {
          reject()
          return
        }

        allow(action)
      },
      process: async ({ getState, action }, dispatch) => {
        const { id, silent } = action

        const idString = typeof action.id === 'object' ? getCombinedDeviceId(action.id) : action.id

        // Only get historical events emitted after the latest event or latest
        // cleared event in the store to avoid duplicate historical events.
        const state = getState()
        const latestEvent = selectLatestEvent(state, idString)
        const latestClearedEvent = selectLatestClearedEvent(state, idString)
        const latestEventTime = Boolean(latestEvent) ? latestEvent.time : ''
        const latestClearedEventTime = Boolean(latestClearedEvent) ? latestClearedEvent.time : ''
        const after =
          (latestEventTime > latestClearedEventTime ? latestEventTime : latestClearedEventTime) ||
          undefined
        const filter = selectEventFilter(state, idString)
        const filterRegExp = Boolean(filter) ? filter.filterRegExp : undefined

        try {
          channel = await onEventsStart([id], filterRegExp, EVENT_TAIL, after)

          channel.on('start', () => dispatch(startEventsSuccess(id, { silent })))
          channel.on('chunk', message => dispatch(getEventSuccess(id, message)))
          channel.on('error', error => dispatch(getEventFailure(id, error)))
          channel.on('close', wasClientRequest =>
            dispatch(closeEvents(id, { silent: wasClientRequest })),
          )

          channel.open()
        } catch (error) {
          if (isUnauthenticatedError(error)) {
            // The user is no longer authenticated; reinitiate the auth flow
            // by refreshing the page.
            window.location.reload()
          } else {
            dispatch(startEventsFailure(id, error))
          }
        }
      },
    }),
    createLogic({
      type: [STOP_EVENTS, START_EVENTS_FAILURE],
      validate: ({ getState, action = {} }, allow, reject) => {
        if (!action.id) {
          reject()
          return
        }

        const id = typeof action.id === 'object' ? getCombinedDeviceId(action.id) : action.id

        // Only proceed if connected.
        const status = selectEntityEventsStatus(getState(), id)
        const connected = status === CONNECTION_STATUS.CONNECTED
        const connecting = status === CONNECTION_STATUS.CONNECTING
        if (!connected && !connecting) {
          reject()
          return
        }

        allow(action)
      },
      process: ({ action }, dispatch, done) => {
        if (channel) {
          try {
            channel.close()
          } catch (error) {
            if (isNetworkError(error) || isTimeoutError(action.payload)) {
              // Set the connection status to `checking` to trigger connection checks
              // and detect possible offline state.
              dispatch(setStatusChecking())

              // In case of a network error, the connection could not be closed
              // since the network connection is disrupted. We can regard this
              // as equivalent to a closed connection.
              return done()
            }
            throw error
          }
        }
        done()
      },
    }),
    createLogic({
      type: [GET_EVENT_MESSAGE_FAILURE, EVENT_STREAM_CLOSED],
      cancelType: [START_EVENTS_SUCCESS, GET_EVENT_MESSAGE_SUCCESS, STOP_EVENTS],
      warnTimeout: 0,
      validate: ({ getState, action = {} }, allow, reject) => {
        if (!action.id) {
          reject()
          return
        }

        const id = typeof action.id === 'object' ? getCombinedDeviceId(action.id) : action.id

        // Only proceed if connected and not interrupted.
        const status = selectEntityEventsStatus(getState(), id)
        const connected = status === CONNECTION_STATUS.CONNECTED
        const interrupted = selectEntityEventsInterrupted(getState(), id)
        if (!connected && interrupted) {
          reject()
        }

        allow(action)
      },
      process: ({ getState, action }, dispatch, done) => {
        const isOnline = selectIsOnlineStatus(getState())

        // If the app is not offline, try to reconnect periodically.
        if (isOnline) {
          const reconnector = setInterval(() => {
            // Only proceed if still disconnected, interrupted and online.
            const state = getState()
            const id = typeof action.id === 'object' ? getCombinedDeviceId(action.id) : action.id
            const status = selectEntityEventsStatus(state, id)
            const disconnected = status === CONNECTION_STATUS.DISCONNECTED
            const interrupted = selectEntityEventsInterrupted(state, id)
            const isOnline = selectIsOnlineStatus(state)
            if (disconnected && interrupted && isOnline) {
              dispatch(startEvents(action.id))
            } else {
              clearInterval(reconnector)
              done()
            }
          }, 5000)
        } else {
          done()
        }
      },
    }),
    createLogic({
      type: SET_CONNECTION_STATUS,
      process: ({ getState, action }, dispatch, done) => {
        const isOnline = action.payload.onlineStatus === ONLINE_STATUS.ONLINE

        if (isOnline) {
          const state = getState()
          for (const id in selectInterruptedStreams(state)) {
            const status = selectEntityEventsStatus(state, id)
            const disconnected = status === CONNECTION_STATUS.DISCONNECTED
            // If the app reconnected to the internet and there is a pending
            // interrupted stream connection, try to reconnect.
            if (disconnected) {
              let ids = id
              // For end devices, it's necessary to retrieve the entity ids object
              // back from the combined id string.
              if (entityName === 'devices' && typeof id === 'string') {
                const selectedDevice = selectDeviceById(state, id)
                if (!selectedDevice || !selectedDevice.ids) {
                  continue
                }
                ids = selectedDevice.ids
              }

              dispatch(dispatch(startEvents(ids)))
            }
          }
        }

        done()
      },
    }),
    createLogic({
      type: SET_EVENT_FILTER,
      process: async ({ action }, dispatch, done) => {
        if (channel) {
          try {
            await channel.close()
          } catch (error) {
            if (isNetworkError(error) || isTimeoutError(action.payload)) {
              dispatch(setStatusChecking())
            } else {
              throw error
            }
          } finally {
            dispatch(startEvents(action.id, { silent: true }))
          }
        }
        done()
      },
    }),
  ]
}

export default createEventsConnectLogics
