<script>
	import { Avatar } from '@skeletonlabs/skeleton';
	import Pong from '$lib/components/Pong.svelte'; // TODO: refaktorer possisjonen til component folder
	//import submitGame from '$lib/gameApiWrapper.js'
	import Game from '../../../src/generated/src/model/Game';
	import { getCookie, parseJwt } from '$lib/auth';
	//import Employee from '../../../src/generated/src/model/Employee';

	let player1Score = 0;
	let player2Score = 0;

	async function createAndPostGame() {
		const cookie = await getCookie();
		if (cookie) {
			const data = await cookie.json();
			const jwt = data.Authorization;
			let id;
			if (jwt) {
				id = parseJwt(jwt).id;
			}
		}
		
		//create game with the employee ids and scores
		//game = new Game(player1Score, player2Score, employee1, employee2)
		//submitGame(game, jwt)
	}
</script>

<form class="flex flex-col items-center justify-center min-h-screen">
	<div class="flex flex-row items-center justify-center w-full">
		<div class="flex flex-col card p-4 m-4">
			<h1>Name1</h1>
			<header class="flex card-header justify-center">
				<Avatar
					src="https://images.unsplash.com/photo-1617296538902-887900d9b592?ixid=M3w0Njc5ODF8MHwxfGFsbHx8fHx8fHx8fDE2ODc5NzExMDB8&ixlib=rb-4.0.3&w=128&h=128&auto=format&fit=crop"
					width="w-32"
					rounded="rounded-full"
				/>
			</header>
			<section class="flex flex-col p-4 items-center">
				<h1>Score: {player1Score}</h1>
				<div class="flex flex-row">
					<button class="btn variant-filled m-1" on:click={() => player1Score++}>+</button>
					<button class="btn variant-filled m-1" on:click={() => player1Score--}>-</button>
				</div>
				<Pong />
			</section>
		</div>
		<div class="flex flex-col card p-4 m-4">
			<h1>Name2</h1>
			<header class="flex card-header justify-center">
				<Avatar
					src="https://images.unsplash.com/photo-1617296538902-887900d9b592?ixid=M3w0Njc5ODF8MHwxfGFsbHx8fHx8fHx8fDE2ODc5NzExMDB8&ixlib=rb-4.0.3&w=128&h=128&auto=format&fit=crop"
					width="w-32"
					rounded="rounded-full"
				/>
			</header>
			<section class="flex flex-col p-4 items-center">
				<h1>Score: {player2Score}</h1>
				<div class="flex flex-row">
					<button class="btn variant-filled m-1" on:click={() => player2Score++}>+</button>
					<button class="btn variant-filled m-1" on:click={() => player2Score--}>-</button>
				</div>
				<Pong />
			</section>
		</div>
	</div>
	<button on:click={createAndPostGame} class="btn variant-filled m-4 max-w-fit">Submit</button>
</form>
