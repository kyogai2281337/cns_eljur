<template>
  <div class="container authPage">
    <div class="auth authPage">
      <h1 class="authPage" style="font-size:100%;text-align: center;">{{authType===0?"Авторизация":"Регистрация"}}</h1>
      <div class="input-container authPage">
        <input class="authPage" type="email" id="email" required placeholder=" ">
        <label class="authPage" for="email">Почта</label>
      </div>
      <div class="input-container authPage">
        <input class="authPage" type="password" id="password" required placeholder=" ">
        <label class="authPage" for="password">Пароль</label>
      </div>
      <div class="input-container authPage" v-if="authType!==0">
        <input class="authPage" type="text" id="surname" required placeholder=" ">
        <label class="authPage" for="surname">Фамилия</label>
      </div>
      <div class="input-container authPage" v-if="authType!==0">
        <input class="authPage" type="text" id="name" required placeholder=" ">
        <label class="authPage" for="name">Имя</label>
      </div>
      <div class="buttons authPage">
        <button class="authPage">Войти</button>
        <button class="authPage">Создать</button>
      </div>
    </div>
    <div class="content authPage">
      Hello World!
    </div>
  </div>
</template>

<style>
  @import '@/assets/css/authPage.css';
</style>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

@Options({
  data: () => {
    return {
      authType: 1
    };
  },
  components: {},
  beforeCreate: function() {
    document.body.className = 'authPage';
  },
})
export default class HomeView extends Vue {}
</script>
