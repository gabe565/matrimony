<template>
  <img
    :src="src"
    class="max-h-[40rem] w-full object-cover pointer-events-none"
  />
</template>

<script setup>
defineProps({
  src: { type: String, default: "" },
});
</script>
