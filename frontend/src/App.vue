<script setup>
import { computed, ref } from "vue";
import AppNav from "./components/AppNav.vue";
import PatternBackground from "./components/PatternBackground.vue";
import HeroSection from "./components/HeroSection.vue";
import ImageSection from "./components/ImageSection.vue";
import WeddingParty from "./components/WeddingParty.vue";
import TextSection from "./components/TextSection.vue";
import ScheduleSection from "./components/ScheduleSection.vue";

const loading = ref(true);
const config = ref({});
fetch("/api/config").then(async (r) => {
  config.value = await r.json();
  loading.value = false;
  if (config.value.info?.name) {
    document.title = config.value.info.name;
  }
});

const sections = computed(() => ({
  hero: {
    component: HeroSection,
    bind: {
      title: config.value.info?.name,
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

const sectionTitles = computed(() =>
  config.value.sections
    ?.map((sectionEntry) =>
      Object.values(sectionEntry)
        .filter((s) => s.title)
        .map((s) => s.title)
    )
    .flat()
);
</script>

<template>
  <pattern-background />
  <app-nav
    :partners="config.partners"
    :loading="loading"
    :pages="sectionTitles"
  />
  <template v-for="sectionEntry in config.sections">
    <template v-for="(section, type) in sectionEntry">
      <component
        :is="sections[type].component"
        v-bind="{ ...section, ...sections[type].bind }"
      />
    </template>
  </template>
</template>
