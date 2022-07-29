<template>
  <div class="relative overflow-hidden" :class="[bgColor]">
    <img
      v-if="src"
      v-lazy="{ src, lifecycle }"
      :alt="alt"
      class="h-full w-full duration-300"
      :class="[contentFit]"
    />
    <div
      v-if="loading || !src"
      class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2"
    >
      <font-awesome-icon
        v-if="loading"
        :icon="['fad', 'spinner-third']"
        class="ml-2 text-4xl dark:text-white animate-spin"
      />
      <font-awesome-icon
        v-else-if="fallbackIcon"
        :icon="fallbackIcon"
        class="ml-2"
        :class="fallbackClass"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const props = defineProps({
  src: { type: String, default: "" },
  alt: { type: String, default: "" },
  bgColor: { type: String, default: "bg-slate-100 dark:bg-slate-800" },
  contentFit: { type: String, default: "object-cover object-center" },
  fallbackIcon: { type: Array, default: () => [] },
  fallbackClass: {
    type: [String, Array, Object],
    default: "text-7xl dark:text-white",
  },
});

const loading = ref(props.src !== "");
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
