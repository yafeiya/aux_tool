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
      <img src="../assets/dataset.png">
      <Tooltip content="点击编辑" placement="top">
        <a @click="edit = true">
          <h4 style="margin-top: -5%;color: #054079">
            <Icon type="ios-create" style="font-size: 20px"/>
            {{this.cardInfo.dataset_name}}
            <Modal v-model="edit" title="修改数据集" :styles="{top: '30px'}"
                   @on-ok="editsucess" @on-cancel="cancelInfo">

            </Modal>
          </h4>
        </a>
      </Tooltip>

    </div>
  </Card>
</template>
<script>
import axios from 'axios';
export default {
  data() {
    return {
      // jsonBaseUrl: "http://localhost:3000",
      modal_intro1:false,//卡片2的介绍、发布、删除与预览
      modal_delete1:false,
      release1:false,
      preview1:false,
      edit:false,
    }
  },
  props:['cardInfo','viewRange', 'nowItem','taskType','addFormItem'],
  inject:['reload' ,'updataPage','jsonBaseUrl', 'pageKind','addFormItemCfg','pageKind'],
  methods: {
    cancelInfo () {
      //提示取消信息
      this.$Message.info('操作取消');
    },
    editsucess () {
      //提示取消信息
      this.$Message.info('编辑成功');
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
