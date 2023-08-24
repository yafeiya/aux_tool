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
        添加
      </Button>
      <Modal v-model="addMenu" width="500" style="margin-top: 0px" @on-ok="okInfo">
        <template #header>
          <p style="color:#4d85ea;text-align:center">
            <Icon type="ios-information-circle"></Icon>
            <span>添加类别</span>
          </p>
        </template>
        <Form ref="menuform" :model="menuform" :label-width="80" :rules="rulemenu" style="width: 400px">
          <FormItem label="类别名" prop="name">
            <Input v-model="menuform.name" placeholder="输入类别名称（必填）"></Input>
          </FormItem>

          <template v-for="(item, index) in menuform.items">
            <FormItem
                v-if="item.status"
                :key="index"
                :label="'任务 ' + item.index"
                :prop="'items.' + index + '.value'"
                :rules="{required: true, message: '任务名不可为空', trigger: 'blur'}">
              <Row>
                <Col span="18">
                  <Input type="text" v-model="item.value" placeholder="输入任务名"></Input>
                </Col>
                <Col span="4" offset="1">
                  <Button @click="handleRemove(index)">Delete</Button>
                </Col>
              </Row>
            </FormItem>
          </template>
          <FormItem>
            <Row>
              <Col span="12">
                <Button type="dashed" long @click="handleAdd" icon="md-add">添加子任务</Button>
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
export default {
  data() {
    return {
      index: 1,
      menuform:{
        name:"",
        items: [
          {
            value: '',
            index: 1,
            status: 1
          }
        ]
      },
      rulemenu:{
        name: [
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
    handleAdd () {
      this.index++;
      this.menuform.items.push({
        value: '',
        index: this.index,
        status: 1
      });
    },
    handleRemove (index) {
      this.menuform.items[index].status = 0;
    },

    okInfo(){
      this.$Message.info('添加成功');
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
