# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoInitTokens** | Pointer to **map[string]interface{}** |  | [optional] 
**CommunitiesActions** | [**CommunitiesActions**](CommunitiesActions.md) |  | 
**Country** | **string** |  | 
**GuestId** | **string** |  | 
**HasCommunityMemberships** | **bool** |  | 
**IsActiveCreator** | **bool** |  | 
**IsRestrictedSession** | **bool** |  | 
**IsSuperFollowSubscriber** | **bool** |  | 
**Language** | **string** |  | 
**OneFactorLoginEligibility** | [**OneFactorLoginEligibility**](OneFactorLoginEligibility.md) |  | 
**SuperFollowersCount** | **int32** |  | 
**SuperFollowsApplicationStatus** | **string** |  | 
**UserFeatures** | [**UserFeatures**](UserFeatures.md) |  | 
**UserId** | **string** |  | 

## Methods

### NewSession

`func NewSession(communitiesActions CommunitiesActions, country string, guestId string, hasCommunityMemberships bool, isActiveCreator bool, isRestrictedSession bool, isSuperFollowSubscriber bool, language string, oneFactorLoginEligibility OneFactorLoginEligibility, superFollowersCount int32, superFollowsApplicationStatus string, userFeatures UserFeatures, userId string, ) *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoInitTokens

`func (o *Session) GetSsoInitTokens() map[string]interface{}`

GetSsoInitTokens returns the SsoInitTokens field if non-nil, zero value otherwise.

### GetSsoInitTokensOk

`func (o *Session) GetSsoInitTokensOk() (*map[string]interface{}, bool)`

GetSsoInitTokensOk returns a tuple with the SsoInitTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoInitTokens

`func (o *Session) SetSsoInitTokens(v map[string]interface{})`

SetSsoInitTokens sets SsoInitTokens field to given value.

### HasSsoInitTokens

`func (o *Session) HasSsoInitTokens() bool`

HasSsoInitTokens returns a boolean if a field has been set.

### GetCommunitiesActions

`func (o *Session) GetCommunitiesActions() CommunitiesActions`

GetCommunitiesActions returns the CommunitiesActions field if non-nil, zero value otherwise.

### GetCommunitiesActionsOk

`func (o *Session) GetCommunitiesActionsOk() (*CommunitiesActions, bool)`

GetCommunitiesActionsOk returns a tuple with the CommunitiesActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunitiesActions

`func (o *Session) SetCommunitiesActions(v CommunitiesActions)`

SetCommunitiesActions sets CommunitiesActions field to given value.


### GetCountry

`func (o *Session) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Session) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Session) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetGuestId

`func (o *Session) GetGuestId() string`

GetGuestId returns the GuestId field if non-nil, zero value otherwise.

### GetGuestIdOk

`func (o *Session) GetGuestIdOk() (*string, bool)`

GetGuestIdOk returns a tuple with the GuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestId

`func (o *Session) SetGuestId(v string)`

SetGuestId sets GuestId field to given value.


### GetHasCommunityMemberships

`func (o *Session) GetHasCommunityMemberships() bool`

GetHasCommunityMemberships returns the HasCommunityMemberships field if non-nil, zero value otherwise.

### GetHasCommunityMembershipsOk

`func (o *Session) GetHasCommunityMembershipsOk() (*bool, bool)`

GetHasCommunityMembershipsOk returns a tuple with the HasCommunityMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCommunityMemberships

`func (o *Session) SetHasCommunityMemberships(v bool)`

SetHasCommunityMemberships sets HasCommunityMemberships field to given value.


### GetIsActiveCreator

`func (o *Session) GetIsActiveCreator() bool`

GetIsActiveCreator returns the IsActiveCreator field if non-nil, zero value otherwise.

### GetIsActiveCreatorOk

`func (o *Session) GetIsActiveCreatorOk() (*bool, bool)`

GetIsActiveCreatorOk returns a tuple with the IsActiveCreator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActiveCreator

