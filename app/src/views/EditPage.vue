<script setup lang="ts">
import { ref } from "vue";
import { File } from "../models/types";
import FileTreeSidebar from "../components/edit/FileTreeSidebar.vue";
import EditForm from "../components/edit/EditForm.vue";

const files = ref<File[]>([
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
]);
</script>
<template>
  <div class="relative md:flex md:flex-row h-full">
    <FileTreeSidebar :files="files" />
    <EditForm />
  </div>
</template>
