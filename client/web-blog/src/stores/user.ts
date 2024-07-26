import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
  state: () => ({
    test: 12345777,
  }),
  getters: {},
  actions: {},
})