<template>
    <div class="container constr">
        <div class="row constr">
            <div class="col-2 constr">
                <input v-model="searchTable" type="searchTable" class="inputS constr center" placeholder="Поиск таблиц">
                <div class="table constr" v-for="(value, key) in tables" :key="key">
                    <p v-if="value.toLowerCase().includes(searchTable.toLowerCase())||searchTable===''" :style="selectedTable === value ? 'background-color: yellow;margin-left: 6%;' : ''" @click="selectTable(value)">{{ value }}</p>
                </div>
            </div>
            <div class="col-3 constr" :class="{'loading-overlay': dataloading}">
                <div class="groupDataBtn constr">
                    <button class="dataBtn" @click="selectAll">Выделить всё</button>
                    <button class="dataBtn" @click="clearSelection">Отмена</button>
                    <button class="dataBtn" @click="refreshTable">Обновить</button>
                    <button class="dataBtn" @click="deleteTable">Удалить</button>
                </div>
                <div v-if="dataloading&&selectedTable!==''" class="loader"></div>
                <div class="table-container constr">
                    <table class="data-table">
                        <thead>
                            <tr>
                                <th v-for="(column, index) in tableSchemas[selectedTable]" :key="index">{{ column }}</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(row, index) in selectedTableData" :key="index">
                                <td v-for="(column, index1) in tableSchemas[selectedTable]" :key="index1">
                                    <input v-if="index1 === 'id'" type="checkbox" v-model="selectedDataWTb[selectedTable]" :value="row[index1]">{{ row[index1] }}
                                </td>
                                
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
            <div class="col-4 constr"></div>
        </div>
    </div>
    <div v-if="isDevMode" style="position: fixed;top:70%;left:0%;font-size: 20px;color: red;background-color: azure;">
        <p>Debug: true</p> <p>Кусочек: {{ chunk }}</p>
        <p>Кешировано в память: {{ Object.values(cacheData).map((arr1:any) => arr1.length).reduce((acc, length) => acc + length, 0) }}</p>
        <p>Записей в тек. тбл: {{ selectedTableData.length }}</p>
    </div>
</template>

<style>
    @import '@/assets/css/constructor.css';
</style>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { getTables, getList } from '@/components/api/admin';

@Options({
    data: () => {
        return {
            isDevMode: (process.env.NODE_ENV === 'development'&&localStorage.getItem('devMode')==='true')||localStorage.getItem('devModeForce') === 'true',
            chunk: 500,
            searchTable: "",
            hidetables: ["users", "roles"],
            tables: [],
            filteredTables: [],
            selectedTable: "",
            selectedTableData: [],
            dataloading: true,
            enddata: false,
            cacheData: {},
            tableSchemas: {
                "groups": {
                    id: "Айди",
                    name: "Имя",
                    max_pairs: "Кол-во пар",
                    specialization: "Специализация",
                },
                "cabinets": {
                    id: "Айди",
                    name: "Название"
                },
                "subjects": {
                    id: "Айди",
                    name: "Название"
                },
                "teachers": {
                    id: "Айди",
                    name: "Имя",
                    capacity: "Капасити"
                },
            },
            selectedDataWTb: {} as { [key: string]: number[] }
        };
    },
    components: {},
    beforeCreate: function() {
        if (document.cookie.includes('auth')) {
            document.location.href = '/home';
        }
        document.body.className = 'constr';
        getTables().then((res: { status: number; data: { tables: string[] } }) => {
            if (res.status === 200) {
                this.tables = res.data.tables;
                this.tables = this.tables.filter((value: string) => {
                    return !this.hidetables.includes(value);
                });
            }
        });
    },
    methods: {
        selectTable: function(table: string) {
            this.selectedTable = table;
            this.processGetData(table);
        },
        processGetData: async function(tableName: string) {
            if (this.cacheData[tableName]) {
                this.selectedTableData = this.cacheData[tableName];
                return;
            } else {
                this.dataloading = true;
                this.selectedTableData = [];
                let page = 1;
                const getData = async () => {
                    getList({ tablename: tableName, page: page, limit: this.chunk }).then((res: { status: number; data: any; }) => {
                        if (res.status === 200) {
                            if (res.data.table === null) {
                                this.dataloading = false;
                                this.selectedTableData = this.selectedTableData.sort((a: any, b: any) => a.id - b.id);
                                this.cacheData[tableName] = this.selectedTableData;
                            } else {
                                this.selectedTableData = this.selectedTableData.concat(res.data.table);
                                page++;
                                getData();
                            }
                        } else {
                            this.cacheData[tableName] = this.selectedTableData;
                            console.log('error');
                        }
                    });
                };
                getData();
            }
        },
        selectAll: function() {
            if (this.selectedTableData.length > 0) {
                this.selectedDataWTb[this.selectedTable] = this.selectedTableData.map((row: any) => row.id);
            }
        },
        clearSelection: function() {
            this.selectedDataWTb[this.selectedTable] = [];
        },
        refreshTable: function() {
            delete this.cacheData[this.selectedTable];
            this.processGetData(this.selectedTable);
        },
        deleteTable: function() {
            delete this.cacheData[this.selectedTable];
            this.selectedTableData = [];
            this.selectedDataWTb[this.selectedTable] = [];
            this.selectedTable = "";
        }
    }
})
export default class constructorPage extends Vue {}
</script>