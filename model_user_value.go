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

// checks if the UserValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserValue{}

// UserValue struct for UserValue
type UserValue struct {
	IdStr string `json:"id_str" validate:"regexp=^[0-9]+$"`
}

type _UserValue UserValue

// NewUserValue instantiates a new UserValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserValue(idStr string) *UserValue {
	this := UserValue{}
	this.IdStr = idStr
	return &this
}

// NewUserValueWithDefaults instantiates a new UserValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserValueWithDefaults() *UserValue {
	this := UserValue{}
	return &this
}

// GetIdStr returns the IdStr field value
func (o *UserValue) GetIdStr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdStr
}

// GetIdStrOk returns a tuple with the IdStr field value
// and a boolean to check if the value has been set.
func (o *UserValue) GetIdStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdStr, true
}

// SetIdStr sets field value
func (o *UserValue) SetIdStr(v string) {
	o.IdStr = v
}

func (o UserValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id_str"] = o.IdStr
	return toSerialize, nil
}

func (o *UserValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id_str",
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

	varUserValue := _UserValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserValue)

	if err != nil {
		return err
	}

	*o = UserValue(varUserValue)

	return err
}

type NullableUserValue struct {
	value *UserValue
	isSet bool
}

func (v NullableUserValue) Get() *UserValue {
	return v.value
}

func (v *NullableUserValue) Set(val *UserValue) {
	v.value = val
	v.isSet = true
}

func (v NullableUserValue) IsSet() bool {
	return v.isSet
}

func (v *NullableUserValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserValue(val *UserValue) *NullableUserValue {
	return &NullableUserValue{value: val, isSet: true}
}

func (v NullableUserValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


