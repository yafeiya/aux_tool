<template>
  <Menu class="menu" width="187px" background-color="#545c64" text-color="#ffffff" ref="menus"
    :collapse-transition="false" :active-name=opnename :open-names=opnename>
    <MenuItem class="backhome" name="backhome" @click="backhome()">
    <Icon type="md-arrow-back" />
    <span>| 返回首页</span>
    </MenuItem>

    <Space>
      <a @click="clearMenu" style="margin-left: 30px">
        <Icon type="ios-arrow-dropup" /> 收起
      </a>
      <a @click="editMenu = true">
        <Icon type="ios-add-circle-outline" style="margin-left: 30px" /> 管理
      </a>
      <Modal v-model="editMenu" width="400" style="margin-top: 0px">
        <template #header>
          <p style="color:#4d85ea;text-align:center">
            <Icon type="ios-information-circle"></Icon>
            <span>编辑类别</span>
          </p>
        </template>
        <Form ref="menuform" :model="menuform" :label-width="80" :rules="rulemenu" style="width: 400px">
          <FormItem label="一级菜单" prop="fartherMenu">
            <Input v-model="menuform.fartherMenu" placeholder="输入一级菜单名称（必填）" style="width: 250px"></Input>
          </FormItem>
          <FormItem label="二级菜单" prop="childreMenu">
            <Input v-model="menuform.childreMenu" placeholder="输入二级菜单名称" style="width: 250px"></Input>
          </FormItem>
          <FormItem label="三级菜单" prop="grandsonMenu"> <!-- 修正这里 -->
            <Input v-model="menuform.grandsonMenu" placeholder="输入三级菜单名称" style="width: 250px"></Input>
          </FormItem>
          <FormItem>
            <Row>
              <Col span="8">
              <Button type="default" long
                @click="canvasMenuAdd(menuform.fartherMenu, menuform.childreMenu, menuform.grandsonMenu)" icon="md-add"
                style="margin-left: 10px">添加</Button>
              </Col>
              <Col span="8">
              <Button type="default" long
                @click="canvasMenuDelete(menuform.fartherMenu, menuform.childreMenu, menuform.grandsonMenu)"
                icon="md-remove" style="margin-left: 30px">删除</Button>
              </Col>
            </Row>
          </FormItem>
        </Form>

      </Modal>
    </Space>

    <MenuItem class="" v-for="item in noChildren(list)" :name="item.name" @click="">
    <span>{{ item.title }}</span>
    </MenuItem>

    <Submenu :name="item.name" v-for="item in hasChildren(list)">
      <template #title>
        <span>{{ item.title }}</span>
      </template>

      <!-- <MenuItem
                      class="secondmenustyle"
                      :name="item.title"
                      @mousedown="startDrag($event,item.title)"
                      >   
                          <span>{{ item.title }}</span>
                      </MenuItem> -->

      <MenuItem class="secondmenustyle" :name="subitem.name" v-for="subitem in noChildren(item.children)"
        @mousedown="startDrag($event, item.title, subitem.title, subitem.title)">
      <span>{{ subitem.title }}</span>
      </MenuItem>

      <Submenu class="secondmenustyle" :name="subitem.name" v-for="subitem in hasChildren(item.children)"
        @mousedown.right="startDrag($event, item.title, subitem.title, subitem.title)">
        <template #title>
          <span>{{ subitem.title }}</span>
        </template>
        <!-- 
                              <MenuItem
                              class="thirdmenustyle"
                              :name="subitem.title"
                              @mousedown="startDrag($event,item.title)"
                              >   
                                  <span>{{ subitem.title }}</span>
                              </MenuItem> -->

        <MenuItem class="thirdmenustyle" :name="lastitem.name" v-for="lastitem in noChildren(subitem.children)"
          @mousedown="startDrag($event, item.title, subitem.title, lastitem.title)">
        <span>{{ lastitem.title }}</span>
        </MenuItem>
      </Submenu>
    </Submenu>
  </Menu>
</template>

<script lang="ts">
import { defineComponent, ref, provide, inject } from 'vue';
import { useRouter } from 'vue-router';
import './reset/reset.less';
import './reset/global.css';
import './index.less';

