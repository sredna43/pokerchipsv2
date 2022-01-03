<script lang="ts">
	import Playertable from '../components/playertable.svelte';
    import Button from '../components/button.svelte';
    import { sResponse, myName } from '../scripts/store';
    import { sendAction } from '../scripts/ws'

	let name: string;
	myName.subscribe((n) => {
		name = n;
	});

    let displayNames: string[];
	sResponse.subscribe((r) => {
        if (r.players) {
            let names = Object.keys(r.players);
			displayNames = names.map((name): string => {
				let displayName = name;
				if (name === r.dealer) {
					displayName += ' (dealer)';
				}
                if (name === r.whose_turn) {
                    displayName = '* ' + displayName
                }
				return displayName;
			});
        }
    });

    function handleLeave(): void {
		sendAction({ action: 'remove_player', name, amount: 0 });
	}
</script>

<p>Table</p>
<p>{name}</p>
<Button text="Leave" onClick={handleLeave} />
<Playertable bind:displayNames />