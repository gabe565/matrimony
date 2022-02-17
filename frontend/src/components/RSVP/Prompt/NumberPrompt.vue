<template>
  <text-field v-model="value" type="number" class="mb-6">
    {{ question.field }}
  </text-field>
</template>

<script setup>
import { computed } from "vue";
import TextField from "../../Forms/TextField.vue";

const props = defineProps({
  question: { type: Object, required: true },
  modelValue: { type: [Number, String], default: 1 },
});

const emit = defineEmits(["update:modelValue"]);

const value = computed({
  get: () => props.modelValue,
  set: (e) => emit("update:modelValue", e),
});

emit("update:modelValue", 1);
</script>

<script>
export default {
  name: "NumberPrompt",
};
</script>
