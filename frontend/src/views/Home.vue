<script setup>
import { computed } from "vue";
import HeroSection from "../components/HeroSection.vue";
import ImageSection from "../components/ImageSection.vue";
import WeddingParty from "../components/WeddingParty.vue";
import TextSection from "../components/TextSection.vue";
import ScheduleSection from "../components/ScheduleSection.vue";
import FaqSection from "../components/FaqSection.vue";
import SpacerSection from "../components/SpacerSection.vue";
import { toSlug } from "../util/toSlug";

const props = defineProps({
  config: { type: Object, default: () => ({}) },
});

const firstAnchor = computed(() => {
  if (props.config.sections) {
    for (const section of props.config.sections) {
      for (const e of Object.values(section)) {
        if (e.title && !e.hideFromNav) {
          return `/#${toSlug(e.title)}`;
        }
      }
    }
  }
  return "";
});

const sectionTypes = computed(() => ({
  hero: {
    component: HeroSection,
    bind: {
      title: props.config.info?.name,
      toAnchor: firstAnchor.value,
    },
  },
  image: {
    component: ImageSection,
  },
  text: {
    component: TextSection,
  },
  party: {
    component: WeddingParty,
  },
  schedule: {
    component: ScheduleSection,
  },
  faq: {
    component: FaqSection,
  },
  spacer: {
    component: SpacerSection,
  },
}));
</script>

<template>
  <template v-for="(sectionEntry, key) in config.sections">
    <spacer-section v-if="sectionEntry.image && key > 2" />
    <template v-for="(section, type) in sectionEntry">
      <component
        :is="sectionTypes[type].component"
        v-bind="{ ...sectionTypes[type].bind, ...section }"
        :section-key="key"
        :class="section.class"
      />
    </template>
    <spacer-section v-if="sectionEntry.image" />
  </template>
</template>
