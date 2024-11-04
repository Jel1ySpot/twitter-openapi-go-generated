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

// checks if the UserHighlightsTweetsTimeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserHighlightsTweetsTimeline{}

// UserHighlightsTweetsTimeline struct for UserHighlightsTweetsTimeline
type UserHighlightsTweetsTimeline struct {
	Timeline Timeline `json:"timeline"`
}

type _UserHighlightsTweetsTimeline UserHighlightsTweetsTimeline

// NewUserHighlightsTweetsTimeline instantiates a new UserHighlightsTweetsTimeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserHighlightsTweetsTimeline(timeline Timeline) *UserHighlightsTweetsTimeline {
	this := UserHighlightsTweetsTimeline{}
	this.Timeline = timeline
	return &this
}

// NewUserHighlightsTweetsTimelineWithDefaults instantiates a new UserHighlightsTweetsTimeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserHighlightsTweetsTimelineWithDefaults() *UserHighlightsTweetsTimeline {
	this := UserHighlightsTweetsTimeline{}
	return &this
}

// GetTimeline returns the Timeline field value
func (o *UserHighlightsTweetsTimeline) GetTimeline() Timeline {
	if o == nil {
		var ret Timeline
		return ret
	}

	return o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value
// and a boolean to check if the value has been set.
func (o *UserHighlightsTweetsTimeline) GetTimelineOk() (*Timeline, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeline, true
}

// SetTimeline sets field value
func (o *UserHighlightsTweetsTimeline) SetTimeline(v Timeline) {
	o.Timeline = v
}

func (o UserHighlightsTweetsTimeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserHighlightsTweetsTimeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeline"] = o.Timeline
	return toSerialize, nil
}

func (o *UserHighlightsTweetsTimeline) UnmarshalJSON(data []byte) (err error) {
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

	varUserHighlightsTweetsTimeline := _UserHighlightsTweetsTimeline{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserHighlightsTweetsTimeline)

	if err != nil {
		return err
	}

	*o = UserHighlightsTweetsTimeline(varUserHighlightsTweetsTimeline)

	return err
}

type NullableUserHighlightsTweetsTimeline struct {
	value *UserHighlightsTweetsTimeline
	isSet bool
}

func (v NullableUserHighlightsTweetsTimeline) Get() *UserHighlightsTweetsTimeline {
	return v.value
}

func (v *NullableUserHighlightsTweetsTimeline) Set(val *UserHighlightsTweetsTimeline) {
	v.value = val
	v.isSet = true
}

func (v NullableUserHighlightsTweetsTimeline) IsSet() bool {
	return v.isSet
}

func (v *NullableUserHighlightsTweetsTimeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserHighlightsTweetsTimeline(val *UserHighlightsTweetsTimeline) *NullableUserHighlightsTweetsTimeline {
	return &NullableUserHighlightsTweetsTimeline{value: val, isSet: true}
}

func (v NullableUserHighlightsTweetsTimeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserHighlightsTweetsTimeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

