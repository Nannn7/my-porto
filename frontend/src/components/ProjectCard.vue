<script setup>
defineProps({
  project: {
    type: Object,
    required: true,
  },
})
</script>

<template>
  <article class="project-card">
    <h3>{{ project.title || project.name || 'Untitled Project' }}</h3>
    <p>{{ project.description || 'No description provided yet.' }}</p>
    <a v-if="project.link" :href="project.link" target="_blank" rel="noopener">View Project</a>
  </article>
</template>

<style scoped>
.project-card {
  border: 1px solid #e2e8f0;
  border-radius: 10px;
  padding: 1rem;
  background: #fff;
  box-shadow: 0 3px 10px rgba(15, 23, 42, 0.08);
}
h3 { margin: 0 0 0.5rem; }
p { color: #334155; }
a { color: #0284c7; text-decoration: none; font-weight: 600; }
</style>
