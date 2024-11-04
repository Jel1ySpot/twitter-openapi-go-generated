# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**AffiliatesHighlightedLabel** | **map[string]interface{}** |  | 
**BusinessAccount** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatorSubscriptionsCount** | Pointer to **int32** |  | [optional] 
**HasGraduatedAccess** | Pointer to **bool** |  | [optional] 
**HasHiddenLikesOnProfile** | Pointer to **bool** |  | [optional] 
**HasNftAvatar** | Pointer to **bool** |  | [optional] 
**HighlightsInfo** | Pointer to [**UserHighlightsInfo**](UserHighlightsInfo.md) |  | [optional] 
**Id** | **string** |  | 
**IsBlueVerified** | **bool** |  | 
**IsProfileTranslatable** | Pointer to **bool** |  | [optional] 
**Legacy** | [**UserLegacy**](UserLegacy.md) |  | 
**LegacyExtendedProfile** | Pointer to [**UserLegacyExtendedProfile**](UserLegacyExtendedProfile.md) |  | [optional] 
**PremiumGiftingEligible** | Pointer to **bool** |  | [optional] 
**Professional** | Pointer to [**UserProfessional**](UserProfessional.md) |  | [optional] 
**ProfileImageShape** | **string** |  | 
**RestId** | **string** |  | 
**SuperFollowEligible** | Pointer to **bool** |  | [optional] 
**SuperFollowedBy** | Pointer to **bool** |  | [optional] 
**SuperFollowing** | Pointer to **bool** |  | [optional] 
**TipjarSettings** | Pointer to [**UserTipJarSettings**](UserTipJarSettings.md) |  | [optional] 
**UserSeedTweetCount** | Pointer to **int32** |  | [optional] 
**VerificationInfo** | Pointer to [**UserVerificationInfo**](UserVerificationInfo.md) |  | [optional] 

## Methods

### NewUser

