<script lang="ts">
	import { LightSwitch } from '@skeletonlabs/skeleton';
	import '../app.postcss';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });
	import { onMount } from 'svelte';
	let name: string | null = null;

	onMount(async () => {
    try {
      const response = await fetch('/api/get-cookie', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      });

      if (response.ok) {
        const data = await response.json();
        const jwt = data.Authorization;
        if (jwt) {
          name = parseJwt(data.Authorization).name;
        }
      } else {
        console.error('Failed to fetch the cookie:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching the cookie:', error);
    }
  });

  function parseJwt (token) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
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
