<script>
import Header from './components/Header.vue';
import Form from './components/Form.vue';

import { dataStructs } from './assets/js/datastructures.js';
export default {
  components: {
    Header: Header,
    Form: Form,
  },
  data() {
    return {
      title: "Index",
      isAuth: true,
      email: "Guest",
      dataStructs: dataStructs,

    }
  },
  methods: {
    changeAuth () {
      this.isAuth = !this.isAuth;
    },

  }

}
</script>

<template>
  <div>
    <Header :title="title" :auth="isAuth"/>
    <Form :dataset="dataStructs.register" form_title="Register" url="http://localhost/api/signup"/>
  </div>
</template>

<style scoped>

</style>
