<template>
  <header
    ref="nav"
    class="fixed top-0 z-30 w-full text-accent-300 body-font duration-300 transition-colors transition-[padding]"
    :class="[atHero ? 'pt-12' : 'bg-primary-500/70 backdrop-blur-md']"
  >
    <!-- eslint-disable vue/no-v-html -->
    <div
      class="fill-accent-500 opacity-20 absolute top-0 w-full lg:hidden -z-10"
      v-html="blueContourTablet"
    />
    <!-- eslint-enable vue/no-v-html -->
    <div
      class="container mx-auto flex flex-wrap p-3 flex-col lg:flex-row items-center"
    >
      <router-link
        class="flex title-font lg:text-5xl font-script items-center hover:text-accent-50 duration-200 ease-out lg:mb-0 lg:pr-4"
        :class="[atHero ? 'text-5xl mb-4' : 'text-3xl mb-2']"
        to="/#top"
      >
        <rings-wedding-icon
          class="lg:mr-3 w-[1em] h-[1em]"
          :class="[atHero ? 'mr-3' : 'mr-2']"
        />
        <span
          v-if="loading"
          class="w-40 h-5 bg-slate-400 animate-pulse rounded-full"
        />
        <span v-else class="flex items-center">
          <template v-for="e in title">
            <component
              :is="e.icon"
              v-if="e.icon"
              class="lg:mx-2 lg:text-lg motion-safe:animate-pulse"
              :class="[atHero ? 'text-lg mx-2' : 'text-sm mx-1.5']"
            />
            <span v-else>{{ e.text }}</span>
          </template>
        </span>
      </router-link>
      <nav
        class="lg:mr-auto lg:ml-4 lg:py-1 lg:pl-4 lg:border-l lg:border-slate-400 flex flex-wrap items-center justify-center pb-3"
      >
        <router-link
          v-for="page in pages"
          :key="page"
          class="px-2 lg:px-4 hover:text-accent-50 duration-200"
          :to="`/#${toSlug(page)}`"
        >
          {{ page }}
        </router-link>
      </nav>
      <transition>
        <matrimony-button
          v-show="!route.path.startsWith('/rsvp')"
          to="/rsvp"
          class="mb-0"
        >
          RSVP
          <template #icon>
            <rsvp-icon class="ml-2" />
          </template>
        </matrimony-button>
      </transition>
    </div>
  </header>
</template>

<script setup>
import { computed, ref } from "vue";
import { useRoute } from "vue-router";
import RingsWeddingIcon from "~icons/matrimony/rings-wedding-thin";
import HeartIcon from "~icons/mdi/heart";
import RsvpIcon from "~icons/solar/calendar-bold";
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
      result.push({ icon: HeartIcon });
    }
  }
  return result;
});

const route = useRoute();
const nav = ref(null);
const calcAtHero = () => {
  let headerCollapse = 100;
  if (route.meta.headerCollapseY) {
    switch (typeof route.meta.headerCollapseY) {
      case "function":
        headerCollapse = route.meta.headerCollapseY();
        break;
      case "number":
        headerCollapse = route.meta.headerCollapseY;
        break;
      default:
        throw new Error(
          `Invalid headerCollapseY. Got ${route.meta.headerCollapseY}`
        );
    }
  }
  let scrollY;
  if (document.body.classList.contains("fixed")) {
    scrollY = parseInt(document.body.style.top || "0", 10) * -1;
  } else {
    scrollY = window.scrollY;
  }
  return scrollY < headerCollapse - (nav.value?.clientHeight || 0);
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

<script>
export default {
  name: "AppNav",
};
</script>
