# UserUnavailable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**Message** | Pointer to **string** |  | [optional] 
**Reason** | **string** |  | 

## Methods

### NewUserUnavailable

`func NewUserUnavailable(typename TypeName, reason string, ) *UserUnavailable`

NewUserUnavailable instantiates a new UserUnavailable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUnavailableWithDefaults

`func NewUserUnavailableWithDefaults() *UserUnavailable`

NewUserUnavailableWithDefaults instantiates a new UserUnavailable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *UserUnavailable) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *UserUnavailable) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *UserUnavailable) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetMessage

`func (o *UserUnavailable) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserUnavailable) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserUnavailable) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserUnavailable) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *UserUnavailable) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UserUnavailable) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UserUnavailable) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


