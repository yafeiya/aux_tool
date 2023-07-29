<template>
  <Menu active-name="numDataBase1-1-火力分配1" theme="dark" width="auto" :open-names="['numDataBase']" @on-select="getMainTable">
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
        name:'paretMenu',
		components: {
		    childMenu
		},
        // 将菜单数组传入
        props: ['pageMenu','pageKind'],
        methods: {
            getMainTable(name) {
                console.info(name)
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