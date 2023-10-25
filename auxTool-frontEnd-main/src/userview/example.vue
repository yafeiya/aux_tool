<style>
#Charts{
  width: 500px;
  height:250px;
  border: 1px solid #729ce3;
  margin: auto;
  margin-top: 10px;
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
              <h4 v-if="row.State==='挂起中'">
                <Icon type="md-pause" />
                挂起中
              </h4>
              <h4 v-if="row.State==='运行中'">
                <Icon type="md-time" />
                运行中
              </h4>
              <h4 v-if="row.State==='已终止'">
                <Icon type="md-power" />
                已终止
              </h4>
            </template>
            <!--时间相关-->
            <template #time="{row}">
              <Time :time="row.Start_time- 60 * 1 * 1000"  />
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
                    <Col span="8">案例号：{{ curRow.Id }}</Col>
                    <Col span="8">案例名: {{ curRow.Example_name }}</Col>
                    <Col span="8">状态: {{ curRow.State }}</Col>
                    <Col span="8">级别: {{ curRow.Rank }}</Col>
                    <Col span="8">数据目录: {{ curRow.Dataset_url }}</Col>
                  </Row>

                </Card>
                <Card>
                  <template #title><strong>模型信息</strong></template>
                  <Row>
                    <Col span="8">模型名：{{ curRow.Model_name }}</Col>
                    <Col span="8">模型类型: {{ curRow.Model_type }}</Col>
                    <Col span="8">迭代次数: {{ curRow.Epoch_num }}</Col>
                    <Col span="8">损失函数: {{ curRow.Loss }}</Col>
                    <Col span="8">优化器: {{ curRow.Optimizer }}</Col>
                    <Col span="8">衰减因子: {{ curRow.Decay }}</Col>
                    <Col span="8">评价指标: {{ curRow.Evalution }}</Col>
                    <Col span="8">保存位置: {{ curRow.Model_url }}</Col>
                  </Row>
                </Card>
                <Card>
                  <template #title><strong>资源需求</strong></template>
                  <Row>
                    <Col span="8">CPU核数：{{ curRow.Cpu_num }}</Col>
                    <Col span="8">GPU算力: {{ curRow.Gpu_num }}</Col>
                    <Col span="8">内存用量: {{ curRow.Memory }}</Col>
                    <Col span="8">启动时间:<Time :time="curRow.Start_time" type="datetime" /></Col>
                    <Col span="8">结束时间:
                      <Time v-if="curRow.end_time!='' " :time="curRow.End_time" type="datetime" />
                      <div v-else> {{ "" }}</div>
                    </Col>
                    <Col span="8">运行时间:{{ curRow.Run_time }}</Col>
                  </Row>
                </Card>
                <template #footer>
                  <Button type="info"  long @click="close">确定</Button>
                </template>
              </Modal>


              <Button type="info" style="margin-right: 5px" @click="LogInfo(row)" v-width=85 >训练过程</Button>
