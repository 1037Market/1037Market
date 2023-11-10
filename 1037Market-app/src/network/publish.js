import {request} from "@/network/request";

export function publishProduct(data) {
    return request({
        url: `/api/product?user=${window.localStorage.getItem('token')}`,
        method: 'post',
        data
    })
}

export function updateProduct(data) {
    return request({
        url: `/api/product/update?user=${window.localStorage.getItem('token')}`,
        method: 'post',
        data
    })
}

export function deleteProduct(productId) {
    return request({
        url: `/api/product?user=${window.localStorage.getItem('token')}&productId=${productId}`,
        method: 'delete',
    })
}

export function sellProduct(productId) {
    return request({
        url: `/api/product/sold?user=${window.localStorage.getItem('token')}&productId=${productId}`,
        method: 'post',
    })
}