# CommunityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**Actions** | [**CommunityActions**](CommunityActions.md) |  | 
**AdminResults** | [**UserResults**](UserResults.md) |  | 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**CreatorResults** | [**UserResults**](UserResults.md) |  | 
**CustomBannerMedia** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultBannerMedia** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | **string** |  | 
**IdStr** | **string** |  | 
**InvitesPolicy** | **string** |  | 
**InvitesResult** | [**CommunityInvitesResult**](CommunityInvitesResult.md) |  | 
**IsPinned** | **bool** |  | 
**JoinPolicy** | **string** |  | 
**JoinRequestsResult** | Pointer to [**CommunityJoinRequestsResult**](CommunityJoinRequestsResult.md) |  | [optional] 
**MemberCount** | **int32** |  | 
**MembersFacepileResults** | [**[]UserResults**](UserResults.md) |  | 
**ModeratorCount** | **int32** |  | 
**Name** | **string** |  | 
**PrimaryCommunityTopic** | Pointer to [**PrimaryCommunityTopic**](PrimaryCommunityTopic.md) |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | 
**Rules** | [**[]CommunityRule**](CommunityRule.md) |  | 
**SearchTags** | **[]string** |  | 
**ShowOnlyUsersToDisplay** | Pointer to **[]string** |  | [optional] 
**Urls** | Pointer to [**CommunityUrls**](CommunityUrls.md) |  | [optional] 
**ViewerRelationship** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCommunityData

`func NewCommunityData(typename TypeName, actions CommunityActions, adminResults UserResults, creatorResults UserResults, description string, idStr string, invitesPolicy string, invitesResult CommunityInvitesResult, isPinned bool, joinPolicy string, memberCount int32, membersFacepileResults []UserResults, moderatorCount int32, name string, role string, rules []CommunityRule, searchTags []string, ) *CommunityData`

NewCommunityData instantiates a new CommunityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunityDataWithDefaults

`func NewCommunityDataWithDefaults() *CommunityData`

NewCommunityDataWithDefaults instantiates a new CommunityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *CommunityData) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *CommunityData) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *CommunityData) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetActions

`func (o *CommunityData) GetActions() CommunityActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CommunityData) GetActionsOk() (*CommunityActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CommunityData) SetActions(v CommunityActions)`

SetActions sets Actions field to given value.


### GetAdminResults

`func (o *CommunityData) GetAdminResults() UserResults`

GetAdminResults returns the AdminResults field if non-nil, zero value otherwise.

### GetAdminResultsOk

`func (o *CommunityData) GetAdminResultsOk() (*UserResults, bool)`

GetAdminResultsOk returns a tuple with the AdminResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminResults

`func (o *CommunityData) SetAdminResults(v UserResults)`

SetAdminResults sets AdminResults field to given value.


### GetCreatedAt

