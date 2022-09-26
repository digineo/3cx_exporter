<template lang="pug">
  .app
    b-navbar( type="dark" variant="info")
      b-navbar-brand
        b-container 
          | 3CX EXPORTER
    b-container 
          
      b-row
        b-col
          Settings.mt-2
        b-col
          Status.mt-2
</template>

<script>
import Settings from './components/Settings.vue'
import Status from './components/Status.vue'

export default {
  name: 'App',
  components: {
    Settings,
    Status
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
