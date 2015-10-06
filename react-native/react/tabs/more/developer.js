'use strict'
/* @flow */

import React, { Component, StyleSheet, TextInput, View, Text } from 'react-native'
import Button from '../../common-adapters/button'
import commonStyles from '../../styles/common'
import * as SearchActions from '../../actions/search'
import { navigateUp } from '../../actions/router'

export default class Developer extends Component {
  constructor (props) {
    super(props)

    this.state = { }
  }

  render () {
    return (
      <View style={styles.container}>
        <Text style={[{textAlign: 'center', marginBottom: 75}, commonStyles.h1]}>Dev settings</Text>
        <TextInput
          style={styles.input}
          placeholder='Some setting'
          value='TODO'
          enablesReturnKeyAutomatically
          returnKeyType='next'
          autoCorrect={false}
          onChangeText={() => { console.log('typing') }}
        />
        <Button
          buttonStyle={{backgroundColor: 'blue'}}
          onPress={ () => this.props.dispatch(SearchActions.pushNewSearch('more')) }
          title='Launch search' />
        <Text onPress={() => this.props.dispatch(navigateUp())}>Back</Text>
      </View>
    )
  }

  static parseRoute (store, currentPath, nextPath) {
    const componentAtTop = {
      title: 'Developer',
      component: Developer,
      hideNavBar: true
    }

    return {
      componentAtTop,
      parseNextRoute: null
    }
  }
}

Developer.propTypes = {
  dispatch: React.PropTypes.func.isRequired
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'stretch',
    backgroundColor: 'red'
  },
  input: {
    height: 40,
    marginBottom: 5,
    marginLeft: 10,
    marginRight: 10,
    borderWidth: 0.5,
    borderColor: '#0f0f0f',
    fontSize: 13,
    padding: 4
  },
  submitWrapper: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    marginTop: 10
  }
})
