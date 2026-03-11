import { defineStore } from 'pinia'
import { loginAdmin, ADMIN_TOKEN_STORAGE_KEY } from '../api/axios'

const ADMIN_USERNAME_STORAGE_KEY = 'myporto_admin_username'
const ADMIN_EXPIRES_AT_STORAGE_KEY = 'myporto_admin_expires_at'

function readFromStorage(key) {
  if (typeof window === 'undefined') {
    return ''
  }

  return window.localStorage.getItem(key) || ''
}

function saveToStorage(key, value) {
  if (typeof window === 'undefined') {
    return
  }

  if (!value) {
    window.localStorage.removeItem(key)
    return
  }

  window.localStorage.setItem(key, value)
}

function extractAuthResponse(payload) {
  if (!payload || typeof payload !== 'object') {
    return { token: '', expiresAt: '', username: '' }
  }

  const directToken = payload.token || payload.access_token || ''
  const directExpires = payload.expires_at || payload.expiresAt || ''
  const directUsername = payload.username || ''

  if (directToken || directExpires || directUsername) {
    return {
      token: directToken,
      expiresAt: directExpires,
      username: directUsername,
    }
  }

  const nested = payload.data
  if (nested && typeof nested === 'object') {
    return extractAuthResponse(nested)
  }

  return { token: '', expiresAt: '', username: '' }
}

export function hasAdminSession() {
  const token = readFromStorage(ADMIN_TOKEN_STORAGE_KEY)
  const expiresAt = readFromStorage(ADMIN_EXPIRES_AT_STORAGE_KEY)

  if (!token || !expiresAt) {
    return false
  }

  const expiresTime = Date.parse(expiresAt)
  if (!Number.isFinite(expiresTime)) {
    return false
  }

  return expiresTime > Date.now()
}

export const useAdminStore = defineStore('admin', {
  state: () => ({
    token: readFromStorage(ADMIN_TOKEN_STORAGE_KEY),
    username: readFromStorage(ADMIN_USERNAME_STORAGE_KEY),
    expiresAt: readFromStorage(ADMIN_EXPIRES_AT_STORAGE_KEY),
    loading: false,
    error: '',
  }),
  getters: {
    isAuthenticated: (state) => {
      if (!state.token || !state.expiresAt) {
        return false
      }

      const expiresTime = Date.parse(state.expiresAt)
      return Number.isFinite(expiresTime) && expiresTime > Date.now()
    },
  },
  actions: {
    async login(username, password) {
      this.loading = true
      this.error = ''

      try {
        const response = await loginAdmin({ username, password })
        const parsed = extractAuthResponse(response)
        const token = parsed.token
        const expiresAt = parsed.expiresAt
        const serverUsername = parsed.username || username

        if (!token || !expiresAt) {
          throw new Error(
            'Login berhasil tapi token tidak ditemukan. Restart backend terbaru lalu coba lagi.'
          )
        }

        this.token = token
        this.username = serverUsername
        this.expiresAt = expiresAt

        saveToStorage(ADMIN_TOKEN_STORAGE_KEY, token)
        saveToStorage(ADMIN_USERNAME_STORAGE_KEY, serverUsername)
        saveToStorage(ADMIN_EXPIRES_AT_STORAGE_KEY, expiresAt)
      } catch (error) {
        this.token = ''
        this.username = ''
        this.expiresAt = ''
        this.error = error?.response?.data?.message || error?.message || 'Login failed'

        saveToStorage(ADMIN_TOKEN_STORAGE_KEY, '')
        saveToStorage(ADMIN_USERNAME_STORAGE_KEY, '')
        saveToStorage(ADMIN_EXPIRES_AT_STORAGE_KEY, '')

        throw error
      } finally {
        this.loading = false
      }
    },
    logout() {
      this.token = ''
      this.username = ''
      this.expiresAt = ''
      this.error = ''

      saveToStorage(ADMIN_TOKEN_STORAGE_KEY, '')
      saveToStorage(ADMIN_USERNAME_STORAGE_KEY, '')
      saveToStorage(ADMIN_EXPIRES_AT_STORAGE_KEY, '')
    },
  },
})
