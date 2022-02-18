<template>
  <teleport to="#modals">
    <div v-if="show" class="fixed top-0 left-0 w-full h-full bg-black/50 z-40">
      <div
        aria-hidden="true"
        class="overflow-y-auto overflow-x-hidden fixed top-1/2 left-1/2 z-50 justify-center items-center h-modal -translate-x-1/2 -translate-y-1/2"
      >
        <div class="relative px-4 w-full max-w-2xl h-full md:h-auto">
          <!-- Modal content -->
          <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
            <!-- Modal header -->
            <div
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
              class="flex items-center p-6 space-x-2 rounded-b border-t border-gray-200 dark:border-gray-600"
            >
              <slot name="footer" :toggle="toggle" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </teleport>
  <slot name="activator" :on="on" :attrs="attrs" />
</template>

<script setup>
import { ref } from "vue";
import MatrimonyButton from "./MatrimonyButton.vue";

const attrs = {};

const show = ref(false);
const toggle = () => {
  show.value = !show.value;
};
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
