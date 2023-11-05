import {request} from "./request";

export function getUserCommentIds(studentId) {
    return request({
        url: studentId ? `/api/comment?studentId=${studentId}`:
            `/api/comment?studentId=${window.localStorage.getItem('studentId')}`,
        method: 'get'
    })
}

export function getCommentDetail(commentId) {
    return request({
        url: `/api/comment/get?commentId=${commentId}`,
        method: 'get'
    })
}