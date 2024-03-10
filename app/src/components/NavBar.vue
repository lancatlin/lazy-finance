<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { logout } from "../utils/api";

const menuOpen = ref(false);

const route = useRoute();
const router = useRouter();
const path = computed(() => route.path);

const toggleMenu = () => {
  menuOpen.value = !menuOpen.value;
};

async function onLogout() {
  await logout();
  router.push({ name: "SignIn" });
}
</script>
<template>
  <nav
    class="relative bg-white shadow p-3 flex flex-row flex-wrap items-center justify-between sm:justify-start"
  >
    <router-link to="/">
      <h1 class="text-xl">Lazy 累記</h1>
    </router-link>
    <!-- Mobile menu button-->
    <div class="block sm:hidden">
      <button
        type="button"
        class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
        aria-controls="mobile-menu"
        aria-expanded="false"
        @click="toggleMenu"
      >
        <span class="sr-only">Open main menu</span>
        <!-- Icon when menu is closed. -->
        <svg
          class="block h-6 w-6"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          aria-hidden="true"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16m-7 6h7"
          />
        </svg>
      </button>
    </div>
    <div
      class="w-full flex flex-col items-center sm:w-auto sm:flex-row sm:block sm:ml-6 bg-white"
      :class="{
        hidden: !menuOpen,
      }"
    >
      <!-- Primary navigation links -->
      <router-link
        to="/"
        class="text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm"
        :class="{ 'font-bold': path === '/' }"
        >Home</router-link
      >
      <router-link
        to="/edit"
        class="text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm"
        :class="{ 'font-bold': path.startsWith('/edit') }"
        active-class="text-gray-900"
        >Edit</router-link
      >
      <router-link
        to="/transactions"
        class="text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm"
        :class="{ 'font-bold': path.startsWith('/transactions') }"
        >Transactions</router-link
      >
      <router-link
        to="/balances"
        class="text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm"
        :class="{ 'font-bold': path.startsWith('/balances') }"
        >Balances</router-link
      >
      <button
        @click="onLogout"
        class="text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm"
      >
        Logout
      </button>
    </div>
  </nav>
</template>
