# UsersResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]UserResults**](UserResults.md) |  | 

## Methods

### NewUsersResponseData

`func NewUsersResponseData(users []UserResults, ) *UsersResponseData`

NewUsersResponseData instantiates a new UsersResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersResponseDataWithDefaults

`func NewUsersResponseDataWithDefaults() *UsersResponseData`

NewUsersResponseDataWithDefaults instantiates a new UsersResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UsersResponseData) GetUsers() []UserResults`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersResponseData) GetUsersOk() (*[]UserResults, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersResponseData) SetUsers(v []UserResults)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


