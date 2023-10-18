<template>
  <Tabs class="config" value="1">
    <TabPane label="属性" name="1">
      <Row align="middle">
        <Button class="view" @click="previewImage">预览{{data.nodeselflabel}}大图</Button>
        <Modal
          v-model="imagemodal"
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
  import { defineComponent, inject, watch, reactive} from 'vue';
  import { Cell } from '@antv/x6/lib';
  import { nodeOpt } from './method';
  import * as echarts from 'echarts'
  import 'echarts-gl'
  import { Button } from 'view-ui-plus';
  // import datas from "../../../results/data.json"
  import {downloadModelFile, getprocessFile} from '../../../../../api/api.js'
  import qs from "qs";

  export default defineComponent({
    name: 'Index',
    data () {
            return {
              actionData:[],
                modal: false,
              imagemodal:false,
              cardInfo: {
                id:'',
                type: '',
                task: '',
                processfile:''
              }
            }
        },
    methods: {
            ok () {
                this.$Message.info('预览完毕');
            },
      previewImage(){
        this.imagemodal=true
        var url = decodeURI(window.location.href);
        var cs_arr = url.split('?')[1];//?后面的
        var id = cs_arr.split('=')[1].split('&')[0];
        var task = cs_arr.split('=')[2].split('&')[0];
        var type = cs_arr.split('=')[3];
        this.cardInfo.id=id,
        this.cardInfo.task=task,
        this.cardInfo.type=type,
        this.cardInfo.processFile='actions.json'
        let data=this.cardInfo
        data = qs.stringify(data)
        console.info("iiiiiiiiiiiiiiiiiiiiii",data)
        getprocessFile(data).then(response => {
          this.actionData=response.data.data.Info
          console.info("rrrrrrrrrr",this.actionData)
        })
      },
      initchart(label){
        // console.log("initchart",readpath.value)
        // clearHander()
        let actdata = this.actionData["actions"]
        console.info("aaaaaaa",actdata)
        let rewarddata = this.datas["reward"]
        let lrdata = this.datas["learning_rate"]
        let count:any = []  //统计轮数

        // let actx:any = ["AIR1","AIR2","AIR3","AIR4"] //act图x坐标飞机种类
        // const acty= ["MSL1","MSL2","MSL3","MSL4","MSL5","MSL6","MSL7"]  //act图y坐标导弹种类
        const x = 4;  //飞机种类数
        const  y = 7; //导弹种类数
        let maxnum = 0
        let MSLnum = new Array(x); // MSL记录数据
        let MSLdata:any = []

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
        let msi:any
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
        // console.log(count)
        // console.log(MSLnum)
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

        //定义lr图表参数
        var option3={
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
        }

        if( nodelabel == '动作分布'){
          myChart.setOption(option1);
          miniChart.setOption(option1);
        }

        else if( nodelabel == '奖励分布'){
          myChart.setOption(option2);
          miniChart.setOption(option2);
        }

        else if( nodelabel == '学习率'){
          myChart.setOption(option3);
          miniChart.setOption(option3);
        }

        window.onresize = function () { // 自适应大小
          myChart.resize();
          miniChart.resize();
        };

      }
        },
    setup() {
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
        console.log(readpath.value)
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
            // initchart(data.nodeselflabel)
            // console.log("inif",readpath.value)
          }else {
            clearHander()
          }
        },
        {
          immediate: false,
          deep: false,
        },
      );




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
