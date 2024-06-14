import ApiClient from '../generated/src/ApiClient'; import EmployeeApi from '../generated/src/api/EmployeeApi';
/**
 * @type {ApiClient}
 */
let clientInstance;

/**
 * @type {EmployeeApi}
 */
let employeeApi;

function getClient() {
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
export function getEmployeeApiClient(jwt) {
	const client = getClient();

	// @ts-ignore TODO: fix
	client.defaultHeaders = { 'Authorization': `Bearer ${jwt}` };
	if (employeeApi) {
		return employeeApi;
	}
	employeeApi = new EmployeeApi(client);
	return employeeApi;
}


/**
 * @param {string} jwt
 */
export function getAllEmployees(jwt) {
	if (!employeeApi) {
		employeeApi = getEmployeeApiClient(jwt);
	}

	return new Promise((resolve, reject) => {
		employeeApi.apiEmployeesGet((/** @type {any} */ err, /** @type {any} */ data) => { // TODO: fix typing
			if (err) {
				console.error(err);
				reject(err);
			} else {
				resolve(data);
			}
		});
	});
}

