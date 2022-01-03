import { writable } from 'svelte/store';
import type { Players } from './models';

export const activeView = writable('home');
export const tableId = writable('');
export const myName = writable('');
export const isHost = writable(false);
export const sResponse = writable({
    requester: '',
    players: {} as Players,
    message: '',
    pot: 0,
    whose_turn: '',
    dealer: '',
    error: false,
})