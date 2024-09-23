<style>
#Charts {
  width: 500px;
  height: 350px;
  border: 1px solid #729ce3;
  margin: auto;
  margin-top: 10px;
}
.layout {
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
}
.layout-home {
  width: 150px;
  height: 30px;
  border-radius: 3px;
  float: left;
  position: relative;
  top: 3px;
  left: 0px;
}
.layout-nav1 {
  width: 620px;
  margin: 0 auto;
  margin-right: -80px;
}
.dev-run-preview .dev-run-preview-edit {
  display: none;
}
</style>
<template>
  <!-- <Button @click="test">1231</Button> -->
  <div class="layout">
    <Layout :style="{ marginLeft: '0px' }">
      <!--头部菜单导航-->
      <Header>
        <Menu
          mode="horizontal"
          theme="dark"
          active-name="example"
          @on-select="toPages"
        >
          <div
            class="layout-home"
            v-color="'white'"
            style="font-size: 20px; margin-left: -15px; margin-top: -2px"
          >
            <Button :type="text" size="small" shape="circle" @click="toHome">
              <Icon type="md-arrow-back" />
            </Button>
            | 返回首页
          </div>
          <div class="layout-nav1">
            <MenuItem name="database">
              <Icon type="ios-navigate"></Icon>
              数据集
            </MenuItem>
            <MenuItem name="modelbase">
              <Icon type="ios-keypad"></Icon>
              算法库
            </MenuItem>
            <MenuItem name="defineFunction">
              <Icon type="ios-analytics"></Icon>
              自定义函数
            </MenuItem>
            <MenuItem name="design">
              <Icon type="ios-paper"></Icon>
              模型设计
            </MenuItem>
            <MenuItem name="example">
              <Icon type="md-arrow-dropright-circle"></Icon>
              实例管理
            </MenuItem>
          </div>
        </Menu>
      </Header>
      <!--主体内容部分-->

      <Content :style="{ padding: '0 16px 16px' }">
        <!--=批量处理按钮-->
        <p style="margin-top: 1%; font-size: 20px">
          批量处理：
          <Button
            type="warning"
            icon="md-power"
            shape="circle"
            v-width="90"
            style="margin-left: 0%"
            @click="updateToState('已终止')"
            >终止</Button
          >
          <Button
            type="success"
            icon="md-play"
            shape="circle"
            v-width="90"
            style="margin-left: 1%"
            @click="updateToState('运行中')"
            >继续</Button
          >
          <Button
            type="primary"
            icon="md-pause"
            shape="circle"
            v-width="90"
            style="margin-left: 1%"
            @click="updateToState('挂起中')"
            >挂起</Button
          >
          <Button
            type="error"
            icon="md-trash"
            shape="circle"
            v-width="90"
            style="margin-left: 1%"
            @click="deleteItem"
            >删除</Button
          >
        </p>

        <div style="margin-top: 1%">
          <!--表格部分-->

          <Table
            border
            ref="selection"
            :columns="columns"
            :data="curItemList"
            @on-selection-change="selectChange"
            style="width: auto"
          >
            <!--根据状态state值显示图标-->

            <template #zhuangtai="{ row }">
              <h4 v-if="row.State === '挂起中'">
                <Icon type="md-pause" />
                挂起中
              </h4>
              <h4 v-if="row.State === '运行中'">
                <Icon type="md-time" />
                运行中
              </h4>
              <h4 v-if="row.State === '已终止'">
                <Icon type="md-power" />
                已终止
              </h4>
            </template>
            <!--时间相关-->
            <template #time="{ row }">
              <Time :time="row.Start_time - 60 * 1 * 1000" type="datetime" />
            </template>
            <!--表格最右列查看详情-->

            <template #details="{ row, index }">
              <Row>
                <Col span="6">
                  <Upload
                    :action="EndUrl().backEndUrl + '/uploadResult'"
                    method="POST"
                    accept=".json"
                    :data="lossData"
                    :show-upload-list="false"
                    :on-success="uploadSuccess"
                  >
                    <Button
                      type="success"
                      style=""
                      @click="uploadfile(row)"
                      v-width="85"
                      >结果上传</Button
                    >
                  </Upload>
                </Col>
                <Col span="6">
                  <Button
                    type="info"
                    style=""
                    @click="itemInfoBtn(row, index)"
                    v-width="85"
                    >详情</Button
                  >

                  <Modal v-model="isItemInfo" width="800" :loading="true">
                    <template #header>
                      <p style="color: #4d85ea; text-align: center">
                        <Icon type="ios-information-circle"></Icon>
                        <span>案例详情</span>
                      </p>
                    </template>
                    <Card>
                      <template #title><strong>实例信息</strong></template>

                      <Row>
                        <Col span="8">案例号：{{ curRow.Id }}</Col>
                        <Col span="8">案例名: {{ curRow.Example_name }}</Col>
                        <Col span="8">类型: {{ curRow.Type }}</Col>
                        <Col span="8">状态: {{ curRow.State }}</Col>
                        <Col span="8">级别: {{ curRow.Rank }}</Col>
                        <Col span="8">任务: {{ curRow.Task }}</Col>
                      </Row>
                    </Card>
                    <Card>
                      <template #title><strong>建模方案</strong></template>
                      <Row>
                        <Col span="8">数据集：{{ curRow.Dataset_name }}</Col>
                        <Col span="8">模型算法: {{ curRow.Model_name }}</Col>
                        <Col span="8"
                          >训练过程分布: {{ curRow.Train_state }}</Col
                        >
                      </Row>
                    </Card>
                    <Card>
                      <template #title><strong>模型超参数</strong></template>
                      <Row>
                        <Col span="8">网络层数: {{ curRow.Network_num }}</Col>
                        <Col span="8">学习率: {{ curRow.Learning_rate }}</Col>
                        <Col span="8">优化器：{{ curRow.Optimizer }}</Col>
                        <Col span="8">迭代次数: {{ curRow.Epoch_num }}</Col>
                        <Col span="8">批大小: {{ curRow.Batch_size }}</Col>
                        <Col span="8">激活函数: {{ curRow.Act_function }}</Col>
                        <Col span="8">衰减因子: {{ curRow.Decay_factor }}</Col>
                        <Col span="8">探索率: {{ curRow.Explore_rate }}</Col>
                        <Col span="8">随机种子: {{ curRow.Radom_seed }}</Col>
                      </Row>
                    </Card>
                    <template #footer>
                      <Button type="info" long @click="close">确定</Button>
                    </template>
                  </Modal>
                </Col>
                <Col span="6">
                  <Button
                    type="info"
                    style=""
                    @click="LogInfo(row)"
                    v-width="85"
                    >训练过程</Button
                  >
                  <!--              <Button type="info" style="margin-right: 5px" @click="isLogInfo=true" v-width=85 >训练过程</Button>-->
                  <Modal
                    v-model="isLogInfo"
                    width="1200"
                    style="margin-top: -80px"
                  >
                    <template #header>
                      <p style="color: #4d85ea; text-align: center">
                        <Icon type="ios-information-circle"></Icon>
                        <span>训练过程可视化</span>
                      </p>
                    </template>

                    <Row>
                      <Col span="12">
                        <div id="Charts" ref="Echarts1"></div>
                      </Col>
                      <Col span="12">
                        <div id="Charts" ref="Echarts2"></div>
                      </Col>
                    </Row>
                    <Row>
                      <Col span="12">
                        <div id="Charts" ref="Echarts3"></div>
                      </Col>
                      <Col span="12">
                        <div
                          id="Charts"
                          style="display: flex; justify-content: center"
                        >
                          <img
                            style="height: 100%"
                            :src="imageUrl"
                            alt="Example Image"
                          />
                        </div>
                      </Col>
                    </Row>

                    <template #footer>
                      <Space :size="50">
                        <Button
                          type="info"
                          long
                          @click="previewImage = true"
                          v-width="100"
                          >方案预览</Button
                        >
                        <Button
                          type="info"
                          long
                          @click="isproxyInfo = true"
                          v-width="100"
                          style="margin-right: 460px"
                          >指标介绍</Button
                        >
                      </Space>
                    </template>
                  </Modal>
                  <ImagePreview
                    v-model="previewImage"
                    :preview-list="urlList"
                  />
                  <Modal
                    v-model="isproxyInfo"
                    width="700"
                    style="margin-top: -10px"
                  >
                    <template #header>
                      <p style="color: #4d85ea; text-align: center">
                        <Icon type="ios-information-circle"></Icon>
                        <span>指标介绍</span>
                      </p>
                    </template>
                    <Card>
                      <template #title
                        ><strong>学习率 (learning-rate)</strong></template
                      >
                      学习率决定了在模型训练过程中每一步参数更新的幅度，如果学习率设置得过小，模型训练会变得非常缓慢，甚至可能在达到最优值之前就停止更新了。相反，如果学习率设置得过大，可能会导致在损失函数空间中来回跳动，甚至可能使得损失不断增大，从而无法收敛到最优解
                    </Card>
                    <Card>
                      <template #title
                        ><strong>回报率 (reward)</strong></template
                      >
                      在强化学习中，奖励（Reward）是一个用于评价智能体在环境中行动的反馈信号。奖励通常是一个标量值，用来表示某个特定时刻智能体的行动的好坏程度。奖励的目的是引导智能体学习如何在环境中选择行动以达到特定的目标。
                    </Card>
                    <Card>
                      <template #title
                        ><strong>网络损失 (loss)</strong></template
                      >
                      损失函数是在机器学习和深度学习中用来度量模型预测输出与实际标签之间差异的一种函数，损失函数接受模型的预测输出和真实标签作为输入，并计算一个数值来表示它们之间的差异。该数值越小，表示模型的预测越接近实际标签，也就是模型的性能越好。
                    </Card>

                    <template #footer>
                      <Button type="info" long @click="isproxyInfo = false"
                        >确定</Button
                      >
                    </template>
                  </Modal>
                </Col>
                <Col span="6">
                  <Button
                    type="info"
                    style=""
                    @click="ModelVal(row)"
                    v-width="85"
                    >模型评价</Button
                  >
                  <Modal
                    v-model="isModelInfo"
                    width="1200"
                    style="margin-top: -80px"
                    @on-ok="closeLogInfo()"
                    @on-cancel="closeLogInfo()"
                  >
                    <template #header>
                      <p style="color: #4d85ea; text-align: center">
                        <Icon type="ios-information-circle"></Icon>
                        <span>模型评价</span>
                      </p>
                    </template>

                    <Row>
                      <Col span="12">
                        <div id="Charts" ref="Echarts4"></div>
                      </Col>
                      <Col span="12">
                        <div id="Charts" ref="Echarts5"></div>
                      </Col>
                    </Row>
                    <Row v-if="!isBigdata">
                      <Col span="12">
                        <div id="Charts" ref="Echarts6"></div>
                      </Col>
                      <Col span="12">
                        <div id="Charts" ref="Echarts7"></div>
                      </Col>
                    </Row>
                  </Modal>
                </Col>
              </Row>
            </template>
          </Table>
        </div>
      </Content>
      <!--分页组件-->
      <!-- <h3>{{ this.itemNum }}</h3> -->
      <Page
        :total="itemNum"
        :page-size="pageSize"
        style="margin-bottom: 1%; text-align: center"
        @on-change="updatePage"
      />
    </Layout>
  </div>
