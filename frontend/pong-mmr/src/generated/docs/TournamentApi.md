# TietoevryPongMmr.TournamentApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apiTournamentsGet**](TournamentApi.md#apiTournamentsGet) | **GET** /api/tournaments | Get all tournaments
[**apiTournamentsIdDelete**](TournamentApi.md#apiTournamentsIdDelete) | **DELETE** /api/tournaments/{id} | Delete a tournament
[**apiTournamentsIdGet**](TournamentApi.md#apiTournamentsIdGet) | **GET** /api/tournaments/{id} | Get a tournament by ID
[**apiTournamentsIdPut**](TournamentApi.md#apiTournamentsIdPut) | **PUT** /api/tournaments/{id} | Update a tournament
[**apiTournamentsPost**](TournamentApi.md#apiTournamentsPost) | **POST** /api/tournaments | Create a new tournament

<a name="apiTournamentsGet"></a>
# **apiTournamentsGet**
> [Tournament] apiTournamentsGet()

Get all tournaments

Get all tournaments

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.TournamentApi();
apiInstance.apiTournamentsGet((error, data, response) => {
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

[**[Tournament]**](Tournament.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="apiTournamentsIdDelete"></a>
# **apiTournamentsIdDelete**
> apiTournamentsIdDelete(id)

Delete a tournament

Delete a tournament

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.TournamentApi();
let id = 56; // Number | Tournament ID

apiInstance.apiTournamentsIdDelete(id, (error, data, response) => {
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
 **id** | **Number**| Tournament ID | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="apiTournamentsIdGet"></a>
# **apiTournamentsIdGet**
> Tournament apiTournamentsIdGet(id)

Get a tournament by ID

Get a tournament by its ID

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.TournamentApi();
let id = 56; // Number | Tournament ID

apiInstance.apiTournamentsIdGet(id, (error, data, response) => {
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
 **id** | **Number**| Tournament ID | 

### Return type

[**Tournament**](Tournament.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="apiTournamentsIdPut"></a>
# **apiTournamentsIdPut**
> Tournament apiTournamentsIdPut(body, id)

Update a tournament

Update a tournament

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.TournamentApi();
let body = new TietoevryPongMmr.Tournament(); // Tournament | Tournament object
let id = 56; // Number | Tournament ID

apiInstance.apiTournamentsIdPut(body, id, (error, data, response) => {
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
 **body** | [**Tournament**](Tournament.md)| Tournament object | 
 **id** | **Number**| Tournament ID | 

### Return type

[**Tournament**](Tournament.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

<a name="apiTournamentsPost"></a>
# **apiTournamentsPost**
> Tournament apiTournamentsPost(body)

Create a new tournament

Create a new tournament

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.TournamentApi();
let body = new TietoevryPongMmr.Tournament(); // Tournament | Tournament object

apiInstance.apiTournamentsPost(body, (error, data, response) => {
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
 **body** | [**Tournament**](Tournament.md)| Tournament object | 

### Return type

[**Tournament**](Tournament.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

