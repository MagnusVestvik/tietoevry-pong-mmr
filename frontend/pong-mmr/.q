<script>
	import { LightSwitch } from '@skeletonlabs/skeleton';
	import '../app.postcss';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	import { onMount } from 'svelte';
	import { getCookie, parseJwt } from '$lib/auth';

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });
	/**
	 * @type{string|null}
	 */
	let name = null;

	onMount(async () => {
		const cookie = await getCookie();
		if (cookie) {
  const data = await cookie.json();
  const jwt = data.Authorization;
  if (jwt) {
    name = parseJwt(jwt).name;
  }
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
					<a class="btn btn-sm variant-ghost-surface" href="/leaderboard"> LeaderBoard </a>
					<a class="btn btn-sm variant-ghost-surface" href="/">Home</a>
					{#if !name}
						<a class="btn btn-sm variant-ghost-surface" href="/login">Sign in</a>
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
