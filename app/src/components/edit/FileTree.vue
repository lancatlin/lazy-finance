<template>
  <ul class="file-tree">
    <li v-for="(node, index) in nodes" :key="index">
      <div v-if="node.type === 'folder'">
        <span>{{ node.name }}</span>
        <FileTree :nodes="node.children" />
      </div>
      <div v-else>
        <span>{{ node.name }}</span>
      </div>
    </li>
  </ul>
</template>

<script setup lang="ts">
import { defineProps } from "vue";
import { Folder, File } from "../../models/types";

defineProps<{
  nodes: (Folder | File)[];
}>();
</script>

<style scoped>
.file-tree {
  list-style-type: none;
  padding-left: 20px;
}
</style>
