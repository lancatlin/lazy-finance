<script setup>
import { ref } from 'vue'

const accounts = ref([
  { name: '', amount: 0, commodity: '' },
  { name: '', amount: 0, commodity: '' },
])

const addAccount = () => {
  accounts.value.push({ name: '', amount: 0, commodity: '' })
}

const removeAccount = (index) => {
  accounts.value.splice(index, 1)
}
</script>

<template>
  <div class="container max-w-md mx-auto">
    <h1 class="text-3xl underline">Lazy 累記</h1>
    <p>Some dummy text</p>
    <div>
      <div class="mb-6">
        <label for="template" class="block mb-2 text-sm font-medium text-gray-900">Template </label>
        <select id="template" class="bg-gray-50 border rounded">
          <option>Option 1</option>
          <option>Option 2</option>
        </select>

      </div>
      <div class="mb-6">
        <label for="amount" class="block mb-2 text-md font-medium text-gray-900">Amount </label>
        <input type="number" id="amount"
          class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline" />
      </div>
      <div class="flex flex-wrap -mx-3 mb-2" v-for="account, index in accounts" v-key="index">
        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <label :for="`account-${index}-name`" class="block mb-2 text-md font-medium text-gray-900">Account {{ index
          }}</label>
          <input type="text" :id="`account-${index}-name`" v-model="account.name"
            class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline" />
        </div>
        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <label :for="`account-${index}-amount`" class="block mb-2 text-md font-medium text-gray-900">Amount</label>
          <input type="number" :id="`account-${index}-amount`" v-model="account.amount"
            class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline" />
        </div>

        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <label :for="`account-${index}-comodity`" class="block mb-2 text-md font-medium text-gray-900">Commodity</label>
          <input type="text" :id="`account-${index}-commodity`" v-model="account.commodity"
            class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline" />
        </div>
        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <button class="block border px-3 py-2 rounded shadow disabled:shadow-none disabled:bg-gray-200 mt-5 hover:white"
            :disabled="accounts.length <= 2" @click="removeAccount(index)">X</button>
        </div>
      </div>
      <button @click="addAccount">Add Account</button>
      <input type="text" id="from" />
      <br>
      <label for="memo">Memo </label>
      <input type="text" id="memo" />
      <br>
      <button class="block hover:color-gray">Submit</button>
    </div>
  </div>
</template>