`func NewUser(typename TypeName, affiliatesHighlightedLabel map[string]interface{}, id string, isBlueVerified bool, legacy UserLegacy, profileImageShape string, restId string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *User) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *User) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *User) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetAffiliatesHighlightedLabel

`func (o *User) GetAffiliatesHighlightedLabel() map[string]interface{}`

GetAffiliatesHighlightedLabel returns the AffiliatesHighlightedLabel field if non-nil, zero value otherwise.

### GetAffiliatesHighlightedLabelOk

`func (o *User) GetAffiliatesHighlightedLabelOk() (*map[string]interface{}, bool)`

GetAffiliatesHighlightedLabelOk returns a tuple with the AffiliatesHighlightedLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliatesHighlightedLabel

`func (o *User) SetAffiliatesHighlightedLabel(v map[string]interface{})`

SetAffiliatesHighlightedLabel sets AffiliatesHighlightedLabel field to given value.


### GetBusinessAccount

`func (o *User) GetBusinessAccount() map[string]interface{}`

GetBusinessAccount returns the BusinessAccount field if non-nil, zero value otherwise.

### GetBusinessAccountOk

`func (o *User) GetBusinessAccountOk() (*map[string]interface{}, bool)`

GetBusinessAccountOk returns a tuple with the BusinessAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAccount

`func (o *User) SetBusinessAccount(v map[string]interface{})`

SetBusinessAccount sets BusinessAccount field to given value.

### HasBusinessAccount

`func (o *User) HasBusinessAccount() bool`

HasBusinessAccount returns a boolean if a field has been set.

### GetCreatorSubscriptionsCount

`func (o *User) GetCreatorSubscriptionsCount() int32`

GetCreatorSubscriptionsCount returns the CreatorSubscriptionsCount field if non-nil, zero value otherwise.

### GetCreatorSubscriptionsCountOk

`func (o *User) GetCreatorSubscriptionsCountOk() (*int32, bool)`

GetCreatorSubscriptionsCountOk returns a tuple with the CreatorSubscriptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorSubscriptionsCount

`func (o *User) SetCreatorSubscriptionsCount(v int32)`

SetCreatorSubscriptionsCount sets CreatorSubscriptionsCount field to given value.

### HasCreatorSubscriptionsCount

`func (o *User) HasCreatorSubscriptionsCount() bool`

HasCreatorSubscriptionsCount returns a boolean if a field has been set.

### GetHasGraduatedAccess

`func (o *User) GetHasGraduatedAccess() bool`

GetHasGraduatedAccess returns the HasGraduatedAccess field if non-nil, zero value otherwise.

### GetHasGraduatedAccessOk

`func (o *User) GetHasGraduatedAccessOk() (*bool, bool)`

GetHasGraduatedAccessOk returns a tuple with the HasGraduatedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGraduatedAccess

`func (o *User) SetHasGraduatedAccess(v bool)`

SetHasGraduatedAccess sets HasGraduatedAccess field to given value.

### HasHasGraduatedAccess

`func (o *User) HasHasGraduatedAccess() bool`

HasHasGraduatedAccess returns a boolean if a field has been set.

### GetHasHiddenLikesOnProfile

`func (o *User) GetHasHiddenLikesOnProfile() bool`

GetHasHiddenLikesOnProfile returns the HasHiddenLikesOnProfile field if non-nil, zero value otherwise.

### GetHasHiddenLikesOnProfileOk

`func (o *User) GetHasHiddenLikesOnProfileOk() (*bool, bool)`

GetHasHiddenLikesOnProfileOk returns a tuple with the HasHiddenLikesOnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHiddenLikesOnProfile

`func (o *User) SetHasHiddenLikesOnProfile(v bool)`

SetHasHiddenLikesOnProfile sets HasHiddenLikesOnProfile field to given value.

### HasHasHiddenLikesOnProfile

`func (o *User) HasHasHiddenLikesOnProfile() bool`

HasHasHiddenLikesOnProfile returns a boolean if a field has been set.

### GetHasNftAvatar

`func (o *User) GetHasNftAvatar() bool`

GetHasNftAvatar returns the HasNftAvatar field if non-nil, zero value otherwise.

### GetHasNftAvatarOk

`func (o *User) GetHasNftAvatarOk() (*bool, bool)`

GetHasNftAvatarOk returns a tuple with the HasNftAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNftAvatar

`func (o *User) SetHasNftAvatar(v bool)`

SetHasNftAvatar sets HasNftAvatar field to given value.

### HasHasNftAvatar

`func (o *User) HasHasNftAvatar() bool`

HasHasNftAvatar returns a boolean if a field has been set.

### GetHighlightsInfo

`func (o *User) GetHighlightsInfo() UserHighlightsInfo`

GetHighlightsInfo returns the HighlightsInfo field if non-nil, zero value otherwise.

### GetHighlightsInfoOk

`func (o *User) GetHighlightsInfoOk() (*UserHighlightsInfo, bool)`

GetHighlightsInfoOk returns a tuple with the HighlightsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightsInfo

`func (o *User) SetHighlightsInfo(v UserHighlightsInfo)`

SetHighlightsInfo sets HighlightsInfo field to given value.

### HasHighlightsInfo

`func (o *User) HasHighlightsInfo() bool`

HasHighlightsInfo returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetIsBlueVerified

`func (o *User) GetIsBlueVerified() bool`

GetIsBlueVerified returns the IsBlueVerified field if non-nil, zero value otherwise.

### GetIsBlueVerifiedOk

`func (o *User) GetIsBlueVerifiedOk() (*bool, bool)`

GetIsBlueVerifiedOk returns a tuple with the IsBlueVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlueVerified

`func (o *User) SetIsBlueVerified(v bool)`

SetIsBlueVerified sets IsBlueVerified field to given value.


### GetIsProfileTranslatable

`func (o *User) GetIsProfileTranslatable() bool`

GetIsProfileTranslatable returns the IsProfileTranslatable field if non-nil, zero value otherwise.

### GetIsProfileTranslatableOk

`func (o *User) GetIsProfileTranslatableOk() (*bool, bool)`

GetIsProfileTranslatableOk returns a tuple with the IsProfileTranslatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfileTranslatable

`func (o *User) SetIsProfileTranslatable(v bool)`

SetIsProfileTranslatable sets IsProfileTranslatable field to given value.

### HasIsProfileTranslatable

`func (o *User) HasIsProfileTranslatable() bool`

HasIsProfileTranslatable returns a boolean if a field has been set.

### GetLegacy

`func (o *User) GetLegacy() UserLegacy`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *User) GetLegacyOk() (*UserLegacy, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *User) SetLegacy(v UserLegacy)`

SetLegacy sets Legacy field to given value.


### GetLegacyExtendedProfile

`func (o *User) GetLegacyExtendedProfile() UserLegacyExtendedProfile`

GetLegacyExtendedProfile returns the LegacyExtendedProfile field if non-nil, zero value otherwise.

### GetLegacyExtendedProfileOk

`func (o *User) GetLegacyExtendedProfileOk() (*UserLegacyExtendedProfile, bool)`

GetLegacyExtendedProfileOk returns a tuple with the LegacyExtendedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyExtendedProfile

`func (o *User) SetLegacyExtendedProfile(v UserLegacyExtendedProfile)`

SetLegacyExtendedProfile sets LegacyExtendedProfile field to given value.

### HasLegacyExtendedProfile

`func (o *User) HasLegacyExtendedProfile() bool`

HasLegacyExtendedProfile returns a boolean if a field has been set.

### GetPremiumGiftingEligible

`func (o *User) GetPremiumGiftingEligible() bool`

GetPremiumGiftingEligible returns the PremiumGiftingEligible field if non-nil, zero value otherwise.

### GetPremiumGiftingEligibleOk

`func (o *User) GetPremiumGiftingEligibleOk() (*bool, bool)`

GetPremiumGiftingEligibleOk returns a tuple with the PremiumGiftingEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumGiftingEligible

`func (o *User) SetPremiumGiftingEligible(v bool)`

SetPremiumGiftingEligible sets PremiumGiftingEligible field to given value.

### HasPremiumGiftingEligible

`func (o *User) HasPremiumGiftingEligible() bool`

HasPremiumGiftingEligible returns a boolean if a field has been set.

