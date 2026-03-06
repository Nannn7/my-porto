import axios from 'axios'

const client = axios.create({
  baseURL: 'https://localhost:8080/api',
})

export async function getProjects() {
  const { data } = await client.get('/projects')
  return data.data || []
}
