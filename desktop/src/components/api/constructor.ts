// src/components/api/constructor.ts

const baseURL = 'https://localhost/api';

const api = {
  async getConstructor(id: string) {
    const response = await fetch(`${baseURL}/constructor/private/constructor/get`, {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        'AuthToken': localStorage.getItem('token') || '',
      },
      body: JSON.stringify({ id }),
    });
    return response.json();
  },

  async deleteConstructor(id: string) {
    const response = await fetch(`${baseURL}/constructor/private/constructor/delete`, {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        'AuthToken': localStorage.getItem('token') || '',
      },
      body: JSON.stringify({ id }),
    });
    return response.json();
  },

  async createConstructor(data: object) {
    const response = await fetch(`${baseURL}/constructor/private/constructor/create`, {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        'AuthToken': localStorage.getItem('token') || '',
      },
      body: JSON.stringify(data),
    });
    return response.json();
  },

  async getConstructorList() {
    const response = await fetch(`${baseURL}/constructor/private/constructor/getlist`, {
      method: 'GET',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        'AuthToken': localStorage.getItem('token') || '',
      },
    });
    return response.json();
  },

  async renameConstructor(id: string, name: string) {
    const response = await fetch(`${baseURL}/constructor/private/constructor/rename`, {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        'AuthToken': localStorage.getItem('token') || '',
      },
      body: JSON.stringify({ id, name }),
    });
    return response.json();
  },

  async saveConstructor(id: string) {
    const response = await fetch(`${baseURL}/constructor/private/constructor/save`, {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        'AuthToken': localStorage.getItem('token') || '',
      },
      body: JSON.stringify({ id }),
    });
    return response.json();
  },

  async loadConstructor(id: string) {
    const response = await fetch(`${baseURL}/constructor/private/constructor/load/`+id, {
      method: 'GET',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        'AuthToken': localStorage.getItem('token') || '',
      },
    });
    return response
  },
};

export { api };
