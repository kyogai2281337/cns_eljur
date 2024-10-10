<template>
  <div class="container filesPage">
    <header class="filesPage-header filesPage">
      <button class="back-button filesPage" @click="homeBtnF">Назад</button>
      <div class="username filesPage" @click="homeBtnF">{{ username }}</div>
    </header>
    <div class="row filesPage">
      <div class="table-files col-2 filesPage">
        <input v-model="searchSch" type="searchTable" class="inputS constr center" placeholder="Поиск расписаний" style="padding: 5px 5%;">
        <div class="table-files-list" v-for="(value, key) in scheduledList" :key="key">
          <p :style="selectedSch === key ? 'background-color: yellow;margin-left: 3%;' : ''" @click="selectSch(key, value)">
            {{key}}
          </p>
        </div>
      </div>
      <div class="files-data col-4 filesPage" :style="selectedSch === '' ? 'background-color: #333333' : ''">
        <header v-if="selectedSch !== ''" style="margin: 3%;margin-top: -1%;display:block;">
          <p>Расписание: {{ schData.name }}</p>
        </header>
        <div v-if="selectedSch !== ''" style="margin: 3%;">
          <p>Сводка:</p>
          <ul>
            <li>Дней: {{ schData.days }}</li>
            <li>Пар: {{ schData.pairs }}</li>
            <li>Лимит недель: {{ schData.weeklimit }}</li>
            <li>Лимит дней: {{ schData.daylimit }}</li>
            <li>Групп: {{ groups.length }}</li>
            <li>Преподавателей: {{ teachers.length }}</li>
            <li>Кабинетов: {{ cabinets.length }}</li>
          </ul>
        </div>
        <footer v-if="selectedSch !== ''" style="margin: 3%;top:80%;position: absolute;">
          <button class="actionSchB" @click="saveFile(schId)">Сохранить</button>
          <button class="actionSchB" @click="deleteSch(schId)">Удалить</button>
          <button class="actionSchB" @click="downloadFile(schId)">Скачать</button>
        </footer>
      </div>
    </div>
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
        role: '',
        scheduledList: [],
        searchSch: '',
        selectedSch: '',
        schData: {},
        schId: '',
        groups: [],
        teachers: [],
        cabinets: [],
      };
    },
    components: {},
    beforeCreate() {
      if (document.cookie.includes('auth')) {
        document.location.href = '#';
      }
      document.body.className = 'filesPage';
    },
    async mounted() {
      this.refreshList()
      try {
        this.parseProfile();
      } catch (error) {
        console.log(error)
      }
    },
    methods: {
      homeBtnF() {
        document.location.href = '#/home';
      },
      async refreshList() {
        const data = await api.getConstructorList();
        this.scheduledList = data.schedules || {};
      },
      async selectSch(sch: string, id: string) {
        this.schId = id;
        if (this.selectedSch === sch) {
          this.selectedSch = '';
        } else {
          this.selectedSch = sch;
          await api.getConstructor(id).then(res => {
            console.log(res)
            this.groups = res.schedule.groups || [];
            this.teachers = res.schedule.teachers || [];
            this.cabinets = res.schedule.cabinets || [];
            this.schData = res.schedule
          }).catch(err => {
            console.log(err)
            return {}
          });
        }
      },
      parseProfile() {
        let userData = localStorage.getItem('profile') || '{}';
        let parsedUserData = JSON.parse(userData) as { first_name?: string, last_name?: string, role?: string };
        this.username = (parsedUserData.first_name || '') + ' ' + (parsedUserData.last_name || '');
        this.role = parsedUserData.role
      },
      async saveFile(fileId: string) {
        const response = await api.saveConstructor(fileId);
        if (response.status !== 200) {
          alert('Ошибка при сохранении файла');
          console.log(response);
          return;
        } else if (response.status === 200) {
          alert('Файл успешно сохранен');
          return;
        } else {
          alert('Ошибка при сохранении файла');
          console.log(response);
          return;
        }
      },
      async downloadFile(fileId: string) {
        try {
          const response = await api.loadConstructor(fileId);
          if (response.status === 404) {
            alert('Файл не найден, сначала сохраните его');
            return;
          }
          const blob = await response.blob();
          const filename = this.selectedSch + '.xlsx';
          var downloadLink = window.document.createElement('a');
          downloadLink.href = window.URL.createObjectURL(new Blob([blob], {type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}));
          downloadLink.download = filename;
          document.body.appendChild(downloadLink);
          downloadLink.click();
          document.body.removeChild(downloadLink);
        } catch (error) {
          console.error('Error downloading file:', error);
        }
      },
      async deleteSch(fileId: string) {
        try {
          if (confirm("Вы уверены, что хотите удалить расписание?")) {
            await api.deleteConstructor(fileId);
            this.refreshList();
          }
        } catch (error) {
          console.error('Error deleting file:', error);
        }
      },
    }
  })
  export default class filesPage extends Vue {}
</script>