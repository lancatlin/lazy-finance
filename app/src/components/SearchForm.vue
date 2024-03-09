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
  <div class="w-full md:w-2/3 px-3">
    <input
      :value="modelValue.searchTerm"
      @input="(event) => handleInput(updateSearchTerm)(event)"
      type="text"
      placeholder="Search term"
      class="w-full shadow border rounded p-2 mb-3"
    />
    <div class="w-full flex flex-row gap-5 justify-center mb-5">
      <div class="">
        <label for="beginDate" class="text-gray-500">From</label>
        <input
          id="beginDate"
          :value="modelValue.beginDate"
          @input="(event) => handleInput(updateBeginDate)(event)"
          type="date"
          class="w-full shadow border rounded p-2"
        />
      </div>
      <div class="">
        <label for="endDate" class="text-gray-500">To</label>
        <input
          id="endDate"
          :value="modelValue.endDate"
          @input="(event) => handleInput(updateEndDate)(event)"
          type="date"
          class="w-full shadow border rounded p-2"
        />
      </div>
    </div>
  </div>
</template>