`func (o *CommunityData) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommunityData) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommunityData) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommunityData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatorResults

`func (o *CommunityData) GetCreatorResults() UserResults`

GetCreatorResults returns the CreatorResults field if non-nil, zero value otherwise.

### GetCreatorResultsOk

`func (o *CommunityData) GetCreatorResultsOk() (*UserResults, bool)`

GetCreatorResultsOk returns a tuple with the CreatorResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorResults

`func (o *CommunityData) SetCreatorResults(v UserResults)`

SetCreatorResults sets CreatorResults field to given value.


### GetCustomBannerMedia

`func (o *CommunityData) GetCustomBannerMedia() map[string]interface{}`

GetCustomBannerMedia returns the CustomBannerMedia field if non-nil, zero value otherwise.

### GetCustomBannerMediaOk

`func (o *CommunityData) GetCustomBannerMediaOk() (*map[string]interface{}, bool)`

GetCustomBannerMediaOk returns a tuple with the CustomBannerMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBannerMedia

`func (o *CommunityData) SetCustomBannerMedia(v map[string]interface{})`

SetCustomBannerMedia sets CustomBannerMedia field to given value.

### HasCustomBannerMedia

`func (o *CommunityData) HasCustomBannerMedia() bool`

HasCustomBannerMedia returns a boolean if a field has been set.

### GetDefaultBannerMedia

`func (o *CommunityData) GetDefaultBannerMedia() map[string]interface{}`

GetDefaultBannerMedia returns the DefaultBannerMedia field if non-nil, zero value otherwise.

### GetDefaultBannerMediaOk

`func (o *CommunityData) GetDefaultBannerMediaOk() (*map[string]interface{}, bool)`

GetDefaultBannerMediaOk returns a tuple with the DefaultBannerMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBannerMedia

`func (o *CommunityData) SetDefaultBannerMedia(v map[string]interface{})`

SetDefaultBannerMedia sets DefaultBannerMedia field to given value.

### HasDefaultBannerMedia

`func (o *CommunityData) HasDefaultBannerMedia() bool`

HasDefaultBannerMedia returns a boolean if a field has been set.

### GetDescription

`func (o *CommunityData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommunityData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommunityData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIdStr

`func (o *CommunityData) GetIdStr() string`

GetIdStr returns the IdStr field if non-nil, zero value otherwise.

### GetIdStrOk

`func (o *CommunityData) GetIdStrOk() (*string, bool)`

GetIdStrOk returns a tuple with the IdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStr

`func (o *CommunityData) SetIdStr(v string)`

SetIdStr sets IdStr field to given value.


### GetInvitesPolicy

`func (o *CommunityData) GetInvitesPolicy() string`

GetInvitesPolicy returns the InvitesPolicy field if non-nil, zero value otherwise.

### GetInvitesPolicyOk

`func (o *CommunityData) GetInvitesPolicyOk() (*string, bool)`

GetInvitesPolicyOk returns a tuple with the InvitesPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitesPolicy

`func (o *CommunityData) SetInvitesPolicy(v string)`

SetInvitesPolicy sets InvitesPolicy field to given value.


### GetInvitesResult

`func (o *CommunityData) GetInvitesResult() CommunityInvitesResult`

GetInvitesResult returns the InvitesResult field if non-nil, zero value otherwise.

### GetInvitesResultOk

`func (o *CommunityData) GetInvitesResultOk() (*CommunityInvitesResult, bool)`

GetInvitesResultOk returns a tuple with the InvitesResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitesResult

`func (o *CommunityData) SetInvitesResult(v CommunityInvitesResult)`

SetInvitesResult sets InvitesResult field to given value.


### GetIsPinned

`func (o *CommunityData) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *CommunityData) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *CommunityData) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.


### GetJoinPolicy

`func (o *CommunityData) GetJoinPolicy() string`

GetJoinPolicy returns the JoinPolicy field if non-nil, zero value otherwise.

### GetJoinPolicyOk

`func (o *CommunityData) GetJoinPolicyOk() (*string, bool)`

GetJoinPolicyOk returns a tuple with the JoinPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinPolicy

`func (o *CommunityData) SetJoinPolicy(v string)`

SetJoinPolicy sets JoinPolicy field to given value.


### GetJoinRequestsResult

`func (o *CommunityData) GetJoinRequestsResult() CommunityJoinRequestsResult`

GetJoinRequestsResult returns the JoinRequestsResult field if non-nil, zero value otherwise.

### GetJoinRequestsResultOk

`func (o *CommunityData) GetJoinRequestsResultOk() (*CommunityJoinRequestsResult, bool)`

GetJoinRequestsResultOk returns a tuple with the JoinRequestsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinRequestsResult

`func (o *CommunityData) SetJoinRequestsResult(v CommunityJoinRequestsResult)`

SetJoinRequestsResult sets JoinRequestsResult field to given value.

### HasJoinRequestsResult

`func (o *CommunityData) HasJoinRequestsResult() bool`

HasJoinRequestsResult returns a boolean if a field has been set.

### GetMemberCount

`func (o *CommunityData) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *CommunityData) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *CommunityData) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetMembersFacepileResults

`func (o *CommunityData) GetMembersFacepileResults() []UserResults`

GetMembersFacepileResults returns the MembersFacepileResults field if non-nil, zero value otherwise.

### GetMembersFacepileResultsOk

`func (o *CommunityData) GetMembersFacepileResultsOk() (*[]UserResults, bool)`

GetMembersFacepileResultsOk returns a tuple with the MembersFacepileResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersFacepileResults

`func (o *CommunityData) SetMembersFacepileResults(v []UserResults)`

SetMembersFacepileResults sets MembersFacepileResults field to given value.


