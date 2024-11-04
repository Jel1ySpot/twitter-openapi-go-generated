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

// checks if the TimelineMessagePrompt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineMessagePrompt{}

// TimelineMessagePrompt struct for TimelineMessagePrompt
type TimelineMessagePrompt struct {
	Typename *TypeName `json:"__typename,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimelineMessagePrompt TimelineMessagePrompt

// NewTimelineMessagePrompt instantiates a new TimelineMessagePrompt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineMessagePrompt() *TimelineMessagePrompt {
	this := TimelineMessagePrompt{}
	return &this
}

// NewTimelineMessagePromptWithDefaults instantiates a new TimelineMessagePrompt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineMessagePromptWithDefaults() *TimelineMessagePrompt {
	this := TimelineMessagePrompt{}
	return &this
}

// GetTypename returns the Typename field value if set, zero value otherwise.
func (o *TimelineMessagePrompt) GetTypename() TypeName {
	if o == nil || IsNil(o.Typename) {
		var ret TypeName
		return ret
	}
	return *o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineMessagePrompt) GetTypenameOk() (*TypeName, bool) {
	if o == nil || IsNil(o.Typename) {
		return nil, false
	}
	return o.Typename, true
}

// HasTypename returns a boolean if a field has been set.
func (o *TimelineMessagePrompt) HasTypename() bool {
	if o != nil && !IsNil(o.Typename) {
		return true
	}

	return false
}

// SetTypename gets a reference to the given TypeName and assigns it to the Typename field.
func (o *TimelineMessagePrompt) SetTypename(v TypeName) {
	o.Typename = &v
}

func (o TimelineMessagePrompt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineMessagePrompt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Typename) {
		toSerialize["__typename"] = o.Typename
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimelineMessagePrompt) UnmarshalJSON(data []byte) (err error) {
	varTimelineMessagePrompt := _TimelineMessagePrompt{}

	err = json.Unmarshal(data, &varTimelineMessagePrompt)

	if err != nil {
		return err
	}

	*o = TimelineMessagePrompt(varTimelineMessagePrompt)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "__typename")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimelineMessagePrompt struct {
	value *TimelineMessagePrompt
	isSet bool
}

func (v NullableTimelineMessagePrompt) Get() *TimelineMessagePrompt {
	return v.value
}

func (v *NullableTimelineMessagePrompt) Set(val *TimelineMessagePrompt) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineMessagePrompt) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineMessagePrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineMessagePrompt(val *TimelineMessagePrompt) *NullableTimelineMessagePrompt {
	return &NullableTimelineMessagePrompt{value: val, isSet: true}
}

func (v NullableTimelineMessagePrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineMessagePrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

