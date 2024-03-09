<template>
  <NavBar />

  <main class="w-full md:w-2/3 mx-auto">
    <router-view class="py-5" />
  </main>
</template>

<script setup lang="ts">
import NavBar from "./components/NavBar.vue";
</script>
