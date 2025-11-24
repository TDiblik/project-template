# ApiV1PublicAuthApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiV1PublicAuthLoginPost**](ApiV1PublicAuthApi.md#apiv1publicauthloginpost) | **POST** /api/v1/public/auth/login |  |
| [**apiV1PublicAuthSignupPost**](ApiV1PublicAuthApi.md#apiv1publicauthsignuppost) | **POST** /api/v1/public/auth/signup |  |



## apiV1PublicAuthLoginPost

> GithubComTDiblikProjectTemplateApiHandlersAuthHandlerResponse apiV1PublicAuthLoginPost(githubComTDiblikProjectTemplateApiHandlersLoginHandlerRequestBody)



### Example

```ts
import {
  Configuration,
  ApiV1PublicAuthApi,
} from '';
import type { ApiV1PublicAuthLoginPostRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: x-user-token
    apiKey: "YOUR API KEY",
  });
  const api = new ApiV1PublicAuthApi(config);

  const body = {
    // GithubComTDiblikProjectTemplateApiHandlersLoginHandlerRequestBody (optional)
    githubComTDiblikProjectTemplateApiHandlersLoginHandlerRequestBody: ...,
  } satisfies ApiV1PublicAuthLoginPostRequest;

  try {
    const data = await api.apiV1PublicAuthLoginPost(body);
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
| **githubComTDiblikProjectTemplateApiHandlersLoginHandlerRequestBody** | [GithubComTDiblikProjectTemplateApiHandlersLoginHandlerRequestBody](GithubComTDiblikProjectTemplateApiHandlersLoginHandlerRequestBody.md) |  | [Optional] |

### Return type

[**GithubComTDiblikProjectTemplateApiHandlersAuthHandlerResponse**](GithubComTDiblikProjectTemplateApiHandlersAuthHandlerResponse.md)

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


## apiV1PublicAuthSignupPost

> GithubComTDiblikProjectTemplateApiHandlersAuthHandlerResponse apiV1PublicAuthSignupPost(githubComTDiblikProjectTemplateApiHandlersSignUpHandlerRequestBody)



### Example

```ts
import {
  Configuration,
  ApiV1PublicAuthApi,
} from '';
import type { ApiV1PublicAuthSignupPostRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: x-user-token
    apiKey: "YOUR API KEY",
  });
  const api = new ApiV1PublicAuthApi(config);

  const body = {
    // GithubComTDiblikProjectTemplateApiHandlersSignUpHandlerRequestBody (optional)
    githubComTDiblikProjectTemplateApiHandlersSignUpHandlerRequestBody: ...,
  } satisfies ApiV1PublicAuthSignupPostRequest;

  try {
    const data = await api.apiV1PublicAuthSignupPost(body);
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
| **githubComTDiblikProjectTemplateApiHandlersSignUpHandlerRequestBody** | [GithubComTDiblikProjectTemplateApiHandlersSignUpHandlerRequestBody](GithubComTDiblikProjectTemplateApiHandlersSignUpHandlerRequestBody.md) |  | [Optional] |

### Return type

[**GithubComTDiblikProjectTemplateApiHandlersAuthHandlerResponse**](GithubComTDiblikProjectTemplateApiHandlersAuthHandlerResponse.md)

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

