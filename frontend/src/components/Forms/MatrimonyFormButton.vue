<template>
  <button
    :type="type"
    class="inline-flex items-center py-2 px-4 font-medium text-sm text-slate-900 bg-accent-100 rounded-lg border border-slate-200 hover:bg-accent-50 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-slate-800 dark:text-slate-300 dark:border-slate-600 dark:hover:text-white dark:hover:bg-slate-700 duration-100"
  >
    <slot />
    <slot name="icon" class="ml-2" />
  </button>
</template>

<script setup>
defineProps({
  type: { type: String, default: "submit" },
});
</script>
