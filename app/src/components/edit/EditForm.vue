<script setup lang="ts">
import { ref, watch } from "vue";
import { getFileContent } from "../../utils/api";

const code = ref(`console.log('Hello, world!')`);
defineEmits({});

const props = defineProps({
  filePath: {
    type: String,
    required: true,
  },
});

watch(
  () => props.filePath,
  async (newPath, oldPath) => {
    if (newPath === oldPath) return;
    code.value = await getFileContent(newPath);
  },
  { immediate: true }
);
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
