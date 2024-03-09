<script setup lang="ts">
import { watch, ref } from "vue";
import { Transaction, Query } from "../models/types";
import { getTxs } from "../utils/api";
import { debounce } from "../utils/debounce";
import SearchForm from "../components/SearchForm.vue";

const transactions = ref<Transaction[]>([]);
const searchCriteria = ref<Query>({
  keyword: "",
  begin: undefined,
  end: undefined,
});

const debouncedGetTxs = debounce(getTxs, 400);

watch(
  searchCriteria,
  async (newCriteria) => {
    const txs = await debouncedGetTxs(newCriteria);
    transactions.value = txs;
  },
  { immediate: true }
);
</script>

<template>
  <div class="flex flex-col items-center">
    <h1 class="text-2xl font-bold mb-5">Transactions</h1>
    <SearchForm v-model="searchCriteria" />
    <p>{{ searchCriteria }}</p>
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
