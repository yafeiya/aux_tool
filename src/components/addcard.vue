<template>
    <Card :bordered="false" style="width:100%;height:100%">
        <div style="text-align:center">
        <!--添加数据集按钮-->
        <a @click="addDialog = true">
            <Icon type="ios-add-circle-outline" style="font-size: 50px;margin-top: 80px;margin-bottom: 5%;"/>
            <h3>添加数据集</h3>
        </a>
        <Modal
        v-model="addDialog"
        title="添加数据集"
        :styles="{top: '30px'}"
        @on-ok="successInfo"
        @on-cancel="cancelInfo">
            <!--填写表单-->
            <Form :model="addFormItem" :label-width="80">
            <FormItem label="名称">
                <Input v-model="addFormItem.name" placeholder="请输入数据集名称..."></Input>
            </FormItem>
            <FormItem label="任务">
                <Input v-model="addFormItem.task" placeholder="请输入任务..."></Input>
            </FormItem>
            <FormItem label="字符集">
                <Input v-model="addFormItem.character" placeholder="请输入字符集..."></Input>
            </FormItem>
            <FormItem label="级别">
                <Select v-model="addFormItem.rank">
                <Option value="beijing">级别1</Option>
                <Option value="shanghai">级别2</Option>
                <Option value="shenzhen">级别3</Option>
                </Select>
            </FormItem>
            <FormItem label="类型">
                <Select v-model="addFormItem.types">
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
            </Upload>
            </Form>
        </Modal>
        </div>
    </Card>
</template>


<script>
import axios from 'axios';
export default {
    data() {
        return {
            addDialog:false,//添加数据集
            addFormItem: {
                name: '',
                task: '',
                character: '',
                rank: '',
                types: '',
                header: '无',
                textarea: ''
            },
        }
    },
   
    methods: {
        successInfo () {
        //提示成功信息
            this.$Message.info('操作成功');
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