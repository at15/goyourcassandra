import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home.vue";
import About from "./views/About.vue";
import Login from "./views/Login.vue";

Vue.use(Router);

// TODO: redirect when not login http://jasonwatmore.com/post/2018/07/14/vue-vuex-user-registration-and-login-tutorial-example
export default new Router({
  routes: [
    {
      path: "/",
      name: "home",
      component: Home
    },
    {
      path: "/about",
      name: "about",
      component: About
    },
    {
      path: "/login",
      name: "login",
      component: Login
    }
  ]
});
