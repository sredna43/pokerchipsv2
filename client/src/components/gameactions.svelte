<script lang="ts">
	import { sendAction } from '../scripts/ws';
	import Button from '../components/button.svelte';
	import { myName, sResponse, isHost } from '../scripts/store';
	export let type = 'cbf';

	let amount = 0;
	let name: string;
	let message: string;
	let host = false;

	myName.subscribe((n) => {
		name = n;
	});

	sResponse.subscribe((r) => {
		message = r.message;
	})

	isHost.subscribe((h) => {
		host = h;
	})

	function check() {
		sendAction({ action: 'check', amount: 0, name });
	}

	function bet() {
		sendAction({ action: 'bet', amount, name });
	}

	function call() {
		sendAction({ action: 'call', amount: 0, name });
	}

	function raise() {
		sendAction({ action: 'raise', amount, name });
	}

    function fold() {
        sendAction({ action: 'fold', amount: 0, name });
    }

	function nextRound() {
		sendAction({ action: 'new_round', amount: 0, name})
	}
</script>

{#if type === 'cbf'}
	<div class={type}>
		<Button text="Check" onClick={check} />
		<Button text="Bet" onClick={bet} />
		<Button text="Fold" onClick={fold} />
	</div>
{:else if type === 'crf'}
	<div class={type}>
		<Button text="Call" onClick={call} />
		<Button text="Raise" onClick={raise} />
		<Button text="Fold" onClick={fold} />
	</div>
{:else if type === 'nyt'}
	<div class={type}>
		<p>It is not your turn</p>
	</div>
{/if}
