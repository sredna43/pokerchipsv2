<script lang="ts">
    export let players: string[];
</script>

<div>
    <ul>
        {#each players as p}
            <li>{p}</li>
        {/each}
    </ul>
</div>