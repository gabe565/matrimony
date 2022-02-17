import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import NotFound from "../views/NotFound.vue";
import Moments from "../views/Moments.vue";
import RSVP from "../views/RSVP.vue";
import { waitForElement } from "../util/waitForElement";
import LookupPage from "../components/RSVP/LookupPage.vue";
import QuestionsPage from "../components/RSVP/QuestionsPage.vue";

const router = [
  {
    path: "/",
    component: Home,
    meta: {
      headerCollapseY: () => window.innerHeight,
    },
  },
  {
    path: "/moments",
    component: Moments,
    meta: {
      headerCollapseY: 100,
    },
  },
  {
    path: "/rsvp",
    redirect: "/rsvp/begin",
    component: RSVP,
    children: [
      {
        path: "begin",
        component: LookupPage,
      },
      {
        path: "questions",
        component: QuestionsPage,
      },
    ],
    meta: {
      headerCollapseY: 100,
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
