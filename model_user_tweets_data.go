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

// checks if the UserTweetsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTweetsData{}

// UserTweetsData struct for UserTweetsData
type UserTweetsData struct {
	User UserTweetsUser `json:"user"`
}

type _UserTweetsData UserTweetsData

// NewUserTweetsData instantiates a new UserTweetsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTweetsData(user UserTweetsUser) *UserTweetsData {
	this := UserTweetsData{}
	this.User = user
	return &this
}

// NewUserTweetsDataWithDefaults instantiates a new UserTweetsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTweetsDataWithDefaults() *UserTweetsData {
	this := UserTweetsData{}
	return &this
}

// GetUser returns the User field value
func (o *UserTweetsData) GetUser() UserTweetsUser {
	if o == nil {
		var ret UserTweetsUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserTweetsData) GetUserOk() (*UserTweetsUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserTweetsData) SetUser(v UserTweetsUser) {
	o.User = v
}

func (o UserTweetsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTweetsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *UserTweetsData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
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

	varUserTweetsData := _UserTweetsData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserTweetsData)

	if err != nil {
		return err
	}

	*o = UserTweetsData(varUserTweetsData)

	return err
}

type NullableUserTweetsData struct {
	value *UserTweetsData
	isSet bool
}

func (v NullableUserTweetsData) Get() *UserTweetsData {
	return v.value
}

func (v *NullableUserTweetsData) Set(val *UserTweetsData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTweetsData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTweetsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTweetsData(val *UserTweetsData) *NullableUserTweetsData {
	return &NullableUserTweetsData{value: val, isSet: true}
}

func (v NullableUserTweetsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTweetsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
