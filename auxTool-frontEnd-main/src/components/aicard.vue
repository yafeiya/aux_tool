<template>
  <Card :bordered="false" style="width:100%;height:232px;display:flex;align-items:center;justify-content:center;">
    <!-- 空白AI卡片，用于占位 -->
    <!-- <button @click="goToAI" style="background: none; border: none; cursor: pointer; outline: none;">
      <img src="/aicard.png" alt="AI Card" style="width:60px;height:60px;" />
    </button> -->
    <!-- <div style="margin-top: 10px; font-size: 16px; color: #666; text-align: center;">智能建模</div> -->
    <a @click= goToAI() style="position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%);">
        <img src="/aicard.png" alt="AI Card" style="width:80px;height:80px;margin-top: 20%;margin-bottom: 8%;" />
        <h3 style="margin-top: 5%;margin-bottom: 18%; color: #2d8cf0; font-size: 20px; font-weight: 550;">{{"智能建模" }}</h3>
      </a>


  </Card>
</template>

<script>
export default {
  name: 'aicard',
  methods: {
    goToAI() {
      // window.location.href = 'http://localhost:3001/';
      window.open('http://localhost:3001/', '_blank');
    }
  }
}
</script>

<style scoped>
/* 保持与addcard一致的样式 */
</style> 