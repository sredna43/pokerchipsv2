<script lang="ts">
	import Button from '../components/button.svelte';
	import Playertable from '../components/playertable.svelte';
	import { tableId, myName, sResponse, isHost } from '../scripts/store';
	import { sendAction } from '../scripts/ws';
    import { sortBy } from 'lodash'

	let tid: string;
	tableId.subscribe((v) => {
		tid = v;
	});

	let name: string;
	myName.subscribe((v) => {
		name = v;
	});

	let host = false;
    let displayNames: string[];
	sResponse.subscribe((r) => {
		if (r.players && r.players[name]?.is_host) {
			isHost.set(true);
			host = true;
		}
        if (r.players) {
			let names = Object.keys(r.players);
			displayNames = names.map((name): string => {
				let displayName = name;
				if (r.players[name]?.is_host) {
					displayName += ' (host)';
				}
				return displayName;
			});
		}
	});

    function handleStart(): void {
        sendAction({action: 'start_game', name, amount: 0});
    }

	function handleLeave(): void {
		sendAction({ action: 'remove_player', name, amount: 0 });
	}
</script>

<svelte:window on:beforeunload={handleLeave} />

<p>Waiting Room</p>
<p>Table ID: {tid}</p>
<p>You are {name}{host ? ' and you are host' : ''}</p>
{#if host}
    <Button text="Start Game" onClick={handleStart} />
{/if}
<Playertable bind:displayNames/>
<Button text="Leave" onClick={handleLeave} />
