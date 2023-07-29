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
.menu-table{
  margin-left: 18%;
  margin-top: 8%;
  font-size:20px;
  margin-bottom: -2%;
}
.layout-nav{
  width: 620px;
  margin: 0 auto;
  margin-right: -80px;
}
.dev-run-preview .dev-run-preview-edit{ display: none }
.main-table {
  margin-top: 2%;
}
</style>
<template>
  <div class="layout">
    <Layout>
      <Layout>
        <!--//固定侧边菜单-->
        <Sider :style="{position: 'fixed', height: '100vh', left: 0, overflow: 'auto'}">
          <!--//侧边导航标题-->
          <Affix v-color="'white'" class="menu-table">
            <Button :type="text" size="small" shape="circle" @click="go"><Icon type="md-arrow-back" /></Button> |  数据集
          </Affix>
          <Divider />
            <Menu active-name="1-1" theme="dark" width="auto" :open-names="['1']">
                <Submenu name="1">
                  <template #title>
                    <Icon type="ios-navigate"></Icon>
                    数值数据集
                  </template>
                  <MenuItem name="1-1">火力分配1</MenuItem>
                  <MenuItem name="1-2">火力分配2</MenuItem>
                </Submenu>
                <Submenu name="2">
                  <template #title>
                    <Icon type="ios-keypad"></Icon>
                    图像数据集
                  </template>
                  <MenuItem name="2-1">目标检测</MenuItem>
                  <MenuItem name="2-2">目标分类</MenuItem>
                </Submenu>
                <Submenu name="3">
                  <template #title>
                    <Icon type="ios-analytics"></Icon>
                    其他数据集
                  </template>
                  <MenuItem name="3-1">待补充</MenuItem>
                </Submenu>
            </Menu>
        </Sider>
        <Layout :style="{marginLeft: '200px'}">
          <!--//上边框菜单栏-->
          <Header>
            <Menu mode="horizontal" theme="dark" active-name="1">
              <div class="layout-home"><Button type="default" ghost>我是数据辅助工具logo</Button></div>
              <div class="layout-nav">
                <MenuItem name="1">
                  <Icon type="ios-navigate"></Icon>
                  数据集
                </MenuItem>
                <MenuItem name="2">
                  <Icon type="ios-keypad"></Icon>
                  模型库
                </MenuItem>
                <MenuItem name="3">
                  <Icon type="ios-analytics"></Icon>
                  自定义函数
                </MenuItem>
                <MenuItem name="4">
                  <Icon type="ios-paper"></Icon>
                  方案设计
                </MenuItem>
                <MenuItem name="5">
                  <Icon type="md-arrow-dropright-circle"></Icon>
                  方案实例
                </MenuItem>
              </div>
            </Menu>
          </Header>
          <Content :style="{padding: '0 16px 16px'}">
            <!--//数据集选项-->
            <Tabs type="card" class="main-table">
              <TabPane label="我的数据集">
                <!--//以栅格形式显示卡片-->
                <Row style="background:#eee;padding:20px" :gutter="16">
                  <!--第一张卡片-->
                  <Col span="6">
                    <Card :bordered="false" style="width:100%;height:100%">
                      <div style="text-align:center">
                        <!--添加数据集按钮-->
                        <Icon type="ios-add-circle-outline" @click="add = true" style="font-size: 50px;margin-top: 80px;margin-bottom: 5%;"/>
                          <Modal
                              v-model="add"
                              title="添加数据集"
                              :styles="{top: '30px'}"
                              @on-ok="ok"
                              @on-cancel="cancel">
                            <!--填写表单-->
                            <Form :model="formItem" :label-width="80">
                              <FormItem label="名称">
                                <Input v-model="formItem.name" placeholder="请输入数据集名称..."></Input>
                              </FormItem>
                              <FormItem label="任务">
                                <Input v-model="formItem.quest" placeholder="请输入任务..."></Input>
                              </FormItem>
                              <FormItem label="字符集">
                                <Input v-model="formItem.character" placeholder="请输入字符集..."></Input>
                              </FormItem>
                              <FormItem label="级别">
                                <Select v-model="formItem.level">
                                  <Option value="beijing">级别1</Option>
                                  <Option value="shanghai">级别2</Option>
                                  <Option value="shenzhen">级别3</Option>
                                </Select>
                              </FormItem>
                              <FormItem label="类型">
                                <Select v-model="formItem.types">
                                  <Option value="beijing">数值数据集</Option>
                                  <Option value="shanghai">图像数据集</Option>
                                  <Option value="shenzhen">文本数据集</Option>
                                  <Option value="shenzhen">其他</Option>
                                </Select>
                              </FormItem>
                              <FormItem label="有无表头">
                                <RadioGroup v-model="formItem.biaotou">
                                  <Radio label="有"></Radio>
                                  <Radio label="无"></Radio>
                                </RadioGroup>
                              </FormItem>
                              <FormItem label="描述">
                                <Input v-model="formItem.textarea" type="textarea" :autosize="{minRows: 2,maxRows: 5}" placeholder="相关描述......"></Input>
                              </FormItem>
                              <Upload multiple type="drag" action="  ">
                                <div style="padding: 5px 0">
                                  <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                                  <p>点击或拖拽上传</p>
                                </div>
                              </Upload>
                            </Form>
                          </Modal>
                        <h3>添加数据集</h3>
                      </div>
                    </Card>
                  </Col>
                  <!--第二张卡片-->
                  <Col span="6">
                    <Card :bordered="false" style="width:100%">
                      <template #title>
                        <Icon type="ios-film-outline"></Icon>
                        决策树算法
                      </template>
                      <template #extra>
                        <!--简介按钮-->
                        <a @click="modal_intro1 = true">
                          <Icon type="ios-help-circle-outline" />简介
                          <Modal
                              v-model="modal_intro1"
                              title="简介">
                            <p style="white-space: pre-wrap">{{'        这里是简介'}}</p>
                          </Modal>
                        </a>
                        <!--删除按钮-->
                        <a @click="modal_delete1 = true">
                          <Icon type="ios-trash-outline" />删除
                          <Modal
                              v-model="modal_delete1"
                              title="删除"
                              @on-ok="ok"
                              @on-cancel="cancel">
                            <p style="white-space: pre-wrap">{{'        是否确定删除该数据集？'}}</p>
                          </Modal>
                        </a>
                        <!--发布按钮-->
                        <a @click="release1 = true">
                          <Icon type="ios-help-circle-outline" />发布
                          <Modal
                              v-model="release1"
                              title="发布"
                              @on-ok="ok"
                              @on-cancel="cancel">
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
                  </Col>
                  <!--第三张卡片 内容和第二张卡片类似-->
                  <Col span="6">
                    <Card :bordered="false" style="width:100%">
                      <template #title>
                        <Icon type="ios-film-outline"></Icon>
                        KNN算法
                      </template>
                      <template #extra>
                        <a @click="modal_intro2 = true">
                          <Icon type="ios-help-circle-outline" />简介
                          <Modal
                              v-model="modal_intro2"
                              title="简介"
                              @on-ok="确定"
                              @on-cancel="取消">
                            <p>这里是简介</p>
                          </Modal>
                        </a>
                        <a @click="modal_delete2 = true">
                          <Icon type="ios-trash-outline" />删除
                          <Modal
                              v-model="modal_delete2"
                              title="删除"
                              @on-ok="ok"
                              @on-cancel="cancel">
                            <p>是否确定删除该数据集？</p>
                          </Modal>
                        </a>
                        <a @click="release2 = true">
                          <Icon type="ios-help-circle-outline" />发布
                          <Modal
                              v-model="release2"
                              title="发布"
                              @on-ok="ok"
                              @on-cancel="cancel">
                            <p style="white-space: pre-wrap">{{'        确定发布到公共数据集？'}}</p>
                          </Modal>
                        </a>
                      </template>
                      <div style="text-align:center">
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
                  </Col>
                </Row>
              </TabPane>
              <TabPane label="公共数据集">这里是公共数据集 参照“我的数据集”内容</TabPane>
            </Tabs>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  </div>
