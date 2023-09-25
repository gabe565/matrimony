<template>
  <matrimony-button
    :href="`https://maps.google.com/?daddr=${location}`"
    target="_blank"
    color="purple"
  >
    Navigate
    <template #icon>
      <navigate-icon class="ml-2" aria-hidden="true" />
    </template>
  </matrimony-button>
</template>

<script setup>
import NavigateIcon from "~icons/solar/map-arrow-up-bold";
import MatrimonyButton from "../MatrimonyButton.vue";

defineProps({
  location: { type: String, default: "" },
});
</script>

<script>
export default {
  name: "NavigateButton",
};
</script>
