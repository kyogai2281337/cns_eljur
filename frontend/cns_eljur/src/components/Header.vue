<script>
  export default {
    props: {
        title: String,
        auth: {
          type: Boolean,
          required: true,
        }, 
    },
    data() {
        return {
            email: "Guest"
        }
    }

  }
</script>

<template>
<header class="header">
    <div class="container">
      <div class="row mh-100 flex-column">
        <div class="col-12">
          <nav class="header__nav d-flex justify-content-between align-items-center">
            <a href="/" class="logo">CNSeljur</a>
            <ul class="header__menu d-flex">
              <li><a href="/" class="header__nav-link">Главная</a></li>
              <li><a href="" class="header__nav-link">Поддержка</a></li>
              <li><a href="" class="header__nav-link">О нас</a></li>
              <div v-if="auth"  class="d-flex">
                <li><a href="" class="header__nav-link">{{ email }}</a></li>
              </div>
              <div v-else class="d-flex">
                <li><a href="" class="header__nav-link">Вход</a></li>
                <li><a href="" class="header__nav-link">Регистрация</a></li>
              </div>

            </ul>
          </nav>
        </div>
        <div class="col-12 flex-grow-1 d-flex flex-column justify-content-center">
          <h1 class="header__title">{{ title }}</h1>
          <p class="header__text">Some more cool text</p>
          <p class="header__text">Another one cool</p>
        </div>
      </div>
    </div>
  </header>
</template>

<style scoped>
.header {
  background: url(../assets/img/exph3.jpg) no-repeat center center /cover;
  background-attachment: fixed;
}

.logo {
  font-size: 24px;
  display: inline-block;
  font-family: var(--r700);
  color: var(--white);
  background: var(--blue);
  border: 3px solid var(--light_text);
  border-radius: 5px;
  padding: 7px 18px;
  text-transform: uppercase;
  line-height: 17px;
  letter-spacing: 1.5px;
  box-shadow: 3px 3px 5px 5px rgba(30, 30, 30, 0.15) inset;
}

.header__nav-link {
  color: var(--light_text);
  font-family: var(--r300);
  font-size: 20px;
  padding: 41px 26px 30px 26px;
  display: block;
  transition: background 300ms, color 300ms;
}

.header__nav-link:hover {
  background: var(--blue_75);
  color: var(--white);
}

.header__title {
  color: var(--light_text);
  font-size: 72px;
  font-family: var(--r900);
  text-transform: uppercase;
  font-weight: lighter;
  letter-spacing: 3px;
  line-height: 53px;
  margin-bottom: 20px;
}

.header__text:first-of-type {
  color: var(--light_text);
  font-size: 36px;
  line-height: 25px;
  margin-bottom: 28px;
  text-transform: uppercase;
}

.header__text:nth-of-type(2) {
  color: var(--light_text);
  font-size: 18px;
  line-height: 22px;
  text-align: justify;
}
</style>