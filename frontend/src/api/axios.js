import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'
export const ADMIN_TOKEN_STORAGE_KEY = 'myporto_admin_token'

const client = axios.create({
  baseURL: API_BASE_URL,
  timeout: 8000,
  headers: {
    'Content-Type': 'application/json',
  },
})

client.interceptors.request.use((config) => {
  if (typeof window === 'undefined') {
    return config
  }

  const token = window.localStorage.getItem(ADMIN_TOKEN_STORAGE_KEY)
  const isAdminEndpoint = config.url?.startsWith('/admin')
  if (token && isAdminEndpoint) {
    config.headers = config.headers || {}
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
})

function unwrapData(payload) {
  if (payload && typeof payload === 'object') {
    if ('data' in payload) {
      return payload.data
    }

    return payload
  }

  return null
}

export async function getProjects() {
  const { data } = await client.get('/projects')
  const responseData = unwrapData(data)

  if (Array.isArray(responseData?.projects)) {
    return responseData.projects
  }

  if (Array.isArray(responseData?.data)) {
    return responseData.data
  }

  if (Array.isArray(responseData)) {
    return responseData
  }

  return []
}

export async function loginAdmin(payload) {
  const { data } = await client.post('/admin/login', payload)
  return unwrapData(data)
}

export async function getAdminProjects() {
  const { data } = await client.get('/admin/projects')
  const responseData = unwrapData(data)
  return Array.isArray(responseData) ? responseData : []
}

export async function createAdminProject(payload) {
  const { data } = await client.post('/admin/projects', payload)
  return unwrapData(data)
}

export async function updateAdminProject(id, payload) {
  const { data } = await client.put(`/admin/projects/${id}`, payload)
  return unwrapData(data)
}

export async function deleteAdminProject(id) {
  const { data } = await client.delete(`/admin/projects/${id}`)
  return unwrapData(data)
}

export async function sendContactMessage(payload) {
  const { data } = await client.post('/contact', payload)
  return unwrapData(data)
}

export { API_BASE_URL }
