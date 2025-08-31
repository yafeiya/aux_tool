<template>
  <div class="bar">
    <Tooltip content="返回" placement="bottom">
      <template #title> </template>
      <Button name="backHome" @click="toDesign" class="item-space" size="small">
        返回
      </Button>
    </Tooltip>
    <Tooltip content="清除 (Cmd + D)" placement="bottom">
      <template #title> </template>
      <Button
        name="delete"
        @click="handleClick"
        class="item-space"
        size="small"
      >
        清除
      </Button>
    </Tooltip>
    <Tooltip content="撤销 (Cmd + Z)" placement="bottom">
      <template #title> </template>
      <Button
        :disabled="canUndo"
        name="undo"
        @click="handleClick"
        class="item-space"
        size="small"
      >
        撤销
      </Button>
    </Tooltip>

    <Tooltip content="复制 (Cmd + C)" placement="bottom">
      <template #title> </template>
      <Button name="copy" @click="handleClick" class="item-space" size="small">
        复制
      </Button>
    </Tooltip>

    <Tooltip content="剪切 (Cmd + X)" placement="bottom">
      <template #title> </template>
      <Button name="cut" @click="handleClick" class="item-space" size="small">
        剪切
      </Button>
    </Tooltip>

    <Tooltip content="粘贴 (Cmd + V)" placement="bottom">
      <template #title> </template>
      <Button name="paste" @click="handleClick" class="item-space" size="small">
        粘贴
      </Button>
    </Tooltip>

    <Tooltip content="保存PNG (Cmd + S)" placement="bottom">
      <template #title> </template>
      <Button
        name="savePNG"
        @click="handleClick"
        class="item-space"
        size="small"
      >
        保存PNG
      </Button>
    </Tooltip>

    <Tooltip content="模型保存" placement="bottom">
      <template #title> </template>
      <Button
        name="modelSave"
        @click="showSaveModal"
        class="model-save"
        size="small"
      >
        模型保存
      </Button>
    </Tooltip>

    <!-- //！！！！！！！！！！！！！在这儿修改！！！！！！！！！！ -->
    <!-- 智能建模助手按钮 -->
    <Tooltip content="打开智能建模助手" placement="bottom">
      <template #title> </template>
      <Button
        name="aiAssistant"
        @click="openAiAssistant"
        class="ai-assistant"
        size="small"
      >
        智能建模助手
      </Button>
    </Tooltip>

    <!-- 开始训练按钮 -->
    <Tooltip content="开始训练模型" placement="bottom">
      <template #title> </template>
      <Button name="run" @click="isRun" class="run" size="small">
        开始训练
      </Button>
    </Tooltip>
    <Modal
      v-model="modal"
      title="即将开始训练"
      @on-ok="RunExample({ mustItem, example })"
      @on-cancel="cancel"
    >
      <p>您已点击开始训练按钮，是否确认开始训练</p>
    </Modal>

    <!-- 模型保存选择弹窗 -->
    <Modal
      v-model="saveModal"
      title="选择保存方式"
      :footer-hide="true"
      width="400"
    >
      <div style="text-align: center; padding: 20px">
        <Button
          type="primary"
          size="large"
          @click="saveCanvas"
          style="margin-right: 20px; width: 120px"
        >
          画布保存
        </Button>
        <Button
          type="success"
          size="large"
          @click="exportFlowChart"
          style="width: 120px"
        >
          流程图导出
        </Button>
      </div>
    </Modal>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue"; // ref, reactive
