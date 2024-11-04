# Go API client for openapi

Twitter OpenAPI(Swagger) specification

## Overview
This API client was generated from [twitter-openapi](https://github.com/fa0311/twitter-openapi) by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/Jel1ySpot/twitter-openapi-go-generated"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://x.com/i/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**GetProfileSpotlightsQuery**](docs/DefaultAPI.md#getprofilespotlightsquery) | **Get** /graphql/{pathQueryId}/ProfileSpotlightsQuery | 
*DefaultAPI* | [**GetTweetResultByRestId**](docs/DefaultAPI.md#gettweetresultbyrestid) | **Get** /graphql/{pathQueryId}/TweetResultByRestId | 
*OtherAPI* | [**Other**](docs/OtherAPI.md#other) | **Get** /other | 
*PostAPI* | [**PostCreateBookmark**](docs/PostAPI.md#postcreatebookmark) | **Post** /graphql/{pathQueryId}/CreateBookmark | 
*PostAPI* | [**PostCreateRetweet**](docs/PostAPI.md#postcreateretweet) | **Post** /graphql/{pathQueryId}/CreateRetweet | 
*PostAPI* | [**PostCreateTweet**](docs/PostAPI.md#postcreatetweet) | **Post** /graphql/{pathQueryId}/CreateTweet | 
*PostAPI* | [**PostDeleteBookmark**](docs/PostAPI.md#postdeletebookmark) | **Post** /graphql/{pathQueryId}/DeleteBookmark | 
*PostAPI* | [**PostDeleteRetweet**](docs/PostAPI.md#postdeleteretweet) | **Post** /graphql/{pathQueryId}/DeleteRetweet | 
*PostAPI* | [**PostDeleteTweet**](docs/PostAPI.md#postdeletetweet) | **Post** /graphql/{pathQueryId}/DeleteTweet | 
*PostAPI* | [**PostFavoriteTweet**](docs/PostAPI.md#postfavoritetweet) | **Post** /graphql/{pathQueryId}/FavoriteTweet | 
*PostAPI* | [**PostUnfavoriteTweet**](docs/PostAPI.md#postunfavoritetweet) | **Post** /graphql/{pathQueryId}/UnfavoriteTweet | 
*TweetAPI* | [**GetBookmarks**](docs/TweetAPI.md#getbookmarks) | **Get** /graphql/{pathQueryId}/Bookmarks | 
*TweetAPI* | [**GetHomeLatestTimeline**](docs/TweetAPI.md#gethomelatesttimeline) | **Get** /graphql/{pathQueryId}/HomeLatestTimeline | 
*TweetAPI* | [**GetHomeTimeline**](docs/TweetAPI.md#gethometimeline) | **Get** /graphql/{pathQueryId}/HomeTimeline | 
*TweetAPI* | [**GetLikes**](docs/TweetAPI.md#getlikes) | **Get** /graphql/{pathQueryId}/Likes | 
*TweetAPI* | [**GetListLatestTweetsTimeline**](docs/TweetAPI.md#getlistlatesttweetstimeline) | **Get** /graphql/{pathQueryId}/ListLatestTweetsTimeline | 
*TweetAPI* | [**GetSearchTimeline**](docs/TweetAPI.md#getsearchtimeline) | **Get** /graphql/{pathQueryId}/SearchTimeline | 
*TweetAPI* | [**GetTweetDetail**](docs/TweetAPI.md#gettweetdetail) | **Get** /graphql/{pathQueryId}/TweetDetail | 
*TweetAPI* | [**GetUserHighlightsTweets**](docs/TweetAPI.md#getuserhighlightstweets) | **Get** /graphql/{pathQueryId}/UserHighlightsTweets | 
*TweetAPI* | [**GetUserMedia**](docs/TweetAPI.md#getusermedia) | **Get** /graphql/{pathQueryId}/UserMedia | 
*TweetAPI* | [**GetUserTweets**](docs/TweetAPI.md#getusertweets) | **Get** /graphql/{pathQueryId}/UserTweets | 
*TweetAPI* | [**GetUserTweetsAndReplies**](docs/TweetAPI.md#getusertweetsandreplies) | **Get** /graphql/{pathQueryId}/UserTweetsAndReplies | 
*UserAPI* | [**GetUserByRestId**](docs/UserAPI.md#getuserbyrestid) | **Get** /graphql/{pathQueryId}/UserByRestId | 
*UserAPI* | [**GetUserByScreenName**](docs/UserAPI.md#getuserbyscreenname) | **Get** /graphql/{pathQueryId}/UserByScreenName | 
*UserListAPI* | [**GetFavoriters**](docs/UserListAPI.md#getfavoriters) | **Get** /graphql/{pathQueryId}/Favoriters | 
*UserListAPI* | [**GetFollowers**](docs/UserListAPI.md#getfollowers) | **Get** /graphql/{pathQueryId}/Followers | 
*UserListAPI* | [**GetFollowersYouKnow**](docs/UserListAPI.md#getfollowersyouknow) | **Get** /graphql/{pathQueryId}/FollowersYouKnow | 
*UserListAPI* | [**GetFollowing**](docs/UserListAPI.md#getfollowing) | **Get** /graphql/{pathQueryId}/Following | 
*UserListAPI* | [**GetRetweeters**](docs/UserListAPI.md#getretweeters) | **Get** /graphql/{pathQueryId}/Retweeters | 
*UsersAPI* | [**GetUsersByRestIds**](docs/UsersAPI.md#getusersbyrestids) | **Get** /graphql/{pathQueryId}/UsersByRestIds | 
*V11GetAPI* | [**GetFriendsFollowingList**](docs/V11GetAPI.md#getfriendsfollowinglist) | **Get** /1.1/friends/following/list.json | 
*V11GetAPI* | [**GetSearchTypeahead**](docs/V11GetAPI.md#getsearchtypeahead) | **Get** /1.1/search/typeahead.json | 
*V11PostAPI* | [**PostCreateFriendships**](docs/V11PostAPI.md#postcreatefriendships) | **Post** /1.1/friendships/create.json | 
*V11PostAPI* | [**PostDestroyFriendships**](docs/V11PostAPI.md#postdestroyfriendships) | **Post** /1.1/friendships/destroy.json | 
*V20GetAPI* | [**GetSearchAdaptive**](docs/V20GetAPI.md#getsearchadaptive) | **Get** /2/search/adaptive.json | 


## Documentation For Models

 - [AdditionalMediaInfo](docs/AdditionalMediaInfo.md)
 - [AdditionalMediaInfoCallToActions](docs/AdditionalMediaInfoCallToActions.md)
 - [AdditionalMediaInfoCallToActionsUrl](docs/AdditionalMediaInfoCallToActionsUrl.md)
 - [AllowDownloadStatus](docs/AllowDownloadStatus.md)
 - [Article](docs/Article.md)
 - [ArticleCoverMedia](docs/ArticleCoverMedia.md)
 - [ArticleCoverMediaColorInfo](docs/ArticleCoverMediaColorInfo.md)
 - [ArticleCoverMediaColorInfoPalette](docs/ArticleCoverMediaColorInfoPalette.md)
 - [ArticleCoverMediaColorInfoPaletteRGB](docs/ArticleCoverMediaColorInfoPaletteRGB.md)
 - [ArticleCoverMediaInfo](docs/ArticleCoverMediaInfo.md)
 - [ArticleLifecycleState](docs/ArticleLifecycleState.md)
 - [ArticleMetadata](docs/ArticleMetadata.md)
 - [ArticleResult](docs/ArticleResult.md)
 - [ArticleResults](docs/ArticleResults.md)
 - [AuthorCommunityRelationship](docs/AuthorCommunityRelationship.md)
 - [BirdwatchEntity](docs/BirdwatchEntity.md)
 - [BirdwatchEntityRef](docs/BirdwatchEntityRef.md)
 - [BirdwatchPivot](docs/BirdwatchPivot.md)
 - [BirdwatchPivotCallToAction](docs/BirdwatchPivotCallToAction.md)
 - [BirdwatchPivotFooter](docs/BirdwatchPivotFooter.md)
 - [BirdwatchPivotNote](docs/BirdwatchPivotNote.md)
 - [BirdwatchPivotSubtitle](docs/BirdwatchPivotSubtitle.md)
 - [BookmarksResponse](docs/BookmarksResponse.md)
 - [BookmarksResponseData](docs/BookmarksResponseData.md)
 - [BookmarksTimeline](docs/BookmarksTimeline.md)
 - [Callback](docs/Callback.md)
 - [ClientEventInfo](docs/ClientEventInfo.md)
 - [CommunitiesActions](docs/CommunitiesActions.md)
 - [Community](docs/Community.md)
 - [CommunityActions](docs/CommunityActions.md)
 - [CommunityData](docs/CommunityData.md)
 - [CommunityDeleteActionResult](docs/CommunityDeleteActionResult.md)
 - [CommunityInvitesResult](docs/CommunityInvitesResult.md)
 - [CommunityJoinActionResult](docs/CommunityJoinActionResult.md)
 - [CommunityJoinRequestsResult](docs/CommunityJoinRequestsResult.md)
 - [CommunityLeaveActionResult](docs/CommunityLeaveActionResult.md)
 - [CommunityPinActionResult](docs/CommunityPinActionResult.md)
 - [CommunityRelationship](docs/CommunityRelationship.md)
 - [CommunityRule](docs/CommunityRule.md)
 - [CommunityUnpinActionResult](docs/CommunityUnpinActionResult.md)
 - [CommunityUrls](docs/CommunityUrls.md)
 - [CommunityUrlsPermalink](docs/CommunityUrlsPermalink.md)
 - [ContentEntryType](docs/ContentEntryType.md)
 - [ContentItemType](docs/ContentItemType.md)
 - [ContentUnion](docs/ContentUnion.md)
 - [ConversationControl](docs/ConversationControl.md)
 - [CoverCta](docs/CoverCta.md)
 - [CreateBookmarkResponse](docs/CreateBookmarkResponse.md)
 - [CreateBookmarkResponseData](docs/CreateBookmarkResponseData.md)
 - [CreateRetweet](docs/CreateRetweet.md)
 - [CreateRetweetResponse](docs/CreateRetweetResponse.md)
 - [CreateRetweetResponseData](docs/CreateRetweetResponseData.md)
 - [CreateRetweetResponseResult](docs/CreateRetweetResponseResult.md)
 - [CreateTweet](docs/CreateTweet.md)
 - [CreateTweetResponse](docs/CreateTweetResponse.md)
 - [CreateTweetResponseData](docs/CreateTweetResponseData.md)
 - [CreateTweetResponseResult](docs/CreateTweetResponseResult.md)
 - [CtaClientEventInfo](docs/CtaClientEventInfo.md)
 - [CursorType](docs/CursorType.md)
 - [DeleteBookmarkResponse](docs/DeleteBookmarkResponse.md)
 - [DeleteBookmarkResponseData](docs/DeleteBookmarkResponseData.md)
 - [DeleteRetweet](docs/DeleteRetweet.md)
 - [DeleteRetweetResponse](docs/DeleteRetweetResponse.md)
 - [DeleteRetweetResponseData](docs/DeleteRetweetResponseData.md)
 - [DeleteRetweetResponseResult](docs/DeleteRetweetResponseResult.md)
 - [DeleteTweetResponse](docs/DeleteTweetResponse.md)
 - [DeleteTweetResponseData](docs/DeleteTweetResponseData.md)
 - [DeleteTweetResponseResult](docs/DeleteTweetResponseResult.md)
 - [DisplayTreatment](docs/DisplayTreatment.md)
 - [DisplayType](docs/DisplayType.md)
 - [Entities](docs/Entities.md)
 - [Error](docs/Error.md)
 - [ErrorExtensions](docs/ErrorExtensions.md)
 - [Errors](docs/Errors.md)
 - [ErrorsData](docs/ErrorsData.md)
 - [ExtMediaAvailability](docs/ExtMediaAvailability.md)
 - [ExtendedEntities](docs/ExtendedEntities.md)
 - [FavoriteTweet](docs/FavoriteTweet.md)
 - [FavoriteTweetResponseData](docs/FavoriteTweetResponseData.md)
 - [FeedbackInfo](docs/FeedbackInfo.md)
 - [FollowResponse](docs/FollowResponse.md)
 - [FollowResponseData](docs/FollowResponseData.md)
 - [FollowResponseResult](docs/FollowResponseResult.md)
 - [FollowResponseUser](docs/FollowResponseUser.md)
 - [FollowTimeline](docs/FollowTimeline.md)
 - [GetBookmarks200Response](docs/GetBookmarks200Response.md)
 - [GetFavoriters200Response](docs/GetFavoriters200Response.md)
 - [GetFollowers200Response](docs/GetFollowers200Response.md)
 - [GetHomeLatestTimeline200Response](docs/GetHomeLatestTimeline200Response.md)
 - [GetLikes200Response](docs/GetLikes200Response.md)
 - [GetListLatestTweetsTimeline200Response](docs/GetListLatestTweetsTimeline200Response.md)
 - [GetProfileSpotlightsQuery200Response](docs/GetProfileSpotlightsQuery200Response.md)
 - [GetRetweeters200Response](docs/GetRetweeters200Response.md)
 - [GetSearchTimeline200Response](docs/GetSearchTimeline200Response.md)
 - [GetTweetDetail200Response](docs/GetTweetDetail200Response.md)
 - [GetTweetResultByRestId200Response](docs/GetTweetResultByRestId200Response.md)
 - [GetUserByRestId200Response](docs/GetUserByRestId200Response.md)
 - [GetUserHighlightsTweets200Response](docs/GetUserHighlightsTweets200Response.md)
 - [GetUsersByRestIds200Response](docs/GetUsersByRestIds200Response.md)
 - [Highlight](docs/Highlight.md)
 - [HomeTimelineHome](docs/HomeTimelineHome.md)
 - [HomeTimelineResponseData](docs/HomeTimelineResponseData.md)
 - [InstructionType](docs/InstructionType.md)
 - [InstructionUnion](docs/InstructionUnion.md)
 - [ItemContentUnion](docs/ItemContentUnion.md)
 - [ItemResult](docs/ItemResult.md)
 - [ListLatestTweetsTimelineResponse](docs/ListLatestTweetsTimelineResponse.md)
 - [ListTweetsTimeline](docs/ListTweetsTimeline.md)
 - [ListTweetsTimelineData](docs/ListTweetsTimelineData.md)
 - [ListTweetsTimelineList](docs/ListTweetsTimelineList.md)
 - [Location](docs/Location.md)
 - [Media](docs/Media.md)
 - [MediaExtended](docs/MediaExtended.md)
 - [MediaOriginalInfo](docs/MediaOriginalInfo.md)
 - [MediaOriginalInfoFocusRect](docs/MediaOriginalInfoFocusRect.md)
 - [MediaResult](docs/MediaResult.md)
 - [MediaResults](docs/MediaResults.md)
 - [MediaSize](docs/MediaSize.md)
 - [MediaSizes](docs/MediaSizes.md)
 - [MediaStats](docs/MediaStats.md)
 - [MediaVideoInfo](docs/MediaVideoInfo.md)
 - [MediaVideoInfoVariant](docs/MediaVideoInfoVariant.md)
 - [MediaVisibilityResults](docs/MediaVisibilityResults.md)
 - [MediaVisibilityResultsBlurredImageInterstitial](docs/MediaVisibilityResultsBlurredImageInterstitial.md)
 - [ModuleEntry](docs/ModuleEntry.md)
 - [ModuleItem](docs/ModuleItem.md)
 - [NoteTweet](docs/NoteTweet.md)
 - [NoteTweetResult](docs/NoteTweetResult.md)
 - [NoteTweetResultData](docs/NoteTweetResultData.md)
 - [NoteTweetResultMedia](docs/NoteTweetResultMedia.md)
 - [NoteTweetResultMediaInlineMedia](docs/NoteTweetResultMediaInlineMedia.md)
 - [NoteTweetResultRichText](docs/NoteTweetResultRichText.md)
 - [NoteTweetResultRichTextTag](docs/NoteTweetResultRichTextTag.md)
 - [OneFactorLoginEligibility](docs/OneFactorLoginEligibility.md)
 - [OtherResponse](docs/OtherResponse.md)
 - [PostCreateBookmark200Response](docs/PostCreateBookmark200Response.md)
 - [PostCreateBookmarkRequest](docs/PostCreateBookmarkRequest.md)
 - [PostCreateBookmarkRequestVariables](docs/PostCreateBookmarkRequestVariables.md)
 - [PostCreateRetweet200Response](docs/PostCreateRetweet200Response.md)
 - [PostCreateRetweetRequest](docs/PostCreateRetweetRequest.md)
 - [PostCreateRetweetRequestVariables](docs/PostCreateRetweetRequestVariables.md)
 - [PostCreateTweet200Response](docs/PostCreateTweet200Response.md)
 - [PostCreateTweetRequest](docs/PostCreateTweetRequest.md)
 - [PostCreateTweetRequestFeatures](docs/PostCreateTweetRequestFeatures.md)
 - [PostCreateTweetRequestVariables](docs/PostCreateTweetRequestVariables.md)
 - [PostCreateTweetRequestVariablesConversationControl](docs/PostCreateTweetRequestVariablesConversationControl.md)
 - [PostCreateTweetRequestVariablesMedia](docs/PostCreateTweetRequestVariablesMedia.md)
 - [PostCreateTweetRequestVariablesMediaMediaEntitiesInner](docs/PostCreateTweetRequestVariablesMediaMediaEntitiesInner.md)
 - [PostCreateTweetRequestVariablesReply](docs/PostCreateTweetRequestVariablesReply.md)
 - [PostDeleteBookmark200Response](docs/PostDeleteBookmark200Response.md)
 - [PostDeleteBookmarkRequest](docs/PostDeleteBookmarkRequest.md)
 - [PostDeleteRetweet200Response](docs/PostDeleteRetweet200Response.md)
 - [PostDeleteRetweetRequest](docs/PostDeleteRetweetRequest.md)
 - [PostDeleteRetweetRequestVariables](docs/PostDeleteRetweetRequestVariables.md)
 - [PostDeleteTweet200Response](docs/PostDeleteTweet200Response.md)
 - [PostDeleteTweetRequest](docs/PostDeleteTweetRequest.md)
 - [PostFavoriteTweet200Response](docs/PostFavoriteTweet200Response.md)
 - [PostFavoriteTweetRequest](docs/PostFavoriteTweetRequest.md)
 - [PostUnfavoriteTweet200Response](docs/PostUnfavoriteTweet200Response.md)
 - [PostUnfavoriteTweetRequest](docs/PostUnfavoriteTweetRequest.md)
 - [PrimaryCommunityTopic](docs/PrimaryCommunityTopic.md)
 - [ProfileResponse](docs/ProfileResponse.md)
 - [ProfileResponseData](docs/ProfileResponseData.md)
 - [QuotedRefResult](docs/QuotedRefResult.md)
 - [QuotedStatusPermalink](docs/QuotedStatusPermalink.md)
 - [Retweet](docs/Retweet.md)
 - [RetweetLegacy](docs/RetweetLegacy.md)
 - [SearchByRawQuery](docs/SearchByRawQuery.md)
 - [SearchTimeline](docs/SearchTimeline.md)
 - [SearchTimelineData](docs/SearchTimelineData.md)
 - [SearchTimelineResponse](docs/SearchTimelineResponse.md)
 - [SelfThread](docs/SelfThread.md)
 - [SensitiveMediaWarning](docs/SensitiveMediaWarning.md)
 - [Session](docs/Session.md)
 - [SocialContextLandingUrl](docs/SocialContextLandingUrl.md)
 - [SocialContextUnion](docs/SocialContextUnion.md)
 - [SocialContextUnionType](docs/SocialContextUnionType.md)
 - [SuperFollowsReplyUserResult](docs/SuperFollowsReplyUserResult.md)
 - [SuperFollowsReplyUserResultData](docs/SuperFollowsReplyUserResultData.md)
 - [SuperFollowsReplyUserResultLegacy](docs/SuperFollowsReplyUserResultLegacy.md)
 - [Text](docs/Text.md)
 - [TextEntity](docs/TextEntity.md)
 - [TextEntityRef](docs/TextEntityRef.md)
 - [TextHighlight](docs/TextHighlight.md)
 - [Timeline](docs/Timeline.md)
 - [TimelineAddEntries](docs/TimelineAddEntries.md)
 - [TimelineAddEntry](docs/TimelineAddEntry.md)
 - [TimelineAddToModule](docs/TimelineAddToModule.md)
 - [TimelineClearCache](docs/TimelineClearCache.md)
 - [TimelineCommunity](docs/TimelineCommunity.md)
 - [TimelineCoverBehavior](docs/TimelineCoverBehavior.md)
 - [TimelineCoverBehaviorUrl](docs/TimelineCoverBehaviorUrl.md)
 - [TimelineGeneralContext](docs/TimelineGeneralContext.md)
 - [TimelineHalfCover](docs/TimelineHalfCover.md)
 - [TimelineMessagePrompt](docs/TimelineMessagePrompt.md)
 - [TimelinePinEntry](docs/TimelinePinEntry.md)
 - [TimelinePrompt](docs/TimelinePrompt.md)
 - [TimelineReplaceEntry](docs/TimelineReplaceEntry.md)
 - [TimelineResponse](docs/TimelineResponse.md)
 - [TimelineShowAlert](docs/TimelineShowAlert.md)
 - [TimelineShowAlertRichText](docs/TimelineShowAlertRichText.md)
 - [TimelineShowCover](docs/TimelineShowCover.md)
 - [TimelineTerminateTimeline](docs/TimelineTerminateTimeline.md)
 - [TimelineTimelineCursor](docs/TimelineTimelineCursor.md)
 - [TimelineTimelineItem](docs/TimelineTimelineItem.md)
 - [TimelineTimelineModule](docs/TimelineTimelineModule.md)
 - [TimelineTopicContext](docs/TimelineTopicContext.md)
 - [TimelineTweet](docs/TimelineTweet.md)
 - [TimelineUser](docs/TimelineUser.md)
 - [TimelineV2](docs/TimelineV2.md)
 - [Timestamp](docs/Timestamp.md)
 - [TopicContext](docs/TopicContext.md)
 - [Tracing](docs/Tracing.md)
 - [Tweet](docs/Tweet.md)
 - [TweetCard](docs/TweetCard.md)
 - [TweetCardLegacy](docs/TweetCardLegacy.md)
 - [TweetCardLegacyBindingValue](docs/TweetCardLegacyBindingValue.md)
 - [TweetCardLegacyBindingValueData](docs/TweetCardLegacyBindingValueData.md)
 - [TweetCardLegacyBindingValueDataImage](docs/TweetCardLegacyBindingValueDataImage.md)
 - [TweetCardPlatform](docs/TweetCardPlatform.md)
 - [TweetCardPlatformAudience](docs/TweetCardPlatformAudience.md)
 - [TweetCardPlatformData](docs/TweetCardPlatformData.md)
 - [TweetCardPlatformDevice](docs/TweetCardPlatformDevice.md)
 - [TweetDetailResponse](docs/TweetDetailResponse.md)
 - [TweetDetailResponseData](docs/TweetDetailResponseData.md)
 - [TweetEditControl](docs/TweetEditControl.md)
 - [TweetEditControlInitial](docs/TweetEditControlInitial.md)
 - [TweetEditPrespective](docs/TweetEditPrespective.md)
 - [TweetFavoritersResponse](docs/TweetFavoritersResponse.md)
 - [TweetFavoritersResponseData](docs/TweetFavoritersResponseData.md)
 - [TweetInterstitial](docs/TweetInterstitial.md)
 - [TweetInterstitialRevealText](docs/TweetInterstitialRevealText.md)
 - [TweetInterstitialText](docs/TweetInterstitialText.md)
 - [TweetInterstitialTextEntity](docs/TweetInterstitialTextEntity.md)
 - [TweetInterstitialTextEntityRef](docs/TweetInterstitialTextEntityRef.md)
 - [TweetLegacy](docs/TweetLegacy.md)
 - [TweetLegacyScopes](docs/TweetLegacyScopes.md)
 - [TweetPreviousCounts](docs/TweetPreviousCounts.md)
 - [TweetResultByRestIdData](docs/TweetResultByRestIdData.md)
 - [TweetResultByRestIdResponse](docs/TweetResultByRestIdResponse.md)
 - [TweetRetweetersResponse](docs/TweetRetweetersResponse.md)
 - [TweetRetweetersResponseData](docs/TweetRetweetersResponseData.md)
 - [TweetTombstone](docs/TweetTombstone.md)
 - [TweetUnavailable](docs/TweetUnavailable.md)
 - [TweetUnion](docs/TweetUnion.md)
 - [TweetView](docs/TweetView.md)
 - [TweetWithVisibilityResults](docs/TweetWithVisibilityResults.md)
 - [TypeName](docs/TypeName.md)
 - [UnfavoriteTweet](docs/UnfavoriteTweet.md)
 - [UnfavoriteTweetResponseData](docs/UnfavoriteTweetResponseData.md)
 - [UnifiedCard](docs/UnifiedCard.md)
 - [Url](docs/Url.md)
 - [UrtEndpointOptions](docs/UrtEndpointOptions.md)
 - [UrtEndpointRequestParams](docs/UrtEndpointRequestParams.md)
 - [User](docs/User.md)
 - [UserFeatures](docs/UserFeatures.md)
 - [UserHighlightsInfo](docs/UserHighlightsInfo.md)
 - [UserHighlightsTweetsData](docs/UserHighlightsTweetsData.md)
 - [UserHighlightsTweetsResponse](docs/UserHighlightsTweetsResponse.md)
 - [UserHighlightsTweetsResult](docs/UserHighlightsTweetsResult.md)
 - [UserHighlightsTweetsTimeline](docs/UserHighlightsTweetsTimeline.md)
 - [UserHighlightsTweetsUser](docs/UserHighlightsTweetsUser.md)
 - [UserLegacy](docs/UserLegacy.md)
 - [UserLegacyExtendedProfile](docs/UserLegacyExtendedProfile.md)
 - [UserLegacyExtendedProfileBirthdate](docs/UserLegacyExtendedProfileBirthdate.md)
 - [UserProfessional](docs/UserProfessional.md)
 - [UserProfessionalCategory](docs/UserProfessionalCategory.md)
 - [UserResponse](docs/UserResponse.md)
 - [UserResponseData](docs/UserResponseData.md)
 - [UserResultByScreenName](docs/UserResultByScreenName.md)
 - [UserResultByScreenNameLegacy](docs/UserResultByScreenNameLegacy.md)
 - [UserResultByScreenNameResult](docs/UserResultByScreenNameResult.md)
 - [UserResultCore](docs/UserResultCore.md)
 - [UserResults](docs/UserResults.md)
 - [UserTipJarSettings](docs/UserTipJarSettings.md)
 - [UserTweetsData](docs/UserTweetsData.md)
 - [UserTweetsResponse](docs/UserTweetsResponse.md)
 - [UserTweetsResult](docs/UserTweetsResult.md)
 - [UserTweetsUser](docs/UserTweetsUser.md)
 - [UserUnavailable](docs/UserUnavailable.md)
 - [UserUnion](docs/UserUnion.md)
 - [UserValue](docs/UserValue.md)
 - [UserVerificationInfo](docs/UserVerificationInfo.md)
 - [UserVerificationInfoReason](docs/UserVerificationInfoReason.md)
 - [UserVerificationInfoReasonDescription](docs/UserVerificationInfoReasonDescription.md)
 - [UserVerificationInfoReasonDescriptionEntities](docs/UserVerificationInfoReasonDescriptionEntities.md)
 - [UserVerificationInfoReasonDescriptionEntitiesRef](docs/UserVerificationInfoReasonDescriptionEntitiesRef.md)
 - [UsersResponse](docs/UsersResponse.md)
 - [UsersResponseData](docs/UsersResponseData.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Accept

- **Type**: API key
- **API key parameter name**: Accept
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Accept and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"Accept": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### AcceptEncoding

- **Type**: API key
- **API key parameter name**: Accept-Encoding
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: AcceptEncoding and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"AcceptEncoding": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### AcceptLanguage

- **Type**: API key
- **API key parameter name**: Accept-Language
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: AcceptLanguage and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"AcceptLanguage": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### ActiveUser

- **Type**: API key
- **API key parameter name**: x-twitter-active-user
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ActiveUser and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"ActiveUser": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### AuthType

- **Type**: API key
- **API key parameter name**: x-twitter-auth-type
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: AuthType and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"AuthType": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### ClientLanguage

- **Type**: API key
- **API key parameter name**: x-twitter-client-language
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ClientLanguage and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"ClientLanguage": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### ClientTransactionId

- **Type**: API key
- **API key parameter name**: x-client-transaction-id
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ClientTransactionId and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"ClientTransactionId": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### ClientUuid

- **Type**: API key
- **API key parameter name**: x-client-uuid
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ClientUuid and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"ClientUuid": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### CookieAuthToken

- **Type**: API key
- **API key parameter name**: auth_token
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: CookieAuthToken and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"CookieAuthToken": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### CookieCt0

- **Type**: API key
- **API key parameter name**: ct0
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: CookieCt0 and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"CookieCt0": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### CookieGt0

- **Type**: API key
- **API key parameter name**: gt0
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: CookieGt0 and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"CookieGt0": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### CsrfToken

- **Type**: API key
- **API key parameter name**: x-csrf-token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: CsrfToken and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"CsrfToken": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### GuestToken

- **Type**: API key
- **API key parameter name**: x-guest-token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: GuestToken and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"GuestToken": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### Referer

- **Type**: API key
- **API key parameter name**: Referer
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Referer and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"Referer": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### SecChUa

- **Type**: API key
- **API key parameter name**: Sec-Ch-Ua
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: SecChUa and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"SecChUa": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### SecChUaMobile

- **Type**: API key
- **API key parameter name**: Sec-Ch-Ua-Mobile
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: SecChUaMobile and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"SecChUaMobile": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### SecChUaPlatform

- **Type**: API key
- **API key parameter name**: Sec-Ch-Ua-Platform
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: SecChUaPlatform and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"SecChUaPlatform": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### SecFetchDest

- **Type**: API key
- **API key parameter name**: Sec-Fetch-Dest
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: SecFetchDest and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"SecFetchDest": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### SecFetchMode

- **Type**: API key
- **API key parameter name**: Sec-Fetch-Mode
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: SecFetchMode and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"SecFetchMode": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### SecFetchSite

- **Type**: API key
- **API key parameter name**: Sec-Fetch-Site
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: SecFetchSite and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"SecFetchSite": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### UserAgent

- **Type**: API key
- **API key parameter name**: user-agent
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: UserAgent and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"UserAgent": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

yuki@yuki0311.com

