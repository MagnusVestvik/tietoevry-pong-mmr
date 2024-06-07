<script>
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { invalidate } from '$app/navigation';

	let email = '';
	let password = '';
	let error = '';

	const handleSubmit = async () => {
		try {
			const response = await fetch('http://localhost:8080/api/login', {
				method: 'POST',
				headers: {
				'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, password })
			});

			if (!response.ok) {
				throw new Error('Login failed');
			}

			const data = await response.json();

			await fetch('/api/set-cookie', {
				method: 'POST',
				headers: {
				'Content-Type': 'application/json'
				},
				body: JSON.stringify({ token: data.token })
			});

			} catch (error) {
			console.error('Error logging in:', error);
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
