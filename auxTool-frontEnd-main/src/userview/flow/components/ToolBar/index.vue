<template>
  <div class="bar">
    <Tooltip content="返回" placement="bottom">
      <template #title>
      </template>
      <Button name="backHome" @click="toDesign" class="item-space" size="small"> 返回 </Button>
    </Tooltip>
    <Tooltip content="清除 (Cmd + D)" placement="bottom">
      <template #title>
      </template>
      <Button name="delete" @click="handleClick" class="item-space" size="small"> 清除 </Button>
    </Tooltip>
    <Tooltip content="撤销 (Cmd + Z)" placement="bottom">
      <template #title>
      </template>
      <Button :disabled="canUndo" name="undo" @click="handleClick" class="item-space" size="small"> 撤销 </Button>
    </Tooltip>
    <Tooltip content="重做 (Cmd + Shift + Z)" placement="bottom">
      <template #title>
      </template>
      <Button :disabled="canRedo" name="redo" @click="handleClick" class="item-space" size="small"> 重做 </Button>
    </Tooltip>

    <Tooltip content="复制 (Cmd + C)" placement="bottom">
      <template #title>
      </template>
      <Button name="copy" @click="handleClick" class="item-space" size="small"> 复制 </Button>
    </Tooltip>

    <Tooltip content="剪切 (Cmd + X)" placement="bottom">
      <template #title>
      </template>
      <Button name="cut" @click="handleClick" class="item-space" size="small"> 剪切 </Button>
    </Tooltip>

    <Tooltip content="粘贴 (Cmd + V)" placement="bottom">
      <template #title>
      </template>
      <Button name="paste" @click="handleClick" class="item-space" size="small"> 粘贴 </Button>
    </Tooltip>

    <Tooltip content="保存PNG (Cmd + S)" placement="bottom">
      <template #title>
      </template>
      <Button name="savePNG" @click="handleClick" class="item-space" size="small"> 保存PNG </Button>
    </Tooltip>

    <Tooltip content="保存SVG (Cmd + S)" placement="bottom">
      <template #title>
      </template>
      <Button name="saveSVG" @click="handleClick" class="item-space" size="small"> 保存SVG </Button>
    </Tooltip>

    <Tooltip content="打印 (Cmd + P)" placement="bottom">
      <template #title>
      </template>
      <Button name="print" @click="handleClick" class="item-space" size="small"> 打印 </Button>
    </Tooltip>

    <Tooltip content="保存" placement="bottom">
      <template #title>
      </template>
      <Button name="save"
      @mousedown="handleClick"
      @mouseup="info"
      class="item-space"
      size="small"
      > 保存 </Button>
    </Tooltip>

    <!-- //！！！！！！！！！！！！！在这儿修改！！！！！！！！！！ -->
    <Tooltip content="运行" placement="bottom">
      <template #title>
      </template>
      <Button name="run"  @click="isRun" class="run" size="small">
        运行
      </Button>
    </Tooltip>
    <Modal
        v-model="modal"
        title="即将运行该方案"
        @on-ok="RunExample({mustItem, example})"
        @on-cancel="cancel">
        <p>您已点击运行按钮，是否确认运行</p>

    </Modal>

  </div>
</template>

