<template>
  <button
    @click="sidebarOpen = !sidebarOpen"
    class="md:hidden bg-slate-200 rounded shadow p-3 absolute top-2 right-2"
  >
    📂
  </button>
  <aside
    class="border md:block bg-slate-200 md:w-64 md:h-screen"
    :class="{ hidden: !sidebarOpen, block: sidebarOpen }"
  >
    <FileTree :nodes="folder" />
  </aside>
</template>

<script setup lang="ts">
import FileTree from "./FileTree.vue";
import { ref } from "vue";
import { Folder } from "../models/types";

const sidebarOpen = ref(false);

const folder = ref<Folder[]>([
  {
    name: "root",
    type: "folder",
    children: [
      {
        name: "journals",
        type: "folder",
        children: [
          {
            name: "2021.j",
            type: "file",
          },
          {
            name: "2022.j",
            type: "file",
          },
          {
            name: "2023.j",
            type: "file",
          },
        ],
      },
      {
        name: "settings.json",
        type: "file",
      },
    ],
  },
]);
</script>
