<script>
	/**
	 * Represents a user.
	 * @typedef {Object} User
	 * @property {string} name - The name of the user.
	 * @property {string} department - The department the user works at.
	 * @property {number} points - The points the user has.
	 * @property {number} [rank] - The rank of the user (optional at first).
	 */

	/**
	 * Generates a list of users with unique ranks based on their points.
	 * @returns {User[]} A list of 10 users with unique ranks.
	 */
	function generateUsers() {
		// List of departments
		const departments = ['Banking', 'Create', 'TechServices'];

		// List of realistic names
		const names = [
			'Frank',
			'Peter',
			'John',
			'Jane',
			'Alice',
			'Bob',
			'Eve',
			'Charlie',
			'Dave',
			'Mallory'
		];

		// Generate random users
		/** @type {User[]} */
		let users = [];
		for (let i = 0; i < 10; i++) {
			users.push({
				name: names[i],
				department: departments[Math.floor(Math.random() * departments.length)],
				points: Math.floor(Math.random() * 100) + 1 // Random points between 1 and 100
			});
		}

		// Sort users by points in descending order and assign ranks
		users
			.sort((a, b) => b.points - a.points)
			.forEach((user, index) => {
				user.rank = index + 1;
			});

		return users;
	}

	// Generate and log the list of users
	const arbitraryUsers = generateUsers();
</script>

<div class="flex flex-col h-full w-full justify-center">
	<table class="table table-hover">
		<thead>
			<tr>
				<th>Rank Number</th>
				<th>Name</th>
				<th>MMR</th>
				<th>Department</th>
			</tr>
		</thead>
		<tbody>
			{#each arbitraryUsers as user}
				<tr>
					<td>{user.rank}</td>
					<td>{user.name}</td>
					<td>{user.points}</td>
					<td>{user.department}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
