<template>
  <div v-if="viewRange === 'private'">
    <div v-if="myCardNum!=0">
      <Row v-for="(i,index1) in this.myCardRowNum" :key="index1"  style="background:#eee;padding:20px" :gutter="16" >
        <Col v-for="(j, index2) in this.getColOfRow(i, this.myCardRowNum, this.myCardColNum, this.myCardNum)"
             :key="index2" :span= "Math.ceil(24 / this.myCardColNum)">
          <div v-if="i==1&&j==1">
            <addCard :view-range="viewRange"  :now-item="nowItem" :task-type="taskType"
                     :add-form-item="addFormItem"/>
          </div>
          <!-- 新增智能建模模块卡片 -->
          <div v-else-if="i==1&&j==2&&this.pageKind=='design'">
            <aicard />
          </div>
          
          <div v-else>
            <contentCard :card-info="this.myCardList[ (i-1) * this.myCardColNum + j - 2 ]" :view-range="viewRange" 
                         :now-item="nowItem" :task-type="taskType" :add-form-item="addFormItem" :card-list="myCardList"/>
                         <!-- :card-name="this.myCardList[ (i-1) * this.myCardColNum + j - 2 ][cardNameFlag]" -->
          </div>
        </Col>
        <!-- 独立addcard始终显示在卡片区最后一格 -->
       <!-- <Col :span="Math.ceil(24 / this.myCardColNum)"
          <addCard :view-range="viewRange" :now-item="nowItem" :task-type="taskType" :add-form-item="addFormItem"/>
        </Col> -->
      </Row>
    </div>
    <div v-else>
      <Row style="background:#eee;padding:20px" :gutter="16">
        <Col :span= "Math.ceil(24 / this.myCardColNum)">
          <addCard :view-range="viewRange"  :now-item="nowItem" :task-type="taskType"
                   :add-form-item="addFormItem"/>
        </Col>
      </Row>
    </div>
  </div>
  <div v-else>
    <Row v-for="(i,index1) in this.publicCardRowNum" :key="index1"  style="background:#eee;padding:20px" :gutter="16" >
      <Col v-for="(j, index2) in this.getColOfRow(i, this.publicCardRowNum, this.publicCardColNum, this.publicCardNum)"
           :key="index2" :span= "Math.ceil(24 / this.publicCardColNum)">
        <contentCard :card-info="this.publicCardList[ (i-1) * this.publicCardColNum + j - 1]" :view-range="viewRange" 
                     :now-item="nowItem" :task-type="taskType" :add-form-item="addFormItem" :card-list="publicCardList"/>
                     <!-- :card-name="this.publicCardList[ (i-1) * this.publicCardColNum + j - 1][cardNameFlag]" -->
      </Col>
    </Row>
  </div>
</template>
<script>
import axios from 'axios';
import addCard from './addcard.vue'
import contentCard from './contentcard.vue'
import aicard from './aicard.vue'
export default {
  data() {
    return {
      // cardName: this.,
      // jsonBaseUrl: "http://localhost:3000",
      // 中间页面卡片列表与数量
      // myCardList: [],
      // publicCardList: [],
      // colnum应该为24的因数
      // myCardNum: 0,
      // myCardRowNum: 0,
      // myCardColNum: 4,
      // publicCardNum: 0,
      // publicCardRowNum: 0,
      // publicCardColNum: 4,
    }
  },
  inject:['jsonBaseUrl', 'pageKind','cardNameFlag'],
  // nowItem左侧选中条目，pageKind当前页面种类{'database'，'modelbase'，'definFunc'，'design'}，viewRange公共的还是私有的{'private','public'}
  props: ['nowItem','viewRange','taskType' ,'myCardList', 'publicCardList', 'myCardNum', 'myCardRowNum', 'myCardColNum',
    'publicCardNum', 'publicCardRowNum', 'publicCardColNum', 'addFormItem'],
  components:{
    addCard,
    contentCard,
    aicard
  },
  mounted() {
    console.info(this.myCardList)
  },  
  methods: {

    // 一列的个数不一定等于列数，因为有可能不够，所以得计算
    // (当前所处行得位置，总行数，总列数，总个数)
    getColOfRow(nowRow, rowNum, colNum, Num) {
      if(nowRow == rowNum) {
        if(this.viewRange == "private") {
          return (Num - ((rowNum-1) * colNum) + 1)
        } else {
          return (Num - ((rowNum-1) * colNum))
        }

      } else {
        return colNum
      }
    }
  },

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
.dev-run-preview .dev-run-preview-edit{ display: none }
.main-table {
  margin-top: 2%;
}
</style>
