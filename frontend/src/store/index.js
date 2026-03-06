import { defineStore } from 'pinia'
<<<<<<< HEAD
import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'https://localhost:8080/api'

const defaultSkills = [
  'Golang',
  'Node.js',
  'Vue.js',
  'REST API Development',
  'MySQL & PostgreSQL',
  'Docker',
]

export const usePortfolioStore = defineStore('portfolio', {
  state: () => ({
    skills: [],
    projects: [],
    skillsLoaded: false,
    projectsLoaded: false,
    loadingProjects: false,
    projectsError: '',
  }),
  actions: {
    async fetchSkills(force = false) {
      if (this.skillsLoaded && !force) {
        return this.skills
      }

      this.skills = defaultSkills
      this.skillsLoaded = true
      return this.skills
    },
    async fetchProjects(force = false) {
      if (this.projectsLoaded && !force) {
        return this.projects
      }

      this.loadingProjects = true
      this.projectsError = ''

      try {
        const { data } = await axios.get(`${API_BASE_URL}/projects`)
        const items = Array.isArray(data?.projects)
          ? data.projects
          : Array.isArray(data)
            ? data
            : []

        this.projects = items
        this.projectsLoaded = true
      } catch (error) {
        this.projectsError = 'Gagal memuat projects dari backend.'
        this.projects = []
        this.projectsLoaded = false
      } finally {
        this.loadingProjects = false
      }

      return this.projects
    },
  },
})

export { API_BASE_URL }
=======
import { getProjects } from '../api/axios'

export const useProjectStore = defineStore('projects', {
  state: () => ({
    projects: [],
  }),
  actions: {
    async fetchProjects() {
      try {
        this.projects = await getProjects()
      } catch {
        this.projects = []
      }
    },
  },
})
>>>>>>> origin/main
