<template>
  <Card :bordered="false" style="width:100%;height: 232px"  v-if="this.cardInfo!=undefined">
    <!-- {{ console.info(this.cardInfo.dataset_name) + " " + console.info(this.viewRange)}} -->
    <template #title>
      <div style="text-align: center">
        <a @click="modal_intro1 = true">
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
    <div style="text-align:center;margin-top: -8%">
      <!--预览图 目前图片来自网页-->
      <img v-if="this.pageKind=='database'" src="../assets/dataset.png">
      <img v-if="this.pageKind=='modelbase'" src="../assets/model.png">
      <img v-if="this.pageKind=='defineFunction'" src="../assets/function.png">
      <img v-if="this.pageKind=='design'" src="../assets/design.png">
      <Tooltip content="点击编辑" placement="top">
        <div v-if="this.pageKind == 'design'">
          <a @click="toindex">
            <h4 style="margin-top: -5%;color: #054079">
              <Icon type="ios-create" style="font-size: 20px"/>
              {{this.cardName}}
            </h4>
          </a>
        </div>
        <div v-else>
          <a @click="editBtn = true">
            <h4 style="margin-top: 8%;color: #054079">
              <Icon type="ios-create" style="font-size: 20px"/>
              {{this.cardName}}
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
import {MenuGroup} from "view-ui-plus";
import parentMenu from "@/components/parentmenu.vue";
import mainTable from "@/components/maintable.vue";
export default {
  data() {
    return {
      // jsonBaseUrl: "http://localhost:3000",
      modal_intro1:false,//卡片2的介绍、发布、删除与预览
      modal_delete1:false,
      release1:false,
      preview1:false,
      editBtn:false,
      editFormItem: {},
      cardName: null,
      ruleValidate: {
        // dataset_name: [
        //   { required: true, message: 'The name cannot be empty', trigger: 'blur' }
        // ],
      },
    }
  },
  props:['cardInfo','viewRange', 'nowItem','taskType','addFormItem', 'cardList'],
  inject:['reload' ,'updataPage','jsonBaseUrl', 'pageKind','addFormItemCfg','getPageContent','cardNameFlag'],
  components: {
    dynamicInput,
  },
  created() {
    this.cardName = this.cardInfo[this.cardNameFlag]
    console.info("cardname: " + this.cardName + " " + this.cardNameFlag )
    this.initVaild()
  },

  methods: {
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
        console.info("vaild: "+i)
        if(i == "released") {
          continue;
        }
        var validItem = [{ required: true, message: '该项必填', trigger: 'blur' }]
        this.ruleValidate[i] = validItem
      }
      console.info(this.ruleValidate)
    },
    editCard (name) {
      //提示取消信息
      this.$refs[name].validate((valid) => {
            if (valid) {
              console.info(this.cardInfo)
              for(var i in this.addFormItemCfg) {
                var item = this.addFormItemCfg[i]
                console.info(item)
                if(item.isEditOnly == true) {
                  console.info("the only is: " + item.title)
                  for(var i in this.cardList) {
                    if(this.cardInfo.id !== this.cardList[i].id && this.cardInfo[item.name] == this.cardList[i][item.name]) {
                      this.$Message["error"]({
                        background: true,
                        content: item.title + "与其他项重复，请仔细检查"
                      });
                      this.getPageContent()
                      // this.updataPage("editRepeat")
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
      // console.info(findUrl + " contentcard")
      axios.put(findUrl, this.cardInfo).then(response => {
        this.updataPage("delete")
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
  watch:{
    // cardInfo(newCardList, oldCardList) {
    //     console.info(newCardList)
    //     console.info(oldCardList)
    // }
  },
}

</script>
