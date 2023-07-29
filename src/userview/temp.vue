<MenuGroup title="数据集" class="menu-table"></MenuGroup>
<Submenu name="1">
  <template #title>
    <Icon type="ios-navigate"></Icon>
    数值数据集
  </template>
  <MenuItem v-for="(item, index) in this.numDataList" :key ="index" :name='nameList[index]'> {{ item}}</MenuItem>
  <!-- <MenuItem name="1-1">火力分配1</MenuItem>
  <MenuItem name="1-2">火力分配2</MenuItem> -->
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