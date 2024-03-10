<script setup lang="ts">
import { ref, reactive } from "vue";
import { useToast } from "vue-toast-notification";
import { LoginRequest } from "../models/types";
import { signIn } from "../utils/api";
import { useRouter } from "vue-router";

const toast = useToast();
const loginRequest = reactive<LoginRequest>({
  email: "",
  password: "",
});
const router = useRouter();

async function onSubmit() {
  try {
    await signIn(loginRequest);
    toast.success("Signed in successfully");
    router.push({ name: "Home" });
  } catch (err) {
    console.log(err);
    toast.error(err as string);
  }
}
</script>
<template>
  <div class="max-w-sm mx-auto">
    <h1 class="text-2xl font-bold mb-5 text-center mt-5">Sign In</h1>
    <form @submit.prevent="onSubmit">
      <div class="mb-5 mx-2">
        <label for="email" class="block mb-2 text-md font-medium text-gray-900"
          >Email</label
        >
        <input
          id="email"
          type="text"
          required="true"
          v-model="loginRequest.email"
          class="w-full shadow bg-gray-50 border rounded py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline"
        />
      </div>
      <div class="mb-5 mx-2">
        <label
          for="password"
          class="block mb-2 text-md font-medium text-gray-900"
          >Password</label
        >
        <input
          id="password"
          type="password"
          required="true"
          v-model="loginRequest.password"
          class="w-full shadow bg-gray-50 border rounded py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:shadow-outline"
        />
      </div>
      <div class="mb-5 mx-2 text-center">
        <button
          type="submit"
          class="py-2 px-4 bg-blue-500 text-white rounded mx-auto hover:bg-blue-400"
        >
          Sign In
        </button>
      </div>
    </form>
  </div>
</template>
