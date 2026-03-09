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
  <section class="contact-grid">
    <article class="contact-card contact-card--info reveal">
      <h1>Contact</h1>
      <p>
        Open to collaboration for web application enhancement, middleware integration initiatives, and product development
        support.
      </p>

      <div class="contact-links">
        <a :href="`mailto:${profile.email}`" class="reveal" style="animation-delay: 120ms">
          <span>Email</span>
          <strong>{{ profile.email }}</strong>
        </a>

        <a :href="profile.whatsappUrl" target="_blank" rel="noopener" class="reveal" style="animation-delay: 180ms">
          <span>WhatsApp</span>
          <strong>Chat directly</strong>
        </a>

        <a :href="profile.linkedinUrl" target="_blank" rel="noopener" class="reveal" style="animation-delay: 240ms">
          <span>LinkedIn</span>
          <strong>Professional profile</strong>
        </a>
      </div>
    </article>

    <article class="contact-card reveal reveal--delay-1">
      <h2>Send a Message</h2>
      <form @submit.prevent="submitForm">
        <label>
          Name
          <input v-model="form.name" type="text" required autocomplete="name" />
        </label>

        <label>
          Email
          <input v-model="form.email" type="email" required autocomplete="email" />
        </label>

        <label>
          Message
          <textarea v-model="form.message" rows="6" required></textarea>
        </label>

        <button type="submit" :disabled="submitting">
          {{ submitting ? 'Sending...' : 'Send Message' }}
        </button>
      </form>

      <p v-if="status" class="status">{{ status }}</p>
    </article>
  </section>
</template>

<style scoped>
.contact-grid {
  display: grid;
  grid-template-columns: minmax(260px, 0.95fr) minmax(320px, 1.05fr);
  gap: 1rem;
}

.contact-card {
  border: 1px solid var(--border);
  background: var(--surface);
  border-radius: var(--radius-lg);
  padding: 1.25rem;
  box-shadow: 0 14px 28px rgba(16, 34, 40, 0.06);
}

.contact-card--info {
  display: grid;
  align-content: start;
  gap: 0.8rem;
}

.contact-card--info p {
  color: var(--muted);
}

.contact-links {
  display: grid;
  gap: 0.65rem;
}

.contact-links a {
  display: grid;
  gap: 0.12rem;
  border-radius: var(--radius-md);
  border: 1px solid var(--border);
  background: var(--surface-soft);
  padding: 0.74rem 0.86rem;
}

.contact-links a:hover {
  border-color: var(--brand);
  transform: translateY(-2px);
}

.contact-links span {
  color: var(--muted);
  font-size: 0.8rem;
}

.contact-links strong {
  color: var(--heading);
  font-size: 0.92rem;
}

h2 {
  margin-bottom: 0.75rem;
}

form {
  display: grid;
  gap: 0.68rem;
}

label {
  display: grid;
  gap: 0.3rem;
  font-weight: 600;
  color: var(--ink);
}

input,
textarea {
  border: 1px solid var(--border);
  border-radius: 12px;
  background: var(--surface-soft);
  padding: 0.6rem 0.72rem;
  color: var(--ink);
}

input:focus,
textarea:focus {
  outline: 2px solid var(--focus);
  border-color: var(--brand);
}

button {
  width: fit-content;
  border: 0;
  border-radius: 12px;
  color: #fff;
  background: linear-gradient(135deg, var(--brand) 0%, var(--brand-deep) 100%);
  font-weight: 700;
  padding: 0.62rem 0.94rem;
  cursor: pointer;
}

button:hover {
  transform: translateY(-1px);
}

button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.status {
  margin-top: 0.75rem;
  border: 1px dashed var(--border);
  border-radius: 12px;
  padding: 0.6rem 0.74rem;
  background: var(--surface-soft);
  color: var(--ink);
}

@media (max-width: 900px) {
  .contact-grid {
    grid-template-columns: 1fr;
  }
}
</style>
