# UserUnion

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
**Message** | Pointer to **string** |  | [optional] 
**Reason** | **string** |  | 

## Methods

### NewUserUnion

`func NewUserUnion(typename TypeName, affiliatesHighlightedLabel map[string]interface{}, id string, isBlueVerified bool, legacy UserLegacy, profileImageShape string, restId string, reason string, ) *UserUnion`

NewUserUnion instantiates a new UserUnion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUnionWithDefaults

`func NewUserUnionWithDefaults() *UserUnion`

NewUserUnionWithDefaults instantiates a new UserUnion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *UserUnion) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *UserUnion) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *UserUnion) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetAffiliatesHighlightedLabel

`func (o *UserUnion) GetAffiliatesHighlightedLabel() map[string]interface{}`

GetAffiliatesHighlightedLabel returns the AffiliatesHighlightedLabel field if non-nil, zero value otherwise.

### GetAffiliatesHighlightedLabelOk

`func (o *UserUnion) GetAffiliatesHighlightedLabelOk() (*map[string]interface{}, bool)`

GetAffiliatesHighlightedLabelOk returns a tuple with the AffiliatesHighlightedLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliatesHighlightedLabel

`func (o *UserUnion) SetAffiliatesHighlightedLabel(v map[string]interface{})`

SetAffiliatesHighlightedLabel sets AffiliatesHighlightedLabel field to given value.


### GetBusinessAccount

`func (o *UserUnion) GetBusinessAccount() map[string]interface{}`

GetBusinessAccount returns the BusinessAccount field if non-nil, zero value otherwise.

### GetBusinessAccountOk

`func (o *UserUnion) GetBusinessAccountOk() (*map[string]interface{}, bool)`

GetBusinessAccountOk returns a tuple with the BusinessAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAccount

`func (o *UserUnion) SetBusinessAccount(v map[string]interface{})`

SetBusinessAccount sets BusinessAccount field to given value.

### HasBusinessAccount

`func (o *UserUnion) HasBusinessAccount() bool`

HasBusinessAccount returns a boolean if a field has been set.

### GetCreatorSubscriptionsCount

`func (o *UserUnion) GetCreatorSubscriptionsCount() int32`

GetCreatorSubscriptionsCount returns the CreatorSubscriptionsCount field if non-nil, zero value otherwise.

### GetCreatorSubscriptionsCountOk

`func (o *UserUnion) GetCreatorSubscriptionsCountOk() (*int32, bool)`

GetCreatorSubscriptionsCountOk returns a tuple with the CreatorSubscriptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorSubscriptionsCount

`func (o *UserUnion) SetCreatorSubscriptionsCount(v int32)`

SetCreatorSubscriptionsCount sets CreatorSubscriptionsCount field to given value.

### HasCreatorSubscriptionsCount

`func (o *UserUnion) HasCreatorSubscriptionsCount() bool`

HasCreatorSubscriptionsCount returns a boolean if a field has been set.

### GetHasGraduatedAccess

`func (o *UserUnion) GetHasGraduatedAccess() bool`

GetHasGraduatedAccess returns the HasGraduatedAccess field if non-nil, zero value otherwise.

### GetHasGraduatedAccessOk

`func (o *UserUnion) GetHasGraduatedAccessOk() (*bool, bool)`

GetHasGraduatedAccessOk returns a tuple with the HasGraduatedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGraduatedAccess

`func (o *UserUnion) SetHasGraduatedAccess(v bool)`

SetHasGraduatedAccess sets HasGraduatedAccess field to given value.

### HasHasGraduatedAccess

`func (o *UserUnion) HasHasGraduatedAccess() bool`

HasHasGraduatedAccess returns a boolean if a field has been set.

### GetHasHiddenLikesOnProfile

`func (o *UserUnion) GetHasHiddenLikesOnProfile() bool`

GetHasHiddenLikesOnProfile returns the HasHiddenLikesOnProfile field if non-nil, zero value otherwise.

### GetHasHiddenLikesOnProfileOk

`func (o *UserUnion) GetHasHiddenLikesOnProfileOk() (*bool, bool)`

GetHasHiddenLikesOnProfileOk returns a tuple with the HasHiddenLikesOnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHiddenLikesOnProfile

`func (o *UserUnion) SetHasHiddenLikesOnProfile(v bool)`

SetHasHiddenLikesOnProfile sets HasHiddenLikesOnProfile field to given value.

### HasHasHiddenLikesOnProfile

`func (o *UserUnion) HasHasHiddenLikesOnProfile() bool`

HasHasHiddenLikesOnProfile returns a boolean if a field has been set.

### GetHasNftAvatar

`func (o *UserUnion) GetHasNftAvatar() bool`

GetHasNftAvatar returns the HasNftAvatar field if non-nil, zero value otherwise.

### GetHasNftAvatarOk

`func (o *UserUnion) GetHasNftAvatarOk() (*bool, bool)`

GetHasNftAvatarOk returns a tuple with the HasNftAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNftAvatar

