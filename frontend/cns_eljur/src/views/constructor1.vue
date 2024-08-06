<template>
    <div class = "header">
      <div class="constructor"></div>
        <div class="const-title">
            <a>Конструктор</a>
        </div>
        <div class="const-input">
        <div class="bat">
          <input class="batnn" type="text" placeholder="Заполнить данные">
          </div>
        <div class="but">
          <input class="butnn" type="text" placeholder="Проверить данные">
          </div>
          </div>
          <div class="const-btn">
            <a class="button" href="construct2.html">Сгенерировать</a>
            <div class="arrow-rectangle" href="construct2.html"></div>
            <img src="@/assets/images/arrow-right.png" class="arrow-img">
          </div>
      </div>
</template>

<script>
export default {
  name: "ComponentOne",
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100..900;1,100..900&display=swap');
@import '@/assets/css/constructor1.css';
</style>
  