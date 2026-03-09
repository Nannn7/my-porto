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
  <section class="reveal grid items-stretch gap-4 [grid-template-columns:1.35fr_minmax(240px,0.65fr)] max-[980px]:grid-cols-1">
    <div class="grid gap-[0.9rem] rounded-lgx border border-border bg-surface p-[1.4rem] shadow-soft max-[600px]:p-4">
      <p class="w-fit rounded-full bg-brand-soft px-[0.6rem] py-[0.24rem] text-[0.74rem] font-bold text-ink">
        Available for Collaboration
      </p>
      <h1 class="text-[clamp(1.6rem,3.2vw,2.25rem)]">{{ profile.name }}</h1>
      <p class="font-bold text-link">{{ profile.headline }}</p>
      <p class="text-muted">{{ profile.summary }}</p>

      <div class="flex flex-wrap gap-2">
        <a
          :href="`mailto:${profile.email}`"
          class="rounded-full border border-border bg-surface-soft px-[0.7rem] py-[0.38rem] text-[0.83rem] font-semibold text-ink hover:-translate-y-px hover:border-brand"
          >{{ profile.email }}</a
        >
        <a
          :href="profile.whatsappUrl"
          target="_blank"
          rel="noopener"
          class="rounded-full border border-border bg-surface-soft px-[0.7rem] py-[0.38rem] text-[0.83rem] font-semibold text-ink hover:-translate-y-px hover:border-brand"
          >WhatsApp</a
        >
        <a
          :href="profile.linkedinUrl"
          target="_blank"
          rel="noopener"
          class="rounded-full border border-border bg-surface-soft px-[0.7rem] py-[0.38rem] text-[0.83rem] font-semibold text-ink hover:-translate-y-px hover:border-brand"
          >LinkedIn</a
        >
      </div>

      <div class="flex flex-wrap gap-[0.6rem]">
        <RouterLink
          to="/projects"
          class="rounded-xl bg-gradient-to-br from-brand to-brand-deep px-[0.92rem] py-[0.62rem] text-[0.88rem] font-bold text-white hover:-translate-y-px"
          >See Work Highlights</RouterLink
        >
        <RouterLink
          to="/contact"
          class="rounded-xl border border-border bg-surface px-[0.92rem] py-[0.62rem] text-[0.88rem] font-bold text-ink hover:-translate-y-px"
          >Let's Talk</RouterLink
        >
      </div>
    </div>

    <aside
      class="reveal reveal--delay-1 grid content-start gap-[0.6rem] rounded-lgx border border-border bg-gradient-to-b from-surface to-surface-soft p-[1.2rem] shadow-soft max-[600px]:p-4"
    >
      <p class="text-[0.8rem] font-bold uppercase tracking-[0.06em] text-muted">Base Location</p>
      <p class="mb-1 text-[1.2rem] font-bold text-heading">{{ profile.location }}</p>
      <p class="text-[0.8rem] font-bold uppercase tracking-[0.06em] text-muted">Focus Areas</p>
      <div class="flex flex-wrap gap-1.5">
        <span
          v-for="item in focusAreas"
          :key="item"
          class="rounded-full bg-brand-soft px-[0.52rem] py-[0.2rem] text-[0.75rem] font-bold text-ink"
          >{{ item }}</span
        >
      </div>
    </aside>
  </section>

  <section class="mt-4 grid grid-cols-4 gap-[0.8rem] max-[980px]:grid-cols-2 max-[600px]:grid-cols-1">
    <article
      v-for="(stat, index) in quickStats"
      :key="stat.label"
      class="reveal rounded-mdx border border-border bg-surface p-[0.9rem] hover:-translate-y-[3px] hover:border-brand"
      :style="statDelayStyle(index)"
    >
      <p class="font-display text-[1.25rem] font-bold text-heading">{{ stat.value }}</p>
      <p class="text-[0.82rem] text-muted">{{ stat.label }}</p>
    </article>
  </section>

  <section
    class="reveal reveal--delay-2 mt-4 rounded-lgx border border-border bg-surface p-[1.2rem] shadow-[0_10px_24px_rgba(16,34,40,0.06)] max-[600px]:p-4"
  >
    <div class="mb-[0.9rem] flex items-center justify-between gap-[0.8rem]">
      <h2>Core Technical Stack</h2>
      <RouterLink to="/about" class="text-[0.86rem] font-bold text-link">View full profile</RouterLink>
    </div>

    <div class="flex flex-wrap gap-[0.45rem]">
      <span
        v-for="(skill, index) in allSkills"
        :key="skill"
        class="reveal rounded-full border border-border bg-surface-soft px-[0.64rem] py-[0.24rem] text-[0.8rem] font-semibold text-ink hover:border-brand"
        :style="skillDelayStyle(index)"
        >{{ skill }}</span
      >
    </div>
  </section>
</template>
