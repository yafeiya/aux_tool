<template>
  <div class="config">
    <ConfigGrid v-show="type === 'grid'" />
    <dataloading v-show="type === 'dataloading'" />
    <datapre v-show="type === 'datapre'" />
    <moudles v-show="type === 'moudles'" />
    <moudletrain v-show="type === 'moudletrain'" />
  </div>
</template>

<script lang="ts">
  import ConfigGrid from './ConfigGrid/index.vue';
  import dataloading from './dataloading/index.vue';
  import datapre from './datapre/index.vue';
  import moudles from './moudles/index.vue';
  import moudletrain from './moudletrain/index.vue';
  // import ConfigEdge from './ConfigEdge/index.vue';
  import FlowGraph from '../../graph/index';
  import './index.less';
  import { defineComponent, ref, provide } from 'vue';
  import { globalGridAttr } from '../../models/global';

  export default defineComponent({
    name: 'Index',
    components: {
      ConfigGrid,
      dataloading,
      datapre,
      moudles,
      moudletrain,
      // ConfigEdge,
    },
    setup() {
      const type = ref('grid');
      const id = ref('');
      // 待优化
      const boundEvent = function (): void {
        const { graph } = FlowGraph;
        graph.on('blank:click', () => {
          type.value = 'grid';
        });
        graph.on('cell:click', ({ cell }) => {
          
          id.value = cell.id;
          if(cell.data.fatherLabel=='数据加载'){
            type.value = 'dataloading';
          }
          if(cell.data.fatherLabel=='数据预处理'){
            type.value = 'datapre';
          }
          if(cell.data.fatherLabel=='模型模板'){
            type.value = 'moudles';
          }
          if(cell.data.fatherLabel=='模型训练'){
            type.value = 'moudletrain';
          }
        });
      };
      boundEvent();
      provide('globalGridAttr', globalGridAttr);
      provide('id', id);

      return {
        type,
        id,
      };
    },
  });

  const list = [
          {
              name:'dataloading',
              label:'数据路径',
              attributes:[
                {
                  name:'dataurl',
                  label:'数据路径',
                },
                {
                  name:'headernum',
                  label:'表头数量',
                },
                {
                  name:'datasetsize',
                  label:'数据集大小',
                },
                {
                  name:'imgsize',
                  label:'图像尺寸',
                },
              ]
          },
          {
              name:'datapre',
              label:'数据预处理',
              attributes:[
                {
                  name:'downsamplescale',
                  label:'下采样因子',
                },
                {
                  name:'augmentationscale',
                  label:'扩充比例',
                },
              ]
          },
          {
              name:'moudles',
              label:'模型模板',
              attributes:[
                {
                  name:'networkdepth',
                  label:'网络深度',
                },
                {
                  name:'classnum',
                  label:'分类类别数',
                },
                {
                  name:'futurerewarddiscount',
                  label:'未来奖励折扣',
                },
              ]
          },
          {
              name:'moudletrain',
              label:'模型训练',
              attributes:[
                {
                  name:'iterations',
                  label:'迭代次数',
                },
                {
                  name:'optimizer',
                  label:'优化器',
                },
                {
                  name:'decayfactor',
                  label:'衰减因子',
                },
                {
                  name:'learningrate',
                  label:'学习率',
                },
              ]
          },
        ]

</script>

<style lang="less" scoped>
.config{
  width: 300px;
}
</style>
