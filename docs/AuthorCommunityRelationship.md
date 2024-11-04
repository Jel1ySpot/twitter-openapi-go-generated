# AuthorCommunityRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunityResults** | [**Community**](Community.md) |  | 
**Role** | Pointer to **string** |  | [optional] 
**UserResults** | Pointer to [**UserResults**](UserResults.md) |  | [optional] 

## Methods

### NewAuthorCommunityRelationship

`func NewAuthorCommunityRelationship(communityResults Community, ) *AuthorCommunityRelationship`

NewAuthorCommunityRelationship instantiates a new AuthorCommunityRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorCommunityRelationshipWithDefaults

`func NewAuthorCommunityRelationshipWithDefaults() *AuthorCommunityRelationship`

NewAuthorCommunityRelationshipWithDefaults instantiates a new AuthorCommunityRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunityResults

`func (o *AuthorCommunityRelationship) GetCommunityResults() Community`

GetCommunityResults returns the CommunityResults field if non-nil, zero value otherwise.

### GetCommunityResultsOk

`func (o *AuthorCommunityRelationship) GetCommunityResultsOk() (*Community, bool)`

GetCommunityResultsOk returns a tuple with the CommunityResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityResults

`func (o *AuthorCommunityRelationship) SetCommunityResults(v Community)`

SetCommunityResults sets CommunityResults field to given value.


### GetRole

`func (o *AuthorCommunityRelationship) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AuthorCommunityRelationship) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AuthorCommunityRelationship) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AuthorCommunityRelationship) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUserResults

`func (o *AuthorCommunityRelationship) GetUserResults() UserResults`

GetUserResults returns the UserResults field if non-nil, zero value otherwise.

### GetUserResultsOk

`func (o *AuthorCommunityRelationship) GetUserResultsOk() (*UserResults, bool)`

GetUserResultsOk returns a tuple with the UserResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResults

`func (o *AuthorCommunityRelationship) SetUserResults(v UserResults)`

SetUserResults sets UserResults field to given value.

### HasUserResults

`func (o *AuthorCommunityRelationship) HasUserResults() bool`

HasUserResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


