<template>
  <Tabs class="config" value="1">
    <TabPane label="超参数" name="1">

      <div
      v-for ="children1 in hasChildren(menu)"
      >
        <div
        v-for ="children2 in children1.children"
        >
          <div
          v-for ="childrenitem in children2.children"
          >
            <div v-if="!childrenitem.children">
              <Row align="middle"
              v-if="globalGridAttr.selflabel === childrenitem.title"
              >
                <Col :span="10">指定{{childrenitem.title}}</Col>
                <Col :span="12">
                  <Select
                    style="width: 100%"
                    v-model="data.nodeLabelname "
                    @on-change="onLabelChange"
                  >
                    <Option
                    v-for="option in childrenitem.content"
                    :value="option.Dataset_name"
                    >
                    {{option.Dataset_name}}
                    </Option>
                  </Select>
                </Col>
              </Row>
            </div>

            <div v-else-if="childrenitem.children">
              <div v-for="subchildrenitem in childrenitem.children">
                <Row align="middle"
                v-if="globalGridAttr.selflabel === subchildrenitem.title"
                >
                  <Col :span="10">指定{{subchildrenitem.title}}</Col>
                  <Col :span="12">
                    <Select
                      style="width: 100%"
                      v-model="data.nodeLabelname "
                      @on-change="onLabelChange"
                    >
                      <Option
                      v-for="option in subchildrenitem.content"
                      :value="option.Dataset_name"
                      >
                      {{option.Dataset_name}}
                      </Option>
                    </Select>
                  </Col>
                </Row>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Row class="params" align="middle">
        <Col :span="8">网络层数</Col>
        <Col :span="14">
          <Input v-model="data.Network_num" style="width: 100%" placeholder="default 256" @change="onNetwork_numChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">学习率</Col>
        <Col :span="14">
          <Input v-model="data.learning_rate" style="width: 100%" placeholder="default 0.001" @change="onlearningrateChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">优化器</Col>
        <Col :span="14">
          <Input v-model="data.Optimizer" style="width: 100%" placeholder="default SGD" @change="onOptimizerChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">迭代次数</Col>
        <Col :span="14">
          <Input v-model="data.Epoch_num" style="width: 100%" placeholder="default 10000" @change="onEpoch_numChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">批大小</Col>
        <Col :span="14">
          <Input v-model="data.batch" style="width: 100%" placeholder="default 256" @change="onbatchChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">激活函数</Col>
        <Col :span="14">
          <Input v-model="data.Act_function" style="width: 100%" placeholder="default RELU" @change="onAct_functionChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">衰减因子</Col>
        <Col :span="14">
          <Input v-model="data.Decay_factor" style="width: 100%" placeholder="default 0.95" @change="onDecay_factorChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">探索率</Col>
        <Col :span="14">
          <Input v-model="data.Explore_rate" style="width: 100%" placeholder="default 0.9" @change="onExplore_rateChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">随机种子</Col>
        <Col :span="14">
          <Input v-model="data.Radom_seed" style="width: 100%" placeholder="default 42" @change="onRadom_seedChange" />
        </Col>
      </Row>

    </TabPane>
    <TabPane label="样式" name="2">
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
    <TabPane label="字体" name="3">
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
  import { defineComponent, inject, watch, reactive, ref} from 'vue';
  import { Cell } from '@antv/x6/lib';
  import { nodeOpt } from './method';
  import { getCard, getMenuInfo} from '../../../../../api/api.js'




  export default defineComponent({
    name: 'Index',

    async setup() {
      const globalGridAttr: any = inject('globalGridAttr');
      const id: any = inject('id');
      let curCel: Cell;

      const sendtoYe = function() {
        data.ModelModal=true
        console.log('ModelModal', data.ModelModal)

      }

      const data = reactive({
        Network_num: '',
        learning_rate: '',
        Epoch_num:'',
        Act_function:'',
        Decay_factor:'',
        Explore_rate:'',
        Radom_seed:'',
        Optimizer: '',
        nodemodeltype:'',
        nodeStrokeWidth: '',
        nodeFill: '',
        nodeFontSize: '',
        nodeLabelname:'',
        nodeselflabel:'',
        ModelModal:false,
        nodemodelbatch:'',
        batch:'',
      })
      watch(
        [() => id.value],
        () => {
          curCel = nodeOpt(id, globalGridAttr);
          data.Network_num = globalGridAttr.Network_num;
          data.learning_rate = globalGridAttr.learning_rate;
          data.Optimizer = globalGridAttr.Optimizer;
          data.Epoch_num = globalGridAttr.Epoch_num;
          data.batch = globalGridAttr.batch;
          data.Act_function = globalGridAttr.Act_function;
          data.Decay_factor = globalGridAttr.Decay_factor;
          data.Explore_rate = globalGridAttr.Explore_rate;
          data.Radom_seed = globalGridAttr.Radom_seed;

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

      let databasecards:any  = ref([])
      // let databaselist:any  = ref([])
      let menu:any = [
        {
          "name" :"modelbase",
          "title": "模型模板",
          "children": []
        },
        ]

      let finalmenu:any = [

        ]

      let modelbase = await getMenuInfo("modelbase")
      menu[0]["children"] = modelbase.data;

      // console.log("显示初始菜单",await menu)
      let i = 0;
      while (menu[0]["children"][i] != null) {
        // console.log(i)
        // console.log(menu[0]["children"][i].title)
        let j=0
        while (menu[0]["children"][i]["children"][j] != null) {
          // console.log(menu[0]["children"][i]["children"][j].title)

          let tip = {
            pageKind: 'modelbase',
            task: menu[0]["children"][i]["children"][j].title,
            Type: menu[0]["children"][i].title,
          }

          databasecards = await getCard(tip)
          if(databasecards.data.data != null){
            let cardList = await databasecards.data.data.modelbase
            // console.log("读取数据库",await cardList)
            menu[0]["children"][i]["children"][j].content = cardList
          }
          finalmenu.push(menu[0]["children"][i]["children"][j])
          j++
        }
        i++
        // console.log("显示最终菜单",await menu)
        // console.log("显示finalmenu",await finalmenu)
      }

      // while (menu[0]["children"][i] != null) {

      // }



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

      const onNetwork_numChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          Network_num: val,
        });
      };
      const onlearningrateChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          learning_rate: val,
        });
      };

      const onOptimizerChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          Optimizer: val,
        });
      };
      const onEpoch_numChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          Epoch_num: val,
        });
      };
      const onbatchChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          batch: val,
        });
      };
      const onAct_functionChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          Act_function: val,
        });
      };
      const onDecay_factorChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          Decay_factor: val,
        });
      };
      const onExplore_rateChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          Explore_rate: val,
        });
      };
      const onRadom_seedChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          Radom_seed: val,
        });
      };
      const onmodeltypeChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          modeltype: val,
        });
      };
      const onmodelbatchChange = (e: any) => {
        const val = e.target.value;
        curCel?.setData({
          modelbatch: val,
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
        onNetwork_numChange,
        onlearningrateChange,
        onOptimizerChange,
        onAct_functionChange,
        onEpoch_numChange,
        onExplore_rateChange,
        onDecay_factorChange,
        onbatchChange,
        onRadom_seedChange,
        onmodelbatchChange,
        onmodeltypeChange,
        menu,
        hasChildren,
        noChildren,
        onLabelChange,
        sendtoYe
      };
    },
  });
</script>

<style lang="less" scoped>

.config{
  width: 300px;
  text-align: center
}
.params{
  margin: 10px;
}

</style>
