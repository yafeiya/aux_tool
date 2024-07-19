import { createRouter,createWebHashHistory } from "vue-router";
import login from '../userview/login.vue'
import database from '../userview//database.vue'
import modelbase from '../userview//modelbase.vue'
import defineFunction from '../userview//defineFunction.vue'
import design from '../userview//design.vue'
import example from '../userview//example.vue'
import regist from '../userview//regist.vue'
import home from '../userview//home.vue'
import paint from '../userview/flow/index.vue'
import temp from '../userview/temp.vue'
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
        path:'/modelbase',
        name:'modelbase',
        //   component:resolve => require(['../components/database.vue'], resolve),
        component:modelbase,
        meta: {
            title: '模型库信息'
      }
    },
    {
        path:'/defineFunction',
        name:'defineFunction',
        //   component:resolve => require(['../components/database.vue'], resolve),
        component:defineFunction,
        meta: {
            title: '自定义函数'
      }
    },
    {
        path:'/design',
        name:'design',
        //   component:resolve => require(['../components/database.vue'], resolve),
        component:design,
        meta: {
            title: '方案设计'
      }
    },
    {
        path:'/example',
        name:'example',
        //   component:resolve => require(['../components/database.vue'], resolve),
        component:example,
        meta: {
            title: '方案实例'
      }
    },
    {
      path:'/index',
      name:'index',
      //   component:resolve => require(['../components/database.vue'], resolve),
      component:paint,
      meta: {
          title: '开始设计'
      }
    },
    {
      path:'/temp',
      name:'temp',
      //   component:resolve => require(['../components/database.vue'], resolve),
      component:temp,
      meta: {
          title: 'test'
      }
    },
    {
        path:'/home',
        name:'home',
        component:home,
        meta: {
            title: '主页'
        }
    }
  ]
  const router = createRouter({
    history:createWebHashHistory(),
    routes
  })
  export default router;
