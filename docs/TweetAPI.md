# \TweetAPI

All URIs are relative to *https://x.com/i/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBookmarks**](TweetAPI.md#GetBookmarks) | **Get** /graphql/{pathQueryId}/Bookmarks | 
[**GetHomeLatestTimeline**](TweetAPI.md#GetHomeLatestTimeline) | **Get** /graphql/{pathQueryId}/HomeLatestTimeline | 
[**GetHomeTimeline**](TweetAPI.md#GetHomeTimeline) | **Get** /graphql/{pathQueryId}/HomeTimeline | 
[**GetLikes**](TweetAPI.md#GetLikes) | **Get** /graphql/{pathQueryId}/Likes | 
[**GetListLatestTweetsTimeline**](TweetAPI.md#GetListLatestTweetsTimeline) | **Get** /graphql/{pathQueryId}/ListLatestTweetsTimeline | 
[**GetSearchTimeline**](TweetAPI.md#GetSearchTimeline) | **Get** /graphql/{pathQueryId}/SearchTimeline | 
[**GetTweetDetail**](TweetAPI.md#GetTweetDetail) | **Get** /graphql/{pathQueryId}/TweetDetail | 
[**GetUserHighlightsTweets**](TweetAPI.md#GetUserHighlightsTweets) | **Get** /graphql/{pathQueryId}/UserHighlightsTweets | 
[**GetUserMedia**](TweetAPI.md#GetUserMedia) | **Get** /graphql/{pathQueryId}/UserMedia | 
[**GetUserTweets**](TweetAPI.md#GetUserTweets) | **Get** /graphql/{pathQueryId}/UserTweets | 
[**GetUserTweetsAndReplies**](TweetAPI.md#GetUserTweetsAndReplies) | **Get** /graphql/{pathQueryId}/UserTweetsAndReplies | 



## GetBookmarks

> GetBookmarks200Response GetBookmarks(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "L7vvM2UluPgWOW4GDvWyvw" // string |  (default to "L7vvM2UluPgWOW4GDvWyvw")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetBookmarks(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetBookmarks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookmarks`: GetBookmarks200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetBookmarks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;L7vvM2UluPgWOW4GDvWyvw&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookmarksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetBookmarks200Response**](GetBookmarks200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHomeLatestTimeline

> GetHomeLatestTimeline200Response GetHomeLatestTimeline(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "HyuV8ml52TYmyUjyrDq1CQ" // string |  (default to "HyuV8ml52TYmyUjyrDq1CQ")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetHomeLatestTimeline(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetHomeLatestTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHomeLatestTimeline`: GetHomeLatestTimeline200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetHomeLatestTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;HyuV8ml52TYmyUjyrDq1CQ&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetHomeLatestTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetHomeLatestTimeline200Response**](GetHomeLatestTimeline200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHomeTimeline

> GetHomeLatestTimeline200Response GetHomeTimeline(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "E6AtJXVPtK7nIHAntKc5fA" // string |  (default to "E6AtJXVPtK7nIHAntKc5fA")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetHomeTimeline(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetHomeTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHomeTimeline`: GetHomeLatestTimeline200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetHomeTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;E6AtJXVPtK7nIHAntKc5fA&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetHomeTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetHomeLatestTimeline200Response**](GetHomeLatestTimeline200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLikes

> GetLikes200Response GetLikes(ctx, pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()





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
	pathQueryId := "px6_YxfWkXo0odY84iqqmw" // string |  (default to "px6_YxfWkXo0odY84iqqmw")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 
	fieldToggles := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetLikes(context.Background(), pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetLikes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLikes`: GetLikes200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetLikes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;px6_YxfWkXo0odY84iqqmw&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetLikesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **fieldToggles** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetLikes200Response**](GetLikes200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListLatestTweetsTimeline

> GetListLatestTweetsTimeline200Response GetListLatestTweetsTimeline(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "f-Lsj0rHCztXcgdo585UUw" // string |  (default to "f-Lsj0rHCztXcgdo585UUw")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetListLatestTweetsTimeline(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetListLatestTweetsTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListLatestTweetsTimeline`: GetListLatestTweetsTimeline200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetListLatestTweetsTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;f-Lsj0rHCztXcgdo585UUw&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetListLatestTweetsTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetListLatestTweetsTimeline200Response**](GetListLatestTweetsTimeline200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchTimeline

> GetSearchTimeline200Response GetSearchTimeline(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "MJpyQGqgklrVl_0X9gNy3A" // string |  (default to "MJpyQGqgklrVl_0X9gNy3A")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetSearchTimeline(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetSearchTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchTimeline`: GetSearchTimeline200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetSearchTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;MJpyQGqgklrVl_0X9gNy3A&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetSearchTimeline200Response**](GetSearchTimeline200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTweetDetail

> GetTweetDetail200Response GetTweetDetail(ctx, pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()





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
	pathQueryId := "nBS-WpgA6ZG0CyNHD517JQ" // string |  (default to "nBS-WpgA6ZG0CyNHD517JQ")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 
	fieldToggles := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetTweetDetail(context.Background(), pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetTweetDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTweetDetail`: GetTweetDetail200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetTweetDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;nBS-WpgA6ZG0CyNHD517JQ&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTweetDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **fieldToggles** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetTweetDetail200Response**](GetTweetDetail200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserHighlightsTweets

> GetUserHighlightsTweets200Response GetUserHighlightsTweets(ctx, pathQueryId).Variables(variables).Features(features).Execute()





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
	pathQueryId := "Z-XscDcWUuMO5HalgHf57A" // string |  (default to "Z-XscDcWUuMO5HalgHf57A")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetUserHighlightsTweets(context.Background(), pathQueryId).Variables(variables).Features(features).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetUserHighlightsTweets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserHighlightsTweets`: GetUserHighlightsTweets200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetUserHighlightsTweets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;Z-XscDcWUuMO5HalgHf57A&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserHighlightsTweetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetUserHighlightsTweets200Response**](GetUserHighlightsTweets200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserMedia

> GetLikes200Response GetUserMedia(ctx, pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()





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
	pathQueryId := "HaouMjBviBKKTYZGV_9qtg" // string |  (default to "HaouMjBviBKKTYZGV_9qtg")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 
	fieldToggles := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetUserMedia(context.Background(), pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetUserMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserMedia`: GetLikes200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetUserMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;HaouMjBviBKKTYZGV_9qtg&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **fieldToggles** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetLikes200Response**](GetLikes200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTweets

> GetLikes200Response GetUserTweets(ctx, pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()





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
	pathQueryId := "Tg82Ez_kxVaJf7OPbUdbCg" // string |  (default to "Tg82Ez_kxVaJf7OPbUdbCg")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 
	fieldToggles := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetUserTweets(context.Background(), pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetUserTweets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTweets`: GetLikes200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetUserTweets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;Tg82Ez_kxVaJf7OPbUdbCg&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTweetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **fieldToggles** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetLikes200Response**](GetLikes200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTweetsAndReplies

> GetLikes200Response GetUserTweetsAndReplies(ctx, pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()





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
	pathQueryId := "HmWGzuzXoI6uFqqX6QNhEg" // string |  (default to "HmWGzuzXoI6uFqqX6QNhEg")
	variables := TODO // map[string]interface{} | 
	features := TODO // map[string]interface{} | 
	fieldToggles := TODO // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TweetAPI.GetUserTweetsAndReplies(context.Background(), pathQueryId).Variables(variables).Features(features).FieldToggles(fieldToggles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TweetAPI.GetUserTweetsAndReplies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTweetsAndReplies`: GetLikes200Response
	fmt.Fprintf(os.Stdout, "Response from `TweetAPI.GetUserTweetsAndReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;HmWGzuzXoI6uFqqX6QNhEg&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTweetsAndRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **features** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **fieldToggles** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**GetLikes200Response**](GetLikes200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