<script lang="ts">
  import { defineComponent, ref } from 'vue'; // ref, reactive
  import { runCanvas,saveCanvas,saveCanvasPNG } from '../../../../api/api.js'
  import FlowGraph from '../../graph';
  import { DataUri } from '@antv/x6';
  import axios from 'axios';
  import qs from "qs";
  import ViewUIPlus from 'view-ui-plus'
  import {EndUrl} from "../../../../../url_config";
  export default defineComponent({
    name: 'Index',
    components: {},
    data() {
      return {
        modal: false,
        nowTime : new Date(),
      // 检查必须拖动到画布上的模块
        mustItem :[
              {
                name:'数据加载',
                flag: false
              },
              {
                name:'模型模板',
                flag: false
              },
              {
                name:'模型训练',
                flag: false
              },
              {
                name:'仿真交互',
                flag: false
              }
                  ],
        example:{
          "id":"",
          "example_name": "",
          "Rank":          "",
          "State":         "运行中",
          "Task":          "",
          "Type":          "",
          "Dataset_name":  "Dataset_name",
          "Model_name":    "string",
          "Train_state":   "1",
          "Metics":        "string",
          "Network_num":   1,
          "Learning_rate": 0.01,
          "Act_function":  "RELu",
          "Radom_seed":    2,
          "Optimizer":     "string",
          "Batch_size":    64,
          "Epoch_num":     10,
          "Explore_rate":  0.1,
          "Decay_factor":  0.1,
          "start_time":(new Date()).getTime() ,
          "end_time": '' ,
      }
      }
    },
    methods: {
      isRun() {
        this.modal = true
      },
      toDesign() {
        this.$router.push('/design')
      },
      info() {
          this.$Message.info('画布已保存');
      },
      RunExample({mustItem, example}) {
        var url = decodeURI(window.location.href);
        var cs_arr = url.split('?')[1];//?后面的
        var iid = cs_arr.split('=')[1].split('&')[0];
        this.example.id=iid;
        example.start_time = (new Date()).getTime()
        example.end_time = ''
        var cs_arr = url.split('?')[1];//?后面的
        var id = cs_arr.split('=')[1].split('&')[0];
        var task = cs_arr.split('=')[2].split('&')[0];
        var type = cs_arr.split('=')[3];
        example.dataset_url=EndUrl().fileUrl+"/"+ type+"/"+task+"/"+id
        example.id =id
        const { graph } = FlowGraph;
        var graphData = graph.toJSON()
        console.info("graphData",graphData)
        const DatasetResult = graphData.cells
            .filter(cell => cell.data.grandLabel === "数据加载"&&(cell.data.fatherLabel.endsWith("数据集")))
            .map(cell => cell.data);
        var Dataset = DatasetResult.length > 0 ? DatasetResult : "没有找到符合条件的数据";
        example.Dataset_name=Dataset[0].name
        const ModelResult = graphData.cells
            .filter(cell => cell.data.grandLabel === "模型模板")
            .map(cell => cell.data);
        var Model = ModelResult.length > 0 ? ModelResult : "没有找到符合条件的数据";
        example.Model_name=Model[0].name
        example. Act_function=Model[0]. Act_function
        example.Decay_factor=Model[0].Decay_factor
        example.Epoch_num=Model[0].Epoch_num
        example.Explore_rate=Model[0].Explore_rate
        example.Optimizer=Model[0].Optimizer
        example.Radom_seed=Model[0].Radom_seed
        example.batch=Model[0].batch
        example.Network_num=Model[0].Network_num
        example.learning_rate=Model[0].learning_rate
        console.info("Model",Model)
        const TrainStateResult = graphData.cells
            .filter(cell => cell.data.grandLabel === "训练过程可视化")
            .map(cell => cell.data);
        var TrainState = TrainStateResult.length > 0 ? TrainStateResult : "没有找到符合条件的数据";
        example.Train_state=TrainState[0].selflabel
        console.info("example",example)
        runCanvas(example).then(res=>{
          this.$Message["success"]({
              background: true,
              content: '运行成功，请在实例管理页面查看详情'
          });

          //以下代码用于存png
          const { graph } = FlowGraph;
          var graphData = graph.toJSON()
          var url = decodeURI(window.location.href);
          var cs_arr = url.split('?')[1];//?后面的
          console.info("url",url);
          var iid = cs_arr.split('=')[1].split('&')[0];
          var cellsToSend = {
            id: iid,
            cells: graphData.cells // 假设 graphData 里有一个 cells 属性
          };
          graph.toPNG(
              (dataUri) => {
                // 将dataUri转换为Blob对象
                const byteString = atob(dataUri.split(',')[1]);
                const mimeString = dataUri.split(',')[0].split(':')[1].split(';')[0];
                const ab = new ArrayBuffer(byteString.length);
                const ia = new Uint8Array(ab);
                for (let i = 0; i < byteString.length; i++) {
                  ia[i] = byteString.charCodeAt(i);
                }
                const blob = new Blob([ab], { type: mimeString });

                // 创建FormData对象
                const formData = new FormData();
                const url = decodeURI(window.location.href);
                const cs_arr = url.split('?')[1]; //?后面的
                const id = cs_arr.split('=')[1].split('&')[0];
                const type = cs_arr.split('=')[2].split('&')[0];
                const task = cs_arr.split('=')[3];
                formData.append('image', blob, 'chartx.png');
                formData.append('id', id); // 替换为实际的id
                formData.append('type', task); // 替换为实际的type
                formData.append('task', type); // 替换为实际的task

                // 打印FormData的内容
                for (let [key, value] of formData.entries()) {
                  console.log(key, value);
                }

                // 使用axios发送POST请求到服务器
                saveCanvasPNG(formData).then(res => {
                  console.log('Success:', res.data);
                });
              },
              {
                backgroundColor: 'white',
                padding: {
                  top: 20,
                  right: 30,
                  bottom: 40,
                  left: 50,
                },
                quality: 1,
              },
          );
          saveCanvas(cellsToSend).then(res=>{
            console.log(res);
          })

          for(var i in mustItem) {
            mustItem[i].flag = false
          }
        })
      },
      cancel () {
        this.$Message.info('已取消');
      }

    },

    setup() {
      const { graph } = FlowGraph;
      const history  = graph;

      const canUndo = ref(history.canUndo());
      const canRedo = ref(history.canRedo());
      const copy = () => {
        const { graph } = FlowGraph;
        const cells = graph.getSelectedCells();
        if (cells.length) {
          graph.copy(cells);
        }
        return false;
      };

      const cut = () => {
        const { graph } = FlowGraph;
        const cells = graph.getSelectedCells();
        if (cells.length) {
          graph.cut(cells);
        }
        return false;
      };

      const paste = () => {
        const { graph } = FlowGraph;
        if (!graph.isClipboardEmpty()) {
          const cells = graph.paste({ offset: 32 });
          graph.cleanSelection();
          graph.select(cells);
        }
        return false;
      };

      const handleClick = (event: Event) => {
        const { graph } = FlowGraph;
        const name = (event.currentTarget as any).name!;
        switch (name) {
          case 'undo':
            graph.history.undo();
            break;
          case 'redo':
            graph.history.redo();
            break;
          case 'delete':
            graph.clearCells();
            break;
          case 'savePNG':
            graph.toPNG(
                (dataUri) => {
                  // 将dataUri转换为Blob对象
                  const byteString = atob(dataUri.split(',')[1]);
                  const mimeString = dataUri.split(',')[0].split(':')[1].split(';')[0];
                  const ab = new ArrayBuffer(byteString.length);
                  const ia = new Uint8Array(ab);
                  for (let i = 0; i < byteString.length; i++) {
                    ia[i] = byteString.charCodeAt(i);
                  }
                  const blob = new Blob([ab], { type: mimeString });

                  // 创建FormData对象
                  const formData = new FormData();
                  const url = decodeURI(window.location.href);
                  const cs_arr = url.split('?')[1]; //?后面的
                  const id = cs_arr.split('=')[1].split('&')[0];
                  const type = cs_arr.split('=')[2].split('&')[0];
                  const task = cs_arr.split('=')[3];
                  formData.append('image', blob, 'chartx.png');
                  formData.append('id', id); // 替换为实际的id
                  formData.append('type', task); // 替换为实际的type
                  formData.append('task', type); // 替换为实际的task

                  // 打印FormData的内容
                  for (let [key, value] of formData.entries()) {
                    console.log(key, value);
                  }

                  // 使用axios发送POST请求到服务器
                  saveCanvasPNG(formData).then(res => {
                    console.log('Success:', res.data);
                  });
                },
                {
                  backgroundColor: 'white',
                  padding: {
                    top: 20,
                    right: 30,
                    bottom: 40,
                    left: 50,
                  },
                  quality: 1,
                },
            );

            break;

          case 'saveSVG':
            graph.toSVG((dataUri: string) => {
              // 下载
              DataUri.downloadDataUri(DataUri.svgToDataUrl(dataUri), 'chart.svg');
            });
            break;
          case 'print':
            graph.printPreview();
            break;
          case 'copy':
            copy();
            break;
          case 'cut':
            cut();
            break;
          case 'paste':
            paste();
            break;
          case 'save':
            var graphData = graph.toJSON()
            var url = decodeURI(window.location.href);
            var cs_arr = url.split('?')[1];//?后面的
            console.info("url",url);
            var iid = cs_arr.split('=')[1].split('&')[0];
            var cellsToSend = {
              id: iid,
              cells: graphData.cells // 假设 graphData 里有一个 cells 属性
            };
            console.log("测试");
            console.log(cellsToSend);
            saveCanvas(cellsToSend).then(res=>{
              console.log(res);
            })
            break;
          default:
            break;
        }
      };

      history.on('change', () => {
        canUndo.value = history.canUndo();
        canRedo.value = history.canRedo();
      });
      graph.bindKey('ctrl+z', () => {
        if (history.canUndo()) {
          history.undo();
        }
        return false;
      });
      graph.bindKey('ctrl+shift+z', () => {
        if (history.canRedo()) {
          history.redo();
        }
        return false;
      });
      graph.bindKey('ctrl+d', () => {
        graph.clearCells();
        return false;
      });
      graph.bindKey('ctrl+s', () => {
        graph.toPNG((datauri: string) => {
          DataUri.downloadDataUri(datauri, 'chart.png');
        });
        return false;
      });
      graph.bindKey('ctrl+p', () => {
        graph.printPreview();
        return false;
      });
      graph.bindKey('ctrl+c', copy);
      graph.bindKey('ctrl+v', paste);
      graph.bindKey('ctrl+x', cut);

      return {
        canUndo,
        canRedo,
        copy,
        cut,
        paste,
        handleClick,
      };
    },
  });
</script>

<style lang="less" scoped>
.bar{
  text-align: left;
}
  .item-space {
    margin-top: 0px;
    margin-left: 15px;
    margin-bottom: 25px;
  }

  .run{
    width: 70px;
    margin-top: 0px;
    margin-left: 20px;
    margin-bottom: 25px;
    background-color: #5cadff;
    color: #fff;
  }
</style>
