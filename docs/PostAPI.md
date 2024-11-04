# \PostAPI

All URIs are relative to *https://x.com/i/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCreateBookmark**](PostAPI.md#PostCreateBookmark) | **Post** /graphql/{pathQueryId}/CreateBookmark | 
[**PostCreateRetweet**](PostAPI.md#PostCreateRetweet) | **Post** /graphql/{pathQueryId}/CreateRetweet | 
[**PostCreateTweet**](PostAPI.md#PostCreateTweet) | **Post** /graphql/{pathQueryId}/CreateTweet | 
[**PostDeleteBookmark**](PostAPI.md#PostDeleteBookmark) | **Post** /graphql/{pathQueryId}/DeleteBookmark | 
[**PostDeleteRetweet**](PostAPI.md#PostDeleteRetweet) | **Post** /graphql/{pathQueryId}/DeleteRetweet | 
[**PostDeleteTweet**](PostAPI.md#PostDeleteTweet) | **Post** /graphql/{pathQueryId}/DeleteTweet | 
[**PostFavoriteTweet**](PostAPI.md#PostFavoriteTweet) | **Post** /graphql/{pathQueryId}/FavoriteTweet | 
[**PostUnfavoriteTweet**](PostAPI.md#PostUnfavoriteTweet) | **Post** /graphql/{pathQueryId}/UnfavoriteTweet | 



## PostCreateBookmark

> PostCreateBookmark200Response PostCreateBookmark(ctx, pathQueryId).PostCreateBookmarkRequest(postCreateBookmarkRequest).Execute()





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
	pathQueryId := "aoDbu3RHznuiSkQ9aNM67Q" // string |  (default to "aoDbu3RHznuiSkQ9aNM67Q")
	postCreateBookmarkRequest := *openapiclient.NewPostCreateBookmarkRequest("aoDbu3RHznuiSkQ9aNM67Q", *openapiclient.NewPostCreateBookmarkRequestVariables("1349129669258448897")) // PostCreateBookmarkRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostCreateBookmark(context.Background(), pathQueryId).PostCreateBookmarkRequest(postCreateBookmarkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostCreateBookmark``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCreateBookmark`: PostCreateBookmark200Response
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostCreateBookmark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;aoDbu3RHznuiSkQ9aNM67Q&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateBookmarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCreateBookmarkRequest** | [**PostCreateBookmarkRequest**](PostCreateBookmarkRequest.md) | body | 

### Return type

[**PostCreateBookmark200Response**](PostCreateBookmark200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateRetweet

> PostCreateRetweet200Response PostCreateRetweet(ctx, pathQueryId).PostCreateRetweetRequest(postCreateRetweetRequest).Execute()





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
	pathQueryId := "ojPdsZsimiJrUGLR1sjUtA" // string |  (default to "ojPdsZsimiJrUGLR1sjUtA")
	postCreateRetweetRequest := *openapiclient.NewPostCreateRetweetRequest("ojPdsZsimiJrUGLR1sjUtA", *openapiclient.NewPostCreateRetweetRequestVariables(false, "1349129669258448897")) // PostCreateRetweetRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostCreateRetweet(context.Background(), pathQueryId).PostCreateRetweetRequest(postCreateRetweetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostCreateRetweet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCreateRetweet`: PostCreateRetweet200Response
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostCreateRetweet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;ojPdsZsimiJrUGLR1sjUtA&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateRetweetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCreateRetweetRequest** | [**PostCreateRetweetRequest**](PostCreateRetweetRequest.md) | body | 

### Return type

[**PostCreateRetweet200Response**](PostCreateRetweet200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateTweet

> PostCreateTweet200Response PostCreateTweet(ctx, pathQueryId).PostCreateTweetRequest(postCreateTweetRequest).Execute()





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
	pathQueryId := "xT36w0XM3A8jDynpkram2A" // string |  (default to "xT36w0XM3A8jDynpkram2A")
	postCreateTweetRequest := *openapiclient.NewPostCreateTweetRequest(*openapiclient.NewPostCreateTweetRequestFeatures(true, true, true, false, true, true, true, true, true, true, false, true, false, true, true, true, true, true, false, true, false, true), "xT36w0XM3A8jDynpkram2A", *openapiclient.NewPostCreateTweetRequestVariables(false, *openapiclient.NewPostCreateTweetRequestVariablesMedia(false), []map[string]interface{}{map[string]interface{}(123)}, "test")) // PostCreateTweetRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostCreateTweet(context.Background(), pathQueryId).PostCreateTweetRequest(postCreateTweetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostCreateTweet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCreateTweet`: PostCreateTweet200Response
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostCreateTweet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;xT36w0XM3A8jDynpkram2A&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateTweetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCreateTweetRequest** | [**PostCreateTweetRequest**](PostCreateTweetRequest.md) | body | 

### Return type

[**PostCreateTweet200Response**](PostCreateTweet200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeleteBookmark

> PostDeleteBookmark200Response PostDeleteBookmark(ctx, pathQueryId).PostDeleteBookmarkRequest(postDeleteBookmarkRequest).Execute()





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
	pathQueryId := "Wlmlj2-xzyS1GN3a6cj-mQ" // string |  (default to "Wlmlj2-xzyS1GN3a6cj-mQ")
	postDeleteBookmarkRequest := *openapiclient.NewPostDeleteBookmarkRequest("Wlmlj2-xzyS1GN3a6cj-mQ", *openapiclient.NewPostCreateBookmarkRequestVariables("1349129669258448897")) // PostDeleteBookmarkRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostDeleteBookmark(context.Background(), pathQueryId).PostDeleteBookmarkRequest(postDeleteBookmarkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostDeleteBookmark``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeleteBookmark`: PostDeleteBookmark200Response
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostDeleteBookmark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;Wlmlj2-xzyS1GN3a6cj-mQ&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeleteBookmarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDeleteBookmarkRequest** | [**PostDeleteBookmarkRequest**](PostDeleteBookmarkRequest.md) | body | 

### Return type

[**PostDeleteBookmark200Response**](PostDeleteBookmark200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeleteRetweet

> PostDeleteRetweet200Response PostDeleteRetweet(ctx, pathQueryId).PostDeleteRetweetRequest(postDeleteRetweetRequest).Execute()





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
	pathQueryId := "iQtK4dl5hBmXewYZuEOKVw" // string |  (default to "iQtK4dl5hBmXewYZuEOKVw")
	postDeleteRetweetRequest := *openapiclient.NewPostDeleteRetweetRequest("iQtK4dl5hBmXewYZuEOKVw", *openapiclient.NewPostDeleteRetweetRequestVariables(false, "1349129669258448897")) // PostDeleteRetweetRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostDeleteRetweet(context.Background(), pathQueryId).PostDeleteRetweetRequest(postDeleteRetweetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostDeleteRetweet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeleteRetweet`: PostDeleteRetweet200Response
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostDeleteRetweet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;iQtK4dl5hBmXewYZuEOKVw&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeleteRetweetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDeleteRetweetRequest** | [**PostDeleteRetweetRequest**](PostDeleteRetweetRequest.md) | body | 

### Return type

[**PostDeleteRetweet200Response**](PostDeleteRetweet200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeleteTweet

> PostDeleteTweet200Response PostDeleteTweet(ctx, pathQueryId).PostDeleteTweetRequest(postDeleteTweetRequest).Execute()





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
	pathQueryId := "VaenaVgh5q5ih7kvyVjgtg" // string |  (default to "VaenaVgh5q5ih7kvyVjgtg")
	postDeleteTweetRequest := *openapiclient.NewPostDeleteTweetRequest("VaenaVgh5q5ih7kvyVjgtg", *openapiclient.NewPostCreateRetweetRequestVariables(false, "1349129669258448897")) // PostDeleteTweetRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostDeleteTweet(context.Background(), pathQueryId).PostDeleteTweetRequest(postDeleteTweetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostDeleteTweet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeleteTweet`: PostDeleteTweet200Response
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostDeleteTweet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;VaenaVgh5q5ih7kvyVjgtg&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeleteTweetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDeleteTweetRequest** | [**PostDeleteTweetRequest**](PostDeleteTweetRequest.md) | body | 

### Return type

[**PostDeleteTweet200Response**](PostDeleteTweet200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFavoriteTweet

> PostFavoriteTweet200Response PostFavoriteTweet(ctx, pathQueryId).PostFavoriteTweetRequest(postFavoriteTweetRequest).Execute()





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
	pathQueryId := "lI07N6Otwv1PhnEgXILM7A" // string |  (default to "lI07N6Otwv1PhnEgXILM7A")
	postFavoriteTweetRequest := *openapiclient.NewPostFavoriteTweetRequest("lI07N6Otwv1PhnEgXILM7A", *openapiclient.NewPostCreateBookmarkRequestVariables("1349129669258448897")) // PostFavoriteTweetRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostFavoriteTweet(context.Background(), pathQueryId).PostFavoriteTweetRequest(postFavoriteTweetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostFavoriteTweet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFavoriteTweet`: PostFavoriteTweet200Response
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostFavoriteTweet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;lI07N6Otwv1PhnEgXILM7A&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostFavoriteTweetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postFavoriteTweetRequest** | [**PostFavoriteTweetRequest**](PostFavoriteTweetRequest.md) | body | 

### Return type

[**PostFavoriteTweet200Response**](PostFavoriteTweet200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUnfavoriteTweet

> PostUnfavoriteTweet200Response PostUnfavoriteTweet(ctx, pathQueryId).PostUnfavoriteTweetRequest(postUnfavoriteTweetRequest).Execute()





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
	pathQueryId := "ZYKSe-w7KEslx3JhSIk5LA" // string |  (default to "ZYKSe-w7KEslx3JhSIk5LA")
	postUnfavoriteTweetRequest := *openapiclient.NewPostUnfavoriteTweetRequest("ZYKSe-w7KEslx3JhSIk5LA", *openapiclient.NewPostCreateRetweetRequestVariables(false, "1349129669258448897")) // PostUnfavoriteTweetRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostUnfavoriteTweet(context.Background(), pathQueryId).PostUnfavoriteTweetRequest(postUnfavoriteTweetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostUnfavoriteTweet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUnfavoriteTweet`: PostUnfavoriteTweet200Response
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostUnfavoriteTweet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathQueryId** | **string** |  | [default to &quot;ZYKSe-w7KEslx3JhSIk5LA&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostUnfavoriteTweetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postUnfavoriteTweetRequest** | [**PostUnfavoriteTweetRequest**](PostUnfavoriteTweetRequest.md) | body | 

### Return type

[**PostUnfavoriteTweet200Response**](PostUnfavoriteTweet200Response.md)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

