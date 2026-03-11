<script setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import {
  createAdminProject,
  deleteAdminProject,
  getAdminProjects,
  updateAdminProject,
} from '../../api/axios'
import { useAdminStore } from '../../store/admin'

const router = useRouter()
const adminStore = useAdminStore()

const projects = ref([])
const loading = ref(false)
const saving = ref(false)
const pageError = ref('')
const formError = ref('')

const form = reactive({
  id: null,
  title: '',
  description: '',
  tech_stack: '',
  image_url: '',
  project_url: '',
})

const isEditing = computed(() => Number.isFinite(form.id) && form.id > 0)

function resetForm() {
  form.id = null
  form.title = ''
  form.description = ''
  form.tech_stack = ''
  form.image_url = ''
  form.project_url = ''
  formError.value = ''
}

function hydrateForm(project) {
  form.id = project.id
  form.title = project.title || ''
  form.description = project.description || ''
  form.tech_stack = project.tech_stack || ''
  form.image_url = project.image_url || ''
  form.project_url = project.project_url || ''
  formError.value = ''
}

async function loadProjects() {
  loading.value = true
  pageError.value = ''

  try {
    projects.value = await getAdminProjects()
  } catch (error) {
    pageError.value = error?.response?.data?.message || error?.message || 'Gagal memuat project.'
    if (error?.response?.status === 401) {
      adminStore.logout()
      await router.push('/admin/login')
    }
  } finally {
    loading.value = false
  }
}

async function submitForm() {
  formError.value = ''
  saving.value = true

  const payload = {
    title: form.title.trim(),
    description: form.description.trim(),
    tech_stack: form.tech_stack.trim(),
    image_url: form.image_url.trim(),
    project_url: form.project_url.trim(),
  }

  try {
    if (isEditing.value) {
      await updateAdminProject(form.id, payload)
    } else {
      await createAdminProject(payload)
    }

    await loadProjects()
    resetForm()
  } catch (error) {
    const serverMessage = error?.response?.data?.error || error?.response?.data?.message
    formError.value = serverMessage || error?.message || 'Gagal menyimpan project.'
    if (error?.response?.status === 401) {
      adminStore.logout()
      await router.push('/admin/login')
    }
  } finally {
    saving.value = false
  }
}

async function removeProject(projectId) {
  const confirmed = window.confirm('Hapus project ini?')
  if (!confirmed) {
    return
  }

  try {
    await deleteAdminProject(projectId)
    if (form.id === projectId) {
      resetForm()
    }
    await loadProjects()
  } catch (error) {
    pageError.value = error?.response?.data?.message || error?.message || 'Gagal menghapus project.'
    if (error?.response?.status === 401) {
      adminStore.logout()
      await router.push('/admin/login')
    }
  }
}

async function logout() {
  adminStore.logout()
  await router.push('/admin/login')
}

onMounted(async () => {
  if (!adminStore.isAuthenticated) {
    await router.push('/admin/login')
    return
  }

  await loadProjects()
})
</script>

