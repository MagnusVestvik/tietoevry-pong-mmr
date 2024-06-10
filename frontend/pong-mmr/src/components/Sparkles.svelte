<script>
	import { onMount } from 'svelte';

	/**
	 * Component inspired by https://ui.aceternity.com/components/sparkles
	 */

	/**
	 * @typedef {Object} Sparkles
	 * @property {number} x
	 * @property {number} y
	 * @property {number} dx
	 * @property {number} dy
	 * @property {string} animationDelay
	 */

	/**
	 * @type {string}
	 */
	export let text;
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
	 * @type {Array<Sparkles>}
	 */
	let sparklesArr = [];

	onMount(() => {
		const sparkleCount = 300;

		for (let i = 0; i < sparkleCount; i++) {
			const sparkle = {
				x: Math.random() * 100,
				y: Math.random() * 100,
				dx: (Math.random() * 10) ^ 10000,
				dy: (Math.random() * 10) ^ 10000,
				animationDelay: Math.random() + 's'
			};
			sparklesArr.push(sparkle);
		}

		const updateSparkles = () => {
			sparklesArr = sparklesArr.map((sparkle) => {
				sparkle.x += sparkle.dx;
				sparkle.y += sparkle.dy;

				if (sparkle.x < 0 || sparkle.x > 100) {
					sparkle.dx = -sparkle.dx;
				}
				if (sparkle.y < 0 || sparkle.y > 100) {
					sparkle.dy = -sparkle.dy;
				}

				if (Math.random() < 0.1) {
					sparkle.dx = Math.random() * 2 - 1;
					sparkle.dy = Math.random() * 2 - 1;
				}

				return sparkle;
			});
		};

		setInterval(updateSparkles, 60);
	});
</script>

<div
	class="relative overflow-hidden"
	style="width: {width}; height: {height}; background-color: {backgroundColor};"
>
	<div class="absolute inset-0 z-10">
		{#each sparklesArr as sparkle}
			<div
				class="w-0.5 h-0.5 bg-white rounded-full absolute animate-blink"
				style="left: {sparkle.x}%; top: {sparkle.y}%; animation-delay: {sparkle.animationDelay};"
			></div>
		{/each}
	</div>
	<h1 class="relative z-20 flex justify-center items-center h-full text-center h1">
		<span
			class="bg-gradient-to-br from-blue-500 to-cyan-300 bg-clip-text text-transparent box-decoration-clone"
			>{text}</span
		>
	</h1>
</div>

<style>
	@keyframes blink {
		0%,
		100% {
			opacity: 0;
		}
		50% {
			opacity: 1;
		}
	}
	.animate-blink {
		animation: blink 2s infinite;
	}
</style>
