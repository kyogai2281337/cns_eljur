// src/components/api/constructor.ts

const baseURL = '/api';

const api = {
    async getConstructor(id: string) {
        const response = await fetch(`${baseURL}/constructor/private/constructor/get`, {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({id}),
        });
        return response.json();
    },

    async deleteConstructor(id: string) {
        const response = await fetch(`${baseURL}/constructor/private/constructor/delete`, {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({id}),
        });
        return response.json();
    },

    async createConstructor(data: object) {
        const response = await fetch(`${baseURL}/constructor/private/constructor/create`, {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
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
            },
            body: JSON.stringify({id, name}),
        });
        return response.json();
    },

    async saveConstructor(id: string) {
        const response = await fetch(`${baseURL}/constructor/private/constructor/save`, {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({id}),
        });
        return response.json();
    },

    async loadConstructor() {
        const response = await fetch(`${baseURL}/constructor/private/constructor/load/66a87970eb671d17a059bbe4`, {
            method: 'GET',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
        });
        return response.json();
    },
};

export {api};
