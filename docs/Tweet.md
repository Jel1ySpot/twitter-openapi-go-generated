# Tweet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | Pointer to [**TypeName**](TypeName.md) |  | [optional] 
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

## Methods

### NewTweet

`func NewTweet(restId string, ) *Tweet`

NewTweet instantiates a new Tweet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetWithDefaults

`func NewTweetWithDefaults() *Tweet`

NewTweetWithDefaults instantiates a new Tweet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *Tweet) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *Tweet) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *Tweet) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.

### HasTypename

`func (o *Tweet) HasTypename() bool`

HasTypename returns a boolean if a field has been set.

### GetArticle

`func (o *Tweet) GetArticle() Article`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *Tweet) GetArticleOk() (*Article, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *Tweet) SetArticle(v Article)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *Tweet) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### GetAuthorCommunityRelationship

`func (o *Tweet) GetAuthorCommunityRelationship() AuthorCommunityRelationship`

GetAuthorCommunityRelationship returns the AuthorCommunityRelationship field if non-nil, zero value otherwise.

### GetAuthorCommunityRelationshipOk

`func (o *Tweet) GetAuthorCommunityRelationshipOk() (*AuthorCommunityRelationship, bool)`

GetAuthorCommunityRelationshipOk returns a tuple with the AuthorCommunityRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorCommunityRelationship

`func (o *Tweet) SetAuthorCommunityRelationship(v AuthorCommunityRelationship)`

SetAuthorCommunityRelationship sets AuthorCommunityRelationship field to given value.

### HasAuthorCommunityRelationship

`func (o *Tweet) HasAuthorCommunityRelationship() bool`

HasAuthorCommunityRelationship returns a boolean if a field has been set.

### GetBirdwatchPivot

`func (o *Tweet) GetBirdwatchPivot() BirdwatchPivot`

GetBirdwatchPivot returns the BirdwatchPivot field if non-nil, zero value otherwise.

### GetBirdwatchPivotOk

`func (o *Tweet) GetBirdwatchPivotOk() (*BirdwatchPivot, bool)`

GetBirdwatchPivotOk returns a tuple with the BirdwatchPivot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirdwatchPivot

`func (o *Tweet) SetBirdwatchPivot(v BirdwatchPivot)`

SetBirdwatchPivot sets BirdwatchPivot field to given value.

### HasBirdwatchPivot

`func (o *Tweet) HasBirdwatchPivot() bool`

HasBirdwatchPivot returns a boolean if a field has been set.

### GetCard

`func (o *Tweet) GetCard() TweetCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *Tweet) GetCardOk() (*TweetCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *Tweet) SetCard(v TweetCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *Tweet) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCommunityRelationship

`func (o *Tweet) GetCommunityRelationship() CommunityRelationship`

GetCommunityRelationship returns the CommunityRelationship field if non-nil, zero value otherwise.

### GetCommunityRelationshipOk

`func (o *Tweet) GetCommunityRelationshipOk() (*CommunityRelationship, bool)`

GetCommunityRelationshipOk returns a tuple with the CommunityRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityRelationship

`func (o *Tweet) SetCommunityRelationship(v CommunityRelationship)`

SetCommunityRelationship sets CommunityRelationship field to given value.

### HasCommunityRelationship

`func (o *Tweet) HasCommunityRelationship() bool`

HasCommunityRelationship returns a boolean if a field has been set.

### GetCommunityResults

`func (o *Tweet) GetCommunityResults() Community`

GetCommunityResults returns the CommunityResults field if non-nil, zero value otherwise.

### GetCommunityResultsOk

`func (o *Tweet) GetCommunityResultsOk() (*Community, bool)`

GetCommunityResultsOk returns a tuple with the CommunityResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityResults

`func (o *Tweet) SetCommunityResults(v Community)`

SetCommunityResults sets CommunityResults field to given value.

### HasCommunityResults

`func (o *Tweet) HasCommunityResults() bool`

HasCommunityResults returns a boolean if a field has been set.

### GetCore

`func (o *Tweet) GetCore() UserResultCore`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *Tweet) GetCoreOk() (*UserResultCore, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *Tweet) SetCore(v UserResultCore)`

SetCore sets Core field to given value.

### HasCore

