<template>
  <Tabs class="config" value="1">
    <TabPane label="属性" name="1">
      <Row class="params" align="middle">
        <Col :span="8">迭代次数</Col>
        <Col :span="14">
          <Input v-model="data.nodeiterations" style="width: 100%" @change="oniterationsChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">优化器</Col>
        <Col :span="14">
          <Input v-model="data.nodeoptimizer" style="width: 100%" @change="onoptimizerChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">衰减因子</Col>
        <Col :span="14">
          <Input v-model="data.nodedecayfactor" style="width: 100%" @change="ondecayfactorChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">学习率</Col>
        <Col :span="14">
          <Input v-model="data.nodelearningrate" style="width: 100%" @change="onlearningrateChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">损失函数</Col>
        <Col :span="14">
          <Input v-model="data.nodeloss" style="width: 100%" @change="onlossChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">评价指标</Col>
        <Col :span="14">
          <Input v-model="data.nodeevalution" style="width: 100%" @change="onevalutionChange" />
        </Col>
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


  
  export default defineComponent({
    name: 'Index',

    setup() {
      const globalGridAttr: any = inject('globalGridAttr');
      const id: any = inject('id');
      let curCel: Cell;
      const data = reactive({
        nodeiterations: '',
        nodeoptimizer: '',
        nodedecayfactor: '',
        nodelearningrate: '',
        nodeloss: '',
        nodeevalution: '',
        nodeStrokeWidth: '',
        nodeFill: '',
        nodeFontSize: '',
      })
      watch(
        [() => id.value],
        () => {
          curCel = nodeOpt(id, globalGridAttr);
          data.nodeiterations = globalGridAttr.nodeiterations;
          data.nodeoptimizer = globalGridAttr.nodeoptimizer;
          data.nodedecayfactor = globalGridAttr.nodedecayfactor;
          data.nodelearningrate = globalGridAttr.nodelearningrate;
          data.nodeloss = globalGridAttr.nodeloss;
          data.nodeevalution = globalGridAttr.nodeevalution;
          data.nodeStrokeWidth = globalGridAttr.nodeStrokeWidth
          data.nodeFill = globalGridAttr.nodeFill
          data.nodeFontSize = globalGridAttr.nodeFontSize
          // curCel?.attr('body/stroke', 'red');
        },
        {
          immediate: false,
          deep: false,
        },
      );



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

      const oniterationsChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          iterations: val,
        });
      };
      const onoptimizerChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          optimizer: val,
        });
      };  

      const ondecayfactorChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          decayfactor: val,
        });
      };  

      const onlearningrateChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          learningrate: val,
        });
      };  
      const onlossChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          loss: val,
        });
      };  
      const onevalutionChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          evalution: val,
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
        oniterationsChange,
        onoptimizerChange,
        ondecayfactorChange,
        onlearningrateChange,
        onlossChange,
        onevalutionChange,
      };
    },
  });
</script>

<style lang="less" scoped>

.config{
  width: 300px;
  text-align: center, 
}
.params{
  margin: 10px;
}

</style>
