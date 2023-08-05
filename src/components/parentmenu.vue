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
      <Button ghost  shape="circle" style="width:80%;margin-left: 0%" >
        <Icon type="ios-add-circle-outline" style="font-size: 110%;margin-left: -6px;"/>
        添加类别
      </Button>
    </Affix>

  </Menu>

</template>
<script>
import childMenu from './childmenu.vue';
export default {
  data() {
    return {
      initNowItem: "numDataBase1-1-火力分配1"
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
