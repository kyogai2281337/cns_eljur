interface Item {
    id?: number;
    value: string;
  }
  
  const dbName = 'constuctorDB';
  
  let db: IDBDatabase | null = null;
  
  const openDB = (storeName: string): Promise<IDBDatabase> => {
    return new Promise((resolve, reject) => {
      const request = indexedDB.open(dbName, 1);
  
      request.onupgradeneeded = (event) => {
        const db = (event.target as IDBOpenDBRequest).result;
        if (!db.objectStoreNames.contains(storeName)) {
          db.createObjectStore(storeName, { keyPath: 'id', autoIncrement: true });
        }
      };
  
      request.onsuccess = (event) => {
        db = (event.target as IDBOpenDBRequest).result;
        resolve(db);
      };
  
      request.onerror = (event) => {
        reject((event.target as IDBOpenDBRequest).error);
      };
    });
  };
  
  export const setData = async (storeName: string, value: string): Promise<void> => {
    if (!db) {
      db = await openDB(storeName);
    }
  
    const transaction = db.transaction([storeName], 'readwrite');
    const store = transaction.objectStore(storeName);
    const request = store.add({ value });
  
    return new Promise((resolve, reject) => {
      request.onsuccess = () => resolve();
      request.onerror = () => reject(request.error);
    });
  };
  
  export const getData = async (storeName: string): Promise<Item[]> => {
    if (!db) {
      db = await openDB(storeName);
    }
  
    const transaction = db.transaction([storeName], 'readonly');
    const store = transaction.objectStore(storeName);
    const request = store.getAll();
  
    return new Promise((resolve, reject) => {
      request.onsuccess = () => resolve(request.result);
      request.onerror = () => reject(request.error);
    });
  };
  
  export const updateData = async (storeName: string, id: number, value: string): Promise<void> => {
    if (!db) {
      db = await openDB(storeName);
    }
  
    const transaction = db.transaction([storeName], 'readwrite');
    const store = transaction.objectStore(storeName);
    const request = store.put({ id, value });
  
    return new Promise((resolve, reject) => {
      request.onsuccess = () => resolve();
      request.onerror = () => reject(request.error);
    });
  };