<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <header class="bg-gray-50 py-10 border-b border-gray-200">
    <div class="wrapper max-w-3xl mx-auto text-center px-4">
      <!-- Título -->
      <h1 class="text-3xl font-extrabold text-gray-900">
        Stock exercise! :)
      </h1>
      <!-- Descripción -->
      <p class="mt-3 text-lg text-gray-600">
        Sample implementation of Vue3 + Typescript + Tailwindcss + Go/Gin + CockroachDB.
      </p>
      <p class="mt-3 text-lg text-gray-600">Made by Angel009</p>
    </div>
  </header>
<!--main space-->
  <RouterView />
</template>
