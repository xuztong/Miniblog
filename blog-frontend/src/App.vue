<script setup>
import { useRouter } from 'vue-router'
import { userInfo, clearUser } from './store/user'

const router = useRouter()

async function handleLogout() {
  // 简化处理：直接清除本地状态跳回首页
  // 不调后端 /tokens/ 接口，因为 httpOnly Cookie 到期会自然失效
  clearUser()
  router.push('/')
}
</script>

<template>
  <a-layout class="layout">
    <a-layout-header class="header">
      <h1 class="logo">MiniBlog</h1>
      
      <a-space size="large">
        
        <router-link to="/">首页</router-link>
        <template v-if="userInfo">
          <router-link to="/admin">管理后台</router-link>
          <a-button type="text" @click="handleLogout">
            {{ userInfo.username }} · 退出
          </a-button>
        </template>
        <template v-else>
          <router-link to="/login">登录</router-link>
      <a-space ></a-space>  
      <a-space size="4" :style="{ alignItems: 'center' }">
    <img src="../public/new.png" width="15px" height="20px">
    <span>1803192985@qq.com</span>
  </a-space>
        </template>
      </a-space>
    </a-layout-header>
    <a-layout-content class="content">
      <router-view />
    </a-layout-content>
  </a-layout>
</template>

<style  scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 24px;
  background-color: #fff;
  border-bottom: 1px solid #e5e6eb;
}

.logo {
  font-size: 18px;
  margin: 0;
}

.content {
  padding: 24px;
}

a {
  text-decoration: none;
  color: #1d2129;
}
</style>