<!--              <Button type="info" style="margin-right: 5px" @click="isLogInfo=true" v-width=85 >训练过程</Button>-->
              <Modal v-model="isLogInfo" width="1200" style="margin-top: -80px">
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

                <template #footer>
                  <Space :size="50">
                    <Button type="info"  long @click="previewImage=true" v-width="100" >方案预览</Button>
                    <Button type="info"  long @click="isproxyInfo=true" v-width="100" style="margin-right: 460px">指标介绍</Button>
                  </Space>
                </template>
              </Modal>
              <ImagePreview v-model="previewImage" :preview-list="urlList" />
              <Modal v-model="isproxyInfo" width="700" style="margin-top: -10px">
                <template #header>
                  <p style="color:#4d85ea;text-align:center">
                    <Icon type="ios-information-circle"></Icon>
                    <span>指标介绍</span>
                  </p>
                </template>
                <Card>
                  <template #title><strong>学习率 (learning-rate)</strong></template>
                  学习率决定了在模型训练过程中每一步参数更新的幅度，如果学习率设置得过小，模型训练会变得非常缓慢，甚至可能在达到最优值之前就停止更新了。相反，如果学习率设置得过大，可能会导致在损失函数空间中来回跳动，甚至可能使得损失不断增大，从而无法收敛到最优解
                </Card>
                <Card>
                  <template #title><strong>回报率 (reward)</strong></template>
                  在强化学习中，奖励（Reward）是一个用于评价智能体在环境中行动的反馈信号。奖励通常是一个标量值，用来表示某个特定时刻智能体的行动的好坏程度。奖励的目的是引导智能体学习如何在环境中选择行动以达到特定的目标。
                </Card>
                <Card>
                  <template #title><strong>网络损失 (loss)</strong></template>
                  损失函数是在机器学习和深度学习中用来度量模型预测输出与实际标签之间差异的一种函数，损失函数接受模型的预测输出和真实标签作为输入，并计算一个数值来表示它们之间的差异。该数值越小，表示模型的预测越接近实际标签，也就是模型的性能越好。
                </Card>

                <template #footer>
                  <Button type="info"  long @click="isproxyInfo=false" >确定</Button>
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
import {getCsvData, getExampleList, updateExample, deleteExample,getprocessFile} from "../api/api.js"

import datas from "@/userview/flow/results/data.json";

