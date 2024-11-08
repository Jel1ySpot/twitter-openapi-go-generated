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
)

// checks if the TimelineV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineV2{}

// TimelineV2 struct for TimelineV2
type TimelineV2 struct {
	Timeline *Timeline `json:"timeline,omitempty"`
}

// NewTimelineV2 instantiates a new TimelineV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineV2() *TimelineV2 {
	this := TimelineV2{}
	return &this
}

// NewTimelineV2WithDefaults instantiates a new TimelineV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineV2WithDefaults() *TimelineV2 {
	this := TimelineV2{}
	return &this
}

// GetTimeline returns the Timeline field value if set, zero value otherwise.
func (o *TimelineV2) GetTimeline() Timeline {
	if o == nil || IsNil(o.Timeline) {
		var ret Timeline
		return ret
	}
	return *o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineV2) GetTimelineOk() (*Timeline, bool) {
	if o == nil || IsNil(o.Timeline) {
		return nil, false
	}
	return o.Timeline, true
}

// HasTimeline returns a boolean if a field has been set.
func (o *TimelineV2) HasTimeline() bool {
	if o != nil && !IsNil(o.Timeline) {
		return true
	}

	return false
}

// SetTimeline gets a reference to the given Timeline and assigns it to the Timeline field.
func (o *TimelineV2) SetTimeline(v Timeline) {
	o.Timeline = &v
}

func (o TimelineV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timeline) {
		toSerialize["timeline"] = o.Timeline
	}
	return toSerialize, nil
}

type NullableTimelineV2 struct {
	value *TimelineV2
	isSet bool
}

func (v NullableTimelineV2) Get() *TimelineV2 {
	return v.value
}

func (v *NullableTimelineV2) Set(val *TimelineV2) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineV2) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineV2(val *TimelineV2) *NullableTimelineV2 {
	return &NullableTimelineV2{value: val, isSet: true}
}

func (v NullableTimelineV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
