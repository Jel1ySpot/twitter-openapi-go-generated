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

// checks if the ListTweetsTimeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTweetsTimeline{}

// ListTweetsTimeline struct for ListTweetsTimeline
type ListTweetsTimeline struct {
	Timeline *Timeline `json:"timeline,omitempty"`
}

// NewListTweetsTimeline instantiates a new ListTweetsTimeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTweetsTimeline() *ListTweetsTimeline {
	this := ListTweetsTimeline{}
	return &this
}

// NewListTweetsTimelineWithDefaults instantiates a new ListTweetsTimeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTweetsTimelineWithDefaults() *ListTweetsTimeline {
	this := ListTweetsTimeline{}
	return &this
}

// GetTimeline returns the Timeline field value if set, zero value otherwise.
func (o *ListTweetsTimeline) GetTimeline() Timeline {
	if o == nil || IsNil(o.Timeline) {
		var ret Timeline
		return ret
	}
	return *o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTweetsTimeline) GetTimelineOk() (*Timeline, bool) {
	if o == nil || IsNil(o.Timeline) {
		return nil, false
	}
	return o.Timeline, true
}

// HasTimeline returns a boolean if a field has been set.
func (o *ListTweetsTimeline) HasTimeline() bool {
	if o != nil && !IsNil(o.Timeline) {
		return true
	}

	return false
}

// SetTimeline gets a reference to the given Timeline and assigns it to the Timeline field.
func (o *ListTweetsTimeline) SetTimeline(v Timeline) {
	o.Timeline = &v
}

func (o ListTweetsTimeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTweetsTimeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timeline) {
		toSerialize["timeline"] = o.Timeline
	}
	return toSerialize, nil
}

type NullableListTweetsTimeline struct {
	value *ListTweetsTimeline
	isSet bool
}

func (v NullableListTweetsTimeline) Get() *ListTweetsTimeline {
	return v.value
}

func (v *NullableListTweetsTimeline) Set(val *ListTweetsTimeline) {
	v.value = val
	v.isSet = true
}

func (v NullableListTweetsTimeline) IsSet() bool {
	return v.isSet
}

func (v *NullableListTweetsTimeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTweetsTimeline(val *ListTweetsTimeline) *NullableListTweetsTimeline {
	return &NullableListTweetsTimeline{value: val, isSet: true}
}

func (v NullableListTweetsTimeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTweetsTimeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


