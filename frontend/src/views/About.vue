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
  <section class="about-grid">
    <article class="about-card about-card--summary reveal">
      <h1>About</h1>
      <p>{{ profile.summary }}</p>
      <p>
        Strong in requirement clarification, cross-functional communication, and transforming business needs into
        reliable technical solutions.
      </p>
    </article>

    <article class="about-card reveal reveal--delay-1">
      <h2>Education</h2>
      <p class="title">{{ education.major }}</p>
      <p>{{ education.institution }}</p>
      <p class="muted">{{ education.grade }} | {{ education.period }}</p>
    </article>

    <article class="about-card reveal reveal--delay-2">
      <h2>Certification</h2>
      <div v-for="cert in certifications" :key="cert.name" class="stack-item">
        <p class="title">{{ cert.name }}</p>
        <p>{{ cert.issuer }}</p>
        <p class="muted">{{ cert.period }}</p>
      </div>
    </article>
  </section>

  <section class="section-card reveal reveal--delay-1">
    <h2>Work Experience</h2>
    <div class="timeline">
      <article v-for="(job, index) in experiences" :key="`${job.company}-${job.role}`" class="timeline-item reveal" :style="delayStyle(index, 120)">
        <div class="timeline-item__head">
          <div>
            <h3>{{ job.role }}</h3>
            <p class="title">{{ job.company }}</p>
            <p class="muted">{{ job.location }}</p>
          </div>
          <p class="muted timeline-item__period">{{ job.period }}</p>
        </div>

        <ul>
          <li v-for="point in job.highlights" :key="point">{{ point }}</li>
        </ul>
      </article>
    </div>
  </section>

  <section class="dual-grid">
    <article class="section-card reveal reveal--delay-2">
      <h2>Organizations</h2>
      <div v-for="(org, index) in organizations" :key="org.name" class="stack-item reveal" :style="delayStyle(index, 160)">
        <p class="title">{{ org.name }}</p>
        <p class="muted">{{ org.period }}</p>
        <ul>
          <li v-for="note in org.notes" :key="note">{{ note }}</li>
        </ul>
      </div>
    </article>

    <article class="section-card reveal reveal--delay-3">
      <h2>Related Courses</h2>
      <div class="pill-wrap">
        <span v-for="(course, index) in relatedCourses" :key="course" class="pill reveal" :style="delayStyle(index, 200)">{{
          course
        }}</span>
      </div>
    </article>
  </section>

  <section class="section-card reveal reveal--delay-2">
    <h2>Skill Matrix</h2>
    <div class="matrix-grid">
      <article class="matrix-item reveal" :style="delayStyle(0, 120)">
        <h3>Language</h3>
        <p>{{ skillGroups.language.join(' | ') }}</p>
      </article>

      <article class="matrix-item reveal" :style="delayStyle(1, 120)">
        <h3>Programming</h3>
        <div class="pill-wrap">
          <span v-for="item in skillGroups.programmingLanguages" :key="item" class="pill">{{ item }}</span>
        </div>
      </article>

      <article class="matrix-item reveal" :style="delayStyle(2, 120)">
        <h3>Framework</h3>
        <div class="pill-wrap">
          <span v-for="item in skillGroups.frameworks" :key="item" class="pill">{{ item }}</span>
        </div>
      </article>

      <article class="matrix-item reveal" :style="delayStyle(3, 120)">
        <h3>Tools and Platforms</h3>
        <div class="pill-wrap">
          <span v-for="item in skillGroups.tools" :key="item" class="pill">{{ item }}</span>
        </div>
      </article>

      <article class="matrix-item reveal" :style="delayStyle(4, 120)">
        <h3>Database</h3>
        <div class="pill-wrap">
          <span v-for="item in skillGroups.databases" :key="item" class="pill">{{ item }}</span>
        </div>
      </article>

      <article class="matrix-item reveal" :style="delayStyle(5, 120)">
        <h3>API and Middleware</h3>
        <div class="pill-wrap">
          <span v-for="item in skillGroups.apiMiddleware" :key="item" class="pill">{{ item }}</span>
        </div>
      </article>
    </div>
  </section>
</template>

<style scoped>
.about-grid {
  display: grid;
  gap: 0.9rem;
  grid-template-columns: 1.5fr 1fr 1fr;
}

.about-card,
.section-card {
  border: 1px solid var(--border);
  background: var(--surface);
  border-radius: var(--radius-lg);
  padding: 1.15rem;
  box-shadow: 0 14px 28px rgba(16, 34, 40, 0.06);
}

.about-card {
  display: grid;
  gap: 0.55rem;
}

.about-card--summary {
  background: linear-gradient(170deg, var(--surface) 0%, var(--surface-soft) 100%);
}

.title {
  font-weight: 700;
  color: var(--heading);
}

.muted {
  color: var(--muted);
  font-size: 0.9rem;
}

.section-card {
  margin-top: 1rem;
}

.timeline {
  display: grid;
  gap: 0.95rem;
}

.timeline-item {
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
  padding: 0.95rem;
  background: var(--surface-soft);
  transition: transform 0.22s ease, border-color 0.22s ease;
}

.timeline-item:hover {
  transform: translateY(-2px);
  border-color: var(--brand);
}

.timeline-item__head {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
}

.timeline-item__period {
  font-weight: 700;
}

ul {
  margin: 0.6rem 0 0;
  padding-left: 1.15rem;
  display: grid;
  gap: 0.3rem;
}

.dual-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.9rem;
}

.stack-item + .stack-item {
  margin-top: 0.85rem;
}

.pill-wrap {
  margin-top: 0.6rem;
  display: flex;
  flex-wrap: wrap;
  gap: 0.4rem;
}

.pill {
  border-radius: 999px;
  background: var(--surface-soft);
  border: 1px solid var(--border);
  padding: 0.2rem 0.56rem;
  font-size: 0.78rem;
  font-weight: 600;
  color: var(--ink);
}

.pill:hover {
  border-color: var(--brand);
}

.matrix-grid {
  margin-top: 0.8rem;
  display: grid;
  gap: 0.8rem;
  grid-template-columns: repeat(2, minmax(0, 1fr));
}

.matrix-item {
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
  padding: 0.9rem;
  background: var(--surface-soft);
  transition: transform 0.22s ease, border-color 0.22s ease;
}

.matrix-item:hover {
  transform: translateY(-2px);
  border-color: var(--brand);
}

.matrix-item h3 {
  margin-bottom: 0.45rem;
  font-size: 1rem;
}

@media (max-width: 1040px) {
  .about-grid {
    grid-template-columns: 1fr;
  }

  .dual-grid,
  .matrix-grid {
    grid-template-columns: 1fr;
  }

  .timeline-item__head {
    flex-direction: column;
  }
}
</style>
