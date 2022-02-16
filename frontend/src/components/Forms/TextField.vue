<template>
  <div class="mb-6">
    <label
      class="block mb-2 text-sm font-medium text-slate-900 dark:text-slate-200"
    >
      <slot />
      <span v-if="required" class="text-red-600">*</span>
      <input
        v-model="value"
        :type="type"
        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 shadow-sm"
        :placeholder="placeholder"
        :required="required"
      />
    </label>
  </div>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  stacked: { type: Boolean, default: false },
  placeholder: { type: String, default: "" },
  type: { type: String, default: "text" },
  modelValue: { type: String, default: "" },
  required: { type: Boolean, default: false },
});

const emit = defineEmits(["update:modelValue"]);

const value = computed({
  get: () => props.modelValue,
  set: (v) => emit("update:modelValue", v),
});
</script>
