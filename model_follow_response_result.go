/*
Twitter OpenAPI

Twitter OpenAPI(Swagger) specification

API version: 0.0.1
Contact: yuki@yuki0311.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the FollowResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowResponseResult{}

// FollowResponseResult struct for FollowResponseResult
type FollowResponseResult struct {
	Typename TypeName       `json:"__typename"`
	Timeline FollowTimeline `json:"timeline"`
}

type _FollowResponseResult FollowResponseResult

// NewFollowResponseResult instantiates a new FollowResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowResponseResult(typename TypeName, timeline FollowTimeline) *FollowResponseResult {
	this := FollowResponseResult{}
	this.Typename = typename
	this.Timeline = timeline
	return &this
}

// NewFollowResponseResultWithDefaults instantiates a new FollowResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowResponseResultWithDefaults() *FollowResponseResult {
	this := FollowResponseResult{}
	return &this
}

// GetTypename returns the Typename field value
func (o *FollowResponseResult) GetTypename() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value
// and a boolean to check if the value has been set.
func (o *FollowResponseResult) GetTypenameOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typename, true
}

// SetTypename sets field value
func (o *FollowResponseResult) SetTypename(v TypeName) {
	o.Typename = v
}

// GetTimeline returns the Timeline field value
func (o *FollowResponseResult) GetTimeline() FollowTimeline {
	if o == nil {
		var ret FollowTimeline
		return ret
	}

	return o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value
// and a boolean to check if the value has been set.
func (o *FollowResponseResult) GetTimelineOk() (*FollowTimeline, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeline, true
}

// SetTimeline sets field value
func (o *FollowResponseResult) SetTimeline(v FollowTimeline) {
	o.Timeline = v
}

func (o FollowResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["__typename"] = o.Typename
	toSerialize["timeline"] = o.Timeline
	return toSerialize, nil
}

func (o *FollowResponseResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"__typename",
		"timeline",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFollowResponseResult := _FollowResponseResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFollowResponseResult)

	if err != nil {
		return err
	}

	*o = FollowResponseResult(varFollowResponseResult)

	return err
}

type NullableFollowResponseResult struct {
	value *FollowResponseResult
	isSet bool
}

func (v NullableFollowResponseResult) Get() *FollowResponseResult {
	return v.value
}

func (v *NullableFollowResponseResult) Set(val *FollowResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowResponseResult(val *FollowResponseResult) *NullableFollowResponseResult {
	return &NullableFollowResponseResult{value: val, isSet: true}
}

func (v NullableFollowResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
