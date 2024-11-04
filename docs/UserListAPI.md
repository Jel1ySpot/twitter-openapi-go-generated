# \UserListAPI

All URIs are relative to *https://x.com/i/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFavoriters**](UserListAPI.md#GetFavoriters) | **Get** /graphql/{pathQueryId}/Favoriters | 
[**GetFollowers**](UserListAPI.md#GetFollowers) | **Get** /graphql/{pathQueryId}/Followers | 
[**GetFollowersYouKnow**](UserListAPI.md#GetFollowersYouKnow) | **Get** /graphql/{pathQueryId}/FollowersYouKnow | 
[**GetFollowing**](UserListAPI.md#GetFollowing) | **Get** /graphql/{pathQueryId}/Following | 
[**GetRetweeters**](UserListAPI.md#GetRetweeters) | **Get** /graphql/{pathQueryId}/Retweeters | 



## GetFavoriters

> GetFavoriters200Response GetFavoriters(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "riUYr3PwuHNe4tCmzjPNrg" // string |  (default to "riUYr3PwuHNe4tCmzjPNrg")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserListAPI.GetFavoriters(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserListAPI.GetFavoriters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoriters`: GetFavoriters200Response
	fmt.Fprintf(os.Stdout, "Response from `UserListAPI.GetFavoriters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;riUYr3PwuHNe4tCmzjPNrg&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetFavoriters200Response**](GetFavoriters200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowers

> GetFollowers200Response GetFollowers(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "OSXFkKmGvfw_6pGgGtkWFg" // string |  (default to "OSXFkKmGvfw_6pGgGtkWFg")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserListAPI.GetFollowers(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserListAPI.GetFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowers`: GetFollowers200Response
	fmt.Fprintf(os.Stdout, "Response from `UserListAPI.GetFollowers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;OSXFkKmGvfw_6pGgGtkWFg&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetFollowers200Response**](GetFollowers200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowersYouKnow

> GetFollowers200Response GetFollowersYouKnow(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "52sUpz5G7XvESPWgKW9i9Q" // string |  (default to "52sUpz5G7XvESPWgKW9i9Q")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserListAPI.GetFollowersYouKnow(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserListAPI.GetFollowersYouKnow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowersYouKnow`: GetFollowers200Response
	fmt.Fprintf(os.Stdout, "Response from `UserListAPI.GetFollowersYouKnow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;52sUpz5G7XvESPWgKW9i9Q&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowersYouKnowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetFollowers200Response**](GetFollowers200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowing

> GetFollowers200Response GetFollowing(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "7oQrdmth4zE3EtD42ZxgOA" // string |  (default to "7oQrdmth4zE3EtD42ZxgOA")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserListAPI.GetFollowing(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserListAPI.GetFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowing`: GetFollowers200Response
	fmt.Fprintf(os.Stdout, "Response from `UserListAPI.GetFollowing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;7oQrdmth4zE3EtD42ZxgOA&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetFollowers200Response**](GetFollowers200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetweeters

> GetRetweeters200Response GetRetweeters(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "rPcOnVhyaTBQrVgPuY7x7A" // string |  (default to "rPcOnVhyaTBQrVgPuY7x7A")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserListAPI.GetRetweeters(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserListAPI.GetRetweeters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRetweeters`: GetRetweeters200Response
	fmt.Fprintf(os.Stdout, "Response from `UserListAPI.GetRetweeters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;rPcOnVhyaTBQrVgPuY7x7A&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetweetersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetRetweeters200Response**](GetRetweeters200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

