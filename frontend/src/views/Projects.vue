<script setup>
import { computed, onMounted } from 'vue'
import { usePortfolioStore } from '../store'
import ProjectCard from '../components/ProjectCard.vue'

const store = usePortfolioStore()

onMounted(() => {
  store.fetchRemoteProjects()
})

const featuredProjects = computed(() => store.featuredProjects)
const remoteProjects = computed(() => store.remoteProjects)
const loadingRemoteProjects = computed(() => store.loadingRemoteProjects)
const remoteProjectsError = computed(() => store.remoteProjectsError)

const delayStyle = (index, base = 0) => ({ animationDelay: `${base + index * 80}ms` })
</script>

<template>
  <section class="section-card reveal">
    <h1>Career Highlights</h1>
    <p class="section-desc">
      Representative works aligned with real responsibilities from internship to full-time professional delivery.
    </p>

    <div class="project-grid">
      <ProjectCard
        v-for="(project, index) in featuredProjects"
        :key="project.title"
        :project="project"
        class="reveal"
        :style="delayStyle(index, 100)"
      />
    </div>
  </section>

  <section class="section-card section-card--soft reveal reveal--delay-1">
    <h2>Additional Projects from API</h2>
    <p class="section-desc">
      This section is connected to backend endpoint and can be used to display external portfolio items.
    </p>

    <p v-if="loadingRemoteProjects" class="state-msg">Loading additional projects...</p>
    <p v-else-if="remoteProjectsError" class="state-msg state-msg--error">{{ remoteProjectsError }}</p>
    <p v-else-if="!remoteProjects.length" class="state-msg">
      No additional project data available yet. You can still review curated highlights above.
    </p>

    <div v-else class="project-grid">
      <ProjectCard
        v-for="(project, index) in remoteProjects"
        :key="project.id || project.slug || project.title || project.name"
        :project="project"
        class="reveal"
        :style="delayStyle(index, 120)"
      />
    </div>
  </section>
</template>

<style scoped>
.section-card {
  border: 1px solid var(--border);
  background: var(--surface);
  border-radius: var(--radius-lg);
  padding: 1.2rem;
  box-shadow: 0 14px 28px rgba(16, 34, 40, 0.06);
}

.section-card + .section-card {
  margin-top: 1rem;
}

.section-card--soft {
  background: linear-gradient(170deg, var(--surface) 0%, var(--surface-soft) 100%);
}

.section-desc {
  color: var(--muted);
  margin-top: 0.4rem;
}

.project-grid {
  margin-top: 1rem;
  display: grid;
  gap: 0.85rem;
  grid-template-columns: repeat(2, minmax(0, 1fr));
}

.state-msg {
  margin-top: 0.8rem;
  padding: 0.7rem 0.85rem;
  border: 1px dashed var(--border);
  border-radius: 12px;
  color: var(--ink);
  background: var(--surface-soft);
}

.state-msg--error {
  border-color: #fdba74;
  color: #9a3412;
  background: #fff7ed;
}

@media (max-width: 920px) {
  .project-grid {
    grid-template-columns: 1fr;
  }
}
</style>
