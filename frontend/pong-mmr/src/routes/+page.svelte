<script>
	import { Autocomplete } from '@skeletonlabs/skeleton';
	import { onDestroy, onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { popup } from '@skeletonlabs/skeleton';
	import Sparkles from '$lib/components/Sparkles.svelte';
	import { getAllEmployees } from '$lib/services/employeeService.js';
	import { getCookie } from '$lib/services/authService.js';
	import { opponent, submitedMatch } from '$lib/store.js';
	import Podium from '$lib/components/Podium.svelte';

	/**
	 * Represents a user.
	 * @typedef {Object} User
	 * @property {string} ID - The id used to recognize a unique user.
	 * @property {string} Name - The name of the user.
	 * @property {string} Email - The user's email address.
	 * @property {number} games - The number of games the user has played.
	 * @property {number} elo - The amount of elo points a given user has.
	 */

	/**
	 * Represents a filtered user with only the Name property.
	 * @typedef {Object} FilteredUser
	 * @property {string} Name - The name of the user.
	 * @property {string} Email - The user's email address.
	 * @property {string} ID - The id used to recognize a unique user.
	 */

	/** @type{Array<User>} */
	let employees = [];

	/** @type{Array<FilteredUser>} */
	let filteredEmployees = [];

	/** @type {{ label: string; value: string; keywords: string; }[]} */
	let userSearchOptions;

	let userSelect = '';

	let hasOpponent = false;

	let emailIdMapping = new Map();

	let submitMatch;

	const user1 = {
		name: 'Mateusz',
		mmr: 1920,
		imgUrl: 'https://i.imgur.com/EXSmx6x.png'
	};

	const user2 = {
		name: 'Magnoice',
		mmr: 2500,
		imgUrl: 'not so magnoice he forgot img url'
	};

	const user3 = {
		name: 'Elliot',
		mmr: 2140,
		imgUrl: 'https://i.imgur.com/KlXWJzE.jpeg'
	};

	const unsubscribeMatch = submitedMatch.subscribe((submitedMatch) => {
		submitMatch = submitedMatch;
	});

	onDestroy(() => {
		unsubscribeMatch();
	});

	onMount(async () => {
		const cookie = await getCookie();
		const auth = cookie?.Authorization;
		if (auth === undefined) {
			return;
		}
		employees = await getAllEmployees(auth)
			.then((response) => {
				return response.body;
			})
			.catch((error) => {
				console.error('Error fetching employees:', error);
			});

		console.log('Employees:', employees);

		filteredEmployees = employees.map((/**@type {User}*/ user) => ({
			Name: user.Name,
			Email: user.Email,
			ID: user.ID
		}));

		emailIdMapping = new Map(
			filteredEmployees.map((user) => [user.Email, { id: user.ID, name: user.Name }])
		);

		userSearchOptions = createUserSearchOptions(filteredEmployees);
	});

	/**
	 * @param filteredUsers {Array<FilteredUser>}
	 */
	function createUserSearchOptions(filteredUsers) {
		const userSearchOptions = [];
		for (const user of filteredUsers) {
			userSearchOptions.push({ label: user.Email, value: user.Name, keywords: 'plain basic' });
		}
		return userSearchOptions;
	}

	let popupSettings = {
		// TODO: add typing
		event: 'focus-click',
		target: 'popupAutocomplete',
		placement: 'bottom'
	};

	/**
	 * @param {{ detail: { label: string; }; }} event
	 */
	function onUserSelect(event) {
		userSelect = event.detail.label;
		const selectedUser = emailIdMapping.get(userSelect);
		opponent.set(selectedUser);
		hasOpponent = !hasOpponent;
	}
</script>

<div class="flex flex-col items-center mt-10 h-screen w-screen">
	<Sparkles
		text="Pong MMR: Find Your Ceiling"
		height="15%"
		width="45%"
		backgroundColor="transparent"
	/>
	<div class="grid-cols-2 m-4 min-w-[40%]">
		<input
			class="input autocomplete p-2"
			type="search"
			name="autocomplete-search"
			bind:value={userSelect}
			placeholder="Find opponent"
			use:popup={popupSettings}
		/>
		<div data-popup="popupAutocomplete">
			<Autocomplete
				bind:input={userSelect}
				options={userSearchOptions}
				on:selection={onUserSelect}
			/>
		</div>
	</div>
	
<Podium {user1} {user2} {user3} />

	<div class="flex items-center card shadow-gray-700 p-10 max-w-[80%] mt-28">
		4-10 ranks here ig
	</div>
	{#if hasOpponent}
		<button class="btn variant-filled m-4" on:click={() => goto('/match')}> Start Match</button>
	{/if}
</div>


{#if submitMatch}
	<div class="flex items-end content-end">
		<aside class="alert variant-ghost-success">
			<!-- Icon -->

			<!--TODO: legg til countdown fra timer -->
			<div>(icon)</div>
			<!-- Message -->
			<div class="alert-message">
				<h3 class="h3">(title)</h3>
				<p>Match was submitted, this improved your mmr by points</p>
				<!--TODO: legg til hvor mange points -->
			</div>
			<!-- Actions -->
			<div class="alert-actions">(buttons)</div>
		</aside>
	</div>
{/if}
