<template>
  <div class="container authPage">
    <div class="auth authPage">
      <h1 class="authPage" style="font-size:100%;text-align: center;">{{authType===0?"Авторизация":"Регистрация"}}</h1>
      <h3 style="font-size:100%;text-align: center;">{{ message }}</h3>
      <div class="input-container authPage">
        <input v-model="email" class="authPage" type="email" id="email" required placeholder=" ">
        <label class="authPage" for="email">Почта</label>
      </div>
      <div class="input-container authPage">
        <input v-model="password" class="authPage" type="password" id="password" required placeholder=" ">
        <label class="authPage" for="password">Пароль</label>
      </div>
      <div class="input-container authPage" v-if="authType!==0">
        <input v-model="firstname" class="authPage" type="text" id="surname" required placeholder=" ">
        <label class="authPage" for="surname">Фамилия</label>
      </div>
      <div class="input-container authPage" v-if="authType!==0">
        <input v-model="lastname" class="authPage" type="text" id="name" required placeholder=" ">
        <label class="authPage" for="name">Имя</label>
      </div>
      <div class="buttons authPage">
        <button class="authPage" @click="auth(0)">Войти</button>
        <button class="authPage" @click="auth(1)">Создать</button>
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
import userApi from '@/components/api/user'

@Options({
  data: () => {
    return {
      authType: 0,
      message: '',
      firstName: '',
      lastName: '',
      email: '',
      password: ''
    };
  },
  components: {},
  beforeCreate: function() {
    if (document.cookie.includes('auth')) {
      document.location.href = '/home';
    }
    document.body.className = 'authPage';
  },
  methods: {
    async auth(type: number): Promise<void> {
      if (this.authType===type) {
        if (this.authType===0) {
          const data = await userApi().signin({email:this.email,password:this.password});
          if (!data.error&&data.data?.status===200) {
            document.cookie = `auth=${data.data?.token}`
            userApi().getProfile().then(dataProfile=>{
              if (!data.error) {
                localStorage.setItem('profile', JSON.stringify(dataProfile.data));
                document.location.href = '/home';
              }
            }).catch(()=>{
              this.message=`Ошибка авторизации`
              setTimeout(() => {
                this.message=``
              }, 3500);
            })
          } else {
            this.message=`Ошибка авторизации`
            setTimeout(() => {
              this.message=``
            }, 3500);
          }
        } else {
          const data = await userApi().signup({email:this.email,password:this.password,first:this.firstname,last:this.lastname});
          if (!data.error) {
            this.authType=0;
          } else {
            this.message=`Ошибка регистрации`
            setTimeout(() => {
              this.message=``
            }, 3500);
          }
        }
      } else {
        this.authType=type
      }
    }
  }
})
export default class AuthPage extends Vue {}
</script>