`func (o *Tweet) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetEditControl

`func (o *Tweet) GetEditControl() TweetEditControl`

GetEditControl returns the EditControl field if non-nil, zero value otherwise.

### GetEditControlOk

`func (o *Tweet) GetEditControlOk() (*TweetEditControl, bool)`

GetEditControlOk returns a tuple with the EditControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditControl

`func (o *Tweet) SetEditControl(v TweetEditControl)`

SetEditControl sets EditControl field to given value.

### HasEditControl

`func (o *Tweet) HasEditControl() bool`

HasEditControl returns a boolean if a field has been set.

### GetEditPrespective

`func (o *Tweet) GetEditPrespective() TweetEditPrespective`

GetEditPrespective returns the EditPrespective field if non-nil, zero value otherwise.

### GetEditPrespectiveOk

`func (o *Tweet) GetEditPrespectiveOk() (*TweetEditPrespective, bool)`

GetEditPrespectiveOk returns a tuple with the EditPrespective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditPrespective

`func (o *Tweet) SetEditPrespective(v TweetEditPrespective)`

SetEditPrespective sets EditPrespective field to given value.

### HasEditPrespective

`func (o *Tweet) HasEditPrespective() bool`

HasEditPrespective returns a boolean if a field has been set.

### GetHasBirdwatchNotes

`func (o *Tweet) GetHasBirdwatchNotes() bool`

GetHasBirdwatchNotes returns the HasBirdwatchNotes field if non-nil, zero value otherwise.

### GetHasBirdwatchNotesOk

`func (o *Tweet) GetHasBirdwatchNotesOk() (*bool, bool)`

GetHasBirdwatchNotesOk returns a tuple with the HasBirdwatchNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBirdwatchNotes

`func (o *Tweet) SetHasBirdwatchNotes(v bool)`

SetHasBirdwatchNotes sets HasBirdwatchNotes field to given value.

### HasHasBirdwatchNotes

`func (o *Tweet) HasHasBirdwatchNotes() bool`

HasHasBirdwatchNotes returns a boolean if a field has been set.

### GetIsTranslatable

`func (o *Tweet) GetIsTranslatable() bool`

GetIsTranslatable returns the IsTranslatable field if non-nil, zero value otherwise.

### GetIsTranslatableOk

`func (o *Tweet) GetIsTranslatableOk() (*bool, bool)`

GetIsTranslatableOk returns a tuple with the IsTranslatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTranslatable

`func (o *Tweet) SetIsTranslatable(v bool)`

SetIsTranslatable sets IsTranslatable field to given value.

### HasIsTranslatable

`func (o *Tweet) HasIsTranslatable() bool`

HasIsTranslatable returns a boolean if a field has been set.

### GetLegacy

`func (o *Tweet) GetLegacy() TweetLegacy`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *Tweet) GetLegacyOk() (*TweetLegacy, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *Tweet) SetLegacy(v TweetLegacy)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *Tweet) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetNoteTweet

`func (o *Tweet) GetNoteTweet() NoteTweet`

GetNoteTweet returns the NoteTweet field if non-nil, zero value otherwise.

### GetNoteTweetOk

`func (o *Tweet) GetNoteTweetOk() (*NoteTweet, bool)`

GetNoteTweetOk returns a tuple with the NoteTweet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteTweet

`func (o *Tweet) SetNoteTweet(v NoteTweet)`

SetNoteTweet sets NoteTweet field to given value.

### HasNoteTweet

`func (o *Tweet) HasNoteTweet() bool`

HasNoteTweet returns a boolean if a field has been set.

### GetPreviousCounts

`func (o *Tweet) GetPreviousCounts() TweetPreviousCounts`

GetPreviousCounts returns the PreviousCounts field if non-nil, zero value otherwise.

### GetPreviousCountsOk

`func (o *Tweet) GetPreviousCountsOk() (*TweetPreviousCounts, bool)`

GetPreviousCountsOk returns a tuple with the PreviousCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCounts

`func (o *Tweet) SetPreviousCounts(v TweetPreviousCounts)`

SetPreviousCounts sets PreviousCounts field to given value.

### HasPreviousCounts

`func (o *Tweet) HasPreviousCounts() bool`

HasPreviousCounts returns a boolean if a field has been set.

### GetQuickPromoteEligibility

`func (o *Tweet) GetQuickPromoteEligibility() map[string]interface{}`

GetQuickPromoteEligibility returns the QuickPromoteEligibility field if non-nil, zero value otherwise.

### GetQuickPromoteEligibilityOk

`func (o *Tweet) GetQuickPromoteEligibilityOk() (*map[string]interface{}, bool)`

GetQuickPromoteEligibilityOk returns a tuple with the QuickPromoteEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickPromoteEligibility

`func (o *Tweet) SetQuickPromoteEligibility(v map[string]interface{})`

SetQuickPromoteEligibility sets QuickPromoteEligibility field to given value.

### HasQuickPromoteEligibility

`func (o *Tweet) HasQuickPromoteEligibility() bool`

HasQuickPromoteEligibility returns a boolean if a field has been set.

### GetQuotedRefResult

`func (o *Tweet) GetQuotedRefResult() QuotedRefResult`

GetQuotedRefResult returns the QuotedRefResult field if non-nil, zero value otherwise.

### GetQuotedRefResultOk

`func (o *Tweet) GetQuotedRefResultOk() (*QuotedRefResult, bool)`

GetQuotedRefResultOk returns a tuple with the QuotedRefResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedRefResult

`func (o *Tweet) SetQuotedRefResult(v QuotedRefResult)`

SetQuotedRefResult sets QuotedRefResult field to given value.

### HasQuotedRefResult

`func (o *Tweet) HasQuotedRefResult() bool`

HasQuotedRefResult returns a boolean if a field has been set.

### GetQuotedStatusResult

`func (o *Tweet) GetQuotedStatusResult() ItemResult`

GetQuotedStatusResult returns the QuotedStatusResult field if non-nil, zero value otherwise.

### GetQuotedStatusResultOk

`func (o *Tweet) GetQuotedStatusResultOk() (*ItemResult, bool)`

GetQuotedStatusResultOk returns a tuple with the QuotedStatusResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedStatusResult

`func (o *Tweet) SetQuotedStatusResult(v ItemResult)`

SetQuotedStatusResult sets QuotedStatusResult field to given value.

### HasQuotedStatusResult

`func (o *Tweet) HasQuotedStatusResult() bool`

HasQuotedStatusResult returns a boolean if a field has been set.

### GetRestId

`func (o *Tweet) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *Tweet) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *Tweet) SetRestId(v string)`

SetRestId sets RestId field to given value.


### GetSource

`func (o *Tweet) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Tweet) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Tweet) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Tweet) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSuperFollowsReplyUserResult

