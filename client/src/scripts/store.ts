import { writable } from 'svelte/store'
import { browser } from '$app/env'

export const messageStore = writable('')

let socket: WebSocket

const getServerAddress = (tableId: string): string => {
    return `ws://localhost:8081/ws?${tableId}`
}

const connect = (tableId): void => {
    if (browser) {
        socket = new WebSocket(getServerAddress(tableId))

        socket.addEventListener('open', (e: MessageEvent) => {
            console.log('opened websocket')
        })

        socket.addEventListener('message', (e: MessageEvent) => {
            messageStore.set(e.data)
        })
    }
}

export const sendMessage = (m): void => {
    if (socket.readyState === WebSocket.OPEN && browser) {
        socket.send(m)
    }
}