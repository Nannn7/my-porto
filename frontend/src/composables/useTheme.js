import { ref } from 'vue'

const STORAGE_KEY = 'my-porto-theme'
const theme = ref('light')
let initialized = false

const getPreferredTheme = () => {
  if (typeof window === 'undefined') {
    return 'light'
  }

  const saved = window.localStorage.getItem(STORAGE_KEY)
  if (saved === 'light' || saved === 'dark') {
    return saved
  }

  if (typeof window.matchMedia === 'function') {
    return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
  }

  return 'light'
}

const applyTheme = (value) => {
  if (typeof document === 'undefined') {
    return
  }

  document.documentElement.setAttribute('data-theme', value)
}

const initTheme = () => {
  if (initialized) {
    return
  }

  theme.value = getPreferredTheme()
  applyTheme(theme.value)
  initialized = true
}

const setTheme = (value) => {
  theme.value = value
  applyTheme(value)

  if (typeof window !== 'undefined') {
    window.localStorage.setItem(STORAGE_KEY, value)
  }
}

const toggleTheme = () => {
  setTheme(theme.value === 'dark' ? 'light' : 'dark')
}

export const useTheme = () => {
  initTheme()

  return {
    theme,
    setTheme,
    toggleTheme,
  }
}
