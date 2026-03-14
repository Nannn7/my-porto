<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useTheme } from '../composables/useTheme'
import { useAdminStore } from '../store/admin'

const navItems = [
  { name: 'Home', to: '/' },
  { name: 'About', to: '/about' },
  { name: 'Highlights', to: '/projects' },
  { name: 'Contact', to: '/contact' },
]

const { theme, toggleTheme } = useTheme()
const isDark = computed(() => theme.value === 'dark')
const router = useRouter()
const adminStore = useAdminStore()
const isAdminAuthenticated = computed(() => adminStore.isAuthenticated)
const adminRoute = computed(() => (isAdminAuthenticated.value ? '/admin/projects' : '/admin/login'))

const logoutAdmin = async () => {
  adminStore.logout()
  await router.push('/')
}
</script>

<template>
  <header class="sticky top-0 z-20 border-b border-border bg-bg-elevated/90 backdrop-blur">
    <div
      class="reveal flex min-h-20 w-full items-center justify-between gap-4 px-4 py-3 md:px-8 lg:px-12 max-[1080px]:flex-wrap max-[1080px]:justify-center max-[1080px]:pb-[0.9rem] max-[1080px]:text-center"
    >
      <RouterLink to="/" class="grid gap-0.5 max-[1080px]:w-full" aria-label="Go to home page">
        <p class="font-display text-[0.96rem] font-bold text-heading">Nanda Surya Diffa</p>
        <p class="text-[0.76rem] text-muted max-[1080px]:hidden">Web Developer, Android Developer, API & Middleware</p>
      </RouterLink>

      <nav class="flex flex-wrap items-center justify-center gap-2 max-[1080px]:w-full" aria-label="Main navigation">
        <RouterLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          class="rounded-full px-3 py-2 text-[0.9rem] font-semibold text-muted hover:-translate-y-px hover:bg-brand-soft hover:text-ink"
          active-class="bg-brand-soft text-ink"
        >
          {{ item.name }}
        </RouterLink>
      </nav>

      <div class="flex items-center gap-2 max-[1080px]:w-full max-[1080px]:justify-center">
        <!-- <RouterLink
          :to="adminRoute"
          class="rounded-full border border-border bg-surface px-3 py-2 text-[0.8rem] font-bold text-ink hover:border-brand"
        >
          {{ isAdminAuthenticated ? 'Dashboard' : 'Admin' }}
        </RouterLink>
        <button
          v-if="isAdminAuthenticated"
          type="button"
          class="cursor-pointer rounded-full border border-border bg-surface px-3 py-2 text-[0.8rem] font-bold text-ink hover:border-brand"
          @click="logoutAdmin"
        >
          Logout
        </button> -->
        <button
          class="cursor-pointer rounded-full border border-border bg-surface px-3 py-2 text-[0.8rem] font-bold text-ink hover:border-brand"
          type="button"
          @click="toggleTheme"
          :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'"
        >
          {{ isDark ? 'Use Light' : 'Use Dark' }}
        </button>
        <RouterLink
          to="/contact"
          class="rounded-full bg-gradient-to-br from-brand to-brand-deep px-4 py-[0.56rem] text-[0.88rem] font-bold text-white shadow-[0_10px_20px_rgba(15,118,110,0.2)] hover:-translate-y-px"
          >Let's Collaborate</RouterLink
        >
      </div>
    </div>
  </header>
</template>
