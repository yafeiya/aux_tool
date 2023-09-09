<style>
#Charts{
  width: 400px;
  height:200px;
  border: 1px solid red;
  margin: auto;
}
.layout{
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
}
.layout-home{
  width: 150px;
  height: 30px;
  border-radius: 3px;
  float: left;
  position: relative;
  top: 3px;
  left: 0px;
}
.layout-nav1{
  width: 620px;
  margin: 0 auto;
  margin-right: -80px;
}
.dev-run-preview .dev-run-preview-edit{ display: none }
</style>
<template>

  <!-- <Button @click="test">1231</Button> -->
  <div class="layout">

    <Layout :style="{marginLeft: '0px'}">
      <!--头部菜单导航-->
      <Header>
        <Menu mode="horizontal" theme="dark" active-name="example" @on-select="toPages">
          <div class="layout-home" v-color="'white'" style="font-size: 20px;margin-left: -15px;margin-top: -2px">
            <Button :type="text" size="small" shape="circle" @click="toHome">
              <Icon type="md-arrow-back" />
            </Button>
            |  返回首页
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
      <!--主体内容部分-->

      <Content :style="{padding: '0 16px 16px'}">


        <!--=批量处理按钮-->
        <p style="margin-top: 1%;font-size: 20px">
          批量处理：
          <Button  type="warning" icon="md-power" shape="circle" v-width=90 style="margin-left: 0%" @click="updateToState('终止')">终止</Button>
          <Button  type="success" icon="md-play"  shape="circle" v-width=90 style="margin-left: 1%" @click="updateToState('运行')">继续</Button>
          <Button  type="primary" icon="md-pause"  shape="circle" v-width=90 style="margin-left: 1%" @click="updateToState('挂起')">挂起</Button>
          <Button  type="error" icon="md-trash"  shape="circle" v-width=90 style="margin-left: 1%" @click="deleteItem">删除</Button>
        </p>

        <div style="margin-top: 1%">
          <!--表格部分-->

          <Table border ref="selection" :columns="columns" :data="curItemList" @on-selection-change="selectChange" style="width: auto">
            <!--根据状态state值显示图标-->

            <template #zhuangtai="{ row }">
              <h4 v-if="row.state==='挂起中'">
                <Icon type="md-pause" />
                挂起中
              </h4>
              <h4 v-if="row.state==='运行中'">
                <Icon type="md-time" />
                运行中
              </h4>
              <h4 v-if="row.state==='已终止'">
                <Icon type="md-power" />
                已终止
              </h4>
            </template>
            <!--时间相关-->
            <template #time="{row}">
              <Time :time="row.post_time" type="datetime" />
            </template>
            <!--表格最右列查看详情-->

            <template #details="{row, index}">

              <Button type="info" style="margin-right: 5px;margin-left: -10%" @click="itemInfoBtn(row, index)" v-width=85 >详情</Button>

              <Modal v-model="isItemInfo" width="800" :loading="true">
                <template #header>
                  <p style="color:#4d85ea;text-align:center">
                    <Icon type="ios-information-circle"></Icon>
                    <span>案例详情</span>
                  </p>
                </template>
                <Card>
                  <template #title><strong>任务信息</strong></template>

                  <Row>
                    <Col span="8">案例号：{{ curRow.id }}</Col>
                    <Col span="8">案例名: {{ curRow.example_name }}</Col>
                    <Col span="8">状态: {{ curRow.state }}</Col>
                    <Col span="8">级别: {{ curRow.rank }}</Col>
                    <Col span="8">数据目录: {{ curRow.dataset_url }}</Col>
                  </Row>

                </Card>
                <Card>
                  <template #title><strong>模型信息</strong></template>
                  <Row>
                    <Col span="8">模型名：{{ curRow.model_name }}</Col>
                    <Col span="8">模型类型: {{ curRow.model_type }}</Col>
                    <Col span="8">迭代次数: {{ curRow.epoch_num }}</Col>
                    <Col span="8">损失函数: {{ curRow.loss }}</Col>
                    <Col span="8">优化器: {{ curRow.optimizer }}</Col>
                    <Col span="8">衰减因子: {{ curRow.decay }}</Col>
                    <Col span="8">评价指标: {{ curRow.evalution }}</Col>
                    <Col span="8">保存位置: {{ curRow.model_url }}</Col>
                  </Row>
                </Card>
                <Card>
                  <template #title><strong>资源需求</strong></template>
                  <Row>
                    <Col span="8">CPU核数：{{ curRow.cpu_num }}</Col>
                    <Col span="8">GPU算力: {{ curRow.gpu_num }}</Col>
                    <Col span="8">内存用量: {{ curRow.memory }}</Col>
                    <Col span="8">启动时间:<Time :time="curRow.start_time" type="datetime" /></Col>
                    <Col span="8">结束时间:
                      <Time v-if="curRow.end_time!='' " :time="curRow.end_time" type="datetime" />
                      <div v-else> {{ "" }}</div>
                    </Col>
                    <Col span="8">运行时间:{{ curRow.run_time }}</Col>
                  </Row>
                </Card>
                <template #footer>
                  <Button type="info"  long @click="close">确定</Button>
                </template>
              </Modal>


              <Button type="info" style="margin-right: 5px" @click="LogInfo(row)" v-width=85 >训练过程</Button>
