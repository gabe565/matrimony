<template>
  <div class="flex flex-col">
    <slot />
  </div>
</template>

<script>
export default {
  name: "RadioButtonGroup",
};
</script>

<style scoped lang="scss">
:slotted(label) {
  &:first-child {
    @apply rounded-b-none;
  }

  @apply mb-0 border-b-0;

  &:last-child {
    @apply rounded-t-none border-b;
  }

  &:only-child {
    @apply rounded-lg;
  }
}
</style>
