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

// checks if the SelfThread type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfThread{}

// SelfThread struct for SelfThread
type SelfThread struct {
	IdStr string `json:"id_str" validate:"regexp=^[0-9]+$"`
}

type _SelfThread SelfThread

// NewSelfThread instantiates a new SelfThread object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfThread(idStr string) *SelfThread {
	this := SelfThread{}
	this.IdStr = idStr
	return &this
}

// NewSelfThreadWithDefaults instantiates a new SelfThread object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfThreadWithDefaults() *SelfThread {
	this := SelfThread{}
	return &this
}

// GetIdStr returns the IdStr field value
func (o *SelfThread) GetIdStr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdStr
}

// GetIdStrOk returns a tuple with the IdStr field value
// and a boolean to check if the value has been set.
func (o *SelfThread) GetIdStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdStr, true
}

// SetIdStr sets field value
func (o *SelfThread) SetIdStr(v string) {
	o.IdStr = v
}

func (o SelfThread) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfThread) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id_str"] = o.IdStr
	return toSerialize, nil
}

func (o *SelfThread) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id_str",
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

	varSelfThread := _SelfThread{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSelfThread)

	if err != nil {
		return err
	}

	*o = SelfThread(varSelfThread)

	return err
}

type NullableSelfThread struct {
	value *SelfThread
	isSet bool
}

func (v NullableSelfThread) Get() *SelfThread {
	return v.value
}

func (v *NullableSelfThread) Set(val *SelfThread) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfThread) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfThread) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfThread(val *SelfThread) *NullableSelfThread {
	return &NullableSelfThread{value: val, isSet: true}
}

func (v NullableSelfThread) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfThread) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
