<template>
  <div class="config">
    <ConfigGrid v-show="type === 'grid'" />
    <dataloading v-show="type === 'dataloading'" />
    <datapre v-show="type === 'datapre'" />
    <moudles v-show="type === 'moudles'" />
    <visualize v-show="type === 'visualize'" />
  </div>
</template>

<script lang="ts">
  import ConfigGrid from './ConfigGrid/index.vue';
  import dataloading from './dataloading/index.vue';
  import datapre from './datapre/index.vue';
  import moudles from './moudles/index.vue';
  import visualize from './visualize/index.vue';
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
      visualize,
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
          else if(cell.data.fatherLabel=='数据预处理'){
            type.value = 'datapre';
          }
          else if(cell.data.fatherLabel=='模型模板'){
            type.value = 'moudles';
          }
          else if(cell.data.fatherLabel=='模型结果可视化'){
            type.value = 'visualize';
          }
          else{
            type.value = 'grid';
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

</script>

<style lang="less" scoped>
.config{
  width: 300px;
}
</style>
