<script setup>
import { computed } from 'vue'
import { usePortfolioStore } from '../store'

const store = usePortfolioStore()

const profile = computed(() => store.profile)
const quickStats = computed(() => store.quickStats)
const allSkills = computed(() => store.allSkills.slice(0, 16))
const focusAreas = computed(() => store.skillGroups.apiMiddleware)

const statDelayStyle = (index) => ({ animationDelay: `${100 + index * 80}ms` })
const skillDelayStyle = (index) => ({ animationDelay: `${220 + index * 45}ms` })
</script>

<template>
  <section class="hero reveal">
    <div class="hero__content">
      <p class="badge">Available for Collaboration</p>
      <h1>{{ profile.name }}</h1>
      <p class="hero__headline">{{ profile.headline }}</p>
      <p class="hero__summary">{{ profile.summary }}</p>

      <div class="hero__links">
        <a :href="`mailto:${profile.email}`" class="chip">{{ profile.email }}</a>
        <a :href="profile.whatsappUrl" target="_blank" rel="noopener" class="chip">WhatsApp</a>
        <a :href="profile.linkedinUrl" target="_blank" rel="noopener" class="chip">LinkedIn</a>
      </div>

      <div class="hero__actions">
        <RouterLink to="/projects" class="btn btn--primary">See Work Highlights</RouterLink>
        <RouterLink to="/contact" class="btn btn--ghost">Let's Talk</RouterLink>
      </div>
    </div>

    <aside class="hero__panel reveal reveal--delay-1">
      <p class="panel__label">Base Location</p>
      <p class="panel__value">{{ profile.location }}</p>
      <p class="panel__label">Focus Areas</p>
      <div class="panel__chips">
        <span v-for="item in focusAreas" :key="item" class="mini-chip">{{ item }}</span>
      </div>
    </aside>
  </section>

  <section class="stats">
    <article v-for="(stat, index) in quickStats" :key="stat.label" class="stat-card reveal" :style="statDelayStyle(index)">
      <p class="stat-card__value">{{ stat.value }}</p>
      <p class="stat-card__label">{{ stat.label }}</p>
    </article>
  </section>

  <section class="skills reveal reveal--delay-2">
    <div class="skills__heading">
      <h2>Core Technical Stack</h2>
      <RouterLink to="/about" class="skills__link">View full profile</RouterLink>
    </div>

    <div class="skills__cloud">
      <span v-for="(skill, index) in allSkills" :key="skill" class="skill-pill reveal" :style="skillDelayStyle(index)">{{
        skill
      }}</span>
    </div>
  </section>
</template>

<style scoped>
.hero {
  display: grid;
  grid-template-columns: 1.35fr minmax(240px, 0.65fr);
  gap: 1rem;
  align-items: stretch;
}

.hero__content,
.hero__panel {
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  background: var(--surface);
  box-shadow: var(--shadow);
}

.hero__content {
  padding: 1.4rem;
  display: grid;
  gap: 0.9rem;
}

.hero__panel {
  padding: 1.2rem;
  background: linear-gradient(170deg, var(--surface) 0%, var(--surface-soft) 100%);
  display: grid;
  gap: 0.6rem;
  align-content: start;
}

.badge {
  width: fit-content;
  font-size: 0.74rem;
  font-weight: 700;
  color: var(--ink);
  background: var(--brand-soft);
  border-radius: 999px;
  padding: 0.24rem 0.6rem;
}

h1 {
  font-size: clamp(1.6rem, 3.2vw, 2.25rem);
}

.hero__headline {
  color: var(--link);
  font-weight: 700;
}

.hero__summary {
  color: var(--muted);
}

.hero__links {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.chip {
  padding: 0.38rem 0.7rem;
  border-radius: 999px;
  border: 1px solid var(--border);
  background: var(--surface-soft);
  font-size: 0.83rem;
  font-weight: 600;
  color: var(--ink);
}

.chip:hover {
  border-color: var(--brand);
  transform: translateY(-1px);
}

.hero__actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.6rem;
}

.btn {
  border-radius: 12px;
  padding: 0.62rem 0.92rem;
  font-weight: 700;
  font-size: 0.88rem;
}

.btn--primary {
  color: #fff;
  background: linear-gradient(140deg, var(--brand) 0%, var(--brand-deep) 100%);
}

.btn--ghost {
  border: 1px solid var(--border);
  color: var(--ink);
  background: var(--surface);
}

.btn:hover {
  transform: translateY(-1px);
}

.panel__label {
  font-size: 0.8rem;
  color: var(--muted);
  text-transform: uppercase;
  letter-spacing: 0.06em;
  font-weight: 700;
}

.panel__value {
  font-size: 1.2rem;
  font-weight: 700;
  margin-bottom: 0.25rem;
  color: var(--heading);
}

.panel__chips {
  display: flex;
  flex-wrap: wrap;
  gap: 0.4rem;
}

.mini-chip {
  font-size: 0.75rem;
  border-radius: 999px;
  padding: 0.2rem 0.52rem;
  background: var(--brand-soft);
  color: var(--ink);
  font-weight: 700;
}

.stats {
  margin-top: 1rem;
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 0.8rem;
}

.stat-card {
  border-radius: var(--radius-md);
  border: 1px solid var(--border);
  background: var(--surface);
  padding: 0.9rem;
  transition: transform 0.25s ease, border-color 0.25s ease;
}

.stat-card:hover {
  transform: translateY(-3px);
  border-color: var(--brand);
}

.stat-card__value {
  font-size: 1.25rem;
  font-family: 'Sora', 'Trebuchet MS', sans-serif;
  font-weight: 700;
  color: var(--heading);
}

.stat-card__label {
  color: var(--muted);
  font-size: 0.82rem;
}

.skills {
  margin-top: 1rem;
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  background: var(--surface);
  padding: 1.2rem;
  box-shadow: 0 10px 24px rgba(16, 34, 40, 0.06);
}

.skills__heading {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 0.8rem;
  margin-bottom: 0.9rem;
}

.skills__link {
  color: var(--link);
  font-weight: 700;
  font-size: 0.86rem;
}

.skills__cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 0.45rem;
}

.skill-pill {
  border-radius: 999px;
  background: var(--surface-soft);
  border: 1px solid var(--border);
  padding: 0.24rem 0.64rem;
  font-size: 0.8rem;
  font-weight: 600;
  color: var(--ink);
}

.skill-pill:hover {
  border-color: var(--brand);
}

@media (max-width: 980px) {
  .hero {
    grid-template-columns: 1fr;
  }

  .stats {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 600px) {
  .hero__content,
  .hero__panel,
  .skills {
    padding: 1rem;
  }

  .stats {
    grid-template-columns: 1fr;
  }
}
</style>
