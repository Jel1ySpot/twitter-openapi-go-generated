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

// checks if the ErrorsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorsData{}

// ErrorsData struct for ErrorsData
type ErrorsData struct {
	User *string `json:"user,omitempty" validate:"regexp=dummy"`
}

// NewErrorsData instantiates a new ErrorsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorsData() *ErrorsData {
	this := ErrorsData{}
	return &this
}

// NewErrorsDataWithDefaults instantiates a new ErrorsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsDataWithDefaults() *ErrorsData {
	this := ErrorsData{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ErrorsData) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsData) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ErrorsData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ErrorsData) SetUser(v string) {
	o.User = &v
}

func (o ErrorsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableErrorsData struct {
	value *ErrorsData
	isSet bool
}

func (v NullableErrorsData) Get() *ErrorsData {
	return v.value
}

func (v *NullableErrorsData) Set(val *ErrorsData) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsData) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsData(val *ErrorsData) *NullableErrorsData {
	return &NullableErrorsData{value: val, isSet: true}
}

func (v NullableErrorsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
