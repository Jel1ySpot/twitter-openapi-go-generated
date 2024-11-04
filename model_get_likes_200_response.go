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

// GetLikes200Response - struct for GetLikes200Response
type GetLikes200Response struct {
	Errors *Errors
	UserTweetsResponse *UserTweetsResponse
}

// ErrorsAsGetLikes200Response is a convenience function that returns Errors wrapped in GetLikes200Response
func ErrorsAsGetLikes200Response(v *Errors) GetLikes200Response {
	return GetLikes200Response{
		Errors: v,
	}
}

// UserTweetsResponseAsGetLikes200Response is a convenience function that returns UserTweetsResponse wrapped in GetLikes200Response
func UserTweetsResponseAsGetLikes200Response(v *UserTweetsResponse) GetLikes200Response {
	return GetLikes200Response{
		UserTweetsResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetLikes200Response) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into UserTweetsResponse
	err = newStrictDecoder(data).Decode(&dst.UserTweetsResponse)
	if err == nil {
		jsonUserTweetsResponse, _ := json.Marshal(dst.UserTweetsResponse)
		if string(jsonUserTweetsResponse) == "{}" { // empty struct
			dst.UserTweetsResponse = nil
		} else {
			if err = validator.Validate(dst.UserTweetsResponse); err != nil {
				dst.UserTweetsResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserTweetsResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Errors = nil
		dst.UserTweetsResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetLikes200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetLikes200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetLikes200Response) MarshalJSON() ([]byte, error) {
	if src.Errors != nil {
		return json.Marshal(&src.Errors)
	}

	if src.UserTweetsResponse != nil {
		return json.Marshal(&src.UserTweetsResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetLikes200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Errors != nil {
		return obj.Errors
	}

	if obj.UserTweetsResponse != nil {
		return obj.UserTweetsResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetLikes200Response struct {
	value *GetLikes200Response
	isSet bool
}

func (v NullableGetLikes200Response) Get() *GetLikes200Response {
	return v.value
}

func (v *NullableGetLikes200Response) Set(val *GetLikes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLikes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLikes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLikes200Response(val *GetLikes200Response) *NullableGetLikes200Response {
	return &NullableGetLikes200Response{value: val, isSet: true}
}

func (v NullableGetLikes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLikes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

