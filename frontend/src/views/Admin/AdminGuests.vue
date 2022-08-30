<template>
  <q-page padding class="q-px-xl">
    <div class="row py-5">
      <div class="col">
        <h1 class="text-h4">Guests</h1>
      </div>
    </div>
    <div class="row py-5 w-full">
      <q-table
        v-model:filter="filters"
        :filter-method="applyFilter"
        :loading="loading"
        :rows="guests"
        :columns="columns"
        :visible-columns="visibleColumns"
        row-key="id"
        :pagination="{ rowsPerPage: 0 }"
        virtual-scroll
        :rows-per-page-options="[0]"
        class="w-full"
        :class="{ 'h-[700px]': !fullscreen }"
        table-header-class="sticky top-0 bg-white z-10"
        table-class="z-0"
        :fullscreen="fullscreen"
      >
        <template #top-right>
          <q-select
            v-model="visibleColumns"
            multiple
            outlined
            dense
            options-dense
            :display-value="$q.lang.table.columns"
            emit-value
            map-options
            :options="columns"
            option-value="name"
            options-cover
            style="min-width: 150px"
          />

          <q-btn flat round dense :icon="fasFilter" class="q-ml-md">
            <q-menu v-model="showFilters" class="q-pa-md">
              <div class="row no-wrap items-center">
                <div class="col">
                  <div class="text-h6">Filters</div>
                </div>
              </div>

              <div class="row">
                <div class="col">
                  <q-input
                    v-model="filters.name"
                    dense
                    debounce="300"
                    label="Name"
                    clearable
                  />

                  <q-input
                    v-model="filters.party"
                    dense
                    debounce="300"
                    label="Party"
                    clearable
                  />

                  <q-toggle
                    v-model="filters.responded"
                    label="Responded"
                    toggle-indeterminate
                  />

                  <q-toggle
                    v-model="filters.attending"
                    label="Attending"
                    toggle-indeterminate
                  />

                  <q-toggle
                    v-model="filters.mail_invite"
                    label="Mail Invite"
                    toggle-indeterminate
                  />
                </div>
              </div>
              <div class="row">
                <div class="col text-right">
                  <q-btn color="primary" @click="showFilters = !showFilters"
                    >Close</q-btn
                  >
                </div>
              </div>
            </q-menu>
          </q-btn>

          <q-btn
            flat
            round
            dense
            :icon="fullscreen ? fasCompress : fasExpand"
            class="q-ml-md"
            @click="fullscreen = !fullscreen"
          />
        </template>
      </q-table>
    </div>
  </q-page>
</template>

<script setup>
import { ref } from "vue";
import {
  fasExpand,
  fasCompress,
  fasFilter,
} from "@quasar/extras/fontawesome-v6";

const loading = ref(true);
const guests = ref([]);
fetch("/api/guest").then(async (r) => {
  loading.value = false;
  guests.value = await r.json();
});

const columns = [
  {
    name: "party",
    label: "Party",
    field: (val) => val.party.id,
  },
  {
    name: "first",
    required: true,
    label: "First Name",
    align: "left",
    field: "first",
    sortable: true,
    headerStyle: "width: 0.1%; white-space: nowrap",
  },
  {
    name: "last",
    label: "Last Name",
    align: "left",
    field: "last",
    sortable: true,
    headerStyle: "width: 0.1%; white-space: nowrap",
  },
  {
    name: "email",
    label: "Email",
    align: "left",
    field: "email",
    format: (val) => val || "",
  },
  {
    name: "rsvp",
    label: "RSVP",
    align: "left",
    field: (val) => val.rsvp?.RSVP || "",
    sortable: true,
  },
  {
    name: "number_attending",
    label: "# Attending",
    field: (val) => val.rsvp?.["Number Attending"] || "",
    sortable: true,
  },
  {
    name: "honeymoon",
    label: "Honeymoon",
    align: "left",
    field: (val) => val.rsvp?.Honeymoon || "",
    sortable: true,
  },
  {
    name: "how_do_you_know_the_couple",
    label: "How do you know the couple?",
    align: "left",
    field: (val) => val.rsvp?.["How do you know the couple"] || "",
    sortable: true,
    hidden: true,
  },
  {
    name: "mail_invite",
    label: "Mail Invite",
    align: "left",
    field: (val) => val.rsvp?.["Mail Invitation"] || "",
    sortable: true,
    hidden: true,
  },
  {
    name: "address",
    label: "Mailing Address",
    align: "left",
    field: (val) => val.rsvp?.["Mailing Address"] || "",
    hidden: true,
  },
  {
    name: "updatedAt",
    label: "Updated",
    align: "left",
    field: "updatedAt",
    format: (val) => new Date(val).toLocaleString(),
    sort: (a, b) => new Date(a) - new Date(b),
    sortable: true,
    sortOrder: "da",
  },
];
const visibleColumns = ref(
  columns.filter((col) => !col.hidden).map((col) => col.name)
);

const fullscreen = ref(false);
const filters = ref({});
const showFilters = ref(false);

const filterName = (filter, row) => {
  if (filter) {
    return `${row.first} ${row.last}`
      .toLowerCase()
      .includes(filter.toLowerCase());
  }
  return true;
};

const filterParty = (filter, row) => {
  if (filter) {
    return row.party.id === Number(filter);
  }
  return true;
};

const filterResponded = (filter, row) => {
  if (filter === true || filter === false) {
    return Boolean(row.rsvp) === filter;
  }
  return true;
};

const filterRsvpBool = (filter, val, trueVal = true, falseVal = false) => {
  switch (filter) {
    case true:
      return val === trueVal;
    case false:
      return val === falseVal;
    default:
      return true;
  }
};

const applyFilter = (rows, terms) => {
  return rows.filter(
    (row) =>
      filterName(terms.name, row) &&
      filterParty(terms.party, row) &&
      filterResponded(terms.responded, row) &&
      filterRsvpBool(terms.attending, row.rsvp?.RSVP, "yes", "no") &&
      filterRsvpBool(
        terms.mail_invite,
        row.rsvp?.["Mail Invitation"],
        "Yes Please!",
        "No Thank You!"
      )
  );
};
</script>

<script>
export default {
  name: "AdminParties",
};
</script>

<style lang="scss">
.row :deep(table) {
  position: sticky;
}
</style>
