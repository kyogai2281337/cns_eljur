<template>
  <header class="header container header__container">
    <h1 class="h1">Главная страница</h1>
    <button
      class="open_menu-btn btn-no-bg"
      @click="isMenuOpen = true"
      id="open_sidemenuBtn"
    >
      <svg
        width="60"
        height="60"
        viewBox="0 0 60 60"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <g clip-path="url(#clip0_262_259)">
          <path
            d="M24.4167 45.1668C24.4167 45.6531 24.6098 46.1194 24.9536 46.4632C25.2974 46.807 25.7638 47.0002 26.25 47.0002H50V43.3335H26.25C25.7638 43.3335 25.2974 43.5266 24.9536 43.8705C24.6098 44.2143 24.4167 44.6806 24.4167 45.1668Z"
            fill="black"
          />
          <path
            d="M11.5 35.1668C11.5 35.6531 11.6932 36.1194 12.037 36.4632C12.3808 36.807 12.8471 37.0002 13.3333 37.0002H50V33.3335H13.3333C13.0926 33.3335 12.8542 33.3809 12.6317 33.4731C12.4093 33.5652 12.2072 33.7002 12.037 33.8705C11.8667 34.0407 11.7317 34.2428 11.6396 34.4652C11.5474 34.6877 11.5 34.9261 11.5 35.1668Z"
            fill="black"
          />
          <path
            d="M22.3333 25.1668C22.3333 25.6531 22.5265 26.1194 22.8703 26.4632C23.2141 26.807 23.6804 27.0002 24.1667 27.0002H50V23.3335H24.1667C23.6804 23.3335 23.2141 23.5267 22.8703 23.8705C22.5265 24.2143 22.3333 24.6806 22.3333 25.1668Z"
            fill="black"
          />
          <path
            d="M11.25 13.3335C11.0092 13.3335 10.7708 13.3809 10.5484 13.4731C10.326 13.5652 10.1239 13.7002 9.95363 13.8705C9.78339 14.0407 9.64834 14.2428 9.55621 14.4652C9.46408 14.6877 9.41666 14.9261 9.41666 15.1668C9.41666 15.4076 9.46408 15.646 9.55621 15.8684C9.64834 16.0908 9.78339 16.293 9.95363 16.4632C10.1239 16.6334 10.326 16.7685 10.5484 16.8606C10.7708 16.9527 11.0092 17.0002 11.25 17.0002H50V13.3335H11.25Z"
            fill="black"
          />
        </g>
      </svg>
    </button>
  </header>

  <main class="main main__container container">
    <div class="text-wrapper">
      <h2 class="h2">Скоро здесь появятся новости!</h2>
      <p class="text text-temp">
        Сервис находится в разработке. Мы активно работаем над проектом и уже
        скоро здесь появится новостная лента. Следите за обновлениями!
      </p>
    </div>

    <p class="text text-temp text-warning">
      Для начала работы перейдите в меню и авторизуйтесь
    </p>
  </main>
  <SideMenu
    v-model:isOpen="isMenuOpen"
    :first_name="first_name"
    :last_name="last_name"
  />
</template>

<script>
import SideMenu from "@/components/sidemenu/sidemenu.vue";
import userApi from "@/components/api/user";

export default {
  components: {
    SideMenu,
  },
  data() {
    return {
      isMenuOpen: false,
      first_name: "",
      last_name: "",
      email: "",
      role: "",
    };
  },
  async mounted() {
    const respProfile = await userApi.getProfile();
    if (respProfile.data) {
      this.first_name = respProfile.data.first_name;
      this.last_name = respProfile.data.last_name;
      this.email = respProfile.data.email;
      this.role = respProfile.data.role;
    } else {
      this.$router.push("/");
    }
  },
};
</script>
