<template>
  <div class="relative overflow-hidden" :class="[bgColor]">
    <img
      v-lazy="{ src, lifecycle }"
      :alt="alt"
      class="h-full w-full duration-300"
      :class="[contentFit]"
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
  bgColor: { type: String, default: "bg-slate-100 dark:bg-slate-800" },
  contentFit: { type: String, default: "object-cover object-center" },
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

<script>
export default {
  name: "LazyImage",
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
