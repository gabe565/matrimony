<template>
  <matrimony-modal>
    <template #title>Add Guest</template>

    <template #activator="{ on, attrs }">
      <matrimony-button v-bind="attrs" @click.prevent="on.click">
        <template #prefix-icon>
          <add-icon class="pr-1 text-xl" aria-hidden="true" />
        </template>
        Add a Guest
      </matrimony-button>
    </template>

    <template #body="{ toggle }">
      <form
        ref="form"
        class="flex flex-wrap"
        @submit.prevent="onSubmit(toggle)"
      >
        <transition>
          <matrimony-alert v-if="error" class="mb-8" @dismiss="error = null">
            {{ error }}
          </matrimony-alert>
        </transition>

        <input-field
          v-model="newGuest.first"
          class="w-full sm:w-1/2 sm:pr-8"
          required
          autofocus
          :errors="v$.first.$errors"
          @input="v$.first.$touch"
          @blur="v$.first.$touch"
        >
          First Name
        </input-field>
        <input-field v-model="newGuest.last" class="w-full sm:w-1/2">
          Last Name
        </input-field>
        <input type="submit" class="hidden" />
      </form>
    </template>

    <template #footer="{ toggle }">
      <matrimony-button color="alt" @click.prevent="toggle">
        <template #prefix-icon>
          <close-icon class="text-xl pr-1" aria-hidden="true" />
        </template>
        Cancel
      </matrimony-button>
      <matrimony-button @click.prevent="onSubmit(toggle)">
        <template #prefix-icon>
          <spinner-icon
            v-if="loading"
            class="pr-1 animate-spin text-xl"
            aria-hidden="true"
          />
          <add-icon class="pr-1 text-xl" aria-hidden="true" />
        </template>
        Add
      </matrimony-button>
    </template>
  </matrimony-modal>
</template>

<script setup>
import { ref } from "vue";
import { useStore } from "vuex";
import { required } from "@vuelidate/validators";
import useVuelidate from "@vuelidate/core";
import AddIcon from "~icons/mdi/account-plus";
import SpinnerIcon from "~icons/gg/spinner";
import CloseIcon from "~icons/mdi/window-close";
import MatrimonyModal from "../MatrimonyModal.vue";
import MatrimonyButton from "../MatrimonyButton.vue";
import InputField from "../Forms/InputField.vue";
import MatrimonyAlert from "../MatrimonyAlert.vue";
import { ErrGeneric } from "../../strings/strings";

const store = useStore();

const newGuest = ref({});
const form = ref(null);
const error = ref(null);
const loading = ref(false);
const rules = {
  first: { required },
};
const v$ = useVuelidate(rules, newGuest);

const onSubmit = async (toggle) => {
  if (!(await v$.value.$validate())) return;

  loading.value = true;
  error.value = null;
  if (!form.value.reportValidity()) return;
  try {
    const j = await store.dispatch("createGuest", newGuest.value);
    if (j.error) {
      error.value = j.error;
    } else {
      toggle();
      newGuest.value = {};
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
  name: "AddGuestModal",
};
</script>
