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
  <article class="card">
    <header class="card__header">
      <div>
        <p v-if="project.tag" class="card__tag">{{ project.tag }}</p>
        <h3>{{ title }}</h3>
      </div>
      <p v-if="project.period" class="card__period">{{ project.period }}</p>
    </header>

    <p class="card__desc">{{ description }}</p>

    <p v-if="project.company" class="card__company">{{ project.company }}</p>

    <ul v-if="outcomes.length" class="card__outcomes">
      <li v-for="(outcome, idx) in outcomes" :key="`${title}-${idx}`">
        {{ outcome }}
      </li>
    </ul>

    <div v-if="stack.length" class="card__stack">
      <span v-for="item in stack" :key="`${title}-${item}`" class="pill">{{ item }}</span>
    </div>

    <a v-if="project.link" :href="project.link" target="_blank" rel="noopener" class="card__link">
      Open Link
    </a>
  </article>
</template>

<style scoped>
.card {
  display: grid;
  gap: 0.8rem;
  padding: 1.05rem;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
  box-shadow: 0 16px 30px rgba(16, 34, 40, 0.06);
  transition: transform 0.25s ease, border-color 0.25s ease;
}

.card:hover {
  transform: translateY(-3px);
  border-color: var(--brand);
}

.card__header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 0.8rem;
}

.card__tag {
  display: inline-flex;
  margin-bottom: 0.3rem;
  padding: 0.18rem 0.54rem;
  border-radius: 999px;
  font-size: 0.72rem;
  font-weight: 700;
  color: var(--ink);
  background: var(--brand-soft);
}

h3 {
  font-size: 1.03rem;
}

.card__period {
  color: var(--muted);
  font-size: 0.84rem;
  font-weight: 700;
  white-space: nowrap;
}

.card__desc,
.card__company {
  color: var(--ink);
}

.card__company {
  font-weight: 700;
  font-size: 0.9rem;
}

.card__outcomes {
  margin: 0;
  padding-left: 1.2rem;
  color: var(--ink);
  display: grid;
  gap: 0.28rem;
}

.card__stack {
  display: flex;
  flex-wrap: wrap;
  gap: 0.4rem;
}

.pill {
  padding: 0.2rem 0.52rem;
  border-radius: 999px;
  font-size: 0.78rem;
  border: 1px solid var(--border);
  background: var(--surface-soft);
  color: var(--ink);
}

.card__link {
  width: fit-content;
  color: #fff;
  background: var(--accent);
  padding: 0.42rem 0.78rem;
  border-radius: 999px;
  font-size: 0.83rem;
  font-weight: 700;
}

.card__link:hover {
  transform: translateY(-1px);
}
</style>
