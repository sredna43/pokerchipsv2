<script lang="ts">
    import { onMount } from 'svelte';
    import { activeView } from '../scripts/appState'
    import Home from '../views/home.svelte'
    let view: string
    onMount(() => {
        activeView.subscribe((v) => {
            view = v
        })
    })
</script>

<h1>Online Poker Chips</h1>
{#if view === 'home'}
    <Home/>
{/if}