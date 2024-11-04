# Entities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hashtags** | **[]map[string]interface{}** |  | 
**Media** | Pointer to [**[]Media**](Media.md) |  | [optional] 
**Symbols** | **[]map[string]interface{}** |  | 
**Timestamps** | Pointer to [**[]Timestamp**](Timestamp.md) |  | [optional] 
**Urls** | [**[]Url**](Url.md) |  | 
**UserMentions** | **[]map[string]interface{}** |  | 

## Methods

### NewEntities

`func NewEntities(hashtags []map[string]interface{}, symbols []map[string]interface{}, urls []Url, userMentions []map[string]interface{}, ) *Entities`

NewEntities instantiates a new Entities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitiesWithDefaults

`func NewEntitiesWithDefaults() *Entities`

NewEntitiesWithDefaults instantiates a new Entities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashtags

`func (o *Entities) GetHashtags() []map[string]interface{}`

GetHashtags returns the Hashtags field if non-nil, zero value otherwise.

### GetHashtagsOk

`func (o *Entities) GetHashtagsOk() (*[]map[string]interface{}, bool)`

GetHashtagsOk returns a tuple with the Hashtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashtags

`func (o *Entities) SetHashtags(v []map[string]interface{})`

SetHashtags sets Hashtags field to given value.


### GetMedia

`func (o *Entities) GetMedia() []Media`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *Entities) GetMediaOk() (*[]Media, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *Entities) SetMedia(v []Media)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *Entities) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetSymbols

`func (o *Entities) GetSymbols() []map[string]interface{}`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *Entities) GetSymbolsOk() (*[]map[string]interface{}, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *Entities) SetSymbols(v []map[string]interface{})`

SetSymbols sets Symbols field to given value.


### GetTimestamps

`func (o *Entities) GetTimestamps() []Timestamp`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *Entities) GetTimestampsOk() (*[]Timestamp, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *Entities) SetTimestamps(v []Timestamp)`

SetTimestamps sets Timestamps field to given value.

### HasTimestamps

`func (o *Entities) HasTimestamps() bool`

HasTimestamps returns a boolean if a field has been set.

### GetUrls

`func (o *Entities) GetUrls() []Url`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Entities) GetUrlsOk() (*[]Url, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Entities) SetUrls(v []Url)`

SetUrls sets Urls field to given value.


### GetUserMentions

`func (o *Entities) GetUserMentions() []map[string]interface{}`

GetUserMentions returns the UserMentions field if non-nil, zero value otherwise.

### GetUserMentionsOk

`func (o *Entities) GetUserMentionsOk() (*[]map[string]interface{}, bool)`

GetUserMentionsOk returns a tuple with the UserMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMentions

`func (o *Entities) SetUserMentions(v []map[string]interface{})`

SetUserMentions sets UserMentions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


