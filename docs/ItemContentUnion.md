# ItemContentUnion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**Highlights** | Pointer to [**Highlight**](Highlight.md) |  | [optional] 
**ItemType** | [**ContentItemType**](ContentItemType.md) |  | 
**PromotedMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**SocialContext** | Pointer to [**SocialContextUnion**](SocialContextUnion.md) |  | [optional] 
**TweetDisplayType** | **string** |  | 
**TweetResults** | [**ItemResult**](ItemResult.md) |  | 
**CursorType** | [**CursorType**](CursorType.md) |  | 
**DisplayTreatment** | Pointer to [**DisplayTreatment**](DisplayTreatment.md) |  | [optional] 
**EntryType** | Pointer to [**ContentEntryType**](ContentEntryType.md) |  | [optional] 
**StopOnEmptyResponse** | Pointer to **bool** |  | [optional] 
**Value** | **string** |  | 
**UserDisplayType** | **string** |  | 
**UserResults** | [**UserResults**](UserResults.md) |  | 

## Methods

### NewItemContentUnion

`func NewItemContentUnion(typename TypeName, itemType ContentItemType, tweetDisplayType string, tweetResults ItemResult, cursorType CursorType, value string, userDisplayType string, userResults UserResults, ) *ItemContentUnion`

NewItemContentUnion instantiates a new ItemContentUnion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemContentUnionWithDefaults

`func NewItemContentUnionWithDefaults() *ItemContentUnion`

NewItemContentUnionWithDefaults instantiates a new ItemContentUnion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *ItemContentUnion) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *ItemContentUnion) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *ItemContentUnion) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetHighlights

