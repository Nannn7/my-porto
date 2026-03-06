<script setup>
<<<<<<< HEAD
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
=======
import { onMounted } from 'vue'
import { useProjectStore } from '../store'
import ProjectCard from '../components/ProjectCard.vue'

const projectStore = useProjectStore()

onMounted(() => {
  projectStore.fetchProjects()
})
>>>>>>> origin/main
</script>

<template>
  <section>
    <h1>Projects</h1>
<<<<<<< HEAD
    <p>Daftar project diambil dari API backend.</p>

    <p v-if="loading">Loading projects...</p>
    <p v-else-if="error">{{ error }}</p>
    <p v-else-if="!projects.length">Belum ada data project yang tersedia.</p>

    <div v-else class="project-grid">
      <ProjectCard v-for="project in projects" :key="project.id || project.name || project.title" :project="project" />
=======
    <div class="grid">
      <ProjectCard
        v-for="project in projectStore.projects"
        :key="project.id"
        :project="project"
      />
>>>>>>> origin/main
    </div>
  </section>
</template>

<style scoped>
<<<<<<< HEAD
.project-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
=======
.grid {
  display: grid;
  gap: 12px;
>>>>>>> origin/main
}
</style>
