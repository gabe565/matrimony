<template>
  <div class="flex flex-col lg:w-1/2 mx-auto">
    <party-member-picker class="mb-12 w-full" />

    <offset-anchor id="questions" ref="questionScrollAnchor" />
    <transition>
      <rsvp-form v-if="questionsActive" />
    </transition>

    <transition>
      <matrimony-alert v-if="error" class="mb-8" @dismiss="error = null">
        {{ error }}
      </matrimony-alert>
    </transition>

    <transition>
      <matrimony-form-button
        v-if="!questionsActive"
        key="next-button"
        type="submit"
        class="mx-auto"
        size="lg"
        @click.prevent="onSubmit"
      >
        Submit RSVP
        <template #icon>
          <spinner-icon v-if="loading" class="animate-spin ml-2" />
          <save-icon v-else class="ml-2" />
        </template>
      </matrimony-form-button>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed, watch, nextTick } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import SaveIcon from "~icons/mdi/content-save";
import SpinnerIcon from "~icons/gg/spinner";
import MatrimonyAlert from "../MatrimonyAlert.vue";
import MatrimonyFormButton from "../Forms/MatrimonyFormButton.vue";
import PartyMemberPicker from "./Prompt/PartyMemberPicker.vue";
import { ErrGeneric } from "../../strings/strings";
import RsvpForm from "./RsvpForm.vue";
import OffsetAnchor from "../OffsetAnchor.vue";

const store = useStore();
const router = useRouter();

const party = computed(() => store.state.persistent.party);
const activeGuestId = computed(() => store.state.activeGuestId);
const questionsActive = computed(() => store.state.activeGuestId !== 0);
const headConfirmed = computed(() => {
  const headId = store.getters.partyHeadGuest?.id;
  return store.state.persistent.responseSaved[headId] === true;
});

if (!party.value.guests) {
  router.replace("/rsvp/begin");
}

const loading = ref(false);
const error = ref(null);

store.dispatch("fetchQuestions").catch((e) => {
  console.error(e);
  error.value = ErrGeneric;
});

const questionScrollAnchor = ref(null);
watch(activeGuestId, async () => {
  error.value = null;
  if (activeGuestId.value !== 0) {
    await nextTick();
    questionScrollAnchor.value.$el.scrollIntoView();
  }
});

const onSubmit = async () => {
  if (!headConfirmed.value) {
    error.value = "Please respond for at least the first party member.";
    return;
  }
  loading.value = true;
  error.value = null;
  try {
    await router.push("/rsvp/finish");
  } catch (err) {
    error.value = err;
  }
};
</script>

<script>
export default {
  name: "QuestionsPage",
};
</script>
