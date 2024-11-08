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

// GetHomeLatestTimeline200Response - struct for GetHomeLatestTimeline200Response
type GetHomeLatestTimeline200Response struct {
	Errors           *Errors
	TimelineResponse *TimelineResponse
}

// ErrorsAsGetHomeLatestTimeline200Response is a convenience function that returns Errors wrapped in GetHomeLatestTimeline200Response
func ErrorsAsGetHomeLatestTimeline200Response(v *Errors) GetHomeLatestTimeline200Response {
	return GetHomeLatestTimeline200Response{
		Errors: v,
	}
}

// TimelineResponseAsGetHomeLatestTimeline200Response is a convenience function that returns TimelineResponse wrapped in GetHomeLatestTimeline200Response
func TimelineResponseAsGetHomeLatestTimeline200Response(v *TimelineResponse) GetHomeLatestTimeline200Response {
	return GetHomeLatestTimeline200Response{
		TimelineResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetHomeLatestTimeline200Response) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into TimelineResponse
	err = newStrictDecoder(data).Decode(&dst.TimelineResponse)
	if err == nil {
		jsonTimelineResponse, _ := json.Marshal(dst.TimelineResponse)
		if string(jsonTimelineResponse) == "{}" { // empty struct
			dst.TimelineResponse = nil
		} else {
			if err = validator.Validate(dst.TimelineResponse); err != nil {
				dst.TimelineResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.TimelineResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Errors = nil
		dst.TimelineResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetHomeLatestTimeline200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetHomeLatestTimeline200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetHomeLatestTimeline200Response) MarshalJSON() ([]byte, error) {
	if src.Errors != nil {
		return json.Marshal(&src.Errors)
	}

	if src.TimelineResponse != nil {
		return json.Marshal(&src.TimelineResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetHomeLatestTimeline200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Errors != nil {
		return obj.Errors
	}

	if obj.TimelineResponse != nil {
		return obj.TimelineResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetHomeLatestTimeline200Response struct {
	value *GetHomeLatestTimeline200Response
	isSet bool
}

func (v NullableGetHomeLatestTimeline200Response) Get() *GetHomeLatestTimeline200Response {
	return v.value
}

func (v *NullableGetHomeLatestTimeline200Response) Set(val *GetHomeLatestTimeline200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHomeLatestTimeline200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHomeLatestTimeline200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHomeLatestTimeline200Response(val *GetHomeLatestTimeline200Response) *NullableGetHomeLatestTimeline200Response {
	return &NullableGetHomeLatestTimeline200Response{value: val, isSet: true}
}

func (v NullableGetHomeLatestTimeline200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHomeLatestTimeline200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
