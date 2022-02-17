<template>
  <matrimony-button
    :href="`https://maps.google.com/?daddr=${location}`"
    target="_blank"
    color="purple"
  >
    Navigate
    <template #icon>
      <font-awesome-icon :icon="['fas', 'diamond-turn-right']" class="ml-2" />
    </template>
  </matrimony-button>
</template>

<script setup>
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
