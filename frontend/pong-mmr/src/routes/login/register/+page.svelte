<script>
	import { createUser } from '$lib/userFunctions';
	import { login } from '$lib/auth';
	import { goto } from '$app/navigation';

	const inputStyle = 'input variant-form-material focus:outline-none h-10 p-4';

	let email = '';
	let firstName = '';
	let lastName = '';
	let password = '';
	let department = '';

	async function submitNewUser() {
		await createUser(
			{
				email: email,
				name: firstName + ' ' + lastName,
				password: password,
				department: department
			},
			async () => {
				const isLoggedIn = await login({ email, password });
				if (isLoggedIn) {
					goto('/');
				}
			}
		);
	}
</script>

<div class="flex flex-col justify-center items-center h-full w-full">
	<form on:submit|preventDefault={submitNewUser}>
		<div class="card flex flex-col justify-center items-center">
			<label class="label m-5">
				<span class="mt-6">Email</span>
				<input class={inputStyle} type="email" bind:value={email} placeholder="Email" required />
				<span class="mt-6">First Name</span>
				<input
					class={inputStyle}
					type="text"
					bind:value={firstName}
					placeholder="First Name"
					required
				/>
				<span class="mt-6">Last Name</span>
				<input
					class={inputStyle}
					type="text"
					bind:value={lastName}
					placeholder="Last Name"
					required
				/>

				<span class="mt-6">Department</span>
				<input
					class={inputStyle}
					type="text"
					bind:value={department}
					placeholder="Department"
					required
				/>

				<span class="mt-6">Password</span>
				<input
					class={inputStyle}
					type="password"
					bind:value={password}
					placeholder="Password"
					required
				/>
			</label>
			<button class="btn variant-soft m-5" type="submit">Create account</button>
		</div>
	</form>
</div>
