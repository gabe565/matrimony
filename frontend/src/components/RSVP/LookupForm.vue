<template>
  <div class="flex flex-wrap lg:w-1/2 mx-auto">
    <text-field
      v-model="lookup.first"
      label="First Name"
      class="w-full sm:w-1/2 sm:pr-8 pb-3"
      stacked
    />
    <text-field
      v-model="lookup.last"
      label="Last Name"
      class="w-full sm:w-1/2 pb-3"
      stacked
    />
    <text-field
      v-model="lookup.email"
      label="Email Address"
      class="w-full pb-3"
      stacked
    />
    <matrimony-button class="ml-auto mr-30" @click.prevent="checkUser">
      Next
      <template #icon>
        <font-awesome-icon :icon="['far', 'arrow-right']" class="ml-2" />
      </template>
    </matrimony-button>
  </div>
</template>

<script setup>
import { reactive, ref } from "vue";
import TextField from "../Forms/TextField.vue";
import MatrimonyButton from "../MatrimonyButton.vue";

const lookup = reactive({
  first: "",
  last: "",
  email: "",
});

const user = ref({});
const checkUser = async () => {
  const params = new URLSearchParams(Object.entries(lookup));
  const r = await fetch(`/api/rsvp/checkUser?${params}`);
  const j = await r.json();
  console.log(j);
};
</script>
