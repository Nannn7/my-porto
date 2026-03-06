<<<<<<< HEAD
<script setup>
import { reactive, ref } from 'vue'
import axios from 'axios'
import { API_BASE_URL } from '../store'

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
    await axios.post(`${API_BASE_URL}/contact`, form)
    status.value = 'Pesan berhasil dikirim.'
    form.name = ''
    form.email = ''
    form.message = ''
  } catch {
    status.value = 'Gagal mengirim pesan. Coba lagi nanti.'
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <section class="contact">
    <h1>Contact</h1>
    <form @submit.prevent="submitForm">
      <label>
        Name
        <input v-model="form.name" type="text" required />
      </label>
      <label>
        Email
        <input v-model="form.email" type="email" required />
      </label>
      <label>
        Message
        <textarea v-model="form.message" rows="6" required />
      </label>
      <button type="submit" :disabled="submitting">
        {{ submitting ? 'Sending...' : 'Send Message' }}
      </button>
    </form>
    <p v-if="status">{{ status }}</p>
  </section>
</template>

<style scoped>
.contact { max-width: 640px; }
form { display: grid; gap: 0.8rem; }
label { display: grid; gap: 0.35rem; }
input, textarea {
  padding: 0.6rem;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
}
button {
  width: fit-content;
  border: 0;
  background: #0369a1;
  color: white;
  padding: 0.6rem 1rem;
  border-radius: 8px;
}
button:disabled { opacity: 0.7; }
</style>
=======
<template>
  <section>
    <h1>Contact</h1>
    <p>Silakan hubungi melalui email.</p>
  </section>
</template>
>>>>>>> origin/main
