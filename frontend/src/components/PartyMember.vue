<template>
  <div class="p-4 hover:scale-125 duration-300 ease-out">
    <div class="flex flex-shrink flex-col items-center text-center">
      <lazy-image
        :src="member.image"
        :alt="`${member.first} ${member.last} Photo`"
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
  </div>
</template>

<script setup>
import LazyImage from "./LazyImage.vue";

defineProps({
  member: { type: Object, default: () => ({}) },
});
</script>

<script>
export default {
  name: "PartyMember",
};
</script>
