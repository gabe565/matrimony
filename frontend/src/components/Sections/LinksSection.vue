<template>
  <section
    class="text-slate-600 dark:text-slate-400 bg-slate-300 dark:bg-slate-700 body-font md:h-60"
  >
    <div
      class="container p-5 mx-auto flex flex-col h-full justify-center items-center"
    >
      <div class="text-center mb-4">
        <offset-anchor :id="toSlug(title)" />
        <section-title>{{ title }}</section-title>
      </div>
      <matrimony-button-group>
        <matrimony-button v-if="!moments.hidden" to="/moments" size="xl">
          {{ moments.content }}
          <template #icon>
            <gallery-icon class="ml-2" aria-hidden="true" />
          </template>
        </matrimony-button>
        <matrimony-button
          v-if="!rsvp.hidden"
          to="/rsvp"
          size="xl"
          color="green"
        >
          {{ rsvp.content }}
          <template #icon>
            <rsvp-icon class="ml-2" aria-hidden="true" />
          </template>
        </matrimony-button>
      </matrimony-button-group>
    </div>
  </section>
</template>

<script setup>
import GalleryIcon from "~icons/solar/gallery-wide-bold";
import RsvpIcon from "~icons/solar/calendar-bold";
import { toSlug } from "../../util/toSlug";
import MatrimonyButton from "../MatrimonyButton.vue";
import MatrimonyButtonGroup from "../MatrimonyButtonGroup.vue";
import OffsetAnchor from "../OffsetAnchor.vue";
import SectionTitle from "../SectionTitle.vue";

defineProps({
  title: { type: String, default: "Links" },
  content: { type: String, default: "" },
  moments: { type: Object, default: () => ({}) },
  rsvp: { type: Object, default: () => ({}) },
});
</script>

<script>
export default {
  name: "LinksSection",
};
</script>
