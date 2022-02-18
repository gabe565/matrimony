<template>
  <matrimony-modal v-model="value" center disable-scroll>
    <lazy-image
      :src="src"
      class="h-screen"
      content-fit="object-scale-down"
      bg-color="bg-transparent"
    />
  </matrimony-modal>
</template>

<script setup>
import { computed } from "vue";
import MatrimonyModal from "../MatrimonyModal.vue";
import LazyImage from "../LazyImage.vue";

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  src: { type: String, default: "" },
});
const emit = defineEmits(["update:modelValue"]);

const value = computed({
  get: () => props.modelValue,
  set: (e) => emit("update:modelValue", e),
});
</script>

<script>
export default {
  name: "PhotoModal",
};
</script>
