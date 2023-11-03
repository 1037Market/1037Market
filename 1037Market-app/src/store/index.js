import { createStore } from 'vuex'
import mutations from './mutations'
import actions from './actions'
import getters from './getters'


const state = {
  user: {
    isLogin: window.localStorage.getItem('token') ? true : false
  },
  cartCount: 0
}

export default createStore({
  state,
  mutations,
  actions,
  getters,

})
