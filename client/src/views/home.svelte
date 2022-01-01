<script lang="ts">
    import { tableId } from '../scripts/appState'
    let tid: string = ''

    function handleJoin(): void {
        console.log("Joining table " + tid)
        tableId.set(tid)
    }
</script>

<div>
    <button>host</button>
    <input bind:value={tid}>
    <button on:click={handleJoin}>join</button>
</div>
