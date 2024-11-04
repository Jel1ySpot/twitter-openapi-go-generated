# CommunityRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**RestId** | **string** |  | 

## Methods

### NewCommunityRule

`func NewCommunityRule(name string, restId string, ) *CommunityRule`

NewCommunityRule instantiates a new CommunityRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunityRuleWithDefaults

`func NewCommunityRuleWithDefaults() *CommunityRule`

NewCommunityRuleWithDefaults instantiates a new CommunityRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CommunityRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommunityRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommunityRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommunityRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CommunityRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommunityRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommunityRule) SetName(v string)`

SetName sets Name field to given value.


### GetRestId

`func (o *CommunityRule) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *CommunityRule) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *CommunityRule) SetRestId(v string)`

SetRestId sets RestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


