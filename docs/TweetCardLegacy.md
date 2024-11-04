# TweetCardLegacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingValues** | [**[]TweetCardLegacyBindingValue**](TweetCardLegacyBindingValue.md) |  | 
**CardPlatform** | Pointer to [**TweetCardPlatformData**](TweetCardPlatformData.md) |  | [optional] 
**Name** | **string** |  | 
**Url** | **string** |  | 
**UserRefsResults** | Pointer to [**[]UserResults**](UserResults.md) |  | [optional] 

## Methods

### NewTweetCardLegacy

`func NewTweetCardLegacy(bindingValues []TweetCardLegacyBindingValue, name string, url string, ) *TweetCardLegacy`

NewTweetCardLegacy instantiates a new TweetCardLegacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetCardLegacyWithDefaults

`func NewTweetCardLegacyWithDefaults() *TweetCardLegacy`

NewTweetCardLegacyWithDefaults instantiates a new TweetCardLegacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingValues

`func (o *TweetCardLegacy) GetBindingValues() []TweetCardLegacyBindingValue`

GetBindingValues returns the BindingValues field if non-nil, zero value otherwise.

### GetBindingValuesOk

`func (o *TweetCardLegacy) GetBindingValuesOk() (*[]TweetCardLegacyBindingValue, bool)`

GetBindingValuesOk returns a tuple with the BindingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingValues

`func (o *TweetCardLegacy) SetBindingValues(v []TweetCardLegacyBindingValue)`

SetBindingValues sets BindingValues field to given value.


### GetCardPlatform

`func (o *TweetCardLegacy) GetCardPlatform() TweetCardPlatformData`

GetCardPlatform returns the CardPlatform field if non-nil, zero value otherwise.

### GetCardPlatformOk

`func (o *TweetCardLegacy) GetCardPlatformOk() (*TweetCardPlatformData, bool)`

GetCardPlatformOk returns a tuple with the CardPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPlatform

`func (o *TweetCardLegacy) SetCardPlatform(v TweetCardPlatformData)`

SetCardPlatform sets CardPlatform field to given value.

### HasCardPlatform

`func (o *TweetCardLegacy) HasCardPlatform() bool`

HasCardPlatform returns a boolean if a field has been set.

### GetName

`func (o *TweetCardLegacy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TweetCardLegacy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TweetCardLegacy) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *TweetCardLegacy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TweetCardLegacy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TweetCardLegacy) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUserRefsResults

`func (o *TweetCardLegacy) GetUserRefsResults() []UserResults`

GetUserRefsResults returns the UserRefsResults field if non-nil, zero value otherwise.

### GetUserRefsResultsOk

`func (o *TweetCardLegacy) GetUserRefsResultsOk() (*[]UserResults, bool)`

GetUserRefsResultsOk returns a tuple with the UserRefsResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRefsResults

`func (o *TweetCardLegacy) SetUserRefsResults(v []UserResults)`

SetUserRefsResults sets UserRefsResults field to given value.

### HasUserRefsResults

`func (o *TweetCardLegacy) HasUserRefsResults() bool`

HasUserRefsResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


