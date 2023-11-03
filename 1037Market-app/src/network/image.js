import {request} from "@/network/request";

export function uploadImage(data) {
    return request({
        url: `/api/image?user=${window.localStorage.getItem('token')}`,
        method: 'post',
        data
    })
}

