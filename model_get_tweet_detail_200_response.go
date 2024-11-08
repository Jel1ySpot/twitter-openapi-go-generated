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

// GetTweetDetail200Response - struct for GetTweetDetail200Response
type GetTweetDetail200Response struct {
	Errors              *Errors
	TweetDetailResponse *TweetDetailResponse
}

// ErrorsAsGetTweetDetail200Response is a convenience function that returns Errors wrapped in GetTweetDetail200Response
func ErrorsAsGetTweetDetail200Response(v *Errors) GetTweetDetail200Response {
	return GetTweetDetail200Response{
		Errors: v,
	}
}

// TweetDetailResponseAsGetTweetDetail200Response is a convenience function that returns TweetDetailResponse wrapped in GetTweetDetail200Response
func TweetDetailResponseAsGetTweetDetail200Response(v *TweetDetailResponse) GetTweetDetail200Response {
	return GetTweetDetail200Response{
		TweetDetailResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetTweetDetail200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Errors
	err = newStrictDecoder(data).Decode(&dst.Errors)
	if err == nil {
		jsonErrors, _ := json.Marshal(dst.Errors)
		if string(jsonErrors) == "{}" { // empty struct
			dst.Errors = nil
		} else {
			if err = validator.Validate(dst.Errors); err != nil {
				dst.Errors = nil
			} else {
				match++
			}
		}
	} else {
		dst.Errors = nil
	}

	// try to unmarshal data into TweetDetailResponse
	err = newStrictDecoder(data).Decode(&dst.TweetDetailResponse)
	if err == nil {
		jsonTweetDetailResponse, _ := json.Marshal(dst.TweetDetailResponse)
		if string(jsonTweetDetailResponse) == "{}" { // empty struct
			dst.TweetDetailResponse = nil
		} else {
			if err = validator.Validate(dst.TweetDetailResponse); err != nil {
				dst.TweetDetailResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.TweetDetailResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Errors = nil
		dst.TweetDetailResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetTweetDetail200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetTweetDetail200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetTweetDetail200Response) MarshalJSON() ([]byte, error) {
	if src.Errors != nil {
		return json.Marshal(&src.Errors)
	}

	if src.TweetDetailResponse != nil {
		return json.Marshal(&src.TweetDetailResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetTweetDetail200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Errors != nil {
		return obj.Errors
	}

	if obj.TweetDetailResponse != nil {
		return obj.TweetDetailResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetTweetDetail200Response struct {
	value *GetTweetDetail200Response
	isSet bool
}

func (v NullableGetTweetDetail200Response) Get() *GetTweetDetail200Response {
	return v.value
}

func (v *NullableGetTweetDetail200Response) Set(val *GetTweetDetail200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTweetDetail200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTweetDetail200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTweetDetail200Response(val *GetTweetDetail200Response) *NullableGetTweetDetail200Response {
	return &NullableGetTweetDetail200Response{value: val, isSet: true}
}

func (v NullableGetTweetDetail200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTweetDetail200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
