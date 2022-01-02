<script lang="ts">
	import Button from '../components/button.svelte';
    import Playertable from '../components/playertable.svelte';
	import { tableId, myName, sResponse } from '../scripts/store';
	import { disconnect, sendAction } from '../scripts/ws';

	let tid: string;
	tableId.subscribe((v) => {
		tid = v;
	});

	let name: string;
	myName.subscribe((v) => {
		name = v;
	});

	let players: string[];
	sResponse.subscribe((r) => {
        if (r.players) {
            players = Object.keys(r.players);
        }
	});

	function handleLeave(): void {
        sendAction({action: 'remove_player', name, amount: 0});
	}
</script>

<p>Waiting Room</p>
<p>Table ID: {tid}</p>
<p>You are: {name}</p>
<Playertable {players} />
<Button text="Leave" onClick={handleLeave} />
