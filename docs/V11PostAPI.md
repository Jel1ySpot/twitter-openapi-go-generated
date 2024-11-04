# \V11PostAPI

All URIs are relative to *https://x.com/i/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCreateFriendships**](V11PostAPI.md#PostCreateFriendships) | **Post** /1.1/friendships/create.json | 
[**PostDestroyFriendships**](V11PostAPI.md#PostDestroyFriendships) | **Post** /1.1/friendships/destroy.json | 



## PostCreateFriendships

> PostCreateFriendships(ctx).IncludeBlockedBy(includeBlockedBy).IncludeBlocking(includeBlocking).IncludeCanDm(includeCanDm).IncludeCanMediaTag(includeCanMediaTag).IncludeExtHasNftAvatar(includeExtHasNftAvatar).IncludeExtIsBlueVerified(includeExtIsBlueVerified).IncludeExtProfileImageShape(includeExtProfileImageShape).IncludeExtVerifiedType(includeExtVerifiedType).IncludeFollowedBy(includeFollowedBy).IncludeMuteEdge(includeMuteEdge).IncludeProfileInterstitialType(includeProfileInterstitialType).IncludeWantRetweets(includeWantRetweets).SkipStatus(skipStatus).UserId(userId).Execute()





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
	includeBlockedBy := int32(56) // int32 |  (default to 1)
	includeBlocking := int32(56) // int32 |  (default to 1)
	includeCanDm := int32(56) // int32 |  (default to 1)
	includeCanMediaTag := int32(56) // int32 |  (default to 1)
	includeExtHasNftAvatar := int32(56) // int32 |  (default to 1)
	includeExtIsBlueVerified := int32(56) // int32 |  (default to 1)
	includeExtProfileImageShape := int32(56) // int32 |  (default to 1)
	includeExtVerifiedType := int32(56) // int32 |  (default to 1)
	includeFollowedBy := int32(56) // int32 |  (default to 1)
	includeMuteEdge := int32(56) // int32 |  (default to 1)
	includeProfileInterstitialType := int32(56) // int32 |  (default to 1)
	includeWantRetweets := int32(56) // int32 |  (default to 1)
	skipStatus := int32(56) // int32 |  (default to 1)
	userId := "userId_example" // string |  (default to "44196397")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.V11PostAPI.PostCreateFriendships(context.Background()).IncludeBlockedBy(includeBlockedBy).IncludeBlocking(includeBlocking).IncludeCanDm(includeCanDm).IncludeCanMediaTag(includeCanMediaTag).IncludeExtHasNftAvatar(includeExtHasNftAvatar).IncludeExtIsBlueVerified(includeExtIsBlueVerified).IncludeExtProfileImageShape(includeExtProfileImageShape).IncludeExtVerifiedType(includeExtVerifiedType).IncludeFollowedBy(includeFollowedBy).IncludeMuteEdge(includeMuteEdge).IncludeProfileInterstitialType(includeProfileInterstitialType).IncludeWantRetweets(includeWantRetweets).SkipStatus(skipStatus).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V11PostAPI.PostCreateFriendships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFriendshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeBlockedBy** | **int32** |  | [default to 1]
 **includeBlocking** | **int32** |  | [default to 1]
 **includeCanDm** | **int32** |  | [default to 1]
 **includeCanMediaTag** | **int32** |  | [default to 1]
 **includeExtHasNftAvatar** | **int32** |  | [default to 1]
 **includeExtIsBlueVerified** | **int32** |  | [default to 1]
 **includeExtProfileImageShape** | **int32** |  | [default to 1]
 **includeExtVerifiedType** | **int32** |  | [default to 1]
 **includeFollowedBy** | **int32** |  | [default to 1]
 **includeMuteEdge** | **int32** |  | [default to 1]
 **includeProfileInterstitialType** | **int32** |  | [default to 1]
 **includeWantRetweets** | **int32** |  | [default to 1]
 **skipStatus** | **int32** |  | [default to 1]
 **userId** | **string** |  | [default to &quot;44196397&quot;]

### Return type

 (empty response body)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDestroyFriendships

> PostDestroyFriendships(ctx).IncludeBlockedBy(includeBlockedBy).IncludeBlocking(includeBlocking).IncludeCanDm(includeCanDm).IncludeCanMediaTag(includeCanMediaTag).IncludeExtHasNftAvatar(includeExtHasNftAvatar).IncludeExtIsBlueVerified(includeExtIsBlueVerified).IncludeExtProfileImageShape(includeExtProfileImageShape).IncludeExtVerifiedType(includeExtVerifiedType).IncludeFollowedBy(includeFollowedBy).IncludeMuteEdge(includeMuteEdge).IncludeProfileInterstitialType(includeProfileInterstitialType).IncludeWantRetweets(includeWantRetweets).SkipStatus(skipStatus).UserId(userId).Execute()





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
	includeBlockedBy := int32(56) // int32 |  (default to 1)
	includeBlocking := int32(56) // int32 |  (default to 1)
	includeCanDm := int32(56) // int32 |  (default to 1)
	includeCanMediaTag := int32(56) // int32 |  (default to 1)
	includeExtHasNftAvatar := int32(56) // int32 |  (default to 1)
	includeExtIsBlueVerified := int32(56) // int32 |  (default to 1)
	includeExtProfileImageShape := int32(56) // int32 |  (default to 1)
	includeExtVerifiedType := int32(56) // int32 |  (default to 1)
	includeFollowedBy := int32(56) // int32 |  (default to 1)
	includeMuteEdge := int32(56) // int32 |  (default to 1)
	includeProfileInterstitialType := int32(56) // int32 |  (default to 1)
	includeWantRetweets := int32(56) // int32 |  (default to 1)
	skipStatus := int32(56) // int32 |  (default to 1)
	userId := "userId_example" // string |  (default to "44196397")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.V11PostAPI.PostDestroyFriendships(context.Background()).IncludeBlockedBy(includeBlockedBy).IncludeBlocking(includeBlocking).IncludeCanDm(includeCanDm).IncludeCanMediaTag(includeCanMediaTag).IncludeExtHasNftAvatar(includeExtHasNftAvatar).IncludeExtIsBlueVerified(includeExtIsBlueVerified).IncludeExtProfileImageShape(includeExtProfileImageShape).IncludeExtVerifiedType(includeExtVerifiedType).IncludeFollowedBy(includeFollowedBy).IncludeMuteEdge(includeMuteEdge).IncludeProfileInterstitialType(includeProfileInterstitialType).IncludeWantRetweets(includeWantRetweets).SkipStatus(skipStatus).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V11PostAPI.PostDestroyFriendships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDestroyFriendshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeBlockedBy** | **int32** |  | [default to 1]
 **includeBlocking** | **int32** |  | [default to 1]
 **includeCanDm** | **int32** |  | [default to 1]
 **includeCanMediaTag** | **int32** |  | [default to 1]
 **includeExtHasNftAvatar** | **int32** |  | [default to 1]
 **includeExtIsBlueVerified** | **int32** |  | [default to 1]
 **includeExtProfileImageShape** | **int32** |  | [default to 1]
 **includeExtVerifiedType** | **int32** |  | [default to 1]
 **includeFollowedBy** | **int32** |  | [default to 1]
 **includeMuteEdge** | **int32** |  | [default to 1]
 **includeProfileInterstitialType** | **int32** |  | [default to 1]
 **includeWantRetweets** | **int32** |  | [default to 1]
 **skipStatus** | **int32** |  | [default to 1]
 **userId** | **string** |  | [default to &quot;44196397&quot;]

### Return type

 (empty response body)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

