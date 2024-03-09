<script setup lang="ts">
import { onMounted, ref } from "vue";
import { Transaction } from "../models/types";
import { getTxs } from "../utils/api";

const transactions = ref<Transaction[]>([]);

onMounted(async () => {
  const txs = await getTxs();
  transactions.value = txs;
});
</script>
<template>
  <div class="flex flex-col items-center">
    <h1 class="text-2xl font-bold mb-5">Transactions</h1>
    <table class="table-auto border">
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
