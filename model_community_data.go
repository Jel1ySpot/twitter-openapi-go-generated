/*
Twitter OpenAPI

Twitter OpenAPI(Swagger) specification

API version: 0.0.1
Contact: yuki@yuki0311.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CommunityData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunityData{}

// CommunityData struct for CommunityData
type CommunityData struct {
	Typename TypeName `json:"__typename"`
	Actions CommunityActions `json:"actions"`
	AdminResults UserResults `json:"admin_results"`
	CreatedAt *int32 `json:"created_at,omitempty"`
	CreatorResults UserResults `json:"creator_results"`
	CustomBannerMedia map[string]interface{} `json:"custom_banner_media,omitempty"`
	DefaultBannerMedia map[string]interface{} `json:"default_banner_media,omitempty"`
	Description string `json:"description"`
	IdStr string `json:"id_str" validate:"regexp=^[0-9]+$"`
	InvitesPolicy string `json:"invites_policy"`
	InvitesResult CommunityInvitesResult `json:"invites_result"`
	IsPinned bool `json:"is_pinned"`
	JoinPolicy string `json:"join_policy"`
	JoinRequestsResult *CommunityJoinRequestsResult `json:"join_requests_result,omitempty"`
	MemberCount int32 `json:"member_count"`
	MembersFacepileResults []UserResults `json:"members_facepile_results"`
	ModeratorCount int32 `json:"moderator_count"`
	Name string `json:"name"`
	PrimaryCommunityTopic *PrimaryCommunityTopic `json:"primary_community_topic,omitempty"`
	Question *string `json:"question,omitempty"`
	Role string `json:"role"`
	Rules []CommunityRule `json:"rules"`
	SearchTags []string `json:"search_tags"`
	ShowOnlyUsersToDisplay []string `json:"show_only_users_to_display,omitempty"`
	Urls *CommunityUrls `json:"urls,omitempty"`
	ViewerRelationship map[string]interface{} `json:"viewer_relationship,omitempty"`
}

type _CommunityData CommunityData

// NewCommunityData instantiates a new CommunityData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityData(typename TypeName, actions CommunityActions, adminResults UserResults, creatorResults UserResults, description string, idStr string, invitesPolicy string, invitesResult CommunityInvitesResult, isPinned bool, joinPolicy string, memberCount int32, membersFacepileResults []UserResults, moderatorCount int32, name string, role string, rules []CommunityRule, searchTags []string) *CommunityData {
	this := CommunityData{}
	this.Typename = typename
	this.Actions = actions
	this.AdminResults = adminResults
	this.CreatorResults = creatorResults
	this.Description = description
	this.IdStr = idStr
	this.InvitesPolicy = invitesPolicy
	this.InvitesResult = invitesResult
	this.IsPinned = isPinned
	this.JoinPolicy = joinPolicy
	this.MemberCount = memberCount
	this.MembersFacepileResults = membersFacepileResults
	this.ModeratorCount = moderatorCount
	this.Name = name
	this.Role = role
	this.Rules = rules
	this.SearchTags = searchTags
	return &this
}

// NewCommunityDataWithDefaults instantiates a new CommunityData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityDataWithDefaults() *CommunityData {
	this := CommunityData{}
	return &this
}

// GetTypename returns the Typename field value
func (o *CommunityData) GetTypename() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetTypenameOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typename, true
}

// SetTypename sets field value
func (o *CommunityData) SetTypename(v TypeName) {
	o.Typename = v
}

// GetActions returns the Actions field value
func (o *CommunityData) GetActions() CommunityActions {
	if o == nil {
		var ret CommunityActions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetActionsOk() (*CommunityActions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value
func (o *CommunityData) SetActions(v CommunityActions) {
	o.Actions = v
}

// GetAdminResults returns the AdminResults field value
func (o *CommunityData) GetAdminResults() UserResults {
	if o == nil {
		var ret UserResults
		return ret
	}

	return o.AdminResults
}

// GetAdminResultsOk returns a tuple with the AdminResults field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetAdminResultsOk() (*UserResults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminResults, true
}

// SetAdminResults sets field value
func (o *CommunityData) SetAdminResults(v UserResults) {
	o.AdminResults = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CommunityData) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityData) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CommunityData) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *CommunityData) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetCreatorResults returns the CreatorResults field value
func (o *CommunityData) GetCreatorResults() UserResults {
	if o == nil {
		var ret UserResults
		return ret
	}

	return o.CreatorResults
}

// GetCreatorResultsOk returns a tuple with the CreatorResults field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetCreatorResultsOk() (*UserResults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorResults, true
}

// SetCreatorResults sets field value
func (o *CommunityData) SetCreatorResults(v UserResults) {
	o.CreatorResults = v
}

// GetCustomBannerMedia returns the CustomBannerMedia field value if set, zero value otherwise.
func (o *CommunityData) GetCustomBannerMedia() map[string]interface{} {
	if o == nil || IsNil(o.CustomBannerMedia) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomBannerMedia
}

// GetCustomBannerMediaOk returns a tuple with the CustomBannerMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityData) GetCustomBannerMediaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomBannerMedia) {
		return map[string]interface{}{}, false
	}
	return o.CustomBannerMedia, true
}

// HasCustomBannerMedia returns a boolean if a field has been set.
func (o *CommunityData) HasCustomBannerMedia() bool {
	if o != nil && !IsNil(o.CustomBannerMedia) {
		return true
	}

	return false
}

// SetCustomBannerMedia gets a reference to the given map[string]interface{} and assigns it to the CustomBannerMedia field.
func (o *CommunityData) SetCustomBannerMedia(v map[string]interface{}) {
	o.CustomBannerMedia = v
}

// GetDefaultBannerMedia returns the DefaultBannerMedia field value if set, zero value otherwise.
func (o *CommunityData) GetDefaultBannerMedia() map[string]interface{} {
	if o == nil || IsNil(o.DefaultBannerMedia) {
		var ret map[string]interface{}
		return ret
	}
	return o.DefaultBannerMedia
}

// GetDefaultBannerMediaOk returns a tuple with the DefaultBannerMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityData) GetDefaultBannerMediaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DefaultBannerMedia) {
		return map[string]interface{}{}, false
	}
	return o.DefaultBannerMedia, true
}

// HasDefaultBannerMedia returns a boolean if a field has been set.
func (o *CommunityData) HasDefaultBannerMedia() bool {
	if o != nil && !IsNil(o.DefaultBannerMedia) {
		return true
	}

	return false
}

// SetDefaultBannerMedia gets a reference to the given map[string]interface{} and assigns it to the DefaultBannerMedia field.
func (o *CommunityData) SetDefaultBannerMedia(v map[string]interface{}) {
	o.DefaultBannerMedia = v
}

// GetDescription returns the Description field value
func (o *CommunityData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CommunityData) SetDescription(v string) {
	o.Description = v
}

// GetIdStr returns the IdStr field value
func (o *CommunityData) GetIdStr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdStr
}

// GetIdStrOk returns a tuple with the IdStr field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetIdStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdStr, true
}

// SetIdStr sets field value
func (o *CommunityData) SetIdStr(v string) {
	o.IdStr = v
}

// GetInvitesPolicy returns the InvitesPolicy field value
func (o *CommunityData) GetInvitesPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvitesPolicy
}

// GetInvitesPolicyOk returns a tuple with the InvitesPolicy field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetInvitesPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvitesPolicy, true
}

// SetInvitesPolicy sets field value
func (o *CommunityData) SetInvitesPolicy(v string) {
	o.InvitesPolicy = v
}

// GetInvitesResult returns the InvitesResult field value
func (o *CommunityData) GetInvitesResult() CommunityInvitesResult {
	if o == nil {
		var ret CommunityInvitesResult
		return ret
	}

	return o.InvitesResult
}

// GetInvitesResultOk returns a tuple with the InvitesResult field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetInvitesResultOk() (*CommunityInvitesResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvitesResult, true
}

// SetInvitesResult sets field value
func (o *CommunityData) SetInvitesResult(v CommunityInvitesResult) {
	o.InvitesResult = v
}

// GetIsPinned returns the IsPinned field value
func (o *CommunityData) GetIsPinned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPinned
}

// GetIsPinnedOk returns a tuple with the IsPinned field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetIsPinnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPinned, true
}

// SetIsPinned sets field value
func (o *CommunityData) SetIsPinned(v bool) {
	o.IsPinned = v
}

// GetJoinPolicy returns the JoinPolicy field value
func (o *CommunityData) GetJoinPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinPolicy
}

// GetJoinPolicyOk returns a tuple with the JoinPolicy field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetJoinPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinPolicy, true
}

// SetJoinPolicy sets field value
func (o *CommunityData) SetJoinPolicy(v string) {
	o.JoinPolicy = v
}

// GetJoinRequestsResult returns the JoinRequestsResult field value if set, zero value otherwise.
func (o *CommunityData) GetJoinRequestsResult() CommunityJoinRequestsResult {
	if o == nil || IsNil(o.JoinRequestsResult) {
		var ret CommunityJoinRequestsResult
		return ret
	}
	return *o.JoinRequestsResult
}

// GetJoinRequestsResultOk returns a tuple with the JoinRequestsResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityData) GetJoinRequestsResultOk() (*CommunityJoinRequestsResult, bool) {
	if o == nil || IsNil(o.JoinRequestsResult) {
		return nil, false
	}
	return o.JoinRequestsResult, true
}

// HasJoinRequestsResult returns a boolean if a field has been set.
func (o *CommunityData) HasJoinRequestsResult() bool {
	if o != nil && !IsNil(o.JoinRequestsResult) {
		return true
	}

	return false
}

// SetJoinRequestsResult gets a reference to the given CommunityJoinRequestsResult and assigns it to the JoinRequestsResult field.
func (o *CommunityData) SetJoinRequestsResult(v CommunityJoinRequestsResult) {
	o.JoinRequestsResult = &v
}

// GetMemberCount returns the MemberCount field value
func (o *CommunityData) GetMemberCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetMemberCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberCount, true
}

// SetMemberCount sets field value
func (o *CommunityData) SetMemberCount(v int32) {
	o.MemberCount = v
}

// GetMembersFacepileResults returns the MembersFacepileResults field value
func (o *CommunityData) GetMembersFacepileResults() []UserResults {
	if o == nil {
		var ret []UserResults
		return ret
	}

	return o.MembersFacepileResults
}

// GetMembersFacepileResultsOk returns a tuple with the MembersFacepileResults field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetMembersFacepileResultsOk() ([]UserResults, bool) {
	if o == nil {
		return nil, false
	}
	return o.MembersFacepileResults, true
}

// SetMembersFacepileResults sets field value
func (o *CommunityData) SetMembersFacepileResults(v []UserResults) {
	o.MembersFacepileResults = v
}

// GetModeratorCount returns the ModeratorCount field value
func (o *CommunityData) GetModeratorCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ModeratorCount
}

// GetModeratorCountOk returns a tuple with the ModeratorCount field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetModeratorCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModeratorCount, true
}

// SetModeratorCount sets field value
func (o *CommunityData) SetModeratorCount(v int32) {
	o.ModeratorCount = v
}

// GetName returns the Name field value
func (o *CommunityData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CommunityData) SetName(v string) {
	o.Name = v
}

// GetPrimaryCommunityTopic returns the PrimaryCommunityTopic field value if set, zero value otherwise.
func (o *CommunityData) GetPrimaryCommunityTopic() PrimaryCommunityTopic {
	if o == nil || IsNil(o.PrimaryCommunityTopic) {
		var ret PrimaryCommunityTopic
		return ret
	}
	return *o.PrimaryCommunityTopic
}

// GetPrimaryCommunityTopicOk returns a tuple with the PrimaryCommunityTopic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityData) GetPrimaryCommunityTopicOk() (*PrimaryCommunityTopic, bool) {
	if o == nil || IsNil(o.PrimaryCommunityTopic) {
		return nil, false
	}
	return o.PrimaryCommunityTopic, true
}

// HasPrimaryCommunityTopic returns a boolean if a field has been set.
func (o *CommunityData) HasPrimaryCommunityTopic() bool {
	if o != nil && !IsNil(o.PrimaryCommunityTopic) {
		return true
	}

	return false
}

// SetPrimaryCommunityTopic gets a reference to the given PrimaryCommunityTopic and assigns it to the PrimaryCommunityTopic field.
func (o *CommunityData) SetPrimaryCommunityTopic(v PrimaryCommunityTopic) {
	o.PrimaryCommunityTopic = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *CommunityData) GetQuestion() string {
	if o == nil || IsNil(o.Question) {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityData) GetQuestionOk() (*string, bool) {
	if o == nil || IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *CommunityData) HasQuestion() bool {
	if o != nil && !IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *CommunityData) SetQuestion(v string) {
	o.Question = &v
}

// GetRole returns the Role field value
func (o *CommunityData) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *CommunityData) SetRole(v string) {
	o.Role = v
}

// GetRules returns the Rules field value
func (o *CommunityData) GetRules() []CommunityRule {
	if o == nil {
		var ret []CommunityRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetRulesOk() ([]CommunityRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *CommunityData) SetRules(v []CommunityRule) {
	o.Rules = v
}

// GetSearchTags returns the SearchTags field value
func (o *CommunityData) GetSearchTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SearchTags
}

// GetSearchTagsOk returns a tuple with the SearchTags field value
// and a boolean to check if the value has been set.
func (o *CommunityData) GetSearchTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchTags, true
}

// SetSearchTags sets field value
func (o *CommunityData) SetSearchTags(v []string) {
	o.SearchTags = v
}

// GetShowOnlyUsersToDisplay returns the ShowOnlyUsersToDisplay field value if set, zero value otherwise.
func (o *CommunityData) GetShowOnlyUsersToDisplay() []string {
	if o == nil || IsNil(o.ShowOnlyUsersToDisplay) {
		var ret []string
		return ret
	}
	return o.ShowOnlyUsersToDisplay
}

// GetShowOnlyUsersToDisplayOk returns a tuple with the ShowOnlyUsersToDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityData) GetShowOnlyUsersToDisplayOk() ([]string, bool) {
	if o == nil || IsNil(o.ShowOnlyUsersToDisplay) {
		return nil, false
	}
	return o.ShowOnlyUsersToDisplay, true
}

// HasShowOnlyUsersToDisplay returns a boolean if a field has been set.
func (o *CommunityData) HasShowOnlyUsersToDisplay() bool {
	if o != nil && !IsNil(o.ShowOnlyUsersToDisplay) {
		return true
	}

	return false
}

// SetShowOnlyUsersToDisplay gets a reference to the given []string and assigns it to the ShowOnlyUsersToDisplay field.
func (o *CommunityData) SetShowOnlyUsersToDisplay(v []string) {
	o.ShowOnlyUsersToDisplay = v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *CommunityData) GetUrls() CommunityUrls {
	if o == nil || IsNil(o.Urls) {
		var ret CommunityUrls
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityData) GetUrlsOk() (*CommunityUrls, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *CommunityData) HasUrls() bool {
	if o != nil && !IsNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given CommunityUrls and assigns it to the Urls field.
func (o *CommunityData) SetUrls(v CommunityUrls) {
	o.Urls = &v
}

// GetViewerRelationship returns the ViewerRelationship field value if set, zero value otherwise.
func (o *CommunityData) GetViewerRelationship() map[string]interface{} {
	if o == nil || IsNil(o.ViewerRelationship) {
		var ret map[string]interface{}
		return ret
	}
	return o.ViewerRelationship
}

// GetViewerRelationshipOk returns a tuple with the ViewerRelationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityData) GetViewerRelationshipOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ViewerRelationship) {
		return map[string]interface{}{}, false
	}
	return o.ViewerRelationship, true
}

// HasViewerRelationship returns a boolean if a field has been set.
func (o *CommunityData) HasViewerRelationship() bool {
	if o != nil && !IsNil(o.ViewerRelationship) {
		return true
	}

	return false
}

// SetViewerRelationship gets a reference to the given map[string]interface{} and assigns it to the ViewerRelationship field.
func (o *CommunityData) SetViewerRelationship(v map[string]interface{}) {
	o.ViewerRelationship = v
}

func (o CommunityData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunityData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["__typename"] = o.Typename
	toSerialize["actions"] = o.Actions
	toSerialize["admin_results"] = o.AdminResults
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["creator_results"] = o.CreatorResults
	if !IsNil(o.CustomBannerMedia) {
		toSerialize["custom_banner_media"] = o.CustomBannerMedia
	}
	if !IsNil(o.DefaultBannerMedia) {
		toSerialize["default_banner_media"] = o.DefaultBannerMedia
	}
	toSerialize["description"] = o.Description
	toSerialize["id_str"] = o.IdStr
	toSerialize["invites_policy"] = o.InvitesPolicy
	toSerialize["invites_result"] = o.InvitesResult
	toSerialize["is_pinned"] = o.IsPinned
	toSerialize["join_policy"] = o.JoinPolicy
	if !IsNil(o.JoinRequestsResult) {
		toSerialize["join_requests_result"] = o.JoinRequestsResult
	}
	toSerialize["member_count"] = o.MemberCount
	toSerialize["members_facepile_results"] = o.MembersFacepileResults
	toSerialize["moderator_count"] = o.ModeratorCount
	toSerialize["name"] = o.Name
	if !IsNil(o.PrimaryCommunityTopic) {
		toSerialize["primary_community_topic"] = o.PrimaryCommunityTopic
	}
	if !IsNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	toSerialize["role"] = o.Role
	toSerialize["rules"] = o.Rules
	toSerialize["search_tags"] = o.SearchTags
	if !IsNil(o.ShowOnlyUsersToDisplay) {
		toSerialize["show_only_users_to_display"] = o.ShowOnlyUsersToDisplay
	}
	if !IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	if !IsNil(o.ViewerRelationship) {
		toSerialize["viewer_relationship"] = o.ViewerRelationship
	}
	return toSerialize, nil
}

func (o *CommunityData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"__typename",
		"actions",
		"admin_results",
		"creator_results",
		"description",
		"id_str",
		"invites_policy",
		"invites_result",
		"is_pinned",
		"join_policy",
		"member_count",
		"members_facepile_results",
		"moderator_count",
		"name",
		"role",
		"rules",
		"search_tags",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCommunityData := _CommunityData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommunityData)

	if err != nil {
		return err
	}

	*o = CommunityData(varCommunityData)

	return err
}

type NullableCommunityData struct {
	value *CommunityData
	isSet bool
}

func (v NullableCommunityData) Get() *CommunityData {
	return v.value
}

func (v *NullableCommunityData) Set(val *CommunityData) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityData) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityData(val *CommunityData) *NullableCommunityData {
	return &NullableCommunityData{value: val, isSet: true}
}

func (v NullableCommunityData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


