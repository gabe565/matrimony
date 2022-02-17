<template>
  <form class="flex flex-col lg:w-1/2 mx-auto" @submit.prevent="submit">
    <Transition>
      <matrimony-alert v-if="error" class="mb-8" @dismiss="error = null">
        {{ error }}
      </matrimony-alert>
    </Transition>

    <party-member-picker class="mb-12 w-full" />

    <TransitionGroup name="list" tag="div" class="relative">
      <component
        :is="components[question[0]]"
        v-for="question in questions"
        :key="question[1].field"
        v-model="responses[question[1].field]"
        :question="question[1]"
        class="pb-3 duration-300 transition-opacity w-full"
      />

      <div v-if="showNext" key="next-button" class="w-full">
        <matrimony-form-button
          size="lg"
          class="float-right mr-30"
          @click.prevent="saveResponses"
        >
          Next
          <template #icon>
            <font-awesome-icon
              :icon="
                loading ? ['fad', 'spinner-third'] : ['far', 'arrow-right']
              "
              :class="{ 'animate-spin': loading }"
              class="ml-2 fa-fw"
            />
          </template>
        </matrimony-form-button>
      </div>
    </TransitionGroup>
  </form>
</template>

<script setup>
import { ref, computed } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import MatrimonyAlert from "../MatrimonyAlert.vue";
import MatrimonyFormButton from "../Forms/MatrimonyFormButton.vue";
import PartyMemberPicker from "./Prompt/PartyMemberPicker.vue";
import AttendancePrompt from "./Prompt/AttendancePrompt.vue";
import TextPrompt from "./Prompt/TextPrompt.vue";
import ChoicePrompt from "./Prompt/ChoicePrompt.vue";
import NumberPrompt from "./Prompt/NumberPrompt.vue";

const store = useStore();
const router = useRouter();

const party = computed(() => store.state.persistent.party);
const responses = computed(() => store.getters.activeResponse);
const questions = computed(() => store.getters.visibleQuestions);

if (!party.value.guests) {
  router.replace("/rsvp/begin");
}

const loading = ref(false);
const error = ref(null);

store.dispatch("fetchQuestions").catch((e) => {
  console.error(e);
  error.value = "Something went wrong. Please try again later.";
});

const saveResponses = async () => {
  let r;
  try {
    r = await store.dispatch("respond", responses.value);
  } catch (err) {
    error.value = "Something went wrong! Please try again later.";
  }
  if (!r.ok) {
    error.value = "Something went wrong! Please try again later.";
  }
};

const components = {
  attendance: AttendancePrompt,
  text: TextPrompt,
  choice: ChoicePrompt,
  number: NumberPrompt,
};

const showNext = computed(() => {
  if (!responses.value) return false;
  for (const question of questions.value) {
    const { field } = question[1];
    if (!(field in responses.value) || responses.value[field] === "") {
      return false;
    }
  }
  return true;
});
</script>

<script>
export default {
  name: "QuestionsPage",
};
</script>
