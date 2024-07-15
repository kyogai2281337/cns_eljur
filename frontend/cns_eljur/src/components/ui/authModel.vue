<template>
    <NotificationsUi ref="notificationsUi" />
    <div v-if="isVisible" class="modal">
      <div class="modal-content">
        <span class="close-btn" @click="closeModal">&times;</span>
        <div class="auth-back">
          <div v-if="isLoginMode" class="log-in">
            <p>Вход</p>
          </div>
          <div v-else class="register">
            <p>Регистрация</p>
            <input v-model="firstName" :class="{ 'input': true, 'invalid': firstNameInvalid }" type="text" placeholder="Имя" />
            <input v-model="lastName" :class="{ 'input': true, 'invalid': lastNameInvalid }" type="text" placeholder="Фамилия" />
          </div>
          <input v-model="email" :class="{ 'input': true, 'invalid': emailInvalid }" type="text" placeholder="Почта" />
          <input v-model="password" :class="{ 'input': true, 'invalid': passwordInvalid }" type="password" placeholder="Пароль" />
          <div class="auth-btn">
            <button @click="loginOrRegister">{{ isLoginMode ? 'Войти' : 'Зарегистрироваться' }}</button>
            <button class="toggle-mode-btn" @click="toggleMode">
              {{ isLoginMode ? 'Регистрация' : 'Вход' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import userApi from '@/components/api/user';
  import NotificationsUi from './notificationsUi.vue';
  
  export default {
    name: 'AuthModal',
    components: {
      NotificationsUi,
    },
    props: {
      isVisible: {
        type: Boolean,
        default: false
      }
    },
    data() {
      return {
        isLoginMode: true,
        firstName: '',
        lastName: '',
        email: '',
        password: '',
        firstNameInvalid: false,
        lastNameInvalid: false,
        emailInvalid: false,
        passwordInvalid: false
      };
    },
    methods: {
      closeModal() {
        this.$emit('close');
      },
      toggleMode() {
        this.isLoginMode = !this.isLoginMode;
        this.firstName = '';
        this.lastName = '';
        this.email = '';
        this.password = '';
        this.resetValidation();
      },
      addNotification(message) {
        const notificationsUi = this.$refs.notificationsUi;
        if (notificationsUi) {
          notificationsUi.addNotification(message);
        }
      },
      loginOrRegister() {
        this.resetValidation();
        if (this.validateFields()) {
          if (this.isLoginMode) {
            userApi.loginUser(this.email, this.password).then(response => {
                if (response.status) {
                  this.closeModal();
                  this.addNotification('Вход выполнен успешно!');
                  userApi.getUserProfile().then(response => {
                        if (response.status) {
                            localStorage.setItem('profile', JSON.stringify(response));
                            document.location.href = '/';
                        } else {
                            console.log('error');
                            this.addNotification('Ошибка получения информации о пользователе');
                        }
                    });
                } else {
                  console.log('error');
                  this.addNotification('Ошибка входа');
                }
            });
          } else {
            userApi.signupUser(this.email, this.password, this.firstName, this.lastName).then(response => {
                if (response.status) {
                  this.addNotification('Регистрация выполнена успешно!');
                  this.toggleMode()
                } else {
                  console.log('error');
                  this.addNotification('Ошибка регистрации');
                }
            });
          }
        }
      },
      validateFields() {
        let isValid = true;
        if (!this.isLoginMode) {
          if (!this.firstName.trim()) {
            this.firstNameInvalid = true;
            isValid = false;
          }
          if (!this.lastName.trim()) {
            this.lastNameInvalid = true;
            isValid = false;
          }
        }
        if (!this.email.trim()) {
          this.emailInvalid = true;
          isValid = false;
        } else if (!this.validateEmail(this.email)) {
          this.emailInvalid = true;
          isValid = false;
        }
        if (!this.password.trim()) {
          this.passwordInvalid = true;
          isValid = false;
        }
        return isValid;
      },
      validateEmail(email) {
        const re = /\S+@\S+\.\S+/;
        return re.test(email);
      },
      resetValidation() {
        this.firstNameInvalid = false;
        this.lastNameInvalid = false;
        this.emailInvalid = false;
        this.passwordInvalid = false;
      }
    }
  };
  </script>
  
  <style scoped>
  .modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .modal-content {
    background-color: white;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.5);
    min-width: 300px;
    max-width: 600px;
    position: relative;
  }
  
  .close-btn {
    position: absolute;
    top: 10px;
    right: 10px;
    cursor: pointer;
    font-size: 24px;
    color: #333;
  }
  
  .auth-back {
    margin-top: 20px;
  }
  
  .log-in,
  .register {
    text-align: center;
    margin-bottom: 10px;
  }
  
  .input {
    width: calc(100% - 20px);
    margin-bottom: 10px;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    font-size: 16px;
  }
  
  .input.invalid {
    border-color: #db3d3d;
  }
  
  .auth-btn {
    text-align: center;
  }
  
  button {
    background-color: #db3d3d;
    color: white;
    border: none;
    padding: 10px 20px;
    border-radius: 5px;
    cursor: pointer;
    font-size: 16px;
  }
  
  .toggle-mode-btn {
    margin-top: 10px;
    background-color: transparent;
    border: none;
    color: #db3d3d;
    cursor: pointer;
  }
  </style>
  