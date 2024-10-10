<template>
  <div class="container homePage">
    <div class="panel homePage">
      <div class="panel-heading homePage">
        <h1 class="title homePage">CNS ELJUR</h1>
        <div class="username homePage">
          <h3 class="username-msg">Добро пожаловать</h3>
          <h3 class="username">{{ username }}</h3>
        </div>
      </div>
      <div class="panel-body homePage">
        <a v-if="role==='superuser'" class="hrefBtn homePage" href="#/db">База данных</a>
        <a v-if="role==='superuser'" class="hrefBtn homePage" href="#/contra">Конструктор</a>
        <a v-if="role==='superuser'" class="hrefBtn homePage" href="#/files">Файлы</a>
        <a class="hrefBtn homePage" href="#/">Выход</a>
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
    try {
      this.parseProfile();
    } catch (error) {
      console.log(error)
    }
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