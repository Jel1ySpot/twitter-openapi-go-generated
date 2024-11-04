# UserLegacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedBy** | Pointer to **bool** |  | [optional] 
**Blocking** | Pointer to **bool** |  | [optional] 
**CanDm** | Pointer to **bool** |  | [optional] 
**CanMediaTag** | Pointer to **bool** |  | [optional] 
**CreatedAt** | **string** |  | 
**DefaultProfile** | **bool** |  | 
**DefaultProfileImage** | **bool** |  | 
**Description** | **string** |  | 
**Entities** | **map[string]interface{}** |  | 
**FastFollowersCount** | **int32** |  | 
**FavouritesCount** | **int32** |  | 
**FollowRequestSent** | Pointer to **bool** |  | [optional] 
**FollowedBy** | Pointer to **bool** |  | [optional] 
**FollowersCount** | **int32** |  | 
**Following** | Pointer to **bool** |  | [optional] 
**FriendsCount** | **int32** |  | 
**HasCustomTimelines** | **bool** |  | 
**IsTranslator** | **bool** |  | 
**ListedCount** | **int32** |  | 
**Location** | **string** |  | 
**MediaCount** | **int32** |  | 
**Muting** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**NormalFollowersCount** | **int32** |  | 
**Notifications** | Pointer to **bool** |  | [optional] 
**PinnedTweetIdsStr** | **[]string** |  | 
**PossiblySensitive** | **bool** |  | 
**ProfileBannerExtensions** | Pointer to **map[string]interface{}** |  | [optional] 
**ProfileBannerUrl** | Pointer to **string** |  | [optional] 
**ProfileImageExtensions** | Pointer to **map[string]interface{}** |  | [optional] 
**ProfileImageUrlHttps** | **string** |  | 
**ProfileInterstitialType** | **string** |  | 
**Protected** | Pointer to **bool** |  | [optional] 
**ScreenName** | **string** |  | 
**StatusesCount** | **int32** |  | 
**TranslatorType** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 
**Verified** | **bool** |  | 
**VerifiedType** | Pointer to **string** |  | [optional] 
**WantRetweets** | Pointer to **bool** |  | [optional] 
**WithheldInCountries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserLegacy

`func NewUserLegacy(createdAt string, defaultProfile bool, defaultProfileImage bool, description string, entities map[string]interface{}, fastFollowersCount int32, favouritesCount int32, followersCount int32, friendsCount int32, hasCustomTimelines bool, isTranslator bool, listedCount int32, location string, mediaCount int32, name string, normalFollowersCount int32, pinnedTweetIdsStr []string, possiblySensitive bool, profileImageUrlHttps string, profileInterstitialType string, screenName string, statusesCount int32, translatorType string, verified bool, ) *UserLegacy`

NewUserLegacy instantiates a new UserLegacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLegacyWithDefaults

`func NewUserLegacyWithDefaults() *UserLegacy`

NewUserLegacyWithDefaults instantiates a new UserLegacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedBy

`func (o *UserLegacy) GetBlockedBy() bool`

GetBlockedBy returns the BlockedBy field if non-nil, zero value otherwise.

### GetBlockedByOk

`func (o *UserLegacy) GetBlockedByOk() (*bool, bool)`

GetBlockedByOk returns a tuple with the BlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBy

`func (o *UserLegacy) SetBlockedBy(v bool)`

SetBlockedBy sets BlockedBy field to given value.

### HasBlockedBy

`func (o *UserLegacy) HasBlockedBy() bool`

HasBlockedBy returns a boolean if a field has been set.

### GetBlocking

`func (o *UserLegacy) GetBlocking() bool`

GetBlocking returns the Blocking field if non-nil, zero value otherwise.

### GetBlockingOk

`func (o *UserLegacy) GetBlockingOk() (*bool, bool)`

GetBlockingOk returns a tuple with the Blocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocking

`func (o *UserLegacy) SetBlocking(v bool)`

SetBlocking sets Blocking field to given value.

### HasBlocking

`func (o *UserLegacy) HasBlocking() bool`

HasBlocking returns a boolean if a field has been set.

### GetCanDm

`func (o *UserLegacy) GetCanDm() bool`

GetCanDm returns the CanDm field if non-nil, zero value otherwise.

### GetCanDmOk

