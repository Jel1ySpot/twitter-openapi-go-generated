# CommunityRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**CommunityActions**](CommunityActions.md) |  | 
**Id** | **string** |  | 
**ModerationState** | **map[string]interface{}** |  | 
**RestId** | **string** |  | 

## Methods

### NewCommunityRelationship

`func NewCommunityRelationship(actions CommunityActions, id string, moderationState map[string]interface{}, restId string, ) *CommunityRelationship`

NewCommunityRelationship instantiates a new CommunityRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunityRelationshipWithDefaults

`func NewCommunityRelationshipWithDefaults() *CommunityRelationship`

NewCommunityRelationshipWithDefaults instantiates a new CommunityRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *CommunityRelationship) GetActions() CommunityActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CommunityRelationship) GetActionsOk() (*CommunityActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CommunityRelationship) SetActions(v CommunityActions)`

SetActions sets Actions field to given value.


### GetId

`func (o *CommunityRelationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommunityRelationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommunityRelationship) SetId(v string)`

SetId sets Id field to given value.


### GetModerationState

`func (o *CommunityRelationship) GetModerationState() map[string]interface{}`

GetModerationState returns the ModerationState field if non-nil, zero value otherwise.

### GetModerationStateOk

`func (o *CommunityRelationship) GetModerationStateOk() (*map[string]interface{}, bool)`

GetModerationStateOk returns a tuple with the ModerationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationState

`func (o *CommunityRelationship) SetModerationState(v map[string]interface{})`

SetModerationState sets ModerationState field to given value.


### GetRestId

`func (o *CommunityRelationship) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *CommunityRelationship) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *CommunityRelationship) SetRestId(v string)`

SetRestId sets RestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


