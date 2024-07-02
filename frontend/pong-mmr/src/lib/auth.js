import { goto } from "$app/navigation";

export async function getCookie() {
	try {
		const response = await fetch('/api/get-cookie', {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		});
		if (response.ok) {
			return response.json();
		}
		throw new Error("failed to retrive cookie");
	} catch (error) {
		console.error(error)
	}
}

/**
 * Decodes a JWT token and returns its payload as a JSON object.
 * @param {string} token - The JWT token to decode.
 * @returns {Object} The decoded payload.
 */
export function parseJwt(token) {
	var base64Url = token.split('.')[1];
	var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
	var jsonPayload = decodeURIComponent(
		window
			.atob(base64)
			.split('')
			.map(function(c) {
				return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
			})
			.join('')
	);

	return JSON.parse(jsonPayload);
}

/**
 *@param {string} email
 *@param {string} password
 *@returns {Promise<boolean>} returns true if login succeed, return false otherwise
 */
export async function login(email, password) {
	try {
		const response = await fetch('http://localhost:8080/api/login', {
			// TODO: Should not be hardcoded. Should also use generated client
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ email, password })
		});

		if (!response.ok) {
			throw new Error('Login failed');
		}

		const data = await response.json();

		await fetch('/api/set-cookie', {
			// TODO: Should not be hardcoded.
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ token: data.token })
		});
	}
	catch (error) {
		console.error('Error logging in: ', error);
		return false;
	}
	return true;
}

