<template>
  <!-- <Alert type="success" show-icon>
      A success prompt
      <template #desc>Content of prompt. Content of prompt. Content of prompt. Content of prompt. </template>
  </Alert> -->
  <div>
    <img src="../assets/背景.jpg" width="100%" height="100%" alt=""  class="background"/>
    <div class="demo-login">
      <Card :bordered="false" style="position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%);width: 28%">
        <template #title><img src="../assets/图标.png" class="logo-size"></template>
        <Login @on-submit="handleSubmit">
          <UserName name="username" />
          <Password name="password" />
          <div class="demo-auto-login">
            <Checkbox v-model="autoLogin" size="large">自动登录</Checkbox>
            <a>忘记密码</a>
          </div>
          <Submit>登录</Submit>

        </Login>
        <Login class="regist-btn" @on-submit="toRegist">
          <Submit>注册</Submit>
        </Login>
       <router-link to="/home">Home</router-link>
      </Card>
    </div>
  </div>
</template>
<script>
import axios from 'axios';
import { isLogin } from '../api/api.js'
import qs from "qs";
export default {
  data () {
    return {
      autoLogin: true,
      username: null,
    }
  },
  methods: {

    handleSubmit (valid, { username, password }) {
      if (valid) {
        var data = {
          name: username,
          pwd: password
        }
        data = qs.stringify(data)
        isLogin(data).then(res => {
          this.result = res.msg
          console.info(res)
          if(res.data.msg == "The username does not exist"){
            this.$Message["error"]({
              background: true,
              content: "用户名不存在"
            });
            console.info("The username does not exist")
          }else if(res.data.msg == "Password error"){
            console.info("Password error")
            this.$Message["error"]({
              background: true,
              content: "密码错误"
            });
          }else{
            this.$router.push('/home')
          }
        })
      }

    },
    toRegist(valid) {
      this.$router.push('/regist')
    },

  }
}
</script>
<style>
.logo-size{
  width: 100%;
  height: 100%;
}
.demo-login{
  padding: 10px;
  height: 500px;
  width: 30%;
  margin-left: auto;
  margin-right: auto;

}
.demo-auto-login{
  margin-bottom: 24px;
  text-align: left;
}
.demo-auto-login a{
  float: right;
}
.regist-btn{
  margin-top: 2%;
}
.background{
  width:100%;
  height:100%;  /**宽高100%是为了图片铺满屏幕 */
  z-index:-1;
  position: absolute;
}
</style>
