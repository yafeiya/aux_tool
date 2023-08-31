<template>
  <Layout>
    <Header style="background: #d1e4ff;height: 250px;">
      <img src="../assets/home页背景图.png" style="height: 100%;position:relative;right:-50px;width: 100%;"/>
      <div style="margin-top: -240px;position:relative;margin-left: 10%">
        <p style="font-size: 40px;color: #4d85ea;">
          <Icon type="md-cog" class="ivu-anim-loop" size="40" />
          数据驱动辅助工具
        </p>
        <p style="font-size: 20px;color: #000000;margin-left: 40px;line-height: 50px">
          ·基于Vue.js 3的与iviews组件设计的前端交互解决方案<br/>
          ·通用数据驱动建模可视化辅助工具
        </p>
      </div>
    </Header>
    <Content class="layout-content">
      <p style="font-size: 40px;color: #515a6e;text-align: center;">
        <Icon type="ios-arrow-dropdown" />
        快速开始
        <Icon type="ios-arrow-dropdown" />
      </p>
      <div class="home-card" style="text-align:center">
        <Grid :col="5" :hover=true>
          <GridItem>
            <Card :bordered="false" style="width: 100%;height: 100%" dis-hover>
              <div>
                <p style="font-size: 120px;">
                  <Icon type="ios-list-box-outline" />
                </p>
                <div style="font-size: 20px;margin-top: -10%;height: 80px;">
                  <Button type="primary" shape="circle" size="large" style="width: 80%" @click="toDatabase" ghost>数据集</Button>
                </div>
              </div>
            </Card>
          </GridItem>
          <GridItem>
            <Card :bordered="false" style="width: 100%;height: 100%" dis-hover>
              <div>
                <p style="font-size: 120px;">
                  <Icon type="logo-buffer" />
                </p>
                <div style="font-size: 20px;margin-top: -10%;height: 80px;">
                  <Button type="primary" shape="circle" size="large" style="width: 80%" @click="toModelbase" ghost>模型库</Button>
                </div>
              </div>
            </Card>
          </GridItem>
          <GridItem>
            <Card :bordered="false" style="width: 100%;height: 100%" dis-hover>
              <div>
                <p style="font-size: 120px;">
                  <Icon type="ios-hammer-outline" />
                </p>
                <div style="font-size: 20px;margin-top: -10%;height: 80px;">
                  <Button type="primary" shape="circle" size="large" style="width: 80%" @click="toDefineFunction" ghost>自定义函数</Button>
                </div>
              </div>
            </Card>
          </GridItem>
          <GridItem>
            <Card :bordered="false" style="width: 100%;height: 100%" dis-hover>
              <div>
                <p style="font-size: 120px;">
                  <Icon type="ios-git-network" />
                </p>
                <div style="font-size: 20px;margin-top: -10%;height: 80px;">
                  <Button type="primary" shape="circle" size="large" style="width: 80%" @click="toDesign" ghost>方案设计</Button>
                </div>
              </div>
            </Card>
          </GridItem>
          <GridItem>
            <Card :bordered="false" style="width: 100%;height: 100%" dis-hover>
              <div>
                <p style="font-size: 120px;">
                  <Icon type="ios-laptop" />
                </p>
                <div style="font-size: 20px;margin-top: -10%;height: 80px;">
                  <Button type="primary" shape="circle" size="large" style="width: 80%" @click="toExample" ghost>方案实例</Button>
                </div>
              </div>
            </Card>
          </GridItem>
        </Grid>
      </div>
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

.layout-base > .ivu-layout {
  margin-bottom: 20px;
}
.layout-content{
  background: white;
  min-height: 120px;
  line-height: 120px;
  height: 120px;
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
  width:1200px;
  /* height: 350px; */
  margin: auto;
  margin-top: 0%;
  margin-bottom: 7%;
  /* padding-bottom: 0px; */
}
</style>
