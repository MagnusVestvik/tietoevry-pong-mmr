<!-- YOU CAN DELETE EVERYTHING IN THIS PAGE -->
<script>
	import { goto } from '$app/navigation';
	import Sparkles from '../components/Sparkles.svelte';
</script>

<a href="/login" on:click|preventDefault={() => goto('/login')}>Login</a>
<div class="flex flex-col items-center h-full w-full">
	<Sparkles
		text="Pong MMR: Find Your Ceiling"
		height="10%"
		width="25%"
		backgroundColor="transparent"
	/>
</div>