`func (o *UserLegacy) GetCanDmOk() (*bool, bool)`

GetCanDmOk returns a tuple with the CanDm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDm

`func (o *UserLegacy) SetCanDm(v bool)`

SetCanDm sets CanDm field to given value.

### HasCanDm

`func (o *UserLegacy) HasCanDm() bool`

HasCanDm returns a boolean if a field has been set.

### GetCanMediaTag

`func (o *UserLegacy) GetCanMediaTag() bool`

GetCanMediaTag returns the CanMediaTag field if non-nil, zero value otherwise.

### GetCanMediaTagOk

`func (o *UserLegacy) GetCanMediaTagOk() (*bool, bool)`

GetCanMediaTagOk returns a tuple with the CanMediaTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMediaTag

`func (o *UserLegacy) SetCanMediaTag(v bool)`

SetCanMediaTag sets CanMediaTag field to given value.

### HasCanMediaTag

`func (o *UserLegacy) HasCanMediaTag() bool`

HasCanMediaTag returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserLegacy) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserLegacy) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserLegacy) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultProfile

`func (o *UserLegacy) GetDefaultProfile() bool`

GetDefaultProfile returns the DefaultProfile field if non-nil, zero value otherwise.

### GetDefaultProfileOk

`func (o *UserLegacy) GetDefaultProfileOk() (*bool, bool)`

GetDefaultProfileOk returns a tuple with the DefaultProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProfile

`func (o *UserLegacy) SetDefaultProfile(v bool)`

SetDefaultProfile sets DefaultProfile field to given value.


### GetDefaultProfileImage

`func (o *UserLegacy) GetDefaultProfileImage() bool`

GetDefaultProfileImage returns the DefaultProfileImage field if non-nil, zero value otherwise.

### GetDefaultProfileImageOk

`func (o *UserLegacy) GetDefaultProfileImageOk() (*bool, bool)`

GetDefaultProfileImageOk returns a tuple with the DefaultProfileImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProfileImage

`func (o *UserLegacy) SetDefaultProfileImage(v bool)`

SetDefaultProfileImage sets DefaultProfileImage field to given value.


### GetDescription

`func (o *UserLegacy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserLegacy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserLegacy) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntities

`func (o *UserLegacy) GetEntities() map[string]interface{}`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *UserLegacy) GetEntitiesOk() (*map[string]interface{}, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *UserLegacy) SetEntities(v map[string]interface{})`

SetEntities sets Entities field to given value.


### GetFastFollowersCount

`func (o *UserLegacy) GetFastFollowersCount() int32`

GetFastFollowersCount returns the FastFollowersCount field if non-nil, zero value otherwise.

### GetFastFollowersCountOk

`func (o *UserLegacy) GetFastFollowersCountOk() (*int32, bool)`

GetFastFollowersCountOk returns a tuple with the FastFollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastFollowersCount

`func (o *UserLegacy) SetFastFollowersCount(v int32)`

SetFastFollowersCount sets FastFollowersCount field to given value.


### GetFavouritesCount

`func (o *UserLegacy) GetFavouritesCount() int32`

GetFavouritesCount returns the FavouritesCount field if non-nil, zero value otherwise.

### GetFavouritesCountOk

`func (o *UserLegacy) GetFavouritesCountOk() (*int32, bool)`

GetFavouritesCountOk returns a tuple with the FavouritesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavouritesCount

`func (o *UserLegacy) SetFavouritesCount(v int32)`

SetFavouritesCount sets FavouritesCount field to given value.


### GetFollowRequestSent

`func (o *UserLegacy) GetFollowRequestSent() bool`

GetFollowRequestSent returns the FollowRequestSent field if non-nil, zero value otherwise.

### GetFollowRequestSentOk

`func (o *UserLegacy) GetFollowRequestSentOk() (*bool, bool)`

GetFollowRequestSentOk returns a tuple with the FollowRequestSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowRequestSent

`func (o *UserLegacy) SetFollowRequestSent(v bool)`

SetFollowRequestSent sets FollowRequestSent field to given value.

### HasFollowRequestSent

`func (o *UserLegacy) HasFollowRequestSent() bool`

HasFollowRequestSent returns a boolean if a field has been set.

### GetFollowedBy

`func (o *UserLegacy) GetFollowedBy() bool`

GetFollowedBy returns the FollowedBy field if non-nil, zero value otherwise.

