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
  <div>
    <h1>Query</h1>
    <table class="table-auto border">
      <tr class="border">
        <th>Date</th>
        <th>Name</th>
        <th>Account</th>
        <th>Amount</th>
      </tr>
      <template v-for="transaction in transactions" :key="transaction.date">
        <tr class="border">
          <td :rowspan="transaction.accounts.length">
            {{ transaction.date.toLocaleDateString() }}
          </td>
          <td :rowspan="transaction.accounts.length">{{ transaction.name }}</td>
          <td>{{ transaction.accounts[0].name }}</td>
          <td>{{ transaction.accounts[0].amount }}</td>
        </tr>
        <tr
          class="border"
          v-for="(account, index) in transaction.accounts.slice(1)"
          :key="`${transaction.date}-${index}`"
        >
          <td>{{ account.name }}</td>
          <td>{{ account.amount }}</td>
        </tr>
      </template>
    </table>
  </div>
</template>
