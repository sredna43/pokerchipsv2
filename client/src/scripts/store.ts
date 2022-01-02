import { writable } from 'svelte/store';

export const activeView = writable('home');
export const tableId = writable('');
export const myName = writable('');
export const sResponse = writable({
    requester: '',
    players: {},
    message: '',
    pot: 0,
    whose_turn: '',
    dealer: '',
    error: false,
})