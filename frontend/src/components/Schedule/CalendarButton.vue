<template>
  <matrimony-button
    :href="`/api/ical/${sectionKey}/${eventKey}`"
    target="_blank"
    color="purple"
  >
    Add to Calendar
    <template #icon>
      <calendar-add-icon class="ml-2" />
    </template>
  </matrimony-button>
</template>

<script setup>
import CalendarAddIcon from "~icons/solar/calendar-add-bold";
import MatrimonyButton from "../MatrimonyButton.vue";

defineProps({
  sectionKey: { type: Number, default: 0 },
  eventKey: { type: Number, default: 0 },
});
</script>

<script>
export default {
  name: "CalendarButton",
};
</script>
