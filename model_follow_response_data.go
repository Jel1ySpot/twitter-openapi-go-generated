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

// checks if the FollowResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowResponseData{}

// FollowResponseData struct for FollowResponseData
type FollowResponseData struct {
	User FollowResponseUser `json:"user"`
}

type _FollowResponseData FollowResponseData

// NewFollowResponseData instantiates a new FollowResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowResponseData(user FollowResponseUser) *FollowResponseData {
	this := FollowResponseData{}
	this.User = user
	return &this
}

// NewFollowResponseDataWithDefaults instantiates a new FollowResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowResponseDataWithDefaults() *FollowResponseData {
	this := FollowResponseData{}
	return &this
}

// GetUser returns the User field value
func (o *FollowResponseData) GetUser() FollowResponseUser {
	if o == nil {
		var ret FollowResponseUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *FollowResponseData) GetUserOk() (*FollowResponseUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *FollowResponseData) SetUser(v FollowResponseUser) {
	o.User = v
}

func (o FollowResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *FollowResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
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

	varFollowResponseData := _FollowResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFollowResponseData)

	if err != nil {
		return err
	}

	*o = FollowResponseData(varFollowResponseData)

	return err
}

type NullableFollowResponseData struct {
	value *FollowResponseData
	isSet bool
}

func (v NullableFollowResponseData) Get() *FollowResponseData {
	return v.value
}

func (v *NullableFollowResponseData) Set(val *FollowResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowResponseData(val *FollowResponseData) *NullableFollowResponseData {
	return &NullableFollowResponseData{value: val, isSet: true}
}

func (v NullableFollowResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

