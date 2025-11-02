<script setup lang="ts">
import { ref, onMounted } from 'vue'

const message = ref('Loading...')

onMounted(async () => {
  const res = await fetch('/api/hello')
  const text = await res.text()
  message.value = text
})
</script>

<template>
  <h1>{{ message }}</h1>
</template>
