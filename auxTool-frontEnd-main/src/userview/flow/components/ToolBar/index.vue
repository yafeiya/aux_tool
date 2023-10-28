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
        @on-ok="runorder"
        @on-cancel="cancel">
        <p>您已点击运行按钮，是否确认运行</p>

    </Modal>

  </div>
</template>

<script lang="ts">
  import { defineComponent, ref } from 'vue'; // ref, reactive
  import FlowGraph from '../../graph';
  import { DataUri } from '@antv/x6';
  import { EndUrl } from '../../../../../url_config'

  import axios from 'axios';
  import {saveCanvas, runCanvas, getDesignsById} from '../../../../api/api.js'
  import qs from "qs";
  import ViewUIPlus from 'view-ui-plus'
  export default defineComponent({
    name: 'Index',
    components: {},
    data() {
      return {
        modal: false,
        nowTime : new Date(),
        example:{
          "example_name": "",
          "rank": "",
          "state": "运行中",
          "cpu_num": 1,
          "gpu_num": 10,
          "post_date": (new Date()).getTime(),
          "post_time": (new Date()).getTime(),
          "run_time": "",
          "dataset_url": "",
          "model_name": "",
          "model_type": "",
          "epoch_num": "",
          "loss": "",
          "optimizer":"",
          "decay": "",
          "evalution": "",
          "model_url": "",
          "memory": "2000MB",
          "start_time":(new Date()).getTime() ,
          "end_time": '' ,
      }
      }
    },
    methods: {
      runorder() {
        const { graph } = FlowGraph;
        var graphData = graph.toJSON();

        var url = decodeURI(window.location.href);
        var cs_arr = url.split('?')[1];//?后面的
        var iid = cs_arr.split('=')[1].split('&')[0];
        var task = cs_arr.split('=')[2].split('&')[0];
        var type = cs_arr.split('=')[3].split('&')[0];
        console.log("task",task)
        console.log("type",type)
        let data = {
          dataset_url:EndUrl().fileUrl+"/"+type+"/"+task+"/"+iid,
          start_time:(new Date()).getTime(),
          id:iid,
          design_name:'1',
          rank:'1',
          cell:JSON.stringify(graphData.cells)
        }
        getDesignsById((data)).then(response => {
          data.design_name = response.data.data.design.Dataset_name
          data.rank=response.data.data.design.Rank
        })
        console.log("data",data);
        setTimeout(()=>{
          let rundata=data
          console.log("ddddddddddddddddata",rundata);
          runCanvas(rundata).then(response => {
          })
        },10);
        this.$Message.success('运行成功')
      },
      isRun() {
        this.modal = true
      },
      toDesign() {
        this.$router.push('/design')
      },
      info () {
          this.$Message.info('画布已保存');
      },
      // runErrorInfo({example}) {
      //   console.info(example)
      //   example.post_date = (new Date()).getTime()
      //   example.post_time = (new Date()).getTime()
      //   example.start_time = (new Date()).getTime()
      //   example.end_time = ''
      //   var postUrl = "http://localhost:3000/example"
      //   axios.post(postUrl,example).then(res=>{
      //     this.$Message["success"]({
      //         background: true,
      //         content: '运行成功，请在方案实例页面查看详情'
      //     });
      //   })
      // },
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
              (dataUri: string) => {
                // 下载
                DataUri.downloadDataUri(dataUri, 'chartx.png');
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
            var graphData = graph.toJSON();
            var url = decodeURI(window.location.href);
            var cs_arr = url.split('?')[1];//?后面的
            console.info("url",url);
            var iid = cs_arr.split('=')[1].split('&')[0];
            let data = {
              cell:JSON.stringify(graphData.cells),
              id:iid
            }
            console.log("ddddddddd",qs.stringify(data));

            saveCanvas((data)).then(response => {
              console.log(response);
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
