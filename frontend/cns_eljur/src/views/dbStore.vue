<template>
  <div class="header" id="app">
    <div class="admin">
      <p class="admin-text">Панель администратора</p>
    </div>
    <div class="admin-rectangle1">
      <input type="text" v-model="searchTable" placeholder="Поиск по таблицам" @input="filterTables">
      <ul>
        <li v-for="table in filteredTables" :key="table" @click="selectTable(table)">
          {{ table }}
        </li>
      </ul>
    </div>
    <div class="admin-rectangle2" @scroll="handleScroll">
      <input type="text" v-model="searchObj" placeholder="Поиск по объектам" @input="filterObjects">
      <ul>
        <li v-for="obj in filteredObjects" :key="obj.id" @click="selectObject(obj)">
          {{ obj.email }}
          <div v-if="selectedObject && selectedObject.id === obj.id">
            <p>ID: {{ selectedObject.id }}</p>
            <p>Email: {{ selectedObject.email }}</p>
            <p>Имя: {{ selectedObject.first_name }}</p>
            <p>Фамилия: {{ selectedObject.last_name }}</p>
            <p>Роль: {{ selectedObject.role ? selectedObject.role.name : 'Нет роли' }}</p>
          </div>
        </li>
      </ul>
    </div>
    <a class="navigation">1-25</a>
    <img src="@/assets/images/arrow-right.png" class="admin-img" />
  </div>
</template>

<script>
import { getTables, getList, getObj } from '@/components/api/admin';

export default {
  name: 'dbStore',
  data() {
    return {
      tables: null,
      filteredTables: [],
      selectedTable: null,
      searchTable: '',
      objects: [],
      filteredObjects: [],
      searchObj: '',
      selectedObject: null,
      page: 1,
      loading: false,
      endOfList: false,
      noDataMessage: 'Пусто'
    };
  },
  async mounted() {
    await this.listTables();
  },
  methods: {
    async listTables() {
      const responseObj = await getTables();
      if (responseObj.status === 200 && responseObj.data) {
        this.tables = responseObj.data.tables;
        this.filteredTables = this.tables;
      } else {
        this.tables = [];
        this.filteredTables = [];
      }
    },
    filterTables() {
      this.filteredTables = this.tables.filter(table =>
        table.toLowerCase().includes(this.searchTable.toLowerCase())
      );
    },
    async selectTable(table) {
      this.selectedTable = table;
      this.page = 1;
      this.objects = [];
      this.endOfList = false;
      await this.loadObjects();
    },
    filterObjects() {
      this.filteredObjects = this.objects.filter(obj =>
        obj.email.toLowerCase().includes(this.searchObj.toLowerCase())
      );
    },
    async loadObjects() {
      if (this.loading || this.endOfList) return;
      this.loading = true;
      const responseObj = await getList({
        tname: this.selectedTable,
        limit: 25,
        page: this.page
      });
      if (responseObj.status === 200 && responseObj.data && responseObj.data.table) {
        if (responseObj.data.table.length) {
          this.objects.push(...responseObj.data.table);
          this.filteredObjects = this.objects;
          this.page += 1;
        } else {
          this.endOfList = true;
        }
      } else {
        this.endOfList = true;
      }
      this.loading = false;
    },
    async handleScroll(event) {
      const bottom = event.target.scrollHeight - event.target.scrollTop === event.target.clientHeight;
      if (bottom) {
        await this.loadObjects();
      }
    },
    async selectObject(obj) {
      const responseObj = await getObj({
        tname: this.selectedTable,
        id: obj.id
      });
      if (responseObj.status === 200 && responseObj.data) {
        this.selectedObject = responseObj.data;
      } else {
        this.selectedObject = {
          id: obj.id,
          email: obj.email,
          first_name: 'Нет данных',
          last_name: 'Нет данных',
          role: { name: 'Нет роли' }
        };
      }
    }
  }
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100..900;1,100..900&display=swap');
@import '@/assets/css/dbStore.css';

.admin-rectangle1 {
  overflow-y: auto;
}

.admin-rectangle2 {
  overflow-y: auto;
}
</style>
