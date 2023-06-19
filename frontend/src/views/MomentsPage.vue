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
              :src="file.thumb"
              class="rounded-lg cursor-pointer"
              @click="showModal(file.src)"
            />
          </div>
        </div>
        <matrimony-button
          class="sticky top-44 lg:top-[5.5rem] -order-1 self-start"
          to="/"
          @click.prevent="$router.back()"
        >
          <template #prefix-icon>
            <arrow-left-icon class="mr-1" />
          </template>
          Back
        </matrimony-button>

        <photo-modal v-model="modalData.show" :src="modalData.src" />
      </div>
    </section>
  </div>
</template>

<script setup>
import { reactive, ref } from "vue";
import ArrowLeftIcon from "~icons/solar/alt-arrow-left-bold";
import LazyImage from "../components/LazyImage.vue";
import MatrimonyButton from "../components/MatrimonyButton.vue";
import PhotoModal from "../components/Moments/PhotoModal.vue";

const files = ref([]);
fetch("/api/moments").then(async (r) => {
  files.value = await r.json();
});

const modalData = reactive({
  show: false,
  src: "",
});
const showModal = (src) => {
  modalData.src = src;
  modalData.show = true;
};
</script>

<script>
export default {
  name: "MomentsPage",
};
</script>
