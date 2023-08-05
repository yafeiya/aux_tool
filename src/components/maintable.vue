<template>
  <Tabs v-if="pageKind!=='defineFunction'" type="card" class="main-table">
    <TabPane :label = this.labelName1>
      <!--        {{ pageKind +" " + nowItem + " /my " + myCardList.length }}-->
      <menuCard :view-range="'private'"  :now-item="this.nowItem" :task-type="taskType"
                :my-card-list="myCardList" :public-card-list="publicCardList"
                :my-card-num="myCardNum" :my-card-row-num="myCardRowNum" :my-card-col-num="myCardColNum"
                :public-card-num="publicCardNum" :public-card-row-num="publicCardRowNum" :public-card-col-num="publicCardColNum"
                :add-form-item="addFormItem"/>
    </TabPane>
    <TabPane :label = this.labelName2>
      <!--        {{ pageKind +" " + nowItem + " /public" + publicCardList.length }}-->
      <menuCard :view-range="'public'"  :now-item="this.nowItem" :task-type="taskType"
                :my-card-list="myCardList" :public-card-list="publicCardList"
                :my-card-num="myCardNum" :my-card-row-num="myCardRowNum" :my-card-col-num="myCardColNum"
                :public-card-num="publicCardNum" :public-card-row-num="publicCardRowNum" :public-card-col-num="publicCardColNum"
                :add-form-item="addFormItem"/>
    </TabPane>
  </Tabs>
  <Tabs v-else type="card" class="main-table">
    <TabPane :label = this.labelName1>
      <!--        {{ pageKind +" " + nowItem + " /my" }}-->
      <menuCard :view-range="'private'"  :now-item="this.nowItem" :task-type="taskType"
                :my-card-list="myCardList" :public-card-list="publicCardList"
                :my-card-num="myCardNum" :my-card-row-num="myCardRowNum" :my-card-col-num="myCardColNum"
                :public-card-num="publicCardNum" :public-card-row-num="publicCardRowNum" :public-card-col-num="publicCardColNum"
                :add-form-item="addFormItem"/>
    </TabPane>
  </Tabs>
</template>
<script>
import menuCard from './menucard.vue';
import axios from 'axios';
export default {

  data() {
    return {
      // myUrl: null,
      // publicUrl: null
      labelName1: "我的",
      labelName2: "公共",
      // jsonBaseUrl: "http://localhost:3000",
    }
  },
  inject:['jsonBaseUrl', 'pageKind'],
  props: ['nowItem', 'taskType', 'myCardList', 'publicCardList', 'myCardNum', 'myCardRowNum', 'myCardColNum',
    'publicCardNum', 'publicCardRowNum', 'publicCardColNum', 'addFormItem'],

  components: {
    menuCard
  },
  mounted() {
    if(this.pageKind === "database") {
      this.labelName1 = "我的数据集"
      this.labelName2 = "公共数据集"
    } else if(this.pageKind === "modelbase") {
      this.labelName1 = "我的模型库"
      this.labelName2 = "基础算法模型"
    } else if(this.pageKind === "design") {
      this.labelName1 = "我的方案"
      this.labelName2 = "公共方案库"
    } else {
      this.labelName1 = "自定义函数"
    }
    // this.getPageContent()
  },
  methods: {

  }
}
</script>
