import { request } from './request';


//添加购物车
export function addCart(productId) {
  return request({
    url: `/api/subscribe?productId=${productId}`,
    method: 'post'
  })
}

// 修改购物车数量， data = {num:1}
export function modifyCart(id, data) {
  return request({
    url: `/api/carts/${id}`,
    method: 'put',
    data
  })
}

//  获取购物车列表
export function getCart(data = '') {
  return request({
    // api/carts?include=goods
    url: '/api/carts?' + data,
  })
}


//  删除购物车
export function deleteCartItem(id) {
  return request({
    url: `/api/carts/${id}`,
    method: 'delete'
  })
}