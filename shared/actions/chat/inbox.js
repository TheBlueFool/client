// @flow
import * as ChatTypes from '../../constants/types/flow-types-chat'
import * as Constants from '../../constants/chat'
import {List, Map} from 'immutable'
import {TlfKeysTLFIdentifyBehavior} from '../../constants/types/flow-types'
import {call, put, select, race} from 'redux-saga/effects'
import {delay} from 'redux-saga'
import {loadMoreMessages, loadedInbox} from './creators'
import {parseFolderNameToUsers} from '../../util/kbfs'
import {requestIdleCallback} from '../../util/idle-callback'
import {singleFixedChannelConfig, takeFromChannelMap} from '../../util/saga'
import {unsafeUnwrap} from '../../constants/types/more'
import {usernameSelector} from '../../constants/selectors'

import type {FinalizedState, InboxState, LoadInbox, ConversationIDKey, MetaData} from '../../constants/chat'
import type {GetInboxLocalRes, ConversationLocal} from '../../constants/types/flow-types-chat'
import type {SagaGenerator, ChannelMap} from '../../constants/types/saga'
import type {TypedState} from '../../constants/reducer'

const {CommonConversationStatus, CommonTLFVisibility, CommonTopicType, LocalConversationErrorType, localGetInboxNonblockLocalRpcChannelMap, LocalMessageUnboxedState} = ChatTypes
const {InboxStateRecord, conversationIDToKey, getSelectedConversation, GetInboxAndUnbox, keyToConversationID, makeSnippet} = Constants

const _metaDataSelector = (state: TypedState) => state.chat.get('metaData')
const followingSelector = (state: TypedState) => state.config.following

let _loadedInboxOnce = false
function * onLoadInboxMaybeOnce (action: LoadInbox): SagaGenerator<any, any> {
  if (!_loadedInboxOnce || action.payload.force) {
    _loadedInboxOnce = true
    yield call(onLoadInbox)
  }
}

