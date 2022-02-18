<template>
  <teleport v-if="value" to="#modals">
    <!-- Backdrop -->
    <div
      class="fixed inset-0 z-40 bg-gray-900 bg-opacity-50 dark:bg-opacity-80"
      @click.stop="toggle"
    />

    <!-- Modal -->
    <div
      :aria-hidden="!value"
      aria-modal="true"
      role="dialog"
      class="flex overflow-x-hidden fixed top-4 right-0 left-0 z-50 justify-center items-center h-modal md:h-full md:inset-0"
      :class="{ 'h-full': center, 'overflow-y-auto': !disableScroll }"
      @click.stop="toggle"
    >
      <slot>
        <div
          class="relative px-4 w-full max-w-2xl h-full md:h-auto"
          :class="{ 'h-auto': center }"
          @click.stop="toggle"
        >
          <!-- Modal content -->
          <div
            class="relative bg-white rounded-lg shadow dark:bg-gray-700"
            @click.stop
          >
            <!-- Modal header -->
            <div
              v-if="$slots.title"
              class="flex justify-between items-start p-5 rounded-t border-b dark:border-gray-600"
            >
              <h3
                class="text-xl font-semibold text-gray-900 lg:text-2xl dark:text-white"
              >
                <slot name="title" />
              </h3>
              <matrimony-button
                color="transparent"
                size="icon"
                @click.prevent="toggle"
              >
                <font-awesome-icon :icon="['fas', 'xmark']" class="text-lg" />
              </matrimony-button>
            </div>
            <!-- Modal body -->
            <div class="p-6 space-y-6">
              <slot name="body" :toggle="toggle" />
            </div>
            <!-- Modal footer -->
            <div
              v-if="$slots.footer"
              class="flex items-center p-6 space-x-2 rounded-b border-t border-gray-200 dark:border-gray-600"
            >
              <slot name="footer" :toggle="toggle" />
            </div>
          </div>
        </div>
      </slot>
    </div>
  </teleport>
  <slot name="activator" :on="on" />
</template>

<script setup>
import { onUnmounted, ref, watch, watchEffect } from "vue";
import { onBeforeRouteUpdate, useRoute, useRouter } from "vue-router";
import MatrimonyButton from "./MatrimonyButton.vue";

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  center: { type: Boolean, default: false },
  disableScroll: { type: Boolean, default: false },
});
const emit = defineEmits(["update:modelValue"]);
const route = useRoute();
const router = useRouter();

const value = ref(props.modelValue);
watchEffect(() => {
  value.value = props.modelValue;
});

const toggle = () => {
  value.value = !value.value;
  emit("update:modelValue", value.value);
};

const hash = "#modal";
const goBack = ref(false);
onBeforeRouteUpdate((to, from) => {
  if (goBack.value && from.hash === hash) {
    goBack.value = false;
    toggle();
  }
});

const bodyClasses = [
  "fixed",
  "left-0",
  "right-0",
  "overscroll-contain",
  "md:overscroll-auto",
];
const preventScroll = async (v, enableScroll) => {
  if (v) {
    await router.push(`${route.path}${hash}`);
    goBack.value = true;
    document.body.style.top = `-${window.scrollY}px`;
    document.body.classList.add(...bodyClasses);
  } else {
    if (goBack.value) {
      goBack.value = false;
      await router.back();
    }
    const scrollY = document.body.style.top;
    document.body.classList.remove(...bodyClasses);
    document.body.style.top = "";
    if (enableScroll) {
      window.scrollTo({
        top: parseInt(scrollY || "0", 10) * -1,
        behavior: "instant",
      });
    }
  }
};
watch(value, preventScroll);
onUnmounted(() => preventScroll(false, false));

const on = {
  click: toggle,
};
</script>

<script>
export default {
  name: "MatrimonyModal",
};
</script>

<style scoped></style>