</template>
<script>
import { MenuGroup, Result } from "view-ui-plus";
import lineChart from "../components/chart/line.vue";
import axios from "axios";
import * as echarts from "echarts";
import {
  getCsvData,
  getExampleList,
  updateExample,
  deleteExample,
  getprocessFile,
  getResultFile,
} from "../api/api.js";

import datas from "@/userview/flow/results/data.json";
import { EndUrl } from "../../url_config";

export default {
  data() {
    return {
      lossData: {
        id: "",
        type: "",
        task: "",
      },
      isModelInfo: false,
      actdata: "",
      resultdata: "",
      rewarddata: "",
      lrdata: "",
      imageUrl: EndUrl().fileUrl,
      modelimageUrl1: "",
      modelimageUrl2: "",
      modelimageUrl3: "",
      modelimageUrl4: "",
      previewImage: false,
      urlList: [""],
      isproxyInfo: false,
      nowTime: new Date().getTime(),
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
      isBigdata: false,
      // 当前行
      curRow: {},
      chartId: "",
      // chartxyz:{},
      // 表头
      columns: [
        {
          type: "selection",
          width: 60,
          align: "center",
        },
        {
          title: "案例号",
          key: "Id",
          width: 100,
          sortable: true,
          align: "center",
        },
        {
          title: "案例名",
          key: "Example_name",
          align: "center",
        },
        {
          title: "级别",
          key: "Rank",
          sortable: true,
          align: "center",
        },
        {
          title: "状态",
          key: "State",
          slot: "zhuangtai",
          align: "center",
        },
        {
          title: "数据集",
          key: "Dataset_name",
          align: "center",
        },
        {
          title: "模型算法",
          key: "Model_name",
          align: "center",
        },
        {
          title: "提交时间",
          key: "Start_time",
          width: 180,
          slot: "time",
          align: "center",
          sortable: true,
        },
        {
          title: "运行时长",
          key: "run_time",
          align: "center",
          sortable: true,
        },
        {
          title: "查看",
          slot: "details",
          width: 400,
          align: "center",
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
      data: [],
      jsonBaseUrl: "http://localhost:3000",
      // pageKind标明当前页的信息（database，modelbase等）
      pageKind: "example",
    };
  },
  components: {
    lineChart,
  },
  inject: ["reload"],

  created() {
    this.getItemInfo();
  },
  methods: {
    EndUrl,
    uploadfile(row) {
      (this.lossData.id = row.Example_id),
        (this.lossData.task = row.Task),
        (this.lossData.type = row.Type);
      console.info("iiiiiiiiiiiiiiiiiiiiii", this.lossData);
    },
    uploadSuccess() {
      this.$Message.success("上传成功");
    },
    getcharts() {
      console.info("data1:", this.actdata);
      console.info("data2:", this.rewarddata);
      console.info("data3:", this.lrdata);
      let psgTimeCharts1 = echarts.init(this.$refs.Echarts1);
      let psgTimeCharts2 = echarts.init(this.$refs.Echarts2);
      let psgTimeCharts3 = echarts.init(this.$refs.Echarts3);
      // let psgTimeCharts4 = echarts.init(this.$refs.Echarts4)
      let count = []; //统计轮数
      let x = 0; //飞机种类数
      let y = 0; //导弹种类数
      let maxnum = 0;
      let MSLnum = []; // MSL记录数据
      let MSLdata = []; //最终画actions图的数据表
      let rewarddata = []; //最终画reward的数据表
      let lossdict = {}; //最终画loss的数据表

      for (var i = 0; i < 100; i++) {
        MSLnum.push(new Array(50)); // 创建二维数组
      }
      // console.log("创建二维数组",MSLnum)
      // 初始为0
      for (var q = 0; q < 100; q++) {
        for (var p = 0; p < 50; p++) {
          MSLnum[q][p] = 0;
        }
      }
      // console.log(MSLnum)
      //统计发射数目
      let msi = [];
      let serial = 12;
      var arr = [",", "[", "]"]; //排除非数字

      while (this.actdata[serial] != null) {
        if (arr.indexOf(this.actdata[serial]) == -1) {
          // console.log(actionDataString[serial])
          //处理数字
          msi.push(Number(this.actdata[serial]));
          // console.log(msi)
          if (msi[1] != null) {
            // console.log("处理msi",msi)
            // console.log(actionDataString[serial])
            if (x < msi[0]) {
              x = msi[0];
              // console.log("显示最大x",x)
            }
            if (y < msi[1]) {
              y = msi[1];
              // console.log("显示最大y",x)
            }
            // console.log("MSLnum[msi[0],msi[1]]",MSLnum[msi[0],msi[1]])
            MSLnum[msi[0]][msi[1]]++;
            // console.log("MSLnum[msi[0],msi[1]]",MSLnum[msi[0],msi[1]])
            msi = [];
          }
        }
        serial++;
      }
      // console.log("处理的数组",MSLnum)
      //转换成最终坐标
      for (var n = 0; n < x; n++) {
        for (var m = 0; m < y; m++) {
          if (MSLnum[n][m] > maxnum) {
            maxnum = MSLnum[n][m];
          }
          MSLdata.push([n, m, MSLnum[n][m]]);
        }
      }
      console.log("MSLdata:", MSLdata);

      let serial2 = 11;
      let rewad = "";
      while (this.rewarddata[serial2] != null) {
        if (this.rewarddata[serial2] != ",") {
          if (this.rewarddata[serial2] == "]") {
            rewarddata.push(Number(rewad));
            break;
          }
          rewad = rewad + this.rewarddata[serial2];
          // console.log(rewad)
        } else if (this.rewarddata[serial2] == ",") {
          rewarddata.push(Number(rewad));
          rewad = "";
          count++; // 记录轮数
        }
        serial2++;
      }
      console.log("rewad is", rewarddata);

      //处理loss
      let lossname = "";
      let lossnamedict = [];
      let serial3 = 9;
      let loss = "";

      while (this.lrdata[serial3] != null) {
        if (this.lrdata[serial3] == '"') {
          serial3++;
          while (this.lrdata[serial3] != '"' && this.lrdata[serial3] != null) {
            lossname = lossname + this.lrdata[serial3];
            // console.log(lossDataString[serial3])
            serial3++;
          }
          serial3 = serial3 + 2;
          lossdict[lossname] = new Array();
          // console.log("outwhile2",lossdict)
          lossnamedict.push(lossname);
        } else if (this.lrdata[serial3] != ",") {
          if (this.lrdata[serial3] == ":") {
            serial3 = serial3 + 2;
            // console.log("inpan：",lossDataString[serial3])
            loss = loss += this.lrdata[serial3];
          } else if (this.lrdata[serial3] == "]") {
            lossdict[lossname].push(Number(loss));
            loss = "";
            lossname = "";
            serial3++;
          } else {
            loss = loss += this.lrdata[serial3];
          }
        } else if (this.lrdata[serial3] == ",") {
          // console.log("loss:",loss)
          lossdict[lossname].push(Number(loss));
          loss = "";
        }
        serial3++;
      }
      console.log("lossdict", lossdict);
      // console.log("lossdata['alpha_loss']:",lossdict['alpha_loss'])

      var option1 = {
        title: {
          text: "火力分配",
          x: "left",
        },
        xAxis3D: {
          type: "category",
        },
        yAxis3D: {
          type: "category",
        },
        zAxis3D: {
          type: "value",
        },
        grid3D: {
          viewControl: {
            //可以控制整个柱状图场景旋转平移等，自行代数数据试试
            projection: "orthographic", //正交投影
          },
        },
        visualMap: {
          calculable: true,
          max: maxnum,
          // dimension: 'Life Expectancy',
          inRange: {
            color: [
              "#313695",
              "#4575b4",
              "#74add1",
              "#abd9e9",
              "#e0f3f8",
              "#ffffbf",
              "#fee090",
              "#fdae61",
              "#f46d43",
              "#d73027",
              "#a50026",
            ],
          },
        },
        dataset: {
          source: MSLdata,
        },
        series: [
          {
            type: "bar3D",
            barSize: 10, //柱子大小
            encode: {
              // 维度的名字默认就是表头的属性名
              x: "Airplane",
              y: "Missile",
              z: "Num",
              tooltip: [0, 1, 2, 3, 4],
            },
          },
        ],
      };
      var option2 = {
        title: {
          text: "奖励曲线",
          x: "left",
        },
        xAxis: {
          data: count,
        },
        grid: {
          left: "5%",
          height: "auto",
          right: "5%",
          bottom: "5%",
          containLabel: true,
        },
        yAxis: {},
        series: [
          {
            name: "y轴",
            type: "line",
            data: rewarddata,
          },
        ],
      };
      var option3 = {
        title: {
          text: "损失函数",
          x: "left",
        },
        xAxis: {
          data: count,
        },
        grid: {
          left: "5%",
          height: "auto",
          right: "5%",
          bottom: "5%",
          containLabel: true,
        },
        yAxis: {},
        series: [],
      };
      var lossnamenum = 0;
      while (lossnamedict[lossnamenum] != null) {
        let a = {
          type: "line",
          data: lossdict[lossnamedict[lossnamenum]],
        };
        option3.series.push(a);
        lossnamenum++;
      }

      psgTimeCharts1.setOption(option1);
      psgTimeCharts2.setOption(option2);
      psgTimeCharts3.setOption(option3);
    },
    getMeticscharts() {
      if (this.isBigdata == false) {
        let psgTimeCharts4 = echarts.init(this.$refs.Echarts4);
        let psgTimeCharts5 = echarts.init(this.$refs.Echarts5);
        let psgTimeCharts6 = echarts.init(this.$refs.Echarts6);
        let psgTimeCharts7 = echarts.init(this.$refs.Echarts7);
        var option4 = {
          tooltip: {
            position: "top",
          },
          title: {
            text: "混淆矩阵",
            x: "left",
          },
          grid: {
            left: "10%",
            right: "4%",
            bottom: "3%",
            containLabel: true,
          },
          xAxis: {
            type: "category",
            data: ["攻击", "突防", "预警", "侦察", "佯攻", "撤退", "干扰"],
            splitLine: { show: false },
            axisLabel: {
              fontSize: 10, // 调整字体大小
              rotate: 45, // 旋转标签
            },
          },
          yAxis: {
            type: "category",
            data: ["攻击", "突防", "预警", "侦察", "佯攻", "撤退", "干扰"],
            splitLine: { show: false },
          },
          visualMap: {
            min: 0,
            max: 0.25,
            calculable: true,
            inRange: {
              color: ["#ffffff", "#565656"], // 从白色到红色
            },
            text: ["高", "低"],
            textStyle: {
              color: "#fff",
            },
          },
          series: [
            {
              name: "热力图",
              type: "heatmap",
              data: this.resultdata.con_mat,
              label: {
                show: true,
              },
              emphasis: {
                itemStyle: {
                  shadowBlur: 10,
                  shadowColor: "#333",
                },
              },
            },
          ],
        };
        var option5 = {
          tooltip: {},
          title: {
            text: "准确率曲线",
            x: "left",
          },

          xAxis: {
            type: "category",
            axisLabel: {
              interval: 4,
              rotate: 0, // 旋转标签
            },
          },
          graphic: [
            {
              type: "text",
              left: "90%", // 右侧位置
              top: "80%", // 垂直居中
              style: {
                text: "epochs",
                font: "12px Microsoft YaHei", // 字体样式
                fill: "#333", // 字体颜色
                textAlign: "center",
              },
            },
          ],
          yAxis: {
            type: "value",
          },
          series: [
            {
              name: "y轴",
              type: "line",
              data: this.resultdata.acc,
            },
          ],
        };
        var option6 = {
          tooltip: {
            position: "top",
          },
          title: {
            text: "意图研判",
            x: "left",
          },
          grid: {
            left: "10%",
            right: "4%",
            bottom: "3%",
            containLabel: true,
          },
          xAxis: {
            type: "category",
            data: [
              "距离",
              "速度",
              "RCS",
              "高度",
              "雷达状态",
              "方位角",
              "机动类型",
            ],
            splitLine: { show: false },
            axisLabel: {
              fontSize: 10, // 调整字体大小
              rotate: 45, // 旋转标签
            },
          },
          yAxis: {
            type: "category",
            data: ["攻击", "突防", "预警", "侦察", "佯攻", "撤退", "干扰"],
            splitLine: { show: false },
          },
          visualMap: {
            min: 0,
            max: 0.25,
            calculable: true,
            inRange: {
              color: ["#ffffff", "#ff0000"], // 从白色到红色
            },
            text: ["高", "低"],
            textStyle: {
              color: "#fff",
            },
          },
          series: [
            {
              name: "热力图",
              type: "heatmap",
              data: this.resultdata.hot,
              label: {
                show: true,
              },
              emphasis: {
                itemStyle: {
                  shadowBlur: 10,
                  shadowColor: "#333",
                },
              },
            },
          ],
        };
        var option7 = {
          title: {
            text: "ROC&AUC曲线",
            left: "left",
          },
          tooltip: {
            trigger: "item",
          },
          xAxis: {
            type: "value",
            name: "假正例率",
            nameLocation: "middle", // 将标签放在中间
            min: 0,
            max: 1,
          },
          yAxis: {
            type: "value",
            name: "真正例率",
            nameLocation: "middle", // 将标签放在中间
            nameGap: 30, // 标签与轴的间距
            min: 0,
            max: 1,
          },
          series: [
            {
              name: "ROC Curve",
              type: "line",
              data: this.resultdata.fpr.map((x, i) => [
                x,
                this.resultdata.tpr[i],
              ]),
              smooth: true,
              lineStyle: {
                color: "blue",
                width: 2,
              },
              symbol: "none",
            },
            {
              name: "Random Guess",
              type: "line",
              data: [
                [0, 0],
                [1, 1],
              ],
              lineStyle: {
                color: "red",
                type: "dashed",
              },
            },
          ],
          graphic: {
            elements: [
              {
                type: "text",
                left: "70%", // 右侧位置
                top: "80%", // 垂直居中
                style: {
                  text: `AUC: 0.82`,
                  fill: "blue",
                  font: "16px Microsoft YaHei",
                },
              },
            ],
          },
          grid: {
            left: "10%",
            right: "10%",
            top: "10%",
            bottom: "10%",
            containLabel: true,
            // 设置宽高比
            aspectRatio: 1,
          },
        };
        psgTimeCharts4.setOption(option4);
        psgTimeCharts5.setOption(option5);
        psgTimeCharts6.setOption(option6);
        psgTimeCharts7.setOption(option7);
      } else {
        let psgTimeCharts4 = echarts.init(this.$refs.Echarts4);
        let psgTimeCharts5 = echarts.init(this.$refs.Echarts5);
        var option4 = {
          tooltip: {},
          title: {
            text: "成功率曲线",
            x: "left",
          },
          xAxis: {
            type: "value", // 将类型改为 'value'
            min: 0, // 设置最小值为 0
            max: 9999, // 设置最大值为 9999
            interval: 1000, // 设置刻度间隔
            axisLabel: {
              formatter: function (value) {
                return value; // 格式化显示
              },
            },
          },
          graphic: [
            {
              type: "text",
              left: "90%", // 右侧位置
              top: "85%", // 垂直居中
              style: {
                text: "epochs",
                font: "12px Microsoft YaHei", // 字体样式
                fill: "#333", // 字体颜色
                textAlign: "center",
              },
            },
          ],
          yAxis: {
            type: "value",
          },
          series: [
            {
              name: "y轴",
              type: "line",
              smooth: true,
              lineStyle: {
                width: 2,
              },
              symbol: "none",
              data: this.resultdata.succ.map((value, index) => [
                index * (9999 / (this.resultdata.succ.length - 1)),
                value,
              ]), // 将数据映射到 0 到 9999
            },
          ],
        };
        var option5 = {
          tooltip: {},
          title: {
            text: "奖励曲线",
            x: "left",
          },
          xAxis: {
            type: "value", // 将类型改为 'value'
            min: 0, // 设置最小值为 0
            max: 9999, // 设置最大值为 9999
            interval: 1000, // 设置刻度间隔
            axisLabel: {
              formatter: function (value) {
                return value; // 格式化显示
              },
            },
          },
          graphic: [
            {
              type: "text",
              left: "90%", // 右侧位置
              top: "85%", // 垂直居中
              style: {
                text: "epochs",
                font: "12px Microsoft YaHei", // 字体样式
                fill: "#333", // 字体颜色
                textAlign: "center",
              },
            },
          ],
          yAxis: {
            type: "value",
          },
          series: [
            {
              name: "y轴",
              type: "line",
              smooth: true,
              lineStyle: {
                width: 2,
              },
              symbol: "none",
              data: this.resultdata.reward.map((value, index) => [
                index * (9999 / (this.resultdata.reward.length - 1)),
                value,
              ]), // 将数据映射到 0 到 9999
            },
          ],
        };
        psgTimeCharts4.setOption(option4);
        psgTimeCharts5.setOption(option5);
      }
    },
    calcTime(startTime, endTime, state) {
      if (state === "on") {
        const currentTime = Date.now(); // 获取当前时间戳
        const runTime = currentTime - parseInt(startTime); // 计算运行时长，确保开始时间戳是整数
        // 将运行时长转换为天、小时、分钟和秒的格式
        // const seconds = Math.floor((runTime / 1000) % 60);
        const minutes = Math.floor((runTime / (1000 * 60)) % 60);
        const hours = Math.floor((runTime / (1000 * 60 * 60)) % 24);
        const days = Math.floor(runTime / (1000 * 60 * 60 * 24));
        return `${days}天${hours}小时${minutes}分钟`;
      } else {
        const runTime = endTime - startTime; // 计算运行时长，确保开始时间戳是整数
        // 将运行时长转换为天、小时、分钟和秒的格式
        // const seconds = Math.floor((runTime / 1000) % 60);
        const minutes = Math.floor((runTime / (1000 * 60)) % 60);
        const hours = Math.floor((runTime / (1000 * 60 * 60)) % 24);
        const days = Math.floor(runTime / (1000 * 60 * 60 * 24));
        return `${days}天${hours}小时${minutes}分钟`;
      }
    },
    itemInfoBtn(row, index) {
      this.isItemInfo = true;
      console.info("itemList", this.itemList);
      for (var i in this.itemList) {
        if (row.id == this.itemList[i].id) {
          this.curRow = this.itemList[i];
        }
      }
    },
    LogInfo(row) {
      this.isLogInfo = true;
      console.info("this.itemList", this.itemList);
      for (var i in this.itemList) {
        if (row.Id == this.itemList[i].Id) {
          this.imageUrl +=
            "/" +
            this.itemList[i].Type +
            "/" +
            this.itemList[i].Task +
            "/" +
            this.itemList[i].Example_id +
            "/image.png";
          console.info("this.imageUrl", this.imageUrl);
          let data1 = {
            type: this.itemList[i].Type,
            task: this.itemList[i].Task,
            id: this.itemList[i].Example_id,
            processFile: "actions.json",
          };
          console.info("data1", data1);
          getprocessFile(data1).then((res) => {
            this.actdata = res.data.data.Info;
            console.info("this.actdata", this.actdata);
          });

          let data2 = {
            type: this.itemList[i].Type,
            task: this.itemList[i].Task,
            id: this.itemList[i].Example_id,
            processFile: "reward.txt",
          };
          getprocessFile(data2).then((res) => {
            this.rewarddata = res.data.data.Info;
          });

          let data3 = {
            type: this.itemList[i].Type,
            task: this.itemList[i].Task,
            id: this.itemList[i].Example_id,
            processFile: "loss.csv",
          };
          getprocessFile(data3).then((res) => {
            this.lrdata = res.data.data.Info;
          });
          setTimeout(() => {
            this.getcharts();
          }, 100);
        }
      }
    },
    ModelVal(row) {
      this.isModelInfo = true;
      console.info("row", row);
      let data = {
        type: row.Type,
        task: row.Task,
        id: row.Example_id,
        processFile: "result.json",
      };
      if (row.Rank === "1级") {
        this.isBigdata = true;
      } else {
        this.isBigdata = false;
      }
      console.info("data", data);
      getResultFile(data).then((res) => {
        if (res.data === "") {
          this.resultdata = "";
        } else {
          this.resultdata = res.data.data.Info;
          console.info("this.resultdata", this.resultdata);
        }
      });

      setTimeout(() => {
        this.getMeticscharts();
      }, 100);
    },
    // 弹窗关闭按钮
    close() {
      this.isItemInfo = false;
      this.isLogInfo = false;
      this.isDataInfo = false;
      this.isproxyInfo = false;
    },
    toHome() {
      this.$router.push("/home");
    },

    toPages(name) {
      var targetUrl = "/" + name;
      this.$router.push(targetUrl);
    },
    putItemState(toState) {
      var putList = [];
      let selectId = "";
      for (var i in this.selections) {
        selectId += this.selections[i].Id + "/";
      }
      let data = {
        id: selectId,
        state: toState,
      };
      console.info("11111111111111111", data);
      updateExample(data).then((response) => {
        console.info("更新成功");
      });
      setTimeout(() => {
        this.getItemInfo();
      }, 20);
      this.$Message.success("修改成功");
      for (var i in this.selections) {
        for (var j in this.itemList) {
          if (this.selections[i].id == this.itemList[j].id) {
            if (toState == "挂起中") {
              this.itemList[j].state = "挂起中";
              this.itemList[j].end_time = this.nowTime;
              // this.itemList[j].run_time = this.calcTime(this.itemList[i].end_time, this.itemList[i].start_time)
            } else if (toState == "运行中") {
              this.itemList[j].state = "运行中";
              this.itemList[j].end_time = "";
              // this.itemList[j].run_time = this.calcTime(this.nowTime, this.itemList[i].start_time)
            } else if (toState == "已终止") {
              this.itemList[j].state = "已终止";
              this.itemList[j].end_time = this.nowTime;
              this.itemList[j].run_time = this.calcTime(
                this.itemList[i].end_time,
                this.itemList[i].start_time
              );
            }
            // var putUrl = findUrl + "/" + this.itemList[j].id
            // putList.push(putPromise)
          }
        }
      }
      console.info("itemList", this.itemList);

      Promise.all(putList)
        .then((result) => {
          console.info("result: " + result);
          this.getItemInfo();
        })
        .catch((error) => {
          console.info(error);
        });
    },
    // tostate param：终止，运行，挂起
    // state:已终止 挂起中 运行中
    updateToState(toState) {
      if (this.selections.length == 0) {
        this.$Message["error"]({
          background: true,
          content: "当前未选中任何对象，请选择需要" + toState + "的对象",
        });
      } else {
        if (toState == "运行中" || toState == "挂起中") {
          for (var i in this.selections) {
            if (this.selections[i].State == "已终止") {
              console.info(
                "this.selections[i].State",
                this.selections[i].State
              );
              this.$Message["error"]({
                background: true,
                content: "选中目标中存在终止对象，请修改后操作",
              });
              return;
            }
          }
        }
        this.putItemState(toState);
      }

      // this.reload()
    },

    deleteItem() {
      if (this.selections.length == 0) {
        this.$Message["error"]({
          background: true,
          content: "当前未选中任何对象，请选择需要删除的对象",
        });
        return;
      }
      let selectId = "";
      for (var i in this.selections) {
        selectId += this.selections[i].Id + "/";
      }
      let data = {
        pagekind: this.pageKind,
        id: selectId,
      };
      console.info("222222222222222222222", data);
      deleteExample(data).then((res) => {
        console.info("删除csv结果：", res);
      });
      this.$Message.success("删除成功");
      setTimeout(() => {
        this.getItemInfo();
      }, 20);
    },
    selectChange(selection) {
      this.selections = selection;
      console.log("已选中数据", this.selections);
    },
    getItemInfo() {
      // var findUrl = this.jsonBaseUrl + "/" + this.pageKind
      this.itemList = [];
      let data = {
        pageKind: this.pageKind,
      };
      getExampleList(data).then((response) => {
        this.itemList = response.data.data.examples;
        console.info("this.itemlist", this.itemList);
        this.itemNum = this.itemList.length;
        for (var i in this.itemList) {
          if (this.itemList[i].State === "已终止") {
            this.itemList[i].run_time = this.calcTime(
              this.itemList[i].Start_time,
              this.itemList[i].End_time,
              "off"
            );
          } else {
            this.itemList[i].run_time = this.calcTime(
              this.itemList[i].Start_time,
              this.itemList[i].End_time,
              "on"
            );
          }
        }
        this.updatePage(1);
      });
    },
    closeLogInfo() {
      console.info("clear");
      let chartRefs = ["Echarts4", "Echarts5", "Echarts6", "Echarts7"];
      chartRefs.forEach((ref) => {
        if (this.$refs[ref]) {
          let chartInstance = echarts.init(this.$refs[ref]);
          chartInstance.clear();
        }
      });
    },
    updatePage(page) {
      this.curItemList = [];
      var len = Math.min(this.itemList.length - 1, this.pageSize * page - 1);
      for (var i = 0 + (page - 1) * this.pageSize; i <= len; i++) {
        this.curItemList.push(this.itemList[i]);
      }
    },
  },
  mounted() {},
};
</script>
