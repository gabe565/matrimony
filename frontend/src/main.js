import { createApp } from "vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { Quasar } from "quasar";
import quasarIconSet from "quasar/icon-set/svg-fontawesome-v5";
import VueLazyLoad from "vue3-lazyload";
import router from "./plugins/router";
import store from "./plugins/store";
import "./index.scss";
import App from "./App.vue";

import "./plugins/fontawesome";

createApp(App)
  .component("font-awesome-icon", FontAwesomeIcon)
  .use(router)
  .use(VueLazyLoad)
  .use(store)
  .use(Quasar, {
    plugins: {},
    iconSet: quasarIconSet,
  })
  .mount("#app");
