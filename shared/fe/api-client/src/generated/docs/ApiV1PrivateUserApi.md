# ApiV1PrivateUserApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiV1PrivateUserMeAvatarPost**](ApiV1PrivateUserApi.md#apiv1privateusermeavatarpost) | **POST** /api/v1/private/user/me/avatar |  |
| [**apiV1PrivateUserMeGet**](ApiV1PrivateUserApi.md#apiv1privateusermeget) | **GET** /api/v1/private/user/me |  |
| [**apiV1PrivateUserMePatch**](ApiV1PrivateUserApi.md#apiv1privateusermepatch) | **PATCH** /api/v1/private/user/me |  |



## apiV1PrivateUserMeAvatarPost

> object apiV1PrivateUserMeAvatarPost(avatar)



### Example

```ts
import {
  Configuration,
  ApiV1PrivateUserApi,
} from '';
import type { ApiV1PrivateUserMeAvatarPostRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: x-user-token
    apiKey: "YOUR API KEY",
  });
  const api = new ApiV1PrivateUserApi(config);

  const body = {
    // Blob
    avatar: BINARY_DATA_HERE,
  } satisfies ApiV1PrivateUserMeAvatarPostRequest;

  try {
    const data = await api.apiV1PrivateUserMeAvatarPost(body);
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
| **avatar** | `Blob` |  | [Defaults to `undefined`] |

### Return type

**object**

### Authorization

[x-user-token](../README.md#x-user-token)

### HTTP request headers

- **Content-Type**: `multipart/form-data`
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


## apiV1PrivateUserMeGet

> GithubComTDiblikProjectTemplateApiHandlersGetUserMeHandlerResponse apiV1PrivateUserMeGet()



### Example

```ts
import {
  Configuration,
  ApiV1PrivateUserApi,
} from '';
import type { ApiV1PrivateUserMeGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: x-user-token
    apiKey: "YOUR API KEY",
  });
  const api = new ApiV1PrivateUserApi(config);

  try {
    const data = await api.apiV1PrivateUserMeGet();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**GithubComTDiblikProjectTemplateApiHandlersGetUserMeHandlerResponse**](GithubComTDiblikProjectTemplateApiHandlersGetUserMeHandlerResponse.md)

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


## apiV1PrivateUserMePatch

> object apiV1PrivateUserMePatch(githubComTDiblikProjectTemplateApiHandlersPatchUserMeHandlerRequest)



### Example

```ts
import {
  Configuration,
  ApiV1PrivateUserApi,
} from '';
import type { ApiV1PrivateUserMePatchRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: x-user-token
    apiKey: "YOUR API KEY",
  });
  const api = new ApiV1PrivateUserApi(config);

  const body = {
    // GithubComTDiblikProjectTemplateApiHandlersPatchUserMeHandlerRequest (optional)
    githubComTDiblikProjectTemplateApiHandlersPatchUserMeHandlerRequest: ...,
  } satisfies ApiV1PrivateUserMePatchRequest;

  try {
    const data = await api.apiV1PrivateUserMePatch(body);
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
| **githubComTDiblikProjectTemplateApiHandlersPatchUserMeHandlerRequest** | [GithubComTDiblikProjectTemplateApiHandlersPatchUserMeHandlerRequest](GithubComTDiblikProjectTemplateApiHandlersPatchUserMeHandlerRequest.md) |  | [Optional] |

### Return type

**object**

### Authorization

[x-user-token](../README.md#x-user-token)

### HTTP request headers

- **Content-Type**: `application/json`, `application/x-www-form-urlencoded`, `application/xml`, `multipart/form-data`
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

