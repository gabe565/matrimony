import { createApp } from "vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import router from "./plugins/router";
import "./index.scss";
import App from "./App.vue";

import "./plugins/fontawesome";

createApp(App)
  .component("font-awesome-icon", FontAwesomeIcon)
  .use(router)
  .mount("#app");
