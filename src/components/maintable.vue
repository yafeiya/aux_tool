<template>
  <Tabs v-if="pageKind!=='defineFunction'" type="card" class="main-table">
    <TabPane :label = this.labelName1>
        <!-- {{ pageKind +" " + nowItem + " /my" }} -->
        <menuCard :view-range="'private'" :page-kind="this.pageKind" :now-item="this.nowItem"/>
    </TabPane>
    <TabPane :label = this.labelName2>
        <menuCard :view-range="'public'" :page-kind="this.pageKind" :now-item="this.nowItem"/>
    </TabPane>
  </Tabs>
  <Tabs v-else type="card" class="main-table">
    <TabPane :label = this.labelName1>
        {{ pageKind +" " + nowItem + " /my" }}
    </TabPane>
  </Tabs>
</template>
<script>
import menuCard from './menucard.vue';
export default {
    
    data() {
        return {
            // myUrl: null,
            // publicUrl: null
            labelName1: "我的",
            labelName2: "公共",
        }
    },
    
    props: ['pageKind', 'nowItem'],

    components: {
       menuCard
    },
    mounted() {
        if(this.pageKind === "database") {
            this.labelName1 = "我的数据集"
            this.labelName2 = "公共数据集"
        } else if(this.pageKind === "modelbase") {
            this.labelName1 = "我的模型库"
            this.labelName2 = "基础算法模型"
        } else if(this.pageKind === "design") {
            this.labelName1 = "我的方案"
            this.labelName2 = "公共方案库"
        } else {
            this.labelName1 = "自定义函数"
        }
    }
}
</script>