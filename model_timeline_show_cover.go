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
	"bytes"
	"fmt"
)

// checks if the TimelineShowCover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineShowCover{}

// TimelineShowCover struct for TimelineShowCover
type TimelineShowCover struct {
	ClientEventInfo ClientEventInfo `json:"clientEventInfo"`
	Cover TimelineHalfCover `json:"cover"`
	Type InstructionType `json:"type"`
}

type _TimelineShowCover TimelineShowCover

// NewTimelineShowCover instantiates a new TimelineShowCover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineShowCover(clientEventInfo ClientEventInfo, cover TimelineHalfCover, type_ InstructionType) *TimelineShowCover {
	this := TimelineShowCover{}
	this.ClientEventInfo = clientEventInfo
	this.Cover = cover
	this.Type = type_
	return &this
}

// NewTimelineShowCoverWithDefaults instantiates a new TimelineShowCover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineShowCoverWithDefaults() *TimelineShowCover {
	this := TimelineShowCover{}
	return &this
}

// GetClientEventInfo returns the ClientEventInfo field value
func (o *TimelineShowCover) GetClientEventInfo() ClientEventInfo {
	if o == nil {
		var ret ClientEventInfo
		return ret
	}

	return o.ClientEventInfo
}

// GetClientEventInfoOk returns a tuple with the ClientEventInfo field value
// and a boolean to check if the value has been set.
func (o *TimelineShowCover) GetClientEventInfoOk() (*ClientEventInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientEventInfo, true
}

// SetClientEventInfo sets field value
func (o *TimelineShowCover) SetClientEventInfo(v ClientEventInfo) {
	o.ClientEventInfo = v
}

// GetCover returns the Cover field value
func (o *TimelineShowCover) GetCover() TimelineHalfCover {
	if o == nil {
		var ret TimelineHalfCover
		return ret
	}

	return o.Cover
}

// GetCoverOk returns a tuple with the Cover field value
// and a boolean to check if the value has been set.
func (o *TimelineShowCover) GetCoverOk() (*TimelineHalfCover, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cover, true
}

// SetCover sets field value
func (o *TimelineShowCover) SetCover(v TimelineHalfCover) {
	o.Cover = v
}

// GetType returns the Type field value
func (o *TimelineShowCover) GetType() InstructionType {
	if o == nil {
		var ret InstructionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimelineShowCover) GetTypeOk() (*InstructionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TimelineShowCover) SetType(v InstructionType) {
	o.Type = v
}

func (o TimelineShowCover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineShowCover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientEventInfo"] = o.ClientEventInfo
	toSerialize["cover"] = o.Cover
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *TimelineShowCover) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientEventInfo",
		"cover",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTimelineShowCover := _TimelineShowCover{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimelineShowCover)

	if err != nil {
		return err
	}

	*o = TimelineShowCover(varTimelineShowCover)

	return err
}

type NullableTimelineShowCover struct {
	value *TimelineShowCover
	isSet bool
}

func (v NullableTimelineShowCover) Get() *TimelineShowCover {
	return v.value
}

func (v *NullableTimelineShowCover) Set(val *TimelineShowCover) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineShowCover) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineShowCover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineShowCover(val *TimelineShowCover) *NullableTimelineShowCover {
	return &NullableTimelineShowCover{value: val, isSet: true}
}

func (v NullableTimelineShowCover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineShowCover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

