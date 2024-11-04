# TweetUnion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**Article** | Pointer to [**Article**](Article.md) |  | [optional] 
**AuthorCommunityRelationship** | Pointer to [**AuthorCommunityRelationship**](AuthorCommunityRelationship.md) |  | [optional] 
**BirdwatchPivot** | Pointer to [**BirdwatchPivot**](BirdwatchPivot.md) |  | [optional] 
**Card** | Pointer to [**TweetCard**](TweetCard.md) |  | [optional] 
**CommunityRelationship** | Pointer to [**CommunityRelationship**](CommunityRelationship.md) |  | [optional] 
**CommunityResults** | Pointer to [**Community**](Community.md) |  | [optional] 
**Core** | Pointer to [**UserResultCore**](UserResultCore.md) |  | [optional] 
**EditControl** | Pointer to [**TweetEditControl**](TweetEditControl.md) |  | [optional] 
**EditPrespective** | Pointer to [**TweetEditPrespective**](TweetEditPrespective.md) |  | [optional] 
**HasBirdwatchNotes** | Pointer to **bool** |  | [optional] 
**IsTranslatable** | Pointer to **bool** |  | [optional] 
**Legacy** | Pointer to [**TweetLegacy**](TweetLegacy.md) |  | [optional] 
**NoteTweet** | Pointer to [**NoteTweet**](NoteTweet.md) |  | [optional] 
**PreviousCounts** | Pointer to [**TweetPreviousCounts**](TweetPreviousCounts.md) |  | [optional] 
**QuickPromoteEligibility** | Pointer to **map[string]interface{}** |  | [optional] 
**QuotedRefResult** | Pointer to [**QuotedRefResult**](QuotedRefResult.md) |  | [optional] 
**QuotedStatusResult** | Pointer to [**ItemResult**](ItemResult.md) |  | [optional] 
**RestId** | **string** |  | 
**Source** | Pointer to **string** |  | [optional] 
**SuperFollowsReplyUserResult** | Pointer to [**SuperFollowsReplyUserResult**](SuperFollowsReplyUserResult.md) |  | [optional] 
**UnifiedCard** | Pointer to [**UnifiedCard**](UnifiedCard.md) |  | [optional] 
**UnmentionData** | Pointer to **map[string]interface{}** |  | [optional] 
**Views** | Pointer to [**TweetView**](TweetView.md) |  | [optional] 
**LimitedActionResults** | Pointer to **map[string]interface{}** |  | [optional] 
**MediaVisibilityResults** | Pointer to [**MediaVisibilityResults**](MediaVisibilityResults.md) |  | [optional] 
**Tweet** | [**Tweet**](Tweet.md) |  | 
**TweetInterstitial** | Pointer to [**TweetInterstitial**](TweetInterstitial.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewTweetUnion

`func NewTweetUnion(typename TypeName, restId string, tweet Tweet, ) *TweetUnion`

NewTweetUnion instantiates a new TweetUnion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetUnionWithDefaults

`func NewTweetUnionWithDefaults() *TweetUnion`

NewTweetUnionWithDefaults instantiates a new TweetUnion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *TweetUnion) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *TweetUnion) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *TweetUnion) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetArticle

`func (o *TweetUnion) GetArticle() Article`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *TweetUnion) GetArticleOk() (*Article, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *TweetUnion) SetArticle(v Article)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *TweetUnion) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### GetAuthorCommunityRelationship

`func (o *TweetUnion) GetAuthorCommunityRelationship() AuthorCommunityRelationship`

GetAuthorCommunityRelationship returns the AuthorCommunityRelationship field if non-nil, zero value otherwise.

### GetAuthorCommunityRelationshipOk

`func (o *TweetUnion) GetAuthorCommunityRelationshipOk() (*AuthorCommunityRelationship, bool)`

GetAuthorCommunityRelationshipOk returns a tuple with the AuthorCommunityRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorCommunityRelationship

