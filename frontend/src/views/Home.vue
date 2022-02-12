<script setup>
import { computed } from "vue";
import HeroSection from "../components/HeroSection.vue";
import ImageSection from "../components/ImageSection.vue";
import WeddingParty from "../components/WeddingParty.vue";
import TextSection from "../components/TextSection.vue";
import ScheduleSection from "../components/ScheduleSection.vue";
import FaqSection from "../components/FaqSection.vue";
import SpacerSection from "../components/SpacerSection.vue";

const props = defineProps({
  config: { type: Object, default: () => ({}) },
});

const sectionTypes = computed(() => ({
  hero: {
    component: HeroSection,
    bind: {
      title: props.config.info?.name,
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

const sections = computed(() => {
  const result = [];
  if (props.config.sections) {
    for (const [key, section] of Object.entries(props.config.sections)) {
      if (section.image && key > 2) {
        result.push({ spacer: {} });
      }

      result.push(section);

      if (section.image) {
        result.push({ spacer: {} });
      }
    }
  }
  return result;
});
</script>

<template>
  <template v-for="sectionEntry in sections">
    <template v-for="(section, type) in sectionEntry">
      <component
        :is="sectionTypes[type].component"
        v-bind="{ ...sectionTypes[type].bind, ...section }"
        :class="section.class"
      />
    </template>
  </template>
</template>