`func (o *Session) SetIsActiveCreator(v bool)`

SetIsActiveCreator sets IsActiveCreator field to given value.


### GetIsRestrictedSession

`func (o *Session) GetIsRestrictedSession() bool`

GetIsRestrictedSession returns the IsRestrictedSession field if non-nil, zero value otherwise.

### GetIsRestrictedSessionOk

`func (o *Session) GetIsRestrictedSessionOk() (*bool, bool)`

GetIsRestrictedSessionOk returns a tuple with the IsRestrictedSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestrictedSession

`func (o *Session) SetIsRestrictedSession(v bool)`

SetIsRestrictedSession sets IsRestrictedSession field to given value.


### GetIsSuperFollowSubscriber

`func (o *Session) GetIsSuperFollowSubscriber() bool`

GetIsSuperFollowSubscriber returns the IsSuperFollowSubscriber field if non-nil, zero value otherwise.

### GetIsSuperFollowSubscriberOk

`func (o *Session) GetIsSuperFollowSubscriberOk() (*bool, bool)`

GetIsSuperFollowSubscriberOk returns a tuple with the IsSuperFollowSubscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperFollowSubscriber

`func (o *Session) SetIsSuperFollowSubscriber(v bool)`

SetIsSuperFollowSubscriber sets IsSuperFollowSubscriber field to given value.


### GetLanguage

`func (o *Session) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Session) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Session) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetOneFactorLoginEligibility

`func (o *Session) GetOneFactorLoginEligibility() OneFactorLoginEligibility`

GetOneFactorLoginEligibility returns the OneFactorLoginEligibility field if non-nil, zero value otherwise.

### GetOneFactorLoginEligibilityOk

`func (o *Session) GetOneFactorLoginEligibilityOk() (*OneFactorLoginEligibility, bool)`

GetOneFactorLoginEligibilityOk returns a tuple with the OneFactorLoginEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneFactorLoginEligibility

`func (o *Session) SetOneFactorLoginEligibility(v OneFactorLoginEligibility)`

SetOneFactorLoginEligibility sets OneFactorLoginEligibility field to given value.


### GetSuperFollowersCount

`func (o *Session) GetSuperFollowersCount() int32`

GetSuperFollowersCount returns the SuperFollowersCount field if non-nil, zero value otherwise.

### GetSuperFollowersCountOk

`func (o *Session) GetSuperFollowersCountOk() (*int32, bool)`

GetSuperFollowersCountOk returns a tuple with the SuperFollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperFollowersCount

`func (o *Session) SetSuperFollowersCount(v int32)`

SetSuperFollowersCount sets SuperFollowersCount field to given value.


### GetSuperFollowsApplicationStatus

`func (o *Session) GetSuperFollowsApplicationStatus() string`

GetSuperFollowsApplicationStatus returns the SuperFollowsApplicationStatus field if non-nil, zero value otherwise.

### GetSuperFollowsApplicationStatusOk

`func (o *Session) GetSuperFollowsApplicationStatusOk() (*string, bool)`

GetSuperFollowsApplicationStatusOk returns a tuple with the SuperFollowsApplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperFollowsApplicationStatus

`func (o *Session) SetSuperFollowsApplicationStatus(v string)`

SetSuperFollowsApplicationStatus sets SuperFollowsApplicationStatus field to given value.


### GetUserFeatures

`func (o *Session) GetUserFeatures() UserFeatures`

GetUserFeatures returns the UserFeatures field if non-nil, zero value otherwise.

### GetUserFeaturesOk

`func (o *Session) GetUserFeaturesOk() (*UserFeatures, bool)`

GetUserFeaturesOk returns a tuple with the UserFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFeatures

`func (o *Session) SetUserFeatures(v UserFeatures)`

SetUserFeatures sets UserFeatures field to given value.


### GetUserId

`func (o *Session) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Session) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Session) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


