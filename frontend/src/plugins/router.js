import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import NotFound from "../views/NotFound.vue";
import { waitForElement } from "../util/waitForElement";

const router = [
  {
    path: "/",
    component: Home,
    meta: {
      headerCollapseY: () => window.innerHeight,
    },
  },
  {
    path: "/:pathMatch(.*)*",
    component: NotFound,
  },
];

export default createRouter({
  history: createWebHistory(),
  routes: router,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }
    if (to.hash) {
      /* eslint-disable-next-line no-async-promise-executor */
      return new Promise(async (resolve) => {
        await waitForElement(to.hash);
        resolve({ el: to.hash });
      });
    }
    return { top: 0 };
  },
});
