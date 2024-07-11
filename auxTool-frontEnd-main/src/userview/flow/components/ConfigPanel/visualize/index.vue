<template>
  <Tabs class="config" value="1">
    <TabPane label="属性" name="1">
      <Row align="middle">
        <Button class="view" @click="modal = true">预览{{data.nodeselflabel}}大图</Button>
        <Modal
          v-model="modal"
          width="1000"
          draggable="true"
          class="chartmodal"
          title="训练图像"
          cancel-text=""
          @on-ok="ok"
          >
          <div id="myChart" :style="{width: '1000px', height: '500px'}"></div>

      </Modal>
      <Button class="view" @click="sendMoneytoYe">缩略图显示</Button>
      </Row>
    </TabPane>

    <TabPane label="节点" name="2">
      <Row align="middle">
        <Col :span="8">边框颜色</Col>
        <Col :span="14">
          <Input type="color" v-model="globalGridAttr.nodeStroke" style="width: 100%" @change="onStrokeChange" />
        </Col>
      </Row>
      <Row align="middle">
        <Col :span="8">边框宽度</Col>
        <Col :span="12">
          <Slider :min="1" :max="5" :step="1" v-model="data.nodeStrokeWidth" @on-input="onStrokeWidthChange" />
        </Col>
        <Col :span="2">
          <div class="result">{{ globalGridAttr.nodeStrokeWidth }}</div>
        </Col>
      </Row>
      <Row align="middle">
        <Col :span="8">填充颜色</Col>
        <Col :span="14">
          <Input type="color" v-model="data.nodeFill" style="width: 100%" @change="onFillChange" />
        </Col>
      </Row>
    </TabPane>
    <TabPane label="文本" name="3">
      <Row align="middle">
        <Col :span="8">字体大小</Col>
        <Col :span="12">
          <Slider v-model="data.nodeFontSize" :min="8" :max="16" :step="1"  @on-input="onFontSizeChange" />
        </Col>
        <Col :span="2">
          <div class="result">{{ globalGridAttr.nodeFontSize }}</div>
        </Col>
      </Row>
      <Row align="middle">
        <Col :span="8">字体颜色</Col>
        <Col :span="14">
          <Input type="color" value="globalGridAttr.nodeColor" style="width: 100%" @change="onColorChange" />
        </Col>
      </Row>
    </TabPane>
  </Tabs>
</template>

