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
:slotted(a, input[type="button"]) {
  @apply first:rounded-r-none border-l-0 border-r mr-0 last:rounded-l-none last:border-r-0 only:rounded-full;
}
</style>
