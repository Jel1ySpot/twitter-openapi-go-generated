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

// checks if the UserResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResults{}

// UserResults struct for UserResults
type UserResults struct {
	Result *UserUnion `json:"result,omitempty"`
}

// NewUserResults instantiates a new UserResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResults() *UserResults {
	this := UserResults{}
	return &this
}

// NewUserResultsWithDefaults instantiates a new UserResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResultsWithDefaults() *UserResults {
	this := UserResults{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UserResults) GetResult() UserUnion {
	if o == nil || IsNil(o.Result) {
		var ret UserUnion
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResults) GetResultOk() (*UserUnion, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UserResults) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given UserUnion and assigns it to the Result field.
func (o *UserResults) SetResult(v UserUnion) {
	o.Result = &v
}

func (o UserResults) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableUserResults struct {
	value *UserResults
	isSet bool
}

func (v NullableUserResults) Get() *UserResults {
	return v.value
}

func (v *NullableUserResults) Set(val *UserResults) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResults) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResults(val *UserResults) *NullableUserResults {
	return &NullableUserResults{value: val, isSet: true}
}

func (v NullableUserResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
