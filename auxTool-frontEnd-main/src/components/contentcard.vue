<template>
  <Card :bordered="false" style="width:100%;height: 232px"  v-if="this.cardInfo!=undefined">
    <!-- {{ console.info(this.cardInfo.dataset_name) + " " + console.info(this.viewRange)}} -->
    <template #title>
      <div style="text-align: center">
        <a @click="modal_intro1 = true" v-if="this.pageKind=='database'">
          <Icon type="ios-help-circle-outline" />管理
          <Modal
              v-model="modal_intro1"
              width="1000"
              style="margin-top: -50px">
            <p style="margin-top: 1%;font-size: 20px;text-align: center;" >
              <Space :size="15">
                <Input style="width: 500px" search enter-button="Search" placeholder="搜索数据集..." @click="handleSearch" />
                <Button  type="warning" icon="md-power" shape="circle" v-width=90 style="margin-left: 0%" @click="inputDatabase">导入</Button>
                <Button  type="success" icon="md-play"  shape="circle" v-width=90 style="margin-left: 1%" @click="updateToState('运行')">导出</Button>
                <Button  type="error" icon="md-pause"  shape="circle" v-width=90 style="margin-left: 1%" @click="updateToState('挂起')">删除</Button>
              </Space>
            </p>
            <Modal
                v-model="input"
                title="导入数据集"
                @on-ok="inputok"
                @on-cancel="cancel">

              <Form ref="inputFormItem" :model="inputFormItem" :label-width="80" :rules="ruleValidate">
                <FormItem label="名称" prop="name">
                  <Input v-model="inputFormItem.name" placeholder="自动获取"></Input>
                </FormItem>
                <FormItem label="级别">
                  <Select v-model="inputFormItem.level" style="width:300px" >
                    <Option v-for="item in inputFormItem.level_list" :value="item.value" :key="item.value">{{ item.label }}</Option>
                  </Select>
                </FormItem>
                <FormItem label="类型" prop="type">
                  <Input v-model="inputFormItem.type" placeholder="自动获取" disabled="True"></Input>
                </FormItem>
                <FormItem label="任务" prop="task">
                  <Input v-model="inputFormItem.task" placeholder="自动获取" disabled="True"></Input>
                </FormItem>
                <FormItem label="表头">
                  <RadioGroup v-model="biaotou">
                    <Radio label="有">
                      <span>有</span>
                    </Radio>
                    <Radio label="无">
                      <span>无</span>
                    </Radio>
                  </RadioGroup>
                </FormItem>

                <Upload  multiple :action="this.uploadUrl" style="margin-left: 80px" method="POST">
                  <Button icon="ios-cloud-upload-outline" style="margin-right: 5px">上传数据集</Button> 支持CSV格式，一次性上传多个文件
                </Upload>

              </Form>

            </Modal>
            <Table border  :columns="columns" :data="tabledata" style="width: auto">
              <template #preview="{row, index }">
                <Button type="primary" v-width=90 style="margin-left: -5px" @click="show(index)">下载</Button>
              </template>
            </Table>

          </Modal>
        </a>
        <a @click="modal_intro1 = true" v-else-if="this.pageKind=='modelbase'">
          <Icon type="ios-help-circle-outline" />管理
          <Modal
              v-model="modal_intro1"
              width="500"
              title="上传与下载"
              style="margin-top: -50px">
            <Space :size="15">
              <Upload  multiple :action="this.uploadUrl" style="margin-left: 80px" method="POST">
                  <Button icon="ios-cloud-upload-outline" style="margin-left: 40px">点击上传</Button>
              </Upload>
                  <Button icon="ios-cloud-download-outline" style="margin-left: 10px;margin-bottom: 8px" @click="downloadModel">点击下载</Button>
            </Space>

          </Modal>
        </a>
        <a @click="toindex" v-else-if="this.pageKind=='design'">
          <Icon type="ios-help-circle-outline" />设计
        </a>
        <a @click="modal_intro1 = true" v-else>
          <Icon type="ios-help-circle-outline" />简介
          <Modal
              v-model="modal_intro1"
              title="简介">
            <p style="white-space: pre-wrap" v-if="this.cancelInfo.description !== null">{{ this.cardInfo.description }}</p>
          </Modal>
        </a>
        <!--删除按钮-->
        <a @click="modal_delete1 = true">
          <Icon type="ios-trash-outline" />删除
          <Modal
              v-model="modal_delete1"
              title="删除"
              @on-ok="deleteCard"
              @on-cancel="cancelInfo">
            <p style="white-space: pre-wrap">{{'        是否确定删除该数据集？'}}</p>
          </Modal>
        </a>
        <!--发布按钮-->
        <a @click="release1 = true" v-if="viewRange=='private'">
          <Icon type="ios-help-circle-outline" />发布
          <Modal
              v-model="release1"
              title="发布"
              @on-ok="uploadCard"
              @on-cancel="cancelInfo">
            <p style="white-space: pre-wrap">{{'        确定发布到公共数据集？'}}</p>
          </Modal>
        </a>


      </div>
    </template>
    <div style="text-align:center">
      <!--预览图 目前图片来自网页-->
      <img v-if="this.pageKind=='database'" src="../assets/dataset.png" style="position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%);">
      <img v-if="this.pageKind=='modelbase'" src="../assets/model.png" style="position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%);">
      <img v-if="this.pageKind=='defineFunction'" src="../assets/function.png" style="position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%);">
      <img v-if="this.pageKind=='design'" src="../assets/design.png" style="position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%);">
      <Tooltip content="点击编辑" placement="bottom">
