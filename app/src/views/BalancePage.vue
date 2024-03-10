<script setup lang="ts">
import { watch, ref } from "vue";
import { useToast } from "vue-toast-notification";
import { Balance, Query } from "../models/types";
import { getBalances } from "../utils/api";
import { debounce } from "../utils/debounce";
import SearchForm from "../components/SearchForm.vue";

const balances = ref<Balance[]>([]);

const toast = useToast();

const query = ref<Query>({
  keyword: "",
  begin: undefined,
  end: undefined,
});

const debouncedGetBalances = debounce(getBalances, 400);

watch(
  query,
  async (newCriteria: Query) => {
    try {
      const balancesData = await debouncedGetBalances(newCriteria);
      balances.value = balancesData;
    } catch (err) {
      console.error(err);
      toast.error(err as string);
    }
  },
  { immediate: true }
);
</script>
<template>
  <div class="flex flex-col items-center mt-5">
    <h1 class="text-2xl font-bold mb-5">Balances</h1>
    <SearchForm v-model="query" />
    <table class="table-auto border">
      <thead class="border bg-gray-400">
        <tr>
          <th class="px-5 py-3">Account</th>
          <th class="px-5 py-3">Balance</th>
        </tr>
      </thead>
      <tbody>
        <tr class="border" v-for="balance in balances" :key="balance.account">
          <td class="px-5 py-3">{{ balance.account }}</td>
          <td class="px-5 py-3">{{ balance.balance }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
