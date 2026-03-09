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
  <section class="reveal rounded-lgx border border-border bg-surface p-[1.2rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]">
    <h1>Career Highlights</h1>
    <p class="mt-[0.4rem] text-muted">
      Representative works aligned with real responsibilities from internship to full-time professional delivery.
    </p>

    <div class="mt-4 grid grid-cols-2 gap-[0.85rem] max-[920px]:grid-cols-1">
      <ProjectCard
        v-for="(project, index) in featuredProjects"
        :key="project.title"
        :project="project"
        class="reveal"
        :style="delayStyle(index, 100)"
      />
    </div>
  </section>

  <section
    class="reveal reveal--delay-1 mt-4 rounded-lgx border border-border bg-gradient-to-b from-surface to-surface-soft p-[1.2rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]"
  >
    <h2>Additional Projects from API</h2>
    <p class="mt-[0.4rem] text-muted">
      This section is connected to backend endpoint and can be used to display external portfolio items.
    </p>

    <p
      v-if="loadingRemoteProjects"
      class="mt-[0.8rem] rounded-xl border border-dashed border-border bg-surface-soft px-[0.85rem] py-[0.7rem] text-ink"
    >
      Loading additional projects...
    </p>
    <p
      v-else-if="remoteProjectsError"
      class="mt-[0.8rem] rounded-xl border border-dashed border-[#fdba74] bg-[#fff7ed] px-[0.85rem] py-[0.7rem] text-[#9a3412]"
    >
      {{ remoteProjectsError }}
    </p>
    <p
      v-else-if="!remoteProjects.length"
      class="mt-[0.8rem] rounded-xl border border-dashed border-border bg-surface-soft px-[0.85rem] py-[0.7rem] text-ink"
    >
      No additional project data available yet. You can still review curated highlights above.
    </p>

    <div v-else class="mt-4 grid grid-cols-2 gap-[0.85rem] max-[920px]:grid-cols-1">
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
