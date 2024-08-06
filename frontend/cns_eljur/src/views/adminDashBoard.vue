<template>
  <div>
    <div class="navbar">
      <div class="links">
        <a href="#">Главная</a>
        <a href="#">Ссылка1</a>
        <a href="#">Ссылка2</a>
        <a href="#">Ссылка3</a>
      </div>
      <div class="title">CNS ELJUR</div>
      <div class="user">
        <a href="#">Пользователь</a>
      </div>
    </div>
    <div class="tiles">
      <div class="tile">
        <div class="tile-title">Название плитки</div>
        <div class="tile-text">Тут может быть текста</div>
        <a href="#" class="tile-button">Продолжить</a>
      </div>
      <div class="tile">
        <div class="tile-title">Название плитки</div>
        <div class="tile-text">Тут может быть текста</div>
        <a href="#" class="tile-button">Продолжить</a>
      </div>
      <div class="tile">
        <div class="tile-title">Название плитки</div>
        <div class="tile-text">Тут может быть текста</div>
        <a href="#" class="tile-button">Продолжить</a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdminDashboard',
  data() {
    return {
      // Define your data properties here
    }
  },
  methods: {
    // Define your methods here
  },
  mounted() {
    // Perform any initialization or data fetching here
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100..900;1,100..900&display=swap');
body {
  background-color: #ffffff;
  color: #000000;
  margin: 0;
  padding: 0;
}

.navbar {
  background-color: #db3d3d;
  padding: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.navbar a {
  color: #ffffff;
  text-decoration: none;
  margin: 0 10px;
}

.navbar .title {
  flex-grow: 1;
  text-align: center;
  color: #ffffff;
  font-size: 24px;
  font-weight: bold;
}

.tiles {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 20px;
}

.tile {
  background-color: #db3d3d;
  color: #ffffff;
  border-radius: 15px;
  width: 250px;
  height: 250px;
  margin: 10px;
  padding: 20px;
  text-align: center;
  box-sizing: border-box;
  position: relative;
  transition: transform 0.3s ease, background 0.3s ease;
}

.tile:hover {
  transform: scale(1.05);
  background: linear-gradient(145deg, #db3d3d, #b83f3f);
}

.tile-title {
  font-size: 20px;
  font-weight: bold;
  margin-top: 20px;
}

.tile-text {
  margin: 20px 0;
  font-size: 16px;
}

.tile-button {
  background-color: #ffffff;
  color: #db3d3d;
  padding: 10px 20px;
  text-decoration: none;
  border-radius: 5px;
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
}
</style>
