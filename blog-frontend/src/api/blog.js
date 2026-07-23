import request from "./request";

export function getBlogList(params){
    return request.get('/blogs/',{ params })
}

export function getBlogDetail(id) {
    return request.get(`/blogs/${id}`)
  }

// 切换文章状态（status: 0=草稿，1=发布）
export function toggleBlogStatus(id, status) {
    return request.post(`/blogs/${id}/status`, { status })
  }
  
  // 删除文章
  export function deleteBlog(id) {
    return request.delete(`/blogs/${id}`)
  }

  export function createBlog(data) {
    return request.post('/blogs/', data)
  }
  
  export function updateBlog(id, data) {
    return request.patch(`/blogs/${id}`, data)
  }