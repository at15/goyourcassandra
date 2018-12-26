<template>
    <div id="app">
        <div id="nav">
            <router-link to="/">Home</router-link>
            |
            <router-link to="/about">About</router-link>
            |
            <router-link to="/login">Login</router-link>
        </div>
        <router-view/>
        <div id="footer">
            <a href="https://github.com/at15/goyourcassandra" target="_blank">GitHub</a>
        </div>
    </div>
</template>

<script>
export default {
  methods: {
    // NOTE: based on https://github.com/ElementUI/element-starter/blob/master/src/App.vue
    startHacking() {
      this.$notify({
        title: "It works!",
        type: "success",
        message:
          "We've laid the ground work for you. It's time for you to build something epic!",
        duration: 5000
      });
    }
  }
};
</script>

<style>
/* https://css-tricks.com/perfect-full-page-background-image/ */
html {
  background: url(https://source.unsplash.com/random) no-repeat center center
    fixed;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}
#app {
  margin-top: 20px;

  /*https://css-tricks.com/snippets/css/a-guide-to-flexbox/*/
  /* app is the container */
  display: flex;

  /* flex-direction is the main axis, column means align vertically */
  flex-direction: column;
  /* justify-content is along the main axis */
  /* space-between: items are evenly distributed in the line; first item is on the start line, last item on the end line */
  justify-content: space-between;
  /* default alignment of items Think of it as the justify-content version for the cross-axis (perpendicular to the main-axis)*/
  align-items: center;
  align-content: center;

  min-height: 600px;
}

#nav {
}

#footer {
}
</style>
