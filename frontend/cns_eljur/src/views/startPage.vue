<template>
  <main>
    <div class="rectangle">
      <header class="header">
        <nav>
          <ul>
            <li><a class="main-tab" @mouseenter="handleMouseEnter(0)" @mouseleave="handleMouseLeave(0)" :class="{'active-tab':activeTab===0}">Главная</a></li>
            <li><a class="about-tab" @mouseenter="handleMouseEnter(1)" @mouseleave="handleMouseLeave(1)" :class="{'active-tab':activeTab===1}">О нас</a></li>
            <li><a class="support-tab" @mouseenter="handleMouseEnter(2)" @mouseleave="handleMouseLeave(2)" :class="{'active-tab':activeTab===2}">Поддержка</a></li>
            <li v-if="!isAuth"><a class="login-tab" @mouseenter="handleMouseEnter(3)" @mouseleave="handleMouseLeave(3)" :class="{'active-tab':activeTab===3}" @click="showModal()">Войти</a></li>
            <li v-else class="user" style="color: black;" @mouseenter="handleMouseEnter(4)" @mouseleave="handleMouseLeave(4)" :class="{'active-tab':activeTab===4}">
              {{username}}
              <ul v-if="activeTab===4" class="dropdown">
                <li><a href="#">Профиль</a></li>
                <li><a href="/admin">Админ панель</a></li>
                <li><a href="/constructor1">Конструктор</a></li>
                <li><a href="/db">База данных</a></li>
                <li><a @click="logout">Выход</a></li>
              </ul>
            </li>
          </ul>
        </nav>
      </header>
      <div class="logo">CNSELJUR</div>
      <div class="main-title">
        <div class="main-title-text">
        <p>Генерация учебного <br> расписания для <br> образовательных <br> организаций</p>
        </div>
        <a class="main-button" href="">Начать работу
          <svg xmlns="http://www.w3.org/2000/svg" class="arrow-right" width="128" height="128" fill="currentColor" viewBox="0 0 16 16">
            <path fill-rule="evenodd" d="M4 8a.5.5 0 0 1 .5-.5h5.793L8.146 5.354a.5.5 0 1 1 .708-.708l3 3a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708-.708L10.293 8.5H4.5A.5.5 0 0 1 4 8"/>
          </svg>
        </a>
        </div>
      <img src="@/assets/images/line1.png" class="line1">
      <img src="@/assets/images/line2.png" class="line2">
    </div>
    <div class="rectangle">Прямоугольник 2</div>
    <div class="rectangle">Прямоугольник 3</div>
  </main>
  <AuthModal 
      :isVisible="isModalVisible" 
      @close="closeModal" 
    />
</template>
<script>
import AuthModal from '../components/ui/authModel.vue';

export default {
  name: 'startPage',
  components: {
    AuthModal
  },
  beforeCreate: function() {
    document.body.className = 'startPage';
  },
  data() {
    return {
      activeTab: 0,
      activeTabInt: null,
      isModalVisible: false,
      isAuth: false,
      username: null
    };
  },
  mounted() {
    let profile = localStorage.getItem('profile')
    if (profile) {
      profile = JSON.parse(profile);
      this.isAuth = true;
      this.username = profile.first_name + ' ' + profile.last_name;
    } else {
      this.isAuth = false;
      this.username = null;
    }
  },
  methods: {
    handleMouseEnter(index) {
      this.activeTab = index;
      clearTimeout(this.activeTabInt);
    },
    handleMouseLeave() {
      this.activeTabInt = setTimeout(() => {
        this.activeTab = 0;
      }, 200);
    },
    showModal() {
      this.isModalVisible = true;
    },
    closeModal() {
      this.isModalVisible = false;
    },
    logout() {
      localStorage.removeItem('profile');
      document.cookie = `token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
      this.isAuth = false;
      this.username = null;
    },
  }
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100..900;1,100..900&display=swap');
@import '@/assets/css/startPage.css';
</style>