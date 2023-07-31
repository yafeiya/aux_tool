<template>
    <div v-if="viewRange === 'private'">
        <Row v-for="(i,index1) in this.myCardRowNum" :key="index1"  style="background:#eee;padding:20px" :gutter="16" >
            <Col v-for="(j, index2) in this.getColOfRow(i, this.myCardRowNum, this.myCardColNum, this.myCardNum)" 
            :key="index2" :span= "Math.ceil(24 / this.myCardColNum)">
                <div v-if="i==1&&j==1">
                    <addCard/>
                </div>
                <div v-else>
                    <contentCard :card-info="this.myCardList[ (i-1) * this.myCardColNum + j - 2 ]" :view-range="viewRange"/>
                </div>
            </Col>
        </Row>
    </div>
    <div v-else>
        <Row v-for="(i,index1) in this.publicCardRowNum" :key="index1"  style="background:#eee;padding:20px" :gutter="16" >
            <Col v-for="(j, index2) in this.getColOfRow(i, this.publicCardRowNum, this.publicCardColNum, this.publicCardNum)" 
            :key="index2" :span= "Math.ceil(24 / this.publicCardColNum)">
                <contentCard :card-info="this.publicCardList[ (i-1) * this.publicCardColNum + j - 1]" :view-range="viewRange"/>
            </Col>
            
        </Row>
    </div>
</template>
<script>
import axios from 'axios';
import addCard from './addcard.vue'
import contentCard from './contentcard.vue'
export default {
    data() {
        return {
            jsonBaseUrl: "http://localhost:3000",
            
            // 中间页面卡片列表与数量
            myCardList: [],
            publicCardList: [],
            // colnum应该为24的因数
            myCardNum: 0,
            myCardRowNum: 0,
            myCardColNum: 4,
            publicCardNum: 0,
            publicCardRowNum: 0,
            publicCardColNum: 4,
        }
    },
    // nowItem左侧选中条目，pageKind当前页面种类{'database'，'modelbase'，'definFunc'，'design'}，viewRange公共的还是私有的{'private','public'}
    props: ['nowItem','pageKind','viewRange'],
    components:{
        addCard,
        contentCard
    },
    mounted() {
        // this.getPageContent()
    },
    updated() {
        this.getPageContent()
    },
    methods: {
        // successInfo () {
        // //提示成功信息
        //     this.$Message.info('操作成功');
        // },
        // cancelInfo () {
        // //提示取消信息
        //     this.$Message.info('操作取消');
        // },
        getPageContent() {
            var findUrl = this.jsonBaseUrl + "/" + this.pageKind + "?task=" + this.nowItem
            console.info(findUrl)
            // this.myCardList = Object.keys(this.myCardList).map(key =>this.myCardList[key]);
            console.info(typeof(x))
            this.myCardList = []
            this.publicCardList = []
            axios.get(findUrl).then(response => {
                var cardList = response.data
                var length = cardList.length
                // this.myCardNum = 0;
                // this.publicCardNum = 0;
                for( var i = 0; i < length;i++) {
                    if(cardList[i].released[0] == '1') {
                        this.myCardList.push(cardList[i] );
                        // this.myCardNum = this.myCardNum + 1
                    }
                    if(cardList[i].released[1] == '1') {
                        this.publicCardList.push(cardList[i]);
                        // this.publicCardNum = this.publicCardNum + 1;
                    }
                }
                this.myCardNum = this.myCardList.length
                this.myCardRowNum = Math.ceil(this.myCardNum / this.myCardColNum)
                this.publicCardNum = this.publicCardList.length
                this.publicCardRowNum =  Math.ceil(this.publicCardNum / this.publicCardColNum)
                console.info(this.myCardNum)
                console.info(this.myCardRowNum)
                console.info(this.publicCardNum )
                console.info(this.publicCardRowNum )
            })
            // (async () => {
            //     this.cardList = (await axios.get(findUrl)).data
            // })()
            // console.log(this.cardList);
        },
        // 一列的个数不一定等于列数，因为有可能不够，所以得计算
        // (当前所处行得位置，总行数，总列数，总个数)
        getColOfRow(nowRow, rowNum, colNum, Num) {
            if(nowRow == rowNum) {
                return (Num - ((rowNum-1) * colNum))
            } else {
                return colNum
            }
        }
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