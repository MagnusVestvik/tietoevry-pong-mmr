<script>
	import { Avatar } from '@skeletonlabs/skeleton';
	import Pong from '$lib/components/Pong.svelte';
	import { getUserId, getName } from '$lib/services/employeeService.js';
	import { submitMatch } from '$lib/services/gameService.js';
	import { getCookie } from '$lib/services/authService.js';
	import { opponent } from '$lib/store';
	import { onDestroy, onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { submitedMatch } from '$lib/store';

	/** @type {[string, string] | null} */
	let player2;
	const unsubscribe = opponent.subscribe((player) => {
		player2 = player;
	});
	onDestroy(() => {
		unsubscribe();
	});

	let matchSubmitted;

	const unsubscribeMatch = submitedMatch.subscribe((submitedMatch) => {
		matchSubmitted = submitedMatch;
	});

	onDestroy(() => {
		unsubscribeMatch();
	});

	let game;

	/** @type {string | undefined} */
	let player1Name = '';

	onMount(async () => {
		if (player2 === null) return;
		const userId = await getUserId();

		game = {
			Player1ID: userId,
			Player2ID: player2.id,
			Score1: 0,
			Score2: 0
		};

		player1Name = await getName();
	});

	async function handleGameSubmit() {
		const cookie = await getCookie();
		const jwt = cookie.Authorization;

		response = await submitMatch(jwt, game, () => {
			// set matchSubmitted to false so that a new match can be submitted later.
			matchSubmitted = false;
			goto('/');
		});
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
				<h1>Score: {game && game.Score1}</h1>
				<div class="flex flex-row">
					<button class="btn variant-filled m-1" on:click={() => game.Score1++}
						>+</button
					>
					<button class="btn variant-filled m-1" on:click={() => game.Score1--}
						>-</button
					>
				</div>
				<Pong />
			</section>
		</div>
		{#if player2}
			<div class="flex flex-col card p-4 m-4">
				<h1>{player2.name}</h1>
				<header class="flex card-header justify-center">
					<Avatar
						src="https://images.unsplash.com/photo-1617296538902-887900d9b592?ixid=M3w0Njc5ODF8MHwxfGFsbHx8fHx8fHx8fDE2ODc5NzExMDB8&ixlib=rb-4.0.3&w=128&h=128&auto=format&fit=crop"
						width="w-32"
						rounded="rounded-full"
					/>
				</header>
				<section class="flex flex-col p-4 items-center">
					<h1>Score: {game && game.Score2}</h1>
					<div class="flex flex-row">
						<button class="btn variant-filled m-1" on:click={() => game && game.Score2++}
							>+</button
						>
						<button class="btn variant-filled m-1" on:click={() => game && game.Score2--}
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