<!--              <Button type="info" style="margin-right: 5px" @click="isLogInfo=true" v-width=85 >训练过程</Button>-->
              <Modal v-model="isLogInfo" width="1000" style="margin-top: -50px">
                <template #header>
                  <p style="color:#4d85ea;text-align:center">
                    <Icon type="ios-information-circle"></Icon>
                    <span>训练过程可视化</span>
                  </p>
                </template>

                <Row>
                  <Col span="12">
                    <div id="Charts" ref="Echarts1" ></div>
                  </Col>
                  <Col span="12">
                    <div id="Charts" ref="Echarts2"></div>
                  </Col>
                </Row>
                <Row>
                  <Col span="12">
                    <div id="Charts" ref="Echarts3" ></div>
                  </Col>
                  <Col span="12">
                    <div id="Charts" ref="Echarts4"></div>
                  </Col>
                </Row>
<!--                <div id="Charts" ref="Echarts" ></div>-->


                <template #footer>
                  <Button type="info"  long @click="close">确定</Button>
                </template>
              </Modal>
              <Button type="info" style="margin-right: -30px" @click="isDataInfo=true" v-width=85 >模型评价</Button>
              <Modal v-model="isDataInfo" width="500" style="margin-top: 100px">
                <template #header>
                  <p style="color:#4d85ea;text-align:center">
                    <Icon type="ios-information-circle"></Icon>
                    <span>案例数据</span>
                  </p>
                </template>
                <p style="text-align:center">
                  <!--以上传组件代替实现打开本地资源管理器-->
                  <Upload action="">
                    <Button icon="ios-albums-outline />" >查看本地数据</Button>
                  </Upload>
                </p>
                <template #footer>
                  <Button type="info"  long @click="close" >确定</Button>
                </template>
              </Modal>

            </template>
          </Table>
        </div>

      </Content>
      <!--分页组件-->
      <!-- <h3>{{ this.itemNum }}</h3> -->
      <Page :total="itemNum" :page-size="pageSize" style="margin-bottom: 1%;text-align: center;" @on-change="updatePage"/>
    </Layout>
  </div>
</template>
<script>
import {MenuGroup, Result} from "view-ui-plus";
import lineChart from '../components/chart/line.vue'
import axios from 'axios';
import * as echarts from 'echarts'
import chartData from "./chartdata.json"

