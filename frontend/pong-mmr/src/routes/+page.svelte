<script>
	import { popup } from '@skeletonlabs/skeleton';
	import Sparkles from '$lib/components/Sparkles.svelte';
	import { Autocomplete } from '@skeletonlabs/skeleton';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { getAllEmployees } from '$lib/getUsers';
	import { getCookie } from '$lib/auth';

	/**
	 * Represents a user.
	 * @typedef {Object} User
	 * @property {string} id - The id used to recognize a unique user.
	 * @property {string} Name - The name of the user.
	 * @property {string} email - The user's email address.
	 * @property {number} games - The number of games the user has played.
	 * @property {number} elo - The amount of elo points a given user has.
	 */

	/**
	 * Represents a filtered user with only the Name property.
	 * @typedef {Object} FilteredUser
	 * @property {string} Name - The name of the user.
	 */

	/** @type{Array<User>} */
	let employees = [];

	/** @type{Array<FilteredUser>} */
	let filteredEmployees = [];

	/** @type {{ label: string; value: string; keywords: string; }[]} */
	let userSearchOptions;

	let userSelect = '';

	let hasOpponent = false;

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

		filteredEmployees = employees.map((/**@type {User}*/ user) => ({
			Name: user.Name
		}));

		userSearchOptions = createUserSearchOptions(filteredEmployees.map((user) => user.Name));
	});

	/**
	 * @param employeeNames {Array<string>}
	 */
	function createUserSearchOptions(employeeNames) {
		const userSearchOptions = [];
		for (const name of employeeNames) {
			userSearchOptions.push({ label: name, value: name, keywords: 'plain basic' });
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
		hasOpponent = !hasOpponent;
	}
</script>

<div class="flex flex-col items-center mt-10 h-screen w-screen">
	<Sparkles
		text="Pong MMR: Find Your Ceiling"
		height="10%"
		width="40%"
		backgroundColor="transparent"
	/>

	<div class="flex items-center card shadow-gray-700 p-10 max-w-[80%] mt-28">
		<div class="grid-cols-2">
			<input
				class="input autocomplete p-2"
				type="search"
				name="autocomplete-search"
				bind:value={userSelect}
				placeholder="Search..."
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
	</div>
	{#if hasOpponent}
		<button class="btn variant-filled m-4" on:click={() => goto('/match')}> Start Match</button>
	{/if}
</div>
