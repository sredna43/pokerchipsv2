<script>
    export let disabled = false;
    export let text = "";
    export let onClick = () => {
        console.log("default on click");
    };
    export let variant = 'primary-filled';
</script>

<button class={variant} on:click={onClick} disabled={disabled} title={disabled ? 'Only host can start a game' : 'Start game'}>{text}</button>
