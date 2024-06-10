<script>
	const inputStyle = 'input variant-form-material focus:outline-none h-10 p-4';
	const spanStyle = 'mt-6';

	let email = '';
	let name = '';
	let surName = '';
	let password = '';
	const createUser = async () => {
		try {
			const response = await fetch('', {
				// TODO: add correct url for creating user.
				method: 'Put',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, name, surName, password })
			});

			if (!response.ok) {
				throw new Error('Failed to create new user');
			}
		} catch (error) {
			console.error('Error creating user: ', error);
		}
	};
</script>

<div class="flex flex-col justify-center items-center h-full w-full">
	<form on:submit|preventDefault={createUser}>
		<div class="card flex flex-col justify-center items-center">
			<label class="label m-5">
				<span class={spanStyle}>Email</span>
				<input class={inputStyle} type="email" bind:value={email} placeholder="Email" required />
				<span class={spanStyle}>Name</span>
				<input class={inputStyle} type="text" bind:value={name} placeholder="User Name" required />
				<span class={spanStyle}>SurName</span>
				<input
					class={inputStyle}
					type="text"
					bind:value={surName}
					placeholder="User Name"
					required
				/>

				<span class={spanStyle}>Password</span>
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
