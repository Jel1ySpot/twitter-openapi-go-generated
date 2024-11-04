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

// PostDeleteRetweet200Response - struct for PostDeleteRetweet200Response
type PostDeleteRetweet200Response struct {
	DeleteRetweetResponse *DeleteRetweetResponse
	Errors                *Errors
}

// DeleteRetweetResponseAsPostDeleteRetweet200Response is a convenience function that returns DeleteRetweetResponse wrapped in PostDeleteRetweet200Response
func DeleteRetweetResponseAsPostDeleteRetweet200Response(v *DeleteRetweetResponse) PostDeleteRetweet200Response {
	return PostDeleteRetweet200Response{
		DeleteRetweetResponse: v,
	}
}

// ErrorsAsPostDeleteRetweet200Response is a convenience function that returns Errors wrapped in PostDeleteRetweet200Response
func ErrorsAsPostDeleteRetweet200Response(v *Errors) PostDeleteRetweet200Response {
	return PostDeleteRetweet200Response{
		Errors: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostDeleteRetweet200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeleteRetweetResponse
	err = newStrictDecoder(data).Decode(&dst.DeleteRetweetResponse)
	if err == nil {
		jsonDeleteRetweetResponse, _ := json.Marshal(dst.DeleteRetweetResponse)
		if string(jsonDeleteRetweetResponse) == "{}" { // empty struct
			dst.DeleteRetweetResponse = nil
		} else {
			if err = validator.Validate(dst.DeleteRetweetResponse); err != nil {
				dst.DeleteRetweetResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.DeleteRetweetResponse = nil
	}

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

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DeleteRetweetResponse = nil
		dst.Errors = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PostDeleteRetweet200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostDeleteRetweet200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostDeleteRetweet200Response) MarshalJSON() ([]byte, error) {
	if src.DeleteRetweetResponse != nil {
		return json.Marshal(&src.DeleteRetweetResponse)
	}

	if src.Errors != nil {
		return json.Marshal(&src.Errors)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostDeleteRetweet200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DeleteRetweetResponse != nil {
		return obj.DeleteRetweetResponse
	}

	if obj.Errors != nil {
		return obj.Errors
	}

	// all schemas are nil
	return nil
}

type NullablePostDeleteRetweet200Response struct {
	value *PostDeleteRetweet200Response
	isSet bool
}

func (v NullablePostDeleteRetweet200Response) Get() *PostDeleteRetweet200Response {
	return v.value
}

func (v *NullablePostDeleteRetweet200Response) Set(val *PostDeleteRetweet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostDeleteRetweet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostDeleteRetweet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostDeleteRetweet200Response(val *PostDeleteRetweet200Response) *NullablePostDeleteRetweet200Response {
	return &NullablePostDeleteRetweet200Response{value: val, isSet: true}
}

func (v NullablePostDeleteRetweet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostDeleteRetweet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
