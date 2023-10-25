<template>
    <Menu
          class="menu"
          width="187px"
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
                      <Input v-model="menuform.grandsonMenu" placeholder="输入三级菜单名称" style="width: 250px"></Input>
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
                  <span>{{ item.title }}</span>
              </MenuItem>

              <Submenu
                  :name="item.name"
                  v-for="item in hasChildren(list)"

              >
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

                      <MenuItem
                      class="secondmenustyle"
                      :name="subitem.name"
                      v-for="subitem in noChildren(item.children)"
                      @mousedown="startDrag($event,item.title,subitem.title,subitem.title)"
                      >
                          <span>{{ subitem.title }}</span>
                      </MenuItem>

                      <Submenu
                          class="secondmenustyle"
                          :name="subitem.name"
                          v-for="subitem in hasChildren(item.children)"
                          @mousedown.right="startDrag($event,item.title,subitem.title,subitem.title)"
                      >
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

                              <MenuItem
                              class="thirdmenustyle"
                              :name="lastitem.name"
                              v-for="lastitem in noChildren(subitem.children)"
                              @mousedown="startDrag($event,item.title,subitem.title,lastitem.title)"
                              >
                                  <span>{{ lastitem.title }}</span>
                              </MenuItem>
                      </Submenu>

              </Submenu>
          </Menu>
</template>

<script lang="ts">
  import { defineComponent, ref, provide, inject} from 'vue';
  import { useRouter } from 'vue-router';
  import './reset/reset.less';
  import './reset/global.css';
  import './index.less';

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

    async setup() {

      const menuform={
        fartherMenu:"",
        childreMenu:"",
        grandsonMenu:"",
      };

      const rulemenu={
        fartherMenu: [
          { required: true, message: '名称必填', trigger: 'blur' }
        ],
      };
      const router = useRouter();
      const heightnum = ref(1050)
      const startDrag: any = inject('startDrag')


      // 编写菜单栏
      const noChildren =(thelist) =>{
        // console.log("type is",typeof thelist)
        // console.log("innoChildren",thelist)
        // console.log("result",thelist.filter((item) => !item.children))
          return thelist.filter((item) => !item.children);

      };
      // 编写菜单栏
      const hasChildren =(thelist) =>{
          return thelist.filter((item) => item.children);
      };


      // const list = inject('list')
      // let list = ref(menulist())

      let list:any  = ref([])
      let menu:any = [
        {
          "name" :"database",
          "title": "数据加载",
          "children": []
        },
        {
          "name" :"datapre",
          "title": "数据预处理",
          "children": []
        },
        {
          "name" :"modelbase",
          "title": "模型模板",
          "children": []

        },
        {
          "name" :"moudlepredictive",
          "title": "模型预测",
          "children": []
        },
        {
          "name" :"resultvisualization",
          "title": "训练过程可视化",
          "children": []
        },
        {
          "name" :"scorefunction",
          "title": "评价指标",
          "children": []
        }

        ]
      let database = await getMenuInfo("database")
      menu[0]["children"] = database.data
      menu[0]["children"].push({
                            "name": "monitordata",
                            "title": "仿真监听数据"
                          })
      menu[0]["children"].push({
                            "name": "modellog",
                            "title": "模型日志数据"
                          })
      let design = await getMenuInfo("design")
      menu[1]["children"] = design.data[0].children
      let modelbase = await getMenuInfo("modelbase")
      menu[2]["children"]= modelbase.data[0].children.concat(modelbase.data[1].children.concat(modelbase.data[2].children))

      menu[3]["children"] = design.data[1].children
      let defineFunction = await getMenuInfo("defineFunction")
      menu[4]["children"] = defineFunction.data[0].children
      menu[5]["children"] = defineFunction.data[1].children

      list = menu

      function clearMenu(){
        list.data = []
        setTimeout(() => {
          list.data = menu
        }, 1);
        console.info("收起ssssssssss")
      }



    //   onMounted(async () => {
    //     try {
    //     const response = await getMenuInfo("designmenu");
    //     data.value = response.data; // 将返回的数据赋值给 ref 定义的响应式数据
    //     } catch (error) {
    //      console.error(error);
    //     }
    //     console.log("data=",data.value)
    //     list = readlist(data.value)
    //     console.log("list=",list)
    //   });



      const editMenu = ref(false);
      const backhome = () =>{
        router.push("/home")
      };

      // console.log("sonsetup")

      return {
        list,
        menuform,
        rulemenu,
        editMenu,
        startDrag,
        clearMenu,
        noChildren,
        hasChildren,
        backhome,
        heightnum,

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
</style>