export default {

  data() {
    return {
      previewImage:false,
      urlList: ["",],
      isproxyInfo:false,
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
          key: 'Id',
          width: 100,
          sortable: true,
          align: 'center'
        },
        {
          title: '案例名',
          key: 'Example_name',
          align: 'center'
        },
        {
          title: '级别',
          key: 'Rank',
          sortable: true,
          align: 'center'
        },
        {
          title: '状态',
          key: 'State',
          slot: 'zhuangtai',
          align: 'center'
        },
        {
          title: 'CPU核数',
          key: 'Cpu_num',
          align: 'center'
        },
        {
          title: 'GPU个数',
          key: 'Gpu_num',
          align: 'center'
        },
        {
          title: '提交时间',
          key: 'Start_time',
          width: 180,
          slot: 'time',
          align: 'center',
          sortable: true,
        },
        {
          title: '运行时长',
          key: 'Run_time',
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


      // var data={
      //   Id :
      //   Type :
      //   Task :
      //   processFile :
      // }
      getprocessFile
      let actdata = chartxyz["actions"]
      let rewarddata = chartxyz["reward"]
      let lrdata = chartxyz["learning_rate"]
      let count=[]  //统计轮数
      const x = 4;  //飞机种类数
      const  y = 7; //导弹种类数
      let maxnum = 0
      let MSLnum = new Array(x); // MSL记录数据
      let MSLdata=[]

      for (var i = 0; i < x; i++) {
        MSLnum[i] = new Array(y); // 创建二维数组
      }
      // console.log(MSLnum)
      //初始为0
      for (var q = 0; q < x; q++) {
        for (var p = 0; p < y; p++) {
          MSLnum[q][p]=0
        }
      }
      // console.log(MSLnum)
      //统计发射数目
      var turn = 0
      let msi=0
      while(lrdata[turn]){
        count.push(turn)
        for (var j=0;j<x;j++){
          msi = actdata[turn][j][1]-1;//序号-1
          MSLnum[j][msi]++
        }
        turn++
      }
      //转换成最终坐标
      for (var n = 0; n < x; n++) {
        for (var m = 0; m < y; m++) {
          if(MSLnum[n][m]>maxnum){maxnum = MSLnum[n][m]}
          MSLdata.push([n,m,MSLnum[n][m]])
        }
      }

      var option1 = {
        title: {
                text: '火力分配',
                x: 'left'
            },
        xAxis3D: {
          type: 'category',
        },
        yAxis3D: {
          type: 'category',
        },
        zAxis3D: {
          type: 'value',
        },
        grid3D: {
          viewControl: {//可以控制整个柱状图场景旋转平移等，自行代数数据试试
            projection: 'orthographic'//正交投影
          }
        },
        visualMap: {
          calculable: true,
          max: maxnum,
          // dimension: 'Life Expectancy',
          inRange: {
            color: ['#313695', '#4575b4', '#74add1', '#abd9e9', '#e0f3f8', '#ffffbf', '#fee090', '#fdae61', '#f46d43', '#d73027', '#a50026']
          }
        },
        dataset: {
          source: MSLdata
        },
        series: [{
          type: 'bar3D',
          barSize: 10,//柱子大小
          encode: {
            // 维度的名字默认就是表头的属性名
            x: 'Airplane',
            y: 'Missile',
            z: 'Num',
            tooltip: [0, 1, 2, 3, 4]
          }
        }]
      };
      var option2={
        title: {
          text: '奖励曲线'
        },
        legend: {
          data: ['reward'],
          x:'right',      //可设定图例在左、右、居中
        },
        xAxis: {
          data: count
        },
        grid: {
          left: '5%',
          height: 'auto',
          right: '5%',
          bottom: '5%',
          containLabel: true
        },
        yAxis:{},
        series: [
          {
            name: "y轴",
            type: "line",
            data: rewarddata
          }
        ]
      };
      var option3={
        title: {
          text: '学习率曲线'
        },
        legend: {
          data: ['learning rate'],
          x:'right',      //可设定图例在左、右、居中
        },
        grid: {
          left: '5%',
          height: 'auto',
          right: '5%',
          bottom: '5%',
          containLabel: true
        },
        xAxis: {
          data: count
        },
        yAxis:{},
        series: [
          {
            name: "y轴",
            type: "line",
            data: lrdata
          }
        ]
      };
      let option4 = {
        title: {
          text: '准确率曲线'
        },
        legend: {
          data: ['acc'],
          x:'right',      //可设定图例在左、右、居中
        },
        tooltip: {},
        grid: {
          left: '5%',
          height: 'auto',
          right: '5%',
          bottom: '5%',
          containLabel: true
        },
        xAxis: {
          data: count
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
          // this.chartId = chartData["chart"][i]["exampleId"];
          // var chartxyz=chartData["chart"][this.chartId]
          console.info("row.id:",this.itemList[i])
          // this.getcharts(chartxyz)
        }

      }

    },
    // 弹窗关闭按钮
    close () {
      this.isItemInfo=false;
      this.isLogInfo=false;
      this.isDataInfo=false;
      this.isproxyInfo=false;
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

            updateExample(data).then(response => {


            })
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
      let selectId=""
      for(var i in this.selections){
        selectId+=this.selections[i].Id+"/"
      }
      let data = {
        pagekind:this.pageKind,
        id:selectId,
      }
      console.info("222222222222222222222",data)
      deleteExample(data).then(res => {
        console.info('删除csv结果：', res)
      })
      this.$Message.success('删除成功')
      setTimeout(() => {
        this.getItemInfo()
      },20)
    },
    selectChange(selection) {
      this.selections = selection;
      console.log('已选中数据', this.selections)
    },
    getItemInfo() {
      // var findUrl = this.jsonBaseUrl + "/" + this.pageKind
      this.itemList = []
      let data={
        pageKind:this.pageKind
      }
      getExampleList(data).then(response => {
        this.itemList = response.data.data.examples
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
      // axios.get(findUrl).then(response => {
      //
      //
      // })
    },
    updatePage(page) {
      this.curItemList = []
      var len = Math.min(this.itemList.length - 1, this.pageSize * page - 1)
      for(var i = 0 + (page-1) * this.pageSize;i <= len;i++) {
        this.curItemList.push(this.itemList[i])
      }
    },
  },
  mounted(){

  },
}

</script>
