<template>
  
  <router-view>
  </router-view>
 
  <!-- <database /> -->
</template>
<script>
  import login from './components/login.vue'
  import database from './components/database.vue'
  export default{
    components: {
      'login': login,
      'database': database
    }
  }
</script>