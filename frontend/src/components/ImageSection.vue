<template>
  <img
    :src="src"
    :alt="alt"
    class="h-[15rem] md:h-[30rem] w-full object-cover pointer-events-none"
  />
</template>

<script setup>
defineProps({
  src: { type: String, default: "" },
  alt: { type: String, default: "" },
});
</script>
