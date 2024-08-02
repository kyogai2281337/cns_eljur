<template>
  <div class="container homePage">
    <div class="panel homePage">
      <div class="panel-heading">
        <h1 class="title">CNS ELJUR</h1>
        <div class="username">
          <h3 class="username-msg">Добро пожаловать</h3>
          <h3 class="username">{{ username }}</h3>
        </div>
      </div>
      <div class="panel-body">
        <a v-if="role==='superuser'" class="hrefBtn" href="/db">База данных</a>
        <a v-if="role==='superuser'" class="hrefBtn" href="/contra">Конструктор</a>
        <a v-if="role==='superuser'" class="hrefBtn" href="/files">Файлы</a>
        <br>
        <a class="hrefBtn" href="/">Выход</a>
      </div>
    </div>
  </div>
</template>

<style>
  @import '@/assets/css/homePage.css';
</style>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

@Options({
  data: () => {
    return {
      username: '',
      role: ''
    };
  },
  components: {},
  beforeCreate() {
    if (document.cookie.includes('auth')) {
      document.location.href = '/home';
    }
    document.body.className = 'homePage';
  },
  async mounted() {
    this.parseProfile();
  },
  methods: {
    parseProfile() {
      let userData = localStorage.getItem('profile') || '{}';
      let parsedUserData = JSON.parse(userData) as { first_name?: string, last_name?: string, role?: string };
      this.username = (parsedUserData.first_name || '') + ' ' + (parsedUserData.last_name || '');
      this.role = parsedUserData.role
    },
  }
})
export default class HomeView extends Vue {}
</script>