<script setup lang="ts">
import { ref, defineEmits } from 'vue'

// Define the custom emit
const emit = defineEmits<{ (e: 'submit', payload: { name: string; email: string; experience: string }): void }>()

// Reactive form data
const form = ref({
  name: '',
  email: '',
  experience: ''
})

// Function called to check that all fields are filled in when form is submitted
function handleSubmit() {
  if (!form.value.name || !form.value.email || !form.value.experience) {
    alert('Please fill out all fields.')
    return
  }

  // Emit the data to be handled by the parent
  emit('submit', { ...form.value })
}
</script>

<template>
  <form @submit.prevent="handleSubmit">
    <div>
      <label for="name">Name:</label>
      <input id="name" type="text" v-model="form.name" placeholder="Your full name" />
    </div>

    <div>
      <label for="email">Email:</label>
      <input id="email" type="email" v-model="form.email" placeholder="Your email address" />
    </div>

    <div>
      <label for="experience">Experience:</label>
      <textarea id="experience" v-model="form.experience" placeholder="Brief description of your experience" rows="5"></textarea>
    </div>

    <button type="submit">Generate Resume</button>
  </form>
</template>
