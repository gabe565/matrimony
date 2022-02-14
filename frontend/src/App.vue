<template>
  <span id="top" class="invisible relative" />
  <pattern-background />
  <app-nav
    :partners="config.partners"
    :loading="loading"
    :pages="sectionTitles"
  />
  <router-view :config="config" />
</template>

<script setup>
import { computed, ref } from "vue";
import AppNav from "./components/AppNav.vue";
import PatternBackground from "./components/PatternBackground.vue";

const loading = ref(true);
const config = ref({});
fetch("/api/config").then(async (r) => {
  config.value = await r.json();
  loading.value = false;
  if (config.value.info?.name) {
    document.title = config.value.info.name;
  }
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
