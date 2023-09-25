<template>
  <Button type="info" @click="download">下载测试</Button>
  <div>
    {{ chartcsvmsg }}
  </div>
</template>

<script>
import {downloadCsvFile} from '../api/api.js'
import qs from "qs";
export default{
  methods:{
    download(){
      //模拟下载时，往后端传的数据
      var csvData = {
        dataset_name: "波士顿房价数据集",
        csvName: "动作表",
        Type: "数值数据集",
        Task: "任务1"
      }
      csvData = qs.stringify(csvData)
      downloadCsvFile(csvData).then(res => {
        console.info("下载URL: ", res.data.data.url)
        // 创建一个虚拟的<a>标签
        const a = document.createElement('a');
        a.href = res.data.data.url;
        a.target = '_blank'; // 在新标签页中打开文件
        a.download = 'file.csv'; // 可以自定义文件名
        document.body.appendChild(a);

        // 模拟用户点击链接以触发下载
        a.click();

        // 清除虚拟<a>标签
        document.body.removeChild(a);
      })
      // window.location.href = res.data.data.url
    }
  }
}
</script>
