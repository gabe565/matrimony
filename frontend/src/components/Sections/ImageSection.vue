<template>
  <lazy-image
    :src="`${API_URL}${src}`"
    :alt="alt"
    class="h-[20rem] md:h-[30rem] lg:h-[40rem] w-full object-cover pointer-events-none"
  />
</template>

<script setup>
import LazyImage from "../LazyImage.vue";
import { API_URL } from "../../config/api";

defineProps({
  src: { type: String, default: "" },
  alt: { type: String, default: "" },
});
</script>

<script>
export default {
  name: "ImageSection",
};
</script>
