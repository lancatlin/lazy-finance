<script setup lang="ts">
import { ref } from "vue";
import { Balance } from "../models/types";

const balances = ref<Balance[]>([
  { account: "assets", balance: "$-200" },
  { account: "expenses", balance: "$200" },
]);
</script>
<template>
  <div class="flex flex-col items-center">
    <h1 class="text-2xl mb-10">Balances</h1>
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
