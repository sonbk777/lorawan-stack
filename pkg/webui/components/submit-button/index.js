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

import Button from '@ttn-lw/components/button'

import PropTypes from '@ttn-lw/lib/prop-types'

class SubmitButton extends React.PureComponent {
  static propTypes = {
    disabled: PropTypes.bool,
    icon: PropTypes.string,
    isSubmitting: PropTypes.bool.isRequired,
    isValidating: PropTypes.bool.isRequired,
    message: PropTypes.message,
  }

  static defaultProps = {
    message: undefined,
  }

  static defaultProps = {
    disabled: false,
    icon: undefined,
  }
  render() {
    const { message, icon, disabled, isSubmitting, isValidating, ...rest } = this.props

    const buttonLoading = isSubmitting || isValidating

    return (
      <Button
        primary
        {...rest}
        type="submit"
        icon={icon}
        message={message}
        disabled={disabled}
        busy={buttonLoading}
      />
    )
  }
}

export default SubmitButton
