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
    url: `/api/user/info?user=${localStorage.getItem('token')}`,
    method: 'post',
    data
  })
}

export function getUser(studentId) {
  return request({
    url: studentId ? `/api/user/info?studentId=${studentId}`:
        `/api/user/info?studentId=${window.localStorage.getItem('studentId')}`,
    method: 'get'
  })
}

export function getUserPublishedProductIds(studentId) {
  return request({
    url: studentId ? `/api/product/student?studentId=${studentId}`:
        `/api/product/student?studentId=${window.localStorage.getItem('studentId')}`,
    method: 'get'
  })
}

