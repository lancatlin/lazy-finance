<template>
  <button
    @click="sidebarOpen = !sidebarOpen"
    class="md:hidden bg-slate-200 rounded shadow p-3 absolute top-2 right-2"
  >
    📂
  </button>
  <aside
    class="border md:block bg-slate-200 md:w-64 md:h-full"
    :class="{ hidden: !sidebarOpen, block: sidebarOpen }"
  >
    <FileTree :nodes="files" />
  </aside>
</template>

<script setup lang="ts">
import FileTree from "./FileTree.vue";
import { ref, PropType } from "vue";
import { File } from "../../models/types";

const sidebarOpen = ref(false);

defineProps({
  files: {
    type: Array as PropType<File[]>,
    required: true,
  },
});
</script>
