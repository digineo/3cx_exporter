<template lang="pug">
b-card(header="Settings")
  b-card-text
    b-form-group(id="host-input-group" label="PBX FQDN" description="PBX hostname and port")
      b-form-input(id="host-input" v-model="hostname")
    b-form-group(id="login-input-group" label="Login" ).mt-2
      b-form-input(id="login-input" v-model="login")
    b-form-group(id="password-input-group" label="Password" ).mt-2
      b-form-input(id="password-input" type="password" v-model="password")
    b-button(variant="primary").mt-2 Save

  
</template>

<script>
export default {
  name: 'settings',
  data(){
    return {
      hostname:"3cx.example.com",
      login:"root",
      password:"root"
    }
  }
}
</script>

<style scoped>

</style>
