<template>
  <matrimony-button href="#" target="_blank">
    Add to Calendar
    <template #icon>
      <font-awesome-icon :icon="['fas', 'calendar-plus']" class="ml-2" />
    </template>
  </matrimony-button>
</template>

<script setup>
import MatrimonyButton from "../MatrimonyButton.vue";
</script>
