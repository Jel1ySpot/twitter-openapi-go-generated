# UserResultByScreenNameResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**Id** | **string** |  | 
**Legacy** | [**UserResultByScreenNameLegacy**](UserResultByScreenNameLegacy.md) |  | 
**Profilemodules** | **map[string]interface{}** |  | 
**RestId** | **string** |  | 

## Methods

### NewUserResultByScreenNameResult

`func NewUserResultByScreenNameResult(typename TypeName, id string, legacy UserResultByScreenNameLegacy, profilemodules map[string]interface{}, restId string, ) *UserResultByScreenNameResult`

NewUserResultByScreenNameResult instantiates a new UserResultByScreenNameResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResultByScreenNameResultWithDefaults

`func NewUserResultByScreenNameResultWithDefaults() *UserResultByScreenNameResult`

NewUserResultByScreenNameResultWithDefaults instantiates a new UserResultByScreenNameResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *UserResultByScreenNameResult) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *UserResultByScreenNameResult) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *UserResultByScreenNameResult) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetId

`func (o *UserResultByScreenNameResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResultByScreenNameResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResultByScreenNameResult) SetId(v string)`

SetId sets Id field to given value.


### GetLegacy

`func (o *UserResultByScreenNameResult) GetLegacy() UserResultByScreenNameLegacy`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *UserResultByScreenNameResult) GetLegacyOk() (*UserResultByScreenNameLegacy, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *UserResultByScreenNameResult) SetLegacy(v UserResultByScreenNameLegacy)`

SetLegacy sets Legacy field to given value.


### GetProfilemodules

`func (o *UserResultByScreenNameResult) GetProfilemodules() map[string]interface{}`

GetProfilemodules returns the Profilemodules field if non-nil, zero value otherwise.

### GetProfilemodulesOk

`func (o *UserResultByScreenNameResult) GetProfilemodulesOk() (*map[string]interface{}, bool)`

GetProfilemodulesOk returns a tuple with the Profilemodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilemodules

`func (o *UserResultByScreenNameResult) SetProfilemodules(v map[string]interface{})`

SetProfilemodules sets Profilemodules field to given value.


### GetRestId

`func (o *UserResultByScreenNameResult) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *UserResultByScreenNameResult) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *UserResultByScreenNameResult) SetRestId(v string)`

SetRestId sets RestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


