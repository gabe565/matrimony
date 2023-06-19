<template>
  <div
    class="p-4 max-w-md mx-auto bg-white rounded-lg border shadow-md sm:p-8 dark:bg-gray-800 dark:border-gray-700"
  >
    <div class="flex justify-between items-center mb-4">
      <h3 class="text-xl font-bold leading-none text-gray-900 dark:text-white">
        Party Members
      </h3>
      <matrimony-button
        size="icon"
        class="text-xl"
        @click.prevent="editing = !editing"
      >
        <done-icon v-if="editing" />
        <users-edit-icon v-else />
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
              <!-- Responding button -->
              <matrimony-form-button
                v-if="active === guest.id"
                @click.prevent="active = 0"
              >
                Responding
                <template #icon>
                  <chevron-down-icon class="ml-0.5 fa-fw" />
                </template>
              </matrimony-form-button>

              <!-- Saved button -->
              <matrimony-form-button
                v-else-if="responseSaved[guest.id]"
                color="green"
                @click.prevent="active = guest.id"
              >
                Saved
                <template #icon>
                  <done-icon class="ml-0.5" />
                </template>
              </matrimony-form-button>

              <!-- Respond button -->
              <matrimony-form-button v-else @click.prevent="active = guest.id">
                Respond
                <template #icon>
                  <edit-icon class="ml-0.5 fa-fw" />
                </template>
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
import UsersEditIcon from "~icons/mdi/account-edit";
import DoneIcon from "~icons/mdi/check";
import ChevronDownIcon from "~icons/mdi/chevron-down";
import EditIcon from "~icons/mdi/pencil";
import MatrimonyFormButton from "../../Forms/MatrimonyFormButton.vue";
import MatrimonyButtonGroup from "../../MatrimonyButtonGroup.vue";
import MatrimonyButton from "../../MatrimonyButton.vue";
import AddGuestModal from "../AddGuestModal.vue";

const store = useStore();

const guests = computed(() => store.state.persistent.party.guests);
const responseSaved = computed(() => store.state.persistent.responseSaved);

const active = computed({
  get: () => store.state.activeGuestId,
  set: (e) => store.commit("setActiveGuestId", e),
});

const editing = ref(false);
</script>

<script>
export default {
  name: "PartyMemberPicker",
};
</script>
