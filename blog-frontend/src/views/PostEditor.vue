<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Message } from '@arco-design/web-vue'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { getBlogDetail, createBlog, updateBlog } from '../api/blog'
import { userInfo } from '../store/user'
import { uploadImage } from '../api/image' 

const route = useRoute()
const router = useRouter()

const isNew = computed(() => route.params.id === 'new')

// 表单状态
const form = ref({
  title: '',
  author: userInfo.value?.username || '',
  summary: '',
  content: '',
})
const tagList = ref([])   // 标签数组，保存时会转成 { Vue: 'Vue' } 格式
const saving = ref(false)

// 编辑模式：拉文章数据填入表单
async function fetchPost() {
  if (isNew.value) return    // 新建模式不需要拉数据
  try {
    const res = await getBlogDetail(route.params.id)
    const post = res.data
    form.value = {
      title: post.title,
      author: post.author,
      summary: post.summary || '',
      content: post.content,
    }
    // 后端 tags 是 { Vue: 'Vue', Go: 'Go' }，只取 key 存成数组
    tagList.value = Object.keys(post.tags || {})
  } catch {
    Message.error('文章加载失败')
  }
}

async function onUploadImg(files, callback) {
  // files 是一个数组，编辑器允许同时拖入多张图片
  const urls = await Promise.all(
    files.map(async (file) => {
      const res = await uploadImage(file)
      return res.data.url
    })
  )
  // callback 把图片 URL 回传给编辑器，编辑器自动插入 ![](url)
  callback(urls)
}

// 保存
async function handleSave() {
  if (!form.value.title.trim()) {
    Message.warning('标题不能为空')
    return
  }
  if (!form.value.content.trim()) {
    Message.warning('正文不能为空')
    return
  }

  // 把 ['Vue', 'Go'] 转回 { Vue: 'Vue', Go: 'Go' }
  const tags = tagList.value.reduce((acc, t) => {
    acc[t] = t
    return acc
  }, {})

  saving.value = true
  try {
    if (isNew.value) {
      await createBlog({ ...form.value, tags })
      Message.success('文章创建成功')
    } else {
      await updateBlog(route.params.id, { ...form.value, tags })
      Message.success('文章保存成功')
    }
    router.push('/admin')
  } catch (err) {
    Message.error(err.response?.data?.message || '保存失败')
  } finally {
    saving.value = false
  }
}

onMounted(fetchPost)
</script>

<template>
  <div class="editor-page">
    <div class="editor-layout">

      <!-- 左边：Markdown 编辑器 -->
      <div class="editor-main">
        <MdEditor
          editor-id="post-editor"
          v-model="form.content"
          style="height: 600px"
          :on-upload-img="onUploadImg"
        />
      </div>

      <!-- 右边：信息栏 -->
      <aside class="editor-sidebar">
        <div class="sidebar-title">
          {{ isNew ? '新建文章' : '编辑文章' }}
        </div>

        <a-form-item label="标题">
          <a-input v-model="form.title" placeholder="请输入文章标题" />
        </a-form-item>

        <a-form-item label="作者">
          <a-input v-model="form.author" placeholder="请输入作者名" />
        </a-form-item>

        <a-form-item label="摘要">
          <a-textarea
            v-model="form.summary"
            placeholder="可选：一段简短描述"
            :auto-size="{ minRows: 3, maxRows: 5 }"
          />
        </a-form-item>

        <a-form-item label="标签">
          <a-select
            v-model="tagList"
            multiple
            allow-create
            placeholder="输入后按 Enter 添加"
            :max-tag-count="5"
          />
        </a-form-item>

        <a-button
          type="primary"
          long
          :loading="saving"
          @click="handleSave"
        >
          {{ isNew ? '发布文章' : '保存修改' }}
        </a-button>
      </aside>

    </div>
  </div>
</template>

<style scoped>
.editor-page {
  padding: 16px;
}

.editor-layout {
  display: flex;
  gap: 16px;
  align-items: flex-start;
}

.editor-main {
  flex: 1;
  min-width: 0;
}

.editor-sidebar {
  width: 260px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.sidebar-title {
  font-size: 16px;
  font-weight: 600;
  padding-bottom: 8px;
  border-bottom: 1px solid #e5e6eb;
}
</style>