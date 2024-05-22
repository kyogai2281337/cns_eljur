<script>
import Header from '../components/Header.vue';
import Form from '../components/Form.vue';

import { dataStructs } from '../assets/js/datastructures.js';
export default {
  components: {
    Header: Header,
    Form: Form,
  },
  data() {
    return {
      title: "login/register",
      isAuth: true,
      Auth: true,
      email: "Guest",
      dataStructs: dataStructs,
    };
  },
  methods: {
    changeAuthForm () {
      this.Auth = !this.Auth;
    },

  }

}
</script>

<template>
  <div>
    <Header :title="title" :auth="isAuth"/>
    <div class="section">
      <div class="container">
        <div class="row mh-100 flex-column justify-content-center">
          <div class="col-12 d-flex flex-column align-items-center">
            <Form v-if="Auth" :dataset="dataStructs.register" form_title="Register" url="http://localhost/api/auth/signup"/>
            <Form v-else :dataset="dataStructs.login" form_title="Login" url="http://localhost/api/auth/signin"/>
            <button @click="changeAuthForm" type="button" class="redirect-button">in/up</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
    .section {
      transition: height 0.3s ease-in-out; /* Плавный переход высоты секции */
      background: url(./assets/img/exph3.jpg) no-repeat center center /cover;
      background-attachment: fixed;
    }
    .redirect-button {
      padding: 8px;
      border-radius: 6px;
      width: 8.5%;
      border: 1px solid rgb(53, 57, 94);
      background-color: rgb(147, 147, 179);
      color: rebeccapurple;
      font-size: 18px;
      margin-top: 50px;
    }
</style>