function * onLoadInbox (): SagaGenerator<any, any> {
  const channelConfig = singleFixedChannelConfig([
    'chat.1.chatUi.chatInboxUnverified',
    'chat.1.chatUi.chatInboxConversation',
    'chat.1.chatUi.chatInboxFailed',
    'finished',
  ])

  const loadInboxChanMap: ChannelMap<any> = localGetInboxNonblockLocalRpcChannelMap(channelConfig, {
    param: {
      identifyBehavior: TlfKeysTLFIdentifyBehavior.chatGui,
      query: {
        computeActiveList: true,
        readOnly: false,
        status: Object.keys(CommonConversationStatus).filter(k => !['ignored', 'blocked'].includes(k)).map(k => CommonConversationStatus[k]),
        tlfVisibility: CommonTLFVisibility.private,
        topicType: CommonTopicType.chat,
        unreadOnly: false,
      },
    },
  })

  const chatInboxUnverified = yield takeFromChannelMap(loadInboxChanMap, 'chat.1.chatUi.chatInboxUnverified')

  if (!chatInboxUnverified) {
    throw new Error("Can't load inbox")
  }

  const metaData = ((yield select(_metaDataSelector)): any)
  const inbox: GetInboxLocalRes = chatInboxUnverified.params.inbox
  const author = yield select(usernameSelector)
  const following = yield select(followingSelector)
  const conversations: List<InboxState> = _inboxToConversations(inbox, author, following || {}, metaData)
  const finalizedState: FinalizedState = _inboxToFinalized(inbox)

  yield put(loadedInbox(conversations))
  if (finalizedState.count()) {
    yield put(({type: 'chat:updateFinalizedState', payload: {finalizedState}}: Constants.UpdateFinalizedState))
  }

  chatInboxUnverified.response.result()

  let finishedCalled = false
  while (!finishedCalled) {
    const incoming: {[key: string]: any} = yield race({
      chatInboxConversation: takeFromChannelMap(loadInboxChanMap, 'chat.1.chatUi.chatInboxConversation'),
      chatInboxFailed: takeFromChannelMap(loadInboxChanMap, 'chat.1.chatUi.chatInboxFailed'),
      finished: takeFromChannelMap(loadInboxChanMap, 'finished'),
      timeout: call(delay, 30000),
    })

    if (incoming.chatInboxConversation) {
      requestIdleCallback(() => {
        incoming.chatInboxConversation.response.result()
      }, {timeout: 100})

      yield call(delay, 1)
      let conversation: ?InboxState = _inboxConversationToInboxState(incoming.chatInboxConversation.params.conv, author, following || {}, metaData)

      // TODO this is ugly, ideally we should just call updateInbox here
      const conv = incoming.chatInboxConversation.params.conv
      const supersedesState: Constants.SupersedesState = _inboxConversationLocalToSupersedesState(conv)
      const supersededByState: Constants.SupersededByState = _inboxConversationLocalToSupersededByState(conv)
      const finalizedState: Constants.FinalizedState = _conversationLocalToFinalized(conv)

      if (supersedesState.count()) {
        yield put(({type: 'chat:updateSupersedesState', payload: {supersedesState}}: Constants.UpdateSupersedesState))
      }
      if (supersededByState.count()) {
        yield put(({type: 'chat:updateSupersededByState', payload: {supersededByState}}: Constants.UpdateSupersededByState))
      }
      if (finalizedState.count()) {
        yield put(({type: 'chat:updateFinalizedState', payload: {finalizedState}}: Constants.UpdateFinalizedState))
      }

      if (conversation) {
        yield put(({type: 'chat:updateInbox', payload: {conversation}}: Constants.UpdateInbox))
        const selectedConversation = yield select(getSelectedConversation)
        if (selectedConversation === conversation.get('conversationIDKey')) {
          // load validated selected
          yield put(loadMoreMessages(selectedConversation, false))
        }
      }
      // find it
    } else if (incoming.chatInboxFailed) {
      console.log('ignoring chatInboxFailed', incoming.chatInboxFailed)
      requestIdleCallback(() => {
        incoming.chatInboxFailed.response.result()
      }, {timeout: 100})

      yield call(delay, 1)
      const error = incoming.chatInboxFailed.params.error
      const conversationIDKey = conversationIDToKey(incoming.chatInboxFailed.params.convID)
      const conversation = new InboxStateRecord({
        conversationIDKey,
        info: null,
        isEmpty: false,
        participants: List([].concat(error.rekeyInfo ? error.rekeyInfo.writerNames : [], error.rekeyInfo ? error.rekeyInfo.readerNames : []).filter(Boolean)),
        snippet: null,
        status: 'unfiled',
        time: error.remoteConv.readerInfo.mtime,
        validated: true,
      })
      yield put(({type: 'chat:updateInbox', payload: {conversation}}: Constants.UpdateInbox))

      switch (error.typ) {
        case LocalConversationErrorType.selfrekeyneeded: {
          yield put({type: 'chat:updateInboxRekeySelf', payload: {conversationIDKey}})
          break
        }
        case LocalConversationErrorType.otherrekeyneeded: {
          const rekeyers = error.rekeyInfo.rekeyers
          yield put({type: 'chat:updateInboxRekeyOthers', payload: {conversationIDKey, rekeyers}})
          break
        }
        default:
          if (__DEV__) {
            console.warn('Inbox error:', error)
          }
      }
    } else if (incoming.finished) {
      finishedCalled = true
      yield put({type: 'chat:updateInboxComplete', payload: undefined})
      break
    } else if (incoming.timeout) {
      console.warn('Inbox loading timed out')
      yield put({type: 'chat:updateInboxComplete', payload: undefined})
      break
    }
  }
}

function _inboxConversationLocalToSupersededByState (convo: ?ConversationLocal): Constants.SupersededByState {
  if (!convo || !convo.info || !convo.info.id || !convo.supersededBy) {
    return Map()
  }

  const conversationIDKey = conversationIDToKey(convo.info.id)
  const supersededBy = _toSupersedeInfo(conversationIDKey, (convo.supersededBy || []))
  return supersededBy ? Map({[conversationIDKey]: supersededBy}) : Map()
}

function _conversationLocalToFinalized (convo: ?ConversationLocal): FinalizedState {
  if (convo && convo.info.id && convo.info.finalizeInfo) {
    return Map({
      [conversationIDToKey(convo.info.id)]: convo.info.finalizeInfo,
    })
  }
  return Map()
}

function * getInboxAndUnbox ({payload: {conversationIDKey}}: GetInboxAndUnbox): SagaGenerator<any, any> {
  const param: ChatTypes.localGetInboxAndUnboxLocalRpcParam = {
    identifyBehavior: TlfKeysTLFIdentifyBehavior.chatGui,
    query: {
      computeActiveList: true,
      convIDs: [keyToConversationID(conversationIDKey)],
      readOnly: true,
      tlfVisibility: CommonTLFVisibility.private,
      topicType: CommonTopicType.chat,
      unreadOnly: false,
    },
  }

  const result: ChatTypes.GetInboxAndUnboxLocalRes = yield call(ChatTypes.localGetInboxAndUnboxLocalRpcPromise, {param})
  const {conversations} = result
  if (conversations && conversations[0]) {
    yield call(updateInbox, conversations[0])
    // inbox loaded so rekeyInfo is now clear
    yield put({payload: {conversationIDKey}, type: 'chat:clearRekey'})
  }
  // TODO maybe we get failures and we should update rekeyinfo? unclear...
}

