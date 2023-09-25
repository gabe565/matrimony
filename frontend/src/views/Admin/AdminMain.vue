<template>
  <q-layout view="hHh LpR fFf">
    <q-header elevated class="bg-primary text-white">
      <q-toolbar>
        <q-btn
          dense
          flat
          round
          :icon="fasBars"
          @click="leftDrawerOpen = !leftDrawerOpen"
        />

        <q-toolbar-title>
          <template v-if="config.partners">
            {{ config.partners.map((e) => e.first).join(" & ") }} -
          </template>
          Matrimony Admin
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above side="left" :width="200">
      <template v-for="page in pages" :key="page.value">
        <q-item v-ripple clickable :to="page.to" class="rounded-r-full">
          <q-item-section avatar>
            <q-icon :name="page.icon" />
          </q-item-section>

          <q-item-section>{{ page.title }}</q-item-section>
        </q-item>
      </template>
    </q-drawer>

    <q-page-container>
      <router-view class="bg-white" />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref } from "vue";
import { fasBars, fasUser } from "@quasar/extras/fontawesome-v6";
import { API_URL } from "../../config/api";

// Import Quasar css
import("quasar/src/css/index.sass");

const leftDrawerOpen = ref(false);

const loading = ref(true);
const config = ref({});
fetch(`${API_URL}/api/config`).then(async (r) => {
  config.value = await r.json();
  loading.value = false;
});

const pages = [{ title: "Guests", to: "/admin/guests", icon: fasUser }];
</script>

<script>
export default {
  name: "AdminMain",
};
</script>
