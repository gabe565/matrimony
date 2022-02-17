<template>
  <text-field v-model="value" class="mb-6">
    {{ question.prompt }}
  </text-field>
</template>

<script setup>
import { computed } from "vue";
import TextField from "../../Forms/TextField.vue";

const props = defineProps({
  question: { type: Object, required: true },
  modelValue: { type: String, default: "" },
});

const emit = defineEmits(["update:modelValue"]);

const value = computed({
  get: () => props.modelValue,
  set: (e) => emit("update:modelValue", e),
});
</script>

<script>
export default {
  name: "TextPrompt",
};
</script>
