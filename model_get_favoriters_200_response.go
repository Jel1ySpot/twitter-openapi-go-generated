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

// GetFavoriters200Response - struct for GetFavoriters200Response
type GetFavoriters200Response struct {
	Errors                  *Errors
	TweetFavoritersResponse *TweetFavoritersResponse
}

// ErrorsAsGetFavoriters200Response is a convenience function that returns Errors wrapped in GetFavoriters200Response
func ErrorsAsGetFavoriters200Response(v *Errors) GetFavoriters200Response {
	return GetFavoriters200Response{
		Errors: v,
	}
}

// TweetFavoritersResponseAsGetFavoriters200Response is a convenience function that returns TweetFavoritersResponse wrapped in GetFavoriters200Response
func TweetFavoritersResponseAsGetFavoriters200Response(v *TweetFavoritersResponse) GetFavoriters200Response {
	return GetFavoriters200Response{
		TweetFavoritersResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetFavoriters200Response) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into TweetFavoritersResponse
	err = newStrictDecoder(data).Decode(&dst.TweetFavoritersResponse)
	if err == nil {
		jsonTweetFavoritersResponse, _ := json.Marshal(dst.TweetFavoritersResponse)
		if string(jsonTweetFavoritersResponse) == "{}" { // empty struct
			dst.TweetFavoritersResponse = nil
		} else {
			if err = validator.Validate(dst.TweetFavoritersResponse); err != nil {
				dst.TweetFavoritersResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.TweetFavoritersResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Errors = nil
		dst.TweetFavoritersResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetFavoriters200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetFavoriters200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetFavoriters200Response) MarshalJSON() ([]byte, error) {
	if src.Errors != nil {
		return json.Marshal(&src.Errors)
	}

	if src.TweetFavoritersResponse != nil {
		return json.Marshal(&src.TweetFavoritersResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetFavoriters200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Errors != nil {
		return obj.Errors
	}

	if obj.TweetFavoritersResponse != nil {
		return obj.TweetFavoritersResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetFavoriters200Response struct {
	value *GetFavoriters200Response
	isSet bool
}

func (v NullableGetFavoriters200Response) Get() *GetFavoriters200Response {
	return v.value
}

func (v *NullableGetFavoriters200Response) Set(val *GetFavoriters200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFavoriters200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFavoriters200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFavoriters200Response(val *GetFavoriters200Response) *NullableGetFavoriters200Response {
	return &NullableGetFavoriters200Response{value: val, isSet: true}
}

func (v NullableGetFavoriters200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFavoriters200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
