<script setup>
import { computed } from "vue";
import HeroSection from "../components/HeroSection.vue";
import ImageSection from "../components/ImageSection.vue";
import WeddingParty from "../components/WeddingParty.vue";
import TextSection from "../components/TextSection.vue";
import ScheduleSection from "../components/ScheduleSection.vue";

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
}));
</script>

<template>
  <template v-for="sectionEntry in config.sections">
    <template v-for="(section, type) in sectionEntry">
      <component
        :is="sectionTypes[type].component"
        v-bind="{ ...section, ...sectionTypes[type].bind }"
      />
    </template>
  </template>
</template>
