# \V20GetAPI

All URIs are relative to *https://x.com/i/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearchAdaptive**](V20GetAPI.md#GetSearchAdaptive) | **Get** /2/search/adaptive.json | 



## GetSearchAdaptive

> GetSearchAdaptive(ctx).IncludeProfileInterstitialType(includeProfileInterstitialType).IncludeBlocking(includeBlocking).IncludeBlockedBy(includeBlockedBy).IncludeFollowedBy(includeFollowedBy).IncludeWantRetweets(includeWantRetweets).IncludeMuteEdge(includeMuteEdge).IncludeCanDm(includeCanDm).IncludeCanMediaTag(includeCanMediaTag).IncludeExtHasNftAvatar(includeExtHasNftAvatar).IncludeExtIsBlueVerified(includeExtIsBlueVerified).IncludeExtVerifiedType(includeExtVerifiedType).IncludeExtProfileImageShape(includeExtProfileImageShape).SkipStatus(skipStatus).CardsPlatform(cardsPlatform).IncludeCards(includeCards).IncludeExtAltText(includeExtAltText).IncludeExtLimitedActionResults(includeExtLimitedActionResults).IncludeQuoteCount(includeQuoteCount).IncludeReplyCount(includeReplyCount).TweetMode(tweetMode).IncludeExtViews(includeExtViews).IncludeEntities(includeEntities).IncludeUserEntities(includeUserEntities).IncludeExtMediaColor(includeExtMediaColor).IncludeExtMediaAvailability(includeExtMediaAvailability).IncludeExtSensitiveMediaWarning(includeExtSensitiveMediaWarning).IncludeExtTrustedFriendsMetadata(includeExtTrustedFriendsMetadata).SendErrorCodes(sendErrorCodes).SimpleQuotedTweet(simpleQuotedTweet).Q(q).QuerySource(querySource).Count(count).RequestContext(requestContext).Pc(pc).SpellingCorrections(spellingCorrections).IncludeExtEditControl(includeExtEditControl).Ext(ext).Execute()





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
	cardsPlatform := "Web-12" // string |  (default to "Web-12")
	includeCards := int32(1) // int32 |  (default to 1)
	includeExtAltText := true // bool |  (default to true)
	includeExtLimitedActionResults := false // bool |  (default to false)
	includeQuoteCount := true // bool |  (default to true)
	includeReplyCount := int32(1) // int32 |  (default to 1)
	tweetMode := "extended" // string |  (default to "extended")
	includeExtViews := true // bool |  (default to true)
	includeEntities := true // bool |  (default to true)
	includeUserEntities := true // bool |  (default to true)
	includeExtMediaColor := true // bool |  (default to true)
	includeExtMediaAvailability := true // bool |  (default to true)
	includeExtSensitiveMediaWarning := true // bool |  (default to true)
	includeExtTrustedFriendsMetadata := true // bool |  (default to true)
	sendErrorCodes := true // bool |  (default to true)
	simpleQuotedTweet := true // bool |  (default to true)
	q := "elon musk" // string |  (default to "elon musk")
	querySource := "trend_click" // string |  (default to "trend_click")
	count := int32(20) // int32 |  (default to 20)
	requestContext := "launch" // string |  (default to "launch")
	pc := int32(1) // int32 |  (default to 1)
	spellingCorrections := int32(1) // int32 |  (default to 1)
	includeExtEditControl := true // bool |  (default to true)
	ext := "mediaStats,highlightedLabel,hasNftAvatar,voiceInfo,birdwatchPivot,enrichments,superFollowMetadata,unmentionInfo,editControl,vibe" // string |  (default to "mediaStats,highlightedLabel,hasNftAvatar,voiceInfo,birdwatchPivot,enrichments,superFollowMetadata,unmentionInfo,editControl,vibe")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.V20GetAPI.GetSearchAdaptive(context.Background()).IncludeProfileInterstitialType(includeProfileInterstitialType).IncludeBlocking(includeBlocking).IncludeBlockedBy(includeBlockedBy).IncludeFollowedBy(includeFollowedBy).IncludeWantRetweets(includeWantRetweets).IncludeMuteEdge(includeMuteEdge).IncludeCanDm(includeCanDm).IncludeCanMediaTag(includeCanMediaTag).IncludeExtHasNftAvatar(includeExtHasNftAvatar).IncludeExtIsBlueVerified(includeExtIsBlueVerified).IncludeExtVerifiedType(includeExtVerifiedType).IncludeExtProfileImageShape(includeExtProfileImageShape).SkipStatus(skipStatus).CardsPlatform(cardsPlatform).IncludeCards(includeCards).IncludeExtAltText(includeExtAltText).IncludeExtLimitedActionResults(includeExtLimitedActionResults).IncludeQuoteCount(includeQuoteCount).IncludeReplyCount(includeReplyCount).TweetMode(tweetMode).IncludeExtViews(includeExtViews).IncludeEntities(includeEntities).IncludeUserEntities(includeUserEntities).IncludeExtMediaColor(includeExtMediaColor).IncludeExtMediaAvailability(includeExtMediaAvailability).IncludeExtSensitiveMediaWarning(includeExtSensitiveMediaWarning).IncludeExtTrustedFriendsMetadata(includeExtTrustedFriendsMetadata).SendErrorCodes(sendErrorCodes).SimpleQuotedTweet(simpleQuotedTweet).Q(q).QuerySource(querySource).Count(count).RequestContext(requestContext).Pc(pc).SpellingCorrections(spellingCorrections).IncludeExtEditControl(includeExtEditControl).Ext(ext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V20GetAPI.GetSearchAdaptive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchAdaptiveRequest struct via the builder pattern


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
 **cardsPlatform** | **string** |  | [default to &quot;Web-12&quot;]
 **includeCards** | **int32** |  | [default to 1]
 **includeExtAltText** | **bool** |  | [default to true]
 **includeExtLimitedActionResults** | **bool** |  | [default to false]
 **includeQuoteCount** | **bool** |  | [default to true]
 **includeReplyCount** | **int32** |  | [default to 1]
 **tweetMode** | **string** |  | [default to &quot;extended&quot;]
 **includeExtViews** | **bool** |  | [default to true]
 **includeEntities** | **bool** |  | [default to true]
 **includeUserEntities** | **bool** |  | [default to true]
 **includeExtMediaColor** | **bool** |  | [default to true]
 **includeExtMediaAvailability** | **bool** |  | [default to true]
 **includeExtSensitiveMediaWarning** | **bool** |  | [default to true]
 **includeExtTrustedFriendsMetadata** | **bool** |  | [default to true]
 **sendErrorCodes** | **bool** |  | [default to true]
 **simpleQuotedTweet** | **bool** |  | [default to true]
 **q** | **string** |  | [default to &quot;elon musk&quot;]
 **querySource** | **string** |  | [default to &quot;trend_click&quot;]
 **count** | **int32** |  | [default to 20]
 **requestContext** | **string** |  | [default to &quot;launch&quot;]
 **pc** | **int32** |  | [default to 1]
 **spellingCorrections** | **int32** |  | [default to 1]
 **includeExtEditControl** | **bool** |  | [default to true]
 **ext** | **string** |  | [default to &quot;mediaStats,highlightedLabel,hasNftAvatar,voiceInfo,birdwatchPivot,enrichments,superFollowMetadata,unmentionInfo,editControl,vibe&quot;]

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

