<template>
  <section class="text-slate-600 bg-white dark:bg-slate-900 body-font">
    <div class="container px-5 py-24 mx-auto">
      <div class="text-center mb-4">
        <span :id="toSlug(title)" class="invisible relative -top-[200px]" />
        <h1
          class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-white"
        >
          {{ title }}
        </h1>
      </div>
      <ol class="relative border-l border-slate-200 dark:border-slate-700">
        <li v-for="event in events" class="mb-10 ml-4">
          <event-entry v-bind="event" />
        </li>
      </ol>
    </div>
  </section>
</template>

<script setup>
import EventEntry from "./EventEntry.vue";
import { toSlug } from "../util/toSlug";

defineProps({
  title: { type: String, default: "" },
  events: { type: Array, default: () => [] },
});
</script>
