<script setup>
import { computed } from 'vue'
import { usePortfolioStore } from '../store'

const store = usePortfolioStore()
const profile = computed(() => store.profile)
</script>

<template>
  <footer class="footer-wrap">
    <div class="footer">
      <div>
        <p class="footer__name">Nanda Surya Diffa</p>
        <p class="footer__caption">
          Developer with strong commitment, fast learning curve, and business-oriented technical delivery.
        </p>
      </div>

      <div class="footer__links">
        <a :href="`mailto:${profile.email}`">{{ profile.email }}</a>
        <a :href="profile.whatsappUrl" target="_blank" rel="noopener">WhatsApp</a>
        <a :href="profile.linkedinUrl" target="_blank" rel="noopener">LinkedIn</a>
      </div>
    </div>
  </footer>
</template>

<style scoped>
.footer-wrap {
  margin-top: 3rem;
  border-top: 1px solid var(--border);
  background: var(--bg-elevated);
}

.footer {
  width: min(1120px, 94vw);
  margin: 0 auto;
  padding: 1.5rem 0 2.5rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1.2rem;
}

.footer__name {
  font-family: 'Sora', 'Trebuchet MS', sans-serif;
  font-weight: 700;
  color: var(--heading);
}

.footer__caption {
  color: var(--muted);
  font-size: 0.92rem;
  max-width: 620px;
}

.footer__links {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  justify-content: flex-end;
  gap: 0.55rem;
}

.footer__links a {
  color: var(--ink);
  border: 1px solid var(--border);
  background: var(--surface);
  padding: 0.46rem 0.72rem;
  border-radius: 999px;
  font-size: 0.86rem;
  font-weight: 600;
}

.footer__links a:hover {
  border-color: var(--brand);
  transform: translateY(-1px);
}

@media (max-width: 900px) {
  .footer {
    flex-direction: column;
    align-items: flex-start;
  }

  .footer__links {
    justify-content: flex-start;
  }
}
</style>
