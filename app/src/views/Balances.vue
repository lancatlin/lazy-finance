<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useToast } from "vue-toast-notification";
import { Balance } from "../models/types";
import { getBalances } from "../utils/api";

const balances = ref<Balance[]>([]);

const toast = useToast();

onMounted(async () => {
  try {
    const balancesData = await getBalances();
    balances.value = balancesData;
  } catch (err) {
    console.error(err);
    toast.error(err as string);
  }
});
</script>
<template>
  <div class="flex flex-col items-center">
    <h1 class="text-2xl font-bold mb-5">Balances</h1>
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
