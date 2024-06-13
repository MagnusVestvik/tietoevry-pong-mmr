# TietoevryPongMmr.GameApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apiGamesGet**](GameApi.md#apiGamesGet) | **GET** /api/games | Get all games
[**apiGamesIdDelete**](GameApi.md#apiGamesIdDelete) | **DELETE** /api/games/{id} | Delete an game by ID
[**apiGamesIdGet**](GameApi.md#apiGamesIdGet) | **GET** /api/games/{id} | Get an game by ID
[**apiGamesPost**](GameApi.md#apiGamesPost) | **POST** /api/games | Create a new game

<a name="apiGamesGet"></a>
# **apiGamesGet**
> [Game] apiGamesGet()

Get all games

Get all games

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.GameApi();
apiInstance.apiGamesGet((error, data, response) => {
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

[**[Game]**](Game.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="apiGamesIdDelete"></a>
# **apiGamesIdDelete**
> apiGamesIdDelete(id)

Delete an game by ID

Delete an game by its ID

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.GameApi();
let id = 56; // Number | Game ID

apiInstance.apiGamesIdDelete(id, (error, data, response) => {
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
 **id** | **Number**| Game ID | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="apiGamesIdGet"></a>
# **apiGamesIdGet**
> Game apiGamesIdGet(id)

Get an game by ID

Get an game by its ID

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.GameApi();
let id = 56; // Number | Game ID

apiInstance.apiGamesIdGet(id, (error, data, response) => {
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
 **id** | **Number**| Game ID | 

### Return type

[**Game**](Game.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="apiGamesPost"></a>
# **apiGamesPost**
> Game apiGamesPost(body)

Create a new game

Create a new game

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.GameApi();
let body = new TietoevryPongMmr.Game(); // Game | Game object

apiInstance.apiGamesPost(body, (error, data, response) => {
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
 **body** | [**Game**](Game.md)| Game object | 

### Return type

[**Game**](Game.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

