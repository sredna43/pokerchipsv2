import { writable } from 'svelte/store';
import type { Res } from './models';

export const activeView = writable('home');
export const tableId = writable('');
export const myName = writable('');
export const isHost = writable(false);
export const sResponse = writable({} as Res);