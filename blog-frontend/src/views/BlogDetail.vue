<template>
 <div class="page">
    <router-link to="/" class="back-link">← 返回首页</router-link>

    <a-alert v-if="errorMsg" type="error" :content="errorMsg" />
    <p v-else-if="loading">加载中...</p>

    <article v-else-if="post">
      <h1 class="title">{{ post.title }}</h1>
      <p class="meta">{{ post.author }} · {{ formatDate(post.created_at) }}</p>
      <MdPreview editor-id="blog-detail" :model-value="post.content" />
    </article>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getBlogDetail } from '../api/blog'
import { MdPreview } from 'md-editor-v3'
import 'md-editor-v3/lib/preview.css'

const route = useRoute()

const post = ref(null)
const loading = ref(false)
const errorMsg = ref('')

async function fetchDetail() {
  loading.value = true
  errorMsg.value = ''
  try {
    const res = await getBlogDetail(route.params.id)
    post.value = res.data
  } catch (err) {
    errorMsg.value = '文章加载失败'
  } finally {
    loading.value = false
  }
}

function formatDate(timestamp) {
  return new Date(timestamp * 1000).toLocaleDateString()
}

onMounted(fetchDetail)

</script>

<style  scoped>
.page {
  max-width: 720px;
  margin: 40px auto;
  padding: 0 16px;
}

.back-link {
  display: inline-block;
  margin-bottom: 16px;
}

.title {
  margin: 0 0 8px;
}

.meta {
  margin: 0 0 16px;
  font-size: 13px;
  color: #86909c;
}
</style>