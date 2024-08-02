<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<template>
    <div class="container constr">
      <div class="row constr">
        <div class="col-2 constr">
          <input v-model="searchTable" type="searchTable" class="inputS constr center" placeholder="Поиск таблиц">
          <div class="table constr" v-for="(value, key) in filteredTables" :key="key">
            <p :style="selectedTable === value ? 'background-color: yellow;margin-left: 6%;' : ''" @click="selectTable(value)">
              {{ value }}
            </p>
          </div>
        </div>
        <div class="col-3 constr" :class="{'loading-overlay': dataloading||selectedTable === ''}">
          <div class="groupDataBtn constr">
            <button class="dataBtn" @click="selectAll">Выделить всё</button>
            <button class="dataBtn" @click="clearSelection">Отмена</button>
            <button class="dataBtn" @click="refreshTable">Обновить</button>
            <button class="dataBtn" @click="deleteTable">Удалить</button>
          </div>
          <div v-if="dataloading && selectedTable !== ''" class="loader"></div>
          <div class="table-container constr">
            <table class="data-table">
              <thead>
                <tr>
                  <th v-for="(column, index) in tableSchemas[selectedTable]" :key="index" @click="openSearchMenu(index);" @mouseleave="searchColumn=''">
                    {{ column }}
                    <div v-if="index === searchColumn" class="search-menu">
                        <input class="search-menu-input" v-model="searchValue" @input="applyColumnFilter(index)" placeholder="Введите значение">
                    </div>
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(row, index) in filteredTableData" :key="index">
                  <td v-for="(column, index1) in tableSchemas[selectedTable]" :key="index1">
                    <input v-if="index1 === 'id'" type="checkbox" v-model="selectedDataWTb[selectedTable]" :value="row[index1]">{{ row[index1] }}
                  </td>
                </tr>
                <tr ref="tableEnd"></tr>
              </tbody>
            </table>
          </div>
        </div>
        <div class="col-4 constr">
            <button @click="console.log(selectedDataWTb)">Get Data</button>
            <button @click="createSch()">Создать расписание</button>
            <input v-model="schName" placeholder="название">
            <input v-model="max_weeks" placeholder="максимальное кол-во недель">
            <input v-model="max_days" placeholder="максимальное кол-во дней в неделе">
            <input v-model="days" placeholder="кол-во дней в расписании">
            <input v-model="pairs" placeholder="кол-во пар в расписании">
            <input v-model="results">
            <p>
              Групп: {{ selectedDataWTb.groups?.length }}
              Планов: {{ selectedDataWTb.subjects?.length }}
              Кабинетов: {{ selectedDataWTb.cabinets?.length }}
              Учителейй: {{ selectedDataWTb.teachers?.length }}
            </p>
        </div>
      </div>
    </div>
    <div v-if="isDevMode" style="position: fixed;top:70%;left:0%;font-size: 20px;color: red;background-color: azure;">
      <p>Debug: true</p>
      <p>Кусочек: {{ chunk }}</p>
      <p>Записей в тек. тбл: {{ selectedTableData.length }}</p>
    </div>
  </template>
  
  <style>
  @import '@/assets/css/constructor.css';
  </style>
  
  <script lang="ts">
  import { Options, Vue } from 'vue-class-component';
  import { openDB } from 'idb';
  import { getTables, getList } from '@/components/api/admin';
  import { api } from '@/components/api/constructor';
  
  @Options({
    data: () => ({
      isDevMode: (process.env.NODE_ENV === 'development' && localStorage.getItem('devMode') === 'true') || localStorage.getItem('devModeForce') === 'true',
      chunk: 500,
      searchTable: "",
      hidetables: ["users", "roles"],
      tables: [] as string[],
      selectedTable: "",
      selectedTableData: [],
      visibleTableData: [],
      filteredTableData: [],
      dataloading: false,
      searchColumn: "",
      searchValue: "",
      tableSchemas: Object.freeze({
        "groups": { id: "Айди", name: "Имя", max_pairs: "Кол-во пар", specialization: "Специализация" },
        "cabinets": { id: "Айди", name: "Название" },
        "subjects": { id: "Айди", name: "Название" },
        "teachers": { id: "Айди", name: "Имя", capacity: "Капасити" },
      }),
      selectedDataWTb: {} as { [key: string]: number[] },
      db: null as any,
      schName:'exampletest1',
      max_weeks:18,
      max_days:4,
      days:6,
      pairs:6,
      results:''
    }),
    async beforeCreate() {
      if (document.cookie.includes('auth')) {
        document.location.href = '/home';
      }
      document.body.className = 'constr';
      this.db = await openDB('tablesDB', 1, {
        upgrade(db) {
          db.createObjectStore('tables');
        },
      });
      const storedTables = await this.db.get('tables', 'tables');
      if (storedTables) {
        this.tables = storedTables.filter((value: string) => !this.hidetables.includes(value));
      } else {
        const res = await getTables();
        if (res.status === 200) {
          this.tables = res.data.tables.filter((value: string) => !this.hidetables.includes(value));
          await this.db.put('tables', JSON.parse(JSON.stringify(this.tables)), 'tables');
        }
      }
    },
    computed: {
      filteredTables() {
        return this.tables.filter((value: string) => value.toLowerCase().includes(this.searchTable.toLowerCase()) || this.searchTable === '');
      }
    },
    methods: {
      async selectTable(table: string) {
        this.selectedTable = table;
        await this.processGetData(table);
      },
      async processGetData(tableName: string) {
        this.dataloading = true;
        const cachedData = await this.db.get('tables', tableName);
        if (cachedData) {
          this.selectedTableData = cachedData;
          this.visibleTableData = this.selectedTableData.slice(0, 100);
          this.filteredTableData = this.visibleTableData;
          this.dataloading = false;
        } else {
          this.selectedTableData = [];
          let page = 1;
          const getData = async () => {
            const res = await getList({ tablename: tableName, page, limit: this.chunk });
            if (res.status === 200) {
              if (res.data.table === null) {
                this.selectedTableData = this.selectedTableData.sort((a: any, b: any) => a.id - b.id);
                await this.db.put('tables', JSON.parse(JSON.stringify(this.selectedTableData)), tableName);
                this.visibleTableData = this.selectedTableData.slice(0, 100);
                this.filteredTableData = this.visibleTableData;
                this.dataloading = false;
              } else {
                this.selectedTableData = this.selectedTableData.concat(res.data.table);
                page++;
                getData();
              }
            } else {
              this.dataloading = false;
              console.log('error');
            }
          };
          await getData();
        }
      },
      setupIntersectionObserver() {
        const observer = new IntersectionObserver((entries) => {
          entries.forEach(entry => {
            if (entry.isIntersecting) {
              this.visibleTableData = this.selectedTableData.slice(0, this.visibleTableData.length + 100);
              this.applyColumnFilter(this.searchColumn);
            }
          });
        }, { threshold: 1.0 });
      
        observer.observe(this.$refs.tableEnd);
      },
      openSearchMenu(column: string) {
        this.searchColumn = column;
        this.searchValue = "";
      },
      applyColumnFilter(column: string) {
        if (this.searchValue) {
            console.log(this.searchValue,column);
          this.filteredTableData = this.selectedTableData.filter((row: any) => row[column] && row[column].toString().toLowerCase().includes(this.searchValue.toLowerCase()));
        } else {
          this.filteredTableData = this.visibleTableData;
        }
      },
      async selectAll() {
        if (this.selectedTableData.length > 0) {
          this.selectedDataWTb[this.selectedTable] = this.selectedTableData.map((row: any) => row.id);
        }
      },
      async clearSelection() {
        this.selectedDataWTb[this.selectedTable] = [];
      },
      async refreshTable() {
        await this.db.delete('tables', this.selectedTable);
        await this.processGetData(this.selectedTable);
      },
      async deleteTable() {
        await this.db.delete('tables', this.selectedTable);
        this.selectedTableData = [];
        this.selectedDataWTb[this.selectedTable] = [];
        this.selectedTable = "";
      },
      async createSch() {
        const newConstructor = await api.createConstructor({
    "name": this.schName,
    "limits": {
        "max_weeks": this.max_weeks,
        "max_days": this.max_days,
        "days": this.days,
        "pairs": this.pairs,
    },
    "groups": this.selectedDataWTb.groups,
    "plans": this.selectedDataWTb.subjects,
    "cabinets": this.selectedDataWTb.cabinets,
    "teachers": this.selectedDataWTb.teachers,
})
        this.results=newConstructor
      }
    },
    mounted() {
      this.setupIntersectionObserver();
    }
  })
  export default class constructorPage extends Vue {}
  </script>
  