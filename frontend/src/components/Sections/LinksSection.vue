<template>
  <section
    class="text-slate-600 dark:text-slate-400 bg-slate-300 dark:bg-slate-700 body-font h-[10rem] md:h-[15rem]"
  >
    <div
      class="container p-5 mx-auto flex flex-col h-full justify-center items-center"
    >
      <div class="text-center mb-4">
        <span :id="toSlug(title)" class="invisible relative -top-[200px]" />
        <h1
          class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-white"
        >
          {{ title }}
        </h1>
      </div>
      <matrimony-button-group>
        <matrimony-button to="/moments" class="text-xl">
          {{ moments.content }}
          <template #icon>
            <font-awesome-icon :icon="['fas', 'grid-2']" class="ml-2" />
          </template>
        </matrimony-button>
        <matrimony-button
          to="/rsvp"
          class="text-xl bg-green-300 hover:!bg-green-200 dark:bg-green-900 dark:hover:!bg-green-800"
        >
          {{ rsvp.content }}
          <template #icon>
            <font-awesome-icon
              :icon="['fas', 'calendar-lines-pen']"
              class="ml-2"
            />
          </template>
        </matrimony-button>
      </matrimony-button-group>
    </div>
  </section>
</template>

<script setup>
import { toSlug } from "../../util/toSlug";
import MatrimonyButton from "../MatrimonyButton.vue";
import MatrimonyButtonGroup from "../MatrimonyButtonGroup.vue";

defineProps({
  title: { type: String, default: "Links" },
  content: { type: String, default: "" },
  moments: { type: Object, default: () => ({}) },
  rsvp: { type: Object, default: () => ({}) },
});
</script>
