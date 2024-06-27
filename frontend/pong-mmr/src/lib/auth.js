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

export async function deleteCookie() {
	try {
		const response = await fetch('/api/delete-cookie', {
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