import { menulist } from './list'
import { readlist } from './readlist'
import * as echarts from 'echarts'
import datas from "./results/data.json"
import { deletePageMenuItem,addPageMenuItem,getCard, getMenuInfo } from '../../api/api.js'

// const getContainerSize = () => {
//   return {
//     width: document.body.offsetWidth - 590,
//     height: document.body.offsetHeight - 110,
//   };
// };
export default defineComponent({
  name: 'Index',

  async setup() {

    const menuform = {
      fartherMenu: "",
      childreMenu: "",
      grandsonMenu: "",
    };

    const rulemenu = {
      fartherMenu: [
        { required: true, message: '名称必填', trigger: 'blur' }
      ],
    };
    // 画布菜单编辑
    const canvasMenuAdd = (fartherMenu, childreMenu, grandsonMenu) => {
      console.log('添加了:', fartherMenu,childreMenu,grandsonMenu,);
      var data = {
        //菜单的添加与删除
        pageKind: "canvas",
        title: fartherMenu,
        children_title: childreMenu,
        gran_title: grandsonMenu
        
      }
      addPageMenuItem(data).then(response => {
        console.info("!111111111",response.data.msg)
      })
      setTimeout(() => {
          getMenuInfo("canvas").then(res => {
          })
          location.reload();
        }, 500);
    };
    const canvasMenuDelete = (fartherMenu, childreMenu, grandsonMenu) => {
      // 删除菜单项的逻辑
      console.log('删除:', { fartherMenu, childreMenu, grandsonMenu });
      var data = {
        //菜单的添加与删除
        pageKind: "canvas",
        title: fartherMenu,
        children_title: childreMenu,
        gran_title: grandsonMenu
      }
      deletePageMenuItem(data).then(response => {
        console.info("!222222",response.data.msg)
      })
      setTimeout(() => {
          getMenuInfo("canvas").then(res => {
          })
          location.reload();
        }, 500);
      
    };

    const router = useRouter();
    const heightnum = ref(1050)
    const startDrag: any = inject('startDrag')


    // 编写菜单栏
    const noChildren = (thelist) => {
      return thelist.filter((item) => !item.children);

    };
    // 编写菜单栏
    const hasChildren = (thelist) => {
      return thelist.filter((item) => item.children);
    };

    // list是响应数组 画布总菜单目录
    let list: any = ref([])

    let menu: any = (await getMenuInfo("canvas")).data
    let database = (await getMenuInfo("database")).data
    console.log("后端获取的menu数据库数据", menu)
    const item_data = menu.find(menuItem => menuItem.title === "数据加载");
    if (item_data) {
      item_data.children.push(...database)
    }

    let modelbase = (await getMenuInfo("modelbase")).data
    // console.log("模型数据",modelbase)
    const item_model = menu.find(menuItem => menuItem.title === "模型模板");
    if (item_model) {
      item_model.children.push(...modelbase)
    }

    let defineFunction = (await getMenuInfo("defineFunction")).data
    const item_train = menu.find(menuItem => menuItem.title === "训练过程可视化");
    if (item_train) {
      item_train.children.push(...defineFunction[0].children)
    }
    const item_eval = menu.find(menuItem => menuItem.title === "评价指标");
    if (item_eval) {
      item_eval.children.push(...defineFunction[1].children)
    }
    list = await menu

    //收起菜单
    let opnename = ref([])
    function clearMenu() {
      this.opnename = [];
      this.$nextTick(() => {
        this.$refs.menus.updateActiveName()
        this.$refs.menus.updateOpened()
        // console.log("clearMenu opnename", opnename)
      });
    }

    const editMenu = ref(false);
    const backhome = () => {
      router.push("/home")
    };

    console.log("打印list", list)

    return {
      list,
      opnename,
      menuform,
      rulemenu,
      editMenu,
      startDrag,
      clearMenu,
      noChildren,
      hasChildren,
      backhome,
      heightnum,
      canvasMenuAdd,
      canvasMenuDelete
    };
  },
});



</script>

<style lang="less" scoped>
.backhome {
  // text-align: center;s
  margin-left: 22px;
}

.secondmenustyle {
  background-color: #eee;
}

.thirdmenustyle {
  background-color: #ddd;
}
</style>