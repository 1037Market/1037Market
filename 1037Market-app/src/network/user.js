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

export async function hashPassword(password) {
  const encoder = new TextEncoder();
  const data = encoder.encode(password);

  // 使用SHA-256算法计算哈希值
  const hashBuffer = await crypto.subtle.digest('SHA-256', data);

  // 将Buffer转换为十六进制字符串
  const hashArray = Array.from(new Uint8Array(hashBuffer));
  const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join('');

  return hashHex;
}

