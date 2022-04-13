// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

import React from 'react'
import { connect } from 'react-redux'
import { Switch, Route } from 'react-router-dom'

import applicationIcon from '@assets/misc/application.svg'

import SideNavigation from '@ttn-lw/components/navigation/side'

import withRequest from '@ttn-lw/lib/components/with-request'
import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'
import NotFoundRoute from '@ttn-lw/lib/components/not-found-route'

import OAuthClientOverview from '@account/views/oauth-client-overview'

import { selectApplicationSiteName } from '@ttn-lw/lib/selectors/env'
import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { mayPerformAdminActions } from '@account/lib/feature-checks'

import { getClient } from '@account/store/actions/clients'

import {
  selectClientFetching,
  selectClientById,
  selectClientError,
} from '@account/store/selectors/clients'

const OAuthClient = props => {
  console.log(props)
  const {
    oauthClientId,
    match: { url: matchedUrl, path },
    oauthClient,
    siteName,
  } = props
  console.log(oauthClientId)
  console.log(oauthClient)
  const name = oauthClient.name || oauthClientId

  return (
    <React.Fragment>
      <IntlHelmet titleTemplate={`%s - ${name} - ${siteName}`} />
      <SideNavigation
        header={{
          icon: applicationIcon,
          iconAlt: sharedMessages.application,
          title: name,
          to: matchedUrl,
        }}
      >
        {mayPerformAdminActions && (
          <SideNavigation.Item
            title={sharedMessages.overview}
            path={matchedUrl}
            icon="overview"
            exact
          />
        )}
        {mayPerformAdminActions && (
          <SideNavigation.Item
            title={sharedMessages.collaborators}
            path={`${matchedUrl}/collaborators`}
            icon="organization"
          />
        )}
        {mayPerformAdminActions && (
          <SideNavigation.Item
            title={sharedMessages.generalSettings}
            path={`${matchedUrl}/general-settings`}
            icon="general_settings"
          />
        )}
      </SideNavigation>
      <Switch>
        <Route exact path={`${path}`} component={OAuthClientOverview} />
        <NotFoundRoute />
      </Switch>
    </React.Fragment>
  )
}

OAuthClient.propTypes = {
  match: PropTypes.match.isRequired,
  oauthClientId: PropTypes.string.isRequired,
}

export default connect(
  (state, props) => ({
    oauthClientId: props.match.params.id,
    fetching: selectClientFetching(state),
    oauthClient: selectClientById(state, props.match.params.id),
    error: selectClientError(state),
    siteName: selectApplicationSiteName(),
  }),
  dispatch => ({
    loadData: id => {
      dispatch(getClient(id, ['name', 'description', 'state', 'state_description']))
    },
  }),
)(withRequest(({ oauthClientId, loadData }) => loadData(oauthClientId))(OAuthClient))
