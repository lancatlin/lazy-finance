<template>
  <button @click="sidebarOpen = !sidebarOpen" class="md:hidden">
    Toggle Sidebar
  </button>
  <aside
    class="border md:block"
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
