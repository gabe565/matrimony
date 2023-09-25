<template>
  <footer
    class="container lg:max-w-[1436px] mx-auto p-4 bg-white shadow md:py-8 dark:bg-gray-800"
  >
    <span class="block text-sm text-gray-500 text-center dark:text-gray-400">
      Made with
      <heart-icon
        inline
        class="inline text-red-500 dark:text-red-700"
        aria-label="love"
      />
      by
      <a href="https://gabecook.com" class="underline" target="_blank"
        >gabe565</a
      >
    </span>
  </footer>
</template>

<script setup>
import HeartIcon from "~icons/mdi/heart";
</script>
