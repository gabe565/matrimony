<template>
  <div
    class="p-4 max-w-xl mx-auto bg-white rounded-lg border shadow-md sm:p-8 dark:bg-gray-800 dark:border-gray-700 mb-10 duration-300"
  >
    <h4
      class="text-xl font-bold leading-none text-gray-900 dark:text-white mb-6"
    >
      Answering for {{ guestName }}
    </h4>

    <transition-group
      name="list"
      tag="form"
      class="relative"
      @submit.prevent="onSubmit"
    >
      <component
        :is="promptComponents[type]?.component"
        v-bind="promptComponents[type]?.bind"
        v-for="[type, question] in questions"
        :key="question.field"
        v-model="responses[question.field]"
        :question="question"
        class="pb-3 duration-300 transition-opacity w-full"
        :required="question.required"
        :errors="v$[question.field].$errors"
        @input="v$[question.field].$touch"
        @blur="v$[question.field].$touch"
      />

      <matrimony-alert
        v-if="error"
        key="alert"
        class="mb-8"
        @dismiss="error = null"
      >
        {{ error }}
      </matrimony-alert>

      <div key="save-button" class="w-full flex">
        <matrimony-form-button type="submit" class="mx-auto" size="lg">
          Save
          <template #icon>
            <font-awesome-icon
              :icon="
                loading ? ['fad', 'spinner-third'] : ['fas', 'floppy-disk']
              "
              :class="{ 'animate-spin': loading }"
              class="ml-2 fa-fw"
            />
          </template>
        </matrimony-form-button>
      </div>
    </transition-group>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import { useStore } from "vuex";
import useVuelidate from "@vuelidate/core";
import { required } from "@vuelidate/validators";
import MatrimonyFormButton from "../Forms/MatrimonyFormButton.vue";
import AttendancePrompt from "./Prompt/AttendancePrompt.vue";
import InputPrompt from "./Prompt/InputPrompt.vue";
import ChoicePrompt from "./Prompt/ChoicePrompt.vue";
import { ErrGeneric } from "../../strings/strings";
import MatrimonyAlert from "../MatrimonyAlert.vue";

const store = useStore();
const questions = computed(() => store.getters.visibleQuestions);
const rules = computed(() => {
  const r = {};
  for (const [, question] of questions.value) {
    r[question?.field] = {};
    if (question.required) {
      r[question?.field].required = required;
    }
  }
  return r;
});
const activeGuestId = computed(() => store.state.activeGuestId);
const responses = computed(
  () => store.state.persistent.responses[activeGuestId.value] || {}
);
const v$ = useVuelidate(rules, responses);

const guestName = computed(() => {
  const guest = store.getters.activeGuest;
  if (guest) {
    return `${guest.first} ${guest.last}`;
  }
  return "";
});

const promptComponents = {
  attendance: { component: AttendancePrompt, bind: {} },
  text: { component: InputPrompt, bind: {} },
  choice: { component: ChoicePrompt, bind: {} },
  number: { component: InputPrompt, bind: { type: "number" } },
};

const error = ref(null);
const loading = ref(false);
const onSubmit = async () => {
  if (!(await v$.value.$validate())) return;

  loading.value = true;
  error.value = null;
  try {
    const j = await store.dispatch("respond", responses.value);
    if (j.error) {
      error.value = j.error;
    } else {
      store.commit("setActiveGuestId", 0);
    }
  } catch (err) {
    console.error(err);
    error.value = ErrGeneric;
  } finally {
    loading.value = false;
  }
};
</script>

<script>
export default {
  name: "RsvpForm",
};
</script>