`func (o *TweetUnion) SetAuthorCommunityRelationship(v AuthorCommunityRelationship)`

SetAuthorCommunityRelationship sets AuthorCommunityRelationship field to given value.

### HasAuthorCommunityRelationship

`func (o *TweetUnion) HasAuthorCommunityRelationship() bool`

HasAuthorCommunityRelationship returns a boolean if a field has been set.

### GetBirdwatchPivot

`func (o *TweetUnion) GetBirdwatchPivot() BirdwatchPivot`

GetBirdwatchPivot returns the BirdwatchPivot field if non-nil, zero value otherwise.

### GetBirdwatchPivotOk

`func (o *TweetUnion) GetBirdwatchPivotOk() (*BirdwatchPivot, bool)`

GetBirdwatchPivotOk returns a tuple with the BirdwatchPivot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirdwatchPivot

`func (o *TweetUnion) SetBirdwatchPivot(v BirdwatchPivot)`

SetBirdwatchPivot sets BirdwatchPivot field to given value.

### HasBirdwatchPivot

`func (o *TweetUnion) HasBirdwatchPivot() bool`

HasBirdwatchPivot returns a boolean if a field has been set.

### GetCard

`func (o *TweetUnion) GetCard() TweetCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *TweetUnion) GetCardOk() (*TweetCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *TweetUnion) SetCard(v TweetCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *TweetUnion) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCommunityRelationship

`func (o *TweetUnion) GetCommunityRelationship() CommunityRelationship`

GetCommunityRelationship returns the CommunityRelationship field if non-nil, zero value otherwise.

### GetCommunityRelationshipOk

`func (o *TweetUnion) GetCommunityRelationshipOk() (*CommunityRelationship, bool)`

GetCommunityRelationshipOk returns a tuple with the CommunityRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityRelationship

`func (o *TweetUnion) SetCommunityRelationship(v CommunityRelationship)`

SetCommunityRelationship sets CommunityRelationship field to given value.

### HasCommunityRelationship

`func (o *TweetUnion) HasCommunityRelationship() bool`

HasCommunityRelationship returns a boolean if a field has been set.

### GetCommunityResults

`func (o *TweetUnion) GetCommunityResults() Community`

GetCommunityResults returns the CommunityResults field if non-nil, zero value otherwise.

### GetCommunityResultsOk

`func (o *TweetUnion) GetCommunityResultsOk() (*Community, bool)`

GetCommunityResultsOk returns a tuple with the CommunityResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityResults

`func (o *TweetUnion) SetCommunityResults(v Community)`

SetCommunityResults sets CommunityResults field to given value.

### HasCommunityResults

`func (o *TweetUnion) HasCommunityResults() bool`

HasCommunityResults returns a boolean if a field has been set.

### GetCore

`func (o *TweetUnion) GetCore() UserResultCore`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *TweetUnion) GetCoreOk() (*UserResultCore, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *TweetUnion) SetCore(v UserResultCore)`

SetCore sets Core field to given value.

### HasCore

