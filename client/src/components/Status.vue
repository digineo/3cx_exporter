<template lang="pug">
b-card(header="Status")
  b-card-text
    b-list-group
      b-list-group-item(class="d-flex justify-content-between align-items-center")
        | Connected

        b-icon(v-if="connected" icon="check-square" scale="2" variant="success")
        b-icon(v-else icon="x-circle" scale="2" variant="danger")

      b-list-group-item(class="d-flex justify-content-between align-items-center") 
        | Last prometeus request
        b {{lastRequest}}
      b-list-group-item(class="d-flex justify-content-between align-items-center") 
        | Uptime
        b {{uptime}}
      b-list-group-item(class="d-flex justify-content-between align-items-center") 
        | Requests count: 
        b {{requestsCount}}

  
</template>

<script>
export default {
  name: 'settings',
  data(){
    return {
        connected: false,
        lastRequest:"24.09.2022 13:00:01",
        uptime: "36h 10m 4s",
        requestsCount: 999
    }
  },
  
}
</script>

<style scoped>

</style>
