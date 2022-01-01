import { writable } from 'svelte/store'
import { browser } from '$app/env'

export const messageStore = writable('')

let socket: WebSocket

const getServerAddress = (tableId: string): string => {
    return `ws://localhost:8081/ws?${tableId}`
}

export const connect = (tableId): void => {
    if (browser) {
        socket = new WebSocket(getServerAddress(tableId))

        socket.addEventListener('open', (e: MessageEvent) => {
            console.log('opened websocket')
            messageStore.set('')
        })

        socket.addEventListener('message', (e: MessageEvent) => {
            messageStore.set(e.data)
        })

        socket.addEventListener('close', () => {
            messageStore.set('DISCONNECTED')
        })
    }
}

export const disconnect = (): void => {
    try {
        console.log('Closing socket connection')
        socket.close()
    } catch (error) {
        console.error(`Failed to disconnect: ${error.Message}`)
    }
}

export const sendMessage = (m): void => {
    if (socket.readyState === WebSocket.OPEN && browser) {
        socket.send(m)
    }
}