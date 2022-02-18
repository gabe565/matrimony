<template>
  <form class="flex flex-col lg:w-1/2 mx-auto" @submit.prevent="submit">
    <div class="text-center pb-7">Enter your name and email to RSVP.</div>

    <Transition>
      <matrimony-alert v-if="error" class="mb-8" @dismiss="error = null">
        {{ error }}
      </matrimony-alert>
    </Transition>

    <div class="flex flex-wrap">
      <text-field
        v-model="query.first"
        class="w-full sm:w-1/2 sm:pr-8"
        required
        autofocus
      >
        First Name
      </text-field>
      <text-field v-model="query.last" class="w-full sm:w-1/2" required>
        Last Name
      </text-field>
    </div>
    <text-field v-model="query.email">
      Email Address
      <template #prefix-icon>
        <font-awesome-icon :icon="['fas', 'envelope']" />
      </template>
    </text-field>

    <div>
      <span class="font-bold">Tip:&nbsp;</span>
      <span v-if="error"
        >Check your spelling and try again. "David" vs "Dave"</span
      >
      <span v-else>You can respond for more guests in the next steps.</span>
    </div>

    <matrimony-form-button class="ml-auto mr-30" size="lg">
      Next
      <template #icon>
        <font-awesome-icon
          :icon="loading ? ['fad', 'spinner-third'] : ['far', 'arrow-right']"
          :class="{ 'animate-spin': loading }"
          class="ml-2 fa-fw"
        />
      </template>
    </matrimony-form-button>
  </form>
</template>

<script setup>
import { computed, ref } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import TextField from "../Forms/TextField.vue";
import MatrimonyAlert from "../MatrimonyAlert.vue";
import MatrimonyFormButton from "../Forms/MatrimonyFormButton.vue";

const router = useRouter();
const store = useStore();

const query = computed(() => store.state.persistent.query);

const loading = ref(false);
const error = ref(null);
const submit = async () => {
  loading.value = true;
  error.value = null;
  store.commit("setQuery", query.value);
  try {
    const j = await store.dispatch("fetchParty");
    if (j.ok) {
      await router.push("/rsvp/questions");
    } else {
      error.value = "No match was found.";
    }
  } catch (err) {
    console.error(err);
    error.value = "Something went wrong. Please try again later.";
  } finally {
    loading.value = false;
  }
};
</script>

<script>
export default {
  name: "LookupPage",
};
</script>
