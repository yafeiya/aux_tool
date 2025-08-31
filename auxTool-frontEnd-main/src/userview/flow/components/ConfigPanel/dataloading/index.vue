<template>
  <Tabs class="config" value="1">
    <TabPane label="属性" name="1">
      <Row align="middle" v-if="globalGridAttr.selflabel === '仿真监听数据'">
        <Upload
          :action="EndUrl().backEndUrl + '/uploadLoss'"
          style="margin-bottom: 10px"
          method="POST"
          accept=".csv"
          :data="lossData"
          :show-upload-list="false"
          :on-success="uploadSuccess"
        >
          <Button
            class="view"
            icon="ios-cloud-upload-outline"
            @click="uploadfile"
            >选择网络损失数据</Button
          >
        </Upload>
        <Upload
          :action="EndUrl().backEndUrl + '/uploadReward'"
          style="margin-bottom: 10px"
          method="POST"
          accept=".txt"
          :data="lossData"
          :show-upload-list="false"
          :on-success="uploadSuccess"
        >
          <Button
            class="view"
            icon="ios-cloud-upload-outline"
            @click="uploadfile"
            >选择仿真奖励数据</Button
          >
        </Upload>
        <Upload
          :action="EndUrl().backEndUrl + '/uploadActions'"
          style="margin-bottom: 10px"
          method="POST"
          accept=".json"
          :data="lossData"
          :show-upload-list="false"
          :on-success="uploadSuccess"
        >
          <Button
            class="view"
            icon="ios-cloud-upload-outline"
            @click="uploadfile"
            >选择训练动作数据</Button
          >
        </Upload>
      </Row>

      <Row align="middle" v-if="globalGridAttr.selflabel === '模型日志数据'">
        <Upload action="//localhost:3000/">
          <Button class="view" icon="ios-cloud-upload-outline" @click=""
            >选择模型日志数据</Button
          >
        </Upload>
      </Row>

      <div v-for="children1 in hasChildren(menu)">
        <div v-for="children2 in children1.children">
          <div v-for="childrenitem in children2.children">
            <Row
              align="middle"
              v-if="
                globalGridAttr.selflabel === childrenitem.title &&
                globalGridAttr.selflabel != '仿真监听数据' &&
                globalGridAttr.selflabel != '模型日志数据'
              "
            >
              <Col :span="10">指定{{ childrenitem.title }}</Col>
              <Col :span="12">
                <Select
                  style="width: 100%"
                  v-model="data.nodeLabelname"
                  @on-change="onLabelChange"
                >
                  <Option
                    v-for="option in childrenitem.content"
                    :value="option.Dataset_name"
                  >
                    {{ option.Dataset_name }}
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

      <div
        v-if="
          globalGridAttr.selflabel != '仿真监听数据' &&
          globalGridAttr.selflabel != '模型日志数据'
        "
      >
        <Row class="params" align="middle">
          <Col :span="8">数据路径</Col>
          <Col :span="14">
            <Input
              v-model="data.nodedataurl"
              style="width: 100%"
              @change="ondataurlChange"
            />
          </Col>
        </Row>
        <Row class="params" align="middle">
          <Col :span="8">数据类型</Col>
          <Col :span="14">
            <Input
              v-model="data.nodedatatype"
              style="width: 100%"
              @change="ondatatypeChange"
            />
          </Col>
        </Row>
        <Row class="params" align="middle">
          <Col :span="8">密级</Col>
          <Col :span="14">
            <Input
              v-model="data.nodesecuritylevel"
              style="width: 100%"
              @change="onsecuritylevelChange"
            />
          </Col>
        </Row>
      </div>
    </TabPane>
    <TabPane label="节点" name="2">
      <Row align="middle">
        <Col :span="8">边框颜色</Col>
        <Col :span="14">
          <Input
            type="color"
            v-model="globalGridAttr.nodeStroke"
            style="width: 100%"
            @change="onStrokeChange"
          />
        </Col>
      </Row>
      <Row align="middle">
        <Col :span="8">边框宽度</Col>
        <Col :span="12">
          <Slider
            :min="1"
            :max="5"
            :step="1"
            v-model="data.nodeStrokeWidth"
            @on-input="onStrokeWidthChange"
          />
        </Col>
        <Col :span="2">
          <div class="result">{{ globalGridAttr.nodeStrokeWidth }}</div>
        </Col>
      </Row>
      <Row align="middle">
        <Col :span="8">填充颜色</Col>
        <Col :span="14">
          <Input
            type="color"
            v-model="data.nodeFill"
            style="width: 100%"
            @change="onFillChange"
          />
        </Col>
      </Row>
    </TabPane>
    <TabPane label="文本" name="3">
      <Row align="middle">
        <Col :span="8">字体大小</Col>
        <Col :span="12">
          <Slider
            v-model="data.nodeFontSize"
            :min="8"
            :max="16"
            :step="1"
            @on-input="onFontSizeChange"
          />
        </Col>
        <Col :span="2">
          <div class="result">{{ globalGridAttr.nodeFontSize }}</div>
        </Col>
      </Row>
      <Row align="middle">
        <Col :span="8">字体颜色</Col>
        <Col :span="14">
          <Input
            type="color"
            value="globalGridAttr.nodeColor"
            style="width: 100%"
            @change="onColorChange"
          />
        </Col>
      </Row>
    </TabPane>
  </Tabs>
