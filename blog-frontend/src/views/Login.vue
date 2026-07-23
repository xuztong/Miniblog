<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { login } from '../api/token'
import { setUser } from '../store/user'

const router = useRouter()

const form = ref({
  username: '',
  password: '',
})
const loading = ref(false)
const errorMsg = ref('')

async function handleLogin() {
  if (!form.value.username || !form.value.password) {
    errorMsg.value = '用户名和密码不能为空'
    return
  }
  loading.value = true
  errorMsg.value = ''
  try {
    const res = await login(form.value)
    setUser(res.data)
    router.push('/admin')
  } catch (err) {
    errorMsg.value = err.response?.data?.message || '登录失败，请稍后重试'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="login-page">
    <div class="login-box">
      <h2>登个录?</h2>
      <a-alert v-if="errorMsg" type="error" :content="errorMsg" style="margin-bottom: 16px" />
      <a-form :model="form" layout="vertical">
        <a-form-item label="用户名">
          <a-input v-model="form.username" placeholder="请输入用户名" allow-clear />
        </a-form-item>
        <a-form-item label="密码">
          <a-input-password v-model="form.password" placeholder="请输入密码" />
        </a-form-item>
        <a-button type="primary" long :loading="loading" @click="handleLogin">
          登录
        </a-button>
      </a-form>
    </div>
  </div>
</template>

<style scoped>
.login-page {
  display: flex;
  justify-content: center;
  padding-top: 80px;
}

.login-box {
  width: 360px;
  padding: 32px;
  border: 1px solid #e5e6eb;
  border-radius: 8px;
}

h2 {
  margin: 0 0 24px;
}
</style>