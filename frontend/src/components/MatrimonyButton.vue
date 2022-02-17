<template>
  <router-link v-if="to" :to="to" :class="classes">
    <slot name="prefix-icon" class="mr-2" />
    <slot />
    <slot name="icon" class="ml-2" />
  </router-link>
  <a v-else :href="href" :class="classes">
    <slot name="prefix-icon" class="mr-2" />
    <slot />
    <slot name="icon" class="ml-2" />
  </a>
</template>

<script setup>
import { computed } from "vue";
import { base, colors, sizes } from "../util/buttonClasses";

const props = defineProps({
  to: { type: String, default: "" },
  href: { type: String, default: "" },
  size: { type: String, default: "md" },
  color: { type: String, default: "blue" },
  disabled: { type: Boolean, default: false },
});

const classes = computed(() => {
  return [base, sizes[props.size], colors[props.color]];
});
</script>

<script>
export default {
  name: "MatrimonyButton",
};
</script>
