<template>
  <div
    class="sticky top-0 -z-10 container mx-auto pb-2 lg:pb-6 h-screen flex flex-col justify-center text-center pt-20 lg:pt-0"
  >
    <div
      class="h-[40rem] lg:h-[48rem] w-full pb-8 relative -mb-32 md:-mb-52 -z-10"
    >
      <img
        v-if="image.src"
        class="absolute top-0 left-0 h-full w-full bg-cover object-cover lg:rounded-[1rem]"
        :src="image.src"
        :alt="title"
      />
      <div
        class="absolute bottom-0 left-0 h-1/2 w-full bg-gradient-to-b from-transparent to-black lg:rounded-[1rem]"
      />
    </div>
    <h1
      class="pb-12 md:pb-16 font-script text-accent-50 text-4xl md:text-7xl whitespace-nowrap"
    >
      We're getting married!
    </h1>
    <font-awesome-icon
      :icon="['fas', 'chevron-down']"
      class="motion-safe:animate-bounce text-accent-50 text-xl pb-4"
    />
  </div>
</template>

<script setup>
defineProps({
  title: { type: String, default: "" },
  image: { type: Object, default: () => ({}) },
});
</script>
