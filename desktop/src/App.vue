<template>
  <router-view/>
</template>

<style>
</style>

<script lang="ts">
  import userApi from '@/components/api/user'
  
  export default {
    mounted() {
      this.updateProfile()
    },
    methods: {
      updateProfile(){
        if (localStorage.getItem('profile')) {
          userApi().getProfile().then(dataProfile=>{
           if (!dataProfile.error) {
             localStorage.setItem('profile', JSON.stringify(dataProfile.data));
             console.log('update profile');
           } else {
            return
           }
          }).catch(()=>{
           console.log('error update profile');
          })
        }
      }
    }
  }
</script>