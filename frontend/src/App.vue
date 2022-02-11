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
  <span id="top" class="invisible relative" />
  <pattern-background />
  <app-nav
    :partners="config.partners"
    :loading="loading"
    :pages="sectionTitles"
  />
  <router-view :config="config" />
</template>
