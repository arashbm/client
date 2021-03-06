// @flow
import React, {Component} from 'react'
import debounce from 'lodash/debounce'
import dumbComponentMap from './dumb-component-map.native'
import type {Props} from './dumb-sheet.render'
import {Box, Text, Input, Button} from '../common-adapters'
import {ScrollView} from 'react-native'
import {globalStyles} from '../styles/style-guide'

class Render extends Component<void, Props, any> {
  state: any;
  _onFilterChange: (a: any) => void;

  constructor (props: Props) {
    super(props)

    this.state = {
      filterShow: false,
    }

    this._onFilterChange = debounce(filter => {
      this.props.onDebugConfigChange({
        dumbFilter: filter,
      })
    }, 300)
  }

  render () {
    const filter = this.props.dumbFilter.toLowerCase()
    const components = []
    const componentsOnly = []
    const parentPropsOnly = []

    Object.keys(dumbComponentMap).forEach(key => {
      if (filter && key.toLowerCase().indexOf(filter) === -1) {
        return
      }

      const map = dumbComponentMap[key]
      const Component = map.component
      Object.keys(map.mocks).forEach((mockKey, idx) => {
        const mock = {...map.mocks[mockKey]}
        const parentProps = mock.parentProps
        mock.parentProps = undefined

        components.push(
          <Box key={mockKey} style={styleBox}>
            <Text type='Body' style={{marginBottom: 5}}>{key}: {mockKey}</Text>
            <Box style={{flex: 1}} {...parentProps}>
              <Component key={mockKey} {...mock} />
            </Box>
          </Box>
        )
        componentsOnly.push(<Component key={mockKey} {...mock} />)
        parentPropsOnly.push(parentProps)
      })
    })

    const ToShow = components[this.props.dumbIndex % components.length]

    if (this.props.dumbFullscreen) {
      return (
        <Box style={{flex: 1}} {...parentPropsOnly[this.props.dumbIndex % components.length]}>
          {componentsOnly[this.props.dumbIndex % components.length]}
        </Box>
      )
    }

    return (
      <Box style={{flex: 1}}>
        <ScrollView style={{flex: 1}}>
          {ToShow}
        </ScrollView>
        <Box style={stylesControls}>
          <Text type='BodySmall'>{this.props.dumbIndex}</Text>
          {this.state.filterShow && <Box style={{...globalStyles.flexBoxColumn, backgroundColor: 'red', width: 200}}><Input style={inputStyle} value={filter} onChangeText={filter => this._onFilterChange(filter.toLowerCase())} /></Box>}
          <Button type='Primary' style={stylesButton} label='...' onClick={() => { this.setState({filterShow: !this.state.filterShow}) }} />
          <Button type='Primary' style={stylesButton} label='<' onClick={() => { this._incremement(false) }} />
          <Button type='Primary' style={stylesButton} label='>' onClick={() => { this._incremement(true) }} />
        </Box>
      </Box>
    )
  }

  _incremement (up: boolean) {
    let next = Math.max(0, this.props.dumbIndex + (up ? 1 : -1))
    this.props.onDebugConfigChange({
      dumbIndex: next,
    })
  }
}

const styleBox = {
  ...globalStyles.flexBoxColumn,
  flex: 1,
  height: 800,
  paddingTop: 20,
  marginTop: 10,
}

const inputStyle = {
  height: 40,
  marginTop: 0,
}

const stylesControls = {
  ...globalStyles.flexBoxRow,
  position: 'absolute',
  top: 0,
  right: 0,
}

const stylesButton = {
  width: 20,
  height: 20,
  overflow: 'hidden',
  padding: 0,
  margin: 0,
  paddingTop: 0,
  paddingLeft: 0,
  paddingRight: 0,
  paddingBottom: 20,
  borderRadius: 10,
}

export default Render
