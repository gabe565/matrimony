<template>
  <div
    v-if="toAnchor"
    class="container absolute top-0 left-1/2 -translate-x-1/2 h-screen"
    aria-hidden="true"
  >
    <router-link :to="toAnchor" class="block w-full h-full" />
  </div>

  <div
    class="sticky top-0 -z-10 container mx-auto pb-2 lg:pb-6 h-screen flex flex-col justify-center text-center pt-20 text-accent-50"
  >
    <div class="h-[40rem] lg:h-[48rem] w-full relative">
      <!-- Image -->
      <!-- TODO: Clean up mobile hero image -->
      <template v-if="image.src">
        <lazy-image
          class="hidden md:block h-full w-full bg-cover object-cover lg:rounded-[1rem]"
          bg-color="bg-slate-500"
          :src="image.src"
          :alt="alt"
        />
        <lazy-image
          class="md:hidden h-full w-full bg-cover object-cover lg:rounded-[1rem]"
          bg-color="bg-slate-500"
          :src="image.src.replace('.jpg', '-mobile.jpg')"
          :alt="alt"
        />
      </template>

      <!-- Gradient -->
      <div
        class="absolute bottom-0 left-0 h-2/3 w-full bg-gradient-to-b from-transparent to-black lg:rounded-[1rem]"
      />

      <!-- Content -->
      <div class="absolute bottom-0 left-0 w-full mx-auto flex flex-col">
        <h1 class="pb-4 font-script text-5xl md:text-7xl whitespace-nowrap">
          {{ greeting }}
        </h1>
        <div v-if="location" class="pb-2 md:text-xl">
          {{ location }}
        </div>
        <template v-if="date">
          <div class="pb-2 md:text-xl">{{ prettyDate }}</div>
          <date-countdown class="pb-8 md:text-xl" :date="date" />
        </template>

        <div class="pb-8 text-xl">
          <arrow-down-icon class="inline motion-safe:animate-bounce" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import ArrowDownIcon from "~icons/solar/alt-arrow-down-linear";
import DateCountdown from "../DateCountdown.vue";
import { formatDate, FullDate } from "../../util/formatDate";
import LazyImage from "../LazyImage.vue";

const props = defineProps({
  alt: { type: String, default: "" },
  image: { type: Object, default: () => ({}) },
  toAnchor: { type: String, default: "" },
  greeting: { type: String, default: "" },
  location: { type: String, default: "" },
  date: { type: String, default: "" },
});

const prettyDate = computed(() => formatDate(props.date, FullDate));
</script>

<script>
export default {
  name: "HeroSection",
};
</script>
