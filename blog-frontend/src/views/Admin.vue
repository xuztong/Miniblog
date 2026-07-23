<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Message } from '@arco-design/web-vue'
import { getBlogList, toggleBlogStatus, deleteBlog } from '../api/blog'

const router = useRouter()
const posts = ref([])
const loading = ref(false)
const togglingId = ref(null)   // 记录"正在切换状态"的文章 id

// ---- 表格的列定义 ----
const columns = [
  { title: '标题', dataIndex: 'title' },
  { title: '作者', dataIndex: 'author', width: 120 },
  { title: '状态', dataIndex: 'status', width: 100, slotName: 'status' },
  { title: '创建时间', dataIndex: 'created_at', width: 130, slotName: 'time' },
  { title: '操作', width: 160, slotName: 'actions' },
]

// ---- 拉取文章列表 ----
async function fetchPosts() {
  loading.value = true
  try {
    const res = await getBlogList({ page_number: 1, page_size: 100 })
    posts.value = res.data.items || []
  } catch {
    Message.error('文章列表加载失败')
  } finally {
    loading.value = false
  }
}

// ---- 切换发布/草稿 ----
async function handleToggleStatus(record, checked) {
  togglingId.value = record.id
  try {
    await toggleBlogStatus(record.id, checked ? 1 : 0)
    record.status = checked ? 1 : 0     // 直接改掉 posts 里这条记录的 status
    Message.success(checked ? '已发布' : '已转为草稿')
  } catch {
    Message.error('操作失败，请重试')
  } finally {
    togglingId.value = null
  }
}

// ---- 删除 ----
async function handleDelete(id) {
  try {
    await deleteBlog(id)
    posts.value = posts.value.filter((p) => p.id !== id)
    Message.success('已删除')
  } catch {
    Message.error('删除失败，请重试')
  }
}

function formatDate(timestamp) {
  return new Date(timestamp * 1000).toLocaleDateString()
}

onMounted(fetchPosts)
</script>

<template>
  <div>
    <div class="admin-header">
      <h2>文章管理</h2>
      <a-button type="primary" @click="router.push('/admin/post/new')">
        写文章
      </a-button>
    </div>

    <a-table
      row-key="id"
      :data="posts"
      :columns="columns"
      :loading="loading"
      :pagination="false"
    >
      <!-- 状态列：用开关显示，可以直接拨动 -->
      <template #status="{ record }">
        <a-switch
          :model-value="record.status === 1"
          :loading="togglingId === record.id"
          checked-text="已发布"
          unchecked-text="草稿"
          @change="(val) => handleToggleStatus(record, val)"
        />
      </template>

      <!-- 时间列：把时间戳格式化成日期 -->
      <template #time="{ record }">
        {{ formatDate(record.created_at) }}
      </template>

      <!-- 操作列：编辑 + 删除 -->
      <template #actions="{ record }">
        <a-space>
          <a-button
            type="text"
            size="small"
            @click="router.push(`/admin/post/${record.id}`)"
          >
            编辑
          </a-button>
          <a-popconfirm
            content="确认删除这篇文章？删除后无法恢复。"
            @ok="handleDelete(record.id)"
          >
            <a-button type="text" status="danger" size="small">
              删除
            </a-button>
          </a-popconfirm>
        </a-space>
      </template>
    </a-table>
  </div>
</template>

<style scoped>
.admin-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

h2 { margin: 0; }
</style>