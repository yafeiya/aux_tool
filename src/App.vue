<template>

  <router-view>
  </router-view>
 
  <!-- <database /> -->
</template>
<script>
  import login from './components/login.vue'
  import database from './components/database.vue'
  import example from "./components/example.vue";

  export default{
    components: {
      'login': login,
      'example': example
    }
  }
</script>
