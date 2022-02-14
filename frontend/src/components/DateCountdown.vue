<template>
  <div>
    <template v-for="e in countdown">
      <span class="pr-1 text-2xl font-medium">{{ e[1] }}</span>
      <span class="pr-2">{{ e[0] }}</span>
    </template>
  </div>
</template>

<script setup>
import { computed, onUnmounted, ref } from "vue";
import { DateTime } from "luxon";

const props = defineProps({
  date: { type: String, default: "" },
  seconds: { type: Boolean, default: false },
});

const now = ref(DateTime.now());
const date = computed(() => DateTime.fromISO(props.date));

const units = computed(() => {
  const v = ["days", "hours", "minutes"];
  if (props.seconds) {
    v.push("seconds");
  }
  return v;
});
const countdown = computed(() =>
  Object.entries(date.value.diff(now.value, units.value).toObject()).map(
    (e) => {
      e[1] = Math.floor(e[1]);
      return e;
    }
  )
);

const interval = setInterval(() => {
  now.value = DateTime.now();
}, 1000);

onUnmounted(() => {
  clearInterval(interval);
});
</script>
