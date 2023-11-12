import { request } from './request';

export function getDetail(id) {
  return request({
    url: `/api/product/get?user=${window.localStorage.getItem('token')}&productId=${id}`,
  })
}
