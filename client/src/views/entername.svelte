<script lang="ts">
	import Input from '../components/input.svelte';
	import Button from '../components/button.svelte';
	import { disconnect, sendAction } from '../scripts/ws';
	import { myName, sResponse } from '../scripts/store';
    import { onMount } from 'svelte';

	let name: string;
    let errorMessage = '';
    let ref;

    onMount(() => {
        ref.focus();
    })

	sResponse.subscribe((r) => {
		if (name !== '' && r.requester == name) {
			if (r.error) {
				errorMessage = r.message;
			}
		}
	});

	function handleName(): void {
		myName.set(name.replace(/\s+/g, ''));
		sendAction({
			name: name.replace(/\s+/g, ''),
			action: 'add_player',
			amount: 0
		});
	}

	function handleBack(): void {
		disconnect();
	}
</script>

<p>Enter Name</p>
<p>{errorMessage}</p>
<Input helperText="Enter Name" bind:val={name} onEnter={handleName} bind:ref/>
<Button text="Go" onClick={handleName} />
<Button text="Back" onClick={handleBack} />
