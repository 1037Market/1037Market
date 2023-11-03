//写通用的网络请求
import axios from 'axios'
import { Notify } from "vant";
import router from '../router'
import Cookies from "js-cookie";

export function request(config) {
  //创建一个实例
  const instance = axios.create({
    baseURL: 'http://franky.pro:7301',
    timeout: 5000
  })

  //拦截器
  //请求拦截
  instance.interceptors.request.use((config) => {
    // 如果有一个接口需要认证才可以访问，就在这统一设置
    const token = window.localStorage.getItem('token');

    if (token) {
      config.headers.Authorization = 'Bearer ' + token;
    }
    //直接放行
    return config;
  }, (err) => {
    //一般请求不会发生错误，不写
  })

  //响应拦截,有两个参数，一个响应数据方法，一个错误方法
  instance.interceptors.response.use((res) => {
    //正常响应直接放行
    //有data返回data，没有返回res
    return res.data ? res.data : res;
  }, (err) => {
    //如果有需要授权才可以访问的接口， 统一去login授权
    if (err.response.status == 401) {
      Notify.fail('请先登录')
      router.push({ path: '/login' })
    }
    //如果有错误，显示错误信息
    Notify(err.response.data.errors[Object.keys(err.response.data.errors)[0]][0])
  })

  //要返回实例
  return instance(config)
}