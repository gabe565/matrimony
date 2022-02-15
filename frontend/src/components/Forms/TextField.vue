<template>
  <div
    class="mt-1 flex w-full"
    :class="[stacked ? 'flex-col' : 'items-center']"
  >
    <div class="whitespace-nowrap pr-2 text-slate-900 dark:text-slate-200">
      {{ label }}
    </div>
    <input
      v-model="value"
      :value="value"
      :type="type"
      class="w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50text-slate-900 dark:text-slate-200"
      :placeholder="placeholder"
    />
  </div>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  stacked: { type: Boolean, default: false },
  label: { type: String, default: "" },
  placeholder: { type: String, default: "" },
  type: { type: String, default: "text" },
  modelValue: { type: String, default: "" },
});

const emit = defineEmits(["update:modelValue"]);

const value = computed({
  get: () => props.modelValue,
  set: (v) => emit("update:modelValue", v),
});
</script>
