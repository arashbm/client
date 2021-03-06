/* @flow */

import React, {Component} from 'react'
import {Icon, Text} from '../common-adapters/index'
import {globalStyles, globalColors} from '../styles/style-guide'
import {stateColors} from '../util/tracker'

import type {HeaderProps} from './header.render'

export default class HeaderRender extends Component {
  props: HeaderProps;
  state: {showCloseWarning: boolean};

  constructor (props: HeaderProps) {
    super(props)
    this.state = {showCloseWarning: false}
  }

  render () {
    const isWarningAboutTrackerShowingUpLater = this.props.loggedIn && !this.props.currentlyFollowing && this.state.showCloseWarning
    const headerText = isWarningAboutTrackerShowingUpLater ? 'You will see this window every time you access this folder.' : this.props.reason

    const trackerStateColors = stateColors(this.props)
    const headerBackgroundColor = isWarningAboutTrackerShowingUpLater ? globalColors.yellow : trackerStateColors.header.background
    const headerTextColor = isWarningAboutTrackerShowingUpLater ? globalColors.brown_60 : trackerStateColors.header.text

    return (
      <div style={styleOuter}>
        <div style={{...styleHeader, backgroundColor: headerBackgroundColor}}>
          <Text type='BodySemibold' lineClamp={2} style={{...styleText, color: headerTextColor, ...(isWarningAboutTrackerShowingUpLater ? {zIndex: 2} : {})}}>{headerText}</Text>
          <Icon type='iconfont-close' style={styleClose}
            onClick={() => this.props.onClose()}
            onMouseEnter={() => this.closeMouseEnter()}
            onMouseLeave={() => this.closeMouseLeave()} />
        </div>
      </div>
    )
  }

  closeMouseEnter () {
    this.setState({showCloseWarning: true})
  }

  closeMouseLeave () {
    this.setState({showCloseWarning: false})
  }
}

const styleOuter = {
  position: 'relative',
}

const styleHeader = {
  ...globalStyles.windowDragging,
  cursor: 'default',
  position: 'absolute',
  top: 0,
  ...globalStyles.flexBoxRow,
  height: 90,
  width: 320,
}

const styleClose = {
  ...globalStyles.clickable,
  ...globalStyles.windowDraggingClickable,
  zIndex: 2,
  position: 'absolute',
  top: 7,
  right: 9,
}

const styleText = {
  ...globalStyles.flexBoxRow,
  flex: 1,
  alignItems: 'center',
  justifyContent: 'center',
  color: globalColors.white,
  marginLeft: 30,
  marginRight: 30,
  marginBottom: 32,
  fontSize: 14,
  textAlign: 'center',
  lineHeight: 'normal',
}
