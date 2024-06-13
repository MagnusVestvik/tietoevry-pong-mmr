import ApiClient from '../generated/src/ApiClient';
import EmployeeApi from '../generated/src/api/EmployeeApi';

// TODO: fix typing
let clientInstance;
let employeeApi;

function getClient() {
	if (clientInstance) {
		return clientInstance;
	}
	clientInstance = new ApiClient();
	clientInstance.basePath = 'http://localhost:8080';
	return clientInstance;
}

export function getEmployeeApiClient() {
	if (employeeApi) {
		return employeeApi;
	}
	employeeApi = new EmployeeApi(getClient());
	return employeeApi;
}

