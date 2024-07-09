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
	return clientInstance;
}

/**
 * @param {string} jwt
 */
export function getGameApiClient() {
	const client = getGameClient();

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
export async function submitGame(game) {
	if (!gameApi) {
		gameApi = getGameApiClient();
	}

	return new Promise((resolve, reject) => {
		gameApi.apiGamesPost(game, (/** @type {any} */ err, /** @type {any} */ _, /** @type {any} */ response) => {
			if (err) {
				console.error(err);
				reject(err);
			} else {
				resolve(response);
			}
		});
	});
}



