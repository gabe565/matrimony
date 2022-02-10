<template>
  <div
    class="fixed top-0 left-0 w-full h-full -z-20 bg-gradient-to-b from-primary-400 to-primary-200"
  >
    <div class="fill-accent-500 opacity-30">
      <!-- eslint-disable vue/no-v-html -->
      <div
        class="absolute bottom-0 left-0 -mb-10 md:-mb-44 -ml-1 w-1/3 md:w-[520px] fill-accent-900"
        v-html="blueContourBottomLeft"
      />
      <div
        class="absolute top-0 right-0 hidden lg:block -mt-20 -mr-16 w-[520px]"
        v-html="blueContourTopRight"
      />

      <div class="absolute top-0 w-full lg:hidden" v-html="blueContourTablet" />
    </div>
  </div>
</template>

<script setup>
/* eslint-disable import/no-unresolved */
import blueContourTablet from "../assets/blue_contour/tablet.svg?raw";
import blueContourTopRight from "../assets/blue_contour/top_right.svg?raw";
import blueContourBottomLeft from "../assets/blue_contour/bottom_left.svg?raw";
/* eslint-enable import/no-unresolved */
</script>
