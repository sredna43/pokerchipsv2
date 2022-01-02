<script lang="ts">
	import { sResponse, myName } from '../scripts/store';

	let name: string;
	myName.subscribe((n) => {
		name = n;
	});

	sResponse.subscribe((r) => {});
</script>

<p>Table</p>
<p>{name}</p>
