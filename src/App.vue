<template>
  
  <router-view v-if="isRouterAlive">
  </router-view>
 
  <!-- <database /> -->
</template>
<script>
  import login from './userview//login.vue'
  import database from './userview/database.vue'
  export default{
    provide () {
      return {
        reload: this.reload,
      }
    },
    data() {
      return {
        isRouterAlive: true
      }
    },
    methods:{
      reload() {
        this.isRouterAlive = false
        this.$nextTick(function () {
          this.isRouterAlive = true
        })
      },
    },
    components: {
      'login': login,
      'database': database
    }
  }
</script>