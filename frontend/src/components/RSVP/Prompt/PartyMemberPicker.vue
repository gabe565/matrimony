<template>
  <div
    class="p-4 max-w-md mx-auto bg-white rounded-lg border shadow-md sm:p-8 dark:bg-gray-800 dark:border-gray-700"
  >
    <div class="flex justify-between items-center mb-4">
      <h3 class="text-xl font-bold leading-none text-gray-900 dark:text-white">
        Party Members
      </h3>
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
                <template v-else>Respond</template>
              </matrimony-form-button>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { useStore } from "vuex";
import MatrimonyFormButton from "../../Forms/MatrimonyFormButton.vue";

const store = useStore();

const guests = computed(() => store.state.persistent.party.guests);

const active = computed({
  get: () => store.state.persistent.activeId,
  set: (e) => store.commit("active", e),
});
</script>

<script>
export default {
  name: "PartyMemberPicker",
};
</script>
