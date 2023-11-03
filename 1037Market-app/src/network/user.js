import { request } from './request';

export function register(data) {
  console.log(data)
  return request({
    url: `/api/user/register`,
    method: 'post',
    data
  })
}

export function login(data) {
  console.log("login", data)
  return request({
    url: '/api/user/login',
    method: 'post',
    data
  })
}

export function getCaptcha(data){
  return request({
    url: `/api/user/register/email?studentId=${data.studentId}`,
    method: 'post'
  })
}

export function logout() {
  return request({
    url: '/api/auth/logout',
    method: 'post'
  })
}

export function updateUser(data){
  return request({
    url: '/api/user/info',
    method: 'post',
    data
  })
}

export function getUser() {
  return request({
    url: '/api/user/info',
    method: 'get'
  })
}

