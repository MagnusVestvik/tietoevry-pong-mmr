import ApiClient from '../generated/src/ApiClient';
import GameApi from '../generated/src/api/GameApi';
import Game from '../generated/src/model/Game';
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
	clientInstance.basePath = 'http://localhost:8080';
	clientInstance.defaultHeaders = { 'Content-Type': 'application/json' };
	return clientInstance;
}

/**
 * @param {string} jwt
 */
export function getGameApiClient(jwt) {
	const client = getGameClient();
	client.defaultHeaders = {
		'Content-Type': 'application/json',
		'Authorization': `Bearer ${jwt}`
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
export async function submitGame(jwt, game) {
	if (!gameApi) {
		gameApi = getGameApiClient(jwt);
	}


	return new Promise((resolve, reject) => {
		gameApi.apiGamesPost(game, (/** @type {any} */ err, /** @type {any} */ _, /** @type {any} */ response) => {
			if (err) {
				console.error(err);
				console.log("the response was: " + JSON.stringify(response));
				reject(err);
			} else {
				console.log("the response was: " + response);
				resolve(response);
			}
		});
	});
}

