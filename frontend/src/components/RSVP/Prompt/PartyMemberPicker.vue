<template>
  <div
    class="p-4 max-w-md mx-auto bg-white rounded-lg border shadow-md sm:p-8 dark:bg-gray-800 dark:border-gray-700"
  >
    <div class="flex justify-between items-center mb-4">
      <h3 class="text-xl font-bold leading-none text-gray-900 dark:text-white">
        Party Members
      </h3>
      <matrimony-button size="icon" @click.prevent="editing = !editing">
        <font-awesome-icon
          :icon="['fas', editing ? 'check' : 'pen-to-square']"
          class="text-xl"
        />
      </matrimony-button>
    </div>
    <div class="flow-root">
      <ul role="list" class="divide-y divide-gray-200 dark:divide-gray-700">
        <li v-for="guest in guests" class="py-3 sm:py-4">
          <div class="flex items-center space-x-4">
            <div class="flex-1 min-w-0">
              <p class="font-medium text-gray-900 truncate dark:text-white">
                {{ guest.first }} {{ guest.last }}
              </p>
            </div>
            <div
              class="inline-flex items-center text-base font-semibold text-gray-900 dark:text-white"
            >
              <matrimony-form-button
                type="button"
                :disabled="active === guest.id"
                @click.prevent="active = guest.id"
              >
                <template v-if="active === guest.id">Answering</template>
                <template v-else-if="responseSaved[guest.id]">Saved</template>
                <template v-else>Respond</template>
              </matrimony-form-button>
            </div>
          </div>
        </li>
      </ul>
    </div>
    <transition>
      <matrimony-button-group v-if="editing" class="duration-300">
        <add-guest-modal />
      </matrimony-button-group>
    </transition>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import { useStore } from "vuex";
import MatrimonyFormButton from "../../Forms/MatrimonyFormButton.vue";
import MatrimonyButtonGroup from "../../MatrimonyButtonGroup.vue";
import MatrimonyButton from "../../MatrimonyButton.vue";
import TextField from "../../Forms/TextField.vue";
import MatrimonyModal from "../../MatrimonyModal.vue";
import AddGuestModal from "../AddGuestModal.vue";

const store = useStore();

const guests = computed(() => store.state.persistent.party.guests);
const responseSaved = computed(() => store.state.persistent.responseSaved);

const active = computed({
  get: () => store.state.persistent.activeId,
  set: (e) => store.commit("active", e),
});

const editing = ref(false);
</script>

<script>
export default {
  name: "PartyMemberPicker",
};
</script>
