<script setup lang="ts">
import { defineProps } from 'vue'

// Accept fileUrl from parent
const props = defineProps<{
  fileUrl: string
}>()
</script>

<template>
  <div v-if="props.fileUrl">
    <a :href="props.fileUrl" download="resume.pdf">
      Download Resume
    </a>
  </div>
  <div v-else>
    <p>Your resume will appear here after generation.</p>
  </div>
</template>
