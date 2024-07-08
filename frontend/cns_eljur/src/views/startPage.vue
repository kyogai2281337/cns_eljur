<template>
  <div class = "header">
    <div class="main">
      <div class = "header-line">
        <div class="main-rectangle">
          <div class="menu-rectangle"></div>
          <a class="main-item" @mouseenter="handleMouseEnter(1)" @mouseleave="handleMouseLeave(1)" href="">Главная</a>
          <a class="menu-item" @mouseenter="handleMouseEnter(2)" @mouseleave="handleMouseLeave(2)" href="">О нас</a>
          <a class="menu-item" @mouseenter="handleMouseEnter(3)" @mouseleave="handleMouseLeave(3)" href="">Поддержка</a>
          <a class="menu-item" @mouseenter="handleMouseEnter(4)" @mouseleave="handleMouseLeave(4)" @click="showModal()">Вход</a>
        </div>
        <div class="cns">
          <p>CNSELJUR</p>
        </div>
        <div class="main-title">
          <p>Генерация учебного <br> расписания для <br> образовательных <br> организаций</p>
        </div>
        <div class="main-btn">
          <a class="main-button" href="">Начать работу</a>
            <div class="arrow-rectangle">
              <img src="@/assets/images/arrow-right.png" class="arrow-img">
          </div>
        </div>
        <p>
          <img src="@/assets/images/line1.png" class="line1">
          <img src="@/assets/images/line2.png" class="line2">
        </p>
      </div>
      </div>
    <div class="info">
        <div class="info-text">
            <p>О нас</p>
        </div>
    </div>
    <div class="support">
      <div class="supp">
        <p>Поддержка</p>
      </div>
    </div>
    <div v-if="isModalVisible" class="modal">
      <div class="modal-content">
        <span class="close-btn" @click="closeModal">&times;</span>
        <div class="auth-back">
          <div class="log-in">
            <p>Вход</p>
          </div>
          <input v-model="loginForm.email" class="input1" type="text" placeholder="Почта" />
          <input v-model="loginForm.password" class="input2" type="password" placeholder="Пароль" />
          <div class="auth-btn">
            <a class="auth-btn1" href="" @click.prevent="login">
              Войти
              <div class="auth-rectangle">
                <img src="@/assets/images/arrow-auth.png" class="auth-img" />
              </div>
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import userApi from '@/components/api/user';

export default {
  name: 'startPage',
  data() {
    return {
      activeTab: 0,
      activeTabInt: null,
      loginForm: {
        email: '',
        password: ''
      },
      isModalVisible: false
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
    login() {
      console.log(this.loginForm);
      userApi.loginUser(this.loginForm.email,this.loginForm.password).then(response => {
        console.log(response);
      });
    }
  },
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100..900;1,100..900&display=swap');
@import '@/assets/css/startPage.css';
</style>
