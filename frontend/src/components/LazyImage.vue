<template>
  <div class="relative overflow-hidden" :class="[bgColor, darkBgColor]">
    <img
      v-lazy="{ src, lifecycle }"
      :alt="alt"
      class="object-cover object-center h-full w-full duration-300"
    />
    <div
      v-if="loading"
      class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2"
    >
      <font-awesome-icon
        :icon="['fad', 'spinner-third']"
        class="ml-2 text-4xl dark:text-white animate-spin"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

defineProps({
  src: { type: String, default: "" },
  alt: { type: String, default: "" },
  bgColor: { type: String, default: "bg-slate-100" },
  darkBgColor: { type: String, default: "dark:bg-slate-800" },
});

const loading = ref(true);
const lifecycle = {
  loading: () => {
    loading.value = true;
  },
  error: () => {
    loading.value = false;
  },
  loaded: () => {
    loading.value = false;
  },
};
</script>

<style scoped>
img[lazy="loading"] {
  @apply opacity-0;
}
img[lazy="loaded"] {
  @apply opacity-100;
}
</style>

<script>
export default {
  name: "LazyImage",
};
</script>
