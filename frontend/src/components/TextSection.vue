<template>
  <section
    class="text-slate-600 dark:text-slate-400 bg-white dark:bg-slate-900 body-font"
  >
    <div class="container p-5 mx-auto">
      <div class="text-center mb-4">
        <span :id="toSlug(title)" class="invisible relative -top-[200px]" />
        <h1
          class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-white"
        >
          {{ title }}
        </h1>
      </div>
      monkyyyyeyyyeyyyeyyy
      <matrimony-markdown :source="content" />
    </div>
  </section>
</template>

<script setup>
import { toSlug } from "../util/toSlug";
import MatrimonyMarkdown from "./MatrimonyMarkdown.vue";

defineProps({
  title: { type: String, default: "" },
  content: { type: String, default: "" },
});
</script>
