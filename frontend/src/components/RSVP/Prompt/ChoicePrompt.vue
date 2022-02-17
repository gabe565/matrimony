<template>
  <div>
    <div class="mb-1 text-sm font-medium text-slate-900 dark:text-slate-200">
      {{ question.prompt }}
    </div>
    <radio-button-group class="mb-4">
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
  readonly: { type: Boolean, default: false },
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
