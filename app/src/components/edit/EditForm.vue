<script setup lang="ts">
import { ref, watch } from "vue";
import { getFileContent, saveFile } from "../../utils/api";
import { useToast } from "vue-toast-notification";

const toast = useToast();
const code = ref(`console.log('Hello, world!')`);
const fileChanged = ref<boolean>(false);
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
    reset();
  },
  { immediate: true }
);

async function reset() {
  try {
    code.value = await getFileContent(props.filePath);
    fileChanged.value = false;
  } catch (err) {
    toast.error("Failed to open file");
    console.error(err);
  }
}

function onChange() {
  fileChanged.value = true;
}

async function save() {
  try {
    await saveFile(props.filePath, code.value);
    toast.success("File saved!");
    fileChanged.value = false;
  } catch (err) {
    console.error(err);
    toast.error("Failed to save file");
  }
}
</script>
<template>
  <div class="w-full h-full p-2">
    <textarea
      class="w-full h-full p-2 rounded bg-slate-100"
      rows="20"
      v-model="code"
      @input="onChange"
    ></textarea>
    <div class="flex flex-row justify-end">
      <button
        class="border rounded shadow px-5 py-2 mx-2 bg-green-300 hover:bg-green-200 disabled:bg-gray-400"
        :disabled="!fileChanged"
        @click="save"
      >
        Save
      </button>
      <button
        class="border rounded shadow px-5 py-2 mx-2 bg-red-300 hover:bg-red-200"
        @click="reset"
      >
        Reset
      </button>
    </div>
  </div>
</template>
