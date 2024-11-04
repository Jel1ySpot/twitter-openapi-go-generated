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

// GetUsersByRestIds200Response - struct for GetUsersByRestIds200Response
type GetUsersByRestIds200Response struct {
	Errors *Errors
	UsersResponse *UsersResponse
}

// ErrorsAsGetUsersByRestIds200Response is a convenience function that returns Errors wrapped in GetUsersByRestIds200Response
func ErrorsAsGetUsersByRestIds200Response(v *Errors) GetUsersByRestIds200Response {
	return GetUsersByRestIds200Response{
		Errors: v,
	}
}

// UsersResponseAsGetUsersByRestIds200Response is a convenience function that returns UsersResponse wrapped in GetUsersByRestIds200Response
func UsersResponseAsGetUsersByRestIds200Response(v *UsersResponse) GetUsersByRestIds200Response {
	return GetUsersByRestIds200Response{
		UsersResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetUsersByRestIds200Response) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into UsersResponse
	err = newStrictDecoder(data).Decode(&dst.UsersResponse)
	if err == nil {
		jsonUsersResponse, _ := json.Marshal(dst.UsersResponse)
		if string(jsonUsersResponse) == "{}" { // empty struct
			dst.UsersResponse = nil
		} else {
			if err = validator.Validate(dst.UsersResponse); err != nil {
				dst.UsersResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.UsersResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Errors = nil
		dst.UsersResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetUsersByRestIds200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetUsersByRestIds200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetUsersByRestIds200Response) MarshalJSON() ([]byte, error) {
	if src.Errors != nil {
		return json.Marshal(&src.Errors)
	}

	if src.UsersResponse != nil {
		return json.Marshal(&src.UsersResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetUsersByRestIds200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Errors != nil {
		return obj.Errors
	}

	if obj.UsersResponse != nil {
		return obj.UsersResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetUsersByRestIds200Response struct {
	value *GetUsersByRestIds200Response
	isSet bool
}

func (v NullableGetUsersByRestIds200Response) Get() *GetUsersByRestIds200Response {
	return v.value
}

func (v *NullableGetUsersByRestIds200Response) Set(val *GetUsersByRestIds200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsersByRestIds200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsersByRestIds200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsersByRestIds200Response(val *GetUsersByRestIds200Response) *NullableGetUsersByRestIds200Response {
	return &NullableGetUsersByRestIds200Response{value: val, isSet: true}
}

func (v NullableGetUsersByRestIds200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsersByRestIds200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

