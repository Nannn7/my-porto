import { defineStore } from 'pinia'
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
