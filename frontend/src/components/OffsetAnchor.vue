<template>
  <span
    :id="id"
    class="invisible relative"
    :class="offset"
    aria-hidden="true"
  />
</template>

<script setup>
defineProps({
  id: { type: String, required: true },
  offset: { type: String, default: "-top-44 lg:-top-[5.5rem]" },
});
</script>

<script>
export default {
  name: "OffsetAnchor",
};
</script>
