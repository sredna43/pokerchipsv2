<script lang="ts">
	export let displayNames: string[];
</script>

<div>
	<ul>
		{#each displayNames as name}
			<li>{name}</li>
		{/each}
	</ul>
</div>