</template>

<script lang="ts">
import { defineComponent, inject, watch, reactive, ref } from "vue";
import { Cell } from "@antv/x6/lib";
import { nodeOpt } from "./method";
import { EndUrl } from "../../../../../../url_config";
import { getCard, getMenuInfo } from "../../../../../api/api.js";

export default defineComponent({
  name: "Index",
  data() {
    return {
      lossData: {
        id: "",
        type: "",
        task: "",
      },
    };
  },
  methods: {
    EndUrl,
    uploadfile() {
      var url = decodeURI(window.location.href);
      var cs_arr = url.split("?")[1]; //?后面的
      var id = cs_arr.split("=")[1].split("&")[0];
      var task = cs_arr.split("=")[2].split("&")[0];
      var type = cs_arr.split("=")[3];
      (this.lossData.id = id),
        (this.lossData.task = task),
        (this.lossData.type = type);
      //console.info("iiiiiiiiiiiiiiiiiiiiii",this.lossData)
    },
    uploadSuccess() {
      this.$Message.success("上传成功");
    },
  },

  async setup() {
    const globalGridAttr: any = inject("globalGridAttr");
    const id: any = inject("id");
    let curCel: Cell;
    let writepath: any = inject("datapath");

    const changewritepath = () => {
      writepath.value = !writepath.value;
      // console.log(writepath.value)
    };

    const data = reactive({
      nodedataurl: "",
      nodedatatype: "",
      nodesecuritylevel: "",
      nodeStrokeWidth: "",
      nodeFill: "",
      nodeFontSize: "",
      nodeLabelname: "",
      nodeselflabel: "",
    });
    watch(
      [() => id.value],
      () => {
        console.log("=== 节点选择监听触发 ===");
        console.log("新选择的节点ID:", id.value);

        curCel = nodeOpt(id, globalGridAttr);
        console.log("节点操作完成，当前节点:", curCel);
        console.log("globalGridAttr更新后:", globalGridAttr);

        data.nodedataurl = globalGridAttr.nodedataurl;
        data.nodedatatype = globalGridAttr.nodedatatype || "";
        data.nodesecuritylevel = globalGridAttr.nodesecuritylevel || "";
        data.nodeStrokeWidth = globalGridAttr.nodeStrokeWidth;
        data.nodeFill = globalGridAttr.nodeFill;
        data.nodeFontSize = globalGridAttr.nodeFontSize;
        data.nodeLabelname = globalGridAttr.nodename;
        data.nodeselflabel = globalGridAttr.selflabel;

        console.log("数据同步完成，当前data状态:", {
          nodedataurl: data.nodedataurl,
          nodedatatype: data.nodedatatype,
          nodesecuritylevel: data.nodesecuritylevel,
          nodeLabelname: data.nodeLabelname,
          nodeselflabel: data.nodeselflabel,
        });

        console.log("节点当前数据:", curCel?.getData());
        console.log("=== 节点选择监听结束 ===");

        // curCel?.attr('body/stroke', 'red');
      },
      {
        immediate: false,
        deep: false,
      }
    );

    let databasecards: any = ref([]);
    // let databaselist:any  = ref([])
    let menu: any = [
      {
        name: "database",
        title: "数据准备",
        children: [],
      },
    ];

    let database = await getMenuInfo("database");
    menu[0]["children"] = database.data;

    // console.log("显示初始菜单",await menu)
    let i = 0;
    while (menu[0]["children"][i] != null) {
      // console.log(i)
      // console.log(menu[0]["children"][i].title)
      let j = 0;
      while (menu[0]["children"][i]["children"][j] != null) {
        // console.log(menu[0]["children"][i]["children"][j].title)

        let tip = {
          pageKind: "database",
          task: menu[0]["children"][i]["children"][j].title,
          Type: menu[0]["children"][i].title,
        };

        databasecards = await getCard(tip);
        if (databasecards.data.data != null) {
          let cardList = await databasecards.data.data.database;
          // console.log("读取数据库",await cardList)
          menu[0]["children"][i]["children"][j].content = cardList;
        }
        j++;
      }
      i++;
      // console.log("显示最终菜单",await menu)
    }

    // console.log("显示最终菜单",await menu)

    const onStrokeChange = (e: any) => {
      const val = e.target.value;
      globalGridAttr.nodeStroke = val;
      curCel?.attr("body/stroke", val);
    };

    const onStrokeWidthChange = (val: number) => {
      // console.log(val)
      globalGridAttr.nodeStrokeWidth = val;
      curCel?.attr("body/strokeWidth", val);
    };

    const onFillChange = (e: any) => {
      const val = e.target.value;
      globalGridAttr.nodeFill = val;
      curCel?.attr("body/fill", val);
    };

    const onFontSizeChange = (val: number) => {
      globalGridAttr.nodeFontSize = val;
      curCel?.attr("text/fontSize", val);
    };

    const onColorChange = (e: any) => {
      const val = e.target.value;
      globalGridAttr.nodeColor = val;
      curCel?.attr("text/fill", val);
    };

    const ondataurlChange = (e: any) => {
      const val = e.target.value;
      curCel?.setData({
        dataurl: val,
      });
    };
    const ondatatypeChange = (e: any) => {
      const val = e.target.value;
      globalGridAttr.nodedatatype = val;
      curCel?.setData({
        datatype: val,
      });
    };

    const onsecuritylevelChange = (e: any) => {
      const val = e.target.value;
      globalGridAttr.nodesecuritylevel = val;
      curCel?.setData({
        securitylevel: val,
      });
    };

    const onLabelChange = async (value) => {
      const val = value;

      console.log("=== 数据集选择开始 ===");
      console.log("用户选择的数据集:", val);
      console.log("当前节点ID:", id.value);
      console.log("当前节点对象:", curCel);

      globalGridAttr.nodename = val;
      curCel?.setData({
        name: val,
      });

      console.log("节点名称已保存到globalGridAttr和curCel");

      // 使用现有getCard接口查询数据库详细信息
      try {
        console.log("开始查询数据库详情，数据集名称:", val);

        // 从已加载的menu数据中查找对应的数据库记录
        let foundDatabaseItem = null;
        let foundInMenu = false;

        // 遍历menu数据查找匹配的数据集
        for (let menuCategory of menu) {
          if (menuCategory.children) {
            for (let taskCategory of menuCategory.children) {
              if (taskCategory.children) {
                for (let item of taskCategory.children) {
                  if (item.content) {
                    for (let dbItem of item.content) {
                      if (dbItem.Dataset_name === val) {
                        foundDatabaseItem = dbItem;
                        foundInMenu = true;
                        console.log("在菜单数据中找到匹配记录:", dbItem);
                        break;
                      }
                    }
                  }
                  if (foundInMenu) break;
                }
              }
              if (foundInMenu) break;
            }
          }
          if (foundInMenu) break;
        }

        if (foundDatabaseItem) {
          console.log("获取到的数据库记录:", foundDatabaseItem);

          // 类型断言，告诉TypeScript这是数据库记录对象
          const dbItem = foundDatabaseItem as any;

          // 自动填充数据路径、数据类型、密级
          if (dbItem.Data_path) {
            console.log("填充数据路径:", dbItem.Data_path);
            data.nodedataurl = dbItem.Data_path;
            globalGridAttr.nodedataurl = dbItem.Data_path;
            curCel?.setData({ dataurl: dbItem.Data_path });
          } else {
            console.log("数据库记录中没有Data_path字段");
          }

          if (dbItem.Character_type) {
            console.log("填充数据类型:", dbItem.Character_type);
            data.nodedatatype = dbItem.Character_type;
            globalGridAttr.nodedatatype = dbItem.Character_type;
            curCel?.setData({ datatype: dbItem.Character_type });
          } else {
            console.log("数据库记录中没有Character_type字段");
          }

          if (dbItem.Rank) {
            console.log("填充密级:", dbItem.Rank);
            data.nodesecuritylevel = dbItem.Rank;
            globalGridAttr.nodesecuritylevel = dbItem.Rank;
            curCel?.setData({ securitylevel: dbItem.Rank });
          } else {
            console.log("数据库记录中没有Rank字段");
          }

          console.log("数据填充完成后的data对象:", {
            nodedataurl: data.nodedataurl,
            nodedatatype: data.nodedatatype,
            nodesecuritylevel: data.nodesecuritylevel,
          });

          console.log("数据填充完成后的globalGridAttr:", {
            nodedataurl: globalGridAttr.nodedataurl,
            nodedatatype: globalGridAttr.nodedatatype,
            nodesecuritylevel: globalGridAttr.nodesecuritylevel,
          });

          console.log("数据填充完成后的节点数据:", curCel?.getData());
          console.log("✅ 数据库详情已自动填充成功");
        } else {
          console.log("❌ 在菜单数据中未找到匹配的数据集:", val);
          console.log("当前menu结构:", menu);
        }
      } catch (error) {
        console.error("❌ 查询数据库详情失败:", error);
        console.error("错误详情:", {
          message: error.message,
          stack: error.stack,
        });
      }

      console.log("=== 数据集选择结束 ===");

      // curCel?.attr('text/text', val);
      // console.log(curCel.data.name)
    };

    const hasChildren = (thelist) => {
      return thelist.filter((item) => item.children);
    };

    const noChildren = (thelist) => {
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
      ondatatypeChange,
      onsecuritylevelChange,
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
.config {
  width: 300px;
  text-align: center;
}
.params {
  margin: 10px;
}
.view {
  width: 300px;
  background-color: #fff;
  margin-top: 0px;
}
</style>
