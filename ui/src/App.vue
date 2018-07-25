<template>
  <div id="app">
   <h1>GoYourCassandra</h1>
    <!--<HelloWorld msg="Welcome to Your Vue.js App"/>-->
    <DB name="DB1"/>
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'

import DB from './components/DB'

export default {
  name: 'app',
  components: {
    // HelloWorld
      DB
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
