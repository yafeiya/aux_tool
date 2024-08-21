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
          <Menu active-name="1-2" theme="dark" width="auto" :open-names="['1']">
            <MenuGroup title="数据集" class="menu-table"></MenuGroup>
            <Submenu name="1">
              <template #title>
                <Icon type="ios-navigate"></Icon>
                数值数据集
              </template>
              <MenuItem name="1-1">火力分配1</MenuItem>
              <MenuItem name="1-2">火力分配2</MenuItem>
            </Submenu>
            <Submenu name="2">
              <template #title>
                <Icon type="ios-keypad"></Icon>
                图像数据集
              </template>
              <MenuItem name="2-1">目标检测</MenuItem>
              <MenuItem name="2-2">目标分类</MenuItem>
            </Submenu>
            <Submenu name="3">
              <template #title>
                <Icon type="ios-analytics"></Icon>
                其他数据集
              </template>
              <MenuItem name="3-1">待补充</MenuItem>
            </Submenu>
          </Menu>
        </Sider>
        <Layout :style="{marginLeft: '200px'}">
          <!--//上边框菜单栏-->
          <Header>
            <Menu mode="horizontal" theme="dark" active-name="1">
              <div class="layout-home"><Button @click="go" type="default" ghost>跳转首页</Button></div>
              <div class="layout-nav">
                <MenuItem name="1">
                  <Icon type="ios-navigate"></Icon>
                  数据集
                </MenuItem>
                <MenuItem name="2">
                  <Icon type="ios-keypad"></Icon>
                  算法库
                </MenuItem>
                <MenuItem name="3">
                  <Icon type="ios-analytics"></Icon>
                  自定义函数
                </MenuItem>
                <MenuItem name="4">
                  <Icon type="ios-paper"></Icon>
                  模型设计
                </MenuItem>
                <MenuItem name="5">
                  <Icon type="md-arrow-dropright-circle"></Icon>
                  实例管理
                </MenuItem>
              </div>
            </Menu>
          </Header>
          <Content :style="{padding: '0 16px 16px'}">
            <!--//数据集选项-->
            <Tabs type="card" class="main-table">
              <TabPane label="我的数据集">这里是我的数据集阿巴阿巴</TabPane>
              <TabPane label="公共数据集">这里是公共数据集阿巴阿巴</TabPane>
            </Tabs>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  </div>
</template>
<script>
import {MenuGroup} from "view-ui-plus";

export default {
  components: {MenuGroup},
  methods: {
    go() {
      this.$router.push('/')
    }
  }
}
</script>
