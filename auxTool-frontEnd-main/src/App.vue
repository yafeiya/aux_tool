<template>
  
  <router-view />
 
  <!-- <database /> -->
</template>
<script>
  import login from './userview//login.vue'
  import database from './userview/database.vue'
  export default{
    components: {
      'login': login,
      'database': database
    }
  }
</script>