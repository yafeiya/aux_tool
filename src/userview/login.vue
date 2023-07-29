<template>
    <!-- <Alert type="success" show-icon>
        A success prompt
        <template #desc>Content of prompt. Content of prompt. Content of prompt. Content of prompt. </template>
    </Alert> -->
    <div class="demo-login" :key="login">
      <img src="../assets/图标.png" class="logo-size">
      <!-- <h1>数据驱动辅助工具</h1> -->
      <Login @on-submit="handleSubmit">
          <UserName name="username" />
          <Password name="password" />
          <div class="demo-auto-login">
              <Checkbox v-model="autoLogin" size="large">自动登录</Checkbox>
              <a>忘记密码</a>
          </div>
          <Submit/>
          
      </Login>
      <Login class="regist-btn" @on-submit="toRegist">
          
          <Submit>注册</Submit>
          
      </Login>
      <router-link to="/home">Home</router-link>
    </div>
  </template>
  <script>
    import axios from 'axios';
    export default {
        data () {
            return {
                autoLogin: true,
                jsonBaseUrl: "http://localhost:3000",
                username: null
            }
        },
        methods: {
            handleSubmit (valid, { username, password }) {
                if (valid) {
                    // this.$Modal.info({
                    //     title: '输入的内容如下：',
                    //     content: 'username: ' + username + ' | password: ' + password
                    // });
                    axios.get(this.jsonBaseUrl + "/users?username="+username).then(response => {
                        console.log(response);
                        var user = response.data
                        if(user.length === 0) {
                            this.$Message["error"]({
                                background: true,
                                content: '登录失败，此用户名不存在'
                            });
                            console.info("username is null")
                        } else {
                            var valPassword = user[0].password
                            if(valPassword === password) {
                                this.$Message["success"]({
                                    background: true,
                                    content: '登录成功，用户'+username
                                });
                                this.$router.push('/home?username=' + username)
                            } else {
                                this.$Message["error"]({
                                    background: true,
                                    content: '登录失败，密码错误，请仔细检查'
                                });
                                console.info("密码错误，登录失败")
                            }
                        }
                        var target = 123
                        
                        // users
                    })
                    .catch(error => {
                        console.log(error);
                    })
                    // console.info(users)
                    // console.log(users)
                    // this.$router.push('/database')
                }
                
            },
            toRegist(valid) {
                this.$router.push('/regist')
            }
        }
    }
  </script>
  <style>
    .logo-size{
      width: 380px;
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
    .regist-btn{
        margin-top: 2%;
    }
  </style>