<script lang="ts">
	import { tableId } from '../scripts/store';
	import Button from '../components/button.svelte';
	import Input from '../components/input.svelte';
	import { connect, getTableExists } from '../scripts/ws';
    
	let tid: string = '';
    let errorMessage = '';

	async function handleJoin(): Promise<void> {
		tableId.set(tid);
        if (await getTableExists(tid)) {
            connect(tid);
        } else {
            errorMessage = `Table with id ${tid} does not exist`
        }
	}
</script>

<div>
    <p>{errorMessage}</p>
	<Button text="Host" />
	<Input bind:val={tid} onEnter={handleJoin} helperText="Enter Table ID" />
	<Button text="Join" onClick={handleJoin} />
</div>