`func (o *UserUnion) SetHasNftAvatar(v bool)`

SetHasNftAvatar sets HasNftAvatar field to given value.

### HasHasNftAvatar

`func (o *UserUnion) HasHasNftAvatar() bool`

HasHasNftAvatar returns a boolean if a field has been set.

### GetHighlightsInfo

`func (o *UserUnion) GetHighlightsInfo() UserHighlightsInfo`

GetHighlightsInfo returns the HighlightsInfo field if non-nil, zero value otherwise.

### GetHighlightsInfoOk

`func (o *UserUnion) GetHighlightsInfoOk() (*UserHighlightsInfo, bool)`

GetHighlightsInfoOk returns a tuple with the HighlightsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightsInfo

`func (o *UserUnion) SetHighlightsInfo(v UserHighlightsInfo)`

SetHighlightsInfo sets HighlightsInfo field to given value.

### HasHighlightsInfo

`func (o *UserUnion) HasHighlightsInfo() bool`

HasHighlightsInfo returns a boolean if a field has been set.

### GetId

`func (o *UserUnion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserUnion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserUnion) SetId(v string)`

SetId sets Id field to given value.


### GetIsBlueVerified

`func (o *UserUnion) GetIsBlueVerified() bool`

GetIsBlueVerified returns the IsBlueVerified field if non-nil, zero value otherwise.

### GetIsBlueVerifiedOk

`func (o *UserUnion) GetIsBlueVerifiedOk() (*bool, bool)`

GetIsBlueVerifiedOk returns a tuple with the IsBlueVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlueVerified

`func (o *UserUnion) SetIsBlueVerified(v bool)`

SetIsBlueVerified sets IsBlueVerified field to given value.


### GetIsProfileTranslatable

`func (o *UserUnion) GetIsProfileTranslatable() bool`

GetIsProfileTranslatable returns the IsProfileTranslatable field if non-nil, zero value otherwise.

### GetIsProfileTranslatableOk

`func (o *UserUnion) GetIsProfileTranslatableOk() (*bool, bool)`

GetIsProfileTranslatableOk returns a tuple with the IsProfileTranslatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfileTranslatable

`func (o *UserUnion) SetIsProfileTranslatable(v bool)`

SetIsProfileTranslatable sets IsProfileTranslatable field to given value.

### HasIsProfileTranslatable

`func (o *UserUnion) HasIsProfileTranslatable() bool`

HasIsProfileTranslatable returns a boolean if a field has been set.

### GetLegacy

