<script lang="ts">
	import type { Player } from '../scripts/models';
	import Icon from '@iconify/svelte';
	import { sendAction } from '../scripts/ws';
	import { sortBy } from 'lodash';

	export let players: Player[];
	export let view = 'waitingroom';
	export let mutable = false;

	function arrowPressed(name: string, spot: number, direction: number) {
		sendAction({ action: 'move_player', name, amount: direction });
	}

	function makeDealer(name: string) {
		sendAction({ action: 'set_dealer', name, amount: 0 });
	}

	function kick(name: string) {
		sendAction({ action: 'remove_player', name, amount: 0 });
	}
</script>

<div class={view}>
	<p>Players:</p>
	{#if view === 'waitingroom' && mutable}
		<ul>
			{#each sortBy(players, (p) => {
				return p.spot;
			}) as { name, spot, is_dealer, is_host }}
				<li>
					<button
						class="arrow"
						on:click={() => {
							arrowPressed(name, spot, -1);
						}}><Icon icon="ic:sharp-arrow-drop-up" inline={true} /></button
					>
					<button
						class="arrow"
						on:click={() => {
							arrowPressed(name, spot, 1);
						}}><Icon icon="ic:baseline-arrow-drop-down" inline={false} /></button
					>
					Seat {spot + 1}: {name}
					{is_dealer ? 'First dealer' : ''}
					<button
						class="dealer"
						on:click={() => {
							makeDealer(name);
						}}>Make First Dealer</button
					>
					{#if !is_host}
						<button class="kick" on:click={() => kick(name)}>kick</button>
						{/if}
				</li>
			{/each}
		</ul>
	{:else if view === 'waitingroom'}
		<ul>
			{#each sortBy(players, (p) => {
				return p.spot;
			}) as { name, spot }}
				<li>Seat {spot + 1}: {name}</li>
			{/each}
		</ul>
	{:else if view === 'table'}
		<ul>
			{#each sortBy(players, (p) => {
				return p.spot;
			}) as { name, spot, chips }}
				<li>Seat {spot + 1} - Name: {name} - Chips: {chips}</li>
			{/each}
		</ul>
	{/if}
</div>
