<template>
  <div class="p-4 hover:scale-125 duration-300 ease-out">
    <matrimony-modal>
      <template #activator="{ on }">
        <div
          class="flex flex-shrink flex-col items-center text-center cursor-pointer"
          v-on="on"
        >
          <lazy-image
            :src="member.image"
            :alt="`Photo of ${member.first} ${member.last}`"
            :fallback-icon="['fas', 'person']"
            fallback-class="text-6xl sm:text-[9em] text-slate-300 dark:text-slate-700"
            class="flex-shrink-0 rounded-lg h-40 sm:h-60 xl:h-72 w-40 sm:w-60 xl:w-72 mb-2 pointer-events-none"
          />
          <h2
            class="title-font font-medium text-lg text-slate-900 dark:text-slate-400"
          >
            {{ member.first }} {{ member.last }}
          </h2>
          <h3 class="text-slate-600 dark:text-slate-500 mb-3">
            {{ member.title }}
          </h3>
        </div>
      </template>

      <template #title>
        <div>{{ member.first }} {{ member.last }}</div>
        <h3 class="font-medium text-lg text-slate-600 dark:text-slate-500">
          {{ member.title }}
        </h3>
      </template>

      <template #body>
        <div class="flex flex-shrink flex-col md:flex-row items-center">
          <lazy-image
            v-if="member.image"
            :src="member.image"
            :alt="`Photo of ${member.first} ${member.last}`"
            class="flex-shrink-0 max-w-full rounded-lg h-80 w-80 mb-2 pointer-events-none"
          />

          <matrimony-markdown :source="textContent" class="px-5" />
        </div>
      </template>
    </matrimony-modal>
  </div>
</template>

<script setup>
import { computed } from "vue";
import LazyImage from "./LazyImage.vue";
import MatrimonyModal from "./MatrimonyModal.vue";
import MatrimonyMarkdown from "./MatrimonyMarkdown.vue";

const props = defineProps({
  member: { type: Object, default: () => ({}) },
});

const textContent = computed(() =>
  props.member.text?.content ? `> ${props.member.text.content}` : ""
);
</script>

<script>
export default {
  name: "PartyMember",
};
</script>
