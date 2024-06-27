<script>
	import { popup } from '@skeletonlabs/skeleton';
	import Sparkles from '$lib/components/Sparkles.svelte';
	import { Autocomplete } from '@skeletonlabs/skeleton';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { getAllEmployees } from '$lib/getUsers';
	import { getCookie } from '$lib/auth';
	import { getCookies } from '$lib/userFunctions';

	/**
	 * Represents a user.
	 * @typedef {Object} User
	 * @property {string} id - The id used to recognize a unique user.
	 * @property {string} Name - The name of the user.
	 * @property {string} email - The users email address
	 * @property {number} games - The number of games the user have played
	 * @property {number} elo - The amount of elo points a given user have
	 */

	/**
	 *@type{Array<User>}
	 */
	let employees = [];

	/**
	 * @type {{ label: string; value: string; keywords: string; }[]}
	 */
	let userSearchOptions;

	let userSelect = '';

	let hasOpponent = false;

	function logAllCookies() {
		const cookies = getCookies();
		console.log(cookies);
	}

	onMount(async () => {
		const cookie = await getCookie();
		const auth = cookie.Authorization;
		employees = await getAllEmployees(auth)
			.then((response) => {
				return response.body;
			})
			.catch((error) => {
				console.error('Error fetching employees:', error);
			});

		employees.forEach((user, index) => {
			employees[index] = Object.fromEntries(
				Object.entries(user).filter(([key, _]) => key === 'Name')
			);
		});
		userSearchOptions = createUserSearchOptions(employees.map((user) => user.Name));
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

<button on:click={logAllCookies} class="btn variant-filled"> Cookie</button>
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
