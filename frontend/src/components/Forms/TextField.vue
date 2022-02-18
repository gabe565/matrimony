<template>
  <div class="mb-6">
    <label
      :for="id"
      class="block mb-1 text-sm font-medium text-slate-900 dark:text-slate-200"
    >
      <slot />
      <span v-if="required && $slots.default" class="text-red-600">*</span>
    </label>
    <div class="relative">
      <div
        v-if="$slots['prefix-icon']"
        class="flex absolute inset-y-0 left-0 items-center pl-3.5 pointer-events-none"
      >
        <slot name="prefix-icon" />
      </div>
      <input
        :id="id"
        ref="input"
        v-model="value"
        :type="type"
        class="bg-gray-50 border border-gray-300 text-gray-900 rounded-xl focus:ring-blue-500 focus:border-blue-500 block w-full dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 shadow-sm"
        :class="[
          ...classes,
          { 'pl-10': $slots['prefix-icon'], 'cursor-not-allowed': disabled },
        ]"
        :placeholder="placeholder"
        :required="required"
        :disabled="disabled"
        :readonly="readonly"
        :autofocus="autofocus"
      />
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from "vue";
import { uniqueId } from "../../util/uniqueId";

const props = defineProps({
  stacked: { type: Boolean, default: false },
  placeholder: { type: String, default: "" },
  type: { type: String, default: "text" },
  modelValue: { type: [String, Number], default: "" },
  required: { type: Boolean, default: false },
  size: { type: String, default: "lg" },
  disabled: { type: Boolean, default: false },
  readonly: { type: Boolean, default: false },
  autofocus: { type: Boolean, default: false },
});

const emit = defineEmits(["update:modelValue"]);

const value = computed({
  get: () => props.modelValue,
  set: (v) => emit("update:modelValue", v),
});

const id = `text-field-${uniqueId()}`;

const sizes = {
  sm: "p-2 text-xs",
  md: "p-2 text-sm",
  lg: "p-3 text-lg",
  xl: "p-4 text-xl",
};

const classes = computed(() => {
  return [sizes[props.size]];
});

const input = ref(null);
onMounted(() => {
  if (props.autofocus) {
    input.value.focus();
  }
});
</script>

<script>
export default {
  name: "TextField",
};
</script>
