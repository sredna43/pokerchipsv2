<script lang="ts">
    import { tableId } from '../scripts/appState'
    import Button from '../components/button.svelte'
    import Input from '../components/input.svelte'
    let tid: string = ''

    function handleJoin(): void {
        tableId.set(tid)
    }
</script>

<div>
    <Button text="Host" />
    <Input bind:val={tid} />
    <Button text="Join" onClick={handleJoin} />
</div>
