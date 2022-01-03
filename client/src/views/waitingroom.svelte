<script lang="ts">
	import Button from '../components/button.svelte';
	import Playertable from '../components/playertable.svelte';
	import { tableId, myName, sResponse, isHost } from '../scripts/store';
	import { sendAction } from '../scripts/ws';
	import type { Player } from 'src/scripts/models';
	import Input from '../components/input.svelte';

	let tid: string;
	tableId.subscribe((v) => {
		tid = v;
	});

	let name: string;
	myName.subscribe((v) => {
		name = v;
	});

	let initialChips: string;
	let host = false;
	let players: Player[];
	let dealer: string;
	sResponse.subscribe((r) => {
		if (r.players && r.players[name]?.is_host) {
			isHost.set(true);
			host = true;
		}
		if (r.players) {
			const names = Object.keys(r.players);
			players = names.map((name): Player => {
				return r.players[name];
			})
		}
		dealer = r.dealer;
	});

	function handleStart(): void {
		if (dealer !== '' && initialChips !== '') {
			sendAction({ action: 'start_game', name, amount: 0 });
		} else {
			console.log("Need to specify first dealer and initial chips")
		}
	}

	function handleInitialChips(): void {
		if (Number(initialChips) > 0){
			sendAction({ action: 'set_initial_chips', name, amount: Number(initialChips)})
		}
	}

	function handleLeave(): void {
		sendAction({ action: 'remove_player', name, amount: 0 });
	}
</script>

<svelte:window on:beforeunload={handleLeave} />

<p>Waiting Room</p>
<p>Table ID: {tid}</p>
<p>You are {name}{host ? ' and you are host' : ', waiting for host to start game'}</p>
{#if host}
	<Button text="Start Game" onClick={handleStart} />
	<Button text="Set Initial Chips" onClick={handleInitialChips} />
	<Input bind:val={initialChips} helperText="Initial chips" />
{/if}
<Playertable bind:players mutable={host}/>
<Button text="Leave" onClick={handleLeave} />
