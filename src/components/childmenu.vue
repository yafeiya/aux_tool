<template>
    <Submenu :name="parentItem.name">
        
        <template #title>
            <Icon :type="parentItem.icon"/>
            <span>{{ parentItem.title }}</span>
        </template>
        <template v-for="item in parentItem.children">
            <childMenu 
                v-if="item.children&&item.children.length!==0" 
                :parent-item="item" 
            >
            </childMenu>
            <menu-item v-else :name="item.title">
                <Icon :type="item.icon"/>
                <span>{{ item.title }}</span>
            </menu-item>
        </template>
    </Submenu>
</template>
 
<script>
	export default{
		name:'childMenu',
		props: {
			parentItem: {
				type: Object,
				default: () => {}
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