`func (o *ItemContentUnion) GetHighlights() Highlight`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *ItemContentUnion) GetHighlightsOk() (*Highlight, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *ItemContentUnion) SetHighlights(v Highlight)`

SetHighlights sets Highlights field to given value.

### HasHighlights

`func (o *ItemContentUnion) HasHighlights() bool`

HasHighlights returns a boolean if a field has been set.

### GetItemType

`func (o *ItemContentUnion) GetItemType() ContentItemType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ItemContentUnion) GetItemTypeOk() (*ContentItemType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ItemContentUnion) SetItemType(v ContentItemType)`

SetItemType sets ItemType field to given value.


### GetPromotedMetadata

`func (o *ItemContentUnion) GetPromotedMetadata() map[string]interface{}`

GetPromotedMetadata returns the PromotedMetadata field if non-nil, zero value otherwise.

### GetPromotedMetadataOk

`func (o *ItemContentUnion) GetPromotedMetadataOk() (*map[string]interface{}, bool)`

GetPromotedMetadataOk returns a tuple with the PromotedMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedMetadata

`func (o *ItemContentUnion) SetPromotedMetadata(v map[string]interface{})`

SetPromotedMetadata sets PromotedMetadata field to given value.

### HasPromotedMetadata

`func (o *ItemContentUnion) HasPromotedMetadata() bool`

HasPromotedMetadata returns a boolean if a field has been set.

### GetSocialContext

`func (o *ItemContentUnion) GetSocialContext() SocialContextUnion`

GetSocialContext returns the SocialContext field if non-nil, zero value otherwise.

### GetSocialContextOk

`func (o *ItemContentUnion) GetSocialContextOk() (*SocialContextUnion, bool)`

GetSocialContextOk returns a tuple with the SocialContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialContext

`func (o *ItemContentUnion) SetSocialContext(v SocialContextUnion)`

SetSocialContext sets SocialContext field to given value.

### HasSocialContext

`func (o *ItemContentUnion) HasSocialContext() bool`

HasSocialContext returns a boolean if a field has been set.

### GetTweetDisplayType

`func (o *ItemContentUnion) GetTweetDisplayType() string`

GetTweetDisplayType returns the TweetDisplayType field if non-nil, zero value otherwise.

### GetTweetDisplayTypeOk

`func (o *ItemContentUnion) GetTweetDisplayTypeOk() (*string, bool)`

GetTweetDisplayTypeOk returns a tuple with the TweetDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweetDisplayType

`func (o *ItemContentUnion) SetTweetDisplayType(v string)`

SetTweetDisplayType sets TweetDisplayType field to given value.


### GetTweetResults

`func (o *ItemContentUnion) GetTweetResults() ItemResult`

GetTweetResults returns the TweetResults field if non-nil, zero value otherwise.

### GetTweetResultsOk

`func (o *ItemContentUnion) GetTweetResultsOk() (*ItemResult, bool)`

GetTweetResultsOk returns a tuple with the TweetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweetResults

`func (o *ItemContentUnion) SetTweetResults(v ItemResult)`

SetTweetResults sets TweetResults field to given value.


### GetCursorType

`func (o *ItemContentUnion) GetCursorType() CursorType`

GetCursorType returns the CursorType field if non-nil, zero value otherwise.

### GetCursorTypeOk

`func (o *ItemContentUnion) GetCursorTypeOk() (*CursorType, bool)`

GetCursorTypeOk returns a tuple with the CursorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorType

`func (o *ItemContentUnion) SetCursorType(v CursorType)`

SetCursorType sets CursorType field to given value.


### GetDisplayTreatment

`func (o *ItemContentUnion) GetDisplayTreatment() DisplayTreatment`

GetDisplayTreatment returns the DisplayTreatment field if non-nil, zero value otherwise.

### GetDisplayTreatmentOk

`func (o *ItemContentUnion) GetDisplayTreatmentOk() (*DisplayTreatment, bool)`

GetDisplayTreatmentOk returns a tuple with the DisplayTreatment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTreatment

`func (o *ItemContentUnion) SetDisplayTreatment(v DisplayTreatment)`

SetDisplayTreatment sets DisplayTreatment field to given value.

### HasDisplayTreatment

`func (o *ItemContentUnion) HasDisplayTreatment() bool`

HasDisplayTreatment returns a boolean if a field has been set.

### GetEntryType

`func (o *ItemContentUnion) GetEntryType() ContentEntryType`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *ItemContentUnion) GetEntryTypeOk() (*ContentEntryType, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *ItemContentUnion) SetEntryType(v ContentEntryType)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *ItemContentUnion) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetStopOnEmptyResponse

`func (o *ItemContentUnion) GetStopOnEmptyResponse() bool`

GetStopOnEmptyResponse returns the StopOnEmptyResponse field if non-nil, zero value otherwise.

### GetStopOnEmptyResponseOk

`func (o *ItemContentUnion) GetStopOnEmptyResponseOk() (*bool, bool)`

GetStopOnEmptyResponseOk returns a tuple with the StopOnEmptyResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnEmptyResponse

`func (o *ItemContentUnion) SetStopOnEmptyResponse(v bool)`

SetStopOnEmptyResponse sets StopOnEmptyResponse field to given value.

### HasStopOnEmptyResponse

`func (o *ItemContentUnion) HasStopOnEmptyResponse() bool`

HasStopOnEmptyResponse returns a boolean if a field has been set.

### GetValue

`func (o *ItemContentUnion) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemContentUnion) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemContentUnion) SetValue(v string)`

SetValue sets Value field to given value.


### GetUserDisplayType

`func (o *ItemContentUnion) GetUserDisplayType() string`

GetUserDisplayType returns the UserDisplayType field if non-nil, zero value otherwise.

### GetUserDisplayTypeOk

`func (o *ItemContentUnion) GetUserDisplayTypeOk() (*string, bool)`

GetUserDisplayTypeOk returns a tuple with the UserDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayType

`func (o *ItemContentUnion) SetUserDisplayType(v string)`

SetUserDisplayType sets UserDisplayType field to given value.


### GetUserResults

`func (o *ItemContentUnion) GetUserResults() UserResults`

GetUserResults returns the UserResults field if non-nil, zero value otherwise.

### GetUserResultsOk

`func (o *ItemContentUnion) GetUserResultsOk() (*UserResults, bool)`

GetUserResultsOk returns a tuple with the UserResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResults

`func (o *ItemContentUnion) SetUserResults(v UserResults)`

SetUserResults sets UserResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


