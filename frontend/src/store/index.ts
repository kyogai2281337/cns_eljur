import { createStore } from "vuex";

interface variables {
  [key: string]: unknown;
}
interface State {
  localData: Record<string, unknown>;
  databases: { [key: string]: IDBDatabase | null };
  variables: variables;
}

const store = createStore<State>({
  state: {
    localData: {},
    databases: {},
    variables: {
      role: "",
      firstName: "",
      lastName: "",
      email: "",
    },
  },
  getters: {
    getLocalData: (state) => state.localData,
    getDatabase: (state) => (dbName: string) => state.databases[dbName] || null,
    getVariable: (state) => (name: string) => state.variables[name] || null,
  },
  mutations: {
    setVariable(state, { name, value }: { name: string; value: unknown }) {
      state.variables[name] = value;
    },
    setLocalData(state, data) {
      state.localData = data;
      localStorage.setItem("localData", JSON.stringify(data));
    },
    addDatabase(
      state,
      { dbName, dbInstance }: { dbName: string; dbInstance: IDBDatabase }
    ) {
      state.databases[dbName] = dbInstance;
    },
  },
  actions: {
    loadLocalData({ commit }) {
      const data = localStorage.getItem("localData");
      if (data) {
        commit("setLocalData", JSON.parse(data));
      }
    },
    createDatabase(
      { commit },
      { dbName, version }: { dbName: string; version: number }
    ) {
      return new Promise<IDBDatabase>((resolve, reject) => {
        const request = indexedDB.open(dbName, version);

        request.onupgradeneeded = () => {
          const db = request.result;
          if (!db.objectStoreNames.contains("defaultStore")) {
            db.createObjectStore("defaultStore", {
              keyPath: "id",
              autoIncrement: true,
            });
          }
        };

        request.onsuccess = () => {
          const db = request.result;
          commit("addDatabase", { dbName, dbInstance: db });
          resolve(db);
        };

        request.onerror = () => reject(request.error);
      });
    },
    addDataToStore(
      { state },
      {
        dbName,
        storeName,
        data,
      }: { dbName: string; storeName: string; data: Record<string, unknown> }
    ) {
      return new Promise<number | undefined>((resolve, reject) => {
        const db = state.databases[dbName];
        if (!db) {
          reject(new Error("Database not found"));
          return;
        }

        const transaction = db.transaction(storeName, "readwrite");
        const store = transaction.objectStore(storeName);
        const request = store.add(data);

        request.onsuccess = () => resolve(request.result as number);
        request.onerror = () => reject(request.error);
      });
    },
    readDataFromStore(
      { state },
      { dbName, storeName }: { dbName: string; storeName: string }
    ) {
      return new Promise<Record<string, unknown>[]>((resolve, reject) => {
        const db = state.databases[dbName];
        if (!db) {
          reject(new Error("Database not found"));
          return;
        }

        const transaction = db.transaction(storeName, "readonly");
        const store = transaction.objectStore(storeName);
        const request = store.getAll();

        request.onsuccess = () => resolve(request.result);
        request.onerror = () => reject(request.error);
      });
    },
    deleteDatabase({ commit }, dbName: string) {
      return new Promise<void>((resolve, reject) => {
        const request = indexedDB.deleteDatabase(dbName);

        request.onsuccess = () => {
          commit("addDatabase", { dbName, dbInstance: null });
          resolve();
        };

        request.onerror = () => reject(request.error);
      });
    },
  },
});

export default store;
