---

# Vuex Store для работы с `localStorage` и `IndexedDB`

Этот store позволяет управлять данными в `localStorage` и `IndexedDB` через `Vuex` для удобного доступа к состоянию и базам данных приложения.

## Интерфейсы

### State

- `localData`: хранит объект с данными, загруженными из `localStorage`.
- `databases`: хранит объект `IDBDatabase` с ключами, представляющими названия баз данных, для работы с `IndexedDB`.

## Методы `Vuex` Store

### Getters

- **`getLocalData`** - Возвращает данные из `localStorage`, хранящиеся в состоянии `Vuex`.
- **`getDatabase`** - Возвращает базу данных из состояния `Vuex` по имени базы данных.

### Mutations

- **`setLocalData`** - Устанавливает данные в `localStorage` и обновляет их в состоянии `Vuex`.
- **`addDatabase`** - Добавляет или обновляет базу данных в состоянии `Vuex`.

### Actions

- **`loadLocalData`** - Загружает данные из `localStorage` и сохраняет их в состоянии `Vuex`.
- **`createDatabase`** - Создаёт новую базу данных в `IndexedDB` с указанной версией. В базу данных добавляется хранилище `defaultStore`, если его ещё нет.
- **`addDataToStore`** - Добавляет объект данных в указанное хранилище (`object store`) базы данных.
- **`readDataFromStore`** - Считывает все данные из указанного хранилища базы данных.
- **`deleteDatabase`** - Удаляет указанную базу данных из `IndexedDB`.

## Примеры использования

### Загрузка данных из `localStorage`

Загружает данные из `localStorage` и обновляет состояние `Vuex`.

```typescript
store.dispatch("loadLocalData").then(() => {
  console.log("Данные загружены из localStorage:", store.getters.getLocalData);
});
```

### Создание базы данных

Создаёт новую базу данных с названием `TestDB` и версией `1`. Добавляет её в состояние `Vuex`.

```typescript
store.dispatch("createDatabase", { dbName: "TestDB", version: 1 })
  .then((db) => {
    console.log("База данных TestDB создана:", db);
  })
  .catch((error) => {
    console.error("Ошибка создания базы данных:", error);
  });
```

### Добавление данных в базу данных

Добавляет объект `{ name: "Антон", age: 25 }` в хранилище `defaultStore` базы данных `TestDB`.

```typescript
store.dispatch("addDataToStore", { dbName: "TestDB", storeName: "defaultStore", data: { name: "Антон", age: 25 } })
  .then((id) => {
    console.log("Данные добавлены в IndexedDB с ID:", id);
  })
  .catch((error) => {
    console.error("Ошибка добавления данных:", error);
  });
```

### Чтение данных из базы данных

Читает все данные из хранилища `defaultStore` базы данных `TestDB`.

```typescript
store.dispatch("readDataFromStore", { dbName: "TestDB", storeName: "defaultStore" })
  .then((data) => {
    console.log("Данные из IndexedDB:", data);
  })
  .catch((error) => {
    console.error("Ошибка чтения данных:", error);
  });
```

### Удаление базы данных

Удаляет базу данных `TestDB` и удаляет её экземпляр из состояния `Vuex`.

```typescript
store.dispatch("deleteDatabase", "TestDB")
  .then(() => {
    console.log("База данных TestDB удалена.");
  })
  .catch((error) => {
    console.error("Ошибка удаления базы данных:", error);
  });
```

---