<template>
  <section class="grid gap-4">
    <header class="flex flex-wrap items-center justify-between gap-3 rounded-lgx border border-border bg-surface p-4 shadow-soft">
      <div>
        <h1 class="text-[1.6rem]">Admin Project Manager</h1>
        <p class="text-muted">Kelola data project yang tampil di portfolio.</p>
      </div>
      <div class="flex items-center gap-2">
        <button
          type="button"
          class="rounded-xl border border-border bg-surface-soft px-3 py-2 text-sm font-semibold text-ink hover:border-brand"
          @click="loadProjects"
        >
          Refresh
        </button>
        <button
          type="button"
          class="rounded-xl border border-border bg-surface-soft px-3 py-2 text-sm font-semibold text-ink hover:border-brand"
          @click="logout"
        >
          Logout
        </button>
      </div>
    </header>

    <section class="grid gap-4 lg:grid-cols-[1.25fr_1fr]">
      <article class="rounded-lgx border border-border bg-surface p-4 shadow-soft">
        <div class="mb-3 flex items-center justify-between gap-2">
          <h2 class="text-[1.2rem]">Daftar Project</h2>
          <span class="rounded-full bg-brand-soft px-2 py-1 text-xs font-bold text-ink">{{ projects.length }} items</span>
        </div>

        <p
          v-if="pageError"
          class="mb-3 rounded-xl border border-dashed border-[#fdba74] bg-[#fff7ed] px-3 py-2 text-sm text-[#9a3412]"
        >
          {{ pageError }}
        </p>

        <p v-if="loading" class="text-sm text-muted">Loading project data...</p>

        <ul v-else-if="projects.length" class="grid gap-3">
          <li
            v-for="project in projects"
            :key="project.id"
            class="rounded-xl border border-border bg-surface-soft p-3"
          >
            <div class="flex flex-wrap items-start justify-between gap-2">
              <div>
                <p class="font-bold text-heading">{{ project.title }}</p>
                <p class="mt-1 text-sm text-muted">{{ project.description }}</p>
                <p v-if="project.tech_stack" class="mt-2 text-xs font-semibold text-ink">
                  Tech: {{ project.tech_stack }}
                </p>
              </div>
              <div class="flex items-center gap-2">
                <button
                  type="button"
                  class="rounded-lg border border-border px-2 py-1 text-xs font-semibold text-ink hover:border-brand"
                  @click="hydrateForm(project)"
                >
                  Edit
                </button>
                <button
                  type="button"
                  class="rounded-lg border border-[#fca5a5] px-2 py-1 text-xs font-semibold text-[#b91c1c] hover:bg-[#fef2f2]"
                  @click="removeProject(project.id)"
                >
                  Delete
                </button>
              </div>
            </div>
          </li>
        </ul>

        <p v-else class="text-sm text-muted">Belum ada project di database.</p>
      </article>

      <article class="rounded-lgx border border-border bg-surface p-4 shadow-soft">
        <div class="mb-3 flex items-center justify-between gap-2">
          <h2 class="text-[1.2rem]">{{ isEditing ? 'Edit Project' : 'Tambah Project' }}</h2>
          <button
            v-if="isEditing"
            type="button"
            class="rounded-lg border border-border px-2 py-1 text-xs font-semibold text-ink hover:border-brand"
            @click="resetForm"
          >
            Batal Edit
          </button>
        </div>

        <form class="grid gap-3" @submit.prevent="submitForm">
          <label class="grid gap-1 text-sm font-semibold text-ink">
            Title
            <input
              v-model="form.title"
              type="text"
              required
              class="rounded-xl border border-border bg-surface-soft px-3 py-2 text-ink focus:border-brand focus:outline-none focus:ring-2 focus:ring-focus"
            />
          </label>

          <label class="grid gap-1 text-sm font-semibold text-ink">
            Description
            <textarea
              v-model="form.description"
              rows="4"
              required
              class="rounded-xl border border-border bg-surface-soft px-3 py-2 text-ink focus:border-brand focus:outline-none focus:ring-2 focus:ring-focus"
            ></textarea>
          </label>

          <label class="grid gap-1 text-sm font-semibold text-ink">
            Tech Stack
            <input
              v-model="form.tech_stack"
              type="text"
              placeholder="Vue, Go, PostgreSQL"
              class="rounded-xl border border-border bg-surface-soft px-3 py-2 text-ink focus:border-brand focus:outline-none focus:ring-2 focus:ring-focus"
            />
          </label>

          <label class="grid gap-1 text-sm font-semibold text-ink">
            Image URL
            <input
              v-model="form.image_url"
              type="url"
              placeholder="https://..."
              class="rounded-xl border border-border bg-surface-soft px-3 py-2 text-ink focus:border-brand focus:outline-none focus:ring-2 focus:ring-focus"
            />
          </label>

          <label class="grid gap-1 text-sm font-semibold text-ink">
            Project URL
            <input
              v-model="form.project_url"
              type="url"
              placeholder="https://..."
              class="rounded-xl border border-border bg-surface-soft px-3 py-2 text-ink focus:border-brand focus:outline-none focus:ring-2 focus:ring-focus"
            />
          </label>

          <button
            type="submit"
            :disabled="saving"
            class="w-fit rounded-xl bg-gradient-to-br from-brand to-brand-deep px-4 py-2 text-sm font-bold text-white hover:-translate-y-px disabled:cursor-not-allowed disabled:opacity-70"
          >
            {{ saving ? 'Menyimpan...' : isEditing ? 'Update Project' : 'Tambah Project' }}
          </button>
        </form>

        <p
          v-if="formError"
          class="mt-3 rounded-xl border border-dashed border-[#fdba74] bg-[#fff7ed] px-3 py-2 text-sm text-[#9a3412]"
        >
          {{ formError }}
        </p>
      </article>
    </section>
  </section>
</template>
