# ApiV1PublicAuthOauthApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiV1PublicAuthOauthReturnGet**](ApiV1PublicAuthOauthApi.md#apiv1publicauthoauthreturnget) | **GET** /api/v1/public/auth/oauth/return |  |



## apiV1PublicAuthOauthReturnGet

> GithubComTDiblikProjectTemplateApiHandlersOAuthPostReturnHandlerResponse apiV1PublicAuthOauthReturnGet(state, code)



### Example

```ts
import {
  Configuration,
  ApiV1PublicAuthOauthApi,
} from '';
import type { ApiV1PublicAuthOauthReturnGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: x-user-token
    apiKey: "YOUR API KEY",
  });
  const api = new ApiV1PublicAuthOauthApi(config);

  const body = {
    // string
    state: state_example,
    // string
    code: code_example,
  } satisfies ApiV1PublicAuthOauthReturnGetRequest;

  try {
    const data = await api.apiV1PublicAuthOauthReturnGet(body);
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
| **state** | `string` |  | [Defaults to `undefined`] |
| **code** | `string` |  | [Defaults to `undefined`] |

### Return type

[**GithubComTDiblikProjectTemplateApiHandlersOAuthPostReturnHandlerResponse**](GithubComTDiblikProjectTemplateApiHandlersOAuthPostReturnHandlerResponse.md)

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