<!--        <div v-if="this.pageKind == 'design'">-->
<!--          <a @click="toindex">-->
<!--            <h4 style="margin-top: 130px;color: #054079;height: 10%">-->
<!--              <Icon type="ios-create" style="font-size: 20px"/>-->
<!--              {{this.cardName}}-->
<!--            </h4>-->
<!--          </a>-->
<!--        </div>-->
        <div>
          <a @click="editBtn = true">
            <h4 style="margin-top: 130px;color: #054079;height: 10%">
              <Icon type="ios-create" style="font-size: 20px"/>
              {{this.cardInfo[cardNameFlag]}}
              <Modal v-model="editBtn" title="编辑" :styles="{top: '30px'}"
                     @on-ok="editCard('cardInfo')" @on-cancel="cancelInfo">
                <Form ref="cardInfo" :model="cardInfo" :label-width="80" :rules="ruleValidate">
                  <FormItem v-for="(item, index) in addFormItemCfg" :label="item.title" :prop="item.name">
                    <div v-if="item.itemType == 'input'">
                      <div v-if="item.default">
                        <Input v-model="cardInfo[item.name]" placeholder="请输入任务..." disabled="True"></Input>
                      </div>
                      <div v-else>
                        <Input v-model="cardInfo[item.name]" :placeholder=item.others[0]></Input>
                      </div>
                    </div>
                    <div v-if="item.itemType == 'bigInput'">
                      <Input v-model="cardInfo[item.name]" :placeholder=item.others[0]
                             type="textarea" :autosize="{minRows: 2,maxRows: 5}">
                      </Input>
                    </div>
                    <div v-if="item.itemType == 'select'">
                      <Select v-model="cardInfo[item.name]">
                        <Option v-for="(option, index2) in item.others" :key="index2" :value="option.value">
                          {{ option.text }}
                        </Option>
                      </Select>
                    </div>
                    <div v-if="item.itemType == 'radio'">
                      <RadioGroup v-model="cardInfo[item.name]">
                        <Radio v-for="(option, index2) in item.others" :label="option.label"></Radio>
                        <!-- <Radio label="无"></Radio> -->
                      </RadioGroup>
                    </div>
                    <div v-if="item.itemType == 'dynamicInput'">
                      <dynamicInput :params-form="cardInfo.params"></dynamicInput>
                    </div>
                  </FormItem>
                  <!--                <Upload multiple type="drag" action=" ">-->
                  <!--                  <div style="padding: 5px 0">-->
                  <!--                    <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>-->
                  <!--                    <p>点击或拖拽上传</p>-->
                  <!--                  </div>-->
                  <!--                </Upload>-->
                </Form>
              </Modal>
            </h4>
          </a>
        </div>
      </Tooltip>

    </div>
  </Card>
