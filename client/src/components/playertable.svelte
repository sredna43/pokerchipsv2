<script lang="ts">
	import type { Player } from 'src/scripts/models';

	export let players: Player[];
	export let view = 'waitingroom';
	export let mutable = false;
</script>

{#if view === 'waitingroom' && mutable}
	<div>
		<ul>
			{#each players as { name, spot }}
				<li>
					<button></button>
					<button></button>
					<button>kick</button>
					Seat {spot + 1}: {name}
				</li>
			{/each}
		</ul>
	</div>
{:else if view === 'waitingroom'}
	<div>
		<ul>
			{#each players as { name, spot }}
				<li>Seat {spot + 1}: {name}</li>
			{/each}
		</ul>
	</div>
{:else if view === 'table'}
	<div>
		<ul>
			{#each players as { name, spot }}
				<li>Seat {spot + 1}: {name}</li>
			{/each}
		</ul>
	</div>
{/if}
