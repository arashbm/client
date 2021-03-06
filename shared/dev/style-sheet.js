// @flow
import React, {Component} from 'react'
import Render from './style-sheet.render'
import {globalColors} from '../styles/style-guide'

class StyleSheet extends Component {
  render () {
    return <Render
      {...this.props}
      colors={globalColors}
    />
  }

  static parseRoute () {
    return {
      componentAtTop: {title: 'Stylesheet'},
    }
  }
}

export default StyleSheet
