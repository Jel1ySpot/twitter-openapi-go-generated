# UserVerificationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsIdentityVerified** | **bool** |  | 
**Reason** | Pointer to [**UserVerificationInfoReason**](UserVerificationInfoReason.md) |  | [optional] 

## Methods

### NewUserVerificationInfo

`func NewUserVerificationInfo(isIdentityVerified bool, ) *UserVerificationInfo`

NewUserVerificationInfo instantiates a new UserVerificationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserVerificationInfoWithDefaults

`func NewUserVerificationInfoWithDefaults() *UserVerificationInfo`

NewUserVerificationInfoWithDefaults instantiates a new UserVerificationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsIdentityVerified

`func (o *UserVerificationInfo) GetIsIdentityVerified() bool`

GetIsIdentityVerified returns the IsIdentityVerified field if non-nil, zero value otherwise.

### GetIsIdentityVerifiedOk

`func (o *UserVerificationInfo) GetIsIdentityVerifiedOk() (*bool, bool)`

GetIsIdentityVerifiedOk returns a tuple with the IsIdentityVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIdentityVerified

`func (o *UserVerificationInfo) SetIsIdentityVerified(v bool)`

SetIsIdentityVerified sets IsIdentityVerified field to given value.


### GetReason

`func (o *UserVerificationInfo) GetReason() UserVerificationInfoReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UserVerificationInfo) GetReasonOk() (*UserVerificationInfoReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UserVerificationInfo) SetReason(v UserVerificationInfoReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UserVerificationInfo) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


