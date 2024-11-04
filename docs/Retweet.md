# Retweet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Legacy** | [**RetweetLegacy**](RetweetLegacy.md) |  | 
**RestId** | **string** |  | 

## Methods

### NewRetweet

`func NewRetweet(legacy RetweetLegacy, restId string, ) *Retweet`

NewRetweet instantiates a new Retweet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetweetWithDefaults

`func NewRetweetWithDefaults() *Retweet`

NewRetweetWithDefaults instantiates a new Retweet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegacy

`func (o *Retweet) GetLegacy() RetweetLegacy`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *Retweet) GetLegacyOk() (*RetweetLegacy, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *Retweet) SetLegacy(v RetweetLegacy)`

SetLegacy sets Legacy field to given value.


### GetRestId

`func (o *Retweet) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *Retweet) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *Retweet) SetRestId(v string)`

SetRestId sets RestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


