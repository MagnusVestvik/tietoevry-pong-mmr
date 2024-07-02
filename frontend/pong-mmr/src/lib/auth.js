import { goto } from "$app/navigation";
import ApiClient from "../generated/src/ApiClient";
import AuthApi from "../generated/src/api/AuthApi";

/**
 * @type {ApiClient}
 */
let clientInstance;

/**
 * @type {AuthApi}
 */
let authApiClient;

function getClient() {
	if (clientInstance) {
		return clientInstance;
	}
	clientInstance = new ApiClient();
	clientInstance.basePath = 'http://localhost:8080';
	return clientInstance;
}


function getAuthApiClient() {
	const client = getClient();

	if (authApiClient) {
		return authApiClient;
	}
	authApiClient = new AuthApi(client);
	return authApiClient;
}

// TODO: fix typing
/**
 * @param {any} user
 */
export async function login(user) {
	return new Promise((resolve, reject) => {
		const authApiClient = getAuthApiClient();
		authApiClient.apiLoginPost(user, async (/** @type {any} */ error, /** @type {any} */ _, /** @type {{ body: { token: any; }; }} */ response) => {
			if (error) {
				console.error('Error logging in: ', error);
				reject(error);
				return;
			}
			const token = response.body.token;
			try {
				const resp = await fetch('/api/set-cookie', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({ token: token })
				});
				resolve(resp.ok);
			} catch (fetchError) {
				console.error('Error setting cookie: ', fetchError);
				reject(fetchError);
			}
		});
	});
}

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



