<template>
  <div class="container mx-auto lg:max-w-[1436px] flex flex-col pt-52 lg:pt-0">
    <section
      class="relative lg:mt-40 text-slate-600 dark:text-slate-400 bg-white dark:bg-slate-900 body-font rounded-t-2xl"
    >
      <font-awesome-icon
        v-if="route.path === '/rsvp/finish'"
        :icon="['fat', 'party-horn']"
        class="absolute left-0 top-0 p-20 text-[20em] opacity-[0.05]"
      />
      <div class="container px-5 py-24 mx-auto">
        <div class="text-center mb-4 text-slate-600 dark:text-slate-300">
          <h1
            class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-slate-200"
          >
            RSVP
          </h1>
          <h2 class="text-3xl font-script title-font">
            {{ partnerNames }}
          </h2>
          <div>
            {{ config.info?.location }}
          </div>
          <div class="mb-8">
            {{ prettyDate }}
          </div>
        </div>

        <router-view />
      </div>
    </section>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { useRoute } from "vue-router";
import { formatDate, FullDate } from "../util/formatDate";

const props = defineProps({
  config: { type: Object, default: () => ({}) },
});

const route = useRoute();

const partnerNames = computed(() =>
  props.config.partners?.map((p) => p.first).join(" & ")
);

const prettyDate = computed(() =>
  formatDate(props.config.info?.date, FullDate)
);
</script>

<script>
export default {
  name: "RSVP",
};
</script>
