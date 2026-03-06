<script setup>
defineProps({
  project: {
    type: Object,
    required: true,
  },
})
</script>

<template>
  <article class="card">
    <h3>{{ project.title }}</h3>
    <p>{{ project.description }}</p>
  </article>
</template>

<style scoped>
.card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 12px;
}
</style>
