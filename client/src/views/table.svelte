<script lang="ts">
	import Playertable from '../components/playertable.svelte';
	import Button from '../components/button.svelte';
	import Gameactions from '../components/gameactions.svelte';
	import { sResponse, myName } from '../scripts/store';
	import { sendAction } from '../scripts/ws';
	import type { Player } from '../scripts/models';

	let name: string;
	myName.subscribe((n) => {
		name = n;
	});

	let availableActions = 'cbf';
	let players: Player[];
	sResponse.subscribe((r) => {
		if (r.table?.players) {
			const names = Object.keys(r.table.players);
			players = names.map((name): Player => {
				return r.table.players[name];
			});
			let myPlayer = r.table?.players[name] || ({} as Player);
			if (r.table.whose_turn === myPlayer.spot) {
				var foldedPlayers = 0;
				players.forEach((p) => {
					if (p.folded) {
						foldedPlayers += 1;
					}
				});
				if (foldedPlayers === (names.length - 1) && !r.table.hand_won) {
					sendAction({ action: 'win_round', amount: 0, name })
				}
				if (r.table.can_check) {
					availableActions = 'cbf';
				} else {
					availableActions = 'crf';
				}
			} else {
				availableActions = 'nyt';
			}
		}
	});

	function handleLeave(): void {
		sendAction({ action: 'remove_player', name, amount: 0 });
	}
</script>

<p>Table</p>
<p>{name}</p>
<Gameactions type={availableActions} />
<Playertable bind:players view="table" />
<Button text="Leave" onClick={handleLeave} />
