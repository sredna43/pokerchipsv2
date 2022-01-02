import { browser } from '$app/env';
import type { TableCheckResponse, Req, Res } from './models';
import { sResponse } from './store';

let socket: WebSocket;
const host = 'localhost:8081';

export function sendAction(r: Req): void {
	sendMessage(JSON.stringify(r));
}

export function connect(tableId: string): void {
	if (browser) {
		socket = new WebSocket(`ws://${host}/ws?${tableId}`);

		socket.addEventListener('open', (e: MessageEvent) => {
			console.log('opened websocket', e);
			parseResponse(`{ "message": "CONNECTED" }`);
		});

		socket.addEventListener('message', (e: MessageEvent) => {
			parseResponse(e.data);
		});

		socket.addEventListener('close', () => {
			parseResponse(`{ "message": "DISCONNECTED" }`);
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

export async function getTableExists(tid: string): Promise<boolean> {
	const response = await fetch(`http://${host}/table/${tid}`);
	const json: TableCheckResponse = await response.json();
    return json.exists;
}

function parseResponse(m: string): void {
	console.log(m);
	const r: Res = JSON.parse(m);
	console.log(r);
	sResponse.set(r);
}

function sendMessage(m: string): void {
	if (socket.readyState === WebSocket.OPEN && browser) {
		socket.send(m);
	}
}