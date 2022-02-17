<template>
  <button :type="type" :class="classes">
    <slot name="prefix-icon" class="ml-2" />
    <slot />
    <slot name="icon" class="ml-2" />
  </button>
</template>

<script setup>
import { computed } from "vue";
import { base, colors, sizes } from "../../util/buttonClasses";

const props = defineProps({
  type: { type: String, default: "submit" },
  size: { type: String, default: "md" },
  color: { type: String, default: "blue" },
});

const classes = computed(() => {
  return [base, sizes[props.size], colors[props.color]];
});
</script>

<script>
export default {
  name: "MatrimonyFormButton",
};
</script>
