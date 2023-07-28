import { createRouter,createWebHashHistory } from "vue-router";
import login from '../components/login.vue'
import database from '../components/database.vue'
import regist from '../components/regist.vue'
import home from '../components/home.vue'
import example from '../components/example.vue'
// import List from '../views/List.vue'
// import Detail from '../views/Detail.vue'
const routes = [
    {
        path:'/',
        name:'login',
      //   component:resolve => require(['../components/login.vue'], resolve),
        component:login,
        meta: {
            title: '登录'
        }
    },
    // {
    //   path:'/login',
    //   name:'login',
    // //   component:resolve => require(['../components/login.vue'], resolve),
    //   component:login,
    //   meta: {
    //     title: '登录'
    //   }
    // },
    {
        path:'/regist',
        name:'regist',
      //   component:resolve => require(['../components/login.vue'], resolve),
        component:regist,
        meta: {
            title: '注册'
        }
    },
    {
        path:'/database',
        name:'database',
        //   component:resolve => require(['../components/database.vue'], resolve),
        component:database,
        meta: {
            title: '数据库信息'
      }
    },
    {
        path:'/home',
        name:'home',
        component:home,
        meta: {
            title: '主页'
        }
    },
    {
        path:'/example',
        name:'example',
        //   component:resolve => require(['../components/database.vue'], resolve),
        component:example,
        meta: {
            title: '示例页'
        },
    }
  ]
  const router = createRouter({
    history:createWebHashHistory(),
    routes
  })
  export default router;