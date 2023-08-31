
<template>
  <div id = "main" ></div>

</template>
 
<script>
import * as echarts from 'echarts';
export default {
  data() {
    return {
      pic_x:[],
      pic_y1:[],
      pic_y2:[],
    }
  },
  methods:{
    fullCloseRandom(n, m) {
      var result = Math.random()*(m+1-n)+n;
      while(result>m) {
          result = Math.random()*(m+1-n)+n;
      }
      return result;
    }
  },
  mounted(){
    for(var i = 1;i <= 1000;i++) {
      this.pic_x.push(i)
      var res1 = Math.sqrt(i) + this.fullCloseRandom(-1, 1)
      var res2 = 0.5 * (Math.sqrt(i) + this.fullCloseRandom(-1, 1))
      this.pic_y1.push(res1)
      this.pic_y2.push(res2)
    }
    let option = {
      title: {
        text: '数据可视化'
      },
      legend: {
        data: ['y1', 'y2']
      },
      grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        // data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        data: this.pic_x
      },
      yAxis: {
        type: 'value',
      },
      series: [
        {
        // data: [820, 932, 901, 934, 1290, 1330, 1320],
        name: 'y1',
        symbol: 'none', //去掉折线上面的小圆点
        data: this.pic_y1,
        type: 'line',
        // areaStyle: {}
        },
        {
        // data: [820, 932, 901, 934, 1290, 1330, 1320],
        name: 'y2',
        symbol: 'none', //去掉折线上面的小圆点
        data: this.pic_y2,
        type: 'line',
        // areaStyle: {}
        },
      ],
      grid:{
        height: 'auto',
      }
};
    console.info("11")
    // 传递一个dom元素
    let psgTimeCharts = echarts.init(document.getElementById("main"))
    console.info("122")
    //传入一个配置项
    psgTimeCharts.setOption(option)
    console.info("333")
    // var charDom = document.getElementById('operationAlarm');
    // var myChart = echarts.init(charDom)
    // operationAlarmOption && myChart.setOption(operationAlarmOption);
  }
}
</script>
 
<style scoped>
  #main {
    width: 600px;
    height:400px;
    margin: auto;
    margin-top: 100px
  }
</style>
