<template>
  <div class="container mx-auto lg:max-w-[1436px] flex flex-col pt-52 lg:pt-0">
    <section
      class="lg:mt-40 text-slate-600 dark:text-slate-400 bg-white dark:bg-slate-900 body-font rounded-t-2xl"
    >
      <div class="container px-5 py-24 mx-auto flex flex-col">
        <div class="text-center mb-4 -order-1">
          <h1
            class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-white"
          >
            Moments
          </h1>
        </div>
        <div class="masonry masonry-sm lg:masonry-md">
          <div v-for="file in files" class="py-3 break-inside">
            <lazy-image
              :src="`/public/moments/.thumb/${file}`"
              class="rounded-lg"
            />
          </div>
        </div>
        <matrimony-button
          class="sticky top-44 lg:top-[5.5rem] -order-1 self-start"
          to="/"
          @click.prevent="$router.back()"
        >
          <template #prefix-icon>
            <font-awesome-icon :icon="['far', 'arrow-left']" class="mr-2" />
          </template>
          Back
        </matrimony-button>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from "vue";
import LazyImage from "../components/LazyImage.vue";
import MatrimonyButton from "../components/MatrimonyButton.vue";

const files = ref([]);
fetch("/api/moments").then(async (r) => {
  files.value = await r.json();
});
</script>

<script>
export default {
  name: "Moments",
};
</script>
