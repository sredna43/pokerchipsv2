<script lang="ts">
	import { tableId } from '../scripts/store';
	import Button from '../components/button.svelte';
	import Input from '../components/input.svelte';
	import { connect, getTableExists, getNewGame } from '../scripts/ws';

	let tid: string = '';
	let errorMessage = '';

	async function handleJoin(): Promise<void> {
		tid = tid.toLowerCase();
		tableId.set(tid);
		if (await getTableExists(tid)) {
			connect(tid);
		} else {
			errorMessage = `Table with id ${tid} does not exist`;
		}
	}

	async function handleHost(): Promise<void> {
		tid = await getNewGame();
		tableId.set(tid);
		connect(tid);
	}
</script>

<div>
	<p>{errorMessage}</p>
	<Button text="Host" onClick={handleHost} />
	<Input bind:val={tid} onEnter={handleJoin} helperText="Table ID" className="uppercase" />
	<Button text="Join" onClick={handleJoin} />
</div>
