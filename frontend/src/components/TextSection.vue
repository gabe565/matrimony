<template>
  <section
    class="text-slate-600 dark:text-slate-400 bg-white dark:bg-slate-900 body-font"
  >
    <div class="container px-5 py-24 mx-auto">
      <div class="text-center mb-4">
        <span :id="toSlug(title)" class="invisible relative -top-[200px]" />
        <h1
          class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-white"
        >
          {{ title }}
        </h1>
      </div>
      <div class="content">
        <vue-markdown :source="content" />
      </div>
    </div>
  </section>
</template>

<script setup>
import VueMarkdown from "vue-markdown-render";
import { toSlug } from "../util/toSlug";

defineProps({
  title: { type: String, default: "" },
  content: { type: String, default: "" },
});
</script>

<style scoped lang="scss">
.content {
  :deep(p) {
    padding-bottom: 1.2rem;
  }
  :deep(a) {
    text-decoration: underline;
  }
}
</style>
