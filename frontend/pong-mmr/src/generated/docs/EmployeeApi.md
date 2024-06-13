# TietoevryPongMmr.EmployeeApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apiEmployeesGet**](EmployeeApi.md#apiEmployeesGet) | **GET** /api/employees | Get all employees
[**apiEmployeesIdDelete**](EmployeeApi.md#apiEmployeesIdDelete) | **DELETE** /api/employees/{id} | Delete an employee by ID
[**apiEmployeesIdGet**](EmployeeApi.md#apiEmployeesIdGet) | **GET** /api/employees/{id} | Get an employee by ID
[**apiEmployeesIdPut**](EmployeeApi.md#apiEmployeesIdPut) | **PUT** /api/employees/{id} | Update an existing employee
[**apiEmployeesPost**](EmployeeApi.md#apiEmployeesPost) | **POST** /api/employees | Create a new employee

<a name="apiEmployeesGet"></a>
# **apiEmployeesGet**
> [Employee] apiEmployeesGet()

Get all employees

Get all employees

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.EmployeeApi();
apiInstance.apiEmployeesGet((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[Employee]**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="apiEmployeesIdDelete"></a>
# **apiEmployeesIdDelete**
> apiEmployeesIdDelete(id)

Delete an employee by ID

Delete an employee by its ID

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.EmployeeApi();
let id = 56; // Number | Employee ID

apiInstance.apiEmployeesIdDelete(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Employee ID | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="apiEmployeesIdGet"></a>
# **apiEmployeesIdGet**
> Employee apiEmployeesIdGet(id)

Get an employee by ID

Get an employee by its ID

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.EmployeeApi();
let id = 56; // Number | Employee ID

apiInstance.apiEmployeesIdGet(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Employee ID | 

### Return type

[**Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="apiEmployeesIdPut"></a>
# **apiEmployeesIdPut**
> apiEmployeesIdPut(body, id)

Update an existing employee

Update an existing employee by its ID

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.EmployeeApi();
let body = new TietoevryPongMmr.UpdateEmployee(); // UpdateEmployee | UpdateEmployee object
let id = 56; // Number | Employee ID

apiInstance.apiEmployeesIdPut(body, id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateEmployee**](UpdateEmployee.md)| UpdateEmployee object | 
 **id** | **Number**| Employee ID | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

<a name="apiEmployeesPost"></a>
# **apiEmployeesPost**
> Employee apiEmployeesPost(body)

Create a new employee

Create a new employee

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.EmployeeApi();
let body = new TietoevryPongMmr.Employee(); // Employee | Employee object

apiInstance.apiEmployeesPost(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Employee**](Employee.md)| Employee object | 

### Return type

[**Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

