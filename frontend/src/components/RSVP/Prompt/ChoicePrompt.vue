<template>
  <div class="mb-4">
    <div
      v-if="question.prompt"
      class="mb-1 text-sm font-medium text-slate-900 dark:text-slate-200"
    >
      {{ question.prompt }}
      <span v-if="required" class="text-red-600"> *</span>
    </div>
    <radio-button-group>
      <big-radio-button
        v-for="choice in question.choices"
        v-model="value"
        :name="id"
        :value="choice.answer"
        :disabled="disabled"
        :readonly="readonly"
      >
        {{ choice.answer }}
      </big-radio-button>
    </radio-button-group>
    <div v-for="error in errors" class="text-red-300 dark:text-red-600">
      {{ error.$message }}
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import BigRadioButton from "../../Forms/BigRadioButton.vue";
import { uniqueId } from "../../../util/uniqueId";
import RadioButtonGroup from "./RadioButtonGroup.vue";

const props = defineProps({
  question: { type: Object, required: true },
  modelValue: { type: String, default: "" },
  disabled: { type: Boolean, default: false },
  required: { type: Boolean, default: false },
  readonly: { type: Boolean, default: false },
  errors: { type: Array, default: () => [] },
});

const emit = defineEmits(["update:modelValue"]);

const value = computed({
  get: () => props.modelValue,
  set: (e) => emit("update:modelValue", e),
});

const id = uniqueId();
</script>

<script>
export default {
  name: "ChoicePrompt",
};
</script>
