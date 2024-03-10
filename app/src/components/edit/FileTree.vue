<template>
  <ul class="file-tree">
    <li v-for="(node, index) in nodes" :key="index">
      <RouterLink :to="`/edit/${node.name}`">
        <span>{{ node.name }}</span>
      </RouterLink>
      <div v-if="node.children">
        <FileTree :nodes="node.children" />
      </div>
    </li>
  </ul>
</template>

<script setup lang="ts">
import { defineProps } from "vue";
import { File } from "../../models/types";

defineProps<{
  nodes: File[];
}>();
</script>

<style scoped>
.file-tree {
  list-style-type: none;
  padding-left: 20px;
}
</style>
