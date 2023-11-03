import { getCart } from 'network/cart'
const actions = {
  //更新数量
  updateCart(context) {
    getCart().then(res => {
      context.commit('addCart', { count: res.data.length || 0 })
    })
  }
}

export default actions