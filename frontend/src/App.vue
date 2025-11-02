<template>
  <div id="app">
    <h1>Resumify</h1>
    <!-- Main resume builder component -->
    <ResumeBuilder />
  </div>
</template>

<script setup lang="ts">
import ResumeBuilder from './components/ResumeBuilder.vue'
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  margin: 2rem;
}
</style>
