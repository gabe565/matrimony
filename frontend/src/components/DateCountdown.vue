<template>
  <div>
    <template v-for="e in countdown">
      <span class="pr-1 text-2xl font-medium">{{ e.value }}</span>
      <span class="pr-2">{{ e.unit }}</span>
    </template>
    <template v-if="inPast">ago</template>
  </div>
</template>

<script setup>
import { computed, onUnmounted, ref } from "vue";
import dayjs from "dayjs";
import duration from "dayjs/plugin/duration";

const props = defineProps({
  date: { type: String, default: "" },
  seconds: { type: Boolean, default: false },
});

const now = ref(dayjs());
const date = computed(() => dayjs(props.date));

dayjs.extend(duration);

const inPast = computed(() => date.value.isBefore(now.value));

const countdown = computed(() => {
  const d = dayjs.duration(date.value.diff(now.value));
  const v = [
    { unit: "years", value: Math.abs(d.years()) },
    { unit: "months", value: Math.abs(d.months()) },
    { unit: "days", value: Math.abs(d.days()) },
    { unit: "hours", value: Math.abs(d.hours()) },
    { unit: "minutes", value: Math.abs(d.minutes()) },
  ];
  if (props.seconds) {
    v.push({ unit: "seconds", value: Math.abs(d.seconds()) });
  }
  // Handle plurals
  for (const e of v) {
    if (e.value === 1) {
      e.unit = e.unit.slice(0, -1);
    }
  }
  // Remove beginning non-zero values
  const firstNonZero = v.findIndex((e) => e.value !== 0);
  return v.slice(firstNonZero);
});

const interval = setInterval(() => {
  now.value = dayjs();
}, 1000);

onUnmounted(() => {
  clearInterval(interval);
});
</script>

<script>
export default {
  name: "DateCountdown",
};
</script>
