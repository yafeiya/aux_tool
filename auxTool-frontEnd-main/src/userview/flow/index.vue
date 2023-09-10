<template>
  <div class="layout">
    <Layout>
      <Sider class="leftsider" hide-trigger>
        <Menu
          class="menu"
          width="200px"
          background-color="#545c64"
          text-color="#ffffff"
          :collapse-transition="false"
          >
              <MenuItem
              class="backhome"
              name="backhome"
              @click="backhome()"
              >
                  <Icon type="md-arrow-back" />
                  <span>| 返回首页</span>
              </MenuItem>
          <Space>
            <a @click="clearMenu" style="margin-left: 30px" >
              <Icon type="ios-arrow-dropup" />   收起
            </a>
            <a @click="editMenu=true">
              <Icon type="ios-add-circle-outline" style="margin-left: 30px"/>   管理
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
              <FormItem label="三级菜单" prop="childreMenu">
                <Input v-model="menuform.childreMenu" placeholder="输入三级菜单名称" style="width: 250px"></Input>
              </FormItem>
              <FormItem>
                <Row>
                  <Col span="8">
                    <Button type="default" long @click="" icon="md-add" style="margin-left: 10px">添加</Button>
                  </Col>
                  <Col span="8">
                    <Button type="default" long @click="" icon="md-add" style="margin-left: 30px">删除</Button>
                  </Col>
                </Row>
              </FormItem>
            </Form>

          </Modal>
          </Space>
              <MenuItem
              class=""
              v-for="item in noChildren(list)"
              :name="item.name"
              @click=""
              >
                  <span>{{ item.label }}</span>
              </MenuItem>

              <Submenu
                  :name="item.name"
                  v-for="item in hasChildren(list)"
                  @mousedown.right="startDrag($event,item.label,item.label)"
              >
                  <template #title>
                      <span>{{ item.label }}</span>
                  </template>

                      <MenuItem
                      class="secondmenustyle"
                      :name="subitem.name"
                      v-for="subitem in noChildren(item.children)"
                      @mousedown="startDrag($event,item.label,subitem.label)"
                      >
                          <span>{{ subitem.label }}</span>
                      </MenuItem>

                      <Submenu
                          class="secondmenustyle"
                          :name="subitem.name"
                          v-for="subitem in hasChildren(item.children)"
                      >
                          <template #title>
                              <span>{{ subitem.label }}</span>
                          </template>

                              <MenuItem
                              class="thirdmenustyle"
                              :name="lastitem.name"
                              v-for="lastitem in noChildren(subitem.children)"
                              @mousedown="startDrag($event,item.label,lastitem.label)"
                              >
                                  <span>{{ lastitem.label }}</span>
                              </MenuItem>
                      </Submenu>

              </Submenu>

          </Menu>
      </Sider>
      <Layout>
          <Header class="graphheader">
            <tool-bar v-if="isReady" />

          </Header>

          <Content
          class="graphcontent"
          :class="{ active: grandpreview }"
          >
            <div
            id="container"
            >

          </div>
          </Content>
          <transition name="move-down">
            <Card
                class="preview"
                v-show="grandpreview" id="miniChart"
                :style="{
                  width: '1500px',
                  height: '300px',
                }"
                :bordered="false"
                >
            </Card>
          </transition>
      </Layout>
      <Sider
      class="rightsider"
      style= "max-width:300px;width:300px;flex:content;" >
        <config-panel v-if="isReady" />
      </Sider>
    </Layout>
  </div>
</template>

