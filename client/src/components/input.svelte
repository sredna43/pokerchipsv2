<script>
    export let ref = null;
    export let val = '';
    export let helperText = '';
    export let className = '';
    export let onEnter = () => {
        console.log(val)
    }
    const keyUp = (e) => {
        e.keyCode === 13 && onEnter()
    }
</script>

<input bind:value={val} placeholder={helperText} class={className} on:keyup={keyUp} bind:this={ref}>