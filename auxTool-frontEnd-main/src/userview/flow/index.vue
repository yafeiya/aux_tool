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
                  @mousedown.right="startDrag($event,item.label)"
              >
                  <template #title> 
                      <span>{{ item.label }}</span>
                  </template>

                      <!-- <MenuItem
                      class="secondmenustyle"
                      :name="item.label"
                      @mousedown="startDrag($event,item.label)"
                      >   
                          <span>{{ item.label }}</span>
                      </MenuItem> -->
                  
                      <MenuItem 
                      class="secondmenustyle"
                      :name="subitem.name"
                      v-for="subitem in noChildren(item.children)"
                      @mousedown="startDrag($event,item.label)"
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
<!-- 
                              <MenuItem
                              class="thirdmenustyle"
                              :name="subitem.label"
                              @mousedown="startDrag($event,item.label)"
                              >   
                                  <span>{{ subitem.label }}</span>
                              </MenuItem> -->

                              <MenuItem
                              class="thirdmenustyle" 
                              :name="lastitem.name"
                              v-for="lastitem in noChildren(subitem.children)"
                              @mousedown="startDrag($event,item.label)"
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
          <Content  class="graphcontent">
            <div id="container" ></div>
          </Content>
          
      </Layout>
      <Sider class="rightsider" style= "max-width:300px;width:300px;flex:content" >
        <config-panel v-if="isReady" />
      </Sider>
    </Layout>
  </div>
</template>

<script lang="ts">
  import { defineComponent, ref, onMounted } from 'vue';
  import { useRouter } from 'vue-router';
  import './reset/reset.less';
  import './reset/global.css';
  import './index.less';
  import FlowGraph from './graph';
  import ToolBar from './components/ToolBar/index.vue';
  import ConfigPanel from './components/ConfigPanel/index.vue';
  import { menulist } from './list'

  
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
      const router = useRouter();
      const maingraph = FlowGraph;
      const isReady = ref(false);
      let fathername = ""
      const list = menulist()

      const initGraph = function () {
        const graph = FlowGraph.init();
        isReady.value = true;
        // const resizeFn = () => {
        //   const { width, height } = getContainerSize();
        //   graph.resize(width, height);
        // };
        // resizeFn();
        // window.addEventListener('resize', resizeFn);
        // return () => {
        //   window.removeEventListener('resize', resizeFn);
        // };
      };


      const startDrag = (e, fatherItem) => {
        // var a = Object.assign({},fatherItem)
        // let b = ""
        // let c = ""
        // const d = fatherItem
        // console.log("fatherItem")
        // console.log(e)
        // console.log(e.target.outerText)
        // console.log(a)
        // console.log(d)
  
        // for (const [key, value] of Object.entries(a)) {
        //     b = b + a[key]
        //     // console.log(b)
        // }
        // c = c+b
        // console.log(c)
        const name = savename(fatherItem)
        // console.log(name)
        var initNode = {
          shape: 'flow-chart-rect',
          width: 80,
          height: 42,
          label: e.target.textContent,
          data:{
            fatherLabel: name,
            selflabel: e.target.textContent,
          }
        } 
        const node = maingraph.graph.createNode(initNode);
        maingraph.dnd.start(node, e);
      };

      onMounted(() => {
          initGraph();
      });

      const savename =(a) =>{
        if (typeof(a) !== "undefined"){
          fathername = a
        }else{
          fathername = fathername
        }
        return fathername;
      }
      
      // 编写菜单栏
      const noChildren =(thelist) =>{
          return thelist.filter((item) => !item.children);
      };
      // 编写菜单栏
      const hasChildren =(thelist) =>{
          return thelist.filter((item) => item.children);
      };


      const backhome = () =>{
        router.push("/home")
      };

      return {
        list,
        isReady,
        savename,
        startDrag,
        noChildren,
        hasChildren,
        backhome,
      };
    },
  });
</script>

<style lang="less" scoped>
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
    height: 1100px;
    background-color: #fff;
  }
  .graphheader{
    height: 40px;
    background-color: #fff;
    }
  .graphcontent{
    height: 1060px;
  }
  .rightsider{
    height: 1100px;
    background-color: #fff;
  }
}
</style>
