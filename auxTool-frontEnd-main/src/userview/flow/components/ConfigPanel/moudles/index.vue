<template>
  <Tabs class="config" value="1">
    <TabPane label="属性" name="1">

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
        <Col :span="8">批量大小</Col>
        <Col :span="14">
          <Input v-model="data.nodemodelbatch" style="width: 100%" @change="onmodelbatchChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">优化器</Col>
        <Col :span="14">
          <Input v-model="data.nodemodeltype" style="width: 100%" @change="onmodeltypeChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">训练轮数</Col>
        <Col :span="14">
          <Input v-model="data.nodemodeltype" style="width: 100%" @change="onmodeltypeChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">学习率</Col>
        <Col :span="14">
          <Input v-model="data.nodemodeltype" style="width: 100%" @change="onmodeltypeChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">衰减因子</Col>
        <Col :span="14">
          <Input v-model="data.nodemodeltype" style="width: 100%" @change="onmodeltypeChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">GPU占用数</Col>
        <Col :span="14">
          <Input v-model="data.nodemodeltype" style="width: 100%" @change="onmodeltypeChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">内存用量</Col>
        <Col :span="14">
          <Input v-model="data.nodemodeltype" style="width: 100%" @change="onmodeltypeChange" />
        </Col>
      </Row>
      <Row class="params" align="middle">
        <Col :span="8">CPU核数</Col>
        <Col :span="14">
          <Input v-model="data.nodemodeltype" style="width: 100%" @change="onmodeltypeChange" />
        </Col>
      </Row>
<!--      <Button class="view" @click="sendtoYe">模型超参数</Button>-->
<!--      <Modal-->
<!--          v-model="data.ModelModal"-->
<!--          title="模型超参数"-->
<!--          width="40px"-->
<!--          >-->
<!--        <Row>-->
<!--          <Col span="12"><strong>批量大小: </strong> 11111</Col>-->
<!--          <Col span="12"><strong>优化器: </strong> 11111</Col>-->
<!--          <Col span="12"><strong>训练轮数: </strong> 11111</Col>-->
<!--          <Col span="12"><strong>学习率: </strong> 11111</Col>-->
<!--          <Col span="12"><strong>衰减因子: </strong> 11111</Col>-->
<!--          <Col span="12"><strong>GPU占用数: </strong> 11111</Col>-->
<!--          <Col span="12"><strong>内存用量: </strong> 11111</Col>-->
<!--          <Col span="12"><strong>CPU核数: </strong> 11111</Col>-->
<!--        </Row>-->
<!--      </Modal>-->
      <!-- <Image src="http://127.0.0.1:5173/机器学习/回归/10/home.png" fit="fill" width="400px" height="220px" alt=""></Image> -->
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
  import { getCard, getMenuInfo} from '../../../../../api/api.js'



  
  export default defineComponent({
    name: 'Index',

    async setup() {
      const globalGridAttr: any = inject('globalGridAttr');
      const id: any = inject('id');
      let curCel: Cell;


      // 向画布页面（祖父）传递预览图参数
      const changeimgpreview: any = inject('changeimgpreview')
      const imgpreview: any = inject('imgpreview')
      var imgpath: any = inject('imgpath')

      // const sendtoYe = function() {
      //   console.log('触发sendtoYe', globalGridAttr.nodename)
      //   for(var children1 of menu[0].children){
      //     // console.log(children1)
      //     for(var children2 of children1.children){
      //       // console.log(children2)
      //       if(children2.content != null){
      //         for(var item of children2.content){
      //           if(item.Dataset_name === globalGridAttr.nodename){
      //             var path = item.Image_path
      //             console.log(path)
      //           }
      //         }
      //       }
      //     }
      //   }
      //   imgpath = path
      //
      //   imgpreview.value = !imgpreview.value
      //   changeimgpreview(imgpreview.value,imgpath)
      //   // console.log(imgpath)
      // }
      const sendtoYe = function() {
        data.ModelModal=true
        console.log('ModelModal', data.ModelModal)

      }

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
        nodeselflabel:'',
        ModelModal:false,
        nodemodelbatch:'',
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
        onnetworkdepthChange,
        onclassnumChange,
        onfuturerewarddiscountChange,
        onmodelurlChange,
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
  text-align: center, 
}
.params{
  margin: 10px;
}

</style>