### GetFollowedByOk

`func (o *UserLegacy) GetFollowedByOk() (*bool, bool)`

GetFollowedByOk returns a tuple with the FollowedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowedBy

`func (o *UserLegacy) SetFollowedBy(v bool)`

SetFollowedBy sets FollowedBy field to given value.

### HasFollowedBy

`func (o *UserLegacy) HasFollowedBy() bool`

HasFollowedBy returns a boolean if a field has been set.

### GetFollowersCount

`func (o *UserLegacy) GetFollowersCount() int32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *UserLegacy) GetFollowersCountOk() (*int32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *UserLegacy) SetFollowersCount(v int32)`

SetFollowersCount sets FollowersCount field to given value.


### GetFollowing

`func (o *UserLegacy) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *UserLegacy) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *UserLegacy) SetFollowing(v bool)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *UserLegacy) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetFriendsCount

`func (o *UserLegacy) GetFriendsCount() int32`

GetFriendsCount returns the FriendsCount field if non-nil, zero value otherwise.

### GetFriendsCountOk

`func (o *UserLegacy) GetFriendsCountOk() (*int32, bool)`

GetFriendsCountOk returns a tuple with the FriendsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendsCount

`func (o *UserLegacy) SetFriendsCount(v int32)`

SetFriendsCount sets FriendsCount field to given value.


### GetHasCustomTimelines

`func (o *UserLegacy) GetHasCustomTimelines() bool`

GetHasCustomTimelines returns the HasCustomTimelines field if non-nil, zero value otherwise.

### GetHasCustomTimelinesOk

`func (o *UserLegacy) GetHasCustomTimelinesOk() (*bool, bool)`

GetHasCustomTimelinesOk returns a tuple with the HasCustomTimelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomTimelines

`func (o *UserLegacy) SetHasCustomTimelines(v bool)`

SetHasCustomTimelines sets HasCustomTimelines field to given value.


### GetIsTranslator

`func (o *UserLegacy) GetIsTranslator() bool`

GetIsTranslator returns the IsTranslator field if non-nil, zero value otherwise.

### GetIsTranslatorOk

`func (o *UserLegacy) GetIsTranslatorOk() (*bool, bool)`

GetIsTranslatorOk returns a tuple with the IsTranslator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTranslator

`func (o *UserLegacy) SetIsTranslator(v bool)`

SetIsTranslator sets IsTranslator field to given value.


### GetListedCount

`func (o *UserLegacy) GetListedCount() int32`

GetListedCount returns the ListedCount field if non-nil, zero value otherwise.

### GetListedCountOk

`func (o *UserLegacy) GetListedCountOk() (*int32, bool)`

GetListedCountOk returns a tuple with the ListedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedCount

`func (o *UserLegacy) SetListedCount(v int32)`

SetListedCount sets ListedCount field to given value.


### GetLocation

`func (o *UserLegacy) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UserLegacy) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UserLegacy) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetMediaCount

`func (o *UserLegacy) GetMediaCount() int32`

GetMediaCount returns the MediaCount field if non-nil, zero value otherwise.

### GetMediaCountOk

`func (o *UserLegacy) GetMediaCountOk() (*int32, bool)`

GetMediaCountOk returns a tuple with the MediaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaCount

`func (o *UserLegacy) SetMediaCount(v int32)`

SetMediaCount sets MediaCount field to given value.


### GetMuting