<script lang="ts">
  import { defineComponent, ref, onMounted, provide, reactive} from 'vue';
  import { useRouter } from 'vue-router';
  import './reset/reset.less';
  import './reset/global.css';
  import './index.less';
  import FlowGraph from './graph';
  import ToolBar from './components/ToolBar/index.vue';
  import ConfigPanel from './components/ConfigPanel/index.vue';
  import { menulist } from './list'
  import * as echarts from 'echarts'
  import datas from "./results/data.json"
  import {getMenuInfo} from '../../api/api.js'


  // const getContainerSize = () => {
  //   return {
  //     width: document.body.offsetWidth - 590,
  //     height: document.body.offsetHeight - 110,
  //   };
  // };
  export default defineComponent({

    name: 'Index',
    components: {
      ToolBar,
      ConfigPanel,
    },

    setup() {

      const menuform={
        fartherMenu:"",
        childreMenu:"",
      };

      const rulemenu={
        fartherMenu: [
          { required: true, message: '名称必填', trigger: 'blur' }
        ],
      };

      const router = useRouter();
      const maingraph = FlowGraph;
      const isReady = ref(false);
      let fathername = ""
      let selfItem = ""
      const heightnum = ref(1050)

      const getInTheatersData = async () => {
      const res = await getMenuInfo("database");
      console.info("88888888888888888",res )};

     let list = ref(menulist())
      function clearMenu(){
        list.value=[]
        setTimeout(() => {
          list.value=menulist()}, 1);
      }

      const grandpreview = ref(false)
      const changepreview = function(value1) {
        grandpreview.value = value1

        if(!grandpreview.value) {
          heightnum.value = 1050
        }else
        {heightnum.value = 750}
        console.log(grandpreview.value)

        console.log(heightnum.value)
      }
      provide('changepreview', changepreview)
      provide('grandpreview', grandpreview)



      const initGraph = function () {
        const graph = FlowGraph.init();
        isReady.value = true;

      };


      const startDrag = (e, fatherItem,selfItem) => {

        const fathername = savefathername(fatherItem)
        const selfname = saveselfname(selfItem)
        // console.log(name)
        var initNode = {
          shape: 'flow-chart-rect',
          width: 80,
          height: 42,
          label: selfname,
          data:{
            fatherLabel: fathername,
            selflabel: selfname,
          }
        }
        const node = maingraph.graph.createNode(initNode);
        maingraph.dnd.start(node, e);
      };

      onMounted(() => {
        initGraph();
        // getInTheatersData();
      });

      const savefathername =(a) =>{
        if (typeof(a) !== "undefined"){
          fathername = a
        }else{
          fathername = fathername
        }
        return fathername;
      }

      const saveselfname =(a) =>{
        if (typeof(a) !== "undefined"){
          selfItem = a
        }else{
          selfItem = selfItem
        }
        return selfItem;
      }
      // 编写菜单栏
      const noChildren =(thelist) =>{
          return thelist.filter((item) => !item.children);
      };
      // 编写菜单栏
      const hasChildren =(thelist) =>{
          return thelist.filter((item) => item.children);
      };

      const editMenu = ref(false);
      const backhome = () =>{
        router.push("/home")
      };

      return {
        list,
        isReady,
        menuform,
        rulemenu,
        startDrag,
        noChildren,
        hasChildren,
        editMenu,
        clearMenu,
        backhome,
        changepreview,
        grandpreview,
        heightnum,
      };
    },
  });
</script>

<style lang="less" scoped>
.editmenu{
  // text-align: center;s
  margin-left: 22px;
}
.backhome{
  // text-align: center;s
  margin-left: 22px;
}
.secondmenustyle{
  background-color:#eee;
}
.thirdmenustyle{
  background-color:#ddd;
}

.layout{
  height: 100%;
  .leftsider{

    background-color: #fff;
    z-index: 3;
  }
  .graphheader{
    height: 40px;
    background-color: #fff;
    }
  .graphcontent{
    height: 1060px;
    background-color: #fff;
  }
  .active{
    height: 760px;
    background-color: #fff;
  }
  .preview{

    align-items: center;
    background-color: #fff;
    z-index: 2;
  }
  .rightsider{
    background-color: #fff;
    z-index: 3;
  }
}
</style>
