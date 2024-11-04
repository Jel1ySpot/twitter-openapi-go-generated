# UserVerificationInfoReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | [**UserVerificationInfoReasonDescription**](UserVerificationInfoReasonDescription.md) |  | 
**OverrideVerifiedYear** | **int32** |  | 
**VerifiedSinceMsec** | **string** |  | 

## Methods

### NewUserVerificationInfoReason

`func NewUserVerificationInfoReason(description UserVerificationInfoReasonDescription, overrideVerifiedYear int32, verifiedSinceMsec string, ) *UserVerificationInfoReason`

NewUserVerificationInfoReason instantiates a new UserVerificationInfoReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserVerificationInfoReasonWithDefaults

`func NewUserVerificationInfoReasonWithDefaults() *UserVerificationInfoReason`

NewUserVerificationInfoReasonWithDefaults instantiates a new UserVerificationInfoReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UserVerificationInfoReason) GetDescription() UserVerificationInfoReasonDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserVerificationInfoReason) GetDescriptionOk() (*UserVerificationInfoReasonDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserVerificationInfoReason) SetDescription(v UserVerificationInfoReasonDescription)`

SetDescription sets Description field to given value.


### GetOverrideVerifiedYear

`func (o *UserVerificationInfoReason) GetOverrideVerifiedYear() int32`

GetOverrideVerifiedYear returns the OverrideVerifiedYear field if non-nil, zero value otherwise.

### GetOverrideVerifiedYearOk

`func (o *UserVerificationInfoReason) GetOverrideVerifiedYearOk() (*int32, bool)`

GetOverrideVerifiedYearOk returns a tuple with the OverrideVerifiedYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideVerifiedYear

`func (o *UserVerificationInfoReason) SetOverrideVerifiedYear(v int32)`

SetOverrideVerifiedYear sets OverrideVerifiedYear field to given value.


### GetVerifiedSinceMsec

`func (o *UserVerificationInfoReason) GetVerifiedSinceMsec() string`

GetVerifiedSinceMsec returns the VerifiedSinceMsec field if non-nil, zero value otherwise.

### GetVerifiedSinceMsecOk

`func (o *UserVerificationInfoReason) GetVerifiedSinceMsecOk() (*string, bool)`

GetVerifiedSinceMsecOk returns a tuple with the VerifiedSinceMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedSinceMsec

`func (o *UserVerificationInfoReason) SetVerifiedSinceMsec(v string)`

SetVerifiedSinceMsec sets VerifiedSinceMsec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


