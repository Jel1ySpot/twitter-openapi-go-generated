# \UserAPI

All URIs are relative to *https://x.com/i/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserByRestId**](UserAPI.md#GetUserByRestId) | **Get** /graphql/{pathQueryId}/UserByRestId | 
[**GetUserByScreenName**](UserAPI.md#GetUserByScreenName) | **Get** /graphql/{pathQueryId}/UserByScreenName | 



## GetUserByRestId

> GetUserByRestId200Response GetUserByRestId(ctx, pathQueryId).Variables(variables).Features(features).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Jel1ySpot/twitter-openapi-go-generated"
)

func main() {
	pathQueryId := "tD8zKvQzwY3kdx5yz6YmOw" // string |  (default to "tD8zKvQzwY3kdx5yz6YmOw")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserByRestId(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserByRestId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserByRestId`: GetUserByRestId200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserByRestId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;tD8zKvQzwY3kdx5yz6YmOw&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByRestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetUserByRestId200Response**](GetUserByRestId200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserByScreenName

> GetUserByRestId200Response GetUserByScreenName(ctx, pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Jel1ySpot/twitter-openapi-go-generated"
)

func main() {
	pathQueryId := "BQ6xjFU6Mgm-WhEP3OiT9w" // string |  (default to "BQ6xjFU6Mgm-WhEP3OiT9w")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 
	fieldToggles := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserByScreenName(context.Background(), pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserByScreenName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserByScreenName`: GetUserByRestId200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserByScreenName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;BQ6xjFU6Mgm-WhEP3OiT9w&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByScreenNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **fieldToggles** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetUserByRestId200Response**](GetUserByRestId200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

