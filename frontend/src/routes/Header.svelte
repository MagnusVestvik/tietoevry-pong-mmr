<script>
	import Pong from '../components/Pong.svelte';
</script>

<header>
	<div class="flex flex-row">
		<h1 class="seymour-one-regular">TietoEvry Pong MMR</h1>
		<Pong />
	</div>
</header>

<style>
	@import url('https://fonts.googleapis.com/css2?family=Seymour+One&display=swap');

	.seymour-one-regular {
		font-family: 'Seymour One', sans-serif;
		font-weight: 400;
		font-style: normal;
	}

	h1 {
		font-family: 'Seymour One', sans-serif;
		font-size: 2rem;
		text-align: center;
	}
</style>