function * updateInbox (conv: ?ConversationLocal): SagaGenerator<any, any> {
  const inboxState = _inboxConversationToInboxState(conv)
  const supersedesState: Constants.SupersedesState = _inboxConversationLocalToSupersedesState(conv)
  const supersededByState: Constants.SupersededByState = _inboxConversationLocalToSupersededByState(conv)
  const finalizedState: Constants.FinalizedState = _conversationLocalToFinalized(conv)

  if (supersedesState.count()) {
    yield put(({type: 'chat:updateSupersedesState', payload: {supersedesState}}: Constants.UpdateSupersedesState))
  }
  if (supersededByState.count()) {
    yield put(({type: 'chat:updateSupersededByState', payload: {supersededByState}}: Constants.UpdateSupersededByState))
  }
  if (finalizedState.count()) {
    yield put(({type: 'chat:updateFinalizedState', payload: {finalizedState}}: Constants.UpdateFinalizedState))
  }

  if (inboxState) {
    yield put(({
      payload: {conversation: inboxState},
      type: 'chat:updateInbox',
    }: Constants.UpdateInbox))
  }
}

function _inboxConversationLocalToSupersedesState (convo: ?ChatTypes.ConversationLocal): Constants.SupersedesState {
  // TODO deep supersedes checking
  if (!convo || !convo.info || !convo.info.id || !convo.supersedes) {
    return Map()
  }

  const conversationIDKey = conversationIDToKey(convo.info.id)
  const supersedes = _toSupersedeInfo(conversationIDKey, (convo.supersedes || []))
  return supersedes ? Map({[conversationIDKey]: supersedes}) : Map()
}

function _toSupersedeInfo (conversationIDKey: ConversationIDKey, supersedeData: Array<ChatTypes.ConversationMetadata>): ?Constants.SupersedeInfo {
  const parsed = supersedeData
    .filter(md => md.idTriple.topicType === CommonTopicType.chat && md.finalizeInfo)
    .map(md => ({
      conversationIDKey: conversationIDToKey(md.conversationID),
      finalizeInfo: unsafeUnwrap(md && md.finalizeInfo),
    }))
  return parsed.length ? parsed[0] : null
}
function _inboxConversationToInboxState (convo: ?ConversationLocal): ?InboxState {
  if (!convo || !convo.info || !convo.info.id) {
    return null
  }

  if (convo.info.visibility !== ChatTypes.CommonTLFVisibility.private) {
    return null
  }

  // We don't support mixed reader/writers
  if (convo.info.tlfName.includes('#')) {
    return null
  }

  const conversationIDKey = conversationIDToKey(convo.info.id)
  let snippet
  let time

  (convo.maxMessages || []).some(message => {
    if (message.state === LocalMessageUnboxedState.valid && message.valid && convo && convo.readerInfo) {
      time = message.valid.serverHeader.ctime || convo.readerInfo.mtime
      snippet = makeSnippet(message.valid.messageBody)
      return !!snippet
    }
    return false
  })

  const participants = List(convo.info.writerNames || [])
  const infoStatus = convo.info ? convo.info.status : 0
  // Go backwards from the value in CommonConversationStatus to its key.
  const status = Constants.ConversationStatusByEnum[infoStatus]

  return new InboxStateRecord({
    conversationIDKey,
    info: convo.info,
    isEmpty: convo.isEmpty,
    participants,
    snippet,
    status,
    time,
    validated: true,
  })
}

function _inboxToConversations (inbox: GetInboxLocalRes, author: ?string, following: {[key: string]: boolean}, metaData: MetaData): List<InboxState> {
  return List((inbox.conversationsUnverified || []).map(convoUnverified => {
    const msgMax = convoUnverified.maxMsgSummaries && convoUnverified.maxMsgSummaries.length && convoUnverified.maxMsgSummaries[0]

    if (!msgMax) {
      return null
    }

    const participants = List(parseFolderNameToUsers(author, msgMax.tlfName).map(ul => ul.username))
    const statusEnum = convoUnverified.metadata.status || 0
    const status = Constants.ConversationStatusByEnum[statusEnum]

    return new InboxStateRecord({
      conversationIDKey: conversationIDToKey(convoUnverified.metadata.conversationID),
      info: null,
      participants,
      snippet: ' ',
      status,
      time: convoUnverified.readerInfo && convoUnverified.readerInfo.mtime,
      validated: false,
    })
  }).filter(Boolean))
}

function _inboxToFinalized (inbox: GetInboxLocalRes): FinalizedState {
  return Map((inbox.conversationsUnverified || []).map(convoUnverified => [
    conversationIDToKey(convoUnverified.metadata.conversationID),
    convoUnverified.metadata.finalizeInfo,
  ]))
}

export {
  onLoadInboxMaybeOnce,
  onLoadInbox,
  getInboxAndUnbox,
  updateInbox,
}
