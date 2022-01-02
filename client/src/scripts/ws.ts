import { browser } from '$app/env';
import { serverRes } from './appState';
import type { Request as R } from './models';

let socket: WebSocket;
const host = 'localhost:8081';

export function sendAction(req: R): void {
	if (socket.readyState === WebSocket.OPEN && browser) {
		socket.send(JSON.stringify(req));
	}
}

export function connect(tableId: string): void {
	if (browser) {
		socket = new WebSocket(`ws://${host}/ws?${tableId}`);

		socket.addEventListener('open', (e: MessageEvent) => {
			console.log('opened websocket', e);
			serverRes.set({});
		});

		socket.addEventListener('message', (e: MessageEvent) => {
			serverRes.set(JSON.parse(e.data));
		});

		socket.addEventListener('close', () => {
			serverRes.set({ message: 'DISCONNECTED' });
		});
	}
}

export function disconnect(): void {
	try {
		console.log('Closing socket connection');
		socket.close();
	} catch (error) {
		console.error(`Failed to disconnect: ${error.Message}`);
	}
}
