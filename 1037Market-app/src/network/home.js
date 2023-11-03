import { request } from './request'

export function getHomeAllData() {
  // 返回promise
  return request({
    url: '/api/index'
  })
}

//携带默认值
export function getHomeGoodsData(type = "recommend", count = 10) {
  if(type === 'recommend'){
    return request({
      url: `/api/product/recommend?count=${count}`
    })
  }
  return request({
    url: `/api/product/category?category=${type}&count=${count}`
  })
}