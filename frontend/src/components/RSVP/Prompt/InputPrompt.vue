<template>
  <input-field v-model="value" class="mb-6">
    {{ question.prompt }}
  </input-field>
</template>

<script setup>
import { computed } from "vue";
import InputField from "../../Forms/InputField.vue";

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
