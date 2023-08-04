<template>
  <Card :bordered="false" style="width:100%;height:232px">
    <div style="text-align:center;top: auto">
      <!--添加数据集按钮-->
      <a @click="addDialog = true" style="position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%);">
        <Icon type="ios-add-circle-outline" style="font-size: 50px;margin-top: 45%;margin-bottom: 8%;"/>
        <h3 style="margin-top: 5%;margin-bottom: 18%;">添加数据集</h3>
      </a>

      <Modal v-model="addDialog" title="添加数据集" :styles="{top: '30px'}"
             @on-ok="creatCard" @on-cancel="cancelInfo">

        <Form :model="addFormItem" :label-width="80">
          <FormItem v-for="(item, index) in addFormItemCfg" :label="item.title">
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
          </FormItem>
          <Upload multiple type="drag" action=" ">
            <div style="padding: 5px 0">
              <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
              <p>点击或拖拽上传</p>
            </div>
          </Upload>
          <!-- <FormItem label="任务">
              <Input v-model="addFormItem.task" placeholder="请输入任务..." disabled="True"></Input>
          </FormItem>
          <FormItem label="字符集">
              <Input v-model="addFormItem.character_type" placeholder="请输入字符集..."></Input>
          </FormItem>
          <FormItem label="级别">
              <Select v-model="addFormItem.rank">
              <Option value="beijing">级别1</Option>
              <Option value="shanghai">级别2</Option>
              <Option value="shenzhen">级别3</Option>
              </Select>
          </FormItem>
          <FormItem label="类型">
              <Select v-model="addFormItem.type">
              <Option value="beijing">数值数据集</Option>
              <Option value="shanghai">图像数据集</Option>
              <Option value="shenzhen">文本数据集</Option>
              <Option value="shenzhen">其他</Option>
              </Select>
          </FormItem>
          <FormItem label="有无表头">
              <RadioGroup v-model="addFormItem.header">
              <Radio label="有"></Radio>
              <Radio label="无"></Radio>
              </RadioGroup>
          </FormItem>
          <FormItem label="描述">
              <Input v-model="addFormItem.description" type="textarea" :autosize="{minRows: 2,maxRows: 5}" placeholder="相关描述......"></Input>
          </FormItem>
          <Upload multiple type="drag" action="  ">
              <div style="padding: 5px 0">
              <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
              <p>点击或拖拽上传</p>
              </div>
          </Upload> -->
        </Form>
      </Modal>

    </div>
  </Card>
  <!-- <Button @click="getFormItem">1231</Button> -->
</template>


<script>
import axios from 'axios';



export default {
  data() {
    return {
      // jsonBaseUrl: "http://localhost:3000",
      addDialog:false,//添加数据集
      // addFormItem: {
      //     released: "00",
      //     data_path: "",
      // },
    }
  },
  inject:['addFormItemCfg', 'updataPage','pageKind', 'jsonBaseUrl'],
  props:['viewRange', 'nowItem', 'taskType', 'addFormItem'],
  // mounted() {
  //     this.getFormItem()
  // },
  methods: {

    creatCard () {
      //提示成功信息
      // console.info("createdCard: " + this.addFormItem.task)
      var findUrl = this.jsonBaseUrl + "/" + this.pageKind + "?task=" + this.nowItem + "&type=" + this.taskType + "&dataset_name=" + this.addFormItem.dataset_name
      console.info(findUrl)
      axios.get(findUrl).then(response=>{
        var myCardList = response.data
        console.info(myCardList.length == 0)
        var newCard = this.addFormItem
        if(myCardList.length == 0) {
          // console.info("10")
          newCard.released = "10";
          // newCard.dataset_name = this.addFormItem.name
          // newCard.type =
          console.info("10")
          axios.post(findUrl, newCard).then(response=>{
            console.info(response)
            this.updataPage("creat")
          })
        } else {
          // newCard = myCardList
          if(myCardList[0].released == "00") {
            newCard.released = "10"
            axios.post(findUrl, newCard).then(response=>{
              console.info(response)
              this.updataPage("creat")
            })
          } else {
            this.$Message["error"]({
              background: true,
              content: "该名称已存在，请仔细检查"
            });
          }
        }

        console.info(newCard)
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
