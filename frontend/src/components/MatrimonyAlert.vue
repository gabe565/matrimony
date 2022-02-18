<template>
  <div
    class="flex w-full p-4 mb-4 bg-red-200 dark:bg-red-800 text-black dark:text-white rounded-lg duration-300 shadow-sm border border-red-300 dark:border-red-700 relative"
  >
    <div class="w-full">
      <slot />
    </div>
    <a href="#" @click.prevent="$emit('dismiss')">
      <font-awesome-icon :icon="['fas', 'xmark']" class="text-xl" />
    </a>
  </div>
</template>

<script setup>
defineEmits(["dismiss"]);
</script>

<script>
export default {
  name: "MatrimonyAlert",
};
</script>
