<template>
  <Tabs class="config" value="1">
    <TabPane label="属性" name="1">
      <Row class="params" align="middle">
        <Col :span="8">下采样因子</Col>
        <Col :span="14">
          <Input v-model="data.nodedownsamplescale" style="width: 100%" @change="ondownsamplescaleChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">扩充比例</Col>
        <Col :span="14">
          <Input v-model="data.nodeaugmentationscale" style="width: 100%" @change="onaugmentationscaleChange" />
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
        nodedownsamplescale: '',
        nodeaugmentationscale: '',
        nodeStrokeWidth: '',
        nodeFill: '',
        nodeFontSize: '',
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
