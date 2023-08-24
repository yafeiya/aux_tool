
<template>
  <Row>
    <Col span="12">
      <div id = "main" ></div>
    </Col>
    <Col span="12">
      <div id = "main2"></div>
    </Col>
  </Row>
  <Row>
    <Col span="12">
      <div id = "main3" ></div>
    </Col>
    <Col span="12">
      <div id = "main4"></div>
    </Col>
  </Row>


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
          height: 'auto',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: this.pic_x
        },
        yAxis: {
          type: 'value',
        },
        series: [
          {
          name: 'y1',
          symbol: 'none', //去掉折线上面的小圆点
          data: this.pic_y1,
          type: 'line',
          // areaStyle: {}
          },
          {
          name: 'y2',
          symbol: 'none', //去掉折线上面的小圆点
          data: this.pic_y2,
          type: 'line',
          // areaStyle: {}
          },
        ],
  };
      console.info("11")
      // 传递一个dom元素
      let psgTimeCharts = echarts.init(document.getElementById("main"))
      let psgTimeCharts2 = echarts.init(document.getElementById("main2"))
      let psgTimeCharts3 = echarts.init(document.getElementById("main3"))
      let psgTimeCharts4 = echarts.init(document.getElementById("main4"))
      console.info("122")
      //传入一个配置项
      psgTimeCharts.setOption(option)
      psgTimeCharts2.setOption(option)
      psgTimeCharts3.setOption(option)
      psgTimeCharts4.setOption(option)

      console.info("333")
      // var charDom = document.getElementById('operationAlarm');
      // var myChart = echarts.init(charDom)
      // operationAlarmOption && myChart.setOption(operationAlarmOption);
    }
  }
  </script>

  <style scoped>
    #main {
      width: 300px;
      height:250px;
      margin: auto;
    }
    #main2 {
      width: 300px;
      height:250px;
      margin: auto;
    }
    #main3 {
      width: 300px;
      height:250px;
      margin: auto;
    }
    #main4 {
      width: 300px;
      height:250px;
      margin: auto;
    }
  </style>
