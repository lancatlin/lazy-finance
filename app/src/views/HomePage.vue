<script setup lang="ts">
import { ref, reactive, onMounted, computed } from "vue";
import { Account, Template, defaultTemplate } from "../models/types";
import { getTemplates } from "../utils/api";
import { applyTemplate } from "../models/validate";
import { newTx } from "../utils/api";
import { useToast } from "vue-toast-notification";

const toast = useToast();

const name = ref<string>("");
const selectedTemplate = ref<Template>(defaultTemplate);

const templates = ref<Template[]>([]);

onMounted(async () => {
  templates.value = [defaultTemplate, ...(await getTemplates())];
  console.log(templates.value);
});

const accounts = reactive<Account[]>([
  { name: "", amount: 0, commodity: "" },
  { name: "", amount: 0, commodity: "" },
]);

function useAmountBinding(account: Account) {
  return computed({
    get: () => (!account.amount ? "" : account.amount.toString()),
    set: (newValue) => {
      account.amount = newValue === "" ? 0 : parseFloat(newValue);
    },
  });
}

const addAccount = () => {
  accounts.push({ name: "", amount: 0, commodity: "" });
};

const removeAccount = (index: number) => {
  accounts.splice(index, 1);
};

const clearForm = () => {
  accounts.splice(
    0,
    accounts.length,
    { name: "", amount: 0, commodity: "" },
    { name: "", amount: 0, commodity: "" }
  );
  name.value = "";
};

const onSubmit = async () => {
  try {
    const result = applyTemplate(
      {
        accounts,
        date: new Date(),
        name: name.value,
      },
      selectedTemplate.value
    );
    console.log(result);
    const tx = await newTx(result);
    console.log("Submitted successfully");
    console.log(tx);
    toast.success(`Transaction "${tx.name}" created`);
    clearForm();
  } catch (e) {
    console.error(e);
    toast.error(e as string); // Cast error to string to satisfy the type of errorMessage
    return;
  }
};
</script>

<template>
  <div class="container max-w-lg mx-auto px-3 mt-5">
    <div>
      <div class="mb-6 flex items-center">
        <label for="template" class="text-sm font-medium text-gray-900 mr-2"
          >Template:
        </label>
        <select
          v-model="selectedTemplate"
          id="template"
          class="bg-gray-50 border rounded p-2"
        >
          <option
            v-for="template in templates"
            :key="template.name"
            :value="template"
          >
            {{ template.name }}
          </option>
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
          :placeholder="selectedTemplate.name"
          class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline"
        />
      </div>

      <!-- Accounts-->
      <div
        class="flex -mx-3 mb-2 items-center"
        v-for="(account, index) in accounts"
        :key="index"
      >
        <div class="flex w-auto">
          <!-- Account name-->
          <div class="w-1/2 px-3 mb-6 md:mb-0">
            <label
              :for="`account-${index}-name`"
              class="block mb-2 text-md font-medium text-gray-900"
              >Account {{ index + 1 }}</label
            >
            <input
              type="text"
              :id="`account-${index}-name`"
              v-model="account.name"
              :placeholder="selectedTemplate?.accounts[index]?.name"
              class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline"
            />
          </div>

          <!-- Amount -->
          <div class="w-1/4 px-3 mb-6 md:mb-0">
            <label
              :for="`account-${index}-amount`"
              class="block mb-2 text-md font-medium text-gray-900"
              >Amount</label
            >
            <input
              type="number"
              :id="`account-${index}-amount`"
              v-model="useAmountBinding(account).value"
              :placeholder="
                selectedTemplate?.accounts[index]?.amount?.toString()
              "
              class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline"
            />
          </div>

          <!-- Commodity -->
          <div class="w-1/4 px-3 mb-6 md:mb-0">
            <label
              :for="`account-${index}-commodity`"
              class="block mb-2 text-md font-medium text-gray-900"
              >Commodity</label
            >
            <input
              type="text"
              :id="`account-${index}-commodity`"
              v-model="account.commodity"
              :placeholder="selectedTemplate?.accounts[index]?.commodity"
              class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline"
            />
          </div>
        </div>

        <!-- Remove button-->
        <div class="w-auto px-3">
          <button
            class="bg-gray-100 hover:bg-white border px-3 py-2 rounded shadow disabled:shadow-none disabled:bg-gray-200 mt-5 hover:white"
            v-if="accounts.length > 2"
            @click="removeAccount(index)"
          >
            X
          </button>
        </div>
      </div>

      <!-- Add account button -->
      <button
        @click="addAccount"
        class="bg-gray-100 hover:bg-white border px-3 py-2 rounded shadow disabled:shadow-none disabled:bg-gray-200 mt-5 hover:white"
      >
        Add Account
      </button>

      <div class="flex flex-row items-center mt-5">
        <button
          @click="onSubmit"
          class="block bg-gray-100 hover:bg-white border px-3 py-2 rounded shadow disabled:shadow-none disabled:bg-gray-200 hover:white mr-2"
        >
          Submit
        </button>
        <input
          id="save"
          type="checkbox"
          class="w-5 h-5 mr-2 rounded focus:ring-blue-500"
        />
        <label for="save" class="mr-2">Save as Template</label>
        <!-- Submit -->
      </div>
    </div>
  </div>
</template>
