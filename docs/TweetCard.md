# TweetCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Legacy** | Pointer to [**TweetCardLegacy**](TweetCardLegacy.md) |  | [optional] 
**RestId** | Pointer to **string** |  | [optional] 

## Methods

### NewTweetCard

`func NewTweetCard() *TweetCard`

NewTweetCard instantiates a new TweetCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetCardWithDefaults

`func NewTweetCardWithDefaults() *TweetCard`

NewTweetCardWithDefaults instantiates a new TweetCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegacy

`func (o *TweetCard) GetLegacy() TweetCardLegacy`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *TweetCard) GetLegacyOk() (*TweetCardLegacy, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *TweetCard) SetLegacy(v TweetCardLegacy)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *TweetCard) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetRestId

`func (o *TweetCard) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *TweetCard) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *TweetCard) SetRestId(v string)`

SetRestId sets RestId field to given value.

### HasRestId

`func (o *TweetCard) HasRestId() bool`

HasRestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


