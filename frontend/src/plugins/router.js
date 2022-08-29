import { createRouter, createWebHistory } from "vue-router";
import PublicPage from "../views/PublicPage.vue";
import HomePage from "../views/HomePage.vue";
import NotFound from "../views/NotFound.vue";
import MomentsPage from "../views/MomentsPage.vue";
import RSVP from "../views/RSVP.vue";
import { waitForElement } from "../util/waitForElement";
import LookupPage from "../components/RSVP/LookupPage.vue";
import QuestionsPage from "../components/RSVP/QuestionsPage.vue";
import FinishRSVP from "../components/RSVP/FinishRSVP.vue";
import AdminMain from "../views/Admin/AdminMain.vue";
import AdminGuests from "../views/Admin/AdminGuests.vue";

const router = [
  {
    path: "/admin",
    component: AdminMain,
    redirect: "/admin/guests",
    children: [
      {
        path: "guests",
        component: AdminGuests,
      },
    ],
  },
  {
    path: "/",
    component: PublicPage,
    children: [
      {
        path: "/",
        component: HomePage,
        meta: {
          headerCollapseY: () => window.innerHeight,
        },
      },
      {
        path: "/moments",
        component: MomentsPage,
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
          {
            path: "finish",
            component: FinishRSVP,
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
    ],
  },
];

export default createRouter({
  history: createWebHistory(),
  routes: router,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return { ...savedPosition, behavior: "instant" };
    }
    if (to.hash) {
      /* eslint-disable-next-line no-async-promise-executor */
      return new Promise(async (resolve) => {
        await waitForElement(to.hash);
        resolve({ el: to.hash });
      });
    }
    return false;
  },
});
