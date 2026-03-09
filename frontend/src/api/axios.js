import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'

const client = axios.create({
  baseURL: API_BASE_URL,
  timeout: 8000,
  headers: {
    'Content-Type': 'application/json',
  },
})

export async function getProjects() {
  const { data } = await client.get('/projects')

  if (Array.isArray(data?.projects)) {
    return data.projects
  }

  if (Array.isArray(data?.data)) {
    return data.data
  }

  if (Array.isArray(data)) {
    return data
  }

  return []
}

export async function sendContactMessage(payload) {
  const { data } = await client.post('/contact', payload)
  return data
}

export { API_BASE_URL }
