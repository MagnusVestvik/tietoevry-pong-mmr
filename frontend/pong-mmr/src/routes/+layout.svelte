<script>
	import { LightSwitch } from '@skeletonlabs/skeleton';
	import '../app.postcss';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	import { onMount } from 'svelte';
	import { getCookie, parseJwt, deleteCookie } from '$lib/services/authService.js';
	import { goto } from '$app/navigation';

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	/** @type{string | undefined} */
	let name = undefined;

	/** @type{{Authorization?:string} | undefined} */
	let cookie = undefined;

	onMount(async () => {
		cookie = await getCookie();
	});

	$: {
		if (cookie?.Authorization) {
			const payload = parseJwt(cookie.Authorization);
			name = payload.name;
		} else {
			name = undefined;
		}
	}
	async function logout() {
		deleteCookie();
		goto('/login');
	}
</script>

<!-- App Shell -->
<AppShell>
	<svelte:fragment slot="header">
		<!-- App Bar -->
		<AppBar>
			<svelte:fragment slot="lead">
				<div class="grid grid-cols-4 gap-3">
					<LightSwitch />
					<a class="btn btn-sm variant-ghost-surface" href="/leaderboard"> Leaderboard </a>
					<a class="btn btn-sm variant-ghost-surface" href="/">Home</a>
					{#if !name}
						<a class="btn btn-sm variant-ghost-surface" href="/login">Sign in</a>
					{:else}
						<button class="btn btn-sm variant-ghost-surface" on:click|preventDefault={logout}
							>Logout</button
						>
					{/if}
				</div>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<div class="flex items-center justify-center space-x-4">
					{#if name}
						<h1 class="text-center">{name}</h1>
					{/if}
					<a
						class="btn btn-icon variant-glass"
						href="https://discord.gg/rjjADEQq"
						target="_blank"
						rel="noreferrer"
					>
						<img src="/images/discordlogo.svg" alt="discord logo" class="h-8 w-8" />
					</a>
					<a
						class="btn btn-icon variant-glass"
						href="https://github.com/MagnusV9/tietoevry-pong-mmr"
						target="_blank"
						rel="noreferrer"
					>
						<img src="/images/github-mark-white.svg" alt="github logo" class="h-8 w-8" />
					</a>
				</div>
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<!-- Page Route Content -->
	<slot />
</AppShell>
