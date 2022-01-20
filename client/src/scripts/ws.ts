import { browser } from '$app/env';
import type { TableCheckResponse, Req, Res, NewTableResponse } from './models';
import { sResponse } from './store';

let socket: WebSocket;
const host = '192.168.0.12:8081';

export function sendAction(r: Req): void {
	sendMessage(JSON.stringify(r));
}

export function connect(tableId: string): void {
	if (browser) {
		socket = new WebSocket(`ws://${host}/ws?${tableId}`);

		socket.addEventListener('open', () => {
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

export async function getNewGame(): Promise<string> {
	const response = await fetch(`http://${host}/new_game`);
	const json: NewTableResponse = await response.json();
	return json.id;
}

function parseResponse(m: string): void {
	console.log(m);
	const r: Res = JSON.parse(m);
	console.log(r);
	sResponse.set(r);
}

function sendMessage(m: string): void {
	if (socket.readyState === WebSocket.OPEN && browser) {
		console.log(m);
		socket.send(m);
	}
}