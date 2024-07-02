import { goto } from '$app/navigation';
import { login } from './auth';

/**
 * @typedef {Object} User
 * @property {string} email - The users email
 * @property {string} name - The users name
 * @property {string} password - The users password
 * @property {string} department - The department the user works in
 */

/**
 * Sends a post request to backend to create the user object and save it in the database
 * @param {User} user - the user object
 * @param{Function} callback - callback function on successfull creation of a user
 */
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


export function getCookies() {
	const cookie = document.cookie;
	return cookie;
}
