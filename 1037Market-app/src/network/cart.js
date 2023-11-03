import { request } from './request';


//添加购物车
export function addCart(productId) {
  return request({
    url: `/api/subscribe?user=${window.localStorage.getItem('token')}&productId=${productId}`,
    method: 'post'
  })
}

//  获取购物车列表
export function getCart(data = '') {
  return request({
    // api/carts?include=goods
    url: `/api/subscribe?userId=${window.localStorage.getItem('studentId')}`,
    method: 'get'
  })
}


//  删除购物车
export function deleteCartItem(productId) {
  return request({
    url: `/api/carts/user=${window.localStorage.getItem('token')}&productId=${productId}`,
    method: 'delete'
  })
}