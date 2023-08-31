<template>
  <Tabs class="config" model-value="1">
    <TabPane label="网格" name="1">
      <Row align="middle" >
        <Col :span="10">网格类型</Col>
        <Col :span="12">
          <Select
            style="width: 100%"
            :value="globalGridAttr.type"
            @on-change="
              (val) => {
                globalGridAttr.type = val;
              }
            "
          >
            <Option :value="GRID_TYPE.DOT">Dot</Option>
            <Option :value="GRID_TYPE.FIXED_DOT">Fixed Dot</Option>
            <Option :value="GRID_TYPE.MESH">Mesh</Option>
            <Option :value="GRID_TYPE.DOUBLE_MESH">Double Mesh</Option>
          </Select>
        </Col>
      </Row>
      <Row align="middle">
        <Col :span="10">网格尺寸</Col>
        <Col :span="10">
          <Slider
            :min="1"
            :max="20"
            :step="1"
            :value="globalGridAttr.size"
            @on-input="
              (val) => {
                globalGridAttr.size = val;
              }
            "
          />
        </Col>
        <Col :span="2">
          <div class="result">{{ globalGridAttr.size }}</div>
        </Col>
      </Row>
      <div v-if="globalGridAttr.type === 'doubleMesh'">
        <Row align="middle">
          <Col :span="10">首级颜色</Col>
          <Col :span="12">
            <Input
              type="color"
              :value="globalGridAttr.color"
              style="width: 100%"
              @change="
                (e) => {
                  globalGridAttr.color = e.target.value;
                }
              "
            />
          </Col>
        </Row>
        <Row align="middle">
          <Col :span="10">首级粗细</Col>
          <Col :span="10">
            <Slider
              :min="0.5"
              :max="10"
              :step="0.5"
              :value="globalGridAttr.thickness"
              @on-input="
                (val) => {
                  globalGridAttr.thickness = val;
                }
              "
            />
          </Col>
          <Col :span="2">
            <div class="result">{{ globalGridAttr.thickness.toFixed(1) }}</div>
          </Col>
        </Row>
        <Row align="middle">
          <Col :span="10">次级颜色</Col>
          <Col :span="12">
            <Input
              type="color"
              :value="globalGridAttr.colorSecond"
              style="width: 100%"
              @change="
                (e) => {
                  globalGridAttr.colorSecond = e.target.value;
                }
              "
            />
          </Col>
        </Row>
        <Row align="middle">
          <Col :span="10">次级粗细</Col>
          <Col :span="10">
            <Slider
              :min="0.5"
              :max="10"
              :step="0.5"
              :value="globalGridAttr.thicknessSecond"
              @on-input="
                (val) => {
                  globalGridAttr.thicknessSecond = val;
                }
              "
            />
          </Col>
          <Col :span="2">
            <div class="result">{{ globalGridAttr.thicknessSecond.toFixed(1) }}</div>
          </Col>
        </Row>
        <Row align="middle">
          <Col :span="10">标尺尺寸</Col>
          <Col :span="10">
            <Slider
              :min="1"
              :max="10"
              :step="1"
              :value="globalGridAttr.factor"
              @on-input="
                (val) => {
                  globalGridAttr.factor = val;
                }
              "
            />
          </Col>
          <Col :span="2">
            <div class="result">{{ globalGridAttr.factor }}</div>
          </Col>
        </Row>
      </div>
      <div v-else>
        <Row align="middle">
          <Col :span="10">网格颜色</Col>
          <Col :span="12">
            <Input
              type="color"
              :value="globalGridAttr.color"
              style="width: 100%"
              @change="
                (e) => {
                  globalGridAttr.color = e.target.value;
                }
              "
            />
          </Col>
        </Row>
        <Row align="middle">
          <Col :span="10">格线粗细</Col>
          <Col :span="10">
            <Slider
              :min="0.5"
              :max="10"
              :step="0.5"
              :value="globalGridAttr.thickness"
              @on-input="
                (val) => {
                  globalGridAttr.thickness = val;
                }
              "
            />
          </Col>
          <Col :span="1">
            <div class="result">{{ globalGridAttr.thickness.toFixed(1) }}</div>
          </Col>
        </Row>
      </div>
    </TabPane>
  </Tabs>
</template>

<script lang="ts">
  import { defineComponent, reactive, inject, watch } from 'vue';
  // import FlowGraph from '@/views/graph'

  import { gridOpt, gridSizeOpt } from './method';

  enum GRID_TYPE_ENUM {
    DOT = 'dot',
    FIXED_DOT = 'fixedDot',
    MESH = 'mesh',
    DOUBLE_MESH = 'doubleMesh',
  }

  export default defineComponent({
    name: 'Index',
    setup() {
      const GRID_TYPE = reactive(GRID_TYPE_ENUM);
      // const REPEAT_TYPE = reactive(REPEAT_TYPE_ENUM);
      const globalGridAttr: any = inject('globalGridAttr');

      gridOpt(globalGridAttr);
      gridSizeOpt(globalGridAttr);

      // 监听网格变化
      watch(
        [
          () => globalGridAttr.type,
          () => globalGridAttr.color,
          () => globalGridAttr.thickness,
          () => globalGridAttr.thicknessSecond,
          () => globalGridAttr.colorSecond,
          () => globalGridAttr.factor,
        ],
        () => {
          gridOpt(globalGridAttr);
        },
        {
          immediate: false,
          deep: false,
        },
      );

      // 监听网格大小变化
      watch(
        [() => globalGridAttr.size],
        () => {
          gridSizeOpt(globalGridAttr);
        },
        {
          immediate: false,
          deep: false,
        },
      );

      return {
        GRID_TYPE,
        globalGridAttr,
      };
    },
  });
</script>

<style lang="less" scoped>
.config{
  height: 500px;
  width: 300px;
  text-align: center, 
}
</style>
