import {request} from "@/network/request";

export function publishProduct(data) {
    return request({
        url: `/api/product?user=${window.localStorage.getItem('token')}`,
        method: 'post',
        data
    })
}