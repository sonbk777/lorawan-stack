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

import Link from '@ttn-lw/components/link'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './row.styl'

class Row extends React.Component {
  @bind
  onClick(evt) {
    const { id, onClick } = this.props

    onClick(id, evt)
  }

  @bind
  onKeyDown(evt) {
    const { id, onClick } = this.props
    if (evt.key === 'Enter') {
      onClick(id, evt)
    }
  }

  @bind
  onMouseDown(evt) {
    const { id, onMouseDown } = this.props

    onMouseDown(id, evt)
  }

  get clickListener() {
    const { body, clickable } = this.props

    if (body && clickable) {
      return this.onClick
    }
  }

  get tabIndex() {
    const { body, clickable } = this.props

    return body && clickable ? 0 : -1
  }

  render() {
    const { className, children, clickable, head, body, footer, linkTo } = this.props

    const rowClassNames = classnames(className, style.row, {
      [style.clickable]: body && clickable,
      [style.rowHead]: head,
      [style.rowBody]: body,
      [style.rowFooter]: footer,
    })

    const Row = linkTo && clickable ? Link : 'div'

    return (
      <Row
        className={rowClassNames}
        onKeyDown={this.onKeyDown}
        onClick={this.clickListener}
        onMouseDown={this.onMouseDown}
        tabIndex={this.tabIndex.toString()}
        to={linkTo}
        role="row"
      >
        {children}
      </Row>
    )
  }
}

Row.propTypes = {
  /** A flag indicating whether the row is wrapping the body of a table. */
  body: PropTypes.bool,
  children: PropTypes.node,
  className: PropTypes.string,
  /** A flag indicating whether the row is clickable. */
  clickable: PropTypes.bool,
  /** A flag indicating whether the row is wrapping the footer of a table. */
  footer: PropTypes.bool,
  /** A flag indicating whether the row is wrapping the head of a table. */
  head: PropTypes.bool,
  /** The identifier of the row. */
  id: PropTypes.number,
  /** The href to be passed as `to` prop to the `<Link />` component that wraps the row. */
  linkTo: PropTypes.string,
  /**
   * Function to be called when the row gets clicked. The identifier of the row
   * is passed as an argument.
   */
  onClick: PropTypes.func,
  onMouseDown: PropTypes.func,
}

Row.defaultProps = {
  children: undefined,
  className: undefined,
  clickable: true,
  head: false,
  body: false,
  footer: false,
  onClick: () => null,
  onMouseDown: () => null,
  id: undefined,
  linkTo: undefined,
}

export default Row
