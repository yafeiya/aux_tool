<template>
    <Layout>
        <Header class="layout-header">数据驱动辅助工具 </Header>
        <Content class="layout-content">
            <Card class="home-card">
                <div style="text-align:center">
                    <Row :gutter="16+3*8">
                        <Col span="8">
                            <transition name="ease">
                            <!-- <div v-show="anim2" class="transition-box">
                                <router-link v-show="anim2" class="transition-button" to="/database">Home</router-link>
                            </div> -->
                            <Button type="info" v-show="anim2" class="transition-button" @click="toDatabase">数据库</Button>
                            </transition>
                        </Col>
                        <Col span="8">
                            <transition name="ease">
                            <Button type="info" v-show="anim2" class="transition-button" @click="toModelbase">模型库</Button>
                            </transition>
                        </Col>
                        <Col span="8">
                            <transition name="ease">
                            <Button type="info" v-show="anim2" class="transition-button" @click="toDefineFunction">自定义函数</Button>
                            </transition>
                        </Col>
                    </Row>
                    <Row :gutter="16+2*8">
                        <Col span="12">
                            <transition name="ease">
                            <Button type="info" v-show="anim2" class="transition-button" @click="toDesign">方案设计</Button>
                            </transition>
                        </Col>
                        <Col span="12">
                            <transition name="ease">
                            <Button type="info" v-show="anim2" class="transition-button" @click="toExample">方案实列</Button>
                            </transition>
                        </Col>
                    </Row>
                </div>
            </Card>
        </Content>

    </Layout>
    
    
</template>
<script>
export default {
    data () {
        return {
            anim1: true,
            anim2: false,
            anim3: true,
            anim4: true,
            anim5: true,
            anim6: true,
            anim7: true,
            anim8: true,
            anim9: true,
            anim10: true,
            anim11: true
        }
    },
    mounted() {
        this.anim2 = true;
    },
    methods: {
        toDatabase(event) {
            this.$router.push('/database')
        },
        toModelbase(event) {
            this.$router.push('/modelbase')
        },
        toDesign(event) {
            this.$router.push('/design')
        },
        toDefineFunction(event) {
            this.$router.push('/defineFunction')
        },
        toExample(event) {
            this.$router.push('/example')
        },

    }
   
}
</script>
<style>
.transition-button {
    
    min-height: 84%;
    width: 100%;
   
    padding: 24px 0;
    margin-bottom: 10px;
    padding: 24px 0;
}
.transition-box {
    margin-top: 32px;
    width: 100%;
    border-radius: 4px;
    background-color: #409EFF;
    text-align: center;
    color: #fff;
    padding: 24px 0;
}
.layout-base{
    border: 1px solid #d7dde4;
    background: #f5f7f9;
    position: relative;
    border-radius: 4px;
    color: white;
    text-align: center;
    overflow: hidden;
}
.layout-base > .ivu-layout {
    margin-bottom: 20px;
}
.layout-con{
    position: relative;
    overflow: hidden;
    text-align: center;
}
.layout-single{
    height: 253px;
    margin-bottom: 48px;
    font-size: 14px;
    color: white;
}
.layout-con:last-child{
    margin: 0;
}
.layout-logo{
    width: 100px;
    height: 30px;
    background: #5b6270;
    border-radius: 3px;
    float: left;
    position: relative;
    top: 15px;
    left: 20px;
}
.layout-nav{
    width: 420px;
    margin: 0 auto;
        margin-right: 20px;
}
.layout-breadcrumb{
    padding: 10px 15px 0;
}
.layout-logo-left{
    width: 90%;
    height: 30px;
    background: #5b6270;
    border-radius: 3px;
    margin: 15px auto;
}
.layout-hide-text .layout-text{
    display: none;
}
.layout-header,
.layout-footer{
    background: #7CBCE9;
    color: white;
}
.layout-content{
    background: white;
    min-height: 120px;
    line-height: 120px;
}
.layout-sider{
    background: #3A9FE8;
    min-height: 120px;
    line-height: 120px;
}
.layout-footer-center{
    text-align: center;
}
.layout iframe{
    width: 100%;
    border: 0;
}
.layout-header-bar{
    background: #fff;
    box-shadow: 0 1px 1px rgba(0,0,0,.1);
}
.menu-icon{
    transition: all .3s;
}
.rotate-icon{
    transform: rotate(-90deg);
}
.menu-item span{
    display: inline-block;
    overflow: hidden;
    width: 69px;
    text-overflow: ellipsis;
    white-space: nowrap;
    vertical-align: bottom;
    transition: width .2s ease .2s;
}
.menu-item i{
    transform: translateX(0px);
    transition: font-size .2s ease, transform .2s ease;
    vertical-align: middle;
    font-size: 16px;
}
.collapsed-menu span{
    width: 0px;
    transition: width .2s ease;
}
.collapsed-menu i{
    transform: translateX(5px);
    transition: font-size .2s ease .2s, transform .2s ease .2s;
    vertical-align: middle;
    font-size: 22px;
}
.home-card {
    width:1000px;
    /* height: 350px; */
    margin: auto;
    margin-top: 7%;
    margin-bottom: 7%;
    /* padding-bottom: 0px; */
}
</style>