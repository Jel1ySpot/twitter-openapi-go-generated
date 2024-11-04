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

// GetUserByRestId200Response - struct for GetUserByRestId200Response
type GetUserByRestId200Response struct {
	Errors *Errors
	UserResponse *UserResponse
}

// ErrorsAsGetUserByRestId200Response is a convenience function that returns Errors wrapped in GetUserByRestId200Response
func ErrorsAsGetUserByRestId200Response(v *Errors) GetUserByRestId200Response {
	return GetUserByRestId200Response{
		Errors: v,
	}
}

// UserResponseAsGetUserByRestId200Response is a convenience function that returns UserResponse wrapped in GetUserByRestId200Response
func UserResponseAsGetUserByRestId200Response(v *UserResponse) GetUserByRestId200Response {
	return GetUserByRestId200Response{
		UserResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetUserByRestId200Response) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into UserResponse
	err = newStrictDecoder(data).Decode(&dst.UserResponse)
	if err == nil {
		jsonUserResponse, _ := json.Marshal(dst.UserResponse)
		if string(jsonUserResponse) == "{}" { // empty struct
			dst.UserResponse = nil
		} else {
			if err = validator.Validate(dst.UserResponse); err != nil {
				dst.UserResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Errors = nil
		dst.UserResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetUserByRestId200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetUserByRestId200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetUserByRestId200Response) MarshalJSON() ([]byte, error) {
	if src.Errors != nil {
		return json.Marshal(&src.Errors)
	}

	if src.UserResponse != nil {
		return json.Marshal(&src.UserResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetUserByRestId200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Errors != nil {
		return obj.Errors
	}

	if obj.UserResponse != nil {
		return obj.UserResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetUserByRestId200Response struct {
	value *GetUserByRestId200Response
	isSet bool
}

func (v NullableGetUserByRestId200Response) Get() *GetUserByRestId200Response {
	return v.value
}

func (v *NullableGetUserByRestId200Response) Set(val *GetUserByRestId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserByRestId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserByRestId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserByRestId200Response(val *GetUserByRestId200Response) *NullableGetUserByRestId200Response {
	return &NullableGetUserByRestId200Response{value: val, isSet: true}
}

func (v NullableGetUserByRestId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserByRestId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

