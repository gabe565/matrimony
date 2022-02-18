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
      <form class="flex flex-wrap" @submit.prevent="add(toggle)">
        <text-field
          v-model="newGuest.first"
          class="w-full sm:w-1/2 sm:pr-8"
          required
          autofocus
        >
          First Name
        </text-field>
        <text-field v-model="newGuest.last" class="w-full sm:w-1/2" required>
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

const store = useStore();

const newGuest = ref({});

const add = async (toggle) => {
  const r = await store.dispatch("createGuest", newGuest.value);
  if (r.ok) {
    toggle();
    newGuest.value = {};
  }
};
</script>

<script>
export default {
  name: "AddGuestModal",
};
</script>
