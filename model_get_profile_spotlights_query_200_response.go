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

// GetProfileSpotlightsQuery200Response - struct for GetProfileSpotlightsQuery200Response
type GetProfileSpotlightsQuery200Response struct {
	Errors          *Errors
	ProfileResponse *ProfileResponse
}

// ErrorsAsGetProfileSpotlightsQuery200Response is a convenience function that returns Errors wrapped in GetProfileSpotlightsQuery200Response
func ErrorsAsGetProfileSpotlightsQuery200Response(v *Errors) GetProfileSpotlightsQuery200Response {
	return GetProfileSpotlightsQuery200Response{
		Errors: v,
	}
}

// ProfileResponseAsGetProfileSpotlightsQuery200Response is a convenience function that returns ProfileResponse wrapped in GetProfileSpotlightsQuery200Response
func ProfileResponseAsGetProfileSpotlightsQuery200Response(v *ProfileResponse) GetProfileSpotlightsQuery200Response {
	return GetProfileSpotlightsQuery200Response{
		ProfileResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetProfileSpotlightsQuery200Response) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into ProfileResponse
	err = newStrictDecoder(data).Decode(&dst.ProfileResponse)
	if err == nil {
		jsonProfileResponse, _ := json.Marshal(dst.ProfileResponse)
		if string(jsonProfileResponse) == "{}" { // empty struct
			dst.ProfileResponse = nil
		} else {
			if err = validator.Validate(dst.ProfileResponse); err != nil {
				dst.ProfileResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ProfileResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Errors = nil
		dst.ProfileResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetProfileSpotlightsQuery200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetProfileSpotlightsQuery200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetProfileSpotlightsQuery200Response) MarshalJSON() ([]byte, error) {
	if src.Errors != nil {
		return json.Marshal(&src.Errors)
	}

	if src.ProfileResponse != nil {
		return json.Marshal(&src.ProfileResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetProfileSpotlightsQuery200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Errors != nil {
		return obj.Errors
	}

	if obj.ProfileResponse != nil {
		return obj.ProfileResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetProfileSpotlightsQuery200Response struct {
	value *GetProfileSpotlightsQuery200Response
	isSet bool
}

func (v NullableGetProfileSpotlightsQuery200Response) Get() *GetProfileSpotlightsQuery200Response {
	return v.value
}

func (v *NullableGetProfileSpotlightsQuery200Response) Set(val *GetProfileSpotlightsQuery200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProfileSpotlightsQuery200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProfileSpotlightsQuery200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProfileSpotlightsQuery200Response(val *GetProfileSpotlightsQuery200Response) *NullableGetProfileSpotlightsQuery200Response {
	return &NullableGetProfileSpotlightsQuery200Response{value: val, isSet: true}
}

func (v NullableGetProfileSpotlightsQuery200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProfileSpotlightsQuery200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
