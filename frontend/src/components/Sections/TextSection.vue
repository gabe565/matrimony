<template>
  <section
    class="text-slate-600 dark:text-slate-400 bg-white dark:bg-slate-900 body-font"
  >
    <div class="container p-5 mx-auto">
      <div class="text-center mb-4">
        <offset-anchor :id="toSlug(title)" />
        <h1
          class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-white"
        >
          {{ title }}
        </h1>
      </div>
      <matrimony-markdown :source="content" />
    </div>
  </section>
</template>

<script setup>
import { toSlug } from "../../util/toSlug";
import MatrimonyMarkdown from "../MatrimonyMarkdown.vue";
import OffsetAnchor from "../OffsetAnchor.vue";

defineProps({
  title: { type: String, default: "" },
  content: { type: String, default: "" },
});
</script>

<script>
export default {
  name: "TextSection",
};
</script>
