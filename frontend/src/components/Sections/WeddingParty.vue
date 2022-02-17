<template>
  <section
    class="text-slate-600 dark:text-white bg-white dark:bg-slate-900 body-font"
  >
    <div class="container p-5 mx-auto">
      <div class="text-center mb-4">
        <offset-anchor :id="toSlug(title)" />
        <section-title>{{ title }}</section-title>
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
import PartyMember from "../PartyMember.vue";
import OffsetAnchor from "../OffsetAnchor.vue";
import { toSlug } from "../../util/toSlug";
import SectionTitle from "../SectionTitle.vue";

defineProps({
  title: { type: String, default: "Wedding Party" },
  members: { type: Array, default: () => [] },
});
</script>

<script>
export default {
  name: "WeddingParty",
};
</script>
