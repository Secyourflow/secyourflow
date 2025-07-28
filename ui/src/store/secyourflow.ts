import { useLocalStorage } from '@vueuse/core'
import { defineStore } from 'pinia'
import type { Ref } from 'vue'

type State = {
  sidebarCollapsed: Ref<boolean>
}

export const useSecyourflowStore = defineStore('secyourflow', {
  state: (): State => ({
    sidebarCollapsed: useLocalStorage('sidebarCollapsed', false)
  }),
  actions: {
    toggleSidebar() {
      this.sidebarCollapsed = !this.sidebarCollapsed
    }
  }
})
