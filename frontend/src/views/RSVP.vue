<template>
  <div class="container mx-auto lg:max-w-[1436px] flex flex-col pt-52 lg:pt-0">
    <section
      class="lg:mt-40 text-slate-600 dark:text-slate-400 bg-white dark:bg-slate-900 body-font rounded-t-2xl"
    >
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
          <div>Enter your name and email to RSVP.</div>
        </div>
        <lookup-form />
      </div>
    </section>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import LookupForm from "../components/RSVP/LookupForm.vue";
import { formatDate, FullDate } from "../util/formatDate";

const props = defineProps({
  config: { type: Object, default: () => ({}) },
});

const partnerNames = computed(() =>
  props.config.partners?.map((p) => p.first).join(" & ")
);

const prettyDate = computed(() =>
  formatDate(props.config.info?.date, FullDate)
);

const questions = ref({});
fetch("/api/rsvp/questions").then(async (r) => {
  questions.value = await r.json();
});
</script>
