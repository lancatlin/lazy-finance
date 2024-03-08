<script setup lang="ts">
import { ref } from "vue";
import { Transaction } from "../models/types";

// convert iso date to Date
function convertDate(date: string) {
  return new Date(date);
}

const transactions = ref<Transaction[]>([
  {
    name: "test",
    date: convertDate("2021-01-01"),
    accounts: [
      {
        name: "expenses",
        amount: 100,
        commodity: "$",
      },
      {
        name: "assets",
        amount: -100,
        commodity: "$",
      },
    ],
  },
]);
</script>
<template>
  <div class="flex flex-col items-center">
    <h1 class="text-2xl font-bold">Transactions</h1>
    <table class="w-full md:w-2/3 table-auto border">
      <tr class="border bg-gray-400">
        <th class="px-5 py-3">Date</th>
        <th class="px-5 py-3">Name</th>
        <th class="px-5 py-3">Account</th>
        <th class="px-5 py-3">Amount</th>
      </tr>
      <template v-for="transaction in transactions" :key="transaction.date">
        <tr class="border">
          <td class="px-5 py-3" :rowspan="transaction.accounts.length">
            {{ transaction.date.toLocaleDateString() }}
          </td>
          <td class="px-5 py-3" :rowspan="transaction.accounts.length">
            {{ transaction.name }}
          </td>
          <td class="px-5 py-3">{{ transaction.accounts[0].name }}</td>
          <td class="px-5 py-3">{{ transaction.accounts[0].amount }}</td>
        </tr>
        <tr
          class="border"
          v-for="(account, index) in transaction.accounts.slice(1)"
          :key="`${transaction.date}-${index}`"
        >
          <td class="px-5 py-3">{{ account.name }}</td>
          <td class="px-5 py-3">{{ account.amount }}</td>
        </tr>
      </template>
    </table>
  </div>
</template>
