# TimelineUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**ItemType** | [**ContentItemType**](ContentItemType.md) |  | 
**SocialContext** | Pointer to [**SocialContextUnion**](SocialContextUnion.md) |  | [optional] 
**UserDisplayType** | **string** |  | 
**UserResults** | [**UserResults**](UserResults.md) |  | 

## Methods

### NewTimelineUser

`func NewTimelineUser(typename TypeName, itemType ContentItemType, userDisplayType string, userResults UserResults, ) *TimelineUser`

NewTimelineUser instantiates a new TimelineUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineUserWithDefaults

`func NewTimelineUserWithDefaults() *TimelineUser`

NewTimelineUserWithDefaults instantiates a new TimelineUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *TimelineUser) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *TimelineUser) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *TimelineUser) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetItemType

`func (o *TimelineUser) GetItemType() ContentItemType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *TimelineUser) GetItemTypeOk() (*ContentItemType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *TimelineUser) SetItemType(v ContentItemType)`

SetItemType sets ItemType field to given value.


### GetSocialContext

`func (o *TimelineUser) GetSocialContext() SocialContextUnion`

GetSocialContext returns the SocialContext field if non-nil, zero value otherwise.

### GetSocialContextOk

`func (o *TimelineUser) GetSocialContextOk() (*SocialContextUnion, bool)`

GetSocialContextOk returns a tuple with the SocialContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialContext

`func (o *TimelineUser) SetSocialContext(v SocialContextUnion)`

SetSocialContext sets SocialContext field to given value.

### HasSocialContext

`func (o *TimelineUser) HasSocialContext() bool`

HasSocialContext returns a boolean if a field has been set.

### GetUserDisplayType

`func (o *TimelineUser) GetUserDisplayType() string`

GetUserDisplayType returns the UserDisplayType field if non-nil, zero value otherwise.

### GetUserDisplayTypeOk

`func (o *TimelineUser) GetUserDisplayTypeOk() (*string, bool)`

GetUserDisplayTypeOk returns a tuple with the UserDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayType

`func (o *TimelineUser) SetUserDisplayType(v string)`

SetUserDisplayType sets UserDisplayType field to given value.


### GetUserResults

`func (o *TimelineUser) GetUserResults() UserResults`

GetUserResults returns the UserResults field if non-nil, zero value otherwise.

### GetUserResultsOk

`func (o *TimelineUser) GetUserResultsOk() (*UserResults, bool)`

GetUserResultsOk returns a tuple with the UserResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResults

`func (o *TimelineUser) SetUserResults(v UserResults)`

SetUserResults sets UserResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


