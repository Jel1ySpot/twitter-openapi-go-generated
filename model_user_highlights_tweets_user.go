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

// checks if the UserHighlightsTweetsUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserHighlightsTweetsUser{}

// UserHighlightsTweetsUser struct for UserHighlightsTweetsUser
type UserHighlightsTweetsUser struct {
	Result UserHighlightsTweetsResult `json:"result"`
}

type _UserHighlightsTweetsUser UserHighlightsTweetsUser

// NewUserHighlightsTweetsUser instantiates a new UserHighlightsTweetsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserHighlightsTweetsUser(result UserHighlightsTweetsResult) *UserHighlightsTweetsUser {
	this := UserHighlightsTweetsUser{}
	this.Result = result
	return &this
}

// NewUserHighlightsTweetsUserWithDefaults instantiates a new UserHighlightsTweetsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserHighlightsTweetsUserWithDefaults() *UserHighlightsTweetsUser {
	this := UserHighlightsTweetsUser{}
	return &this
}

// GetResult returns the Result field value
func (o *UserHighlightsTweetsUser) GetResult() UserHighlightsTweetsResult {
	if o == nil {
		var ret UserHighlightsTweetsResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *UserHighlightsTweetsUser) GetResultOk() (*UserHighlightsTweetsResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *UserHighlightsTweetsUser) SetResult(v UserHighlightsTweetsResult) {
	o.Result = v
}

func (o UserHighlightsTweetsUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserHighlightsTweetsUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *UserHighlightsTweetsUser) UnmarshalJSON(data []byte) (err error) {
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

	varUserHighlightsTweetsUser := _UserHighlightsTweetsUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserHighlightsTweetsUser)

	if err != nil {
		return err
	}

	*o = UserHighlightsTweetsUser(varUserHighlightsTweetsUser)

	return err
}

type NullableUserHighlightsTweetsUser struct {
	value *UserHighlightsTweetsUser
	isSet bool
}

func (v NullableUserHighlightsTweetsUser) Get() *UserHighlightsTweetsUser {
	return v.value
}

func (v *NullableUserHighlightsTweetsUser) Set(val *UserHighlightsTweetsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUserHighlightsTweetsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUserHighlightsTweetsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserHighlightsTweetsUser(val *UserHighlightsTweetsUser) *NullableUserHighlightsTweetsUser {
	return &NullableUserHighlightsTweetsUser{value: val, isSet: true}
}

func (v NullableUserHighlightsTweetsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserHighlightsTweetsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


