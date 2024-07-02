<script>
	import { goto } from '$app/navigation';
	import { login } from '$lib/auth';

	let email = '';
	let password = '';

	async function handleLogin() {
		const isLoggedIn = await login({ email, password });
		if (isLoggedIn) {
			goto('/');
		}
	}
</script>

<div class="flex flex-col justify-center items-center h-full w-full">
	<h1 class="h1 mb-36"><span class="gradient-heading">Ping-Pong MMR: Find Your Ceiling</span></h1>
	<form on:submit|preventDefault={handleLogin}>
		<div class="flex flex-col justify-center items-center card rounded-md">
			<label class="label m-5">
				<span>Email</span>
				<input
					class="input variant-form-material focus:outline-none h-10 p-4"
					type="email"
					bind:value={email}
					placeholder="Email"
					required
				/>
				<span> Password</span>
				<input
					class="input variant-form-material focus:outline-none h-10 p-4"
					type="password"
					bind:value={password}
					placeholder="Password"
					required
				/>
			</label>
			<button class="btn variant-soft m-5" type="submit">Login</button>
		</div>
	</form>
	<div>
		<p>
			Don't have an account?
			<a
				href="/login/register"
				on:click|preventDefault={() => goto('/login/register')}
				class="underline anchor"
			>
				Sign up here
			</a>
		</p>
	</div>
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
