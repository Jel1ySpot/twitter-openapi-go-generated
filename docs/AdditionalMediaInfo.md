# AdditionalMediaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallToActions** | Pointer to [**AdditionalMediaInfoCallToActions**](AdditionalMediaInfoCallToActions.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Embeddable** | Pointer to **bool** |  | [optional] 
**Monetizable** | **bool** |  | 
**SourceUser** | Pointer to [**UserResultCore**](UserResultCore.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAdditionalMediaInfo

`func NewAdditionalMediaInfo(monetizable bool, ) *AdditionalMediaInfo`

NewAdditionalMediaInfo instantiates a new AdditionalMediaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalMediaInfoWithDefaults

`func NewAdditionalMediaInfoWithDefaults() *AdditionalMediaInfo`

NewAdditionalMediaInfoWithDefaults instantiates a new AdditionalMediaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallToActions

`func (o *AdditionalMediaInfo) GetCallToActions() AdditionalMediaInfoCallToActions`

GetCallToActions returns the CallToActions field if non-nil, zero value otherwise.

### GetCallToActionsOk

`func (o *AdditionalMediaInfo) GetCallToActionsOk() (*AdditionalMediaInfoCallToActions, bool)`

GetCallToActionsOk returns a tuple with the CallToActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallToActions

`func (o *AdditionalMediaInfo) SetCallToActions(v AdditionalMediaInfoCallToActions)`

SetCallToActions sets CallToActions field to given value.

### HasCallToActions

`func (o *AdditionalMediaInfo) HasCallToActions() bool`

HasCallToActions returns a boolean if a field has been set.

### GetDescription

`func (o *AdditionalMediaInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdditionalMediaInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdditionalMediaInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdditionalMediaInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmbeddable

`func (o *AdditionalMediaInfo) GetEmbeddable() bool`

GetEmbeddable returns the Embeddable field if non-nil, zero value otherwise.

### GetEmbeddableOk

`func (o *AdditionalMediaInfo) GetEmbeddableOk() (*bool, bool)`

GetEmbeddableOk returns a tuple with the Embeddable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddable

`func (o *AdditionalMediaInfo) SetEmbeddable(v bool)`

SetEmbeddable sets Embeddable field to given value.

### HasEmbeddable

`func (o *AdditionalMediaInfo) HasEmbeddable() bool`

HasEmbeddable returns a boolean if a field has been set.

### GetMonetizable

`func (o *AdditionalMediaInfo) GetMonetizable() bool`

GetMonetizable returns the Monetizable field if non-nil, zero value otherwise.

### GetMonetizableOk

`func (o *AdditionalMediaInfo) GetMonetizableOk() (*bool, bool)`

GetMonetizableOk returns a tuple with the Monetizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonetizable

`func (o *AdditionalMediaInfo) SetMonetizable(v bool)`

SetMonetizable sets Monetizable field to given value.


### GetSourceUser

`func (o *AdditionalMediaInfo) GetSourceUser() UserResultCore`

GetSourceUser returns the SourceUser field if non-nil, zero value otherwise.

### GetSourceUserOk

`func (o *AdditionalMediaInfo) GetSourceUserOk() (*UserResultCore, bool)`

GetSourceUserOk returns a tuple with the SourceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUser

`func (o *AdditionalMediaInfo) SetSourceUser(v UserResultCore)`

SetSourceUser sets SourceUser field to given value.

### HasSourceUser

`func (o *AdditionalMediaInfo) HasSourceUser() bool`

HasSourceUser returns a boolean if a field has been set.

### GetTitle

`func (o *AdditionalMediaInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdditionalMediaInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdditionalMediaInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdditionalMediaInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


