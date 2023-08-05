<template>
  <Tabs class="config" value="1">
    <TabPane label="属性" name="1">
      <Row class="canshu" align="middle">
        <Col :span="8">参数1</Col>
        <Col :span="14">
          <Input v-model="data.nodecanshu1" style="width: 100%" @change="oncanshu1Change" />
        </Col>
      </Row>
      <Row class="canshu" align="middle">
        <Col :span="8">参数2</Col>
        <Col :span="14">
          <Input v-model="data.nodecanshu2" style="width: 100%" @change="oncanshu2Change" />
        </Col>
      </Row>
      <Row class="canshu" align="middle">
        <Col :span="8">参数3</Col>
        <Col :span="14">
          <Input v-model="data.nodecanshu3" style="width: 100%" @change="oncanshu3Change" />
        </Col>
      </Row>
      <Row class="canshu" align="middle">
        <Col :span="8">参数4</Col>
        <Col :span="14">
          <Input v-model="data.nodecanshu4" style="width: 100%" @change="oncanshu4Change" />
        </Col>
      </Row>
      <Row class="canshu" align="middle">
        <Col :span="8">参数5</Col>
        <Col :span="14">
          <Input v-model="data.nodecanshu5" style="width: 100%" @change="oncanshu5Change" />
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
        nodecanshu1: '',
        nodecanshu2: '',
        nodecanshu3: '',
        nodecanshu4: '',
        nodecanshu5: '',
        nodeStrokeWidth: '',
        nodeFill: '',
        nodeFontSize: '',
      })
      watch(
        [() => id.value],
        () => {
          curCel = nodeOpt(id, globalGridAttr);
          data.nodecanshu1 = globalGridAttr.nodecanshu1;
          data.nodecanshu2 = globalGridAttr.nodecanshu2;
          data.nodecanshu3 = globalGridAttr.nodecanshu3;
          data.nodecanshu4 = globalGridAttr.nodecanshu4;
          data.nodecanshu5 = globalGridAttr.nodecanshu5;
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

      const oncanshu1Change = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          canshu1: val,
        });
      };
      const oncanshu2Change = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          canshu2: val,
        });
      };  

      const oncanshu3Change = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          canshu3: val,
        });
      };  

      const oncanshu4Change = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          canshu4: val,
        });
      };  

      const oncanshu5Change = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          canshu5: val,
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
        oncanshu1Change,
        oncanshu2Change,
        oncanshu3Change,
        oncanshu4Change,
        oncanshu5Change,

      };
    },
  });
</script>

<style lang="less" scoped>

.config{
  width: 300px;
  text-align: center, 
}
.canshu{
  margin: 10px;
}

</style>
