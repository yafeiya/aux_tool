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
<template>
  <div class="layout">
    <Layout>
      <Layout>
        <!--//固定侧边菜单-->
        <Sider :style="{position: 'fixed', height: '100vh', left: 0, overflow: 'auto'}">
          <!-- pageKind={"database","defineFunction","design","example"} -->
          <parentMenu :pageMenu="menu" :pageKind="database">
          </parentMenu>
        </Sider>
        <Layout :style="{marginLeft: '200px'}">
          <!--//上边框菜单栏-->
          <Header>
            <Menu mode="horizontal" theme="dark" active-name="1">
              <div class="layout-home"><Button @click="toHome" type="default" ghost>跳转首页</Button></div>
              <div class="layout-nav">
                <MenuItem name="1">
                  <Icon type="ios-navigate"></Icon>
                  数据集
                </MenuItem>
                <MenuItem name="2">
                  <Icon type="ios-keypad"></Icon>
                  模型库
                </MenuItem>
                <MenuItem name="3">
                  <Icon type="ios-analytics"></Icon>
                  自定义函数
                </MenuItem>
                <MenuItem name="4">
                  <Icon type="ios-paper"></Icon>
                  方案设计
                </MenuItem>
                <MenuItem name="5">
                  <Icon type="md-arrow-dropright-circle"></Icon>
                  方案实例
                </MenuItem>
              </div>
            </Menu>
          </Header>
          <Content :style="{padding: '0 16px 16px'}">
            <!--//数据集选项-->
            <Tabs type="card" class="main-table">
              <TabPane label="我的数据集">这里是我的数据集火力分配1</TabPane>
              <TabPane label="公共数据集">这里是公共数据集火力分配1</TabPane>
            </Tabs>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  </div>
</template>
<script>
import {MenuGroup} from "view-ui-plus";
import parentMenu from '../components/parentmenu.vue';
export default {
  data() {
    return {
      menu:[
					{
						name: 'numDataBase',
						title: '数值数据集',
						icon: 'ios-navigate',
            children:[
							{
								name: 'numDataBase1-1',
								title: '火力分配1',
								// icon: 'ios-document-outline',
							},{
								name: 'numDataBase1-2',
								title: '火力分配2',
								// icon: 'md-bulb',
							}
						]
					},
					{
						name: 'imgDataBase',
						title: '图像数据集',
						icon: 'ios-keypad',
            children:[
							{
								name: 'imgDataBase1-1',
								title: '目标检测',
								// icon: 'ios-document-outline',
							},{
								name: 'imgDataBase1-2',
								title: '目标分类',
								// icon: 'md-bulb',
							}
						]
					},
					{
						name: 'others',
						title: '其他数据集',
						icon: 'ios-analytics',
						children:[
							{
								name: '',
								title: '待补充',
								// icon: 'ios-document-outline',
							}
						]
					},
				],
    }
  },
  components: {
    MenuGroup,
    parentMenu,
  },
  methods: {
    toHome() {
      this.$router.push('/home')
    }
  }
}
</script>
