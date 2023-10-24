<template>
  <div class="layout">
    <Layout >
      
      <Sider class="leftsider" hide-trigger >
        <Scroll
        height = 1100
        >
        <suspense>
          <Menu v-if="isReady"  />
        </suspense>
      </Scroll>
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
          <Content
          class="cardsize"
          >
          <transition name="move-down">
            <Card
                class="preview"
                v-show="grandpreview" id="miniChart" 
                :style="{
                  width: '1800px', 
                  height: '300px', 
                }"
                :bordered="false"
                >
            </Card>
          </transition>
          <transition name="move-down">
            <Card
                  class="preview"
                  v-show="imgpreview" id="modelpre" 
                  :style="{
                    width: '1800px', 
                    height: '300px', 
                  }"
                  :bordered="false"
                  >
                  <suspense>
                    <Image 
                    class = "img"
                    :src = conbineimgpath(imgpath)
                    fit="fill" 
                    width="500px" 
                    height="300px" 
                    alt=""
                    >
                    </Image>
                  </suspense>
              </Card>
          </transition>
        </Content>
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
  import { defineComponent, ref, onMounted, provide,onBeforeMount} from 'vue';
  import { useRouter } from 'vue-router';
  import './reset/reset.less';
  import './reset/global.css';
  import './index.less';
  import FlowGraph from './graph';
  import ToolBar from './components/ToolBar/index.vue';
  import ConfigPanel from './components/ConfigPanel/index.vue';
  import Menu from './menu.vue';
  import { menulist } from './list'
  import { readlist } from './readlist'
  import * as echarts from 'echarts'
  import datas from "./results/data.json"
  import { getMenuInfo } from '../../api/api.js'
  
  
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
      Menu,
    },

    // data(){
    //   return {
    //     menu:[],
    //   }
    // },
    // beforeCreate(){
    //   getMenuInfo("designmenu").then(res => {
    //     this.menu = res.data
    //     console.info("mainmenu=",this.menu)
    //   })
    // },

    setup() {
      const maingraph = FlowGraph;
      const isReady = ref(false);
      // console.log("fathersetup")

      let fathername = ""
      let selfItem = ""
      let grandname =""
      const heightnum = ref(1050)

      const grandpreview = ref(false)
      const changepreview = function(value1) {
        grandpreview.value = value1

        if(!grandpreview.value) {
          heightnum.value = 1050
        }else 
        {heightnum.value = 750}
        // console.log(grandpreview.value)

        // console.log(heightnum.value)
      }
      provide('changepreview', changepreview)
      provide('grandpreview', grandpreview)

      const imgpreview = ref(false)
      var imgpath = ref("")
      const changeimgpreview = function(value1,value2) {
        imgpreview.value = value1
        imgpath = value2
        // console.log("接受的地址:",imgpath)

        if(!imgpreview.value) {
          heightnum.value = 1050
        }else 
        {heightnum.value = 750}

        // console.log("heightnum is:",heightnum.value)
      }
      provide('changeimgpreview', changeimgpreview)
      provide('imgpreview', imgpreview)
      provide('imgpath', imgpath)

      // "http://127.0.0.1:5173/" +
      const conbineimgpath =(a) =>{
        if(a != null) {
          // console.log(new URL(`http://127.0.0.1:5173/${imgpath}`))
          return new URL(`http://127.0.0.1:5173/${imgpath}`).href;
        }
      }


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


      const startDrag = (e, grandItem,fatherItem,selfItem) => {
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
        const grandname = savegrandname(grandItem)
        const fathername = savefathername(fatherItem)
        const selfname = saveselfname(selfItem)
        // console.log(name)
        var initNode = {
          shape: 'flow-chart-rect',
          width: 80,
          height: 42,
          label: selfname,
          data:{
            grandLabel: grandname,
            fatherLabel: fathername,
            selflabel: selfname,
          }
        } 
        const node = maingraph.graph.createNode(initNode);
        maingraph.dnd.start(node, e);
      };

      provide('startDrag',startDrag)

      // onBeforeMount(async () => {
      //   try {
      //   const response = await getMenuInfo("designmenu");
      //   data.value = response.data; // 将返回的数据赋值给 ref 定义的响应式数据
      //   } catch (error) {
      //    console.error(error);
      //   }
      //   console.log("data=",data.value)
      //   list = readlist(data.value)
      //   console.log("list=",list)    
      // });


      onMounted(() => { 
        initGraph();
      });

      // watch(
      //   [() => heightnum.value],
      //   () => {
      //     // graph.resize(1600,height.value);
      //   },
      //   {
      //     immediate: false,
      //     deep: false,
      //   },
      // );
      const savegrandname =(a) =>{
        if (typeof(a) !== "undefined"){
          grandname = a
        }else{
          grandname = grandname
        }
        return grandname;
      }
      

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

      return {
        isReady,
        startDrag,
        changepreview, 
        grandpreview,
        heightnum,
        changeimgpreview,
        imgpreview,
        imgpath,
        conbineimgpath
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
    
    background-color: #fff;
    z-index: 3;
  }
  .graphheader{
    height: 40px;
    background-color: #fff;
    width: 1800px;
    align-items: center;
    }
  .graphcontent{
    height: 1060px;
    background-color: #fff;
    width: 1800px;
    align-items: center;
  }
  .active{
    height: 760px;
    background-color: #fff;
  }
  .preview{
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #fff;
    z-index: 2;
    margin: 0;
  }
  .rightsider{
    background-color: #fff;
    z-index: 3;
  }
  .img{
    margin: auto;
    align-items: center;
  }
  .cardsize{
    width: 100%;
    align-items: center;
  }
}
</style>
