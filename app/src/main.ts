import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";
import ToastPlugin from "vue-toast-notification";
import "vue-toast-notification/dist/theme-bootstrap.css";

createApp(App)
  .use(router)
  .use(ToastPlugin, { position: "bottom" })
  .mount("#app");
