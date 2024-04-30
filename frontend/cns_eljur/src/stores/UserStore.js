import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', () => {
  const userData = {
    firstName: "User",
    lastName: "Anonymous",
    email: "None",
    
  }

  return { count, doubleCount, increment }
})
