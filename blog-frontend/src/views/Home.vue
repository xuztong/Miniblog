<script setup>
import { ref, onMounted } from 'vue'
import { getBlogList } from '../api/blog'
import { useRouter } from 'vue-router'

const router = useRouter()
const posts = ref([])
const loading = ref(false)
const errorMsg = ref('')
const total = ref(0)
const pageNumber = ref(1)
const pageSize = ref(5)
const keywords = ref('')

function goToDetail(id) {
  router.push(`/blog/${id}`)
}

async function fetchPosts() {
  loading.value = true
  errorMsg.value = ''
  try {
    const res = await getBlogList({
      page_number: pageNumber.value,
      page_size: pageSize.value,
      keywords: keywords.value,
    })
    posts.value = res.data.items || []
    total.value = res.data.total || 0
  } catch (err) {
    errorMsg.value = '文章加载失败，请稍后再试'
  } finally {
    loading.value = false
  }
}

function handlePageChange(page) {
  pageNumber.value = page
  fetchPosts()
}



function handleSearch(value) {
  keywords.value = value
  pageNumber.value = 1
  fetchPosts()
}

onMounted(fetchPosts)

</script>

<template>
  <div class="page">
    <header class="page-header">
 <img src="../assets/logo.gif" alt="" width="70" height="50">
  <a-input-search
    placeholder="搜索文章标题"
    allow-clear
    search-button
    style="width: 500px"
    @search="handleSearch"
  />
</header>

    <br>
    <br>
    <a-alert v-if="errorMsg" type="error" :content="errorMsg" />
    <p v-if="loading" class="loading-tip">数据加载中...</p>
    <ul v-else class="post-list">
      <li v-for="post in posts" :key="post.id" class="post-card"  @click="goToDetail(post.id)" >
        <h2 class="=post-title">{{ post.title }}</h2>
        <p class="post-meta">作者：{{ post.author }}
          <a-tag v-if="post.status === 0" color="orange" size="small">草稿</a-tag>
        </p>
        <p class="post-summary">{{ post.summary }}</p>
        <p >{{ post.tag }}</p>
      </li>
    </ul>
  </div>
  
  <div class="pagination">
  <a-pagination
    :total="total"
    :current="pageNumber"
    :page-size="pageSize"
    show-total
    @change="handlePageChange"
  />
</div>
</template>
<style scoped>
.page{
  max-width: 720px;
  margin: 40px auto;
  padding: 0 16px;
  font-family: system-ui, -apple-system, 'PingFang SC', sans-serif;
  color: #1d2129;
}
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.btn {
  padding: 8px 16px;
  font-size: 14px;
  color: #fff;
  background-color: #165dff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.btn:hover {
  background-color: #4080ff;
}

.loading-tip {
  color: #86909c;
}

.post-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.post-card {
  padding: 16px;
  border: 1px solid #e5e6eb;
  border-radius: 8px;
}

.post-title {
  margin: 0 0 8px;
  font-size: 18px;
}

.post-meta {
  margin: 0 0 8px;
  font-size: 13px;
  color: #86909c;
}

.post-summary {
  margin: 0;
  color: #4e5969;
  line-height: 1.6;
}

.pagination {
  display: flex;
  justify-content: center;
  margin-top: 24px;
}

.post-card:hover {
  border-color: #165dff;
}
</style>