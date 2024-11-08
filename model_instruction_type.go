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

// InstructionType the model 'InstructionType'
type InstructionType string

// List of InstructionType
const (
	TIMELINE_ADD_ENTRIES        InstructionType = "TimelineAddEntries"
	TIMELINE_ADD_TO_MODULE      InstructionType = "TimelineAddToModule"
	TIMELINE_CLEAR_CACHE        InstructionType = "TimelineClearCache"
	TIMELINE_PIN_ENTRY          InstructionType = "TimelinePinEntry"
	TIMELINE_REPLACE_ENTRY      InstructionType = "TimelineReplaceEntry"
	TIMELINE_SHOW_ALERT         InstructionType = "TimelineShowAlert"
	TIMELINE_TERMINATE_TIMELINE InstructionType = "TimelineTerminateTimeline"
	TIMELINE_SHOW_COVER         InstructionType = "TimelineShowCover"
)

// All allowed values of InstructionType enum
var AllowedInstructionTypeEnumValues = []InstructionType{
	"TimelineAddEntries",
	"TimelineAddToModule",
	"TimelineClearCache",
	"TimelinePinEntry",
	"TimelineReplaceEntry",
	"TimelineShowAlert",
	"TimelineTerminateTimeline",
	"TimelineShowCover",
}

func (v *InstructionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InstructionType(value)
	for _, existing := range AllowedInstructionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InstructionType", value)
}

// NewInstructionTypeFromValue returns a pointer to a valid InstructionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstructionTypeFromValue(v string) (*InstructionType, error) {
	ev := InstructionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstructionType: valid values are %v", v, AllowedInstructionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstructionType) IsValid() bool {
	for _, existing := range AllowedInstructionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InstructionType value
func (v InstructionType) Ptr() *InstructionType {
	return &v
}

type NullableInstructionType struct {
	value *InstructionType
	isSet bool
}

func (v NullableInstructionType) Get() *InstructionType {
	return v.value
}

func (v *NullableInstructionType) Set(val *InstructionType) {
	v.value = val
	v.isSet = true
}

func (v NullableInstructionType) IsSet() bool {
	return v.isSet
}

func (v *NullableInstructionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstructionType(val *InstructionType) *NullableInstructionType {
	return &NullableInstructionType{value: val, isSet: true}
}

func (v NullableInstructionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstructionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
