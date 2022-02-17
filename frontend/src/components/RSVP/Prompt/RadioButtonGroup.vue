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
  @apply first:rounded-b-none mb-0 border-b-0 last:rounded-t-none last:border-b only:rounded-lg;
}
</style>
