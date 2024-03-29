<template>
  <span
    class="flex absolute -left-3 justify-center items-center w-6 h-6 bg-blue-200 rounded-full ring-8 ring-white dark:ring-gray-900 dark:bg-blue-900"
  >
    <calendar-icon
      class="w-3 h-3 text-blue-600 dark:text-blue-400"
      aria-hidden="true"
    />
  </span>
  <h3 class="text-lg font-semibold text-slate-900 dark:text-white">
    {{ title }}
  </h3>
  <time
    class="mb-1 text-sm font-normal leading-none text-slate-400 dark:text-slate-500"
  >
    {{ startTime }}
  </time>
  <p class="text-base font-normal text-slate-500 dark:text-slate-400">
    {{ description }}
  </p>
  <div
    v-if="dressCode"
    class="text-base font-normal text-slate-500 dark:text-slate-400"
  >
    <span class="font-semibold">Dress Code: </span> {{ dressCode }}
  </div>
  <div
    v-if="location"
    class="text-base font-normal text-slate-500 dark:text-slate-400"
  >
    <span class="font-semibold">Address: </span>
    <a
      :href="`https://maps.google.com/?q=${location}`"
      target="_blank"
      class="underline"
    >
      {{ location }}
    </a>
  </div>
  <matrimony-button-group class="mt-4">
    <template v-for="button in buttons">
      <component
        :is="buttonTypes[button.type].component"
        v-if="buttonTypes[button.type]"
        v-bind="buttonTypes[button.type].bind"
      />
    </template>
  </matrimony-button-group>
</template>

<script setup>
import { computed } from "vue";
import CalendarIcon from "~icons/solar/calendar-minimalistic-bold";
import MatrimonyButtonGroup from "./MatrimonyButtonGroup.vue";
import NavigateButton from "./Schedule/NavigateButton.vue";
import CalendarButton from "./Schedule/CalendarButton.vue";
import { formatDate } from "../util/formatDate";

const props = defineProps({
  title: { type: String, default: "" },
  description: { type: String, default: "" },
  location: { type: String, default: "" },
  buttons: { type: Array, default: () => [] },
  dressCode: { type: String, default: "" },
  start: { type: String, default: "" },
  sectionKey: { type: Number, default: 0 },
  eventKey: { type: Number, default: 0 },
});

const buttonTypes = computed(() => ({
  calendar: {
    component: CalendarButton,
    bind: {
      sectionKey: props.sectionKey,
    },
  },
  navigate: {
    component: NavigateButton,
    bind: {
      location: props.location,
    },
  },
}));

const startTime = computed(() => formatDate(props.start));
</script>

<script>
export default {
  name: "EventEntry",
};
</script>
