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
	"fmt"
)

// TypeName the model 'TypeName'
type TypeName string

// List of TypeName
const (
	TIMELINE_TWEET                           TypeName = "TimelineTweet"
	TIMELINE_TIMELINE_ITEM                   TypeName = "TimelineTimelineItem"
	TIMELINE_USER                            TypeName = "TimelineUser"
	TIMELINE_TIMELINE_CURSOR                 TypeName = "TimelineTimelineCursor"
	TWEET_WITH_VISIBILITY_RESULTS            TypeName = "TweetWithVisibilityResults"
	CONTEXTUAL_TWEET_INTERSTITIAL            TypeName = "ContextualTweetInterstitial"
	TIMELINE_TIMELINE_MODULE                 TypeName = "TimelineTimelineModule"
	TWEET_TOMBSTONE                          TypeName = "TweetTombstone"
	TIMELINE_PROMPT                          TypeName = "TimelinePrompt"
	TIMELINE_MESSAGE_PROMPT                  TypeName = "TimelineMessagePrompt"
	TIMELINE_COMMUNITY                       TypeName = "TimelineCommunity"
	TWEET_UNAVAILABLE                        TypeName = "TweetUnavailable"
	TWEET                                    TypeName = "Tweet"
	USER                                     TypeName = "User"
	USER_UNAVAILABLE                         TypeName = "UserUnavailable"
	COMMUNITY                                TypeName = "Community"
	COMMUNITY_DELETE_ACTION_UNAVAILABLE      TypeName = "CommunityDeleteActionUnavailable"
	COMMUNITY_JOIN_ACTION                    TypeName = "CommunityJoinAction"
	COMMUNITY_LEAVE_ACTION_UNAVAILABLE       TypeName = "CommunityLeaveActionUnavailable"
	COMMUNITY_TWEET_PIN_ACTION_UNAVAILABLE   TypeName = "CommunityTweetPinActionUnavailable"
	COMMUNITY_TWEET_UNPIN_ACTION_UNAVAILABLE TypeName = "CommunityTweetUnpinActionUnavailable"
	COMMUNITY_INVITES_UNAVAILABLE            TypeName = "CommunityInvitesUnavailable"
	COMMUNITY_JOIN_REQUESTS_UNAVAILABLE      TypeName = "CommunityJoinRequestsUnavailable"
	API_IMAGE                                TypeName = "ApiImage"
)

// All allowed values of TypeName enum
var AllowedTypeNameEnumValues = []TypeName{
	"TimelineTweet",
	"TimelineTimelineItem",
	"TimelineUser",
	"TimelineTimelineCursor",
	"TweetWithVisibilityResults",
	"ContextualTweetInterstitial",
	"TimelineTimelineModule",
	"TweetTombstone",
	"TimelinePrompt",
	"TimelineMessagePrompt",
	"TimelineCommunity",
	"TweetUnavailable",
	"Tweet",
	"User",
	"UserUnavailable",
	"Community",
	"CommunityDeleteActionUnavailable",
	"CommunityJoinAction",
	"CommunityLeaveActionUnavailable",
	"CommunityTweetPinActionUnavailable",
	"CommunityTweetUnpinActionUnavailable",
	"CommunityInvitesUnavailable",
	"CommunityJoinRequestsUnavailable",
	"ApiImage",
}

func (v *TypeName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeName(value)
	for _, existing := range AllowedTypeNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeName", value)
}

// NewTypeNameFromValue returns a pointer to a valid TypeName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeNameFromValue(v string) (*TypeName, error) {
	ev := TypeName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypeName: valid values are %v", v, AllowedTypeNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeName) IsValid() bool {
	for _, existing := range AllowedTypeNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TypeName value
func (v TypeName) Ptr() *TypeName {
	return &v
}

type NullableTypeName struct {
	value *TypeName
	isSet bool
}

func (v NullableTypeName) Get() *TypeName {
	return v.value
}

func (v *NullableTypeName) Set(val *TypeName) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeName) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeName(val *TypeName) *NullableTypeName {
	return &NullableTypeName{value: val, isSet: true}
}

func (v NullableTypeName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
