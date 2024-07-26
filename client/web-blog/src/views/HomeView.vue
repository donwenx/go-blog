<template>
  <div class="home">
    home页面{{ test }}
    <button class="btn" @click="onclickLogin">跳转login</button>
    <a-space wrap>
      <a-button type="primary">Primary Button</a-button>
      <a-button>Default Button</a-button>
      <a-button type="dashed">Dashed Button</a-button>
      <a-button type="text">Text Button</a-button>
      <a-button type="link">Link Button</a-button>
    </a-space>
  </div>
</template>
<script setup lang="ts">
import { useUserStore } from '../stores/user'
import { useRouter } from 'vue-router'
const store = useUserStore()
const test = store.$state.test
const router = useRouter()

const onclickLogin = () => {
  router.push({ name: 'login' })
}
</script>

<style scoped lang="scss">
.home {
  .btn {
    background-color: #66ccff;
  }
}
</style>