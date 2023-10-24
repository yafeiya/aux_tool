<template>
  <Tabs class="config" value="1">
    <TabPane label="属性" name="1">
      <Row
      align="middle"
      v-if="globalGridAttr.selflabel === '仿真监听数据'"
      >
        <Upload  :action="EndUrl().backEndUrl+'/uploadLoss'"
                 style="margin-bottom: 10px"
                 method="POST"
                 accept=".csv"
                 :data="lossData"
                 :show-upload-list=false
                 :on-success="uploadSuccess"
        >
          <Button class="view" icon="ios-cloud-upload-outline" @click="uploadfile">选择网络损失数据</Button>
        </Upload>
        <Upload  :action="EndUrl().backEndUrl+'/uploadReward'"
                 style="margin-bottom: 10px"
                 method="POST"
                 accept=".txt"
                 :data="lossData"
                 :show-upload-list=false
                 :on-success="uploadSuccess"
        >
          <Button class="view" icon="ios-cloud-upload-outline" @click="uploadfile">选择仿真奖励数据</Button>
        </Upload>
        <Upload  :action="EndUrl().backEndUrl+'/uploadActions'"
                 style="margin-bottom: 10px"
                 method="POST"
                 accept=".json"
                 :data="lossData"
                 :show-upload-list=false
                 :on-success="uploadSuccess"
        >
          <Button class="view" icon="ios-cloud-upload-outline" @click="uploadfile">选择训练动作数据</Button>
        </Upload>
      </Row>

      <Row
      align="middle"
      v-if="globalGridAttr.selflabel === '模型日志数据'"
      >
        <Upload action="//localhost:3000/">
          <Button class="view" icon="ios-cloud-upload-outline" @click="">选择模型日志数据</Button>
        </Upload>
      </Row>

      <div
      v-for ="children1 in hasChildren(menu)"
      >
        <div
        v-for ="children2 in children1.children"
        >
          <div
          v-for ="childrenitem in children2.children"
          >

            <Row align="middle"
            v-if="(globalGridAttr.selflabel === childrenitem.title)&&(globalGridAttr.selflabel != '仿真监听数据')&&(globalGridAttr.selflabel != '模型日志数据')"
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

      <div v-if="(globalGridAttr.selflabel != '仿真监听数据')&&(globalGridAttr.selflabel != '模型日志数据')">

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
    </div>

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
  import { defineComponent, inject, watch, reactive, ref} from 'vue';
  import { Cell } from '@antv/x6/lib';
  import { nodeOpt } from './method';
  import {EndUrl} from "../../../../../../url_config";
  import { getCard, getMenuInfo} from '../../../../../api/api.js'


  export default defineComponent({
    name: 'Index',
    data() {
      return {
       lossData: {
          id:'',
          type: '',
          task: '',
        }

      }
    },
    methods: {
      EndUrl,
      uploadfile(){
        var url = decodeURI(window.location.href);
        var cs_arr = url.split('?')[1];//?后面的
        var id = cs_arr.split('=')[1].split('&')[0];
        var task = cs_arr.split('=')[2].split('&')[0];
        var type = cs_arr.split('=')[3];
        this.lossData.id=id,
        this.lossData.task=task,
        this.lossData.type=type
        // console.info("iiiiiiiiiiiiiiiiiiiiii",this.lossData)

      },
      uploadSuccess(){
        this.$Message.success('上传成功')
        
      },
    },

    async setup() {


      const globalGridAttr: any = inject('globalGridAttr');
      const id: any = inject('id');
      let curCel: Cell;
      let writepath :any =inject('datapath')



      const changewritepath = () => {
        writepath.value = !writepath.value
        // console.log(writepath.value)
      }

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
          data.nodeStrokeWidth = globalGridAttr.nodeStrokeWidth;
          data.nodeFill = globalGridAttr.nodeFill;
          data.nodeFontSize = globalGridAttr.nodeFontSize;
          data.nodeLabelname = globalGridAttr.nodename;
          data.nodeselflabel = globalGridAttr.selflabel;
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
          "name" :"database",
          "title": "数据加载",
          "children": []
        },
        ]

      let database = await getMenuInfo("database")
      menu[0]["children"] = database.data;

      // console.log("显示初始菜单",await menu)
      let i = 0;
      while (menu[0]["children"][i] != null) {
        // console.log(i)
        // console.log(menu[0]["children"][i].title)
        let j=0
        while (menu[0]["children"][i]["children"][j] != null) {
          // console.log(menu[0]["children"][i]["children"][j].title)

          let tip = {
            pageKind: 'database',
            task: menu[0]["children"][i]["children"][j].title,
            Type: menu[0]["children"][i].title,
          }

          databasecards = await getCard(tip)
          if(databasecards.data.data != null){
            let cardList = await databasecards.data.data.database
            // console.log("读取数据库",await cardList)
            menu[0]["children"][i]["children"][j].content = cardList
          }
          j++
        }
        i++
        // console.log("显示最终菜单",await menu)
      }

      // console.log("显示最终菜单",await menu)



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
        menu,
        hasChildren,
        noChildren,
        changewritepath,
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
.view{
  width: 300px;
  background-color: #fff;
  margin-top: 0px;

}
</style>
