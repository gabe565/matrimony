<template>
  <section class="text-slate-600 bg-white dark:bg-slate-900 body-font">
    <div class="container p-5 mx-auto max-w-prose">
      <div class="text-center mb-4">
        <span :id="toSlug(title)" class="invisible relative -top-[200px]" />
        <h1
          class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-white"
        >
          {{ title }}
        </h1>
      </div>
      <ol class="relative border-l border-slate-200 dark:border-slate-700">
        <li v-for="(event, key) in events" class="mb-10 ml-6">
          <event-entry
            v-bind="event"
            :section-key="sectionKey"
            :event-key="key"
          />
        </li>
      </ol>
    </div>
  </section>
</template>

<script setup>
import EventEntry from "../EventEntry.vue";
import { toSlug } from "../../util/toSlug";

defineProps({
  title: { type: String, default: "" },
  events: { type: Array, default: () => [] },
  sectionKey: { type: Number, default: 0 },
});
</script>
