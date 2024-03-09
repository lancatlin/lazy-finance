<script setup lang="ts">
import { ref } from "vue";

const code = ref(`console.log('Hello, world!')`);
defineEmits({});
</script>
<template>
  <div class="w-full h-full p-2">
    <textarea
      class="w-full h-full p-2 rounded bg-slate-100"
      rows="20"
      v-model="code"
    ></textarea>
    <div class="flex flex-row justify-end">
      <button
        class="border rounded shadow px-5 py-2 mx-2 bg-green-300 hover:bg-green-200"
      >
        Save
      </button>
      <button
        class="border rounded shadow px-5 py-2 mx-2 bg-red-300 hover:bg-red-200"
      >
        Reset
      </button>
    </div>
  </div>
</template>
