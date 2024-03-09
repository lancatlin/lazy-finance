<script setup lang="ts">
import { defineProps, defineEmits } from "vue";

const props = defineProps({
  modelValue: {
    type: Object,
    required: true,
  },
});

const emit = defineEmits(["update:modelValue"]);

const handleInput = (callback: (value: string) => void) => (event: Event) => {
  callback((event.target as HTMLInputElement).value);
};

const updateSearchTerm = (searchTerm?: string) => {
  emit("update:modelValue", { ...props.modelValue, searchTerm });
};

const updateBeginDate = (beginDate?: string) => {
  emit("update:modelValue", { ...props.modelValue, beginDate });
};

const updateEndDate = (endDate?: string) => {
  emit("update:modelValue", { ...props.modelValue, endDate });
};
</script>

<template>
  <div class="w-full">
    <div>
      <input
        :value="modelValue.searchTerm"
        @input="(event) => handleInput(updateSearchTerm)(event)"
        type="text"
        placeholder="Search term"
      />
    </div>
    <input
      :value="modelValue.beginDate"
      @input="(event) => handleInput(updateBeginDate)(event)"
      type="date"
    />
    <input
      :value="modelValue.endDate"
      @input="(event) => handleInput(updateEndDate)(event)"
      type="date"
    />
  </div>
</template>
