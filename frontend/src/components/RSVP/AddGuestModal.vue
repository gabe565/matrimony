<template>
  <matrimony-modal>
    <template #title>Add Guest</template>

    <template #activator="{ on, attrs }">
      <matrimony-button v-bind="attrs" @click.prevent="on.click">
        <template #prefix-icon>
          <font-awesome-icon :icon="['fas', 'plus']" class="text-xl" />
        </template>
        Add a Guest
      </matrimony-button>
    </template>

    <template #body="{ toggle }">
      <form ref="form" class="flex flex-wrap" @submit.prevent="add(toggle)">
        <transition>
          <matrimony-alert v-if="error" class="mb-8" @dismiss="error = null">
            {{ error }}
          </matrimony-alert>
        </transition>

        <text-field
          v-model="newGuest.first"
          class="w-full sm:w-1/2 sm:pr-8"
          required
          autofocus
        >
          First Name
        </text-field>
        <text-field v-model="newGuest.last" class="w-full sm:w-1/2">
          Last Name
        </text-field>
        <input type="submit" class="hidden" />
      </form>
    </template>

    <template #footer="{ toggle }">
      <matrimony-button color="alt" @click.prevent="toggle">
        <template #prefix-icon>
          <font-awesome-icon :icon="['fas', 'xmark']" class="text-xl" />
        </template>
        Cancel
      </matrimony-button>
      <matrimony-button @click.prevent="add(toggle)">
        <template #prefix-icon>
          <font-awesome-icon :icon="['fas', 'plus']" class="text-xl" />
        </template>
        Add
      </matrimony-button>
    </template>
  </matrimony-modal>
</template>

<script setup>
import { ref } from "vue";
import { useStore } from "vuex";
import MatrimonyModal from "../MatrimonyModal.vue";
import MatrimonyButton from "../MatrimonyButton.vue";
import TextField from "../Forms/TextField.vue";
import MatrimonyAlert from "../MatrimonyAlert.vue";
import { ErrGeneric } from "../../strings/strings";

const store = useStore();

const newGuest = ref({});
const form = ref(null);
const error = ref(null);

const add = async (toggle) => {
  error.value = null;
  if (!form.value.reportValidity()) return;
  let r;
  try {
    r = await store.dispatch("createGuest", newGuest.value);
  } catch (err) {
    console.error(err);
    error.value = ErrGeneric;
    return;
  }
  if (r.ok) {
    toggle();
    newGuest.value = {};
  } else {
    error.value = ErrGeneric;
  }
};
</script>

<script>
export default {
  name: "AddGuestModal",
};
</script>
