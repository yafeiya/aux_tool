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
            <Menu mode="horizontal" theme="dark" active-name="defineFunction" @on-select="toPages">
              <div class="layout-home">
                <p style="font-size: 20px;color: #ffffff;width: 250px">
                  <Icon type="md-cog" class="ivu-anim-loop" size="23" />
                  数据驱动辅助工具
                </p>
              </div>
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
            <mainTable :nowItem="nowItem" :task-type="taskType"
                       :my-card-list="myCardList" :public-card-list="publicCardList"
                       :my-card-num="myCardNum" :my-card-row-num="myCardRowNum" :my-card-col-num="myCardColNum"
                       :public-card-num="publicCardNum" :public-card-row-num="publicCardRowNum" :public-card-col-num="publicCardColNum"
                       :add-form-item="addFormItem"/>
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
import axios from 'axios';
export default {
  data() {
    return {
      menu:[
        {
          name: 'lossFunction',
          title: '损失函数',
          icon: 'ios-navigate',
          children:[
            {
              name: 'L1',
              title: 'L1',
              // icon: 'ios-document-outline',
            },{
              name: 'L2',
              title: 'L2',
              // icon: 'md-bulb',
            }
          ]
        },
        {
          name: 'evalution',
          title: '评价指标',
          icon: 'ios-keypad',
          children:[
            {
              name: 'reward',
              title: 'reward',
              // icon: 'ios-document-outline',
            },
          ]
        },
      ],
      // pageKind标明当前页的信息（database，modelbase等）
      // menuInfo表明选中的是菜单的哪一项。
      addFormItemCfg:[
        {
          title: '名称',
          name: 'dataset_name',
          value: {dataset_name: ''},
          default: false,
          itemType: 'input',
          isEditOnly: true,
          others:["请输入数据集名称..."]
        },
        {
          title: '参数',
          name: 'params',
          value: {params: ''},
          default: false,
          itemType: 'dynamicInput',
          isEditOnly: true,
          others:[['',''],['','']]
        },
        {
          title: '别名',
          name: 'alias',
          value: {alias: ''},
          default: false,
          itemType: 'input',
          isEditOnly: true,
          others:["请输入别名..."]
        },
        {
          title: '级别',
          name: 'rank',
          value: {rank: ''},
          default: false,
          itemType: 'select',
          isEditOnly: false,
          others:[{'value':'1级', 'text': '级别1'},{'value':'2级', 'text': '级别2'},{'value':'3级', 'text': '级别3'}]
        },
        {
          title: '类型',
          name: 'type',
          value: {type: ''},
          default: true,
          itemType: 'input',
          isEditOnly: false,
          others:['taskType']
        },
        {
          title: '范数',
          name: 'task',
          value: {task: ''},
          default: true,
          itemType: 'input',
          isEditOnly: false,
          others:['nowItem']
        },
        {
          title: '函数体',
          name: 'function_body',
          value: {function_body: ''},
          default: false,
          itemType: 'input',
          isEditOnly: true,
          others:["请输入函数体..."]
        },
        {
          title: '语言',
          name: 'lan',
          value: {lan: ''},
          default: false,
          itemType: 'select',
          isEditOnly: false,
          others:[{'value':'C++', 'text': 'C++'},{'value':'Python', 'text': 'Python'},{'value':'Java', 'text': 'Java'}]
        },
        {
          title: '代码',
          name: 'code',
          value: {code: ''},
          default: false,
          itemType: 'bigInput',
          isEditOnly: false,
          others:["输入代码......"]
        },
        {
          title: '描述',
          name: 'description',
          value: {description: ''},
          default: false,
          itemType: 'bigInput',
          isEditOnly: false,
          others:["相关描述......"]
        },

      ],
      addFormItem: {
        released: "00",
        data_path: "",
      },
      // pageKind标明当前页的信息（database，modelbase等）
      // taskType表明nowItem父级名字
      // nowItem表明选中的是菜单的哪一项。
      // cardNameFlag用来标识作为卡片名称的属性
      cardNameFlag: "dataset_name",
      pageKind: 'defineFunction',
      taskType: null,
      nowItem: null,
      jsonBaseUrl: "http://localhost:3000",
      myCardList: [],
      publicCardList: [],
      myCardNum: 0,
      myCardRowNum: 0,
      myCardColNum: 6,
      publicCardNum: 0,
      publicCardRowNum: 0,
      publicCardColNum: 6,
    }
  },
  provide(){
    return {
      updataPage: this.updataPage,
      addFormItemCfg: this.addFormItemCfg,
      pageKind: this.pageKind,
      // taskType: this.taskType,
      // nowItem: this.nowItem,
      jsonBaseUrl: this.jsonBaseUrl,
      getPageContent: this.getPageContent,
      cardNameFlag: this.cardNameFlag,
      // myCardList: this.myCardList,
      // myCardNum: this.myCardNum,
      // myCardRowNum: this.myCardRowNum,
      // myCardColNum: this.myCardColNum,

      // publicCardList: this.publicCardList,
      // publicCardNum: this.publicCardNum,
      // publicCardRowNum: this.publicCardRowNum,
      // publicCardColNum: this.publicCardColNum,
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
    updataPage(actionType) {
      if(actionType == "delete") {
        // console.info("this is delete")
        this.$Message["success"]({
          background: true,
          content: "删除成功"
        });
      } else if(actionType == "upload") {
        // console.info("this is delete")
        this.$Message["success"]({
          background: true,
          content: "发布成功"
        });
      } else if(actionType == "creat") {
        this.$Message["success"]({
          background: true,
          content: "新建成功"
        });
      } else if(actionType == "edit") {
        this.$Message["success"]({
          background: true,
          content: "编辑成功"
        });
      }
      this.myCardList= []
      this.publicCardList=[]
      this.getPageContent()
    },
    getNowItem(name) {
      // taskType-nowItem
      var myName = name.split('-')
      this.nowItem = myName[1]
      this.taskType = myName[0]
      console.info(this.nowItem)
      console.info(this.taskType)

      this.getPageContent()
      this.getFormItem()
    },
    toPages(name) {
      var targetUrl = "/" + name
      this.$router.push(targetUrl)
    },
    getPageContent() {
      this.myCardList= []
      this.publicCardList=[]
      var findUrl = this.jsonBaseUrl + "/" + this.pageKind + "?task=" + this.nowItem + "&type=" + this.taskType
      console.info(findUrl)
      axios.get(findUrl).then(response => {
        var cardList = response.data
        console.info(cardList)
        var length = cardList.length

        for(var i = 0; i < length;i++) {
          if(cardList[i].released[0] == '1') {
            this.myCardList.push(cardList[i] );
          }
          if(cardList[i].released[1] == '1') {
            this.publicCardList.push(cardList[i]);
          }
          if(cardList[i].released == "00") {
            var findDeleteUrl = this.jsonBaseUrl + "/" + this.pageKind + "/" + cardList[i].id
            // console.info(findUrl)
            axios.delete(findDeleteUrl).then(response=>{
              console.info("delete success")
            })
          }
        }
        this.myCardNum = this.myCardList.length
        this.myCardRowNum = Math.ceil((this.myCardNum+1) / this.myCardColNum)
        this.publicCardNum = this.publicCardList.length
        this.publicCardRowNum =  Math.ceil(this.publicCardNum / this.publicCardColNum)
      })
    },
    getFormItem() {
      for(var index in this.addFormItemCfg) {
        // console.info(this.addFormItemCfg[index].name)
        this.addFormItem = {...this.addFormItem, ...this.addFormItemCfg[index].value}
        if(this.addFormItemCfg[index].default == true) {
          var name = this.addFormItemCfg[index].name
          // var len =
          for(var j in this.addFormItemCfg[index].others) {
            if(j != 0) {
              this.addFormItem[name] = this.addFormItem[name] + '-'
            }
            var text = this.addFormItemCfg[index].others[j]
            this.addFormItem[name] = this[text]
          }
        }
      }
    },
  }
}
</script>
