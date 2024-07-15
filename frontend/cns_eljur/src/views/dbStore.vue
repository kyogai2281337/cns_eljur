<template>
  <main class="dbStore">
    <div class="sidebar">
      <h3 class="bar-item">Редактор 3000</h3>
      <br>
      <h3 class="bar-item">Выберите таблицу</h3>
      <div v-for="table in filteredTables" :key="table" @click="selectTable(table)">
        <a class="bar-item"><h3>{{ table }}</h3></a>
      </div>
    </div>
    <div class="content">
      <div class="container" style="background-color: whitesmoke;">
        <h1>База данных</h1>
      </div>
      <div class="container">
        <h2 style="left: 10px;">Таблица: {{ curTable ? curTable : 'не открыта' }}</h2>
        <table class="database-table" v-if="curTable">
          <thead>
            <tr>
              <th>Записи</th>
              <th>Развернутые данные</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>
                <input type="text" v-model="searchTable" placeholder="Поиск по записям" />
                <ul>
                  <li v-for="object in filteredObjects" :key="object.id" @click="selectObject(object)" :style="selectedObject === object.id ? 'background-color: yellow;' : ''">
                    <div v-for="(value, key) in object" :key="key">
                      {{ key }}: {{ value }}
                    </div>
                  </li>
                </ul>
              </td>
              <td class="fixed-data">
                <div v-if="object">
                  <table>
                    <thead>
                      <tr>
                        <th>Ключ</th>
                        <th>Значение</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="(value, key) in object" :key="key" @click="selectLineF(key)">
                        <td>{{ key }}</td>
                        <td>
                          <template v-if="key === selectLine">
                            <input type="text" v-model="object[key]" />
                          </template>
                          <template v-else>
                            {{ value }}
                          </template>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                  <div v-if="selectLine !== ''">
                    <button @click="saveChanges">Сохранить</button>
                    <button @click="cancelChanges">Отменить</button>
                  </div>
                </div>
                <div v-else>
                  <p>Выберите запись для отображения данных</p>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<script>
import { getTables, getList, getObj } from '@/components/api/admin';

export default {
  name: 'dbStore',
  beforeCreate: function() {
    document.body.className = 'dbStore';
  },
  data() {
    return {
      curTable: null,
      tables: null,
      object: null,
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
      noDataMessage: 'Пусто',
      selectLine: '',
      originalValue: '',
      noEdit:['id','role','deleted']
    };
  },
  async mounted() {
    await this.gettbls();
  },
  methods: {
    async gettbls() {
      const res = await getTables();
      if (res.status === 200) {
        this.tables = res.data.tables;
        this.filteredTables = this.tables;
      }
    },
    async selectTable(table) {
      this.curTable = table;
      const res = await getList({"tname": this.curTable, "limit": 25, "page": this.page});
      if (res.status === 200) {
        this.objects = res.data.table;
        this.filteredObjects = this.objects;
      }
    },
    async selectObject(object) {
      this.selectedObject = object.id;
      const res = await getObj({"tname": this.curTable, "id": object.id});
      this.object = res.data;
    },
    selectLineF(key) {
      if (this.noEdit.includes(key)) {
        return;
      } else {
        this.selectLine = key;
        this.originalValue = this.object[key];
      }
    },
    cancelChanges() {
      if (confirm('Отменить изменения?')) {
        this.object[this.selectLine] = this.originalValue;
        this.selectLine = '';
      }
    },
    saveChanges() {
      if (confirm('Сохранить изменения?')) {
        // Логика сохранения изменений на сервере
        this.selectLine = '';
      }
    }
  }
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100..900;1,100..900&display=swap');
@import '@/assets/css/dbStore.css';
</style>
