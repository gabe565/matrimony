<template>
  <div class="relative flex flex-col lg:w-1/2 mx-auto">
    <h4
      class="text-2xl font-bold leading-none text-gray-900 dark:text-white mx-auto py-6"
    >
      Thank you! Your RSVP has been sent.
    </h4>
    <matrimony-button-group class="mx-auto">
      <matrimony-button size="xl" @click.prevent="router.back()">
        <template #prefix-icon>
          <edit-icon class="mr-1" aria-hidden="true" />
        </template>
        Edit RSVP
      </matrimony-button>
      <matrimony-button to="/" size="xl">
        <template #prefix-icon>
          <home-icon class="mr-1" aria-hidden="true" />
        </template>
        Back to site
      </matrimony-button>
    </matrimony-button-group>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import EditIcon from "~icons/mdi/pencil";
import HomeIcon from "~icons/mdi/home";
import MatrimonyButton from "../MatrimonyButton.vue";
import MatrimonyButtonGroup from "../MatrimonyButtonGroup.vue";

const router = useRouter();
</script>

<script>
export default {
  name: "FinishRSVP",
};
</script>
