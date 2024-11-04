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

// checks if the FollowResponseUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowResponseUser{}

// FollowResponseUser struct for FollowResponseUser
type FollowResponseUser struct {
	Result FollowResponseResult `json:"result"`
}

type _FollowResponseUser FollowResponseUser

// NewFollowResponseUser instantiates a new FollowResponseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowResponseUser(result FollowResponseResult) *FollowResponseUser {
	this := FollowResponseUser{}
	this.Result = result
	return &this
}

// NewFollowResponseUserWithDefaults instantiates a new FollowResponseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowResponseUserWithDefaults() *FollowResponseUser {
	this := FollowResponseUser{}
	return &this
}

// GetResult returns the Result field value
func (o *FollowResponseUser) GetResult() FollowResponseResult {
	if o == nil {
		var ret FollowResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetResultOk() (*FollowResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *FollowResponseUser) SetResult(v FollowResponseResult) {
	o.Result = v
}

func (o FollowResponseUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowResponseUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *FollowResponseUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varFollowResponseUser := _FollowResponseUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFollowResponseUser)

	if err != nil {
		return err
	}

	*o = FollowResponseUser(varFollowResponseUser)

	return err
}

type NullableFollowResponseUser struct {
	value *FollowResponseUser
	isSet bool
}

func (v NullableFollowResponseUser) Get() *FollowResponseUser {
	return v.value
}

func (v *NullableFollowResponseUser) Set(val *FollowResponseUser) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowResponseUser) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowResponseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowResponseUser(val *FollowResponseUser) *NullableFollowResponseUser {
	return &NullableFollowResponseUser{value: val, isSet: true}
}

func (v NullableFollowResponseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowResponseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


