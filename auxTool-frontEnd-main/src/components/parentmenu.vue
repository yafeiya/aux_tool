<template>
  <Menu :active-name=this.initNowItem theme="dark" width="auto" :open-names="[this.pageMenu[0].name]" @on-select="getNowItem">

    <Affix v-color="'white'" style="font-size: 20px;margin-left: 18%;margin-top: 8%;margin-bottom: -2%;">
      <Button :type="text" size="small" shape="circle" @click="toHome"><Icon type="md-arrow-back" /></Button> |  返回首页
    </Affix>
    <Divider />

    <!--    <Affix v-color="'white'" style="font-size: 20px;margin-left: 15%;margin-top: 5%;margin-bottom: -4%;">-->
    <!--      <Button ghost size="large" shape="circle"><Icon type="ios-add-circle-outline" style="font-size: 110%"/> 添加类别</Button>-->
    <!--    </Affix>-->
    <!--    <Divider />-->
    <template v-for="item in pageMenu">
      <childMenu v-if="item.children&&item.children.length!==0" :parent-item="item">
      </childMenu>
      <menu-item v-else :name="item.title">
        <Icon :type="item.icon"/>
        <span>{{ item.title }}</span>
      </menu-item>

    </template>
    <!--菜单栏下添加按钮-->
    <Affix v-color="'white'" style="font-size: 20px;margin-left: 15%;margin-top: 5%;margin-bottom: -4%;">
      <Button ghost  shape="circle" style="width:80%;margin-left: 0%"  @click="addMenu=true" >
        <Icon type="ios-add-circle-outline" style="font-size: 110%;margin-left: -6px;"/>
        管理
      </Button>
      <Modal v-model="addMenu" width="400" style="margin-top: 0px" @on-ok="okInfo">
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

<!--          <template v-for="(item, index) in menuform.items">-->
<!--            <FormItem-->
<!--                v-if="item.status"-->
<!--                :key="index"-->
<!--                :label="'任务 ' + item.index"-->
<!--                :prop="'items.' + index + '.value'"-->
<!--                :rules="{required: true, message: '任务名不可为空', trigger: 'blur'}">-->
<!--              <Row>-->
<!--                <Col span="18">-->
<!--                  <Input type="text" v-model="item.value" placeholder="输入任务名"></Input>-->
<!--                </Col>-->
<!--                <Col span="4" offset="1">-->
<!--                  <Button @click="handleRemove(index)">Delete</Button>-->
<!--                </Col>-->
<!--              </Row>-->
<!--            </FormItem>-->
<!--          </template>-->
          <FormItem>
            <Row>
              <Col span="8">
                <Button type="default" long @click="handleMenu(menuform.fartherMenu,menuform.childreMenu,true)" icon="md-add" style="margin-left: 10px">添加</Button>
              </Col>
              <Col span="8">
                <Button type="default" long @click="handleMenu(menuform.fartherMenu,menuform.childreMenu,false)" icon="md-add" style="margin-left: 30px">删除</Button>
              </Col>
            </Row>
          </FormItem>
        </Form>

      </Modal>
    </Affix>

  </Menu>

</template>
<script>
import childMenu from './childmenu.vue';
import {sendMenu} from "../api/api.js";
// import axios from "axios";
import qs from "qs";
import {addMenuItem, isLogin} from '../api/api.js'
export default {
  data() {
    return {
      index: 1,
      menuform:{
        fartherMenu:"",
        childreMenu:""
      },
      rulemenu:{
        fartherMenu: [
          { required: true, message: '名称必填', trigger: 'blur' }
        ],
      },
      initNowItem: "numDataBase1-1-任务1",
      addMenu:false
    }
  },
  name:'parentMenu',
  components: {
    childMenu
  },
  inject:['pageKind'],
  // 将菜单数组传入
  props: ['pageMenu'],
  created() {
    this.initNowItem = this.pageMenu[0].title + "-" + this.pageMenu[0].children[0].title
    // console.info("created: " + this.initNowItem)
    this.nowItemInit()
  },
  methods: {

    handleMenu (fatherMenu, childrenMenu,flag) {
      var data = {
        //菜单的添加与删除
        FatherMenu: fatherMenu,
        ChildrenMenu: childrenMenu,
        Add_or_Del:flag
      }
      data = qs.stringify(data)
      sendMenu(data).then(res => {
            console.info(res)
        // TODO:需要刷新页面或者菜单
    })
    },

    toHome() {
      this.$router.push('/home')
    },
    getNowItem(name) {
      this.$emit('getNowItem',name);//select事件触发后，自动触发getNowItem事件
    },
    nowItemInit() {
      // console.info(this.initNowItem)
      this.$emit('getNowItem',this.initNowItem);
    }
  }
}

</script>
<style>
.dev-run-preview .dev-run-preview-edit{ display: none }
</style>
