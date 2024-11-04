# \DefaultAPI

All URIs are relative to *https://x.com/i/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProfileSpotlightsQuery**](DefaultAPI.md#GetProfileSpotlightsQuery) | **Get** /graphql/{pathQueryId}/ProfileSpotlightsQuery | 
[**GetTweetResultByRestId**](DefaultAPI.md#GetTweetResultByRestId) | **Get** /graphql/{pathQueryId}/TweetResultByRestId | 



## GetProfileSpotlightsQuery

> GetProfileSpotlightsQuery200Response GetProfileSpotlightsQuery(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "-0XdHI-mrHWBQd8-oLo1aA" // string |  (default to "-0XdHI-mrHWBQd8-oLo1aA")
	variables := TODO // map[string]interface{} | 
	features := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetProfileSpotlightsQuery(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetProfileSpotlightsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileSpotlightsQuery`: GetProfileSpotlightsQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetProfileSpotlightsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;-0XdHI-mrHWBQd8-oLo1aA&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileSpotlightsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetProfileSpotlightsQuery200Response**](GetProfileSpotlightsQuery200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTweetResultByRestId

> GetTweetResultByRestId200Response GetTweetResultByRestId(ctx, pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()





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
	pathQueryId := "7xflPyRiUxGVbJd4uWmbfg" // string |  (default to "7xflPyRiUxGVbJd4uWmbfg")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 
	fieldToggles := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTweetResultByRestId(context.Background(), pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTweetResultByRestId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTweetResultByRestId`: GetTweetResultByRestId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTweetResultByRestId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;7xflPyRiUxGVbJd4uWmbfg&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTweetResultByRestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **fieldToggles** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetTweetResultByRestId200Response**](GetTweetResultByRestId200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

