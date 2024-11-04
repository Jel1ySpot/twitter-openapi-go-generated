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

// checks if the UserResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResponseData{}

// UserResponseData struct for UserResponseData
type UserResponseData struct {
	User *UserResults `json:"user,omitempty"`
}

// NewUserResponseData instantiates a new UserResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponseData() *UserResponseData {
	this := UserResponseData{}
	return &this
}

// NewUserResponseDataWithDefaults instantiates a new UserResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseDataWithDefaults() *UserResponseData {
	this := UserResponseData{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserResponseData) GetUser() UserResults {
	if o == nil || IsNil(o.User) {
		var ret UserResults
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseData) GetUserOk() (*UserResults, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserResponseData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserResults and assigns it to the User field.
func (o *UserResponseData) SetUser(v UserResults) {
	o.User = &v
}

func (o UserResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUserResponseData struct {
	value *UserResponseData
	isSet bool
}

func (v NullableUserResponseData) Get() *UserResponseData {
	return v.value
}

func (v *NullableUserResponseData) Set(val *UserResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponseData(val *UserResponseData) *NullableUserResponseData {
	return &NullableUserResponseData{value: val, isSet: true}
}

func (v NullableUserResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

