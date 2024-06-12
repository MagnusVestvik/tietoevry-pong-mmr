<script>
	const inputStyle = 'input variant-form-material focus:outline-none h-10 p-4';
	const spanStyle = 'mt-6';

	let email = '';
	let name = '';
	let lastName = '';
	let password = '';
	let department = '';
	const createUser = async () => {
		try {
			name = name + " " + lastName;
			const response = await fetch('http://localhost:8080/api/employees', {
				// TODO: add correct url for creating user.
				method: 'Post',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, name, password, department })
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
				<span class={spanStyle}>First Name</span>
				<input class={inputStyle} type="text" bind:value={name} placeholder="First Name" required />
				<span class={spanStyle}>Last Name</span>
				<input
					class={inputStyle}
					type="text"
					bind:value={lastName}
					placeholder="Last Name"
					required
				/>

				<span class={spanStyle}>Department</span>
				<input
					class={inputStyle}
					type="text"
					bind:value={department}
					placeholder="Department"
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
