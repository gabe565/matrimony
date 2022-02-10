<template>
  <header
    ref="nav"
    class="fixed top-0 z-40 w-full text-accent-300 body-font duration-300 transition-colors transition-[padding]"
    :class="[atHero ? 'bg-primary-500/70 backdrop-blur-md' : 'pt-12']"
  >
    <div
      class="container mx-auto flex flex-wrap p-3 flex-col lg:flex-row items-center"
    >
      <a
        class="flex title-font text-5xl font-script items-center hover:text-accent-50 duration-200 ease-out mb-4 lg:mb-0 lg:pr-4"
        href="/"
      >
        <font-awesome-icon :icon="['fat', 'rings-wedding']" class="mr-3" />
        <span
          v-if="loading"
          class="w-40 h-5 bg-slate-400 animate-pulse rounded-full"
        />
        <span v-else class="flex items-center">
          <template v-for="e in title">
            <font-awesome-icon
              v-if="e.icon"
              :icon="['fas', e.icon]"
              class="mx-2 text-xl motion-safe:animate-pulse"
            />
            <span v-else>{{ e.text }}</span>
          </template>
        </span>
      </a>
      <nav
        class="lg:mr-auto lg:ml-4 lg:py-1 lg:pl-4 lg:border-l lg:border-slate-400 flex flex-wrap items-center justify-center"
      >
        <a
          v-for="page in pages"
          :key="page"
          class="px-2 lg:px-4 hover:text-accent-50 duration-200"
          :href="`#${toSlug(page)}`"
        >
          {{ page }}
        </a>
      </nav>
      <matrimony-button href="#">
        RSVP
        <template #icon>
          <font-awesome-icon :icon="['far', 'arrow-right']" class="ml-2" />
        </template>
      </matrimony-button>
    </div>
    <!-- eslint-disable vue/no-v-html -->
    <div
      class="fill-accent-500 opacity-20 absolute top-0 w-full lg:hidden"
      v-html="blueContourTablet"
    />
    <!-- eslint-enable vue/no-v-html -->
  </header>
</template>

<script setup>
import { computed, ref } from "vue";
import { throttle } from "../util/throttle";
import MatrimonyButton from "./MatrimonyButton.vue";
import { passiveEventHandlerSupported } from "../util/passiveEventHandlerSupported";
import { toSlug } from "../util/toSlug";
/* eslint-disable-next-line import/no-unresolved */
import blueContourTablet from "../assets/blue_contour/tablet.svg?raw";

const props = defineProps({
  partners: { type: Array, default: () => [] },
  pages: { type: Array, default: () => [] },
  loading: { type: Boolean, default: false },
});

const title = computed(() => {
  const result = [];
  for (const [key, partner] of props.partners.entries()) {
    result.push({ text: partner.first });
    if (key !== props.partners.length - 1) {
      result.push({ icon: "heart" });
    }
  }
  return result;
});

const nav = ref(null);
const calcAtHero = () => {
  return window.scrollY > window.innerHeight - (nav.value?.clientHeight || 0);
};
const atHero = ref(calcAtHero());
document.addEventListener(
  "scroll",
  throttle(() => {
    atHero.value = calcAtHero();
  }, 100),
  passiveEventHandlerSupported ? { passive: true } : false
);
</script>
