<template>
  <div>
    <div class="mb-1 text-sm font-medium text-slate-900 dark:text-slate-200">
      {{ question.prompt }}
    </div>
    <radio-button-group class="mb-6">
      <big-radio-button
        v-model="value"
        name="attendance"
        value="yes"
        :disabled="disabled"
        :readonly="readonly"
      >
        {{ question.yes.answer }}
      </big-radio-button>
      <big-radio-button
        v-model="value"
        name="attendance"
        value="no"
        :disabled="disabled"
        :readonly="readonly"
      >
        {{ question.no.answer }}
      </big-radio-button>
    </radio-button-group>
  </div>
</template>

<script setup>
import { computed } from "vue";
import BigRadioButton from "../../Forms/BigRadioButton.vue";
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
</script>

<script>
export default {
  name: "AttendancePrompt",
};
</script>
