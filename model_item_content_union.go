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
	"gopkg.in/validator.v2"
)

// ItemContentUnion - struct for ItemContentUnion
type ItemContentUnion struct {
	TimelineCommunity      *TimelineCommunity
	TimelineMessagePrompt  *TimelineMessagePrompt
	TimelinePrompt         *TimelinePrompt
	TimelineTimelineCursor *TimelineTimelineCursor
	TimelineTweet          *TimelineTweet
	TimelineUser           *TimelineUser
}

// TimelineCommunityAsItemContentUnion is a convenience function that returns TimelineCommunity wrapped in ItemContentUnion
func TimelineCommunityAsItemContentUnion(v *TimelineCommunity) ItemContentUnion {
	return ItemContentUnion{
		TimelineCommunity: v,
	}
}

// TimelineMessagePromptAsItemContentUnion is a convenience function that returns TimelineMessagePrompt wrapped in ItemContentUnion
func TimelineMessagePromptAsItemContentUnion(v *TimelineMessagePrompt) ItemContentUnion {
	return ItemContentUnion{
		TimelineMessagePrompt: v,
	}
}

// TimelinePromptAsItemContentUnion is a convenience function that returns TimelinePrompt wrapped in ItemContentUnion
func TimelinePromptAsItemContentUnion(v *TimelinePrompt) ItemContentUnion {
	return ItemContentUnion{
		TimelinePrompt: v,
	}
}

// TimelineTimelineCursorAsItemContentUnion is a convenience function that returns TimelineTimelineCursor wrapped in ItemContentUnion
func TimelineTimelineCursorAsItemContentUnion(v *TimelineTimelineCursor) ItemContentUnion {
	return ItemContentUnion{
		TimelineTimelineCursor: v,
	}
}

// TimelineTweetAsItemContentUnion is a convenience function that returns TimelineTweet wrapped in ItemContentUnion
func TimelineTweetAsItemContentUnion(v *TimelineTweet) ItemContentUnion {
	return ItemContentUnion{
		TimelineTweet: v,
	}
}

