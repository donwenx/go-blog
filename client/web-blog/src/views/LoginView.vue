<template>
  <div>
    login页面{{ test }}
  </div>
</template>
<script setup lang="ts">
import { useUserStore } from '../stores/user'
const store = useUserStore()
const test = store.$state.test
</script>

<style></style>