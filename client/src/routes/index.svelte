<script lang="ts">
import { onMount } from 'svelte';

    import { connect, disconnect, sendMessage, messageStore } from '../scripts/ws'
    let title = "Online Poker Chips"
    let isConnected = false
    let connectButtonText = 'connect'
    let tableId = ''
    let m = ''
    let serverData = ''

    onMount(() => {
        messageStore.subscribe((data) => {
            serverData = data
        })
    })

    const handleConnection = () => {
        if (!isConnected) {
            connect(tableId)
        } else {
            disconnect()
        }
        isConnected = !isConnected
        connectButtonText = isConnected ? 'disconnect' : 'connect'
    }

    const handleSendMessage = () => {
        sendMessage(m)
    }
</script>

<h1>{title}</h1>
<p>Table Code: </p>
{#if !isConnected}
    <input bind:value={tableId}>
{:else}
    <p>{tableId}</p>
    <p>Send message: </p>
    <input bind:value={m}>
    <button on:click={handleSendMessage}>Send</button>
    <p>Message from server: {serverData}</p>
{/if}
<button on:click={handleConnection}>{connectButtonText}</button>