<template>
  <div class="header">
    <div class="main">
      <div class="header-line">
        <div class="main-rectangle">
          <a
            class="menu-item"
            href=""
            v-for="(item, index) in menuItems"
            :key="index"
            @mouseenter="handleMouseEnter(index)"
            @mouseleave="handleMouseLeave(index)"
            :class="{ 'active-tab': activeTab === index }"
            @click.prevent="item === 'Вход' ? showModal() : ''"
            >{{ item }}</a
          >
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
          <img src="@/assets/arrow-right.png" class="arrow-img" />
        </div>
      </div>
      <p>
        <img src="@/assets/line1.png" class="line1" />
        <img src="@/assets/line2.png" class="line2" />
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
          <input class="input1" type="text" placeholder="Логин" />
          <input class="input2" type="password" placeholder="Пароль" />
          <div class="auth-btn">
            <a class="auth-btn1" href="" @click.prevent="login">
              Войти
              <div class="auth-rectangle">
                <img src="@/assets/arrow-auth.png" class="auth-img" />
              </div>
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      activeTab: 0,
      activeTabInt: null,
      isLogin: false,
      username: null,
      menuItems: ['Главная', 'О нас', 'Поддержка', 'Вход'],
      cnsText: 'CNSELJUR',
      mainTitle: 'Генерация учебного <br> расписания для <br> образовательных <br> организаций',
      startWorkText: 'Начать работу',
      infoText: 'О нас',
      supportText: 'Поддержка',
      isModalVisible: false
    }
  },
  methods: {
    startWork() {},
    showModal() {
      this.isModalVisible = true
    },
    closeModal() {
      this.isModalVisible = false
    },
    handleMouseEnter(index) {
      this.activeTab = index
      clearTimeout(this.activeTabInt)
    },
    handleMouseLeave() {
      this.activeTabInt = setTimeout(() => {
        this.activeTab = 0
      }, 200)
    },
    login() {}
  }
}
</script>
