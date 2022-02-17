<template>
  <section
    class="text-slate-600 dark:text-slate-400 bg-white dark:bg-slate-900 body-font"
  >
    <div class="container px-5 mx-auto">
      <div class="text-center mb-4">
        <offset-anchor :id="toSlug(title)" />
        <h1
          class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-white"
        >
          {{ title }}
        </h1>
      </div>
      <div class="masonry md:masonry-sm lg:masonry-md">
        <div v-for="item in items" class="py-3 break-inside">
          <div
            class="p-6 bg-white rounded-lg border border-gray-200 shadow-md dark:bg-gray-800 dark:border-gray-700"
          >
            <h5
              class="mb-2 text-xl font-bold tracking-tight text-gray-900 dark:text-white"
            >
              {{ item.title }}
            </h5>
            <div class="font-normal text-gray-700 dark:text-gray-400">
              <matrimony-markdown :source="item.content" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { toSlug } from "../../util/toSlug";
import MatrimonyMarkdown from "../MatrimonyMarkdown.vue";
import OffsetAnchor from "../OffsetAnchor.vue";

defineProps({
  title: { type: String, default: "" },
  items: { type: Array, default: () => [] },
});
</script>

<script>
export default {
  name: "FaqSection",
};
</script>
