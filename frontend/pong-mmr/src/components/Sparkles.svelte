<script>
	import { onMount } from 'svelte';

	/**
	 * @type {string}
	 */
	export let height;

	/**
	 * @type {string}
	 */
	export let width;

	/**
	 * @type {string}
	 */
	export let backgroundColor;

	/**
	 * @type {string[]}
	 */
	let sparklesArr = [];

	onMount(() => {
		let div = document.createElement('div');
		div.className = 'w-1 h-1 bg-white rounded-full';
		sparklesArr = Array(100)
			.fill('')
			.map(() => div.outerHTML);
	});
</script>

<div style="width: {width}; height: {height}; background-color: {backgroundColor};">
	{#each sparklesArr as sparkle}
		{@html sparkle}
	{/each}
</div>
