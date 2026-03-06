<<<<<<< HEAD
<script setup>
import { onMounted, computed } from 'vue'
import { usePortfolioStore } from '../store'

const store = usePortfolioStore()

const profile = {
  name: 'Muhammad Porto',
  role: 'Full Stack Developer',
  location: 'Jakarta, Indonesia',
  intro:
    'Saya membangun aplikasi web end-to-end dengan fokus pada backend yang kuat, frontend yang responsif, dan pengalaman pengguna yang sederhana.',
}

onMounted(() => {
  store.fetchSkills()
})

const skills = computed(() => store.skills)
</script>

<template>
  <section class="hero">
    <p class="hero__role">{{ profile.role }}</p>
    <h1>{{ profile.name }}</h1>
    <p class="hero__location">📍 {{ profile.location }}</p>
    <p>{{ profile.intro }}</p>
  </section>

  <section class="skills">
    <h2>Skills</h2>
    <ul>
      <li v-for="skill in skills" :key="skill">{{ skill }}</li>
    </ul>
  </section>
</template>

<style scoped>
.hero {
  padding: 2rem;
  border-radius: 12px;
  background: linear-gradient(120deg, #e0f2fe, #f8fafc);
}
.hero__role { color: #0369a1; font-weight: 600; margin-bottom: 0.3rem; }
.hero__location { color: #334155; }
.skills { margin-top: 2rem; }
ul { display: flex; flex-wrap: wrap; gap: 0.7rem; padding: 0; list-style: none; }
li { background: #f1f5f9; padding: 0.45rem 0.8rem; border-radius: 99px; }
</style>
=======
<template>
  <section>
    <h1>Home</h1>
    <p>Selamat datang di portfolio.</p>
  </section>
</template>
>>>>>>> origin/main
