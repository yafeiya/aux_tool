<template>
  <Card :bordered="false" style="width:100%;height:232px">
    <div style="text-align:center;top: auto">
      <!--添加数据集按钮-->
      <a @click="addDialog = true" style="position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%);">
        <Icon type="ios-add-circle-outline" style="font-size: 50px;margin-top: 45%;margin-bottom: 8%;"/>
        <h3 style="margin-top: 5%;margin-bottom: 18%;"> {{this.addBtnName }}</h3>
      </a>

      <Modal v-model="addDialog" :title="this.addBtnName" :styles="{top: '30px'}"
             @on-ok="creatCard('addFormItem')" @on-cancel="cancelInfo">

        <Form ref="addFormItem" :model="addFormItem" :label-width="80" :rules="ruleValidate">
          <FormItem v-for="(item, index) in addFormItemCfg" :label="item.title" :prop="item.name">

            <div v-if="item.itemType == 'input'">
              <div v-if="item.default">
                <Input v-model="addFormItem[item.name]" placeholder="请输入任务..." disabled="True"></Input>
              </div>
              <div v-else>
                <Input v-model="addFormItem[item.name]" :placeholder=item.others[0]></Input>
              </div>
            </div>

            <div v-if="item.itemType == 'bigInput'">
              <Input v-model="addFormItem[item.name]" :placeholder=item.others[0]
                     type="textarea" :autosize="{minRows: 2,maxRows: 5}">
              </Input>
            </div>
            <div v-if="item.itemType == 'select'">
              <Select v-model="addFormItem[item.name]">
                <Option v-for="(option, index2) in item.others" :key="index2" :value="option.value">
                  {{ option.text }}
                </Option>
              </Select>
            </div>
            <div v-if="item.itemType == 'radio'">
              <RadioGroup v-model="addFormItem[item.name]">
                <Radio v-for="(option, index2) in item.others" :label="option.label"></Radio>
                <!-- <Radio label="无"></Radio> -->
              </RadioGroup>
            </div>

            <div v-if="item.itemType == 'dynamicInput'">
              <dynamicInput :params-form="item.others"
                            :add-form-item="addFormItem"/>
            </div>
          </FormItem>
        </Form>
      </Modal>

    </div>
  </Card>
  <!-- <Button @click="getFormItem">1231</Button> -->
</template>


<script>
import axios from 'axios';
import dynamicInput from './dynamicinput.vue'
import {updataCard, createCard, addDesignCard, addDefinefunctionCard} from "@/api/api";


export default {
  data() {
    return {
      // jsonBaseUrl: "http://localhost:3000",
      addBtnName: "添加",
      addDialog:false,//添加数据集
      // addFormItem: {
      //     released: "00",
      //     data_path: "",
      // },
      ruleValidate: {
        // dataset_name: [
        //   { required: true, message: 'The name cannot be empty', trigger: 'blur' }
        // ],
      },
    }
  },
  inject:['addFormItemCfg', 'updataPage','pageKind', 'jsonBaseUrl'],
  props:['viewRange', 'nowItem', 'taskType', 'addFormItem'],
  // mounted() {
  //     this.getFormItem()
  // },
  components: {
    dynamicInput,
  },
  created() {
    this.initVaild()
  },
  mounted() {
    var tmp = ''
    if(this.pageKind == "database") {
      tmp =  "数据集"
    } else if(this.pageKind == "defineFunction") {
      tmp = "函数"
    } else if(this.pageKind == "design") {
      tmp = "方案"
    } else if(this.pageKind == "modelbase") {
      tmp = "模型"
    }
    this.addBtnName = this.addBtnName + tmp
  },
  methods: {
    initVaild() {
      for(var i in this.addFormItem) {
        // console.info("vaild: "+i)
        if(i == "released") {
          continue;
        }
        var validItem = [{ required: true, message: '该项必填', trigger: 'blur' }]
        this.ruleValidate[i] = validItem
        if(this.pageKind=="defineFunction"){
          this.ruleValidate['Params']=[{}]
        }
      }
    },
    creatCard (name) {

      //提示成功信息
      // console.info("createdCard: " + this.addFormItem.task)
      // var form = this.addFormItem
      this.$refs[name].validate((valid) => {
        console.info("this.addFormItemthis.addFormItemthis.addFormItem",this.addFormItem)
        if (valid) {
          var data = {
            lan: this.addFormItem["Lan"],
            id: this.addFormItem["Id"],
            pageKind: this.pageKind,
            dataset_name: this.addFormItem["Dataset_name"],
            task: this.nowItem,
            type: this.taskType,
            rank: this.addFormItem["Rank"],
            character_type: this.addFormItem["Character_type"],
            header: this.addFormItem["Header"],
            description: this.addFormItem["Description"],
            code: this.addFormItem["Code"],
            released:"11",
            alias:this.addFormItem["Alias"],
            function_body:this.addFormItem["Function_body"],
            params:''
          }
          if(data.pageKind=="defineFunction"){
            for(var i in this.addFormItem["params"]){
              data.params+=this.addFormItem["params"][i].value[0]
              data.params+=','
              data.params+=this.addFormItem["params"][i].value[1]
              if(i==this.addFormItem["params"].length-1){
                break
              }
              else
                data.params+='|'
            }
          }
          console.info("dddddddddddddaat",data)
          if(data.pageKind=="database"||data.pageKind=="modelbase"){
            createCard(data).then(response => {
              if(response.data.msg=="fail"){
                this.$Message["error"]({
                  background: true,
                  content: "该名称已存在，请仔细检查"
                });
              }
              else{
                this.updataPage("creat")
                // console.info("11111111111",response)
              }
            })
          }
          //这是前两个页面的新建卡片
         else{
           if(data.pageKind=="design"){
             addDesignCard(data).then(response => {
               // console.info('22222222222222222222222data:',data)
               if(response.data.msg=="fail"){
                 this.$Message["error"]({
                   background: true,
                   content: "该名称已存在，请仔细检查"
                 });
               }
               else{
                 this.updataPage("creat")
                 // console.info("11111111111",response)
               }
             })
           }
           else{
             addDefinefunctionCard(data).then(response => {
               // console.info('22222222222222222222222data:',data)
               if(response.data.msg=="fail"){
                 this.$Message["error"]({
                   background: true,
                   content: "该名称已存在，请仔细检查"
                 });
               }
               else{
                 this.updataPage("creat")
                 // console.info("11111111111",response)
               }
             })
           }
          }
          //这是后两个页面的新建卡片

        } else {
          this.$Message.error('添加失败，请检查必填项！');
        }
      })


    },
    cancelInfo () {
      //提示取消信息
      this.$Message.info('操作取消');
    },
  },

}
</script>

<style>
.layout{
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
}
.layout-home{
  width: 100px;
  height: 30px;
  border-radius: 3px;
  float: left;
  position: relative;
  top: 3px;
  left: 0px;
}
.dev-run-preview .dev-run-preview-edit{ display: none }
.main-table {
  margin-top: 2%;
}
</style>
