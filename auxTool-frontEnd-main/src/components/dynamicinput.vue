<template>
  <Form ref="formDynamic" :model="formDynamic" :label-width="80" style="width: 300px">
    <template v-for="(item, index) in formDynamic.items">
      <FormItem
          v-if="item.status"
          :key="index"
          :prop="'items.' + index + '.value'">
        <Row style="width:450px">
          <Col span="6">
            <Input type="text" v-model="item.value[0]" placeholder="Enter something..."></Input>
          </Col>
          <Col span="6" offset="1">
            <Input type="text" v-model="item.value[1]" placeholder="Enter something..."></Input>
          </Col>

          <Col span="2" offset="1">

            <Button @click="handleAdd">添加</Button>
          </Col>
          <Col span="2" offset="2">
            <Button @click="handleRemove(index)">删除</Button>
            <!--            <Button @click="handleAdd">添加</Button>-->
          </Col>
        </Row>
      </FormItem>
    </template>

  </Form>
</template>
<script>
import {Button, Col} from "view-ui-plus";

export default {
  components: {Button, Col},
  data () {
    return {
      index: 1,
      formDynamic: {
        items:[
        ],
      }
    }
  },

  props: ['paramsForm','addFormItem'],
  mounted() {
    this.initFormDynamic()
  },
  beforeUpdate() {
    this.addFormItem['params']=this.formDynamic.items
    console.info("this.addFormItem['params']",this.addFormItem['params'])
  },
  methods: {
    initFormDynamic() {
      this.index = 1;
      for(var i in this.paramsForm) {
        var paramsItem = {value:[], index: i, status: 1}
        paramsItem.value = this.paramsForm[i]
        this.formDynamic.items.push(paramsItem)
        this.index = this.index + 1
        // console.info(i +" "+ this.paramsForm[i][0] +" " )
      }
    },

    handleAdd () {
      this.index++;
      this.formDynamic.items.push({
        value: {},
        index: this.index,
        status: 1
      });
      console.info("this.formDynamicthis.formDynamic",this.formDynamic)
    },
    handleRemove (index) {
      this.formDynamic.items[index].status = 0;
    }
  }
}
</script>