</template>
<script>
import {MenuGroup} from "view-ui-plus";

export default {
  components: {MenuGroup},
  methods: {
    go() {
      this.$router.push('/')
    },
    ok () {
      //提示成功信息
      this.$Message.info('操作成功');
    },
    cancel () {
      //提示取消信息
      this.$Message.info('操作取消');
    }
  },
  data () {
    return {
      border: true,
      hover: true,
      add:false,//添加数据集
      modal_intro1:false,//卡片2的介绍、发布、删除与预览
      modal_delete1:false,
      release1:false,
      preview1:false,
      modal_intro2:false,//卡片3的介绍、发布、删除 卡片3的预览采用卡片2的内容
      modal_delete2:false,
      release2:false,
      //添加数据集操作的表单内容
      formItem: {
        name: '',
        quest: '',
        character: '',
        level: '',
        types: '',
        biaotou: '无',
        textarea: ''
      },
      //预览内容-表格形式
      columns: [
        {
          title: 'Name',
          key: 'name'
        },
        {
          title: 'Age',
          key: 'age'
        },
        {
          title: 'Address',
          key: 'address'
        }
      ],
      data: [
        {
          name: 'John Brown',
          age: 18,
          address: 'New York No. 1 Lake Park',
          date: '2016-10-03'
        },
        {
          name: 'Jim Green',
          age: 24,
          address: 'London No. 1 Lake Park',
          date: '2016-10-01'
        }
      ]
    }
  }
}

</script>
