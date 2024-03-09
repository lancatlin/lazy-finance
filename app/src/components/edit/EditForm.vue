<script setup lang="ts">
import { ref } from "vue";

const code = ref(`console.log('Hello, world!')`);
</script>
<template>
  <div class="w-full">
    <textarea
      class="w-full p-2 rounded m-2 bg-slate-100"
      v-model="code"
    ></textarea>
  </div>
</template>
