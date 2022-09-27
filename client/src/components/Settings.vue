<template lang="pug">
b-card(header="Settings")
  b-card-text
    b-form-group(id="host-input-group" label="PBX FQDN" description="PBX hostname and port")
      b-form-input(id="host-input" v-model="hostname")
    b-form-group(id="login-input-group" label="Login" ).mt-2
      b-form-input(id="login-input" v-model="login")
    b-form-group(id="password-input-group" label="Password" ).mt-2
      b-form-input(id="password-input" type="password" v-model="password")
    b-container
      b-row 
        b-col
          b-button(variant="primary" @click="saveSettings").mt-2 Save
      b-row.mt-2
        b-col
          b-alert(:show="showOk"  variant="success") Settings saved 
          b-alert(:show="showError"  variant="danger") Settings not saved {{error}} 

          

  
</template>

<script>
import {GetConfig, SetConfig} from "../services/api"

export default {
  name: 'settings',
  data(){
    return {
      hostname:"",
      login:"",
      password:"",
      showOk:false,
      showError:false,
      error:""
    }
  },
  methods:{
    async saveSettings(){
      try {
        await SetConfig(this.hostname,this.login,this.password)
        this.showOk = true
        console.log("here1")

      } catch(e){
        console.log("here1")
        this.error = e
        this.showError = true
      }

      
    }
  },

  async mounted(){
    const settings = await GetConfig()
    this.hostname = settings.Hostname
    this.login = settings.Username
    this.password = settings.Password
  }
}
</script>

<style scoped>

</style>
