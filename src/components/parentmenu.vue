<template>
  <Menu :active-name=this.initNowItem theme="dark" width="auto" :open-names="[this.pageMenu[0].name]" @on-select="getNowItem">
    <MenuGroup title="数据集" class="menu-table"></MenuGroup>
    
    <template v-for="item in pageMenu">  
      <childMenu v-if="item.children&&item.children.length!==0" :parent-item="item">
      </childMenu>
      <menu-item v-else :name="item.name + '-' + item.title">
        <Icon :type="item.icon"/>
        <span>{{ item.title }}</span>
      </menu-item>
    </template>
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
    name:'paretMenu',
    components: {
        childMenu
    },
    // 将菜单数组传入
    props: ['pageMenu','pageKind'],
    mounted() {
        this.initNowItem = this.pageMenu[0].children[0].title
        this.nowItemInit()
    },
    methods: {
        getNowItem(name) {
            this.$emit('getNowItem',name);//select事件触发后，自动触发getNowItem事件
        },
        nowItemInit() {
            console.info(this.initNowItem)
            this.$emit('getNowItem',this.initNowItem);
        }
    }
}
        
</script>
<style>
.layout{
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
}
.layout-home{
  width: 100px;
  height: 30px;
  border-radius: 3px;
  float: left;
  position: relative;
  top: 3px;
  left: 0px;
}
.menu-table{
  margin-left: 15%;
  font-size:100px;
}
.layout-nav{
  width: 620px;
  margin: 0 auto;
  margin-right: -80px;
}
.dev-run-preview .dev-run-preview-edit{ display: none }
.main-table {
  margin-top: 2%;
}
</style>