/**
 * @typedef {Object} Match
 * @property {number} Employee1Score - Score of the first employee
 * @property {number} Employee2Score - Score of the second employee
 * @property {[
 *   {[K in keyof Employee as `Employee1${Capitalize<string & K>}`]: Employee[K]},
 *   {[K in keyof Employee as `Employee2${Capitalize<string & K>}`]: Employee[K]}
 * ]} Employees - Array of two employees
 */
