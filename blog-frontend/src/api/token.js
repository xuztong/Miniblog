import request from './request'

export function login(data) {
  return request.post('/tokens/', data)
}

export function logout(refreshToken) {
  return request.delete('/tokens/', {
    headers: { 'X-REFRESH-TOKEN': refreshToken },
  })
}