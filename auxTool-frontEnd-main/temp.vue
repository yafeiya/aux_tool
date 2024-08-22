<template>
  <div>
    <button @click="getModelData">获取模型数据</button>
    <pre>{{ modelData }}</pre>
  </div>
</template>

<script>
export default {
  data() {
    return {
      cells: [
        {
          attrs: {
            text: {
              text: "策略学习"
            }
          },
          data: {
            classnum: "768",
            fatherLabel: "强化学习",
            futurerewarddiscount: "45",
            grandLabel: "模型模板",
            modelbatch: "56",
            modelurl: "46q4",
            name: "A3C",
            networkdepth: "7867",
            selflabel: "策略学习"
          }
        },
        // 其他对象...
      ],
      modelData: null
    };
  },
  methods: {
    getModelData() {
      const result = this.cells
        .filter(cell => cell.data.grandLabel === "模型模板")
        .map(cell => cell.data);

      this.modelData = result.length > 0 ? result : "没有找到符合条件的数据";
    }
  }
};
</script>

<style scoped>
/* 自定义样式 */
</style>
