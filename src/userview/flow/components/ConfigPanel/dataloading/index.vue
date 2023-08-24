<template>
  <Tabs class="config" value="1">
    <TabPane label="属性" name="1">

      <div
      v-for ="item in hasChildren(list)"
      >
        <div
        v-for ="childrenitem in item.children"
        >
        <div v-if="!childrenitem.children">
          <Row align="middle" 
          v-if="globalGridAttr.selflabel === childrenitem.label"
          >
            <Col :span="10">指定{{childrenitem.label}}</Col>
            <Col :span="12">
              <Select
                style="width: 100%"
                v-model="data.nodeLabelname "
                @on-change="onLabelChange"
              >
                <Option 
                v-for="option in childrenitem.content"
                :value="option.name"
                >
                {{option.name}}
                </Option>
              </Select>
            </Col>
          </Row>
        </div>

        <!-- <div v-if="childrenitem.children">
          <div v-for="subchildrenitem in childrenitem.children">
            <Row align="middle" 
            v-if="globalGridAttr.selflabel === subchildrenitem.label"
            >
              <Col :span="10">指定{{subchildrenitem.label}}</Col>
              <Col :span="12">
                <Select
                  style="width: 100%"
                  v-model="data.nodeLabelname "
                  @on-change="onLabelChange"
                >
                  <Option 
                  v-for="option in subchildrenitem.content"
                  :value="option.name"
                  >
                  {{option.name}}
                  </Option>
                </Select>
              </Col>
            </Row>
          </div>
        </div> -->
        </div>
      </div>

      <!-- <Row align="middle" 
      v-if="globalGridAttr.selflabel === '数值数据集'"
      >
        <Col :span="10">指定数值数据集</Col>
        <Col :span="12">
          <Select
            style="width: 100%"
            v-model="data.nodeLabelname "
            @on-change="onLabelChange"
          >
            <Option 
            v-for="option in selectoptions2"
            :value="option.name"
            >
            {{option.name}}
            </Option>
          </Select>
        </Col>
      </Row> -->



      <Row class="params" align="middle">
        <Col :span="8">数据路径</Col>
        <Col :span="14">
          <Input v-model="data.nodedataurl" style="width: 100%" @change="ondataurlChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">表头数量</Col>
        <Col :span="14">
          <Input v-model="data.nodeheadernum" style="width: 100%" @change="onheadernumChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">数据集大小</Col>
        <Col :span="14">
          <Input v-model="data.nodedatasetsize" style="width: 100%" @change="ondatasetsizeChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">图像尺寸</Col>
        <Col :span="14">
          <Input v-model="data.nodeimgsize" style="width: 100%" @change="onimgsizeChange" />
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
  import { menulist } from '../../../list'


  
  export default defineComponent({
    name: 'Index',

    setup() {
      const globalGridAttr: any = inject('globalGridAttr');
      const id: any = inject('id');
      let curCel: Cell;
      const list = menulist()


      const data = reactive({
        nodedataurl: '',
        nodeheadernum: '',
        nodedatasetsize: '',
        nodeimgsize: '',
        nodeStrokeWidth: '',
        nodeFill: '',
        nodeFontSize: '',
        nodeLabelname:'',
        nodeselflabel:''
      })
      watch(
        [() => id.value],
        () => {
          curCel = nodeOpt(id, globalGridAttr);
          data.nodedataurl = globalGridAttr.nodedataurl;
          data.nodeheadernum = globalGridAttr.nodeheadernum;
          data.nodedatasetsize = globalGridAttr.nodedatasetsize;
          data.nodeimgsize = globalGridAttr.nodeimgsize;
          data.nodeStrokeWidth = globalGridAttr.nodeStrokeWidth
          data.nodeFill = globalGridAttr.nodeFill
          data.nodeFontSize = globalGridAttr.nodeFontSize
          data.nodeLabelname = globalGridAttr.nodename 
          data.nodeselflabel = globalGridAttr.selflabel
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

      const ondataurlChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          dataurl: val,
        });
      };
      const onheadernumChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          headernum: val,
        });
      };  

      const ondatasetsizeChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          datasetsize: val,
        });
      };  

      const onimgsizeChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          imgsize: val,
        });
      }; 
      
      const onLabelChange = (value) => {
        
        const val = value;
        globalGridAttr.nodename = val;
        curCel?.setData({
          name: val,
        });
        // curCel?.attr('text/text', val);
        // console.log(curCel.data.name)
      };

      const hasChildren =(thelist) =>{
          return thelist.filter((item) => item.children);
      };

      const noChildren =(thelist) =>{
          return thelist.filter((item) => !item.children);
      };


      return {
        globalGridAttr,
        data,
        onStrokeChange,
        onStrokeWidthChange,
        onFillChange,
        onFontSizeChange,
        onColorChange,
        ondataurlChange,
        onheadernumChange,
        ondatasetsizeChange,
        onimgsizeChange,
        onLabelChange,
        list,
        hasChildren,
        noChildren,
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
