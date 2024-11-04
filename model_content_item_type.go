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

// ContentItemType the model 'ContentItemType'
type ContentItemType string

// List of ContentItemType
const (
	TIMELINE_TWEET ContentItemType = "TimelineTweet"
	TIMELINE_TIMELINE_CURSOR ContentItemType = "TimelineTimelineCursor"
	TIMELINE_USER ContentItemType = "TimelineUser"
	TIMELINE_PROMPT ContentItemType = "TimelinePrompt"
	TIMELINE_MESSAGE_PROMPT ContentItemType = "TimelineMessagePrompt"
	TIMELINE_COMMUNITY ContentItemType = "TimelineCommunity"
)

// All allowed values of ContentItemType enum
var AllowedContentItemTypeEnumValues = []ContentItemType{
	"TimelineTweet",
	"TimelineTimelineCursor",
	"TimelineUser",
	"TimelinePrompt",
	"TimelineMessagePrompt",
	"TimelineCommunity",
}

func (v *ContentItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContentItemType(value)
	for _, existing := range AllowedContentItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContentItemType", value)
}

// NewContentItemTypeFromValue returns a pointer to a valid ContentItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContentItemTypeFromValue(v string) (*ContentItemType, error) {
	ev := ContentItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContentItemType: valid values are %v", v, AllowedContentItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContentItemType) IsValid() bool {
	for _, existing := range AllowedContentItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContentItemType value
func (v ContentItemType) Ptr() *ContentItemType {
	return &v
}

type NullableContentItemType struct {
	value *ContentItemType
	isSet bool
}

func (v NullableContentItemType) Get() *ContentItemType {
	return v.value
}

func (v *NullableContentItemType) Set(val *ContentItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableContentItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableContentItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentItemType(val *ContentItemType) *NullableContentItemType {
	return &NullableContentItemType{value: val, isSet: true}
}

func (v NullableContentItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
