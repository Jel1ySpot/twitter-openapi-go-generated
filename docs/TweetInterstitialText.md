# TweetInterstitialText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]TweetInterstitialTextEntity**](TweetInterstitialTextEntity.md) |  | 
**Rtl** | **bool** |  | 
**Text** | **string** |  | 

## Methods

### NewTweetInterstitialText

`func NewTweetInterstitialText(entities []TweetInterstitialTextEntity, rtl bool, text string, ) *TweetInterstitialText`

NewTweetInterstitialText instantiates a new TweetInterstitialText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetInterstitialTextWithDefaults

`func NewTweetInterstitialTextWithDefaults() *TweetInterstitialText`

NewTweetInterstitialTextWithDefaults instantiates a new TweetInterstitialText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *TweetInterstitialText) GetEntities() []TweetInterstitialTextEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *TweetInterstitialText) GetEntitiesOk() (*[]TweetInterstitialTextEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *TweetInterstitialText) SetEntities(v []TweetInterstitialTextEntity)`

SetEntities sets Entities field to given value.


### GetRtl

`func (o *TweetInterstitialText) GetRtl() bool`

GetRtl returns the Rtl field if non-nil, zero value otherwise.

### GetRtlOk

`func (o *TweetInterstitialText) GetRtlOk() (*bool, bool)`

GetRtlOk returns a tuple with the Rtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtl

`func (o *TweetInterstitialText) SetRtl(v bool)`

SetRtl sets Rtl field to given value.


### GetText

`func (o *TweetInterstitialText) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TweetInterstitialText) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TweetInterstitialText) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


