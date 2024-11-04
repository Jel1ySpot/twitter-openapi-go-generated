# UserProfessional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**[]UserProfessionalCategory**](UserProfessionalCategory.md) |  | 
**ProfessionalType** | **string** |  | 
**RestId** | **string** |  | 

## Methods

### NewUserProfessional

`func NewUserProfessional(category []UserProfessionalCategory, professionalType string, restId string, ) *UserProfessional`

NewUserProfessional instantiates a new UserProfessional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfessionalWithDefaults

`func NewUserProfessionalWithDefaults() *UserProfessional`

NewUserProfessionalWithDefaults instantiates a new UserProfessional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *UserProfessional) GetCategory() []UserProfessionalCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UserProfessional) GetCategoryOk() (*[]UserProfessionalCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UserProfessional) SetCategory(v []UserProfessionalCategory)`

SetCategory sets Category field to given value.


### GetProfessionalType

`func (o *UserProfessional) GetProfessionalType() string`

GetProfessionalType returns the ProfessionalType field if non-nil, zero value otherwise.

### GetProfessionalTypeOk

`func (o *UserProfessional) GetProfessionalTypeOk() (*string, bool)`

GetProfessionalTypeOk returns a tuple with the ProfessionalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalType

`func (o *UserProfessional) SetProfessionalType(v string)`

SetProfessionalType sets ProfessionalType field to given value.


### GetRestId

`func (o *UserProfessional) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *UserProfessional) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *UserProfessional) SetRestId(v string)`

SetRestId sets RestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