`func (o *UserUnion) GetLegacy() UserLegacy`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *UserUnion) GetLegacyOk() (*UserLegacy, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *UserUnion) SetLegacy(v UserLegacy)`

SetLegacy sets Legacy field to given value.


### GetLegacyExtendedProfile

`func (o *UserUnion) GetLegacyExtendedProfile() UserLegacyExtendedProfile`

GetLegacyExtendedProfile returns the LegacyExtendedProfile field if non-nil, zero value otherwise.

### GetLegacyExtendedProfileOk

`func (o *UserUnion) GetLegacyExtendedProfileOk() (*UserLegacyExtendedProfile, bool)`

GetLegacyExtendedProfileOk returns a tuple with the LegacyExtendedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyExtendedProfile

`func (o *UserUnion) SetLegacyExtendedProfile(v UserLegacyExtendedProfile)`

SetLegacyExtendedProfile sets LegacyExtendedProfile field to given value.

### HasLegacyExtendedProfile

`func (o *UserUnion) HasLegacyExtendedProfile() bool`

HasLegacyExtendedProfile returns a boolean if a field has been set.

### GetPremiumGiftingEligible

`func (o *UserUnion) GetPremiumGiftingEligible() bool`

GetPremiumGiftingEligible returns the PremiumGiftingEligible field if non-nil, zero value otherwise.

### GetPremiumGiftingEligibleOk

`func (o *UserUnion) GetPremiumGiftingEligibleOk() (*bool, bool)`

GetPremiumGiftingEligibleOk returns a tuple with the PremiumGiftingEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumGiftingEligible

`func (o *UserUnion) SetPremiumGiftingEligible(v bool)`

SetPremiumGiftingEligible sets PremiumGiftingEligible field to given value.

### HasPremiumGiftingEligible

`func (o *UserUnion) HasPremiumGiftingEligible() bool`

HasPremiumGiftingEligible returns a boolean if a field has been set.

### GetProfessional

`func (o *UserUnion) GetProfessional() UserProfessional`

GetProfessional returns the Professional field if non-nil, zero value otherwise.

### GetProfessionalOk

`func (o *UserUnion) GetProfessionalOk() (*UserProfessional, bool)`

GetProfessionalOk returns a tuple with the Professional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessional

`func (o *UserUnion) SetProfessional(v UserProfessional)`

SetProfessional sets Professional field to given value.

### HasProfessional

`func (o *UserUnion) HasProfessional() bool`

HasProfessional returns a boolean if a field has been set.

### GetProfileImageShape

`func (o *UserUnion) GetProfileImageShape() string`

GetProfileImageShape returns the ProfileImageShape field if non-nil, zero value otherwise.

### GetProfileImageShapeOk

`func (o *UserUnion) GetProfileImageShapeOk() (*string, bool)`

GetProfileImageShapeOk returns a tuple with the ProfileImageShape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageShape

`func (o *UserUnion) SetProfileImageShape(v string)`

SetProfileImageShape sets ProfileImageShape field to given value.


### GetRestId

`func (o *UserUnion) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *UserUnion) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *UserUnion) SetRestId(v string)`

SetRestId sets RestId field to given value.


### GetSuperFollowEligible

`func (o *UserUnion) GetSuperFollowEligible() bool`

GetSuperFollowEligible returns the SuperFollowEligible field if non-nil, zero value otherwise.

### GetSuperFollowEligibleOk

`func (o *UserUnion) GetSuperFollowEligibleOk() (*bool, bool)`

GetSuperFollowEligibleOk returns a tuple with the SuperFollowEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperFollowEligible

`func (o *UserUnion) SetSuperFollowEligible(v bool)`

SetSuperFollowEligible sets SuperFollowEligible field to given value.

### HasSuperFollowEligible

`func (o *UserUnion) HasSuperFollowEligible() bool`

HasSuperFollowEligible returns a boolean if a field has been set.

### GetSuperFollowedBy

`func (o *UserUnion) GetSuperFollowedBy() bool`

GetSuperFollowedBy returns the SuperFollowedBy field if non-nil, zero value otherwise.

### GetSuperFollowedByOk

`func (o *UserUnion) GetSuperFollowedByOk() (*bool, bool)`

GetSuperFollowedByOk returns a tuple with the SuperFollowedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperFollowedBy

`func (o *UserUnion) SetSuperFollowedBy(v bool)`

SetSuperFollowedBy sets SuperFollowedBy field to given value.

### HasSuperFollowedBy

`func (o *UserUnion) HasSuperFollowedBy() bool`

HasSuperFollowedBy returns a boolean if a field has been set.

### GetSuperFollowing

`func (o *UserUnion) GetSuperFollowing() bool`

GetSuperFollowing returns the SuperFollowing field if non-nil, zero value otherwise.

### GetSuperFollowingOk

`func (o *UserUnion) GetSuperFollowingOk() (*bool, bool)`

GetSuperFollowingOk returns a tuple with the SuperFollowing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperFollowing

`func (o *UserUnion) SetSuperFollowing(v bool)`

SetSuperFollowing sets SuperFollowing field to given value.

### HasSuperFollowing

`func (o *UserUnion) HasSuperFollowing() bool`

HasSuperFollowing returns a boolean if a field has been set.

### GetTipjarSettings

`func (o *UserUnion) GetTipjarSettings() UserTipJarSettings`

GetTipjarSettings returns the TipjarSettings field if non-nil, zero value otherwise.

### GetTipjarSettingsOk

`func (o *UserUnion) GetTipjarSettingsOk() (*UserTipJarSettings, bool)`

GetTipjarSettingsOk returns a tuple with the TipjarSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipjarSettings

`func (o *UserUnion) SetTipjarSettings(v UserTipJarSettings)`

SetTipjarSettings sets TipjarSettings field to given value.

### HasTipjarSettings

`func (o *UserUnion) HasTipjarSettings() bool`

HasTipjarSettings returns a boolean if a field has been set.

### GetUserSeedTweetCount

`func (o *UserUnion) GetUserSeedTweetCount() int32`

GetUserSeedTweetCount returns the UserSeedTweetCount field if non-nil, zero value otherwise.

### GetUserSeedTweetCountOk

`func (o *UserUnion) GetUserSeedTweetCountOk() (*int32, bool)`

GetUserSeedTweetCountOk returns a tuple with the UserSeedTweetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSeedTweetCount

`func (o *UserUnion) SetUserSeedTweetCount(v int32)`

SetUserSeedTweetCount sets UserSeedTweetCount field to given value.

### HasUserSeedTweetCount

`func (o *UserUnion) HasUserSeedTweetCount() bool`

HasUserSeedTweetCount returns a boolean if a field has been set.

### GetVerificationInfo

`func (o *UserUnion) GetVerificationInfo() UserVerificationInfo`

GetVerificationInfo returns the VerificationInfo field if non-nil, zero value otherwise.

### GetVerificationInfoOk

`func (o *UserUnion) GetVerificationInfoOk() (*UserVerificationInfo, bool)`

GetVerificationInfoOk returns a tuple with the VerificationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationInfo

`func (o *UserUnion) SetVerificationInfo(v UserVerificationInfo)`

SetVerificationInfo sets VerificationInfo field to given value.

### HasVerificationInfo

`func (o *UserUnion) HasVerificationInfo() bool`

HasVerificationInfo returns a boolean if a field has been set.

### GetMessage

`func (o *UserUnion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserUnion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserUnion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserUnion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *UserUnion) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UserUnion) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UserUnion) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


