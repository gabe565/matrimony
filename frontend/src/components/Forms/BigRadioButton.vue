<template>
  <label
    class="block font-medium text-gray-900 dark:text-gray-300 inline-flex items-center p-4 font-medium text-slate-900 bg-gray-50 rounded-xl border border-slate-200 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-slate-800 dark:text-slate-300 dark:border-slate-600 duration-100"
    :class="[
      disabled
        ? 'cursor-not-allowed'
        : 'cursor-pointer hover:bg-gray-100 hover:text-blue-700  dark:hover:text-white dark:hover:bg-slate-700',
    ]"
  >
    <input
      v-model="model"
      type="radio"
      :name="name"
      :value="value"
      class="w-4 h-4 mr-4"
      :required="required"
      :disabled="disabled"
      :readonly="readonly"
    />
    <slot />
  </label>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  name: { type: String, default: "" },
  value: { type: String, required: true },
  modelValue: { type: String, required: true },
  required: { type: Boolean, default: false },
  disabled: { type: Boolean, default: false },
  readonly: { type: Boolean, default: false },
});

const emit = defineEmits(["update:modelValue"]);

const model = computed({
  get: () => props.modelValue,
  set: (e) => emit("update:modelValue", e),
});
</script>

<script>
export default {
  name: "BigRadioButton",
};
</script>