### GetModeratorCount

`func (o *CommunityData) GetModeratorCount() int32`

GetModeratorCount returns the ModeratorCount field if non-nil, zero value otherwise.

### GetModeratorCountOk

`func (o *CommunityData) GetModeratorCountOk() (*int32, bool)`

GetModeratorCountOk returns a tuple with the ModeratorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeratorCount

`func (o *CommunityData) SetModeratorCount(v int32)`

SetModeratorCount sets ModeratorCount field to given value.


### GetName

`func (o *CommunityData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommunityData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommunityData) SetName(v string)`

SetName sets Name field to given value.


### GetPrimaryCommunityTopic

`func (o *CommunityData) GetPrimaryCommunityTopic() PrimaryCommunityTopic`

GetPrimaryCommunityTopic returns the PrimaryCommunityTopic field if non-nil, zero value otherwise.

### GetPrimaryCommunityTopicOk

`func (o *CommunityData) GetPrimaryCommunityTopicOk() (*PrimaryCommunityTopic, bool)`

GetPrimaryCommunityTopicOk returns a tuple with the PrimaryCommunityTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCommunityTopic

`func (o *CommunityData) SetPrimaryCommunityTopic(v PrimaryCommunityTopic)`

SetPrimaryCommunityTopic sets PrimaryCommunityTopic field to given value.

### HasPrimaryCommunityTopic

`func (o *CommunityData) HasPrimaryCommunityTopic() bool`

HasPrimaryCommunityTopic returns a boolean if a field has been set.

### GetQuestion

`func (o *CommunityData) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *CommunityData) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *CommunityData) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *CommunityData) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetRole

`func (o *CommunityData) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CommunityData) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CommunityData) SetRole(v string)`

SetRole sets Role field to given value.


### GetRules

`func (o *CommunityData) GetRules() []CommunityRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CommunityData) GetRulesOk() (*[]CommunityRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CommunityData) SetRules(v []CommunityRule)`

SetRules sets Rules field to given value.


### GetSearchTags

`func (o *CommunityData) GetSearchTags() []string`

GetSearchTags returns the SearchTags field if non-nil, zero value otherwise.

### GetSearchTagsOk

`func (o *CommunityData) GetSearchTagsOk() (*[]string, bool)`

GetSearchTagsOk returns a tuple with the SearchTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTags

`func (o *CommunityData) SetSearchTags(v []string)`

SetSearchTags sets SearchTags field to given value.


### GetShowOnlyUsersToDisplay

`func (o *CommunityData) GetShowOnlyUsersToDisplay() []string`

GetShowOnlyUsersToDisplay returns the ShowOnlyUsersToDisplay field if non-nil, zero value otherwise.

### GetShowOnlyUsersToDisplayOk

`func (o *CommunityData) GetShowOnlyUsersToDisplayOk() (*[]string, bool)`

GetShowOnlyUsersToDisplayOk returns a tuple with the ShowOnlyUsersToDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnlyUsersToDisplay

`func (o *CommunityData) SetShowOnlyUsersToDisplay(v []string)`

SetShowOnlyUsersToDisplay sets ShowOnlyUsersToDisplay field to given value.

### HasShowOnlyUsersToDisplay

`func (o *CommunityData) HasShowOnlyUsersToDisplay() bool`

HasShowOnlyUsersToDisplay returns a boolean if a field has been set.

### GetUrls

`func (o *CommunityData) GetUrls() CommunityUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *CommunityData) GetUrlsOk() (*CommunityUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *CommunityData) SetUrls(v CommunityUrls)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *CommunityData) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetViewerRelationship

`func (o *CommunityData) GetViewerRelationship() map[string]interface{}`

GetViewerRelationship returns the ViewerRelationship field if non-nil, zero value otherwise.

### GetViewerRelationshipOk

`func (o *CommunityData) GetViewerRelationshipOk() (*map[string]interface{}, bool)`

GetViewerRelationshipOk returns a tuple with the ViewerRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerRelationship

`func (o *CommunityData) SetViewerRelationship(v map[string]interface{})`

SetViewerRelationship sets ViewerRelationship field to given value.

### HasViewerRelationship

`func (o *CommunityData) HasViewerRelationship() bool`

HasViewerRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


