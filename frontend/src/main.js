import { createApp } from "vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
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
  .mount("#app");
