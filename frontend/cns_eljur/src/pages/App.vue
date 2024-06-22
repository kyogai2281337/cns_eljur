<template>
  <div class="header">
    <div class="main">
      <div class="header-line">
        <div class="main-rectangle">
          <template v-if="!isLogin">
            <a
              class="menu-item"
              href=""
              v-for="(item, index) in menuItems"
              :key="index"
              @mouseenter="handleMouseEnter(index)"
              @mouseleave="handleMouseLeave(index)"
              :class="{ 'active-tab': activeTab === index }"
              @click.prevent="item === 'Вход' ? showModal() : ''"
            >{{ item }}</a>
          </template>
          <template v-else>
            <a
              class="menu-item"
              href=""
              v-for="(item, index) in menuItems.slice(0, -1)"
              :key="index"
              @mouseenter="handleMouseEnter(index)"
              @mouseleave="handleMouseLeave(index)"
              :class="{ 'active-tab': activeTab === index }"
            >{{ item }}</a>
            <a class="menu-item" @click="toggleMenu" @mouseenter="handleMouseEnter(menuItems.length)"
              @mouseleave="handleMouseLeave(menuItems.length - 1)"
              :class="{ 'active-tab': activeTab === 4 }"
              >
              {{ username }}
              <transition name="fade">
                <div v-if="showMenu" class="dropdown-menu" ref="dropdownMenu">
                  <a @click.prevent="logout">Выход</a>
                </div>
              </transition>
            </a>
          </template>
        </div>
        <div class="cns">
          <p>{{ cnsText }}</p>
        </div>
        <div class="main-title">
          <p v-html="mainTitle"></p>
        </div>
        <div class="main-btn">
          <a class="main-button" href="" @click.prevent="startWork">{{ startWorkText }}</a>
          <div class="arrow-rectangle"></div>
          <img src="@/assets/images/main/arrow-right.png" class="arrow-img" />
        </div>
      </div>
      <p>
        <img src="@/assets/images/main/line1.png" class="line1" />
        <img src="@/assets/images/main/line2.png" class="line2" />
      </p>
    </div>
    <div class="info">
      <div class="info-text">
        <p>{{ infoText }}</p>
      </div>
    </div>
    <div class="support">
      <div class="supp">
        <p>{{ supportText }}</p>
      </div>
    </div>
    <!-- Модальное окно для входа -->
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
                <img src="@/assets/images/main/arrow-auth.png" class="auth-img" />
              </div>
            </a>
          </div>
        </div>
      </div>
      <div class="notifications">
        <div v-for="(notification, index) in notifications" :key="index" :class="['notification', notification.type]" @click="dismissNotification(index)">
          {{ notification.message }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ApiService from '@/components/ApiService';

export default {
  name: 'App',
  data() {
    return {
      activeTab: 0,
      activeTabInt: null,
      isLogin: false,
      username: '',
      menuItems: ['Главная', 'О нас', 'Поддержка', 'Вход'],
      cnsText: 'CNSELJUR',
      mainTitle: 'Генерация учебного <br> расписания для <br> образовательных <br> организаций',
      startWorkText: 'Начать работу',
      infoText: 'О нас',
      supportText: 'Поддержка',
      isModalVisible: false,
      showMenu: false,
      loginForm: {
        email: '',
        password: ''
      },
      notifications: []
    };
  },
  methods: {
    mounted() {
      const token = localStorage.getItem('token');
      if (token) {
        this.isLogin = true;
        this.username = 'Пидорас'; // Здесь ваше имя пользователя
      }
    },
    startWork() {
      // Реализация вашей логики начала работы
    },
    showModal() {
      this.isModalVisible = true;
    },
    closeModal() {
      this.isModalVisible = false;
    },
    handleMouseEnter(index) {
      this.activeTab = index;
      clearTimeout(this.activeTabInt);
    },
    handleMouseLeave() {
      this.activeTabInt = setTimeout(() => {
        this.activeTab = 0;
      }, 200);
    },
    async login() {
      try {
        const response = await ApiService.post('/api/auth/signin', {
          email: this.loginForm.email,
          password: this.loginForm.password
        });
        localStorage.setItem('token', response.token);
        this.isLogin = true;
        this.username = 'Пидорас'; // Здесь ваше имя пользователя
        this.showMessage('success', 'Успешный вход');
        this.closeModal();
      } catch (error) {
        this.showMessage('error', 'Ошибка входа. Проверьте почту и пароль.');
      }
    },
    showMessage(type, message) {
      this.notifications.push({ type, message });
      setTimeout(() => {
        this.dismissNotification(this.notifications.length - 1);
      }, 5000);
    },
    dismissNotification(index) {
      this.notifications.splice(index, 1);
    },
    toggleMenu() {
      this.showMenu = !this.showMenu;
      if (this.showMenu) {
        setTimeout(() => {
          document.addEventListener('click', this.closeDropdownMenu);
        }, 100);
      } else {
        document.removeEventListener('click', this.closeDropdownMenu);
      }
    },
    closeDropdownMenu(event) {
      if (!this.$refs.dropdownMenu.contains(event.target)) {
        this.showMenu = false;
        document.removeEventListener('click', this.closeDropdownMenu);
      }
    },
    logout() {
      localStorage.removeItem('token');
      this.isLogin = false;
      this.username = '';
    }
  }
};
</script>

<style src="@/assets/css/main.css"></style>
