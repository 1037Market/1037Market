import { request } from './request'

let seed = Date.now()

export function getHomeAllData() {
  // 返回promise
  return request({
    url: '/api/index'
  })
}

//携带默认值
export function getHomeGoodsData(type = "推荐", startIndex = 0, count = 4) {
  if(type === '推荐'){
    return request({
      url: `/api/product/recommend?count=${count}&seed=${seed}&startIndex=${startIndex}`
    })
  }
  return request({
    url: `/api/product/category?category=${type}&count=${count}`
  })
}

export function getSearchData(keyword){
  return request({
    url: `/api/product/query?keyword=${keyword}`
  })
}

export function refresh(){
  seed = Date.now()
}