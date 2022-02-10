<template>
  <section
    class="text-slate-600 dark:text-white bg-white dark:bg-slate-900 body-font"
  >
    <div class="container px-5 py-24 mx-auto">
      <div class="text-center mb-4">
        <h1
          class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-white"
        >
          Wedding Party
        </h1>
      </div>
      <div class="grid grid-cols-2 mx-auto">
        <div
          v-for="(member, key) in members"
          :class="[key % 2 === 0 ? 'justify-self-end' : 'justify-self-start']"
        >
          <party-member
            :member="member"
            :class="[key % 2 === 0 ? 'origin-right' : 'origin-left']"
          />
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import PartyMember from "./PartyMember.vue";

defineProps({
  members: { type: Array, default: () => [] },
});
</script>
