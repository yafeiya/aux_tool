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
              :index="item.path" 
              v-for="item in noChildren()"
              :name="item.name"
              @click="changeMenu(item)"
              >   
                  <Icon type="md-arrow-back" />
                  <span>{{ item.label }}</span>
              </MenuItem>
              <Submenu
                  :name="item.label"
                  v-for="item in hasChildren()"
                  :key="item.path"
              >
                  <template #title> 
                      <span>{{ item.label }}</span>
                  </template>
                  <MenuGroup
                  class="MenuGroup" 
                  :title="item.label"
                  @mousedown="startDrag"
                  >
                      <MenuItem 
                      :name="item.name"
                      v-for="(subitem, subIndex) in item.children"
                      :key="subIndex"
                      @mousedown="startDrag"
                      >
                          <span>{{ subitem.label }}</span>
                      </MenuItem> 
                  </MenuGroup>
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

      const startDrag = (e) => {
      // console.log(e)
      // console.log(e.target.outerText)
      const node = maingraph.graph.createNode({
        shape: 'flow-chart-rect',
        width: 80,
        height: 42,
        label: e.target.outerText,
      });
        maingraph.dnd.start(node, e);
      };

      onMounted(() => {
          initGraph();
      });

      const list = [
          {
              path: '/home',
              name: 'home',
              label: '| 返回首页',
              icon: '',
              url: '/home'
          },
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
                      name: 'SVD',
                      label: 'SVD',
                      icon: '',
                      url: ''
                  },
                  {
                      name: 'FCN',
                      label: 'FCN',
                      icon: '',
                      url: ''
                  },
                  {
                      name: 'YOLOv3',
                      label: 'YOLOv3',
                      icon: '',
                      url: ''
                  },
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
          {
              name: 'moudlepredictive',
              label: '模型预测',
              icon: '',
              url: '',
              children: [
                  {
                      name: '',
                      label: '',
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
                      name: '',
                      label: '',
                      icon: '',
                      url: ''
                  },
              ]
          },
          
      ];
      // 编写菜单栏
      const noChildren =() =>{
          return list.filter((item) => !item.children);
      };
      // 编写菜单栏
      const hasChildren =() =>{
          return list.filter((item) => item.children);
      };

      const changeMenu = (item) =>{
                router.push({
                    name:item.name,
                })
            };

      return {
        isReady,
        startDrag,
        noChildren,
        hasChildren,
        changeMenu,
      };
    },
  });
</script>

<style lang="less" scoped>
.backhome{
  // text-align: center;s
  margin-left: 22px;
}
.MenuGroup{
  background-color:#eee;
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