`func (o *TweetUnion) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetEditControl

`func (o *TweetUnion) GetEditControl() TweetEditControl`

GetEditControl returns the EditControl field if non-nil, zero value otherwise.

### GetEditControlOk

`func (o *TweetUnion) GetEditControlOk() (*TweetEditControl, bool)`

GetEditControlOk returns a tuple with the EditControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditControl

`func (o *TweetUnion) SetEditControl(v TweetEditControl)`

SetEditControl sets EditControl field to given value.

### HasEditControl

`func (o *TweetUnion) HasEditControl() bool`

HasEditControl returns a boolean if a field has been set.

### GetEditPrespective

`func (o *TweetUnion) GetEditPrespective() TweetEditPrespective`

GetEditPrespective returns the EditPrespective field if non-nil, zero value otherwise.

### GetEditPrespectiveOk

`func (o *TweetUnion) GetEditPrespectiveOk() (*TweetEditPrespective, bool)`

GetEditPrespectiveOk returns a tuple with the EditPrespective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditPrespective

`func (o *TweetUnion) SetEditPrespective(v TweetEditPrespective)`

SetEditPrespective sets EditPrespective field to given value.

### HasEditPrespective

`func (o *TweetUnion) HasEditPrespective() bool`

HasEditPrespective returns a boolean if a field has been set.

### GetHasBirdwatchNotes

`func (o *TweetUnion) GetHasBirdwatchNotes() bool`

GetHasBirdwatchNotes returns the HasBirdwatchNotes field if non-nil, zero value otherwise.

### GetHasBirdwatchNotesOk

`func (o *TweetUnion) GetHasBirdwatchNotesOk() (*bool, bool)`

GetHasBirdwatchNotesOk returns a tuple with the HasBirdwatchNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBirdwatchNotes

`func (o *TweetUnion) SetHasBirdwatchNotes(v bool)`

SetHasBirdwatchNotes sets HasBirdwatchNotes field to given value.

### HasHasBirdwatchNotes

`func (o *TweetUnion) HasHasBirdwatchNotes() bool`

HasHasBirdwatchNotes returns a boolean if a field has been set.

### GetIsTranslatable

`func (o *TweetUnion) GetIsTranslatable() bool`

GetIsTranslatable returns the IsTranslatable field if non-nil, zero value otherwise.

### GetIsTranslatableOk

`func (o *TweetUnion) GetIsTranslatableOk() (*bool, bool)`

GetIsTranslatableOk returns a tuple with the IsTranslatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTranslatable

`func (o *TweetUnion) SetIsTranslatable(v bool)`

SetIsTranslatable sets IsTranslatable field to given value.

### HasIsTranslatable

`func (o *TweetUnion) HasIsTranslatable() bool`

HasIsTranslatable returns a boolean if a field has been set.

### GetLegacy

`func (o *TweetUnion) GetLegacy() TweetLegacy`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *TweetUnion) GetLegacyOk() (*TweetLegacy, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *TweetUnion) SetLegacy(v TweetLegacy)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *TweetUnion) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetNoteTweet

`func (o *TweetUnion) GetNoteTweet() NoteTweet`

GetNoteTweet returns the NoteTweet field if non-nil, zero value otherwise.

### GetNoteTweetOk

`func (o *TweetUnion) GetNoteTweetOk() (*NoteTweet, bool)`

GetNoteTweetOk returns a tuple with the NoteTweet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteTweet

`func (o *TweetUnion) SetNoteTweet(v NoteTweet)`

SetNoteTweet sets NoteTweet field to given value.

### HasNoteTweet

`func (o *TweetUnion) HasNoteTweet() bool`

HasNoteTweet returns a boolean if a field has been set.

### GetPreviousCounts

`func (o *TweetUnion) GetPreviousCounts() TweetPreviousCounts`

GetPreviousCounts returns the PreviousCounts field if non-nil, zero value otherwise.

### GetPreviousCountsOk

`func (o *TweetUnion) GetPreviousCountsOk() (*TweetPreviousCounts, bool)`

GetPreviousCountsOk returns a tuple with the PreviousCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCounts

`func (o *TweetUnion) SetPreviousCounts(v TweetPreviousCounts)`

SetPreviousCounts sets PreviousCounts field to given value.

### HasPreviousCounts

`func (o *TweetUnion) HasPreviousCounts() bool`

HasPreviousCounts returns a boolean if a field has been set.

### GetQuickPromoteEligibility

`func (o *TweetUnion) GetQuickPromoteEligibility() map[string]interface{}`

GetQuickPromoteEligibility returns the QuickPromoteEligibility field if non-nil, zero value otherwise.

### GetQuickPromoteEligibilityOk

`func (o *TweetUnion) GetQuickPromoteEligibilityOk() (*map[string]interface{}, bool)`

GetQuickPromoteEligibilityOk returns a tuple with the QuickPromoteEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickPromoteEligibility

`func (o *TweetUnion) SetQuickPromoteEligibility(v map[string]interface{})`

SetQuickPromoteEligibility sets QuickPromoteEligibility field to given value.

### HasQuickPromoteEligibility

`func (o *TweetUnion) HasQuickPromoteEligibility() bool`

HasQuickPromoteEligibility returns a boolean if a field has been set.

### GetQuotedRefResult

`func (o *TweetUnion) GetQuotedRefResult() QuotedRefResult`

GetQuotedRefResult returns the QuotedRefResult field if non-nil, zero value otherwise.

### GetQuotedRefResultOk

`func (o *TweetUnion) GetQuotedRefResultOk() (*QuotedRefResult, bool)`

GetQuotedRefResultOk returns a tuple with the QuotedRefResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedRefResult

`func (o *TweetUnion) SetQuotedRefResult(v QuotedRefResult)`

SetQuotedRefResult sets QuotedRefResult field to given value.

### HasQuotedRefResult

`func (o *TweetUnion) HasQuotedRefResult() bool`

HasQuotedRefResult returns a boolean if a field has been set.

### GetQuotedStatusResult

`func (o *TweetUnion) GetQuotedStatusResult() ItemResult`

GetQuotedStatusResult returns the QuotedStatusResult field if non-nil, zero value otherwise.

### GetQuotedStatusResultOk

`func (o *TweetUnion) GetQuotedStatusResultOk() (*ItemResult, bool)`

GetQuotedStatusResultOk returns a tuple with the QuotedStatusResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedStatusResult

`func (o *TweetUnion) SetQuotedStatusResult(v ItemResult)`

SetQuotedStatusResult sets QuotedStatusResult field to given value.

### HasQuotedStatusResult

`func (o *TweetUnion) HasQuotedStatusResult() bool`

HasQuotedStatusResult returns a boolean if a field has been set.

### GetRestId

`func (o *TweetUnion) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *TweetUnion) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *TweetUnion) SetRestId(v string)`

SetRestId sets RestId field to given value.


### GetSource

`func (o *TweetUnion) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TweetUnion) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TweetUnion) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TweetUnion) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSuperFollowsReplyUserResult

