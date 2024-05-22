<script>
import Header from '../components/Header.vue';

import { dataStructs } from '../assets/js/datastructures.js';
export default {
  components: {
    Header: Header
  },
  data() {
    return {
      title: "Index",
      isAuth: false,
      email: null,
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
  </div>
</template>

<style scoped>

</style>
