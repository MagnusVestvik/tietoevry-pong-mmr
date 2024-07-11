export async function createUser(user, callback) {
	try {
		const response = await fetch('http://localhost:8080/api/employees', {
			method: 'Post',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				email: user.email,
				name: user.name,
				password: user.password,
				department: user.department
			})
		});

		if (!response.ok) {
			throw new Error('Failed to create new user');
		}
		console.info('User successfully created with ' + response.ok);

		if (callback && typeof callback === 'function') {
			callback();
		}


	}
	catch (error) {
		console.error('Error creating user: ', error);
	}
};

export async function submitMatch(jwt, match, callback) {
	const response = await fetch('http://localhost:8080/api/games', { // dobbeltsjekk at det er riktig url
		method: 'Post',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': `Bearer ${jwt}`
		},
		body: JSON.stringify(match),

	});
	if (!response.ok) {
		throw new Error('Failed to submit match');
	}
	console.info('Match was submitted succesfully with: ' + response.ok);

	if (callback && typeof callback === 'function') {
		callback();
	}

}