`func (o *Tweet) GetSuperFollowsReplyUserResult() SuperFollowsReplyUserResult`

GetSuperFollowsReplyUserResult returns the SuperFollowsReplyUserResult field if non-nil, zero value otherwise.

### GetSuperFollowsReplyUserResultOk

`func (o *Tweet) GetSuperFollowsReplyUserResultOk() (*SuperFollowsReplyUserResult, bool)`

GetSuperFollowsReplyUserResultOk returns a tuple with the SuperFollowsReplyUserResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperFollowsReplyUserResult

`func (o *Tweet) SetSuperFollowsReplyUserResult(v SuperFollowsReplyUserResult)`

SetSuperFollowsReplyUserResult sets SuperFollowsReplyUserResult field to given value.

### HasSuperFollowsReplyUserResult

`func (o *Tweet) HasSuperFollowsReplyUserResult() bool`

HasSuperFollowsReplyUserResult returns a boolean if a field has been set.

### GetUnifiedCard

`func (o *Tweet) GetUnifiedCard() UnifiedCard`

GetUnifiedCard returns the UnifiedCard field if non-nil, zero value otherwise.

### GetUnifiedCardOk

`func (o *Tweet) GetUnifiedCardOk() (*UnifiedCard, bool)`

GetUnifiedCardOk returns a tuple with the UnifiedCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedCard

`func (o *Tweet) SetUnifiedCard(v UnifiedCard)`

SetUnifiedCard sets UnifiedCard field to given value.

### HasUnifiedCard

`func (o *Tweet) HasUnifiedCard() bool`

HasUnifiedCard returns a boolean if a field has been set.

### GetUnmentionData

`func (o *Tweet) GetUnmentionData() map[string]interface{}`

GetUnmentionData returns the UnmentionData field if non-nil, zero value otherwise.

### GetUnmentionDataOk

`func (o *Tweet) GetUnmentionDataOk() (*map[string]interface{}, bool)`

GetUnmentionDataOk returns a tuple with the UnmentionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmentionData

`func (o *Tweet) SetUnmentionData(v map[string]interface{})`

SetUnmentionData sets UnmentionData field to given value.

### HasUnmentionData

`func (o *Tweet) HasUnmentionData() bool`

HasUnmentionData returns a boolean if a field has been set.

### GetViews

`func (o *Tweet) GetViews() TweetView`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *Tweet) GetViewsOk() (*TweetView, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *Tweet) SetViews(v TweetView)`

SetViews sets Views field to given value.

### HasViews

`func (o *Tweet) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


