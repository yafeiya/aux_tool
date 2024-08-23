<template>
  <div ref="barChart" style="width: 100%; height: 400px"></div>
</template>

<script>
import * as echarts from "echarts";

export default {
  name: "BarChart",
  data() {
    return {
      chartData: {
        categories: [

        ],
        series: [
          {
            name: "系列 1",
            data: [0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8],
            color: "#ff4d4f", // 红色
          },
          {
            name: "系列 2",
            data: [0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9],
            color: "#4d88ff", // 蓝色
          },
        ],
      },
    };
  },
  mounted() {
    this.initChart();
  },
  methods: {
    initChart() {
      const chart = echarts.init(this.$refs.barChart);
      this.chartData.categories = Array.from({ length: 50 }, (v, i) => (i + 1) * 32); // 生成 50 个数
      const option = {
        tooltip: {},
        legend: {
          data: this.chartData.series.map((s) => s.name),
        },
        xAxis: {
          type: "category",
          data: this.chartData.categories,
          axisLabel: {
            interval: 0,
            rotate: 45, // 旋转标签
          },
        },
        yAxis: {
          type: "value",
        },
        series: this.chartData.series.map((s) => ({
          name: s.name,
          type: "line",
          data: s.data,
          itemStyle: {
            color: s.color,
          },
        })),
      };
      chart.setOption(option);
    },
  },
};
</script>

<style scoped></style>
