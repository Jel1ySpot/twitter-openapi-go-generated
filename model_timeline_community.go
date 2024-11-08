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

// checks if the TimelineCommunity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineCommunity{}

// TimelineCommunity struct for TimelineCommunity
type TimelineCommunity struct {
	Typename             *TypeName `json:"__typename,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimelineCommunity TimelineCommunity

// NewTimelineCommunity instantiates a new TimelineCommunity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineCommunity() *TimelineCommunity {
	this := TimelineCommunity{}
	return &this
}

// NewTimelineCommunityWithDefaults instantiates a new TimelineCommunity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineCommunityWithDefaults() *TimelineCommunity {
	this := TimelineCommunity{}
	return &this
}

// GetTypename returns the Typename field value if set, zero value otherwise.
func (o *TimelineCommunity) GetTypename() TypeName {
	if o == nil || IsNil(o.Typename) {
		var ret TypeName
		return ret
	}
	return *o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCommunity) GetTypenameOk() (*TypeName, bool) {
	if o == nil || IsNil(o.Typename) {
		return nil, false
	}
	return o.Typename, true
}

// HasTypename returns a boolean if a field has been set.
func (o *TimelineCommunity) HasTypename() bool {
	if o != nil && !IsNil(o.Typename) {
		return true
	}

	return false
}

// SetTypename gets a reference to the given TypeName and assigns it to the Typename field.
func (o *TimelineCommunity) SetTypename(v TypeName) {
	o.Typename = &v
}

func (o TimelineCommunity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineCommunity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Typename) {
		toSerialize["__typename"] = o.Typename
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimelineCommunity) UnmarshalJSON(data []byte) (err error) {
	varTimelineCommunity := _TimelineCommunity{}

	err = json.Unmarshal(data, &varTimelineCommunity)

	if err != nil {
		return err
	}

	*o = TimelineCommunity(varTimelineCommunity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "__typename")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimelineCommunity struct {
	value *TimelineCommunity
	isSet bool
}

func (v NullableTimelineCommunity) Get() *TimelineCommunity {
	return v.value
}

func (v *NullableTimelineCommunity) Set(val *TimelineCommunity) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineCommunity) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineCommunity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineCommunity(val *TimelineCommunity) *NullableTimelineCommunity {
	return &NullableTimelineCommunity{value: val, isSet: true}
}

func (v NullableTimelineCommunity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineCommunity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
