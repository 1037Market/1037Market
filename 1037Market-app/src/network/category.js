import { request } from './request'

export function getCategoryData() {
  // 返回promise
  return request({
    url: '/api/product/categories'
  })
}