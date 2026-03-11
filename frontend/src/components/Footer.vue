<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { usePortfolioStore } from '../store'
import { useAdminStore } from '../store/admin'

const store = usePortfolioStore()
const profile = computed(() => store.profile)
const router = useRouter()
const adminStore = useAdminStore()
const isAdminAuthenticated = computed(() => adminStore.isAuthenticated)
const adminRoute = computed(() => (isAdminAuthenticated.value ? '/admin/projects' : '/admin/login'))

const logoutAdmin = async () => {
  adminStore.logout()
  await router.push('/')
}
</script>

<template>
  <footer class="mt-12 border-t border-border bg-bg-elevated">
    <div class="flex w-full items-center justify-between gap-5 px-4 pb-10 pt-6 md:px-8 lg:px-12 max-[900px]:flex-col max-[900px]:items-start">
      <div>
        <p class="font-display font-bold text-heading">Nanda Surya Diffa</p>
        <p class="max-w-[620px] text-[0.92rem] text-muted">
          Developer with strong commitment, fast learning curve, and business-oriented technical delivery.
        </p>
      </div>

      <div class="flex flex-wrap items-center justify-end gap-[0.55rem] max-[900px]:justify-start">
        <RouterLink
          :to="adminRoute"
          class="rounded-full border border-border bg-surface px-[0.72rem] py-[0.46rem] text-[0.86rem] font-semibold text-ink hover:-translate-y-px hover:border-brand"
        >
          {{ isAdminAuthenticated ? 'Admin Dashboard' : 'Admin Login' }}
        </RouterLink>
        <button
          v-if="isAdminAuthenticated"
          type="button"
          class="cursor-pointer rounded-full border border-border bg-surface px-[0.72rem] py-[0.46rem] text-[0.86rem] font-semibold text-ink hover:-translate-y-px hover:border-brand"
          @click="logoutAdmin"
        >
          Logout
        </button>
        <a
          :href="`mailto:${profile.email}`"
          class="rounded-full border border-border bg-surface px-[0.72rem] py-[0.46rem] text-[0.86rem] font-semibold text-ink hover:-translate-y-px hover:border-brand"
          >{{ profile.email }}</a
        >
        <a
          :href="profile.whatsappUrl"
          target="_blank"
          rel="noopener"
          class="rounded-full border border-border bg-surface px-[0.72rem] py-[0.46rem] text-[0.86rem] font-semibold text-ink hover:-translate-y-px hover:border-brand"
          >WhatsApp</a
        >
        <a
          :href="profile.linkedinUrl"
          target="_blank"
          rel="noopener"
          class="rounded-full border border-border bg-surface px-[0.72rem] py-[0.46rem] text-[0.86rem] font-semibold text-ink hover:-translate-y-px hover:border-brand"
          >LinkedIn</a
        >
      </div>
    </div>
  </footer>
</template>
