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

// checks if the Community type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Community{}

// Community struct for Community
type Community struct {
	Result CommunityData `json:"result"`
}

type _Community Community

// NewCommunity instantiates a new Community object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunity(result CommunityData) *Community {
	this := Community{}
	this.Result = result
	return &this
}

// NewCommunityWithDefaults instantiates a new Community object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityWithDefaults() *Community {
	this := Community{}
	return &this
}

// GetResult returns the Result field value
func (o *Community) GetResult() CommunityData {
	if o == nil {
		var ret CommunityData
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *Community) GetResultOk() (*CommunityData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *Community) SetResult(v CommunityData) {
	o.Result = v
}

func (o Community) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Community) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *Community) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varCommunity := _Community{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommunity)

	if err != nil {
		return err
	}

	*o = Community(varCommunity)

	return err
}

type NullableCommunity struct {
	value *Community
	isSet bool
}

func (v NullableCommunity) Get() *Community {
	return v.value
}

func (v *NullableCommunity) Set(val *Community) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunity) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunity(val *Community) *NullableCommunity {
	return &NullableCommunity{value: val, isSet: true}
}

func (v NullableCommunity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
