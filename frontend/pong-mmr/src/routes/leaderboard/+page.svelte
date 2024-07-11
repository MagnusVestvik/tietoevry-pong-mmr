<script>
	import { getCookie } from '$lib/auth';
	import { getEmployees } from '$lib/game';
    import { onMount } from 'svelte';

    /**
     * Represents an employee.
     * @typedef {Object} Employee
     * @property {string} name - The name of the employee.
     * @property {string} department - The department the employee works at.
     * @property {number} elo - The elo of the employee.
     * @property {string} email - The email of the employee.
     * @property {number} games - The number of games the employee has played.
     * @property {string} password - The password of the employee.
     */

    /** @type {Employee[]} */
    let employees = [];

    onMount(async () => {
		let cookie = await getCookie()
		employees = await getEmployees(cookie?.Authorization)
	});

</script>


<div class="flex flex-col h-full w-full justify-center">
	<table class="table table-hover">
		<thead>
			<tr>
				<th>Name</th>
				<th>MMR (Elo)</th>
				<th>Department</th>
			</tr>
		</thead>
		<tbody>
			{#each employees as employee}
				<tr>
					<td>{employee.Name}</td>
					<td>{employee.Elo}</td>
					<td>{employee.Department}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
