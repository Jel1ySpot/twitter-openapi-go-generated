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

// checks if the UnfavoriteTweet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnfavoriteTweet{}

// UnfavoriteTweet struct for UnfavoriteTweet
type UnfavoriteTweet struct {
	UnfavoriteTweet string `json:"unfavorite_tweet"`
}

type _UnfavoriteTweet UnfavoriteTweet

// NewUnfavoriteTweet instantiates a new UnfavoriteTweet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnfavoriteTweet(unfavoriteTweet string) *UnfavoriteTweet {
	this := UnfavoriteTweet{}
	this.UnfavoriteTweet = unfavoriteTweet
	return &this
}

// NewUnfavoriteTweetWithDefaults instantiates a new UnfavoriteTweet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnfavoriteTweetWithDefaults() *UnfavoriteTweet {
	this := UnfavoriteTweet{}
	return &this
}

// GetUnfavoriteTweet returns the UnfavoriteTweet field value
func (o *UnfavoriteTweet) GetUnfavoriteTweet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnfavoriteTweet
}

// GetUnfavoriteTweetOk returns a tuple with the UnfavoriteTweet field value
// and a boolean to check if the value has been set.
func (o *UnfavoriteTweet) GetUnfavoriteTweetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnfavoriteTweet, true
}

// SetUnfavoriteTweet sets field value
func (o *UnfavoriteTweet) SetUnfavoriteTweet(v string) {
	o.UnfavoriteTweet = v
}

func (o UnfavoriteTweet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnfavoriteTweet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unfavorite_tweet"] = o.UnfavoriteTweet
	return toSerialize, nil
}

func (o *UnfavoriteTweet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"unfavorite_tweet",
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

	varUnfavoriteTweet := _UnfavoriteTweet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnfavoriteTweet)

	if err != nil {
		return err
	}

	*o = UnfavoriteTweet(varUnfavoriteTweet)

	return err
}

type NullableUnfavoriteTweet struct {
	value *UnfavoriteTweet
	isSet bool
}

func (v NullableUnfavoriteTweet) Get() *UnfavoriteTweet {
	return v.value
}

func (v *NullableUnfavoriteTweet) Set(val *UnfavoriteTweet) {
	v.value = val
	v.isSet = true
}

func (v NullableUnfavoriteTweet) IsSet() bool {
	return v.isSet
}

func (v *NullableUnfavoriteTweet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnfavoriteTweet(val *UnfavoriteTweet) *NullableUnfavoriteTweet {
	return &NullableUnfavoriteTweet{value: val, isSet: true}
}

func (v NullableUnfavoriteTweet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnfavoriteTweet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

