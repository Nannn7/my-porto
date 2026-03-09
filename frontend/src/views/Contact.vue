<script setup>
import { computed, reactive, ref } from 'vue'
import { usePortfolioStore } from '../store'
import { sendContactMessage } from '../api/axios'

const store = usePortfolioStore()
const profile = computed(() => store.profile)

const form = reactive({
  name: '',
  email: '',
  message: '',
})

const submitting = ref(false)
const status = ref('')

const submitForm = async () => {
  submitting.value = true
  status.value = ''

  try {
    await sendContactMessage(form)
    status.value = 'Message sent successfully. Thank you for reaching out.'
    form.name = ''
    form.email = ''
    form.message = ''
  } catch {
    status.value = 'Message failed to send from API endpoint. You can still contact me directly via email or WhatsApp.'
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <section class="grid gap-4 [grid-template-columns:minmax(260px,0.95fr)_minmax(320px,1.05fr)] max-[900px]:grid-cols-1">
    <article
      class="reveal grid content-start gap-[0.8rem] rounded-lgx border border-border bg-surface p-[1.25rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]"
    >
      <h1>Contact</h1>
      <p class="text-muted">
        Open to collaboration for web application enhancement, middleware integration initiatives, and product development
        support.
      </p>

      <div class="grid gap-[0.65rem]">
        <a
          :href="`mailto:${profile.email}`"
          class="reveal grid gap-[0.12rem] rounded-mdx border border-border bg-surface-soft px-[0.86rem] py-[0.74rem] hover:-translate-y-0.5 hover:border-brand"
          style="animation-delay: 120ms"
        >
          <span class="text-[0.8rem] text-muted">Email</span>
          <strong class="text-[0.92rem] text-heading">{{ profile.email }}</strong>
        </a>

        <a
          :href="profile.whatsappUrl"
          target="_blank"
          rel="noopener"
          class="reveal grid gap-[0.12rem] rounded-mdx border border-border bg-surface-soft px-[0.86rem] py-[0.74rem] hover:-translate-y-0.5 hover:border-brand"
          style="animation-delay: 180ms"
        >
          <span class="text-[0.8rem] text-muted">WhatsApp</span>
          <strong class="text-[0.92rem] text-heading">Chat directly</strong>
        </a>

        <a
          :href="profile.linkedinUrl"
          target="_blank"
          rel="noopener"
          class="reveal grid gap-[0.12rem] rounded-mdx border border-border bg-surface-soft px-[0.86rem] py-[0.74rem] hover:-translate-y-0.5 hover:border-brand"
          style="animation-delay: 240ms"
        >
          <span class="text-[0.8rem] text-muted">LinkedIn</span>
          <strong class="text-[0.92rem] text-heading">Professional profile</strong>
        </a>
      </div>
    </article>

    <article class="reveal reveal--delay-1 rounded-lgx border border-border bg-surface p-[1.25rem] shadow-[0_14px_28px_rgba(16,34,40,0.06)]">
      <h2 class="mb-3">Send a Message</h2>
      <form class="grid gap-[0.68rem]" @submit.prevent="submitForm">
        <label class="grid gap-[0.3rem] font-semibold text-ink">
          Name
          <input
            v-model="form.name"
            type="text"
            required
            autocomplete="name"
            class="rounded-xl border border-border bg-surface-soft px-[0.72rem] py-[0.6rem] text-ink focus:border-brand focus:outline-none focus:ring-2 focus:ring-focus"
          />
        </label>

        <label class="grid gap-[0.3rem] font-semibold text-ink">
          Email
          <input
            v-model="form.email"
            type="email"
            required
            autocomplete="email"
            class="rounded-xl border border-border bg-surface-soft px-[0.72rem] py-[0.6rem] text-ink focus:border-brand focus:outline-none focus:ring-2 focus:ring-focus"
          />
        </label>

        <label class="grid gap-[0.3rem] font-semibold text-ink">
          Message
          <textarea
            v-model="form.message"
            rows="6"
            required
            class="rounded-xl border border-border bg-surface-soft px-[0.72rem] py-[0.6rem] text-ink focus:border-brand focus:outline-none focus:ring-2 focus:ring-focus"
          ></textarea>
        </label>

        <button
          type="submit"
          :disabled="submitting"
          class="w-fit cursor-pointer rounded-xl bg-gradient-to-br from-brand to-brand-deep px-[0.94rem] py-[0.62rem] font-bold text-white hover:-translate-y-px disabled:cursor-not-allowed disabled:opacity-70"
        >
          {{ submitting ? 'Sending...' : 'Send Message' }}
        </button>
      </form>

      <p
        v-if="status"
        class="mt-3 rounded-xl border border-dashed border-border bg-surface-soft px-[0.74rem] py-[0.6rem] text-ink"
      >
        {{ status }}
      </p>
    </article>
  </section>
</template>
