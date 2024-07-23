import ApiClient from "../../generated/src/ApiClient.js";
import GameApi from "../../generated/src/api/GameApi.js";
/**
 * @type {ApiClient}
 */
let clientInstance;

/**
 * @type {GameApi}
 */
let gameApi;

function getGameClient() {
	if (clientInstance) {
		return clientInstance;
	}
	clientInstance = new ApiClient();
	clientInstance.basePath = "http://localhost:8080";
	clientInstance.defaultHeaders = { "Content-Type": "application/json" };
	return clientInstance;
}

/**
 * @param {string} jwt
 */
export function getGameApiClient(jwt) {
	const client = getGameClient();
	client.defaultHeaders = {
		"Content-Type": "application/json",
		"Authorization": `Bearer ${jwt}`,
	};

	// @ts-ignore TODO: fix
	if (gameApi) {
		return gameApi;
	}
	gameApi = new GameApi(client);
	return gameApi;
}

/**
 * @param {Game} game
 * @param {string} jwt
 */
export function submitGame(jwt, game) {
	if (!gameApi) {
		gameApi = getGameApiClient(jwt);
	}

	game = JSON.parse(game);
	console.log(game);

	return new Promise((resolve, reject) => {
		gameApi.apiGamesPost(
			game,
			(
				/** @type {any} */ err,
				/** @type {any} */ _,
				/** @type {any} */ response,
			) => {
				if (err) {
					console.error(err);
					console.log(
						"the response was: " +
						JSON.stringify(
							response,
						),
					);
					reject(err);
				} else {
					console.log(
						"the response was: " + response,
					);
					resolve(response);
				}
			},
		);
	});
}
export async function createUser(user, callback) {
	try {
		const response = await fetch(
			"http://localhost:8080/api/employees",
			{
				method: "Post",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					email: user.email,
					name: user.name,
					password: user.password,
					department: user.department,
				}),
			},
		);

		if (!response.ok) {
			throw new Error("Failed to create new user");
		}
		console.info("User successfully created with " + response.ok);

		if (callback && typeof callback === "function") {
			callback();
		}
	} catch (error) {
		console.error("Error creating user: ", error);
	}
}

export async function submitMatch(jwt, match, callback) {
	const response = await fetch("http://localhost:8080/api/games", {
		method: "Post",
		headers: {
			"Content-Type": "application/json",
			"Authorization": `Bearer ${jwt}`,
		},
		body: JSON.stringify(match),
	});
	if (!response.ok) {
		throw new Error("Failed to submit match");
	}
	console.info("Match was submitted succesfully with: " + response.ok);

	if (callback && typeof callback === "function") {
		callback();
	}
}

export async function getEmployees(jwt, callback) {
	let employees = [];
	try {
		const response = await fetch(
			"http://localhost:8080/api/employees",
			{
				headers: {
					"Authorization": `Bearer ${jwt}`,
				},
			},
		);

		if (!response.ok) {
			throw new Error("Failed to create new user");
		}

		employees = await response.json();

		if (callback && typeof callback === "function") {
			callback();
		}

		return employees;
	} catch (error) {
		console.error("Error creating user: ", error);
	}
}
