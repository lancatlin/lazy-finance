<script setup lang="ts">
import { ref, reactive } from "vue";
import { Account, Template } from "../models/types";
import validate from "../models/validate";

const name = ref<string>("");
const templateId = ref<string>("default");
const errorMessage = ref<string>("");

const template: Template = {
  accounts: [
    { name: "expenses", commodity: "$", amount: 50 },
    { name: "assets:cash", commodity: "$", amount: null },
  ],
};

const accounts = reactive<Account[]>([
  { name: "", amount: null, commodity: "" },
  { name: "", amount: null, commodity: "" },
]);

const addAccount = () => {
  accounts.push({ name: "", amount: null, commodity: "" });
};

const removeAccount = (index: number) => {
  accounts.splice(index, 1);
};

const onSubmit = () => {
  try {
    const result = validate({
      accounts,
      name: name.value,
      template,
    });
    console.log(result);
    console.log("Submitted successfully");
    errorMessage.value = "";
  } catch (e) {
    console.error(e);
    errorMessage.value = e as string; // Cast error to string to satisfy the type of errorMessage
    return;
  }
};
</script>

<template>
  <div class="container max-w-lg mx-auto px-3">
    <h1 class="text-3xl underline mb-5">Lazy 累記</h1>
    <div>
      <div class="mb-6 flex items-center">
        <label for="template" class="text-sm font-medium text-gray-900 mr-2"
          >Template:
        </label>
        <select
          v-model="templateId"
          id="template"
          class="bg-gray-50 border rounded"
        >
          <option>default</option>
          <option>restaurant</option>
        </select>
      </div>
      <div class="mb-6">
        <label for="name" class="block mb-2 text-md font-medium text-gray-900"
          >Name</label
        >
        <input
          v-model="name"
          type="text"
          id="name"
          class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline"
        />
      </div>
      <div
        class="flex flex-wrap -mx-3 mb-2"
        v-for="(account, index) in accounts"
        :key="index"
      >
        <div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
          <label
            :for="`account-${index}-name`"
            class="block mb-2 text-md font-medium text-gray-900"
            >Account {{ index + 1 }}</label
          >
          <input
            type="text"
            :id="`account-${index}-name`"
            v-model="account.name"
            :placeholder="template.accounts[index]?.name"
            class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline"
          />
        </div>
        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <label
            :for="`account-${index}-amount`"
            class="block mb-2 text-md font-medium text-gray-900"
            >Amount</label
          >
          <input
            type="number"
            :id="`account-${index}-amount`"
            v-model="account.amount"
            :placeholder="template.accounts[index]?.amount?.toString()"
            class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline"
          />
        </div>

        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <label
            :for="`account-${index}-commodity`"
            class="block mb-2 text-md font-medium text-gray-900"
            >Commodity</label
          >
          <input
            type="text"
            :id="`account-${index}-commodity`"
            v-model="account.commodity"
            :placeholder="template.accounts[index]?.commodity"
            class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline"
          />
        </div>
        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <button
            class="block bg-gray-100 hover:bg-white border px-3 py-2 rounded shadow disabled:shadow-none disabled:bg-gray-200 mt-5 hover:white"
            v-if="accounts.length > 2"
            @click="removeAccount(index)"
          >
            X
          </button>
        </div>
      </div>
      <button
        @click="addAccount"
        class="block bg-gray-100 hover:bg-white border px-3 py-2 rounded shadow disabled:shadow-none disabled:bg-gray-200 mt-5 hover:white"
      >
        Add Account
      </button>
      <div
        v-if="errorMessage"
        class="mt-4 text-red-600 border border-red-600 p-3 rounded"
      >
        {{ errorMessage }}
      </div>
      <button
        @click="onSubmit"
        class="block bg-gray-100 hover:bg-white border px-3 py-2 rounded shadow disabled:shadow-none disabled:bg-gray-200 mt-5 hover:white"
      >
        Submit
      </button>
    </div>
  </div>
</template>
