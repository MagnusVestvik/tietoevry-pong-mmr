import ApiClient from "../../generated/src/ApiClient.js";
import AuthApi from "../../generated/src/api/AuthApi.js";

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
	clientInstance.basePath = "http://localhost:8080";
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

/**
 * @typedef {Object} User
 * @property {string} email - The users email
 * @property {string} name - The users name
 * @property {string} password - The users password
 * @property {string} department - The department the user works in
 */

/**
 * @param {User} user
 */
export async function login(user) {
	return new Promise((resolve, reject) => {
		const authApiClient = getAuthApiClient();

		authApiClient.apiLoginPost(
			user,
			async (
				/** @type {Error} */ error,
				/** @type {undefined} */ _,
				/** @type {{ body: { token: any; }; }} */ response,
			) => {
				console.log("login response", response);
				if (error) {
					console.error(
						"Error logging in: ",
						error,
					);
					reject(error);
					return;
				}
				const token = response.body.token;
				try {
					const resp = await fetch(
						"/api/set-cookie",
						{
							method: "POST",
							headers: {
								"Content-Type":
									"application/json",
							},
							body: JSON.stringify({
								token: token,
							}),
						},
					);
					resolve(resp.ok);
				} catch (fetchError) {
					console.error(
						"Error setting cookie: ",
						fetchError,
					);
					reject(fetchError);
				}
			},
		);
	});
}

/**
 * Retrieves a cookie from the server.
 *
 * @returns {Promise<{Authorization?: string} | undefined>} A promise that resolves to an object with an optional Authorization property if successful, or undefined if an error occurs.
 * @throws {Error} If the server responds with a non-OK status.
 */
export async function getCookie() {
	try {
		const response = await fetch("/api/get-cookie", {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
			},
		});
		if (response.ok) {
			return response.json();
		}
		throw new Error("failed to retrive cookie");
	} catch (error) {
		console.error(error);
	}
}

/**
 * Decodes a JWT token and returns its payload as a JSON object.
 * @param {string} token - The JWT token to decode.
 * @returns {{name: string}} The decoded payload.
 */
export function parseJwt(token) {
	var base64Url = token.split(".")[1];
	var base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
	var jsonPayload = decodeURIComponent(
		window
			.atob(base64)
			.split("")
			.map(function (c) {
				return "%" +
					("00" + c.charCodeAt(0).toString(16))
						.slice(-2);
			})
			.join(""),
	);

	return JSON.parse(jsonPayload);
}

export async function deleteCookie() {
	try {
		const response = await fetch("/api/delete-cookie", {
			method: "DELETE",
		});
		if (response.ok) {
			console.log("Cookie deleted successfully");
		} else {
			console.error("Failed to delete cookie");
		}
	} catch (error) {
		console.error("Error:", error);
	}
}
