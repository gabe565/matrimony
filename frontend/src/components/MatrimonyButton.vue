<template>
  <router-link v-if="to" :to="to">
    <slot />
    <slot name="icon" class="ml-2" />
  </router-link>
  <a v-else :href="href">
    <slot />
    <slot name="icon" class="ml-2" />
  </a>
</template>

<script setup>
defineProps({
  to: { type: String, default: "" },
  href: { type: String, default: "" },
});
</script>

<style scoped lang="scss">
a {
  @apply inline-flex items-center py-2 px-4 text-sm font-medium text-slate-900 bg-accent-100 rounded-lg border border-slate-200 hover:bg-slate-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-slate-800 dark:text-slate-300 dark:border-slate-600 dark:hover:text-white dark:hover:bg-slate-700;
}
</style>