`func (o *TweetUnion) GetSuperFollowsReplyUserResult() SuperFollowsReplyUserResult`

GetSuperFollowsReplyUserResult returns the SuperFollowsReplyUserResult field if non-nil, zero value otherwise.

### GetSuperFollowsReplyUserResultOk

`func (o *TweetUnion) GetSuperFollowsReplyUserResultOk() (*SuperFollowsReplyUserResult, bool)`

GetSuperFollowsReplyUserResultOk returns a tuple with the SuperFollowsReplyUserResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperFollowsReplyUserResult

`func (o *TweetUnion) SetSuperFollowsReplyUserResult(v SuperFollowsReplyUserResult)`

SetSuperFollowsReplyUserResult sets SuperFollowsReplyUserResult field to given value.

### HasSuperFollowsReplyUserResult

`func (o *TweetUnion) HasSuperFollowsReplyUserResult() bool`

HasSuperFollowsReplyUserResult returns a boolean if a field has been set.

### GetUnifiedCard

`func (o *TweetUnion) GetUnifiedCard() UnifiedCard`

GetUnifiedCard returns the UnifiedCard field if non-nil, zero value otherwise.

### GetUnifiedCardOk

`func (o *TweetUnion) GetUnifiedCardOk() (*UnifiedCard, bool)`

GetUnifiedCardOk returns a tuple with the UnifiedCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedCard

`func (o *TweetUnion) SetUnifiedCard(v UnifiedCard)`

SetUnifiedCard sets UnifiedCard field to given value.

### HasUnifiedCard

`func (o *TweetUnion) HasUnifiedCard() bool`

