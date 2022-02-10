<template>
  <img
    :src="src"
    :alt="alt"
    class="h-[40rem] w-full object-cover pointer-events-none"
  />
</template>

<script setup>
defineProps({
  src: { type: String, default: "" },
  alt: { type: String, default: "" },
});
</script>
