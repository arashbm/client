// @flow

// This file is auto-generated by client/protocol/Makefile.
import * as gregor1 from './flow-types-gregor'

import engine from '../../engine'
import type {$Exact} from './more'
export type int = number
export type int64 = number
export type uint = number
export type uint64 = number
export type long = number
export type double = number
export type bytes = any
export type RPCError = {
  code: number,
  desc: string
}
export type Conversation = {
  metadata: ConversationMetadata,
  readerInfo?: ?ConversationReaderInfo,
  maxHeaders?: ?Array<MessageServerHeader>,
}

export type ConversationID = uint64

export type ConversationIDTriple = {
  tlfid: TLFID,
  topicType: TopicType,
  topicID: TopicID,
}

export type ConversationMetadata = {
  idTriple: ConversationIDTriple,
  conversationID: ConversationID,
}

export type ConversationReaderInfo = {
  mtime: gregor1.Time,
  readMsgid: MessageID,
  maxMsgid: MessageID,
}

export type EncryptedData = {
  v: int,
  e: bytes,
  n: bytes,
}

export type InboxView = {
  conversations?: ?Array<Conversation>,
  pagination?: ?Pagination,
}

export type MessageBoxed = {
  serverHeader?: ?MessageServerHeader,
  clientHeader: MessageClientHeader,
  headerSignature: SignatureInfo,
  bodyCiphertext: EncryptedData,
  bodySignature: SignatureInfo,
  keyGeneration: int,
}

export type MessageClientHeader = {
  conv: ConversationIDTriple,
  tlfName: string,
  messageType: MessageType,
  prev?: ?Array<MessagePreviousPointer>,
  sender: gregor1.UID,
  senderDevice: gregor1.DeviceID,
}

export type MessageID = uint

export type MessagePreviousPointer = {
  id: MessageID,
  hash: bytes,
}

export type MessageServerHeader = {
  messageType: MessageType,
  messageID: MessageID,
  sender: gregor1.UID,
  senderDevice: gregor1.DeviceID,
  supersededBy: MessageID,
  ctime: gregor1.Time,
}

export type MessageType =
    0 // NONE_0
  | 1 // TEXT_1
  | 2 // ATTACHMENT_2
  | 3 // EDIT_3
  | 4 // DELETE_4
  | 5 // METADATA_5

export type NewMessagePayload = {
  Action: string,
  convID: ConversationID,
  message: MessageBoxed,
}

export type Pagination = {
  next: bytes,
  previous: bytes,
  num: int,
  last: boolean,
}

export type SignatureInfo = {
  v: int,
  s: bytes,
  k: bytes,
}

export type TLFID = bytes

export type ThreadID = bytes

export type ThreadViewBoxed = {
  messages?: ?Array<MessageBoxed>,
  pagination?: ?Pagination,
}

export type TopicID = bytes

export type TopicType =
    0 // NONE_0
  | 1 // CHAT_1

export type remoteGetInboxRemoteRpcParam = $Exact<{
  pagination: (null | Pagination)
}>

type remoteGetInboxRemoteResult = InboxView

export function remoteGetInboxRemoteRpc (request: $Exact<{
  param: remoteGetInboxRemoteRpcParam,
  waitingHandler?: (waiting: boolean, method: string, sessionID: string) => void,
  incomingCallMap?: incomingCallMapType,
  callback?: (null | (err: ?any, response: remoteGetInboxRemoteResult) => void)}>) {
  engine.rpc({...request, method: 'remote.getInboxRemote'})
}
export type remoteGetThreadRemoteRpcParam = $Exact<{
  conversationID: ConversationID,
  markAsRead: boolean,
  pagination: (null | Pagination)
}>

type remoteGetThreadRemoteResult = ThreadViewBoxed

export function remoteGetThreadRemoteRpc (request: $Exact<{
  param: remoteGetThreadRemoteRpcParam,
  waitingHandler?: (waiting: boolean, method: string, sessionID: string) => void,
  incomingCallMap?: incomingCallMapType,
  callback?: (null | (err: ?any, response: remoteGetThreadRemoteResult) => void)}>) {
  engine.rpc({...request, method: 'remote.getThreadRemote'})
}
export type remoteMarkAsReadRpcParam = $Exact<{
  conversationID: ConversationID,
  msgID: MessageID
}>

export function remoteMarkAsReadRpc (request: $Exact<{
  param: remoteMarkAsReadRpcParam,
  waitingHandler?: (waiting: boolean, method: string, sessionID: string) => void,
  incomingCallMap?: incomingCallMapType,
  callback?: (null | (err: ?any) => void)}>) {
  engine.rpc({...request, method: 'remote.markAsRead'})
}
export type remoteNewConversationRemoteRpcParam = $Exact<{
  idTriple: ConversationIDTriple
}>

type remoteNewConversationRemoteResult = ConversationID

export function remoteNewConversationRemoteRpc (request: $Exact<{
  param: remoteNewConversationRemoteRpcParam,
  waitingHandler?: (waiting: boolean, method: string, sessionID: string) => void,
  incomingCallMap?: incomingCallMapType,
  callback?: (null | (err: ?any, response: remoteNewConversationRemoteResult) => void)}>) {
  engine.rpc({...request, method: 'remote.newConversationRemote'})
}
export type remotePostRemoteRpcParam = $Exact<{
  conversationID: ConversationID,
  messageBoxed: MessageBoxed
}>

type remotePostRemoteResult = MessageID

export function remotePostRemoteRpc (request: $Exact<{
  param: remotePostRemoteRpcParam,
  waitingHandler?: (waiting: boolean, method: string, sessionID: string) => void,
  incomingCallMap?: incomingCallMapType,
  callback?: (null | (err: ?any, response: remotePostRemoteResult) => void)}>) {
  engine.rpc({...request, method: 'remote.postRemote'})
}
export type rpc =
    remoteGetInboxRemoteRpc
  | remoteGetThreadRemoteRpc
  | remoteMarkAsReadRpc
  | remoteNewConversationRemoteRpc
  | remotePostRemoteRpc

export type incomingCallMapType = $Exact<{
  'keybase.1.remote.getInboxRemote'?: (
    params: $Exact<{
      pagination: (null | Pagination)
    }>,
    response: {
      error: (err: RPCError) => void,
      result: (result: remoteGetInboxRemoteResult) => void
    }
  ) => void,
  'keybase.1.remote.getThreadRemote'?: (
    params: $Exact<{
      conversationID: ConversationID,
      markAsRead: boolean,
      pagination: (null | Pagination)
    }>,
    response: {
      error: (err: RPCError) => void,
      result: (result: remoteGetThreadRemoteResult) => void
    }
  ) => void,
  'keybase.1.remote.postRemote'?: (
    params: $Exact<{
      conversationID: ConversationID,
      messageBoxed: MessageBoxed
    }>,
    response: {
      error: (err: RPCError) => void,
      result: (result: remotePostRemoteResult) => void
    }
  ) => void,
  'keybase.1.remote.newConversationRemote'?: (
    params: $Exact<{
      idTriple: ConversationIDTriple
    }>,
    response: {
      error: (err: RPCError) => void,
      result: (result: remoteNewConversationRemoteResult) => void
    }
  ) => void,
  'keybase.1.remote.markAsRead'?: (
    params: $Exact<{
      conversationID: ConversationID,
      msgID: MessageID
    }>,
    response: {
      error: (err: RPCError) => void,
      result: () => void
    }
  ) => void
}>

