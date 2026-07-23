import { ref } from 'vue'

// 启动时从 localStorage 里恢复上次登录的用户信息
// 如果没登录过，JSON.parse(null) 返回 null，ref(null) 表示"未登录"
export const userInfo = ref(JSON.parse(localStorage.getItem('userInfo')))

export function setUser(info) {
  userInfo.value = info
  localStorage.setItem('userInfo', JSON.stringify(info))
}

export function clearUser() {
  userInfo.value = null
  localStorage.removeItem('userInfo')
}