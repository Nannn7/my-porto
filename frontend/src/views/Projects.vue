<script setup>
import { onMounted, computed } from 'vue'
import { usePortfolioStore } from '../store'
import ProjectCard from '../components/ProjectCard.vue'

const store = usePortfolioStore()

onMounted(() => {
  store.fetchProjects()
})

const projects = computed(() => store.projects)
const loading = computed(() => store.loadingProjects)
const error = computed(() => store.projectsError)
</script>

<template>
  <section>
    <h1>Projects</h1>
    <p>Daftar project diambil dari API backend.</p>

    <p v-if="loading">Loading projects...</p>
    <p v-else-if="error">{{ error }}</p>
    <p v-else-if="!projects.length">Belum ada data project yang tersedia.</p>

    <div v-else class="project-grid">
      <ProjectCard v-for="project in projects" :key="project.id || project.name || project.title" :project="project" />
    </div>
  </section>
</template>

<style scoped>
.project-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
}
</style>
