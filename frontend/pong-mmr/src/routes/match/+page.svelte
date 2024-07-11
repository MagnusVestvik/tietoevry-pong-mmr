<script>
	import { Avatar } from '@skeletonlabs/skeleton';
	import Pong from '$lib/components/Pong.svelte';
	import { getUserId, getName } from '$lib/userFunctions';
	import { opponent } from '$lib/store';
	import { onDestroy, onMount } from 'svelte';
	import { submitGame } from '$lib/gameApiWrapper';
	import { getCookie } from '$lib/auth';
	import Game from '../../generated/src/model/Game';
	import Score from '../../generated/src/model/Score';

	/** @type {[string, string] | null} */
	let player2;
	const unsubscribe = opponent.subscribe((player) => {
		player2 = player; // TODO: endre slik at player 2 er objekt og ikke array
	});
	onDestroy(() => {
		unsubscribe();
	});

	let game;

	/** @type {string | undefined} */
	let player1Name = '';

	// TODO: tror feil er at body ikke blir sendt som et json object...
	onMount(async () => {
		if (player2 === null) return;
		const userId = await getUserId();
		game = Game.constructFromObject({
			player1ID: userId,
			player2ID: player2[0],
			player2Score: Score.constructFromObject({ score: 0 }),
			player1Score: Score.constructFromObject({ score: 0 })
		});

		player1Name = await getName();
		console.log('player2', player2);
	});

	async function handleGameSubmit() {
		if (game === null) return;
		console.log('submitting game', game);
		const cookie = await getCookie();
		const auth = cookie?.Authorization;
		const response = await submitGame(auth, game);
		console.log('responded with', response);
	}
</script>

<form class="flex flex-col items-center justify-center min-h-screen">
	<div class="flex flex-row items-center justify-center w-full">
		<div class="flex flex-col card p-4 m-4">
			<h1>{player1Name}</h1>
			<header class="flex card-header justify-center">
				<Avatar
					src="https://images.unsplash.com/photo-1617296538902-887900d9b592?ixid=M3w0Njc5ODF8MHwxfGFsbHx8fHx8fHx8fDE2ODc5NzExMDB8&ixlib=rb-4.0.3&w=128&h=128&auto=format&fit=crop"
					width="w-32"
					rounded="rounded-full"
				/>
			</header>
			<section class="flex flex-col p-4 items-center">
				<h1>Score: {0}</h1>
				<div class="flex flex-row">
					<button class="btn variant-filled m-1" on:click={() => game && game.p1Score.score++}
						>+</button
					>
					<button class="btn variant-filled m-1" on:click={() => game && game.p1Score.score--}
						>-</button
					>
				</div>
				<Pong />
			</section>
		</div>
		{#if player2}
			<div class="flex flex-col card p-4 m-4">
				<h1>{player2[1]}</h1>
				<header class="flex card-header justify-center">
					<Avatar
						src="https://images.unsplash.com/photo-1617296538902-887900d9b592?ixid=M3w0Njc5ODF8MHwxfGFsbHx8fHx8fHx8fDE2ODc5NzExMDB8&ixlib=rb-4.0.3&w=128&h=128&auto=format&fit=crop"
						width="w-32"
						rounded="rounded-full"
					/>
				</header>
				<section class="flex flex-col p-4 items-center">
					<h1>Score: {0}</h1>
					<div class="flex flex-row">
						<button class="btn variant-filled m-1" on:click={() => game && game.p2Score.score++}
							>+</button
						>
						<button class="btn variant-filled m-1" on:click={() => game && game.p2Score.score--}
							>-</button
						>
					</div>
					<Pong />
				</section>
			</div>
		{/if}
	</div>
	<button on:click={handleGameSubmit} class="btn variant-filled m-4 max-w-fit">Submit</button>
</form>
