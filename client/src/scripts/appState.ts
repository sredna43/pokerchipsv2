import { writable } from 'svelte/store';

export const activeView = writable('home');
export const tableId = writable('');
export const serverRes = writable({});
export const myName = writable('');