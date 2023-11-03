import { request } from './request'

export function getCategoryData() {
  // 返回promise
  return request({
    url: '/api/product/categories'
  })
}

export function getCategoryGoods(category, count=10) {
  return request({
    url: `/api/goods?category=${category}&count=${count}`,
  })
}
