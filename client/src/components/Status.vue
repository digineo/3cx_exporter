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
        | Citrix License Key
        b {{key}}
      b-list-group-item(class="d-flex justify-content-between align-items-center") 
        | Citrix Version
        b {{version}}
      b-list-group-item(class="d-flex justify-content-between align-items-center") 
        | Requests count: 
        b {{requestsCount}}

  
</template>

<script>
import {GetStatus} from "../services/api"
import formatDate from '../services/formatDate'
import FormatDate from "../services/formatDate"

export default {
  name: 'settings',
  data(){
    return {
        connected: false,
        lastRequest:"24.09.2022 13:00:01",
        requestsCount: 999,
        key: "",
        version:""
    }
  },
  methods: {
    async getStatus(){
      const res = await GetStatus()
      this.connected = res.connected
      this.lastRequest = formatDate(res.last_request)
      this.requestsCount = res.requests_count
      this.key = res.key
      this.version = res.version
    }
  },
  async mounted(){
    setInterval(async () => {
      await this.getStatus()
    }, 2000);
    

  }
  
}
</script>

<style scoped>

</style>
