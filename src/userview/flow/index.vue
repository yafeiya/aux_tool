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
              v-for="item in noChildren()"
              :name="item.name"
              @click=""
              >   
                  <span>{{ item.label }}</span>
              </MenuItem>

              <Submenu
                  :name="item.name"
                  v-for="item in hasChildren()"
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
                      v-for="subitem in nohasChildren(item.children)"
                      @mousedown="startDrag($event,item.label)"
                      >
                          <span>{{ subitem.label }}</span>
                      </MenuItem> 

                      <Submenu
                          class="secondmenustyle"
                          :name="subitem.name"
                          v-for="subitem in stillhasChildren(item.children)"
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
                              v-for="lastitem in nohasChildren(subitem.children)"
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
  import { Addon } from '@antv/x6'

  
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
      const list = [
          // {
          //     path: '/home',
          //     name: 'home',
          //     label: '| 返回首页',
          //     icon: '',
          //     url: '/home',
          // },
          {
              name: 'dataloading',
              label: '数据加载',
              icon: '',
              children: [
                  {
                      name: 'csvdir',
                      label: 'csv文件',
                      icon: '',
                      url: ''
                  },
                  {
                      name: 'imgdir',
                      label: '图片文件夹',
                      icon: '',
                      url: ''
                  },
                  {
                      name: 'numdataset',
                      label: '数值数据集',
                      icon: '',
                      url: ''
                  },
                  {
                      name: 'imgdataset',
                      label: '图像数据集',
                      icon: '',
                      url: ''
                  }
              ]

          },
          {
              name: 'datapre',
              label: '数据预处理',
              icon: '',
              url: '',
              children: [
                  {
                      name: 'binary',
                      label: '二值化',
                      icon: '',
                      url: ''
                  },
                  {
                      name: 'downsample',
                      label: '下采样',
                      icon: '',
                      url: ''
                  },
                  {
                      name: 'dataaugmentation',
                      label: '数据扩充',
                      icon: '',
                      url: ''
                  },
              ]
          },
          {
              name: 'moudles',
              label: '模型模板',
              icon: '',
              url: '',
              children: [
                  {
                    name: 'machinelearning',
                    label: '机器学习',
                    icon: '',
                    url: '',
                    children: [
                        {
                          name: 'SVD',
                          label: 'SVD',
                          icon: '',
                          url: ''
                      },
                    ]
                  },
                  {
                    name: 'deeplearning',
                    label: '深度学习',
                    icon: '',
                    url: '',
                    children: [
                        {
                          name: 'YOLOv3',
                          label: 'YOLOv3',
                          icon: '',
                          url: ''
                      },
                      {
                          name: 'FCN',
                          label: 'FCN',
                          icon: '',
                          url: ''
                      },
                    ]
                  },
                  {
                    name: 'reinforcementlearning',
                    label: '强化学习',
                    icon: '',
                    url: '',
                    children:[                  
                      {
                          name: 'SAC',
                          label: 'SAC',
                          icon: '',
                          url: ''
                      },
                      {
                          name: 'DQN',
                          label: 'DQN',
                          icon: '',
                          url: ''
                      },
                    ]
                  },
              ]
          },
          {
              name: 'moudletrain',
              label: '模型训练',
              icon: '',
              url: '',
              children: [
                  {
                      name: '模型训练',
                      label: '模型训练',
                      icon: '',
                      url: ''
                  },
              ]
          },
          {
              name: 'moudlepredictive',
              label: '模型预测',
              icon: '',
              url: '',
              children: [
                  {
                      name: 'batchdeduction',
                      label: '批量推演',
                      icon: '',
                      url: ''
                  },
              ]
          },
          {
              name: 'resultvisualization',
              label: '模型结果可视化',
              icon: '',
              url: '',
              children: [
                  {
                      name: 'tensorBoard',
                      label: 'tensorBoard',
                      icon: '',
                      url: ''
                  },
              ]
          },
          {
              name: 'simulation',
              label: '仿真交互',
              icon: '',
              url: '',
              children: [
                  {
                      name: 'UDPmonitor',
                      label: 'UDP监听',
                      icon: '',
                      url: ''
                  },
              ]
          },
          
      ];

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
          label: e.target.outerText,
          data:{
            fatherLabel: name,
            selflabel: e.target.outerText,
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
      const noChildren =() =>{
          return list.filter((item) => !item.children);
      };
      // 编写菜单栏
      const hasChildren =() =>{
          return list.filter((item) => item.children);
      };

      const stillhasChildren =(thelist) =>{
          return thelist.filter((item) => item.children);
      };

      const nohasChildren =(thelist) =>{
          return thelist.filter((item) => !item.children);
      };




      const backhome = () =>{
        router.push("/home")
      };

      return {
        isReady,
        savename,
        startDrag,
        noChildren,
        hasChildren,
        backhome,
        stillhasChildren,
        nohasChildren,
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
