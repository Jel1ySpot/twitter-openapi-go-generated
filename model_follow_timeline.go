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

// checks if the FollowTimeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowTimeline{}

// FollowTimeline struct for FollowTimeline
type FollowTimeline struct {
	Timeline Timeline `json:"timeline"`
}

type _FollowTimeline FollowTimeline

// NewFollowTimeline instantiates a new FollowTimeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowTimeline(timeline Timeline) *FollowTimeline {
	this := FollowTimeline{}
	this.Timeline = timeline
	return &this
}

// NewFollowTimelineWithDefaults instantiates a new FollowTimeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowTimelineWithDefaults() *FollowTimeline {
	this := FollowTimeline{}
	return &this
}

// GetTimeline returns the Timeline field value
func (o *FollowTimeline) GetTimeline() Timeline {
	if o == nil {
		var ret Timeline
		return ret
	}

	return o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value
// and a boolean to check if the value has been set.
func (o *FollowTimeline) GetTimelineOk() (*Timeline, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeline, true
}

// SetTimeline sets field value
func (o *FollowTimeline) SetTimeline(v Timeline) {
	o.Timeline = v
}

func (o FollowTimeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowTimeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeline"] = o.Timeline
	return toSerialize, nil
}

func (o *FollowTimeline) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timeline",
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

	varFollowTimeline := _FollowTimeline{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFollowTimeline)

	if err != nil {
		return err
	}

	*o = FollowTimeline(varFollowTimeline)

	return err
}

type NullableFollowTimeline struct {
	value *FollowTimeline
	isSet bool
}

func (v NullableFollowTimeline) Get() *FollowTimeline {
	return v.value
}

func (v *NullableFollowTimeline) Set(val *FollowTimeline) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowTimeline) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowTimeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowTimeline(val *FollowTimeline) *NullableFollowTimeline {
	return &NullableFollowTimeline{value: val, isSet: true}
}

func (v NullableFollowTimeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowTimeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

