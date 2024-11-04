# \UsersAPI

All URIs are relative to *https://x.com/i/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsersByRestIds**](UsersAPI.md#GetUsersByRestIds) | **Get** /graphql/{pathQueryId}/UsersByRestIds | 



## GetUsersByRestIds

> GetUsersByRestIds200Response GetUsersByRestIds(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "lc85bOG5T3IIS4u485VtBg" // string |  (default to "lc85bOG5T3IIS4u485VtBg")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUsersByRestIds(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersByRestIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByRestIds`: GetUsersByRestIds200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersByRestIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;lc85bOG5T3IIS4u485VtBg&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByRestIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetUsersByRestIds200Response**](GetUsersByRestIds200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

