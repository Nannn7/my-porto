<script setup>
import { computed } from 'vue'

const props = defineProps({
  project: {
    type: Object,
    required: true,
  },
})

const title = computed(() => props.project.title || props.project.name || 'Untitled work item')
const description = computed(() => props.project.description || 'No description available yet.')
const stack = computed(() => (Array.isArray(props.project.stack) ? props.project.stack : []))
const outcomes = computed(() => (Array.isArray(props.project.outcomes) ? props.project.outcomes : []))
</script>

<template>
  <article
    class="grid gap-[0.8rem] rounded-mdx border border-border bg-surface p-[1.05rem] shadow-[0_16px_30px_rgba(16,34,40,0.06)] hover:-translate-y-[3px] hover:border-brand"
  >
    <header class="flex items-start justify-between gap-[0.8rem]">
      <div>
        <p
          v-if="project.tag"
          class="mb-[0.3rem] inline-flex rounded-full bg-brand-soft px-[0.54rem] py-[0.18rem] text-[0.72rem] font-bold text-ink"
        >
          {{ project.tag }}
        </p>
        <h3 class="text-[1.03rem]">{{ title }}</h3>
      </div>
      <p v-if="project.period" class="whitespace-nowrap text-[0.84rem] font-bold text-muted">{{ project.period }}</p>
    </header>

    <p class="text-ink">{{ description }}</p>

    <p v-if="project.company" class="text-[0.9rem] font-bold text-ink">{{ project.company }}</p>

    <ul v-if="outcomes.length" class="grid gap-[0.28rem] pl-[1.2rem] text-ink">
      <li v-for="(outcome, idx) in outcomes" :key="`${title}-${idx}`">
        {{ outcome }}
      </li>
    </ul>

    <div v-if="stack.length" class="flex flex-wrap gap-[0.4rem]">
      <span
        v-for="item in stack"
        :key="`${title}-${item}`"
        class="rounded-full border border-border bg-surface-soft px-[0.52rem] py-[0.2rem] text-[0.78rem] text-ink"
        >{{ item }}</span
      >
    </div>

    <a
      v-if="project.link"
      :href="project.link"
      target="_blank"
      rel="noopener"
      class="w-fit rounded-full bg-accent px-[0.78rem] py-[0.42rem] text-[0.83rem] font-bold text-white hover:-translate-y-px"
    >
      Open Link
    </a>
  </article>
</template>
