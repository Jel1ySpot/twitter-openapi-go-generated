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
	"gopkg.in/validator.v2"
	"fmt"
)

// GetTweetResultByRestId200Response - struct for GetTweetResultByRestId200Response
type GetTweetResultByRestId200Response struct {
	Errors *Errors
	TweetResultByRestIdResponse *TweetResultByRestIdResponse
}

// ErrorsAsGetTweetResultByRestId200Response is a convenience function that returns Errors wrapped in GetTweetResultByRestId200Response
func ErrorsAsGetTweetResultByRestId200Response(v *Errors) GetTweetResultByRestId200Response {
	return GetTweetResultByRestId200Response{
		Errors: v,
	}
}

// TweetResultByRestIdResponseAsGetTweetResultByRestId200Response is a convenience function that returns TweetResultByRestIdResponse wrapped in GetTweetResultByRestId200Response
func TweetResultByRestIdResponseAsGetTweetResultByRestId200Response(v *TweetResultByRestIdResponse) GetTweetResultByRestId200Response {
	return GetTweetResultByRestId200Response{
		TweetResultByRestIdResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetTweetResultByRestId200Response) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into TweetResultByRestIdResponse
	err = newStrictDecoder(data).Decode(&dst.TweetResultByRestIdResponse)
	if err == nil {
		jsonTweetResultByRestIdResponse, _ := json.Marshal(dst.TweetResultByRestIdResponse)
		if string(jsonTweetResultByRestIdResponse) == "{}" { // empty struct
			dst.TweetResultByRestIdResponse = nil
		} else {
			if err = validator.Validate(dst.TweetResultByRestIdResponse); err != nil {
				dst.TweetResultByRestIdResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.TweetResultByRestIdResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Errors = nil
		dst.TweetResultByRestIdResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetTweetResultByRestId200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetTweetResultByRestId200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetTweetResultByRestId200Response) MarshalJSON() ([]byte, error) {
	if src.Errors != nil {
		return json.Marshal(&src.Errors)
	}

	if src.TweetResultByRestIdResponse != nil {
		return json.Marshal(&src.TweetResultByRestIdResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetTweetResultByRestId200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Errors != nil {
		return obj.Errors
	}

	if obj.TweetResultByRestIdResponse != nil {
		return obj.TweetResultByRestIdResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetTweetResultByRestId200Response struct {
	value *GetTweetResultByRestId200Response
	isSet bool
}

func (v NullableGetTweetResultByRestId200Response) Get() *GetTweetResultByRestId200Response {
	return v.value
}

func (v *NullableGetTweetResultByRestId200Response) Set(val *GetTweetResultByRestId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTweetResultByRestId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTweetResultByRestId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTweetResultByRestId200Response(val *GetTweetResultByRestId200Response) *NullableGetTweetResultByRestId200Response {
	return &NullableGetTweetResultByRestId200Response{value: val, isSet: true}
}

func (v NullableGetTweetResultByRestId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTweetResultByRestId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

