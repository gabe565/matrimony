<template>
  <div class="container mx-auto lg:max-w-[1436px] flex flex-col pt-52 lg:pt-0">
    <section
      class="lg:mt-40 text-slate-600 dark:text-slate-400 bg-white dark:bg-slate-900 body-font rounded-t-2xl"
    >
      <div class="container px-5 py-24 mx-auto">
        <div class="text-center mb-4">
          <h1
            class="text-4xl font-script title-font mb-2 text-slate-900 dark:text-white"
          >
            404 Not Found
          </h1>
        </div>
        <span>
          I hate to be the bearer of bad news, but the requested file could not
          be located.
          <router-link to="/" class="underline"
            >Click here to go home.</router-link
          >
        </span>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "NotFound",
};
</script>
