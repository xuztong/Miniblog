import axios from 'axios'

// const request=axios.create({
//     baseURL: '/vblog/apps/v1',
//     timeout: 5000,
// })

const request = axios.create({
    baseURL: '/api',
    timeout: 5000,
  })

export default request