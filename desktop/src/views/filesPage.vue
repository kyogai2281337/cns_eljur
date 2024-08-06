<template>
    <div class="container filesPage">
        <header>
            <button>Обновить</button>
            <a href="#/home">Назад</a>
        </header>
        <table>
            <thead>
                <tr>
                    <th>Имя файла</th>
                    <th>Айди</th>
                    <th>Статус</th>
                    <th>Действия</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <th>test1</th>
                    <td>66a87970eb671d17a059bbe4</td>
                    <td>ok</td>
                    <td>
                        <button>Скачать</button>
                        <button>Удалить</button>
                        <button>Переименовать</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
  
  <style>
    @import '@/assets/css/filesPage.css';
  </style>
  
  <script lang="ts">
  import { Options, Vue } from 'vue-class-component';
  import { api } from '@/components/api/constructor'
  
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
      document.body.className = 'filesPage';
    },
    async mounted() {
      1
    },
    methods: {
      
    }
  })
  export default class filesPage extends Vue {}
  </script>