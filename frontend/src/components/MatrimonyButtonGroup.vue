<template>
  <div class="root">
    <slot />
  </div>
</template>

<script>
export default {
  name: "MatrimonyButtonGroup",
};
</script>

<style scoped lang="scss">
.root > {
  :slotted(*) {
    &:first-child {
      @apply rounded-r-none border-l-0;
    }

    @apply border-l-0 border-r mr-0;

    &:last-child {
      @apply rounded-l-none border-r-0;
    }

    &:only-child {
      @apply rounded-full;
    }
  }
}
</style>
