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
.layout-nav1{
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
          <parentMenu :pageMenu="this.menu" :pageKind="this.pageKind" @getNowItem="getNowItem">
          </parentMenu>
        </Sider>
        <Layout :style="{marginLeft: '200px'}">
          <!--//上边框菜单栏-->
          <Header>
            <Menu mode="horizontal" theme="dark" active-name="design" @on-select="toPages">
              <div class="layout-home"><Button @click="toHome" type="default" ghost>跳转首页</Button></div>
              <div class="layout-nav1">
                <MenuItem name="database">
                  <Icon type="ios-navigate"></Icon>
                  数据集
                </MenuItem>
                <MenuItem name="modelbase">
                  <Icon type="ios-keypad"></Icon>
                  模型库
                </MenuItem>
                <MenuItem name="defineFunction">
                  <Icon type="ios-analytics"></Icon>
                  自定义函数
                </MenuItem>
                <MenuItem name="design">
                  <Icon type="ios-paper"></Icon>
                  方案设计
                </MenuItem>
                <MenuItem name="example">
                  <Icon type="md-arrow-dropright-circle"></Icon>
                  方案实例
                </MenuItem>
              </div>
            </Menu>
          </Header>
          <Content :style="{padding: '0 16px 16px'}">
            <!--//数据集选项-->
            <!-- <Tabs type="card" class="main-table">
              <TabPane label="我的数据集">这里是我的数据集火力分配1</TabPane>
              <TabPane label="公共数据集">这里是公共数据集火力分配1</TabPane>
            </Tabs> -->
            <mainTable :pageKind = "pageKind" :nowItem="nowItem"/>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  </div>
</template>
<script>
import {MenuGroup} from "view-ui-plus";
import parentMenu from '../components/parentmenu.vue';
import mainTable from '../components/maintable.vue';
export default {
  data() {
    return {
      menu:[
					{
						name: 'A领域',
						title: 'A领域',
						icon: 'ios-navigate',
            children:[
							{
								name: '红蓝对抗火力分配',
								title: '蓝对抗火力分配',
								// icon: 'ios-document-outline',
							}
						]
					},
					{
						name: 'B领域',
						title: 'B领域',
						icon: 'ios-keypad',
            children:[
							{
								name: '作战意图研判',
								title: '作战意图研判',
								// icon: 'ios-document-outline',
							},
                            {
								name: '战术决策',
								title: '战术决策',
								// icon: 'ios-document-outline',
							},
						]
					},
                    {
						name: 'C领域',
						title: 'C领域',
						icon: 'ios-navigate',
            children:[
							{
								name: 'XXXX',
								title: 'XXXX',
								// icon: 'ios-document-outline',
							}
						]
					},
			],
        // pageKind标明当前页的信息（database，modelbase等）
        // nowItem表明选中的是菜单的哪一项。
      pageKind: 'design',
      nowItem: null, 
    }
  },
  components: {
    MenuGroup,
    parentMenu,
    mainTable,
  },
  methods: {
    toHome() {
      this.$router.push('/home')
    },
    getNowItem(nowItem) {
      this.nowItem = nowItem
    },
    toPages(name) {
      var targetUrl = "/" + name
      this.$router.push(targetUrl)
    }
  }
}
</script>
