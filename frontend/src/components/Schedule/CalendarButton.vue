<template>
  <matrimony-button
    :href="`/api/ical/${sectionKey}/${eventKey}`"
    target="_blank"
  >
    Add to Calendar
    <template #icon>
      <font-awesome-icon :icon="['fas', 'calendar-plus']" class="ml-2" />
    </template>
  </matrimony-button>
</template>

<script setup>
import MatrimonyButton from "../MatrimonyButtonLink.vue";

defineProps({
  sectionKey: { type: Number, default: 0 },
  eventKey: { type: Number, default: 0 },
});
</script>
