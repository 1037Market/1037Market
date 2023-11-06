import {request} from "./request";

export function getRoomIds() {
    return request({
        url: `/api/chat/sessions?studentId=${window.localStorage.getItem('studentId')}`,
        method: 'get'
    })
}

export function getRoomInfo(roomId) {
    return request({
        url: `/api/chat/userInfos?sessionId=${roomId}`,
        method: 'get'
    })
}

export function createRoom(toId) {
    return request({
        url: `/api/chat/session?studentId1=${window.localStorage.getItem('studentId')}&studentId2=${toId}`,
        method: 'get'
    })
}

export function getNewestMessageIn(roomId) {
    return request({
        url: `/api/chat/message?sessionId=${roomId}`,
        method: 'get'
    })
}

export function getMessageContent(messageId) {
    return request({
        url: `/api/chat/content?messageId=${messageId}`,
        method: 'get'
    })
}

export function getNMessages(roomId, idx, cnt) {
    return request({
        url: `/api/chat/messages?sessionId=${roomId}&index=${idx}&count=${cnt}`,
        method: 'get'
    })
}

export function sendMessage(data) {
    return request({
        url: `/api/chat/send?user=${window.localStorage.getItem('token')}`,
        method: 'post',
        data
    })
}