export default {

  data() {
    return {
      nowTime: (new Date()).getTime(),
      runTime: 0,
      // time2: (new Date()).getTime() - 86400 * 3 * 1000 + 2000,
      itemNum: 0,
      pageSize: 7,
      selections: [],
      // 当前显示的item集合（分页后）
      curItemList: [],
      isItemInfo: false,
      isLogInfo: false,
      isDataInfo: false,
      // 当前行
      curRow: {},
      chartId: "",
      // chartxyz:{},
// 表头
      columns: [
        {
          type: 'selection',
          width: 60,
          align: 'center'
        },
        {
          title: '案例号',
          key: 'id',
          width: 100,
          sortable: true,
          align: 'center'
        },
        {
          title: '案例名',
          key: 'example_name',
          align: 'center'
        },
        {
          title: '级别',
          key: 'rank',
          sortable: true,
          align: 'center'
        },
        {
          title: '状态',
          key: 'state',
          slot: 'zhuangtai',
          align: 'center'
        },
        {
          title: 'CPU核数',
          key: 'cpu_num',
          align: 'center'
        },
        {
          title: 'GPU个数',
          key: 'gpu_num',
          align: 'center'
        },
        {
          title: '提交时间',
          key: 'submit_time',
          width: 180,
          slot: 'time',
          align: 'center',
          sortable: true,
        },
        {
          title: '运行时长',
          key: 'run_time',
          align: 'center',
          sortable: true,
        },
        {
          title: '查看',
          slot: 'details',
          width: 280,
          align: 'center'
        },
        // {
        //   title: '实例信息',
        //   slot: 'exampleInfo',
        //   width: 280,
        //   align: 'center'
        // }
      ],
// 表内数据
      itemList: [],
      data: [

      ],
      jsonBaseUrl: "http://localhost:3000",
      // pageKind标明当前页的信息（database，modelbase等）
      pageKind: 'example',
    }
  },
  components:{
    lineChart,
  },
  inject:['reload'],

  created() {
    this.getItemInfo();

  },
  methods: {
    getcharts(chartxyz){
      let psgTimeCharts1 = echarts.init(this.$refs.Echarts1)
      let psgTimeCharts2 = echarts.init(this.$refs.Echarts2)
      let psgTimeCharts3 = echarts.init(this.$refs.Echarts3)
      let psgTimeCharts4 = echarts.init(this.$refs.Echarts4)
      let option1 = {
      title: {
        text: '数据可视化'
      },
      tooltip: {},
      grid: {
        left: '3%',
        height: 'auto',
        right: '4%',
        bottom: '3%',
        containLabel: true
      },
      xAxis: {
        data: chartxyz["reward"],
      },
      yAxis: {
      },
      series: [
        {
          name: 'y1',
          symbol: 'none', //去掉折线上面的小圆点
          data: chartxyz["reward"],
          type: 'line',
          // areaStyle: {}
        },
      ],
    };
      let option2 = {
        title: {
          text: '数据可视化'
        },
        tooltip: {},
        grid: {
          left: '3%',
          height: 'auto',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        xAxis: {
          data: chartxyz["learning_rate"],
        },
        yAxis: {
        },
        series: [
          {
            name: 'y1',
            symbol: 'none', //去掉折线上面的小圆点
            data: chartxyz["learning_rate"],
            type: 'line',
            // areaStyle: {}
          },
        ],
      };
      let option3 = {
        title: {
          text: '数据可视化'
        },
        tooltip: {},
        grid: {
          left: '3%',
          height: 'auto',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        xAxis: {
          data: chartxyz["loss"],
        },
        yAxis: {
        },
        series: [
          {
            name: 'y1',
            symbol: 'none', //去掉折线上面的小圆点
            data: chartxyz["loss"],
            type: 'line',
            // areaStyle: {}
          },
        ],
      };
      let option4 = {
        title: {
          text: '数据可视化'
        },
        tooltip: {},
        grid: {
          left: '3%',
          height: 'auto',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        xAxis: {
          data: chartxyz["Accuracy"],
        },
        yAxis: {
        },
        series: [
          {
            name: 'y1',
            symbol: 'none', //去掉折线上面的小圆点
            data: chartxyz["Accuracy"],
            type: 'line',
            // areaStyle: {}
          },
        ],
      };
      psgTimeCharts1.setOption(option1)
      psgTimeCharts2.setOption(option2)
      psgTimeCharts3.setOption(option3)
      psgTimeCharts4.setOption(option4)
    console.info("1111111111",option1)
  },

    calcTime(newTime, oldTime) {
      var timeDiff = newTime - oldTime
      var timeDiff = Math.floor(timeDiff / 1000)
      var min = Math.floor(timeDiff / 60)
      var hour = Math.floor(min / 60)
      min = min % 60
      var ans = 0;
      if(hour == 0) {
        ans = min + "分钟"
      } else {
        ans = hour + "小时" + min + "分钟"
      }
      return ans
    },
    itemInfoBtn(row, index) {
      this.isItemInfo=true;
      for(var i in this.itemList) {
        if(row.id == this.itemList[i].id) {
          this.curRow = this.itemList[i];
        }
      }

    },
    LogInfo(row) {
      this.isLogInfo=true;

      for(var i in this.itemList) {
        if(row.id == this.itemList[i].id) {
          this.chartId = chartData["chart"][i]["exampleId"];
          var chartxyz=chartData["chart"][this.chartId]
          console.info("this.chartId:",this.chartId)
          console.info("chartxyz:",chartxyz)
          this.getcharts(chartxyz)
        }

      }

    },
    // 弹窗关闭按钮
    close () {
      this.isItemInfo=false;
      this.isLogInfo=false;
      this.isDataInfo=false;
    },
    toHome() {
      this.$router.push('/home')
    },

    toPages(name) {
      var targetUrl = "/" + name
      this.$router.push(targetUrl)
    },
    putItemState(toState) {
      var findUrl = this.jsonBaseUrl + '/' + this.pageKind
      var putList = []
      for(var i in this.selections) {
        for (var j in this.itemList) {
          //
          if(this.selections[i].id == this.itemList[j].id) {
            if(toState == "挂起"){
              this.itemList[j].state = "挂起中"
              this.itemList[j].end_time = this.nowTime
              this.itemList[j].run_time = this.calcTime(this.itemList[i].end_time, this.itemList[i].start_time)
            } else if(toState == "运行") {
              this.itemList[j].state = "运行中"
              this.itemList[j].end_time = ""
              this.itemList[j].run_time = this.calcTime(this.nowTime, this.itemList[i].start_time)
            } else if(toState == "终止") {
              this.itemList[j].state = "已终止"
              this.itemList[j].end_time = this.nowTime
              this.itemList[j].run_time = this.calcTime(this.itemList[i].end_time, this.itemList[i].start_time)
            }

            var putUrl = findUrl + "/" + this.itemList[j].id
            var putPromise = new Promise((resolve,reject)=>{
              axios.put(putUrl, this.itemList[j]).then(response=>{
                // console.info("in:" + response)
              })
            })
            putList.push(putPromise)

          }
        }
      }
      Promise.all(putList).then((result) =>{
        console.info("result: " + result)
        this.getItemInfo()
      }).catch((error) => {
        console.info(error)
      })
    },
    // tostate param：终止，运行，挂起
    // state:已终止 挂起中 运行中
    updateToState(toState) {
      if(this.selections.length == 0) {
        this.$Message["error"]({
          background: true,
          content: "当前未选中任何对象，请选择需要" + toState + "的对象"
        });
      } else {
        if(toState == "运行" || toState == "挂起") {
          for(var i in this.selections) {
            if(this.selections[i].state == "已终止") {
              this.$Message["error"]({
                background: true,
                content: "选中目标中存在终止对象，请修改后操作"
              });
              return
            }
          }
        }


        this.putItemState(toState)
      }

      // this.reload()
    },

    deleteItem() {
      if(this.selections.length == 0) {
        this.$Message["error"]({
          background: true,
          content: "当前未选中任何对象，请选择需要删除的对象"
        });
        return;
      }
      var findUrl = this.jsonBaseUrl + '/' + this.pageKind
      var deleteList = []
      for(var i in this.selections) {
        for (var j in this.itemList) {
          //
          if(this.selections[i].id == this.itemList[j].id) {

            var deleteUrl = findUrl + "/" + this.itemList[j].id
            var deletePromise = new Promise((resolve,reject)=>{
              axios.delete(deleteUrl, this.itemList[j]).then(response=>{
                console.info("in:" + response)
                this.itemList.splice(j, 1)
                this.getItemInfo()
              })
            })
            deleteList.push(deletePromise)

          }
        }
      }
      Promise.all(deleteList).then((result) =>{
        console.info("result: " + result)

      }).catch((error) => {
        console.info(error)
      })
      // this.reload()
    },
    selectChange(selection) {
      this.selections = selection;
      console.log('已选中数据', this.selections)
    },
    getItemInfo() {
      var findUrl = this.jsonBaseUrl + "/" + this.pageKind
      this.itemList = []
      axios.get(findUrl).then(response => {
        this.itemList = response.data
        console.info(this.itemList)
        this.itemNum = this.itemList.length
        for(var i in this.itemList) {
          if(this.itemList[i].state == "运行中") {
            this.itemList[i].run_time = this.calcTime(this.nowTime, this.itemList[i].start_time)
            console.info(this.itemList[i].run_time)
            console.info(this.nowTime)
            console.info(this.itemList[i].start_time)
            this.itemList[i].end_time = ""
          } else if(this.itemList[i].state == "挂起中" || this.itemList[i].state == "已终止") {
            this.itemList[i].run_time = this.calcTime(this.itemList[i].end_time, this.itemList[i].start_time)
          }
        }
        this.updatePage(1)

      })
    },
    updatePage(page) {
      console.info(this.itemNum)
      this.curItemList = []
      var len = Math.min(this.itemList.length - 1, this.pageSize * page - 1)
      for(var i = 0 + (page-1) * this.pageSize;i <= len;i++) {
        this.curItemList.push(this.itemList[i])
      }
      console.info(this.curItemList)
    },
  },
  mounted(){

  },
}

</script>
