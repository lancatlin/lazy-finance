<script setup lang="ts">
import { defineProps, defineEmits } from "vue";
import { Query } from "../models/types";

const props = defineProps({
  modelValue: {
    type: Object as () => Query,
    required: true,
  },
});

const emit = defineEmits(["update:modelValue"]);

const handleInput = (callback: (value: string) => void) => (event: Event) => {
  callback((event.target as HTMLInputElement).value);
};

const updateKeyword = (keyword?: string) => {
  emit("update:modelValue", { ...props.modelValue, keyword });
};

const updateBeginDate = (begin?: string) => {
  emit("update:modelValue", { ...props.modelValue, begin });
};

const updateEndDate = (end?: string) => {
  emit("update:modelValue", { ...props.modelValue, end });
};
</script>

<template>
  <div class="w-full md:w-2/3 px-3">
    <input
      :value="modelValue.keyword"
      @input="(event) => handleInput(updateKeyword)(event)"
      type="text"
      placeholder="Search term"
      class="w-full shadow border rounded p-2 mb-3"
    />
    <div class="w-full flex flex-row gap-5 justify-center mb-5">
      <div class="">
        <label for="beginDate" class="text-gray-500">From</label>
        <input
          id="beginDate"
          :value="modelValue.begin"
          @input="(event) => handleInput(updateBeginDate)(event)"
          type="date"
          class="w-full shadow border rounded p-2"
        />
      </div>
      <div class="">
        <label for="endDate" class="text-gray-500">To</label>
        <input
          id="endDate"
          :value="modelValue.end"
          @input="(event) => handleInput(updateEndDate)(event)"
          type="date"
          class="w-full shadow border rounded p-2"
        />
      </div>
    </div>
  </div>
</template>