HasUnifiedCard returns a boolean if a field has been set.

### GetUnmentionData

`func (o *TweetUnion) GetUnmentionData() map[string]interface{}`

GetUnmentionData returns the UnmentionData field if non-nil, zero value otherwise.

### GetUnmentionDataOk

`func (o *TweetUnion) GetUnmentionDataOk() (*map[string]interface{}, bool)`

GetUnmentionDataOk returns a tuple with the UnmentionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmentionData

`func (o *TweetUnion) SetUnmentionData(v map[string]interface{})`

SetUnmentionData sets UnmentionData field to given value.

### HasUnmentionData

`func (o *TweetUnion) HasUnmentionData() bool`

HasUnmentionData returns a boolean if a field has been set.

### GetViews

`func (o *TweetUnion) GetViews() TweetView`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *TweetUnion) GetViewsOk() (*TweetView, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *TweetUnion) SetViews(v TweetView)`

SetViews sets Views field to given value.

### HasViews

`func (o *TweetUnion) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetLimitedActionResults

`func (o *TweetUnion) GetLimitedActionResults() map[string]interface{}`

GetLimitedActionResults returns the LimitedActionResults field if non-nil, zero value otherwise.

### GetLimitedActionResultsOk

`func (o *TweetUnion) GetLimitedActionResultsOk() (*map[string]interface{}, bool)`

GetLimitedActionResultsOk returns a tuple with the LimitedActionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedActionResults

`func (o *TweetUnion) SetLimitedActionResults(v map[string]interface{})`

SetLimitedActionResults sets LimitedActionResults field to given value.

### HasLimitedActionResults

`func (o *TweetUnion) HasLimitedActionResults() bool`

HasLimitedActionResults returns a boolean if a field has been set.

### GetMediaVisibilityResults

`func (o *TweetUnion) GetMediaVisibilityResults() MediaVisibilityResults`

GetMediaVisibilityResults returns the MediaVisibilityResults field if non-nil, zero value otherwise.

### GetMediaVisibilityResultsOk

`func (o *TweetUnion) GetMediaVisibilityResultsOk() (*MediaVisibilityResults, bool)`

GetMediaVisibilityResultsOk returns a tuple with the MediaVisibilityResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaVisibilityResults

`func (o *TweetUnion) SetMediaVisibilityResults(v MediaVisibilityResults)`

SetMediaVisibilityResults sets MediaVisibilityResults field to given value.

### HasMediaVisibilityResults

`func (o *TweetUnion) HasMediaVisibilityResults() bool`

HasMediaVisibilityResults returns a boolean if a field has been set.

### GetTweet

`func (o *TweetUnion) GetTweet() Tweet`

GetTweet returns the Tweet field if non-nil, zero value otherwise.

### GetTweetOk

`func (o *TweetUnion) GetTweetOk() (*Tweet, bool)`

GetTweetOk returns a tuple with the Tweet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweet

`func (o *TweetUnion) SetTweet(v Tweet)`

SetTweet sets Tweet field to given value.


### GetTweetInterstitial

`func (o *TweetUnion) GetTweetInterstitial() TweetInterstitial`

GetTweetInterstitial returns the TweetInterstitial field if non-nil, zero value otherwise.

### GetTweetInterstitialOk

`func (o *TweetUnion) GetTweetInterstitialOk() (*TweetInterstitial, bool)`

GetTweetInterstitialOk returns a tuple with the TweetInterstitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweetInterstitial

`func (o *TweetUnion) SetTweetInterstitial(v TweetInterstitial)`

SetTweetInterstitial sets TweetInterstitial field to given value.

### HasTweetInterstitial

`func (o *TweetUnion) HasTweetInterstitial() bool`

HasTweetInterstitial returns a boolean if a field has been set.

### GetReason

`func (o *TweetUnion) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TweetUnion) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TweetUnion) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TweetUnion) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


