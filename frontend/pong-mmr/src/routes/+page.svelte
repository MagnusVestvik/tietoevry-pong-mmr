<script>
	import { popup } from '@skeletonlabs/skeleton';
	import Sparkles from '../components/Sparkles.svelte';
	import { Autocomplete } from '@skeletonlabs/skeleton';
	import { onMount } from 'svelte';
	import { getAllEmployees } from '$lib/getUsers';
	import { getCookie } from '$lib/auth';

	//  TODO: Lag type for dette.
	/**
	 * Represents a user.
	 * @typedef {Object} UserOption
	 * @property {string} label - The name that shows up on the search.
	 * @property {string} value - the value a corresponding name gets when pressing enter on the search
	 * @property {string} keywordStyle - The styling of a keyword.
	 * @property {object} meta - A meta object that contains information about the player.
	 * @param {any} users
	 */
	function fillUserOptions(users) {
		return [].fill(users);
	}

	function generateUserOption(name, keywordStyle, meta) {
		return {
			label: name,
			value: name,
			keywordStyle: keywordStyle,
			meta: meta
		};
	}

	onMount(async () => {
		const cookie = await getCookie();
		const auth = cookie.Authorization;
		console.log(auth);

		console.log('magnuiuuus');
		userOptions = await getAllEmployees(auth);
	});

	let userOptions;
	/**
	 *@type{string}
	 */
	let inputPopupDemo = '';

	let popupSettings = {
		event: 'focus-click',
		target: 'popupAutocomplete',
		placement: 'bottom'
	};

	function onPopupDemoSelect(event) {
		inputPopupDemo = event.detail.label;
	}
</script>

<div class="flex flex-col items-center mt-10 h-screen w-screen">
	<Sparkles
		text="Pong MMR: Find Your Ceiling"
		height="10%"
		width="40%"
		backgroundColor="transparent"
	/>
	<div>
		<button on:click={() => console.log(userOptions)}> </button>
	</div>

	<div class="flex items-center card shadow-gray-700 p-10 max-w-[80%] mt-28">
		<div class="grid-cols-2">
			<input
				class="input autocomplete p-2"
				type="search"
				name="autocomplete-search"
				bind:value={inputPopupDemo}
				placeholder="Search..."
				use:popup={popupSettings}
			/>
			<div data-popup="popupAutocomplete">
				<Autocomplete
					bind:input={inputPopupDemo}
					options={userOptions}
					on:selection={onPopupDemoSelect}
				/>
			</div>
		</div>
	</div>
</div>
