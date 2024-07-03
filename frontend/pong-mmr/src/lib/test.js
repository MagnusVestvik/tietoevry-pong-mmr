/**
 * @typedef {Object} Employee
 * @property {number} Id - The employee's ID
 * @property {string} Name - The employee's name
 * @property {string} Email - The employee's email
 * @property {string} Department - The employee's department
 * @property {string} Password - The employee's password
 */

/**
 * @typedef {Object} TestMatch
 * @property {number} Employee1Score - Score of the first employee
 * @property {number} Employee2Score - Score of the second employee
 * @property {[
 *   Employee & {[K in keyof Employee as `Employee1${K}`]: Employee[K]},
 *   Employee & {[K in keyof Employee as `Employee2${K}`]: Employee[K]}
 * ]} Employees - Array of two employees
 */

/**
 * @type {TestMatch}
 */
const testMatch = {
	Employee1Score: 0,
	Employee2Score: 0,
	Employees: [
		{
			Employee2Id: 0,
			Employee1Name: "string",
			Employee1Email: "string",
			Employee1Department: "string",
			Employee1Password: "string",
		},
		{
			Employee2Id: 0,
			Employee2Name: "string",
			Employee2Email: "string",
			Employee2Department: "string",
			Employee2Password: "string"
		}
	]
};

