<template>
    <Card :bordered="false" style="width:100%">
        <template #title>
            <Icon type="ios-film-outline"></Icon>
            <!-- {{ this.cardInfo }} -->
            {{this.cardInfo.dataset_name}}
        </template>
        <template #extra>
        <!--简介按钮-->
        <a @click="modal_intro1 = true">
            <Icon type="ios-help-circle-outline" />简介
            <Modal
                v-model="modal_intro1"
                title="简介">
            <p style="white-space: pre-wrap">{{ this.cardInfo.description }}</p>
            </Modal>
        </a>
        <!--删除按钮-->
        <a @click="modal_delete1 = true">
            <Icon type="ios-trash-outline" />删除
            <Modal
                v-if="this.viewRange == 'private'"
                v-model="modal_delete1"
                title="删除"
                @on-ok="deleteCard"
                @on-cancel="cancelInfo">
            <p style="white-space: pre-wrap">{{'        是否确定删除该数据集？'}}</p>
            </Modal>
<!-- 
            <Modal
                v-else
                v-model="modal_delete1"
                title="删除"
                @on-ok="deleteCard()"
                @on-cancel="cancelInfo">
            <p style="white-space: pre-wrap">{{'        是否确定删除该数据集？'}}</p>
            </Modal> -->
        </a>
        <!--发布按钮-->
        <a @click="release1 = true">
            <Icon type="ios-help-circle-outline" />发布
            <Modal
                v-model="release1"
                title="发布"
                @on-ok="successInfo"
                @on-cancel="cancelInfo">
            <p style="white-space: pre-wrap">{{'        确定发布到公共数据集？'}}</p>
            </Modal>
        </a>
        </template>
        <div style="text-align:center">
        <!--预览图 目前图片来自网页-->
        <img src="https://dev-file.iviewui.com/stJXDnKhL5qEBD0dHSDDTKbdnptK6mV5/large">
        <Button type="default" shape="circle" icon="md-eye" @click="preview1 = true">点击预览</Button>
        <Modal
            v-model="preview1"
            title="数据集预览"
            @on-ok="确定"
            @on-cancel="取消">
            <Table :columns="columns" :data="data"></Table>
        </Modal>
        </div>
    </Card>
</template>
<script>
import axios from 'axios';
export default {
    data() {
        return {
            jsonBaseUrl: "http://localhost:3000",
            modal_intro1:false,//卡片2的介绍、发布、删除与预览
            modal_delete1:false,
            release1:false,
            preview1:false,
        }
    },
    props:['cardInfo','viewRange'],
    methods: {
        successInfo () {
        //提示成功信息
            this.$Message.info('操作成功');
        },
        cancelInfo () {
        //提示取消信息
            this.$Message.info('操作取消');
        },
        deleteCard() {
            if(this.viewRange == "private") {
                console.info(this.cardInfo.released[0])
                console.info(this.cardInfo.released[0])
                console.info(this.cardInfo.released[0])
                this.cardInfo.released = "01"
                console.info("111")
            } else {
                console.info("222")
                this.cardInfo.released[1] = '0'
                console.info("222")
            }
            var findUrl = this.jsonBaseUrl + "/" + this.pageKind + "?id=" + this.cardInfo.id
            axios.put(findUrl, this.cardInfo.released).then(response => {
                console.info(response)
            })
        }
    },
}

</script>