<script lang="ts">
    export let val: string = "0";
    export let max: string = "100";
</script>

<input type="range" min="0" {max} bind:value={val}>