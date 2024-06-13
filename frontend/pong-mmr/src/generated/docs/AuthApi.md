# TietoevryPongMmr.AuthApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apiLoginPost**](AuthApi.md#apiLoginPost) | **POST** /api/login | Login

<a name="apiLoginPost"></a>
# **apiLoginPost**
> apiLoginPost(body)

Login

Authenticates a user and returns a JWT token

### Example
```javascript
import {TietoevryPongMmr} from 'tietoevry_pong_mmr';

let apiInstance = new TietoevryPongMmr.AuthApi();
let body = new TietoevryPongMmr.LoginRequest(); // LoginRequest | Login Request

apiInstance.apiLoginPost(body, (error, data, response) => {
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
 **body** | [**LoginRequest**](LoginRequest.md)| Login Request | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

