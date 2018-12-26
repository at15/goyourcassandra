<template>
    <div class="login-form">
        <!--http://element-cn.eleme.io/#/zh-CN/component/form-->
        <el-form ref="form" :model="form" label-width="80px">
            <el-form-item label="Bookmark">
                <el-select v-model="form.bookmark" v-on:change="selectBookmark" placeholder="Select Bookmark">
                    <el-option label="no" value="no"></el-option>
                    <el-option v-for="(bk, name) in bookmarks" :key="name" :label="name" :value="name"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="Host">
                <el-input v-model="form.host"></el-input>
            </el-form-item>
            <el-form-item label="Username">
                <el-input v-model="form.username"></el-input>
            </el-form-item>
            <!-- TODO: allow show password -->
            <!-- https://simedia.tech/blog/show-hide-password-input-values-with-vue-js/ -->
            <el-form-item label="Password">
                <el-input type="password" v-model="form.password"></el-input>
            </el-form-item>
            <el-form-item label="Keyspace">
                <el-select v-model="form.keyspace" placeholder="Select Keyspace">
                    <el-option v-for="(ks, index) in form.keyspaces" :key="index" :label="ks" :value="ks"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="connectDB">Connect</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Login",
  data() {
    return {
      form: {
        bookmark: "",
        host: "127.0.0.1",
        username: "cassandra",
        password: "cassandra",
        keyspace: "system",
        keyspaces: ["system"],
      },
      bookmarks: {}
    };
  },
  mounted() {
    console.log("mounted!");
    axios
      .get("/api/bookmark")
      .then(
        res => {
          console.log(res.data);
          this.bookmarks = res.data;
        },
        err => {
          console.warn(err);
        }
      )
      .catch(e => {
        console.warn(e);
      });
  },
  methods: {
    connectDB() {
      console.log("connect to db");
    },
    selectBookmark() {
      console.log("book mark selected", this.form.bookmark);
      if (this.form.bookmark === "" || this.form.bookmark === "no") {
        console.log("reset fields");
        this.form.host = "";
        this.form.username = "";
        this.form.keyspace = "";
        this.form.keyspaces = ["system"];
        return
      }
      let bm = this.bookmarks[this.form.bookmark];
      this.form.host = bm.host;
      this.form.username = bm.username;
      this.form.password = bm.password;
      this.form.keyspace = bm.keyspace;
      this.form.keyspaces = bm.keyspaces;
    }
  }
};
</script>

<style scoped>
.login-form {
  padding: 50px 50px 20px 30px;

  border-radius: 20px;

  background-color: rgba(255, 255, 255, 0.8);
}
</style>
