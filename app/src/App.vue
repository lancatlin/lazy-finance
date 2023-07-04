<script setup>
import { ref, reactive } from 'vue'

const memo = ref('')
const templateId = ref('default')
const errorMessage = ref('')

const template = {
  accounts: [
    { name: 'expenses', commodity: '$', amount: 50 },
    { name: 'assets:cash', commodity: '$' },
  ]
}

const accounts = reactive([
  { name: '', amount: null, commodity: '' },
  { name: '', amount: null, commodity: '' },
])

const addAccount = () => {
  accounts.push({ name: '', amount: null, commodity: '' })
}

const removeAccount = (index) => {
  accounts.splice(index, 1)
}

const validateInput = () => {
  let nullCount = 0;
  let totalAmount = 0;
  const result = accounts.map((account, index) => {
    const output = {}
    output.name = account.name || template.accounts[index]?.name;
    output.amount = account.amount || template.accounts[index]?.amount || null;
    output.commodity = account.commodity || template.accounts[index]?.commodity;
    if (output.amount === null) nullCount++;
    else totalAmount += output.amount;
    if (!output.name) {
      throw `Account ${index + 1} cannot be empty.`
    }
    return output
  });
  console.log(result)

  if (nullCount > 1) {
    throw 'Only one account can have a empty amount.'
  }
  else if (nullCount === 0 && totalAmount !== 0) {
    throw 'The total amount of every account should be 0.'
  }
  return result;
}

const onSubmit = () => {
  try {
    const result = validateInput()
    console.log(result)
    console.log('Submitted successfully')
    errorMessage.value = ''
  } catch (e) {
    errorMessage.value = e
    return
  }
}
</script>

<template>
  <div class="container max-w-lg mx-auto px-3">
    <h1 class="text-3xl underline">Lazy 累記</h1>
    <p>Some dummy text</p>
    <div>
      <div class="mb-6">
        <label for="template" class="block mb-2 text-sm font-medium text-gray-900">Template </label>
        <select v-model="templateId" id="template" class="bg-gray-50 border rounded">
          <option>default</option>
          <option>restarant</option>
        </select>

      </div>
      <div class="flex flex-wrap -mx-3 mb-2" v-for="account, index in accounts" v-key="index">
        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <label :for="`account-${index}-name`" class="block mb-2 text-md font-medium text-gray-900">Account {{ index + 1
          }}</label>
          <input type="text" :id="`account-${index}-name`" v-model="account.name"
            :placeholder="template.accounts[index]?.name"
            class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline" />
        </div>
        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <label :for="`account-${index}-amount`" class="block mb-2 text-md font-medium text-gray-900">Amount</label>
          <input type="number" :id="`account-${index}-amount`" v-model="account.amount"
            :placeholder="template.accounts[index]?.amount"
            class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline" />
        </div>

        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <label :for="`account-${index}-comodity`" class="block mb-2 text-md font-medium text-gray-900">Commodity</label>
          <input type="text" :id="`account-${index}-commodity`" v-model="account.commodity"
            :placeholder="template.accounts[index]?.commodity"
            class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline" />
        </div>
        <div class="w-full md:w-1/4 px-3 mb-6 md:mb-0">
          <button
            class="block bg-gray-100 hover:bg-white border px-3 py-2 rounded shadow disabled:shadow-none disabled:bg-gray-200 mt-5 hover:white"
            v-if="accounts.length > 2" @click="removeAccount(index)">X</button>
        </div>
      </div>
      <button @click="addAccount"
        class="block bg-gray-100 hover:bg-white border px-3 py-2 rounded shadow disabled:shadow-none disabled:bg-gray-200 mt-5 hover:white">Add
        Account</button>
      <input type="text" id="from" />
      <div class="mb-6">
        <label for="memo" class="block mb-2 text-md font-medium text-gray-900">Memo</label>
        <input v-model="memo" type="text" id="memo"
          class="shadow bg-gray-50 border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline" />
      </div>
      <div v-if="errorMessage" class="mt-4 text-red-600 border border-red-600 p-3 rounded">
        {{ errorMessage }}
      </div>
      <button @click="onSubmit"
        class="block bg-gray-100 hover:bg-white border px-3 py-2 rounded shadow disabled:shadow-none disabled:bg-gray-200 mt-5 hover:white">Submit</button>
    </div>
  </div>
</template>