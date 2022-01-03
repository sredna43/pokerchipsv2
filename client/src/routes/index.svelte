<script lang="ts">
	import { myName, sResponse } from '../scripts/store';
	import { dev } from '$app/env';
	import Home from '../views/home.svelte';
	import Entername from '../views/entername.svelte';
	import Waitingroom from '../views/waitingroom.svelte';
	import Table from '../views/table.svelte';
	import Button from '../components/button.svelte';
	import { disconnect } from '../scripts/ws';

	let view = 'home';
	let name: string;
	myName.subscribe((n) => {
		name = n;
	});
	sResponse.subscribe((r) => {
		let message = r.message;
		switch (message) {
			case 'CONNECTED': {
				view = 'entername';
				break;
			}
			case 'DISCONNECTED': {
				view = 'home';
				break;
			}
            case 'START_GAME': {
                view = 'table';
                break;
            }
			default: {
				break;
			}
		}
		if (name !== '' && r.requester === name) {
			if (message.includes('joined')) {
				view = 'waitingroom';
			}
			if (message.includes('left')) {
				disconnect();
			}
		}
	});

	async function handleReset() {
		const response = await fetch('http://localhost:8081/reset_test');
		const json = await response.json();
		console.log('Reset state: ', json);
	}
</script>

{#if dev}
	<Button text="reset test table" onClick={handleReset} />
{/if}

<h1>Online Poker Chips</h1>

{#if view === 'home'}
	<Home />
{:else if view === 'entername'}
	<Entername />
{:else if view === 'waitingroom'}
	<Waitingroom />
{:else if view === 'table'}
	<Table />
{/if}
