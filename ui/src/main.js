import Vue from "vue";
import ElementUI from "element-ui";
import locale from "element-ui/lib/locale/lang/en";
import "element-ui/lib/theme-chalk/index.css";
import App from "./App.vue";
import router from "./router";
import Axios from 'axios'

Vue.prototype.$http = Axios; // https://scotch.io/tutorials/vue-authentication-and-route-handling-using-vue-router
Vue.config.productionTip = false;
Vue.use(ElementUI, { locale });

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