<script lang="ts">
  import { defineComponent, inject, watch, reactive,ref,provide} from 'vue';
  import { Cell } from '@antv/x6/lib';
  import { nodeOpt } from './method';
  import * as echarts from 'echarts'
  import 'echarts-gl'
  import { Button } from 'view-ui-plus';
  import datas from "../../../results/data.json"
  import {downloadModelFile, getprocessFile} from '../../../../../api/api.js'
  import qs from "qs";
  import datapre from '../datapre/index.vue';

  export default defineComponent({
    name: 'Index',

    data () {
            return {
                modal: false,
              }
            },
    methods: {
            ok () {
                this.$Message.info('预览完毕');
            },
        },

    async setup() {
      const globalGridAttr: any = inject('globalGridAttr');
      const id: any = inject('id');
      let curCel: Cell;
      let readpath :any =inject('datapath')

      // 向画布页面（祖父）传递预览图参数
      const changepreview: any = inject('changepreview')
      const preview: any = inject('grandpreview')
      const sendMoneytoYe = function() {
        preview.value = !preview.value
        changepreview(preview.value)
        // console.log(readpath.value)
      }

      const data = reactive({
        nodedownsamplescale: '',
        nodeaugmentationscale: '',
        nodeStrokeWidth: '',
        nodeFill: '',
        nodeFontSize: '',
        nodeselflabel:'',
        nodedataurl:'',
      })
      watch(
        [() => id.value],
        () => {
          curCel = nodeOpt(id, globalGridAttr);
          data.nodedownsamplescale = globalGridAttr.nodedownsamplescale;
          data.nodeaugmentationscale = globalGridAttr.nodeaugmentationscale;
          data.nodeStrokeWidth = globalGridAttr.nodeStrokeWidth
          data.nodeFill = globalGridAttr.nodeFill
          data.nodeFontSize = globalGridAttr.nodeFontSize
          data.nodeselflabel = globalGridAttr.selflabel
          data.nodedataurl = globalGridAttr.dataurl



          // 判断图像路径
          if(readpath.value == true){
            // console.log("需要bashinfo")
            bashinfo()
            // console.log("需要initchart")
          }else {
            clearHander()
          }
        },
        {
          immediate: false,
          deep: false,
        },
      );

      let databaseinfo:any = ref([])
      let databaseinfo2:any = ref([])
      let databaseinfo3:any = ref([])
      var actionDataString:any = ''
      var rewardDataString:any = ''
      var lossDataString:any = ''
      let tip = {
                id:'',
                type: '',
                task: '',
                processFile:'actions.json'
              }
      let tip2 = {
                id:'',
                type: '',
                task: '',
                processFile:'reward.txt'
              }
      let tip3 = {
                id:'',
                type: '',
                task: '',
                processFile:'loss.csv'
              }
      var url = decodeURI(window.location.href);
      // console.log(url)
      var cs_arr = url.split('?')[1];//?后面的
      tip.id = cs_arr.split('=')[1].split('&')[0];
      tip.task = cs_arr.split('=')[2].split('&')[0];
      tip.type = cs_arr.split('=')[3];
      tip2.id = tip.id
      tip2.task = tip.task
      tip2.type = tip.type
      tip3.id = tip.id
      tip3.task = tip.task
      tip3.type = tip.type
      qs.stringify(tip)
      // console.info("after——stringify",tip)

      databaseinfo = await getprocessFile(tip)
      databaseinfo2 = await getprocessFile(tip2)
      databaseinfo3 = await getprocessFile(tip3)
      // console.log("databaseinfo is", await getprocessFile(tip))
          //后端返回response.data.data.Info即为[[1,1][2,2]]
      if(databaseinfo.data.data!=null){
        actionDataString= await databaseinfo.data.data.Info
        // console.log("actionDataString:",await databaseinfo.data.data.Info)
      }
      if(databaseinfo2.data.data!=null){
        rewardDataString= await databaseinfo2.data.data.Info
      }
      if(databaseinfo3.data.data!=null){
        lossDataString= await databaseinfo3.data.data.Info
      }

      // console.info("actionData is", await databaseinfo.data.data.Info)
      // console.log("actdata is", await actionDataString)
      // console.log("rewarddata is",await rewardDataString)
      // console.log("lrdata is",await lossDataString)


      const initchart = (label) => {
        // console.log("执行initchart")
        // console.log("initchart",readpath.value)
        clearHander()
        // let actdata = datas["actions"]
        // let rewarddata = datas["reward"]
        // let lrdata = datas["learning_rate"]
        let count:any = []  //统计轮数

        let x = 0;  //飞机种类数
        let y = 0; //导弹种类数
        let maxnum = 0
        let MSLnum:any = []; // MSL记录数据
        let MSLdata:any = []   //最终画actions图的数据表
        let rewarddata:any =[] //最终画reward的数据表
        let lossdict:any = {}  //最终画loss的数据表


        //处理actions
        for (var i = 0; i < 100; i++) {
          MSLnum.push(new Array(50)) // 创建二维数组
        }
        // console.log("创建二维数组",MSLnum)
        // 初始为0
        for (var q = 0; q < 100; q++) {
          for (var p = 0; p < 50; p++) {
            MSLnum[q][p]=0
          }
        }
        // console.log("初始为0",MSLnum)
        // 统计发射数目
        let msi:any = []
        let serial = 12
        var arr:any = [',','[',']'] //排除非数字
        while(actionDataString[serial] != null){
          if(arr.indexOf(actionDataString[serial])==-1){
            // console.log(actionDataString[serial])
            //处理数字
            msi.push(Number(actionDataString[serial]))
            // console.log(msi)
            if(msi[1]!=null){
              // console.log("处理msi",msi)
              // console.log(actionDataString[serial])
              if(x<msi[0]){
                x=msi[0]
                // console.log("显示最大x",x)
              }
              if(y<msi[1]){
                y=msi[1]
                // console.log("显示最大y",x)
              }
              // console.log("MSLnum[msi[0],msi[1]]",MSLnum[msi[0],msi[1]])
              MSLnum[msi[0]][msi[1]]++
              // console.log("MSLnum[msi[0],msi[1]]",MSLnum[msi[0],msi[1]])
              msi = []
            }
          }
          serial++
        }
        // console.log("处理的数组",MSLnum)
        //转换成最终坐标
        for (var n = 0; n < x; n++) {
          for (var m = 0; m < y; m++) {
            if(MSLnum[n][m]>maxnum){maxnum = MSLnum[n][m]}
            MSLdata.push([n,m,MSLnum[n][m]])
          }
        }
      // console.log(count)
      // console.log("MSLdata:",MSLdata)


        //处理reward
        let serial2 = 11
        let rewad:any = ''
        while(rewardDataString[serial2] != null){
          if(rewardDataString[serial2] != ','){
            if(rewardDataString[serial2] == ']'){
              rewarddata.push(Number(rewad))
              break
            }
            rewad= rewad + rewardDataString[serial2]
            // console.log(rewad)
          }
          else if(rewardDataString[serial2] == ','){
            rewarddata.push(Number(rewad))
            rewad = ''
            count++ // 记录轮数
          }

          serial2++
        }
        // console.log("rewad is",rewarddata)


        //处理loss
        let lossname = ''
        let lossnamedict:any = []
        let serial3 = 9
        let loss = ''

        while(lossDataString[serial3] != null){
          if(lossDataString[serial3] == '"'){
            serial3++
            while(lossDataString[serial3] != '"'&&lossDataString[serial3] != null){
              lossname = lossname + lossDataString[serial3]
              // console.log(lossDataString[serial3])
              serial3++
            }
            serial3=serial3+2
            lossdict[lossname] = new Array()
            // console.log("outwhile2",lossdict)
            lossnamedict.push(lossname)
          }
          else if(lossDataString[serial3] != ','){
            if(lossDataString[serial3] == ':'){
              serial3 = serial3+2
              // console.log("inpan：",lossDataString[serial3])
              loss = loss+=lossDataString[serial3]
            }
            else if(lossDataString[serial3] == ']'){
              lossdict[lossname].push(Number(loss))
              loss = ''
              lossname = ''
              serial3++
            }
            else{
              loss = loss+=lossDataString[serial3]
            }
          }
          else if(lossDataString[serial3] == ','){
            // console.log("loss:",loss)
            lossdict[lossname].push(Number(loss))
            loss = ''
          }
          serial3++
        }
        // console.log("lossdict",lossdict)
        // console.log("lossdata['alpha_loss']:",lossdict['alpha_loss'])


        let myChart = echarts.init(document.getElementById("myChart"));
        let miniChart = echarts.init(document.getElementById("miniChart"));
        // 绘制图表
        const nodelabel = label


        //定义act图表参数
        var option1 = {
            // title: {
            //         text: '3D堆叠柱状图',
            //         x: 'center'
            //     },
              xAxis3D: {
                  type: 'category',
                  // axisLine:{
                  //     lineStyle:{
                  //         color:'black',
                  //         width:2
                  //     }
                  // },
              },
              yAxis3D: {
                  type: 'category',
                  // axisLine:{
                  //     lineStyle:{
                  //       width:1,
                  //       color:'black',
                  //     }
                  // },
              },
              zAxis3D: {
                  type: 'value',
                  // axisLine:{
                  //     lineStyle:{
                  //         color:'black',
                  //         width:1
                  //     }
                  // },
              },
              grid3D: {
                  viewControl: {//可以控制整个柱状图场景旋转平移等，自行代数数据试试
                      // alpha: 0,
                      // beta: 0,
                      // minAlpha: 0,//x轴旋转
                      // maxAlpha: 0,
                      // minBeta: 0,//y轴旋转
                      // maxBeta: 0,
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

        //定义reward图表参数
        var option2={
            xAxis: {
              data: count
            },
            yAxis:{},
            series: [
              {
                name: "y轴",
                type: "line",
                data: rewarddata
              }
            ]
          }


        //定义loss图表参数
        var option3 :any ={
            xAxis: {
              data: count
            },
            yAxis:{},
            series: [

            ]
          }

        var lossnamenum = 0
        while(lossnamedict[lossnamenum]!=null){
          let a={
              type: "line",
              data: lossdict[lossnamedict[lossnamenum]]
          }
          option3.series.push(a)
          lossnamenum++
        }

        // console.log(option3.series)

        if( nodelabel == '动作分布'){
          myChart.setOption(option1);
          miniChart.setOption(option1);
          // console.log("执行动作分布")
        }

        else if( nodelabel == '奖励分布'){
          myChart.setOption(option2);
          miniChart.setOption(option2);
          // console.log("执行奖励分布")
        }

        else if( nodelabel == '学习率'){
          myChart.setOption(option3);
          miniChart.setOption(option3);
          // console.log("执行学习率")
        }

        window.onresize = function () { // 自适应大小
          myChart.resize();
          miniChart.resize();
        };

      }

      async function bashinfo() {
        // console.log("执行bashinfo")
        databaseinfo = await getprocessFile(tip)
        databaseinfo2 = await getprocessFile(tip2)
        databaseinfo3 = await getprocessFile(tip3)
        // console.log("databaseinfo is", await getprocessFile(tip))
            //后端返回response.data.data.Info即为[[1,1][2,2]]
        if(databaseinfo.data.data!=null){
          actionDataString= await databaseinfo.data.data.Info
          // console.log("actionDataString:",await databaseinfo.data.data.Info)
        }
        if(databaseinfo2.data.data!=null){
          rewardDataString= await databaseinfo2.data.data.Info
        }
        if(databaseinfo3.data.data!=null){
          lossDataString= await databaseinfo3.data.data.Info
        }
        initchart(data.nodeselflabel)
      }

      const clearHander = () => {
        let myChart = echarts.init(document.getElementById("myChart"));
        let miniChart = echarts.init(document.getElementById("miniChart"));
        myChart.clear()
        miniChart.clear()
    }

      const onStrokeChange = (e: any) => {
        const val = e.target.value;
        globalGridAttr.nodeStroke = val;
        curCel?.attr('body/stroke', val);
      };

      const onStrokeWidthChange = (val: number) => {
        // console.log(val)
        globalGridAttr.nodeStrokeWidth = val;
        curCel?.attr('body/strokeWidth', val);
      };

      const onFillChange = (e: any) => {
        const val = e.target.value;
        globalGridAttr.nodeFill = val;
        curCel?.attr('body/fill', val);
      };

      const onFontSizeChange = (val: number) => {
        globalGridAttr.nodeFontSize = val;
        curCel?.attr('text/fontSize', val);
      };

      const onColorChange = (e: any) => {
        const val = e.target.value;
        globalGridAttr.nodeColor = val;
        curCel?.attr('text/fill', val);
      };

      const ondownsamplescaleChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          downsamplescale: val,
        });
      };
      const onaugmentationscaleChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          augmentationscale: val,
        });
      };


      return {
        globalGridAttr,
        data,
        onStrokeChange,
        onStrokeWidthChange,
        onFillChange,
        onFontSizeChange,
        onColorChange,
        ondownsamplescaleChange,
        onaugmentationscaleChange,
        // initchart,
        sendMoneytoYe,
        bashinfo
      };
    },
  });
</script>

<style lang="less" scoped>

.config{
  width: 300px;
  text-align: center
}
.view{
  width: 300px;
  background-color: #fff;
  margin-top: 10px;

}
.miniChart{
  margin-top: 0px;
  margin-left: 10px;
}


</style>
