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

// checks if the CtaClientEventInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CtaClientEventInfo{}

// CtaClientEventInfo struct for CtaClientEventInfo
type CtaClientEventInfo struct {
	Action string `json:"action"`
}

type _CtaClientEventInfo CtaClientEventInfo

// NewCtaClientEventInfo instantiates a new CtaClientEventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCtaClientEventInfo(action string) *CtaClientEventInfo {
	this := CtaClientEventInfo{}
	this.Action = action
	return &this
}

// NewCtaClientEventInfoWithDefaults instantiates a new CtaClientEventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCtaClientEventInfoWithDefaults() *CtaClientEventInfo {
	this := CtaClientEventInfo{}
	return &this
}

// GetAction returns the Action field value
func (o *CtaClientEventInfo) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *CtaClientEventInfo) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *CtaClientEventInfo) SetAction(v string) {
	o.Action = v
}

func (o CtaClientEventInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CtaClientEventInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

func (o *CtaClientEventInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varCtaClientEventInfo := _CtaClientEventInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCtaClientEventInfo)

	if err != nil {
		return err
	}

	*o = CtaClientEventInfo(varCtaClientEventInfo)

	return err
}

type NullableCtaClientEventInfo struct {
	value *CtaClientEventInfo
	isSet bool
}

func (v NullableCtaClientEventInfo) Get() *CtaClientEventInfo {
	return v.value
}

func (v *NullableCtaClientEventInfo) Set(val *CtaClientEventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCtaClientEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCtaClientEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCtaClientEventInfo(val *CtaClientEventInfo) *NullableCtaClientEventInfo {
	return &NullableCtaClientEventInfo{value: val, isSet: true}
}

func (v NullableCtaClientEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCtaClientEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
