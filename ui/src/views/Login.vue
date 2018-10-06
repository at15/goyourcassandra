<template>
    <div class="login-form">
        <!--http://element-cn.eleme.io/#/zh-CN/component/form-->
        <el-form ref="form" :model="form" label-width="80px">
            <el-form-item label="Host">
                <el-input v-model="form.host"></el-input>
            </el-form-item>
            <el-form-item label="Username">
                <el-input v-model="form.username"></el-input>
            </el-form-item>
            <!-- TODO: allow show password -->
            <el-form-item label="Password">
                <el-input type="password" v-model="form.password"></el-input>
            </el-form-item>
            <el-form-item label="Keyspace">
                <el-select v-model="form.keyspace" placeholder="Select Keyspace">
                    <el-option label="system" value="system"></el-option>
                    <el-option label="app" value="app"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit">立即创建</el-button>
                <el-button>取消</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      form: {
        host: "127.0.0.1",
        username: "cassandra",
        password: "cassandra",
        date1: "",
        date2: "",
        delivery: false,
        type: [],
        resource: "",
        desc: ""
      }
    };
  },
  methods: {
    onSubmit() {
      console.log("submit!");
    }
  }
};
</script>

<style scoped>
.login-form {
    /*display: flex;*/
    /*align-items: center;*/
    /*justify-content: center;*/
}
</style>
