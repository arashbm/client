// @flow
import Platform, {OS} from '../constants/platform'
import React from 'react'
import Text from './text'
import _ from 'lodash'
import type {Props} from './meta'
import {View} from 'react-native'
import {globalColors} from '../styles/style-guide'

const isAndroid = Platform.OS_ANDROID === OS

const Meta = ({title, style}: Props) => (
  <View style={{
    borderRadius: 1,
    paddingLeft: 2,
    paddingRight: 2,
    paddingTop: isAndroid ? 1 : 2,
    paddingBottom: 1,
    alignSelf: 'flex-start',
    ..._.omit(style, ['color']),
  }}>
    <Text type='Header' style={{
      color: style && style.color || globalColors.white,
      fontSize: 10,
      height: 11,
      lineHeight: 11,
    }}>{title && title.toUpperCase()}</Text>
  </View>
)

export default Meta
