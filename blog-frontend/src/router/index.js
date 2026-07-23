import { createRouter, createWebHistory } from 'vue-router'
import { userInfo } from '../store/user'
import Home from '../views/Home.vue'
import BlogDetail from '../views/BlogDetail.vue'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'
import PostEditor from '../views/PostEditor.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/blog/:id', component: BlogDetail },
  { path: '/login', component: Login },
  { path: '/admin', component: Admin },
  { path: '/admin/post/:id', component: PostEditor },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// 路由守卫：每次跳转之前都会执行这个函数
router.beforeEach((to, from, next) => {
  const needsLogin = to.path.startsWith('/admin')
  if (needsLogin && !userInfo.value) {
    next('/login')
  } else {
    next()
  }
})

export default router