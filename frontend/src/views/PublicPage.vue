<template>
  <offset-anchor id="top" offset="" />
  <pattern-background />
  <app-nav
    :partners="config.partners"
    :loading="loading"
    :pages="sectionTitles"
    :show-rsvp="config.rsvp?.enabled"
  />
  <router-view :config="config" />
  <app-footer />
  <div id="modals" />
</template>

<script setup>
import { computed, ref } from "vue";
import AppNav from "../components/AppNav.vue";
import PatternBackground from "../components/PatternBackground.vue";
import AppFooter from "../components/AppFooter.vue";
import OffsetAnchor from "../components/OffsetAnchor.vue";

const loading = ref(true);
const config = ref({});
fetch("/api/config").then(async (r) => {
  config.value = await r.json();
  loading.value = false;
});

const sectionTitles = computed(() => {
  const result = [];
  if (config.value.sections) {
    for (const sectionEntry of config.value.sections) {
      for (const section of Object.values(sectionEntry)) {
        if (section.title && !section.hideFromNav) {
          result.push(section.title);
        }
      }
    }
  }
  return result;
});
</script>
