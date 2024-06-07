<script>
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { invalidate } from '$app/navigation';
	export { load } from './+page.server.js';

	let email = '';
	let password = '';
	let error = '';

	const handleSubmit = async () => {
		try {
			const response = await fetch('/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, password })
			});

			if (!response.ok) {
				throw new Error('Login failed');
			}

			await invalidate('/');

			goto('/');
		} catch (err) {
			error = err.message;
		}
	};
</script>

<div class="flex flex-col justify-center items-center h-full w-full">
	<h1 class="h1 mb-36"><span class="gradient-heading">Ping-Pong MMR: Find Your Ceiling</span></h1>
	<form on:submit|preventDefault={handleSubmit}>
		<div class="flex flex-col justify-center items-center card rounded-md">
			<label class="label m-5">
				<span>Email</span>
				<input
					class="input variant-form-material"
					type="email"
					bind:value={email}
					placeholder="  Email"
					required
				/>
				<span> Password</span>
				<input
					class="input variant-form-material"
					type="password"
					bind:value={password}
					placeholder="  Password"
					required
				/>
			</label>
			<button class="btn variant-soft m-5" type="submit">Login</button>
		</div>
	</form>
</div>

<style>
	.gradient-heading {
		@apply bg-clip-text text-transparent box-decoration-clone;
		/* Direction */
		@apply bg-gradient-to-br;
		/* Color Stops */
		@apply from-primary-500 via-tertiary-500 to-secondary-500;
	}
</style>
