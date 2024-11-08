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

// TweetUnion - struct for TweetUnion
type TweetUnion struct {
	Tweet                      *Tweet
	TweetTombstone             *TweetTombstone
	TweetUnavailable           *TweetUnavailable
	TweetWithVisibilityResults *TweetWithVisibilityResults
}

// TweetAsTweetUnion is a convenience function that returns Tweet wrapped in TweetUnion
func TweetAsTweetUnion(v *Tweet) TweetUnion {
	return TweetUnion{
		Tweet: v,
	}
}

// TweetTombstoneAsTweetUnion is a convenience function that returns TweetTombstone wrapped in TweetUnion
func TweetTombstoneAsTweetUnion(v *TweetTombstone) TweetUnion {
	return TweetUnion{
		TweetTombstone: v,
	}
}

// TweetUnavailableAsTweetUnion is a convenience function that returns TweetUnavailable wrapped in TweetUnion
func TweetUnavailableAsTweetUnion(v *TweetUnavailable) TweetUnion {
	return TweetUnion{
		TweetUnavailable: v,
	}
}

// TweetWithVisibilityResultsAsTweetUnion is a convenience function that returns TweetWithVisibilityResults wrapped in TweetUnion
func TweetWithVisibilityResultsAsTweetUnion(v *TweetWithVisibilityResults) TweetUnion {
	return TweetUnion{
		TweetWithVisibilityResults: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TweetUnion) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Tweet
	err = newStrictDecoder(data).Decode(&dst.Tweet)
	if err == nil {
		jsonTweet, _ := json.Marshal(dst.Tweet)
		if string(jsonTweet) == "{}" { // empty struct
			dst.Tweet = nil
		} else {
			if err = validator.Validate(dst.Tweet); err != nil {
				dst.Tweet = nil
			} else {
				match++
			}
		}
	} else {
		dst.Tweet = nil
	}

	// try to unmarshal data into TweetTombstone
	err = newStrictDecoder(data).Decode(&dst.TweetTombstone)
	if err == nil {
		jsonTweetTombstone, _ := json.Marshal(dst.TweetTombstone)
		if string(jsonTweetTombstone) == "{}" { // empty struct
			dst.TweetTombstone = nil
		} else {
			if err = validator.Validate(dst.TweetTombstone); err != nil {
				dst.TweetTombstone = nil
			} else {
				match++
			}
		}
	} else {
		dst.TweetTombstone = nil
	}

	// try to unmarshal data into TweetUnavailable
	err = newStrictDecoder(data).Decode(&dst.TweetUnavailable)
	if err == nil {
		jsonTweetUnavailable, _ := json.Marshal(dst.TweetUnavailable)
		if string(jsonTweetUnavailable) == "{}" { // empty struct
			dst.TweetUnavailable = nil
		} else {
			if err = validator.Validate(dst.TweetUnavailable); err != nil {
				dst.TweetUnavailable = nil
			} else {
				match++
			}
		}
	} else {
		dst.TweetUnavailable = nil
	}

	// try to unmarshal data into TweetWithVisibilityResults
	err = newStrictDecoder(data).Decode(&dst.TweetWithVisibilityResults)
	if err == nil {
		jsonTweetWithVisibilityResults, _ := json.Marshal(dst.TweetWithVisibilityResults)
		if string(jsonTweetWithVisibilityResults) == "{}" { // empty struct
			dst.TweetWithVisibilityResults = nil
		} else {
			if err = validator.Validate(dst.TweetWithVisibilityResults); err != nil {
				dst.TweetWithVisibilityResults = nil
			} else {
				match++
			}
		}
	} else {
		dst.TweetWithVisibilityResults = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Tweet = nil
		dst.TweetTombstone = nil
		dst.TweetUnavailable = nil
		dst.TweetWithVisibilityResults = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TweetUnion)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TweetUnion)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TweetUnion) MarshalJSON() ([]byte, error) {
	if src.Tweet != nil {
		return json.Marshal(&src.Tweet)
	}

	if src.TweetTombstone != nil {
		return json.Marshal(&src.TweetTombstone)
	}

	if src.TweetUnavailable != nil {
		return json.Marshal(&src.TweetUnavailable)
	}

	if src.TweetWithVisibilityResults != nil {
		return json.Marshal(&src.TweetWithVisibilityResults)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TweetUnion) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Tweet != nil {
		return obj.Tweet
	}

	if obj.TweetTombstone != nil {
		return obj.TweetTombstone
	}

	if obj.TweetUnavailable != nil {
		return obj.TweetUnavailable
	}

	if obj.TweetWithVisibilityResults != nil {
		return obj.TweetWithVisibilityResults
	}

	// all schemas are nil
	return nil
}

type NullableTweetUnion struct {
	value *TweetUnion
	isSet bool
}

func (v NullableTweetUnion) Get() *TweetUnion {
	return v.value
}

func (v *NullableTweetUnion) Set(val *TweetUnion) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetUnion) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetUnion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetUnion(val *TweetUnion) *NullableTweetUnion {
	return &NullableTweetUnion{value: val, isSet: true}
}

func (v NullableTweetUnion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetUnion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
