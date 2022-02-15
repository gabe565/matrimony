<template>
  <div class="root">
    <slot />
  </div>
</template>

<style scoped lang="scss">
.root > {
  :slotted(*) {
    &:first-child {
      @apply rounded-r-none border-l;
    }

    @apply border-l-0;

    &:last-child {
      @apply rounded-l-none;
    }

    &:only-child {
      @apply rounded-lg;
    }
  }
}
</style>
