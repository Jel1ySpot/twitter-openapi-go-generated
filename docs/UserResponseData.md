# UserResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserResults**](UserResults.md) |  | [optional] 

## Methods

### NewUserResponseData

`func NewUserResponseData() *UserResponseData`

NewUserResponseData instantiates a new UserResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseDataWithDefaults

`func NewUserResponseDataWithDefaults() *UserResponseData`

NewUserResponseDataWithDefaults instantiates a new UserResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserResponseData) GetUser() UserResults`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserResponseData) GetUserOk() (*UserResults, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserResponseData) SetUser(v UserResults)`

SetUser sets User field to given value.

### HasUser

`func (o *UserResponseData) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


