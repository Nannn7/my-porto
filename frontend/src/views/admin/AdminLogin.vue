<script setup>
import { reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAdminStore } from '../../store/admin'

const router = useRouter()
const route = useRoute()
const adminStore = useAdminStore()

const form = reactive({
  username: '',
  password: '',
})

const submitError = ref('')

const submit = async () => {
  submitError.value = ''

  try {
    await adminStore.login(form.username, form.password)
    const redirectPath =
      typeof route.query.redirect === 'string' && route.query.redirect.startsWith('/admin')
        ? route.query.redirect
        : '/admin/projects'

    await router.push(redirectPath)
  } catch (error) {
    submitError.value = error?.response?.data?.error || adminStore.error || 'Login failed'
  }
}
</script>

<template>
  <section class="mx-auto mt-8 w-full max-w-[480px] rounded-lgx border border-border bg-surface p-6 shadow-soft">
    <h1 class="text-[1.6rem]">Admin Login</h1>
    <p class="mt-1 text-muted">Login untuk mengelola data project portfolio.</p>

    <form class="mt-5 grid gap-3" @submit.prevent="submit">
      <label class="grid gap-1 text-sm font-semibold text-ink">
        Username
        <input
          v-model="form.username"
          type="text"
          required
          autocomplete="username"
          class="rounded-xl border border-border bg-surface-soft px-3 py-2 text-ink focus:border-brand focus:outline-none focus:ring-2 focus:ring-focus"
        />
      </label>

      <label class="grid gap-1 text-sm font-semibold text-ink">
        Password
        <input
          v-model="form.password"
          type="password"
          required
          autocomplete="current-password"
          class="rounded-xl border border-border bg-surface-soft px-3 py-2 text-ink focus:border-brand focus:outline-none focus:ring-2 focus:ring-focus"
        />
      </label>

      <button
        type="submit"
        :disabled="adminStore.loading"
        class="mt-1 w-fit rounded-xl bg-gradient-to-br from-brand to-brand-deep px-4 py-2 text-sm font-bold text-white hover:-translate-y-px disabled:cursor-not-allowed disabled:opacity-70"
      >
        {{ adminStore.loading ? 'Logging in...' : 'Login' }}
      </button>
    </form>

    <p
      v-if="submitError"
      class="mt-4 rounded-xl border border-dashed border-[#fdba74] bg-[#fff7ed] px-3 py-2 text-sm text-[#9a3412]"
    >
      {{ submitError }}
    </p>
  </section>
</template>
