<!-- frontend/src/App.vue -->
<script setup lang="ts">
import ResumeBuilder from './components/ResumeBuilder.vue'
</script>

<template>
  <div id="app">
    <h1>Resumify</h1>
    <ResumeBuilder />
  </div>
</template>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  margin-top: 2rem;
}
</style>
