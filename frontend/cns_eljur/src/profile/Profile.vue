<script>
import Header from '../components/Header.vue';
import Get from '../components/Get.vue';

import { dataStructs } from '../assets/js/datastructures.js';
import { makeFetch } from '@/assets/js/request_forgery';
export default {
  components: {
    Header: Header,
    Get: Get,
  },
  data() {
    return {
      title: "Profile",
      isAuth: true,
      email: null,
      dataStructs: dataStructs,

    }
  },
  methods: {
    changeAuth () {
      this.isAuth = !this.isAuth;
    },
    async get(url) {
      return await makeFetch(url);
    },
    async logout() {
      await this.get("/api/auth/private/logout");
    },
    async deleteUser() {
      await this.get("/api/auth/private/delete");
    }
  }

}
</script>

<template>
  <div>
    <Header :title="title" :auth="isAuth"/>
    <div class="section">
      <div class="container">
        <div class="row mh-100 flex-column justify-content-around">
          <div class="col-12 d-flex flex-column align-items-center">
            <Get list_title="Your profile" url="/api/auth/private/profile"/>
            <a @click="logout()">logout</a>
            <a @click="deleteUser()">delete</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>