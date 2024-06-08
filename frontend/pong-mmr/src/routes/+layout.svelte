<script>
	import { LightSwitch } from '@skeletonlabs/skeleton';
	import '../app.postcss';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });
</script>

<!-- App Shell -->
<AppShell>
	<svelte:fragment slot="header">
		<!-- App Bar -->
		<AppBar>
			<svelte:fragment slot="lead">
				<div class="grid grid-cols-3 gap-3">
					<LightSwitch />
					<a class="btn btn-sm variant-ghost-surface" href="/leaderboard"> LeaderBoard </a>
					<a class="btn btn-sm variant-ghost-surface" href="/">Home</a>
				</div>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<div class="grid grid-cols-2 gap-5 logo-cloud">
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
