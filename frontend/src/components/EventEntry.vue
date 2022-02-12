<template>
  <span
    class="flex absolute -left-3 justify-center items-center w-6 h-6 bg-blue-200 rounded-full ring-8 ring-white dark:ring-gray-900 dark:bg-blue-900"
  >
    <font-awesome-icon
      :icon="['fas', 'calendar']"
      class="w-3 h-3 text-blue-600 dark:text-blue-400"
    />
  </span>
  <time
    class="mb-1 text-sm font-normal leading-none text-slate-400 dark:text-slate-500"
  >
    {{ startTime }}
  </time>
  <h3 class="text-lg font-semibold text-slate-900 dark:text-white">
    {{ title }}
  </h3>
  <p class="text-base font-normal text-slate-500 dark:text-slate-400">
    {{ description }}
  </p>
  <div
    v-if="dressCode"
    class="text-base font-normal text-slate-500 dark:text-slate-400"
  >
    <span class="font-semibold">Dress Code: </span> {{ dressCode }}
  </div>
  <template v-if="location">
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
    <matrimony-button
      :href="`https://maps.google.com/?daddr=${location}`"
      target="_blank"
      class="mt-4"
    >
      Navigate
      <template #icon>
        <font-awesome-icon :icon="['fas', 'diamond-turn-right']" class="ml-2" />
      </template>
    </matrimony-button>
  </template>
</template>

<script setup>
import { computed } from "vue";
import MatrimonyButton from "./MatrimonyButton.vue";

const props = defineProps({
  title: { type: String, default: "" },
  description: { type: String, default: "" },
  location: { type: String, default: "" },
  showMap: { type: Boolean, default: false },
  dressCode: { type: String, default: "" },
  start: { type: String, default: "" },
  end: { type: String, default: "" },
});

const formatDate = (str) => {
  if (!str) {
    return null;
  }
  const date = new Date(str);
  return date.toLocaleString("en-US", {
    weekday: "short",
    month: "long",
    day: "numeric",
    year: "numeric",
    hour: "numeric",
    minute: "numeric",
    hour12: true,
  });
};

const startTime = computed(() => formatDate(props.start));
</script>
