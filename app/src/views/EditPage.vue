<script setup lang="ts">
import { ref } from "vue";
import { Folder } from "../models/types";
import FileTreeSidebar from "../components/FIleTreeSidebar.vue";

const folder = ref<Folder>({
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
});
</script>
<template>
  <div class="relative md:flex md:flex-row">
    <FileTreeSidebar />
    <div>
      <h1>Edit</h1>
      <p>Edit page</p>
    </div>
  </div>
</template>