import { runCanvas, saveCanvas, saveCanvasPNG } from "../../../../api/api.js";
import FlowGraph from "../../graph";
import { DataUri } from "@antv/x6";
import axios from "axios";
import qs from "qs";
import ViewUIPlus from "view-ui-plus";
import { EndUrl } from "../../../../../url_config";
export default defineComponent({
  name: "Index",
  components: {},
  data() {
    return {
      modal: false,
      saveModal: false,
      nowTime: new Date(),
      // 检查必须拖动到画布上的模块
      mustItem: [
        {
          name: "数据准备",
          flag: false,
        },
        {
          name: "模型选择",
          flag: false,
        },
        {
          name: "模型训练",
          flag: false,
        },
        {
          name: "仿真交互",
          flag: false,
        },
      ],
      example: {
        id: "",
        example_name: "",
        Rank: "",
        State: "运行中",
        Task: "",
        Type: "",
        Dataset_name: "Dataset_name",
        Model_name: "string",
        Train_state: "1",
        Metics: "string",
        Network_num: 1,
        Learning_rate: 0.01,
        Act_function: "RELu",
        Radom_seed: 2,
        Optimizer: "string",
        Batch_size: 64,
        Epoch_num: 10,
        Explore_rate: 0.1,
        Decay_factor: 0.1,
        start_time: new Date().getTime(),
        end_time: "",
      },
    };
  },
  methods: {
    isRun() {
      this.modal = true;
    },

    RunExample({ mustItem, example }) {
      var url = decodeURI(window.location.href);
      var cs_arr = url.split("?")[1]; //?后面的
      var iid = cs_arr.split("=")[1].split("&")[0];
      this.example.id = iid;
      example.start_time = new Date().getTime();
      example.end_time = "";
      var cs_arr = url.split("?")[1]; //?后面的
      var id = cs_arr.split("=")[1].split("&")[0];
      var task = cs_arr.split("=")[2].split("&")[0];
      var type = cs_arr.split("=")[3];
      example.dataset_url =
        EndUrl().fileUrl + "/" + type + "/" + task + "/" + id;
      example.id = id;
      const { graph } = FlowGraph;
      var graphData = graph.toJSON();
      console.info("graphData", graphData);
      const DatasetResult = graphData.cells
        .filter(
          (cell) =>
            cell.data &&
            cell.data.grandLabel &&
            cell.data.grandLabel === "数据准备" &&
            cell.data.fatherLabel.endsWith("数据集")
        )
        .map((cell) => cell.data);
      var Dataset =
        DatasetResult.length > 0 ? DatasetResult : "没有找到符合条件的数据";
      example.Dataset_name = Dataset[0].name;
      const ModelResult = graphData.cells
        .filter(
          (cell) =>
            cell.data &&
            cell.data.grandLabel &&
            cell.data.grandLabel === "模型选择"
        )
        .map((cell) => cell.data);
      var Model =
        ModelResult.length > 0 ? ModelResult : "没有找到符合条件的数据";
      example.Model_name = Model[0].name;
      example.Act_function = Model[0].Act_function;
      example.Decay_factor = Model[0].Decay_factor;
      example.Epoch_num = Model[0].Epoch_num;
      example.Explore_rate = Model[0].Explore_rate;
      example.Optimizer = Model[0].Optimizer;
      example.Radom_seed = Model[0].Radom_seed;
      example.batch = Model[0].batch;
      example.Network_num = Model[0].Network_num;
      example.learning_rate = Model[0].learning_rate;
      console.info("Model", Model);
      const TrainStateResult = graphData.cells
        .filter(
          (cell) =>
            cell.data &&
            cell.data.grandLabel &&
            cell.data.grandLabel === "训练监控"
        )
        .map((cell) => cell.data);
      var TrainState =
        TrainStateResult.length > 0
          ? TrainStateResult
          : "没有找到符合条件的数据";
      example.Train_state = TrainState[0].selflabel;
      console.info("example", example);
      runCanvas(example).then((res) => {
        this.$Message["success"]({
          background: true,
          content: "训练已开始，请在实例管理页面查看详情",
        });

        //以下代码用于存png
        const { graph } = FlowGraph;
        var graphData = graph.toJSON();
        var url = decodeURI(window.location.href);
        var cs_arr = url.split("?")[1]; //?后面的
        console.info("url", url);
        var iid = cs_arr.split("=")[1].split("&")[0];
        var cellsToSend = {
          id: iid,
          cells: graphData.cells, // 假设 graphData 里有一个 cells 属性
        };
        graph.toPNG(
          (dataUri) => {
            // 将dataUri转换为Blob对象
            const byteString = atob(dataUri.split(",")[1]);
            const mimeString = dataUri
              .split(",")[0]
              .split(":")[1]
              .split(";")[0];
            const ab = new ArrayBuffer(byteString.length);
            const ia = new Uint8Array(ab);
            for (let i = 0; i < byteString.length; i++) {
              ia[i] = byteString.charCodeAt(i);
            }
            const blob = new Blob([ab], { type: mimeString });

            // 创建FormData对象
            const formData = new FormData();
            const url = decodeURI(window.location.href);
            const cs_arr = url.split("?")[1]; //?后面的
            const id = cs_arr.split("=")[1].split("&")[0];
            const type = cs_arr.split("=")[2].split("&")[0];
            const task = cs_arr.split("=")[3];
            formData.append("image", blob, "chartx.png");
            formData.append("id", id); // 替换为实际的id
            formData.append("type", task); // 替换为实际的type
            formData.append("task", type); // 替换为实际的task

            // 打印FormData的内容
            for (let [key, value] of formData.entries()) {
              console.log(key, value);
            }

            // 使用axios发送POST请求到服务器
            saveCanvasPNG(formData).then((res) => {
              console.log("Success:", res.data);
            });
          },
          {
            backgroundColor: "white",
            padding: {
              top: 20,
              right: 30,
              bottom: 40,
              left: 50,
            },
            quality: 1,
          }
        );
        saveCanvas(cellsToSend).then((res) => {
          console.log(res);
        });

        for (var i in mustItem) {
          mustItem[i].flag = false;
        }
      });
    },
    cancel() {
      this.$Message.info("已取消");
    },
    // 返回到设计页面
    toDesign() {
      this.$router.push("/design");
    },
    // 打开智能建模助手
    openAiAssistant() {
      // 在新窗口中打开AI工具（4001端口）
      const aiToolUrl =
        window.location.protocol + "//" + window.location.hostname + ":4001/";
      window.open(
        aiToolUrl,
        "_blank"
        // "width=1200,height=800,scrollbars=yes,resizable=yes"
      );
      this.$Message.info("正在打开智能建模助手...");
    },
    // 显示保存选择弹窗
    showSaveModal() {
      this.saveModal = true;
    },
    // 画布保存（原有功能）
    saveCanvas() {
      this.saveModal = false;
      const { graph } = FlowGraph;
      var graphData = graph.toJSON();
      var url = decodeURI(window.location.href);
      var cs_arr = url.split("?")[1]; //?后面的
      console.info("url", url);
      var iid = cs_arr.split("=")[1].split("&")[0];
      var cellsToSend = {
        id: iid,
        cells: graphData.cells, // 假设 graphData 里有一个 cells 属性
      };
      console.log("保存画布");
      console.log(cellsToSend);
      saveCanvas(cellsToSend).then((res) => {
        console.log(res);
        this.$Message.success("画布已保存");
      });
    },
    // 导出流程图（合并Mermaid和配置）
    exportFlowChart() {
      this.saveModal = false;
      const { graph } = FlowGraph;
      var graphData = graph.toJSON();
      var url = decodeURI(window.location.href);
      var cs_arr = url.split("?")[1];
      var id = cs_arr.split("=")[1].split("&")[0];

      // 提取节点和连接线
      const nodes = graphData.cells.filter((cell) => cell.shape !== "edge");
      const edges = graphData.cells.filter((cell) => cell.shape === "edge");

      // 生成完整的导出内容
      let exportContent = "";

      // 1. 添加Mermaid流程图部分
      exportContent += "1. 画布mermaid流程图\n";
      exportContent += "graph TD\n";

      // 添加节点定义
      nodes.forEach((node, index) => {
        const nodeId = `node${index + 1}`;
        const nodeLabel =
          node.data?.selflabel || node.attrs?.text?.text || `节点${index + 1}`;

        // 添加节点，不使用样式
        exportContent += `    ${nodeId}["${nodeLabel}"]\n`;

        // 存储节点ID映射
        node.mermaidId = nodeId;
      });

      exportContent += "\n";

      // 添加连接线
      edges.forEach((edge) => {
        const sourceNode = nodes.find((node) => node.id === edge.source.cell);
        const targetNode = nodes.find((node) => node.id === edge.target.cell);

        if (sourceNode && targetNode) {
          exportContent += `    ${sourceNode.mermaidId} --> ${targetNode.mermaidId}\n`;
        }
      });

      exportContent += "\n";

      // 2. 添加流程图配置部分
      exportContent += "2. 流程图配置\n";
      exportContent += "节点属性配置：\n\n";

      // 遍历每个节点，导出其属性配置
      nodes.forEach((node, index) => {
        const nodeId = `node${index + 1}`;
        const nodeLabel =
          node.data?.selflabel || node.attrs?.text?.text || `节点${index + 1}`;

        exportContent += `【${nodeId} - ${nodeLabel}】\n`;

        // 导出节点的基本信息
        if (node.data) {
          if (node.data.grandLabel) {
            exportContent += `  类型: ${node.data.grandLabel}\n`;
          }
          if (node.data.fatherLabel) {
            exportContent += `  父级标签: ${node.data.fatherLabel}\n`;
          }
          if (node.data.selflabel) {
            exportContent += `  节点标签: ${node.data.selflabel}\n`;
          }

          // 导出数据相关属性
          if (node.data.dataurl) {
            exportContent += `  数据路径: ${node.data.dataurl}\n`;
          }
          if (node.data.datatype) {
            exportContent += `  数据类型: ${node.data.datatype}\n`;
          }
          if (node.data.securitylevel) {
            exportContent += `  密级: ${node.data.securitylevel}\n`;
          }

          // 导出模型相关属性
          if (node.data.learning_rate) {
            exportContent += `  学习率: ${node.data.learning_rate}\n`;
          }
          if (node.data.batch) {
            exportContent += `  批次大小: ${node.data.batch}\n`;
          }
          if (node.data.Epoch_num) {
            exportContent += `  训练轮数: ${node.data.Epoch_num}\n`;
          }
          if (node.data.Act_function) {
            exportContent += `  激活函数: ${node.data.Act_function}\n`;
          }
          if (node.data.Optimizer) {
            exportContent += `  优化器: ${node.data.Optimizer}\n`;
          }
          if (node.data.Network_num) {
            exportContent += `  网络层数: ${node.data.Network_num}\n`;
          }
          if (node.data.Radom_seed) {
            exportContent += `  随机种子: ${node.data.Radom_seed}\n`;
          }
          if (node.data.Explore_rate) {
            exportContent += `  探索率: ${node.data.Explore_rate}\n`;
          }
          if (node.data.Decay_factor) {
            exportContent += `  衰减因子: ${node.data.Decay_factor}\n`;
          }

          // 导出其他数据属性
          Object.keys(node.data).forEach((key) => {
            if (
              ![
                "grandLabel",
                "fatherLabel",
                "selflabel",
                "dataurl",
                "datatype",
                "securitylevel",
                "learning_rate",
                "batch",
                "Epoch_num",
                "Act_function",
                "Optimizer",
                "Network_num",
                "Radom_seed",
                "Explore_rate",
                "Decay_factor",
              ].includes(key)
            ) {
              if (
                node.data[key] !== undefined &&
                node.data[key] !== null &&
                node.data[key] !== ""
              ) {
                exportContent += `  ${key}: ${node.data[key]}\n`;
              }
            }
          });
        }

        exportContent += "\n";
      });

      // 导出连接关系配置
      if (edges.length > 0) {
        exportContent += "连接关系配置：\n";
        edges.forEach((edge, index) => {
          const sourceNode = nodes.find((node) => node.id === edge.source.cell);
          const targetNode = nodes.find((node) => node.id === edge.target.cell);

          if (sourceNode && targetNode) {
            const sourceName =
              sourceNode.data?.selflabel ||
              sourceNode.attrs?.text?.text ||
              "未知节点";
            const targetName =
              targetNode.data?.selflabel ||
              targetNode.attrs?.text?.text ||
              "未知节点";
            exportContent += `  连接${
              index + 1
            }: ${sourceName} → ${targetName}\n`;
          }
        });
      }

      // 下载为txt文件
      const blob = new Blob([exportContent], {
        type: "text/plain;charset=utf-8",
      });
      const url_download = URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = url_download;
      a.download = `flowchart_${id}_${new Date().getTime()}.txt`;
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
      URL.revokeObjectURL(url_download);

      this.$Message.success("流程图已导出");
    },
  },

  setup() {
    const { graph } = FlowGraph;
    const history = graph;

    const canUndo = ref(history.canUndo());

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
        case "undo":
          graph.history.undo();
          break;

        case "delete":
          graph.clearCells();
          break;
        case "savePNG":
          graph.toPNG(
            (dataUri) => {
              // 将dataUri转换为Blob对象
              const byteString = atob(dataUri.split(",")[1]);
              const mimeString = dataUri
                .split(",")[0]
                .split(":")[1]
                .split(";")[0];
              const ab = new ArrayBuffer(byteString.length);
              const ia = new Uint8Array(ab);
              for (let i = 0; i < byteString.length; i++) {
                ia[i] = byteString.charCodeAt(i);
              }
              const blob = new Blob([ab], { type: mimeString });

              // 创建FormData对象
              const formData = new FormData();
              const url = decodeURI(window.location.href);
              const cs_arr = url.split("?")[1]; //?后面的
              const id = cs_arr.split("=")[1].split("&")[0];
              const type = cs_arr.split("=")[2].split("&")[0];
              const task = cs_arr.split("=")[3];
              formData.append("image", blob, "chartx.png");
              formData.append("id", id); // 替换为实际的id
              formData.append("type", task); // 替换为实际的type
              formData.append("task", type); // 替换为实际的task

              // 打印FormData的内容
              for (let [key, value] of formData.entries()) {
                console.log(key, value);
              }

              // 使用axios发送POST请求到服务器
              saveCanvasPNG(formData).then((res) => {
                console.log("Success:", res.data);
              });
            },
            {
              backgroundColor: "white",
              padding: {
                top: 20,
                right: 30,
                bottom: 40,
                left: 50,
              },
              quality: 1,
            }
          );

          break;

        case "copy":
          copy();
          break;
        case "cut":
          cut();
          break;
        case "paste":
          paste();
          break;

        default:
          break;
      }
    };

    history.on("change", () => {
      canUndo.value = history.canUndo();
    });
    graph.bindKey("ctrl+z", () => {
      if (history.canUndo()) {
        history.undo();
      }
      return false;
    });

    graph.bindKey("ctrl+d", () => {
      graph.clearCells();
      return false;
    });
    graph.bindKey("ctrl+s", () => {
      graph.toPNG((datauri: string) => {
        DataUri.downloadDataUri(datauri, "chart.png");
      });
      return false;
    });

    graph.bindKey("ctrl+c", copy);
    graph.bindKey("ctrl+v", paste);
    graph.bindKey("ctrl+x", cut);

    return {
      canUndo,
      copy,
      cut,
      paste,
      handleClick,
    };
  },
});
</script>

<style lang="less" scoped>
.bar {
  text-align: left;
}
.item-space {
  margin-top: 0px;
  margin-left: 15px;
  margin-bottom: 25px;
}

.model-save {
  width: 100px;
  margin-top: 0px;
  margin-left: 15px;
  margin-bottom: 25px;
  background-color: #f56c6c;
  color: #fff;
  border: none;
}

.model-save:hover {
  background-color: #f78989;
  color: #fff;
}

.ai-assistant {
  width: 120px;
  margin-top: 0px;
  margin-left: 15px;
  margin-bottom: 25px;
  background-color: #67c23a;
  color: #fff;
  border: none;
}

.ai-assistant:hover {
  background-color: #85ce61;
  color: #fff;
}

.run {
  width: 90px;
  margin-top: 0px;
  margin-left: 15px;
  margin-bottom: 25px;
  background-color: #5cadff;
  color: #fff;
}
</style>
