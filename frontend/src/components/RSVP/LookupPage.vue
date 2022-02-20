<template>
  <form class="flex flex-col lg:w-1/2 mx-auto" @submit.prevent="submit">
    <div class="text-center pb-7">
      Enter your name and email to RSVP.<br />
      You can respond for more guests in the next steps.
    </div>

    <transition>
      <matrimony-alert v-if="error" class="mb-8" @dismiss="error = null">
        {{ error }}
      </matrimony-alert>
    </transition>

    <div class="flex flex-wrap">
      <input-field
        v-model="query.first"
        class="w-full sm:w-1/2 sm:pr-8"
        required
        autofocus
        :errors="v$.first.$errors"
        @input="v$.first.$touch"
        @blur="v$.first.$touch"
      >
        First Name
      </input-field>
      <input-field v-model="query.last" class="w-full sm:w-1/2">
        Last Name
      </input-field>
    </div>
    <input-field
      v-model="query.email"
      :errors="v$.email.$errors"
      @blur="v$.email.$touch"
    >
      Email Address
      <template #prefix-icon>
        <font-awesome-icon :icon="['fas', 'envelope']" />
      </template>
    </input-field>

    <matrimony-form-button type="submit" class="ml-auto mr-30" size="lg">
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
import { required, email } from "@vuelidate/validators";
import useVuelidate from "@vuelidate/core";
import InputField from "../Forms/InputField.vue";
import MatrimonyAlert from "../MatrimonyAlert.vue";
import MatrimonyFormButton from "../Forms/MatrimonyFormButton.vue";
import { ErrGeneric, ErrNoUserMatch } from "../../strings/strings";

const router = useRouter();
const store = useStore();

const query = computed(() => store.state.persistent.query);
const rules = {
  first: { required },
  email: { email },
};
const v$ = useVuelidate(rules, query);

const loading = ref(false);
const error = ref(null);
const submit = async () => {
  if (!(await v$.value.$validate())) return;

  loading.value = true;
  error.value = null;
  store.commit("setQuery", query.value);
  try {
    const j = await store.dispatch("fetchParty");
    if (j.ok) {
      await router.push("/rsvp/questions");
    } else {
      error.value = ErrNoUserMatch;
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
  name: "LookupPage",
};
</script>
