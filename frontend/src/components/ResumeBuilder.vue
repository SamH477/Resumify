<!-- frontend/src/components/ResumeBuilder.vue -->
<script setup lang="ts">
import { ref } from 'vue'
import ResumeForm from './ResumeForm.vue'
import ResumeDownload from './ResumeDownload.vue'

const downloadUrl = ref('')

// Function that handles the data emitted from ResumeForm
async function handleSubmit(formData: { name: string; email: string; experience: string }) {
  try {
    const response = await fetch('http://localhost:8080/api/resume', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    })

    const blob = await response.blob()
    downloadUrl.value = URL.createObjectURL(blob)
  } catch (error) {
    console.error('Error generating resume:', error)
  }
}
</script>

<template>
  <div class="resume-builder">
    <ResumeForm @submit="handleSubmit" />
    <ResumeDownload :fileUrl="downloadUrl" />
  </div>
</template>

<style scoped>
.resume-builder {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
}
</style>
