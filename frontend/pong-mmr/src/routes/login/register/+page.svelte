<script>
	const inputStyle = 'input variant-form-material focus:outline-none h-10 p-4';
	const spanStyle = 'mt-6';

	let email = '';
	let userName = '';
	let password = '';
	const createUser = () => {
		'';
	};
</script>

<div class="flex flex-col justify-center items-center h-full w-full">
	<form on:submit|preventDefault={createUser}>
		<div class="card flex flex-col justify-center items-center">
			<label class="label m-5">
				<span class={spanStyle}>Email</span>
				<input class={inputStyle} type="email" bind:value={email} placeholder="Email" required />
				<span class={spanStyle}>User Name</span>
				<input
					class={inputStyle}
					type="text"
					bind:value={userName}
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
