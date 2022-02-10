import { createApp } from "vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import "./index.scss";
import App from "./App.vue";

import "./plugins/fontawesome";

createApp(App).component("font-awesome-icon", FontAwesomeIcon).mount("#app");
