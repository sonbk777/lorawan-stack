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

import React from 'react'
import classnames from 'classnames'
import bind from 'autobind-decorator'

import Tooltip from '@ttn-lw/components/tooltip'

import PropTypes from '@ttn-lw/lib/prop-types'

import Tag from '..'

import style from './group.styl'

const measureWidth = element => {
  if (!element) {
    return 0
  }

  return element.current.clientWidth
}

// The width in pixels for the left tag.
const LEFT_TAG_WIDTH = 40
// The space between the tags.
const TAG_SPACE_WIDTH = 3

class TagGroup extends React.Component {
  static propTypes = {
    className: PropTypes.string,
    tagMaxWidth: PropTypes.number.isRequired,
    tags: PropTypes.arrayOf(PropTypes.shape(Tag.PropTypes)).isRequired,
  }

  static defaultProps = {
    className: undefined,
  }

  constructor(props) {
    super(props)

    this.state = {
      left: 0,
    }

    this.element = React.createRef()
  }

  componentDidMount() {
    window.addEventListener('resize', this.handleWindowResize)

    this.handleWindowResize()
  }

  componentDidUpdate(prevProps) {
    const props = this.props

    // Calculate fit on any props change.
    if (prevProps.tags !== props.tags) {
      this.checkTagsFit()
    }
  }

  checkTagsFit() {
    this.handleWindowResize()
  }

  componentWillUnmount() {
    window.removeEventListener('resize', this.handleWindowResize)
  }

  @bind
  handleWindowResize() {
    const { tags, tagMaxWidth } = this.props

    const containerWidth = measureWidth(this.element)
    const totalTagCount = tags.length
    const possibleFitCount = Math.floor(containerWidth / tagMaxWidth) || 1

    // Count for the left tag and paddings between tags.
    const leftTagWidth = totalTagCount !== possibleFitCount ? LEFT_TAG_WIDTH : 0
    const spaceWidth = possibleFitCount > 1 ? possibleFitCount * TAG_SPACE_WIDTH : 0

    const finalAvailableWidth = containerWidth - leftTagWidth - spaceWidth
    const finalLeft = Math.floor(finalAvailableWidth / tagMaxWidth) || 1

    this.setState({
      left: totalTagCount - finalLeft,
    })
  }

  render() {
    const { className, tags } = this.props
    const { left } = this.state

    const ts = tags.slice(0, tags.length - left)
    const leftGroup = <div className={style.leftGroup}>{tags.slice(tags.length - left)}</div>

    return (
      <div ref={this.element} className={classnames(className, style.group)}>
        {ts}
        {left > 0 && (
          <Tooltip content={leftGroup}>
            <Tag content={`+${left}`} />
          </Tooltip>
        )}
      </div>
    )
  }
}

export default TagGroup
