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

    <Tooltip content="复制 (Cmd + Shift + Z)" placement="bottom">
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
      <Button name="run" @click="runorder" class="run" size="small"> 运行 </Button>
    </Tooltip>

  </div>
</template>

<script lang="ts">
  import { defineComponent, ref } from 'vue'; // ref, reactive
  import FlowGraph from '../../graph';
  import { DataUri } from '@antv/x6';
  import axios from 'axios';
  import ViewUIPlus from 'view-ui-plus'
  export default defineComponent({
    name: 'Index',
    components: {},

    methods: {
      toDesign() {
        this.$router.push('/design')
      },
      info () {
          this.$Message.info('画布已保存');
      },
    },

    setup() {
      const { graph } = FlowGraph;
      const history  = graph;

      const canUndo = ref(history.canUndo());
      const canRedo = ref(history.canRedo());

      //！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！在此修改运行函数
      const runorder =() =>{
          console.log("在此修改运行函数")
      }


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
            var graphData = graph.toJSON()
            var url = decodeURI(window.location.href);
            var cs_arr = url.split('?')[1];//?后面的
            var iid = cs_arr.split('=')[1];
            console.log("测试");
            console.log(graphData);
            axios.patch('http://localhost:3000/design/'+iid, {'cells':graphData.cells})
                .then(response => {
                  console.log(response);
                })
                .catch(error => {
                  console.log(error);
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
        runorder,
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
