<script setup>
import { computed, onMounted } from 'vue'
import { usePortfolioStore } from '../store'

const store = usePortfolioStore()

const experiences = [
  {
    company: 'PT Inovasi Digital Nusantara',
    role: 'Backend Engineer',
    period: '2022 - Sekarang',
    details: 'Membangun dan mengelola API untuk produk SaaS serta optimasi query database.',
  },
  {
    company: 'Freelance',
    role: 'Full Stack Developer',
    period: '2020 - 2022',
    details: 'Mengembangkan website company profile dan dashboard internal untuk UMKM.',
  },
]

const education = [
  {
    institution: 'Universitas Indonesia',
    major: 'S1 Ilmu Komputer',
    period: '2016 - 2020',
  },
]

const organizations = [
  'Himpunan Mahasiswa Ilmu Komputer - Divisi Riset & Teknologi',
  'Komunitas Developer Jakarta - Mentor Backend Dasar',
]

const certifications = [
  'Google Associate Cloud Engineer',
  'Dicoding: Belajar Fundamental Front-End Web Development',
]

onMounted(() => {
  store.fetchSkills()
})

const skills = computed(() => store.skills)
</script>

<template>
  <section>
    <h1>About Me</h1>

    <h2>Pengalaman Kerja</h2>
    <div v-for="item in experiences" :key="`${item.company}-${item.role}`" class="card">
      <h3>{{ item.role }} - {{ item.company }}</h3>
      <p>{{ item.period }}</p>
      <p>{{ item.details }}</p>
    </div>

    <h2>Pendidikan</h2>
    <div v-for="item in education" :key="item.institution" class="card">
      <h3>{{ item.major }}</h3>
      <p>{{ item.institution }} • {{ item.period }}</p>
    </div>

    <h2>Organisasi</h2>
    <ul>
      <li v-for="org in organizations" :key="org">{{ org }}</li>
    </ul>

    <h2>Sertifikasi</h2>
    <ul>
      <li v-for="cert in certifications" :key="cert">{{ cert }}</li>
    </ul>

    <h2>Skills</h2>
    <ul>
      <li v-for="skill in skills" :key="skill">{{ skill }}</li>
    </ul>
  </section>
</template>

<style scoped>
section { display: grid; gap: 0.8rem; }
.card {
  border: 1px solid #e2e8f0;
  border-radius: 10px;
  padding: 1rem;
}
ul { margin: 0; }
</style>
