<script lang="ts">
	import Playertable from '../components/playertable.svelte';
	import Button from '../components/button.svelte';
	import { sResponse, myName } from '../scripts/store';
	import { sendAction } from '../scripts/ws';
	import type { Player } from 'src/scripts/models';

	let name: string;
	myName.subscribe((n) => {
		name = n;
	});

	let players: Player[];
	sResponse.subscribe((r) => {
		if (r.players) {
			const names = Object.keys(r.players);
			players = names.map((name): Player => {
				return r.players[name];
			})
		}
	});

	function handleLeave(): void {
		sendAction({ action: 'remove_player', name, amount: 0 });
	}
</script>

<p>Table</p>
<p>{name}</p>
<Button text="Leave" onClick={handleLeave} />
<Playertable bind:players view="table" />
