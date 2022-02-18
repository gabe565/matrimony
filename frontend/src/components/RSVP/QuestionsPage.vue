<template>
  <form class="flex flex-col lg:w-1/2 mx-auto" @submit.prevent="submit">
    <Transition>
      <matrimony-alert v-if="error" class="mb-8" @dismiss="error = null">
        {{ error }}
      </matrimony-alert>
    </Transition>

    <party-member-picker class="mb-12 w-full" />

    <Transition>
      <div
        v-if="hasActive"
        :key="guestName"
        class="p-4 max-w-xl mx-auto bg-white rounded-lg border shadow-md sm:p-8 dark:bg-gray-800 dark:border-gray-700 mb-10 duration-300"
      >
        <h4
          class="text-xl font-bold leading-none text-gray-900 dark:text-white mb-6"
        >
          Answering for {{ guestName }}
        </h4>
        <TransitionGroup name="list" tag="div" class="relative">
          <component
            :is="components[question[0]]"
            v-for="question in questions"
            :key="question[1].field"
            v-model="responses[question[1].field]"
            :question="question[1]"
            class="pb-3 duration-300 transition-opacity w-full"
          />
          <div key="next-button" class="w-full flex">
            <matrimony-form-button
              key="next-button"
              class="mx-auto"
              size="lg"
              :disabled="!completed"
              @click.prevent="save"
            >
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
        </TransitionGroup>
      </div>
    </Transition>

    <Transition>
      <matrimony-form-button
        v-if="allResponsesDone"
        key="next-button"
        class="mx-auto"
        size="lg"
        @click.prevent="finish"
      >
        Finish
        <template #icon>
          <font-awesome-icon
            :icon="loading ? ['fad', 'spinner-third'] : ['fas', 'check']"
            :class="{ 'animate-spin': loading }"
            class="ml-2 fa-fw"
          />
        </template>
      </matrimony-form-button>
    </Transition>
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

const guestName = computed(() => {
  const guest = store.getters.activeGuest;
  if (guest) {
    return `${guest.first} ${guest.last}`;
  }
  return "";
});
const party = computed(() => store.state.persistent.party);
const responses = computed(() => store.getters.activeResponse);
const allResponsesDone = computed(
  () =>
    Object.values(store.state.persistent.responseSaved).filter((e) => e)
      .length === Object.values(store.state.persistent.responses).length
);
const questions = computed(() => store.getters.visibleQuestions);
const hasActive = computed(() => store.state.persistent.activeId !== 0);

if (!party.value.guests) {
  router.replace("/rsvp/begin");
}

const loading = ref(false);
const error = ref(null);

store.dispatch("fetchQuestions").catch((e) => {
  console.error(e);
  error.value = "Something went wrong. Please try again later.";
});

const save = async () => {
  let r;
  try {
    r = await store.dispatch("respond", responses.value);
  } catch (err) {
    console.error(err);
    error.value = "Something went wrong! Please try again later.";
    return;
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

const completed = computed(() => {
  if (!responses.value) return false;
  for (const question of questions.value) {
    const { field } = question[1];
    if (!(field in responses.value) || responses.value[field] === "") {
      return false;
    }
  }
  return true;
});

const finish = async () => {
  await router.push("/rsvp/finish");
};
</script>

<script>
export default {
  name: "QuestionsPage",
};
</script>