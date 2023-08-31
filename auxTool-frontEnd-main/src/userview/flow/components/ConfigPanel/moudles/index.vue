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

        <div v-else-if="childrenitem.children">
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
        </div>
        </div>
      </div>


      <Row class="params" align="middle">
        <Col :span="8">网络深度</Col>
        <Col :span="14">
          <Input v-model="data.nodenetworkdepth" style="width: 100%" @change="onnetworkdepthChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">分类类别数</Col>
        <Col :span="14">
          <Input v-model="data.nodeclassnum" style="width: 100%" @change="onclassnumChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">未来奖励折扣</Col>
        <Col :span="14">
          <Input v-model="data.nodefuturerewarddiscount" style="width: 100%" @change="onfuturerewarddiscountChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">模型路径</Col>
        <Col :span="14">
          <Input v-model="data.nodemodelurl" style="width: 100%" @change="onmodelurlChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">任务类型</Col>
        <Col :span="14">
          <Input v-model="data.nodemodeltype" style="width: 100%" @change="onmodeltypeChange" />
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
        nodenetworkdepth: '',
        nodeclassnum: '',
        nodefuturerewarddiscount: '',
        nodemodelurl:'',
        nodemodeltype:'',
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
          data.nodenetworkdepth = globalGridAttr.nodenetworkdepth;
          data.nodeclassnum = globalGridAttr.nodeclassnum;
          data.nodefuturerewarddiscount = globalGridAttr.nodefuturerewarddiscount;
          data.nodemodelurl = globalGridAttr.nodemodelurl;
          data.nodemodeltype = globalGridAttr.nodemodeltype;
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

      const onnetworkdepthChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          networkdepth: val,
        });
      };
      const onclassnumChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          classnum: val,
        });
      };  

      const onfuturerewarddiscountChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          futurerewarddiscount: val,
        });
      };  
      const onmodelurlChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          modelurl: val,
        });
      };  
      const onmodeltypeChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          modeltype: val,
        });
      };  

      const hasChildren =(thelist) =>{
          return thelist.filter((item) => item.children);
      };

      const noChildren =(thelist) =>{
          return thelist.filter((item) => !item.children);
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



      return {
        globalGridAttr,
        data,
        onStrokeChange,
        onStrokeWidthChange,
        onFillChange,
        onFontSizeChange,
        onColorChange,
        onnetworkdepthChange,
        onclassnumChange,
        onfuturerewarddiscountChange,
        onmodelurlChange,
        onmodeltypeChange,
        list,
        hasChildren,
        noChildren,
        onLabelChange
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