</template>
<script>
import axios from 'axios';
import dynamicInput from "@/components/dynamicinput.vue";
import {FormItem, MenuGroup} from "view-ui-plus";
import parentMenu from "@/components/parentmenu.vue";
import mainTable from "@/components/maintable.vue";
import lineChart from "@/components/chart/line.vue";
import { EndUrl } from '../../url_config'
export default {
  data() {
    return {
      uploadUrl:EndUrl().backEndUrl+'/upload',
      input: false,
      inputFormItem: {
        name: '',
        type: '',
        level:'',
        task:'',
        biaotou:'',
        level_list:[
          {
            value: "等级1",
            label:"等级1"
          },
          {
            value: "等级2",
            label:"等级2"
          },
          {
            value: "等级3",
            label:"等级3"
          }
        ],
      },
      columns: [
        {
          type: 'selection',
          width: 60,
          align: 'center'
        },
        {
          title: '数据名',
          key: 'name',
          align: 'center'
        },
        {
          title: '级别',
          key: 'level',
          align: 'center',
          sortable: true
        },
        {
          title: '所属任务',
          key: 'task',
          align: 'center'
        },
        {
          title: '导入时间',
          key: 'input_time',
          align: 'center',
          sortable: true
        },
        {
          title: '统计信息',
          children:[
            {
              title: '表头数量',
              key: 'header_num',
              align: 'center',
              sortable: true
            },
            {
              title: '数据长度',
              key: 'data_len',
              align: 'center',
              sortable: true
            },
            {
              title: '数据类型',
              key: 'data_type',
              align: 'center',
            },
          ],
          align: 'center'
        },
        {
          title: '预览',
          slot: 'preview',
          align: 'center',
        },
      ],
      tabledata: [
        {
          name: '动作表',
          level: 1,
          task: "目标检测",
          input_time: "2023-9-1",
          header_num: 10,
          data_len: 10,
          data_type: "int",
        },
        {
          name: '状态表',
          level: 1,
          task: "目标检测",
          input_time: "2023-9-2",
          header_num: 12,
          data_len: 13,
          data_type: "char",
        },
        {
          name: '结果表',
          level: 1,
          task: "目标检测",
          input_time: "2023-9-3",
          header_num: 15,
          data_len: 16,
          data_type: "string",
        },
      ],
      searchdata:[],
      // jsonBaseUrl: "http://localhost:3000",
      modal_intro1:false,//卡片2的介绍、发布、删除与预览
      modal_delete1:false,
      release1:false,
      preview1:false,
      editBtn:false,
      editFormItem: {},
      // cardName: null,
      ruleValidate: {
      },
    }
  },
  props:['cardInfo','viewRange', 'nowItem','taskType','addFormItem', 'cardList'],
  inject:['reload' ,'updataPage','jsonBaseUrl', 'pageKind','addFormItemCfg','getPageContent','cardNameFlag'],
  components: {
    FormItem,
    lineChart,
    dynamicInput,
  },

  created() {
    console.info("contentcard--carddecripyion: " +  this.cardInfo.description)
    this.cardName = this.cardInfo[this.cardNameFlag]
    console.info("contentcard--cardname: " + this.cardName)
    this.initVaild()
  },
  // updated() {
  //   this.cardName = this.cardInfo[this.cardNameFlag]
  // },
  methods: {
    handleSearch(){
      this.searchdata = this.tabledata.filter(item => {
        return item.label.includes(this.searchText)
      })
    },
    inputDatabase(){
        this.input=true
        this.inputFormItem.type=this.taskType
        this.inputFormItem.task=this.nowItem
    },
    inputok(){
      console.info("1111111111111")
    },

    //跳转画布设计
    toindex() {
      console.info()
      this.$router.push('/index?id=' + this.cardInfo.id )
    },
    cancelInfo () {
      //提示取消信息
      this.getPageContent()
      this.$Message.info('操作取消');
    },
    initVaild() {
      for(var i in this.cardInfo) {
        // console.info("vaild: "+i)
        if(i == "released") {
          continue;
        }
        var validItem = [{ required: true, message: '该项必填', trigger: 'blur' }]
        this.ruleValidate[i] = validItem
      }
      // console.info(this.ruleValidate)
    },
    editCard (name) {
      //提示取消信息
      var tmpData = this.cardInfo
      this.$refs[name].validate((valid) => {
            if (valid) {
              console.info(this.cardInfo)
              for(var i in this.addFormItemCfg) {
                var item = this.addFormItemCfg[i]
                console.info(item)
                if(item.isEditOnly == true) {
                  console.info("the only is: " + item.title)
                  for(var j in this.cardList) {
                    if(this.cardInfo.id !== this.cardList[j].id && this.cardInfo[item.name] == this.cardList[j][item.name]) {

                      this.$Message["error"]({
                        background: true,
                        content: item.title + "与其他项重复，请仔细检查"
                      });

                      // this.cardInfo = tmpData

                      this.getPageContent()
                      // this.updataPage("editRepeat")
                      // console.info("over")
                      return;
                    }
                  }
                }
              }
              var putUrl = this.jsonBaseUrl + "/" + this.pageKind + "/" + this.cardInfo.id
              axios.put(putUrl, this.cardInfo).then(res=>{
                this.cardName = this.cardInfo[this.cardNameFlag]
                this.updataPage("edit")
              })
              // this.$Message.info('编辑成功');
            }
            else {
              this.$Message.error('编辑失败，请检查必填项！');
            }
      })

    },
    deleteCard() {
      if(this.viewRange == "private") {
        if(this.cardInfo.released === "10") {
          this.cardInfo.released = "00"
        } else if(this.cardInfo.released === "11") {
          this.cardInfo.released = "01"
        }
      } else {
        if(this.cardInfo.released === "01") {
          this.cardInfo.released = "00"
        } else if(this.cardInfo.released === "11") {
          this.cardInfo.released = "10"
        }
      }
      var findUrl = this.jsonBaseUrl + "/" + this.pageKind + "/" + this.cardInfo.id
      console.info(findUrl + " contentcard")
      console.info(this.cardInfo)
      axios.put(findUrl, this.cardInfo).then(response => {
        // this.cardName = this.cardInfo[this.cardNameFlag]
        this.updataPage("delete")
        console.info(response.data)
      })
    },
    uploadCard() {
      if(this.viewRange == "private") {
        if(this.cardInfo.released === "11") {
          this.$Message["error"]({
            background: true,
            content: "该项已发布，请仔细检查后操作"
          });
        } else if(this.cardInfo.released === "10") {
          var findUrl = this.jsonBaseUrl + "/" + this.pageKind + "/" + this.cardInfo.id
          // console.info(findUrl + " contentcard")
          this.cardInfo.released = "11"
          axios.put(findUrl, this.cardInfo).then(response => {
            this.updataPage("upload")
          })
        }
      }
    },
  },


}

</script>
