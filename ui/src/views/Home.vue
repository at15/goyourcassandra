<template>
  <div class="home">
    <!--<img src="../assets/logo.png">-->
    <!--<HelloWorld msg="Welcome to Your Vue.js App"/>-->
    <DB name="DB1"/>
  </div>
</template>

<script>
// @ is an alias to /src
import DB from "@/components/DB.vue";

export default {
  name: "home",
  components: {
    DB
  }
};
</script>
