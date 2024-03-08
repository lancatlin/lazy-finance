<script setup lang="ts"></script>
<template>
  <nav class="bg-white shadow">
    <div class="max-w-7xl mx-auto px-2 sm:px-4 lg:px-8">
      <div class="relative flex items-center justify-between h-16">
        <div
          class="flex-1 flex items-center justify-center sm:items-stretch sm:justify-start"
        >
          <div class="flex-shrink-0 flex items-center">
            <router-link to="/">
              <h1 class="text-xl items-center">Lazy 累記</h1>
            </router-link>
            <!-- <img
              class="block lg:hidden h-8 w-auto"
              src="/img/logo.svg"
              alt="logo"
            /> -->
          </div>
          <div class="hidden sm:block sm:ml-6">
            <div class="flex space-x-4">
              <router-link
                to="/"
                class="text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium"
                active-class="text-gray-900"
                >Home</router-link
              >
              <router-link
                to="/edit"
                class="text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium"
                active-class="text-gray-900"
                >Edit</router-link
              >
              <router-link
                to="/transactions"
                class="text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium"
                active-class="text-gray-900"
                >Transactions</router-link
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>