### GetProfessional

`func (o *User) GetProfessional() UserProfessional`

GetProfessional returns the Professional field if non-nil, zero value otherwise.

### GetProfessionalOk

`func (o *User) GetProfessionalOk() (*UserProfessional, bool)`

GetProfessionalOk returns a tuple with the Professional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessional

`func (o *User) SetProfessional(v UserProfessional)`

SetProfessional sets Professional field to given value.

### HasProfessional

`func (o *User) HasProfessional() bool`

HasProfessional returns a boolean if a field has been set.

### GetProfileImageShape

`func (o *User) GetProfileImageShape() string`

GetProfileImageShape returns the ProfileImageShape field if non-nil, zero value otherwise.

### GetProfileImageShapeOk

`func (o *User) GetProfileImageShapeOk() (*string, bool)`

GetProfileImageShapeOk returns a tuple with the ProfileImageShape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageShape

`func (o *User) SetProfileImageShape(v string)`

SetProfileImageShape sets ProfileImageShape field to given value.


### GetRestId

`func (o *User) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *User) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *User) SetRestId(v string)`

SetRestId sets RestId field to given value.


### GetSuperFollowEligible

`func (o *User) GetSuperFollowEligible() bool`

GetSuperFollowEligible returns the SuperFollowEligible field if non-nil, zero value otherwise.

### GetSuperFollowEligibleOk

`func (o *User) GetSuperFollowEligibleOk() (*bool, bool)`

GetSuperFollowEligibleOk returns a tuple with the SuperFollowEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperFollowEligible

`func (o *User) SetSuperFollowEligible(v bool)`

SetSuperFollowEligible sets SuperFollowEligible field to given value.

### HasSuperFollowEligible

`func (o *User) HasSuperFollowEligible() bool`

HasSuperFollowEligible returns a boolean if a field has been set.

### GetSuperFollowedBy

`func (o *User) GetSuperFollowedBy() bool`

GetSuperFollowedBy returns the SuperFollowedBy field if non-nil, zero value otherwise.

### GetSuperFollowedByOk

`func (o *User) GetSuperFollowedByOk() (*bool, bool)`

GetSuperFollowedByOk returns a tuple with the SuperFollowedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperFollowedBy

`func (o *User) SetSuperFollowedBy(v bool)`

SetSuperFollowedBy sets SuperFollowedBy field to given value.

### HasSuperFollowedBy

`func (o *User) HasSuperFollowedBy() bool`

HasSuperFollowedBy returns a boolean if a field has been set.

### GetSuperFollowing

`func (o *User) GetSuperFollowing() bool`

GetSuperFollowing returns the SuperFollowing field if non-nil, zero value otherwise.

### GetSuperFollowingOk

`func (o *User) GetSuperFollowingOk() (*bool, bool)`

GetSuperFollowingOk returns a tuple with the SuperFollowing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperFollowing

`func (o *User) SetSuperFollowing(v bool)`

SetSuperFollowing sets SuperFollowing field to given value.

### HasSuperFollowing

`func (o *User) HasSuperFollowing() bool`

HasSuperFollowing returns a boolean if a field has been set.

### GetTipjarSettings

`func (o *User) GetTipjarSettings() UserTipJarSettings`

GetTipjarSettings returns the TipjarSettings field if non-nil, zero value otherwise.

### GetTipjarSettingsOk

`func (o *User) GetTipjarSettingsOk() (*UserTipJarSettings, bool)`

GetTipjarSettingsOk returns a tuple with the TipjarSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipjarSettings

`func (o *User) SetTipjarSettings(v UserTipJarSettings)`

SetTipjarSettings sets TipjarSettings field to given value.

### HasTipjarSettings

`func (o *User) HasTipjarSettings() bool`

HasTipjarSettings returns a boolean if a field has been set.

### GetUserSeedTweetCount

`func (o *User) GetUserSeedTweetCount() int32`

GetUserSeedTweetCount returns the UserSeedTweetCount field if non-nil, zero value otherwise.

### GetUserSeedTweetCountOk

`func (o *User) GetUserSeedTweetCountOk() (*int32, bool)`

GetUserSeedTweetCountOk returns a tuple with the UserSeedTweetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSeedTweetCount

`func (o *User) SetUserSeedTweetCount(v int32)`

SetUserSeedTweetCount sets UserSeedTweetCount field to given value.

### HasUserSeedTweetCount

`func (o *User) HasUserSeedTweetCount() bool`

HasUserSeedTweetCount returns a boolean if a field has been set.

### GetVerificationInfo

`func (o *User) GetVerificationInfo() UserVerificationInfo`

GetVerificationInfo returns the VerificationInfo field if non-nil, zero value otherwise.

### GetVerificationInfoOk

`func (o *User) GetVerificationInfoOk() (*UserVerificationInfo, bool)`

GetVerificationInfoOk returns a tuple with the VerificationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationInfo

`func (o *User) SetVerificationInfo(v UserVerificationInfo)`

SetVerificationInfo sets VerificationInfo field to given value.

### HasVerificationInfo

`func (o *User) HasVerificationInfo() bool`

HasVerificationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


