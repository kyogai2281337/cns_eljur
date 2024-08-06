<template>
  <router-view/>
</template>

<style>
</style>

<script lang="ts">
import { defineComponent } from 'vue';
import userApi from '@/components/api/user';

export default defineComponent({
  name: 'App',
  mounted() {
    try {
      this.updateProfile();
    } catch (error) {
      console.log(error);
    }
  },
  methods: {
    async updateProfile() {
      if (localStorage.getItem('profile')) {
        try {
          const dataProfile = await userApi().getProfile();
          if (!dataProfile.error) {
            localStorage.setItem('profile', JSON.stringify(dataProfile.data));
            console.log('update profile');
          }
        } catch {
          console.log('error update profile');
        }
      }
    }
  }
});
</script>