// TimelineUserAsItemContentUnion is a convenience function that returns TimelineUser wrapped in ItemContentUnion
func TimelineUserAsItemContentUnion(v *TimelineUser) ItemContentUnion {
	return ItemContentUnion{
		TimelineUser: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ItemContentUnion) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TimelineCommunity
	err = newStrictDecoder(data).Decode(&dst.TimelineCommunity)
	if err == nil {
		jsonTimelineCommunity, _ := json.Marshal(dst.TimelineCommunity)
		if string(jsonTimelineCommunity) == "{}" { // empty struct
			dst.TimelineCommunity = nil
		} else {
			if err = validator.Validate(dst.TimelineCommunity); err != nil {
				dst.TimelineCommunity = nil
			} else {
				match++
			}
		}
	} else {
		dst.TimelineCommunity = nil
	}

	// try to unmarshal data into TimelineMessagePrompt
	err = newStrictDecoder(data).Decode(&dst.TimelineMessagePrompt)
	if err == nil {
		jsonTimelineMessagePrompt, _ := json.Marshal(dst.TimelineMessagePrompt)
		if string(jsonTimelineMessagePrompt) == "{}" { // empty struct
			dst.TimelineMessagePrompt = nil
		} else {
			if err = validator.Validate(dst.TimelineMessagePrompt); err != nil {
				dst.TimelineMessagePrompt = nil
			} else {
				match++
			}
		}
	} else {
		dst.TimelineMessagePrompt = nil
	}

	// try to unmarshal data into TimelinePrompt
	err = newStrictDecoder(data).Decode(&dst.TimelinePrompt)
	if err == nil {
		jsonTimelinePrompt, _ := json.Marshal(dst.TimelinePrompt)
		if string(jsonTimelinePrompt) == "{}" { // empty struct
			dst.TimelinePrompt = nil
		} else {
			if err = validator.Validate(dst.TimelinePrompt); err != nil {
				dst.TimelinePrompt = nil
			} else {
				match++
			}
		}
	} else {
		dst.TimelinePrompt = nil
	}

	// try to unmarshal data into TimelineTimelineCursor
	err = newStrictDecoder(data).Decode(&dst.TimelineTimelineCursor)
	if err == nil {
		jsonTimelineTimelineCursor, _ := json.Marshal(dst.TimelineTimelineCursor)
		if string(jsonTimelineTimelineCursor) == "{}" { // empty struct
			dst.TimelineTimelineCursor = nil
		} else {
			if err = validator.Validate(dst.TimelineTimelineCursor); err != nil {
				dst.TimelineTimelineCursor = nil
			} else {
				match++
			}
		}
	} else {
		dst.TimelineTimelineCursor = nil
	}

	// try to unmarshal data into TimelineTweet
	err = newStrictDecoder(data).Decode(&dst.TimelineTweet)
	if err == nil {
		jsonTimelineTweet, _ := json.Marshal(dst.TimelineTweet)
		if string(jsonTimelineTweet) == "{}" { // empty struct
			dst.TimelineTweet = nil
		} else {
			if err = validator.Validate(dst.TimelineTweet); err != nil {
				dst.TimelineTweet = nil
			} else {
				match++
			}
		}
	} else {
		dst.TimelineTweet = nil
	}

	// try to unmarshal data into TimelineUser
	err = newStrictDecoder(data).Decode(&dst.TimelineUser)
	if err == nil {
		jsonTimelineUser, _ := json.Marshal(dst.TimelineUser)
		if string(jsonTimelineUser) == "{}" { // empty struct
			dst.TimelineUser = nil
		} else {
			if err = validator.Validate(dst.TimelineUser); err != nil {
				dst.TimelineUser = nil
			} else {
				match++
			}
		}
	} else {
		dst.TimelineUser = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TimelineCommunity = nil
		dst.TimelineMessagePrompt = nil
		dst.TimelinePrompt = nil
		dst.TimelineTimelineCursor = nil
		dst.TimelineTweet = nil
		dst.TimelineUser = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ItemContentUnion)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ItemContentUnion)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ItemContentUnion) MarshalJSON() ([]byte, error) {
	if src.TimelineCommunity != nil {
		return json.Marshal(&src.TimelineCommunity)
	}

	if src.TimelineMessagePrompt != nil {
		return json.Marshal(&src.TimelineMessagePrompt)
	}

	if src.TimelinePrompt != nil {
		return json.Marshal(&src.TimelinePrompt)
	}

	if src.TimelineTimelineCursor != nil {
		return json.Marshal(&src.TimelineTimelineCursor)
	}

	if src.TimelineTweet != nil {
		return json.Marshal(&src.TimelineTweet)
	}

	if src.TimelineUser != nil {
		return json.Marshal(&src.TimelineUser)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ItemContentUnion) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.TimelineCommunity != nil {
		return obj.TimelineCommunity
	}

	if obj.TimelineMessagePrompt != nil {
		return obj.TimelineMessagePrompt
	}

	if obj.TimelinePrompt != nil {
		return obj.TimelinePrompt
	}

	if obj.TimelineTimelineCursor != nil {
		return obj.TimelineTimelineCursor
	}

	if obj.TimelineTweet != nil {
		return obj.TimelineTweet
	}

	if obj.TimelineUser != nil {
		return obj.TimelineUser
	}

	// all schemas are nil
	return nil
}

type NullableItemContentUnion struct {
	value *ItemContentUnion
	isSet bool
}

func (v NullableItemContentUnion) Get() *ItemContentUnion {
	return v.value
}

func (v *NullableItemContentUnion) Set(val *ItemContentUnion) {
	v.value = val
	v.isSet = true
}

func (v NullableItemContentUnion) IsSet() bool {
	return v.isSet
}

func (v *NullableItemContentUnion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemContentUnion(val *ItemContentUnion) *NullableItemContentUnion {
	return &NullableItemContentUnion{value: val, isSet: true}
}

func (v NullableItemContentUnion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemContentUnion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
