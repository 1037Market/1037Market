import {request} from './request'

let seed = Date.now()

//携带默认值
export function getHomeGoodsData(type = "推荐", startIndex = 0, count = 6, sign = 0) {
    if (type === '推荐') {
        return request({
            url: `/api/product/recommend?count=${count}&seed=${seed}&startIndex=${startIndex}&sign=${sign}`
        })
    }
    return request({
        url: `/api/product/category?category=${type}&count=${count}&startIndex=${startIndex + 1}&sign=${sign}`
    })
}

export function getSearchData(keyword, sign) {
    return request({
        url: `/api/product/query?keyword=${keyword}&sign=${sign}`
    })
}

export function refresh() {
    seed = Date.now()
}