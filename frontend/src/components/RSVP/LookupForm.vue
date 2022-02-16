<template>
  <form class="flex flex-wrap lg:w-1/2 mx-auto" @submit.prevent="checkUser">
    <text-field
      v-model="lookup.first"
      label="First Name"
      class="w-full sm:w-1/2 sm:pr-8"
      required
    />
    <text-field
      v-model="lookup.last"
      label="Last Name"
      class="w-full sm:w-1/2"
      required
    />
    <text-field v-model="lookup.email" label="Email Address" class="w-full" />
    <matrimony-form-button class="ml-auto mr-30">
      Next
      <template #icon>
        <font-awesome-icon
          :icon="loading ? ['fad', 'spinner-third'] : ['far', 'arrow-right']"
          :class="{ 'animate-spin': loading }"
          class="ml-2 fa-fw"
        />
      </template>
    </matrimony-form-button>
  </form>
</template>

<script setup>
import { reactive, ref } from "vue";
import TextField from "../Forms/TextField.vue";
import MatrimonyFormButton from "../Forms/MatrimonyFormButton.vue";

const lookup = reactive({
  first: "",
  last: "",
  email: "",
});

const loading = ref(false);
const foundUser = ref(null);
const checkUser = async () => {
  loading.value = true;
  foundUser.value = null;
  const params = new URLSearchParams(Object.entries(lookup));
  try {
    const r = await fetch(`/api/rsvp/checkUser?${params}`);
    foundUser.value = r.ok;
  } catch (error) {
    console.error(error);
    foundUser.value = false;
  } finally {
    loading.value = false;
  }
};
</script>
