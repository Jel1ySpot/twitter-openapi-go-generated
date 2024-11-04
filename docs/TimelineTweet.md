# TimelineTweet

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

## Methods

### NewTimelineTweet

`func NewTimelineTweet(typename TypeName, itemType ContentItemType, tweetDisplayType string, tweetResults ItemResult, ) *TimelineTweet`

NewTimelineTweet instantiates a new TimelineTweet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineTweetWithDefaults

`func NewTimelineTweetWithDefaults() *TimelineTweet`

NewTimelineTweetWithDefaults instantiates a new TimelineTweet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *TimelineTweet) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *TimelineTweet) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *TimelineTweet) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetHighlights

`func (o *TimelineTweet) GetHighlights() Highlight`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *TimelineTweet) GetHighlightsOk() (*Highlight, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *TimelineTweet) SetHighlights(v Highlight)`

SetHighlights sets Highlights field to given value.

### HasHighlights

`func (o *TimelineTweet) HasHighlights() bool`

HasHighlights returns a boolean if a field has been set.

### GetItemType

`func (o *TimelineTweet) GetItemType() ContentItemType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *TimelineTweet) GetItemTypeOk() (*ContentItemType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *TimelineTweet) SetItemType(v ContentItemType)`

SetItemType sets ItemType field to given value.


### GetPromotedMetadata

`func (o *TimelineTweet) GetPromotedMetadata() map[string]interface{}`

GetPromotedMetadata returns the PromotedMetadata field if non-nil, zero value otherwise.

### GetPromotedMetadataOk

`func (o *TimelineTweet) GetPromotedMetadataOk() (*map[string]interface{}, bool)`

GetPromotedMetadataOk returns a tuple with the PromotedMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedMetadata

`func (o *TimelineTweet) SetPromotedMetadata(v map[string]interface{})`

SetPromotedMetadata sets PromotedMetadata field to given value.

### HasPromotedMetadata

`func (o *TimelineTweet) HasPromotedMetadata() bool`

HasPromotedMetadata returns a boolean if a field has been set.

### GetSocialContext

`func (o *TimelineTweet) GetSocialContext() SocialContextUnion`

GetSocialContext returns the SocialContext field if non-nil, zero value otherwise.

### GetSocialContextOk

`func (o *TimelineTweet) GetSocialContextOk() (*SocialContextUnion, bool)`

GetSocialContextOk returns a tuple with the SocialContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialContext

`func (o *TimelineTweet) SetSocialContext(v SocialContextUnion)`

SetSocialContext sets SocialContext field to given value.

### HasSocialContext

`func (o *TimelineTweet) HasSocialContext() bool`

HasSocialContext returns a boolean if a field has been set.

### GetTweetDisplayType

`func (o *TimelineTweet) GetTweetDisplayType() string`

GetTweetDisplayType returns the TweetDisplayType field if non-nil, zero value otherwise.

### GetTweetDisplayTypeOk

`func (o *TimelineTweet) GetTweetDisplayTypeOk() (*string, bool)`

GetTweetDisplayTypeOk returns a tuple with the TweetDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweetDisplayType

`func (o *TimelineTweet) SetTweetDisplayType(v string)`

SetTweetDisplayType sets TweetDisplayType field to given value.


### GetTweetResults

`func (o *TimelineTweet) GetTweetResults() ItemResult`

GetTweetResults returns the TweetResults field if non-nil, zero value otherwise.

### GetTweetResultsOk

`func (o *TimelineTweet) GetTweetResultsOk() (*ItemResult, bool)`

GetTweetResultsOk returns a tuple with the TweetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweetResults

`func (o *TimelineTweet) SetTweetResults(v ItemResult)`

SetTweetResults sets TweetResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


