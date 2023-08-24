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
      <div class="miniChart" id="miniChart" :style="{width: '310px', height: '250px'}"></div>
        
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
  import { defineComponent, inject, watch, reactive, onMounted } from 'vue';
  import { Cell } from '@antv/x6/lib';
  import { nodeOpt } from './method';
  import * as echarts from 'echarts'
  import { Button } from 'view-ui-plus';

  export default defineComponent({
    name: 'Index',

    data () {
            return {
                modal: false
            }
        },
    methods: {
            ok () {
                this.$Message.info('预览完毕');
            },
        },

    setup() {
      const globalGridAttr: any = inject('globalGridAttr');
      const id: any = inject('id');
      let curCel: Cell;
      let Buttonid = false

      const data = reactive({
        nodedownsamplescale: '',
        nodeaugmentationscale: '',
        nodeStrokeWidth: '',
        nodeFill: '',
        nodeFontSize: '',
        nodeselflabel:''
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
          initchart(data.nodeselflabel)
        },
        {
          immediate: false,
          deep: false,
        },
      );

      const changeButton = () => {
        Buttonid = !Buttonid
        console.log(Buttonid)
      }

      const actdata = [
            ["act1", "act2", "act3", "act4", "act5", "act6", "act7"],
            [8, 15, 31, 13, 15, 22, 11],
          ]
      const rewarddata = [
          ["re1", "re2", "re3", "re4", "re5", "re6", "re7"],
          [1, 2, 3, 4, 5, 6, 7],
        ]
      const lrdata = [
          ["lr1", "lr2", "lr3", "lr4", "lr5", "lr6", "lr7"],
          [0.1, 0.1, 0.1, 0.1, 0.01, 0.01, 0.01],
        ]  

      const initchart = (label) => {
        let myChart = echarts.init(document.getElementById("myChart"));
        let miniChart = echarts.init(document.getElementById("miniChart"));
        // 绘制图表
        const nodelabel = label

        if( nodelabel == '动作分布'){
          myChart.setOption({
            xAxis: {
              data: actdata[0]
            },
            yAxis:{},
            series: [
              {
                name: "y轴",
                type: "line",
                data: actdata[1]
              }
            ]
          });
        }

        else if( nodelabel == '奖励分布'){
          myChart.setOption({
            xAxis: {
              data: rewarddata[0]
            },
            yAxis:{},
            series: [
              {
                name: "y轴",
                type: "line",
                data: rewarddata[1]
              }
            ]
          }); 
        }

        else if( nodelabel == '学习率'){
          myChart.setOption({
            xAxis: {
              data: lrdata[0]
            },
            yAxis:{},
            series: [
              {
                name: "y轴",
                type: "line",
                data: lrdata[1]
              }
            ]
          }); 
        }


        if( nodelabel == '动作分布'){
          miniChart.setOption({
            xAxis: {
              data: actdata[0]
            },
            yAxis:{},
            series: [
              {
                name: "y轴",
                type: "line",
                data: actdata[1]
              }
            ]
          });
        }

        else if( nodelabel == '奖励分布'){
          miniChart.setOption({
            xAxis: {
              data: rewarddata[0]
            },
            yAxis:{},
            series: [
              {
                name: "y轴",
                type: "line",
                data: rewarddata[1]
              }
            ]
          }); 
        }

        else if( nodelabel == '学习率'){
          miniChart.setOption({
            xAxis: {
              data: lrdata[0]
            },
            yAxis:{},
            series: [
              {
                name: "y轴",
                type: "line",
                data: lrdata[1]
              }
            ]
          }); 
        }
        
        window.onresize = function () { // 自适应大小
          myChart.resize();
          miniChart.resize();
        };
      
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
        changeButton,
        Buttonid,
      };
    },
  });
</script>

<style lang="less" scoped>

.config{
  width: 300px;
  text-align: center, 
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
