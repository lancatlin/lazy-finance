<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { File } from "../models/types";
import FileTreeSidebar from "../components/edit/FileTreeSidebar.vue";
import EditForm from "../components/edit/EditForm.vue";
import { getFileList } from "../utils/api";
import { useRoute } from "vue-router";

const route = useRoute();
const filePath = computed(() => route.params.path as string);

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

onMounted(async () => {
  const fileList = await getFileList();
  files.value = fileList;
});
</script>
<template>
  <div class="relative md:flex md:flex-row h-full">
    <FileTreeSidebar :files="files" />
    <EditForm :filePath="filePath" />
  </div>
</template>
