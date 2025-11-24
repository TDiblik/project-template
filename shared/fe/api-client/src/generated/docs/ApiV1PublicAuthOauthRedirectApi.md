# ApiV1PublicAuthOauthRedirectApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiV1PublicAuthOauthRedirectFacebookGet**](ApiV1PublicAuthOauthRedirectApi.md#apiv1publicauthoauthredirectfacebookget) | **GET** /api/v1/public/auth/oauth/redirect/facebook |  |
| [**apiV1PublicAuthOauthRedirectGithubGet**](ApiV1PublicAuthOauthRedirectApi.md#apiv1publicauthoauthredirectgithubget) | **GET** /api/v1/public/auth/oauth/redirect/github |  |
| [**apiV1PublicAuthOauthRedirectGoogleGet**](ApiV1PublicAuthOauthRedirectApi.md#apiv1publicauthoauthredirectgoogleget) | **GET** /api/v1/public/auth/oauth/redirect/google |  |
| [**apiV1PublicAuthOauthRedirectSpotifyGet**](ApiV1PublicAuthOauthRedirectApi.md#apiv1publicauthoauthredirectspotifyget) | **GET** /api/v1/public/auth/oauth/redirect/spotify |  |



## apiV1PublicAuthOauthRedirectFacebookGet

> GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse apiV1PublicAuthOauthRedirectFacebookGet(redirectBackToAfterOauth)



### Example

```ts
import {
  Configuration,
  ApiV1PublicAuthOauthRedirectApi,
} from '';
import type { ApiV1PublicAuthOauthRedirectFacebookGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: x-user-token
    apiKey: "YOUR API KEY",
  });
  const api = new ApiV1PublicAuthOauthRedirectApi(config);

  const body = {
    // 'index' | 'settings' (optional)
    redirectBackToAfterOauth: redirectBackToAfterOauth_example,
  } satisfies ApiV1PublicAuthOauthRedirectFacebookGetRequest;

  try {
    const data = await api.apiV1PublicAuthOauthRedirectFacebookGet(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **redirectBackToAfterOauth** | `index`, `settings` |  | [Optional] [Defaults to `undefined`] [Enum: index, settings] |

### Return type

[**GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse**](GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse.md)

### Authorization

[x-user-token](../README.md#x-user-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | ok |  -  |
| **400** | invalid request |  -  |
| **401** | user unauthenticated |  -  |
| **403** | user unauthorized |  -  |
| **404** | not found |  -  |
| **409** | conflicting request |  -  |
| **500** | internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiV1PublicAuthOauthRedirectGithubGet

> GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse apiV1PublicAuthOauthRedirectGithubGet(redirectBackToAfterOauth)



### Example

```ts
import {
  Configuration,
  ApiV1PublicAuthOauthRedirectApi,
} from '';
import type { ApiV1PublicAuthOauthRedirectGithubGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: x-user-token
    apiKey: "YOUR API KEY",
  });
  const api = new ApiV1PublicAuthOauthRedirectApi(config);

  const body = {
    // 'index' | 'settings' (optional)
    redirectBackToAfterOauth: redirectBackToAfterOauth_example,
  } satisfies ApiV1PublicAuthOauthRedirectGithubGetRequest;

  try {
    const data = await api.apiV1PublicAuthOauthRedirectGithubGet(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **redirectBackToAfterOauth** | `index`, `settings` |  | [Optional] [Defaults to `undefined`] [Enum: index, settings] |

### Return type

[**GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse**](GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse.md)

### Authorization

[x-user-token](../README.md#x-user-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | ok |  -  |
| **400** | invalid request |  -  |
| **401** | user unauthenticated |  -  |
| **403** | user unauthorized |  -  |
| **404** | not found |  -  |
| **409** | conflicting request |  -  |
| **500** | internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiV1PublicAuthOauthRedirectGoogleGet

> GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse apiV1PublicAuthOauthRedirectGoogleGet(redirectBackToAfterOauth)



### Example

```ts
import {
  Configuration,
  ApiV1PublicAuthOauthRedirectApi,
} from '';
import type { ApiV1PublicAuthOauthRedirectGoogleGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: x-user-token
    apiKey: "YOUR API KEY",
  });
  const api = new ApiV1PublicAuthOauthRedirectApi(config);

  const body = {
    // 'index' | 'settings' (optional)
    redirectBackToAfterOauth: redirectBackToAfterOauth_example,
  } satisfies ApiV1PublicAuthOauthRedirectGoogleGetRequest;

  try {
    const data = await api.apiV1PublicAuthOauthRedirectGoogleGet(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **redirectBackToAfterOauth** | `index`, `settings` |  | [Optional] [Defaults to `undefined`] [Enum: index, settings] |

### Return type

[**GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse**](GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse.md)

### Authorization

[x-user-token](../README.md#x-user-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | ok |  -  |
| **400** | invalid request |  -  |
| **401** | user unauthenticated |  -  |
| **403** | user unauthorized |  -  |
| **404** | not found |  -  |
| **409** | conflicting request |  -  |
| **500** | internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiV1PublicAuthOauthRedirectSpotifyGet

> GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse apiV1PublicAuthOauthRedirectSpotifyGet(redirectBackToAfterOauth)



### Example

```ts
import {
  Configuration,
  ApiV1PublicAuthOauthRedirectApi,
} from '';
import type { ApiV1PublicAuthOauthRedirectSpotifyGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: x-user-token
    apiKey: "YOUR API KEY",
  });
  const api = new ApiV1PublicAuthOauthRedirectApi(config);

  const body = {
    // 'index' | 'settings' (optional)
    redirectBackToAfterOauth: redirectBackToAfterOauth_example,
  } satisfies ApiV1PublicAuthOauthRedirectSpotifyGetRequest;

  try {
    const data = await api.apiV1PublicAuthOauthRedirectSpotifyGet(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **redirectBackToAfterOauth** | `index`, `settings` |  | [Optional] [Defaults to `undefined`] [Enum: index, settings] |

### Return type

[**GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse**](GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse.md)

### Authorization

[x-user-token](../README.md#x-user-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | ok |  -  |
| **400** | invalid request |  -  |
| **401** | user unauthenticated |  -  |
| **403** | user unauthorized |  -  |
| **404** | not found |  -  |
| **409** | conflicting request |  -  |
| **500** | internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

