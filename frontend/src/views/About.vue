<script setup>
import { computed } from 'vue'
import { usePortfolioStore } from '../store'

const store = usePortfolioStore()

const profile = computed(() => store.profile)
const experiences = computed(() => store.experiences)
const education = computed(() => store.education)
const organizations = computed(() => store.organizations)
const certifications = computed(() => store.certifications)
const relatedCourses = computed(() => store.relatedCourses)
const skillGroups = computed(() => store.skillGroups)

const delayStyle = (index, base = 0) => ({ animationDelay: `${base + index * 80}ms` })
</script>

<template>
  <section class="grid gap-[0.9rem] [grid-template-columns:1.5fr_1fr_1fr] max-[1040px]:grid-cols-1">
    <article class="reveal grid gap-[0.55rem] rounded-lgx border border-border bg-gradient-to-b from-surface to-surface-soft p-[1.15rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]">
      <h1>About</h1>
      <p>{{ profile.summary }}</p>
      <p>
        Strong in requirement clarification, cross-functional communication, and transforming business needs into
        reliable technical solutions.
      </p>
    </article>

    <article class="reveal reveal--delay-1 grid gap-[0.55rem] rounded-lgx border border-border bg-surface p-[1.15rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]">
      <h2>Education</h2>
      <p class="font-bold text-heading">{{ education.major }}</p>
      <p>{{ education.institution }}</p>
      <p class="text-[0.9rem] text-muted">{{ education.grade }} | {{ education.period }}</p>
    </article>

    <article class="reveal reveal--delay-2 grid gap-[0.55rem] rounded-lgx border border-border bg-surface p-[1.15rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]">
      <h2>Certification</h2>
      <div
        v-for="(cert, index) in certifications"
        :key="cert.name"
        class="grid gap-1"
        :class="index > 0 ? 'mt-[0.85rem]' : ''"
      >
        <p class="font-bold text-heading">{{ cert.name }}</p>
        <p>{{ cert.issuer }}</p>
        <p class="text-[0.9rem] text-muted">{{ cert.period }}</p>
      </div>
    </article>
  </section>

  <section class="reveal reveal--delay-1 mt-4 rounded-lgx border border-border bg-surface p-[1.15rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]">
    <h2>Work Experience</h2>
    <div class="mt-[0.95rem] grid gap-[0.95rem]">
      <article
        v-for="(job, index) in experiences"
        :key="`${job.company}-${job.role}`"
        class="reveal rounded-mdx border border-border bg-surface-soft p-[0.95rem] hover:-translate-y-0.5 hover:border-brand"
        :style="delayStyle(index, 120)"
      >
        <div class="flex justify-between gap-4 max-[1040px]:flex-col">
          <div>
            <h3>{{ job.role }}</h3>
            <p class="font-bold text-heading">{{ job.company }}</p>
            <p class="text-[0.9rem] text-muted">{{ job.location }}</p>
          </div>
          <p class="text-[0.9rem] font-bold text-muted">{{ job.period }}</p>
        </div>

        <ul class="mt-[0.6rem] grid gap-[0.3rem] pl-[1.15rem]">
          <li v-for="point in job.highlights" :key="point">{{ point }}</li>
        </ul>
      </article>
    </div>
  </section>

  <section class="mt-4 grid grid-cols-2 gap-[0.9rem] max-[1040px]:grid-cols-1">
    <article class="reveal reveal--delay-2 rounded-lgx border border-border bg-surface p-[1.15rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]">
      <h2>Organizations</h2>
      <div
        v-for="(org, index) in organizations"
        :key="org.name"
        class="reveal grid gap-1"
        :class="index > 0 ? 'mt-[0.85rem]' : ''"
        :style="delayStyle(index, 160)"
      >
        <p class="font-bold text-heading">{{ org.name }}</p>
        <p class="text-[0.9rem] text-muted">{{ org.period }}</p>
        <ul class="mt-1 grid gap-[0.3rem] pl-[1.15rem]">
          <li v-for="note in org.notes" :key="note">{{ note }}</li>
        </ul>
      </div>
    </article>

    <article class="reveal reveal--delay-3 rounded-lgx border border-border bg-surface p-[1.15rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]">
      <h2>Related Courses</h2>
      <div class="mt-[0.6rem] flex flex-wrap gap-1.5">
        <span
          v-for="(course, index) in relatedCourses"
          :key="course"
          class="reveal rounded-full border border-border bg-surface-soft px-[0.56rem] py-[0.2rem] text-[0.78rem] font-semibold text-ink hover:border-brand"
          :style="delayStyle(index, 200)"
          >{{ course }}</span
        >
      </div>
    </article>
  </section>

  <section class="reveal reveal--delay-2 mt-4 rounded-lgx border border-border bg-surface p-[1.15rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]">
    <h2>Skill Matrix</h2>
    <div class="mt-[0.8rem] grid grid-cols-2 gap-[0.8rem] max-[1040px]:grid-cols-1">
      <article
        class="reveal rounded-mdx border border-border bg-surface-soft p-[0.9rem] hover:-translate-y-0.5 hover:border-brand"
        :style="delayStyle(0, 120)"
      >
        <h3 class="mb-[0.45rem] text-base">Language</h3>
        <p>{{ skillGroups.language.join(' | ') }}</p>
      </article>

      <article
        class="reveal rounded-mdx border border-border bg-surface-soft p-[0.9rem] hover:-translate-y-0.5 hover:border-brand"
        :style="delayStyle(1, 120)"
      >
        <h3 class="mb-[0.45rem] text-base">Programming</h3>
        <div class="flex flex-wrap gap-1.5">
          <span
            v-for="item in skillGroups.programmingLanguages"
            :key="item"
            class="rounded-full border border-border bg-surface-soft px-[0.56rem] py-[0.2rem] text-[0.78rem] font-semibold text-ink hover:border-brand"
            >{{ item }}</span
          >
        </div>
      </article>

      <article
        class="reveal rounded-mdx border border-border bg-surface-soft p-[0.9rem] hover:-translate-y-0.5 hover:border-brand"
        :style="delayStyle(2, 120)"
      >
        <h3 class="mb-[0.45rem] text-base">Framework</h3>
        <div class="flex flex-wrap gap-1.5">
          <span
            v-for="item in skillGroups.frameworks"
            :key="item"
            class="rounded-full border border-border bg-surface-soft px-[0.56rem] py-[0.2rem] text-[0.78rem] font-semibold text-ink hover:border-brand"
            >{{ item }}</span
          >
        </div>
      </article>

      <article
        class="reveal rounded-mdx border border-border bg-surface-soft p-[0.9rem] hover:-translate-y-0.5 hover:border-brand"
        :style="delayStyle(3, 120)"
      >
        <h3 class="mb-[0.45rem] text-base">Tools and Platforms</h3>
        <div class="flex flex-wrap gap-1.5">
          <span
            v-for="item in skillGroups.tools"
            :key="item"
            class="rounded-full border border-border bg-surface-soft px-[0.56rem] py-[0.2rem] text-[0.78rem] font-semibold text-ink hover:border-brand"
            >{{ item }}</span
          >
        </div>
      </article>

      <article
        class="reveal rounded-mdx border border-border bg-surface-soft p-[0.9rem] hover:-translate-y-0.5 hover:border-brand"
        :style="delayStyle(4, 120)"
      >
        <h3 class="mb-[0.45rem] text-base">Database</h3>
        <div class="flex flex-wrap gap-1.5">
          <span
            v-for="item in skillGroups.databases"
            :key="item"
            class="rounded-full border border-border bg-surface-soft px-[0.56rem] py-[0.2rem] text-[0.78rem] font-semibold text-ink hover:border-brand"
            >{{ item }}</span
          >
        </div>
      </article>

      <article
        class="reveal rounded-mdx border border-border bg-surface-soft p-[0.9rem] hover:-translate-y-0.5 hover:border-brand"
        :style="delayStyle(5, 120)"
      >
        <h3 class="mb-[0.45rem] text-base">API and Middleware</h3>
        <div class="flex flex-wrap gap-1.5">
          <span
            v-for="item in skillGroups.apiMiddleware"
            :key="item"
            class="rounded-full border border-border bg-surface-soft px-[0.56rem] py-[0.2rem] text-[0.78rem] font-semibold text-ink hover:border-brand"
            >{{ item }}</span
          >
        </div>
      </article>
    </div>
  </section>
</template>
