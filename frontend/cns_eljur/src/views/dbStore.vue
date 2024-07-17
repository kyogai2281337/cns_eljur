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
      <div class="container fixed-header">
        <h1>База данных</h1>
      </div>
      <div class="container">
        <h2 style="left: 10px;">Таблица: {{ curTable ? curTable : 'не открыта' }}</h2>
        <button @click="openModal">Создать запись</button>
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

    <div v-if="showModal" class="modal">
      <div class="modal-content">
        <span class="close" @click="closeModal">&times;</span>
        <h2>Создание новой записи в таблице {{ curTable }}</h2>
        <form @submit.prevent="createRecord">
          <div v-for="(value, key) in tableSchema[curTable]" :key="key">
            <label :for="key">{{ key }}</label>
            <input type="text" v-model="newRecord[key]" :id="key" />
          </div>
          <button type="submit">Создать</button>
        </form>
      </div>
    </div>
  </main>
</template>

<script>
import { getTables, getList, getObj, create } from '@/components/api/admin';

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
      noEdit: ['id', 'role', 'deleted'],
      showModal: false,
      newRecord: {},
      tableSchema: {
        users: {
          id: 'number',
          email: 'string',
          first_name: 'string',
          last_name: 'string',
          role: 'object',
          deleted: 'boolean'
        },
      }
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
      const res = await getList({ tablename: this.curTable, limit: 25, page: this.page });
      if (res.status === 200) {
        this.objects = res.data.table;
        this.filteredObjects = this.objects;
      }
    },
    async selectObject(object) {
      this.selectedObject = object.id;
      const res = await getObj({ tablename: this.curTable, id: object.id });
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
        this.selectLine = '';
      }
    },
    openModal() {
      this.newRecord = {};
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
    },
    async createRecord() {
      const res = await create({ Table: this.curTable, Data: this.newRecord });
      if (res.status === 200) {
        this.selectTable(this.curTable);
        this.closeModal();
      }
    }
  }
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100..900;1,100..900&display=swap');
@import '@/assets/css/dbStore.css';

.fixed-header {
  position: sticky;
  top: 0;
  background-color: whitesmoke;
  z-index: 2;
}

.modal {
  display: block;
  position: fixed;
  z-index: 3;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0,0,0);
  background-color: rgba(0,0,0,0.4);
}

.modal-content {
  background-color: #fefefe;
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}
</style>
