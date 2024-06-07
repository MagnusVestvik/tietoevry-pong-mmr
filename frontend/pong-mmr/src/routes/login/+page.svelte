<script>
</script>

<div class="flex flex-col justify-center items-center h-full w-full">
	<h1 class="h1 mb-36"><span class="gradient-heading">Ping-Pong MMR: Find Your Ceiling</span></h1>
	<div class="flex flex-col justify-center items-center card rounded-md">
		<label class="label m-5">
			<span>User Name</span>
			<input class="input" type="text" placeholder="  User Name" />
			<span> Password</span>
			<input class="input" type="password" placeholder="  Password" />
		</label>
		<button class="btn variant-soft m-5">Login</button>
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
