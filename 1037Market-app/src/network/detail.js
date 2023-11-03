import { request } from './request';

export function getDetail(id) {
  return request({
    url: `/api/product/get?productId=${id}`,
  })
}