`func (o *UserLegacy) GetMuting() bool`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *UserLegacy) GetMutingOk() (*bool, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *UserLegacy) SetMuting(v bool)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *UserLegacy) HasMuting() bool`

HasMuting returns a boolean if a field has been set.

### GetName

`func (o *UserLegacy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserLegacy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserLegacy) SetName(v string)`

SetName sets Name field to given value.


### GetNormalFollowersCount

`func (o *UserLegacy) GetNormalFollowersCount() int32`

GetNormalFollowersCount returns the NormalFollowersCount field if non-nil, zero value otherwise.

### GetNormalFollowersCountOk

`func (o *UserLegacy) GetNormalFollowersCountOk() (*int32, bool)`

GetNormalFollowersCountOk returns a tuple with the NormalFollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalFollowersCount

`func (o *UserLegacy) SetNormalFollowersCount(v int32)`

SetNormalFollowersCount sets NormalFollowersCount field to given value.


### GetNotifications

`func (o *UserLegacy) GetNotifications() bool`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *UserLegacy) GetNotificationsOk() (*bool, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *UserLegacy) SetNotifications(v bool)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *UserLegacy) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetPinnedTweetIdsStr

`func (o *UserLegacy) GetPinnedTweetIdsStr() []string`

GetPinnedTweetIdsStr returns the PinnedTweetIdsStr field if non-nil, zero value otherwise.

### GetPinnedTweetIdsStrOk

`func (o *UserLegacy) GetPinnedTweetIdsStrOk() (*[]string, bool)`

GetPinnedTweetIdsStrOk returns a tuple with the PinnedTweetIdsStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedTweetIdsStr

`func (o *UserLegacy) SetPinnedTweetIdsStr(v []string)`

SetPinnedTweetIdsStr sets PinnedTweetIdsStr field to given value.


### GetPossiblySensitive

`func (o *UserLegacy) GetPossiblySensitive() bool`

GetPossiblySensitive returns the PossiblySensitive field if non-nil, zero value otherwise.

### GetPossiblySensitiveOk

`func (o *UserLegacy) GetPossiblySensitiveOk() (*bool, bool)`

GetPossiblySensitiveOk returns a tuple with the PossiblySensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossiblySensitive

`func (o *UserLegacy) SetPossiblySensitive(v bool)`

SetPossiblySensitive sets PossiblySensitive field to given value.


### GetProfileBannerExtensions

`func (o *UserLegacy) GetProfileBannerExtensions() map[string]interface{}`

GetProfileBannerExtensions returns the ProfileBannerExtensions field if non-nil, zero value otherwise.

### GetProfileBannerExtensionsOk

`func (o *UserLegacy) GetProfileBannerExtensionsOk() (*map[string]interface{}, bool)`

GetProfileBannerExtensionsOk returns a tuple with the ProfileBannerExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileBannerExtensions

`func (o *UserLegacy) SetProfileBannerExtensions(v map[string]interface{})`

SetProfileBannerExtensions sets ProfileBannerExtensions field to given value.

### HasProfileBannerExtensions

`func (o *UserLegacy) HasProfileBannerExtensions() bool`

HasProfileBannerExtensions returns a boolean if a field has been set.

### GetProfileBannerUrl

`func (o *UserLegacy) GetProfileBannerUrl() string`

GetProfileBannerUrl returns the ProfileBannerUrl field if non-nil, zero value otherwise.

### GetProfileBannerUrlOk

`func (o *UserLegacy) GetProfileBannerUrlOk() (*string, bool)`

GetProfileBannerUrlOk returns a tuple with the ProfileBannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileBannerUrl

`func (o *UserLegacy) SetProfileBannerUrl(v string)`

SetProfileBannerUrl sets ProfileBannerUrl field to given value.

### HasProfileBannerUrl

`func (o *UserLegacy) HasProfileBannerUrl() bool`

HasProfileBannerUrl returns a boolean if a field has been set.

### GetProfileImageExtensions

`func (o *UserLegacy) GetProfileImageExtensions() map[string]interface{}`

GetProfileImageExtensions returns the ProfileImageExtensions field if non-nil, zero value otherwise.

### GetProfileImageExtensionsOk

`func (o *UserLegacy) GetProfileImageExtensionsOk() (*map[string]interface{}, bool)`

GetProfileImageExtensionsOk returns a tuple with the ProfileImageExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageExtensions

`func (o *UserLegacy) SetProfileImageExtensions(v map[string]interface{})`

SetProfileImageExtensions sets ProfileImageExtensions field to given value.

### HasProfileImageExtensions

`func (o *UserLegacy) HasProfileImageExtensions() bool`

HasProfileImageExtensions returns a boolean if a field has been set.

### GetProfileImageUrlHttps

`func (o *UserLegacy) GetProfileImageUrlHttps() string`

GetProfileImageUrlHttps returns the ProfileImageUrlHttps field if non-nil, zero value otherwise.

### GetProfileImageUrlHttpsOk

`func (o *UserLegacy) GetProfileImageUrlHttpsOk() (*string, bool)`

GetProfileImageUrlHttpsOk returns a tuple with the ProfileImageUrlHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageUrlHttps

`func (o *UserLegacy) SetProfileImageUrlHttps(v string)`

SetProfileImageUrlHttps sets ProfileImageUrlHttps field to given value.


### GetProfileInterstitialType

`func (o *UserLegacy) GetProfileInterstitialType() string`

GetProfileInterstitialType returns the ProfileInterstitialType field if non-nil, zero value otherwise.

### GetProfileInterstitialTypeOk

`func (o *UserLegacy) GetProfileInterstitialTypeOk() (*string, bool)`

GetProfileInterstitialTypeOk returns a tuple with the ProfileInterstitialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileInterstitialType

`func (o *UserLegacy) SetProfileInterstitialType(v string)`

SetProfileInterstitialType sets ProfileInterstitialType field to given value.


### GetProtected

`func (o *UserLegacy) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *UserLegacy) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *UserLegacy) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *UserLegacy) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetScreenName

`func (o *UserLegacy) GetScreenName() string`

GetScreenName returns the ScreenName field if non-nil, zero value otherwise.

### GetScreenNameOk

`func (o *UserLegacy) GetScreenNameOk() (*string, bool)`

GetScreenNameOk returns a tuple with the ScreenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenName

`func (o *UserLegacy) SetScreenName(v string)`

SetScreenName sets ScreenName field to given value.


### GetStatusesCount

`func (o *UserLegacy) GetStatusesCount() int32`

GetStatusesCount returns the StatusesCount field if non-nil, zero value otherwise.

### GetStatusesCountOk

`func (o *UserLegacy) GetStatusesCountOk() (*int32, bool)`

GetStatusesCountOk returns a tuple with the StatusesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesCount

`func (o *UserLegacy) SetStatusesCount(v int32)`

SetStatusesCount sets StatusesCount field to given value.


### GetTranslatorType

`func (o *UserLegacy) GetTranslatorType() string`

GetTranslatorType returns the TranslatorType field if non-nil, zero value otherwise.

### GetTranslatorTypeOk

`func (o *UserLegacy) GetTranslatorTypeOk() (*string, bool)`

GetTranslatorTypeOk returns a tuple with the TranslatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatorType

`func (o *UserLegacy) SetTranslatorType(v string)`

SetTranslatorType sets TranslatorType field to given value.


### GetUrl

`func (o *UserLegacy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserLegacy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserLegacy) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserLegacy) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVerified

`func (o *UserLegacy) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserLegacy) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserLegacy) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVerifiedType

`func (o *UserLegacy) GetVerifiedType() string`

GetVerifiedType returns the VerifiedType field if non-nil, zero value otherwise.

### GetVerifiedTypeOk

`func (o *UserLegacy) GetVerifiedTypeOk() (*string, bool)`

GetVerifiedTypeOk returns a tuple with the VerifiedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedType

`func (o *UserLegacy) SetVerifiedType(v string)`

SetVerifiedType sets VerifiedType field to given value.

### HasVerifiedType

`func (o *UserLegacy) HasVerifiedType() bool`

HasVerifiedType returns a boolean if a field has been set.

### GetWantRetweets

`func (o *UserLegacy) GetWantRetweets() bool`

GetWantRetweets returns the WantRetweets field if non-nil, zero value otherwise.

### GetWantRetweetsOk

`func (o *UserLegacy) GetWantRetweetsOk() (*bool, bool)`

GetWantRetweetsOk returns a tuple with the WantRetweets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantRetweets

`func (o *UserLegacy) SetWantRetweets(v bool)`

SetWantRetweets sets WantRetweets field to given value.

### HasWantRetweets

`func (o *UserLegacy) HasWantRetweets() bool`

HasWantRetweets returns a boolean if a field has been set.

### GetWithheldInCountries

`func (o *UserLegacy) GetWithheldInCountries() []string`

GetWithheldInCountries returns the WithheldInCountries field if non-nil, zero value otherwise.

### GetWithheldInCountriesOk

`func (o *UserLegacy) GetWithheldInCountriesOk() (*[]string, bool)`

GetWithheldInCountriesOk returns a tuple with the WithheldInCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldInCountries

`func (o *UserLegacy) SetWithheldInCountries(v []string)`

SetWithheldInCountries sets WithheldInCountries field to given value.

### HasWithheldInCountries

`func (o *UserLegacy) HasWithheldInCountries() bool`

HasWithheldInCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


