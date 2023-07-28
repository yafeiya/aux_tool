<template>
    <div class="demo-login" :key="login">
      <img src="../assets/图标.png" class="logo-size">
      <h1>数据驱动辅助工具</h1>
      <Login @on-submit="handleSubmit">
          <UserName name="username" />
          <Password name="password" />
          <div class="demo-auto-login">
              <Checkbox v-model="autoLogin" size="large">自动登录</Checkbox>
              <a>忘记密码</a>
          </div>
          <Submit value="123"/>
          <router-link to="/database">Home</router-link>
      </Login>
    </div>
  </template>
  <script>
    export default {
        data () {
            return {
                autoLogin: true
            }
        },
        methods: {
            handleSubmit (valid, { username, password }) {
                if (valid) {
                    this.$Modal.info({
                        title: '输入的内容如下：',
                        content: 'username: ' + username + ' | password: ' + password
                    });
                }
                this.$router.push('/database')
            }
        }
    }
  </script>
  <style>
    .logo-size{
      width: 80px;
      height: 80px;
    }
    .demo-login{
        width: 400px;
        height: 200px;
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        margin: auto;
        /* margin-left: 100px; */
    }
    .demo-auto-login{
        margin-bottom: 24px;
        text-align: left;
    }
    .demo-auto-login a{
        float: right;
    }
  </style>