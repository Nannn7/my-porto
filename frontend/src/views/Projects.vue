<script setup>
import { onMounted } from 'vue'
import { useProjectStore } from '../store'
import ProjectCard from '../components/ProjectCard.vue'

const projectStore = useProjectStore()

onMounted(() => {
  projectStore.fetchProjects()
})
</script>

<template>
  <section>
    <h1>Projects</h1>
    <div class="grid">
      <ProjectCard
        v-for="project in projectStore.projects"
        :key="project.id"
        :project="project"
      />
    </div>
  </section>
</template>

<style scoped>
.grid {
  display: grid;
  gap: 12px;
}
</style>
