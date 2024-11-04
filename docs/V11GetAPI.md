# \V11GetAPI

All URIs are relative to *https://x.com/i/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFriendsFollowingList**](V11GetAPI.md#GetFriendsFollowingList) | **Get** /1.1/friends/following/list.json | 
[**GetSearchTypeahead**](V11GetAPI.md#GetSearchTypeahead) | **Get** /1.1/search/typeahead.json | 



## GetFriendsFollowingList

> GetFriendsFollowingList(ctx).IncludeProfileInterstitialType(includeProfileInterstitialType).IncludeBlocking(includeBlocking).IncludeBlockedBy(includeBlockedBy).IncludeFollowedBy(includeFollowedBy).IncludeWantRetweets(includeWantRetweets).IncludeMuteEdge(includeMuteEdge).IncludeCanDm(includeCanDm).IncludeCanMediaTag(includeCanMediaTag).IncludeExtHasNftAvatar(includeExtHasNftAvatar).IncludeExtIsBlueVerified(includeExtIsBlueVerified).IncludeExtVerifiedType(includeExtVerifiedType).IncludeExtProfileImageShape(includeExtProfileImageShape).SkipStatus(skipStatus).Cursor(cursor).UserId(userId).Count(count).WithTotalCount(withTotalCount).Execute()





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
	includeProfileInterstitialType := int32(1) // int32 |  (default to 1)
	includeBlocking := int32(1) // int32 |  (default to 1)
	includeBlockedBy := int32(1) // int32 |  (default to 1)
	includeFollowedBy := int32(1) // int32 |  (default to 1)
	includeWantRetweets := int32(1) // int32 |  (default to 1)
	includeMuteEdge := int32(1) // int32 |  (default to 1)
	includeCanDm := int32(1) // int32 |  (default to 1)
	includeCanMediaTag := int32(1) // int32 |  (default to 1)
	includeExtHasNftAvatar := int32(1) // int32 |  (default to 1)
	includeExtIsBlueVerified := int32(1) // int32 |  (default to 1)
	includeExtVerifiedType := int32(1) // int32 |  (default to 1)
	includeExtProfileImageShape := int32(1) // int32 |  (default to 1)
	skipStatus := int32(1) // int32 |  (default to 1)
	cursor := int32(-1) // int32 |  (default to -1)
	userId := "44196397" // string |  (default to "44196397")
	count := int32(3) // int32 |  (default to 3)
	withTotalCount := true // bool |  (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.V11GetAPI.GetFriendsFollowingList(context.Background()).IncludeProfileInterstitialType(includeProfileInterstitialType).IncludeBlocking(includeBlocking).IncludeBlockedBy(includeBlockedBy).IncludeFollowedBy(includeFollowedBy).IncludeWantRetweets(includeWantRetweets).IncludeMuteEdge(includeMuteEdge).IncludeCanDm(includeCanDm).IncludeCanMediaTag(includeCanMediaTag).IncludeExtHasNftAvatar(includeExtHasNftAvatar).IncludeExtIsBlueVerified(includeExtIsBlueVerified).IncludeExtVerifiedType(includeExtVerifiedType).IncludeExtProfileImageShape(includeExtProfileImageShape).SkipStatus(skipStatus).Cursor(cursor).UserId(userId).Count(count).WithTotalCount(withTotalCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V11GetAPI.GetFriendsFollowingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFriendsFollowingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeProfileInterstitialType** | **int32** |  | [default to 1]
 **includeBlocking** | **int32** |  | [default to 1]
 **includeBlockedBy** | **int32** |  | [default to 1]
 **includeFollowedBy** | **int32** |  | [default to 1]
 **includeWantRetweets** | **int32** |  | [default to 1]
 **includeMuteEdge** | **int32** |  | [default to 1]
 **includeCanDm** | **int32** |  | [default to 1]
 **includeCanMediaTag** | **int32** |  | [default to 1]
 **includeExtHasNftAvatar** | **int32** |  | [default to 1]
 **includeExtIsBlueVerified** | **int32** |  | [default to 1]
 **includeExtVerifiedType** | **int32** |  | [default to 1]
 **includeExtProfileImageShape** | **int32** |  | [default to 1]
 **skipStatus** | **int32** |  | [default to 1]
 **cursor** | **int32** |  | [default to -1]
 **userId** | **string** |  | [default to &quot;44196397&quot;]
 **count** | **int32** |  | [default to 3]
 **withTotalCount** | **bool** |  | [default to true]

### Return type

 (empty response body)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchTypeahead

> GetSearchTypeahead(ctx).IncludeExtIsBlueVerified(includeExtIsBlueVerified).IncludeExtVerifiedType(includeExtVerifiedType).IncludeExtProfileImageShape(includeExtProfileImageShape).Q(q).Src(src).ResultType(resultType).Execute()





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
	includeExtIsBlueVerified := int32(1) // int32 |  (default to 1)
	includeExtVerifiedType := int32(1) // int32 |  (default to 1)
	includeExtProfileImageShape := int32(1) // int32 |  (default to 1)
	q := "test" // string |  (default to "test")
	src := "search_box" // string |  (default to "search_box")
	resultType := "events,users,topics" // string |  (default to "events,users,topics")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.V11GetAPI.GetSearchTypeahead(context.Background()).IncludeExtIsBlueVerified(includeExtIsBlueVerified).IncludeExtVerifiedType(includeExtVerifiedType).IncludeExtProfileImageShape(includeExtProfileImageShape).Q(q).Src(src).ResultType(resultType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V11GetAPI.GetSearchTypeahead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchTypeaheadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeExtIsBlueVerified** | **int32** |  | [default to 1]
 **includeExtVerifiedType** | **int32** |  | [default to 1]
 **includeExtProfileImageShape** | **int32** |  | [default to 1]
 **q** | **string** |  | [default to &quot;test&quot;]
 **src** | **string** |  | [default to &quot;search_box&quot;]
 **resultType** | **string** |  | [default to &quot;events,users,topics&quot;]

### Return type

 (empty response body)

### Authorization

[Accept](../README.md#Accept), [ClientLanguage](../README.md#ClientLanguage), [Referer](../README.md#Referer), [SecFetchDest](../README.md#SecFetchDest), [SecChUaPlatform](../README.md#SecChUaPlatform), [SecFetchMode](../README.md#SecFetchMode), [CsrfToken](../README.md#CsrfToken), [ClientUuid](../README.md#ClientUuid), [BearerAuth](../README.md#BearerAuth), [GuestToken](../README.md#GuestToken), [SecChUa](../README.md#SecChUa), [CookieGt0](../README.md#CookieGt0), [ClientTransactionId](../README.md#ClientTransactionId), [ActiveUser](../README.md#ActiveUser), [CookieCt0](../README.md#CookieCt0), [UserAgent](../README.md#UserAgent), [AcceptLanguage](../README.md#AcceptLanguage), [SecFetchSite](../README.md#SecFetchSite), [AuthType](../README.md#AuthType), [CookieAuthToken](../README.md#CookieAuthToken), [SecChUaMobile](../README.md#SecChUaMobile), [AcceptEncoding](../README.md#AcceptEncoding)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

