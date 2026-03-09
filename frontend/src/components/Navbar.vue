<script setup>
import { computed } from 'vue'
import { useTheme } from '../composables/useTheme'

const navItems = [
  { name: 'Home', to: '/' },
  { name: 'About', to: '/about' },
  { name: 'Highlights', to: '/projects' },
  { name: 'Contact', to: '/contact' },
]

const { theme, toggleTheme } = useTheme()
const isDark = computed(() => theme.value === 'dark')
</script>

<template>
  <header class="navbar-wrap">
    <div class="navbar reveal">
      <RouterLink to="/" class="brand" aria-label="Go to home page">
        <p class="brand__name">Nanda Surya Diffa</p>
        <p class="brand__role">Web Developer, Android Developer, API & Middleware</p>
      </RouterLink>

      <nav class="nav-links" aria-label="Main navigation">
        <RouterLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          class="nav-link"
          active-class="nav-link--active"
        >
          {{ item.name }}
        </RouterLink>
      </nav>

      <div class="nav-actions">
        <button class="theme-btn" type="button" @click="toggleTheme" :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'">
          {{ isDark ? 'Use Light' : 'Use Dark' }}
        </button>
        <RouterLink to="/contact" class="contact-btn">Let's Collaborate</RouterLink>
      </div>
    </div>
  </header>
</template>

<style scoped>
.navbar-wrap {
  position: sticky;
  top: 0;
  z-index: 20;
  backdrop-filter: blur(8px);
  background: var(--bg-elevated);
  border-bottom: 1px solid var(--border);
}

.navbar {
  width: min(1120px, 94vw);
  margin: 0 auto;
  min-height: 80px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.75rem 0;
}

.brand {
  display: grid;
  gap: 0.1rem;
}

.brand__name {
  font-family: 'Sora', 'Trebuchet MS', sans-serif;
  font-weight: 700;
  font-size: 0.96rem;
  color: var(--heading);
}

.brand__role {
  color: var(--muted);
  font-size: 0.76rem;
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  flex-wrap: wrap;
  justify-content: center;
}

.nav-link {
  color: var(--muted);
  padding: 0.48rem 0.74rem;
  border-radius: 999px;
  font-weight: 600;
  font-size: 0.9rem;
  transition: background-color 0.22s ease, color 0.22s ease, transform 0.22s ease;
}

.nav-link:hover {
  background: var(--brand-soft);
  color: var(--ink);
  transform: translateY(-1px);
}

.nav-link--active {
  color: var(--ink);
  background: var(--brand-soft);
}

.nav-actions {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.theme-btn {
  border: 1px solid var(--border);
  background: var(--surface);
  color: var(--ink);
  border-radius: 999px;
  padding: 0.5rem 0.8rem;
  font-weight: 700;
  font-size: 0.8rem;
  cursor: pointer;
}

.theme-btn:hover {
  border-color: var(--brand);
}

.contact-btn {
  color: #fff;
  background: linear-gradient(135deg, var(--brand) 0%, var(--brand-deep) 100%);
  border-radius: 999px;
  padding: 0.56rem 0.98rem;
  font-weight: 700;
  font-size: 0.88rem;
  box-shadow: 0 10px 20px rgba(15, 118, 110, 0.2);
}

.contact-btn:hover {
  transform: translateY(-1px);
}

@media (max-width: 1080px) {
  .navbar {
    flex-wrap: wrap;
    justify-content: center;
    text-align: center;
    padding-bottom: 0.9rem;
  }

  .brand,
  .nav-links,
  .nav-actions {
    width: 100%;
    justify-content: center;
  }

  .brand__role {
    display: none;
  }
}
</style>
