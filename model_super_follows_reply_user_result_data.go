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

// checks if the SuperFollowsReplyUserResultData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuperFollowsReplyUserResultData{}

// SuperFollowsReplyUserResultData struct for SuperFollowsReplyUserResultData
type SuperFollowsReplyUserResultData struct {
	Typename TypeName                          `json:"__typename"`
	Legacy   SuperFollowsReplyUserResultLegacy `json:"legacy"`
}

type _SuperFollowsReplyUserResultData SuperFollowsReplyUserResultData

// NewSuperFollowsReplyUserResultData instantiates a new SuperFollowsReplyUserResultData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuperFollowsReplyUserResultData(typename TypeName, legacy SuperFollowsReplyUserResultLegacy) *SuperFollowsReplyUserResultData {
	this := SuperFollowsReplyUserResultData{}
	this.Typename = typename
	this.Legacy = legacy
	return &this
}

// NewSuperFollowsReplyUserResultDataWithDefaults instantiates a new SuperFollowsReplyUserResultData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuperFollowsReplyUserResultDataWithDefaults() *SuperFollowsReplyUserResultData {
	this := SuperFollowsReplyUserResultData{}
	return &this
}

// GetTypename returns the Typename field value
func (o *SuperFollowsReplyUserResultData) GetTypename() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value
// and a boolean to check if the value has been set.
func (o *SuperFollowsReplyUserResultData) GetTypenameOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typename, true
}

// SetTypename sets field value
func (o *SuperFollowsReplyUserResultData) SetTypename(v TypeName) {
	o.Typename = v
}

// GetLegacy returns the Legacy field value
func (o *SuperFollowsReplyUserResultData) GetLegacy() SuperFollowsReplyUserResultLegacy {
	if o == nil {
		var ret SuperFollowsReplyUserResultLegacy
		return ret
	}

	return o.Legacy
}

// GetLegacyOk returns a tuple with the Legacy field value
// and a boolean to check if the value has been set.
func (o *SuperFollowsReplyUserResultData) GetLegacyOk() (*SuperFollowsReplyUserResultLegacy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Legacy, true
}

// SetLegacy sets field value
func (o *SuperFollowsReplyUserResultData) SetLegacy(v SuperFollowsReplyUserResultLegacy) {
	o.Legacy = v
}

func (o SuperFollowsReplyUserResultData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuperFollowsReplyUserResultData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["__typename"] = o.Typename
	toSerialize["legacy"] = o.Legacy
	return toSerialize, nil
}

func (o *SuperFollowsReplyUserResultData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"__typename",
		"legacy",
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

	varSuperFollowsReplyUserResultData := _SuperFollowsReplyUserResultData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuperFollowsReplyUserResultData)

	if err != nil {
		return err
	}

	*o = SuperFollowsReplyUserResultData(varSuperFollowsReplyUserResultData)

	return err
}

type NullableSuperFollowsReplyUserResultData struct {
	value *SuperFollowsReplyUserResultData
	isSet bool
}

func (v NullableSuperFollowsReplyUserResultData) Get() *SuperFollowsReplyUserResultData {
	return v.value
}

func (v *NullableSuperFollowsReplyUserResultData) Set(val *SuperFollowsReplyUserResultData) {
	v.value = val
	v.isSet = true
}

func (v NullableSuperFollowsReplyUserResultData) IsSet() bool {
	return v.isSet
}

func (v *NullableSuperFollowsReplyUserResultData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuperFollowsReplyUserResultData(val *SuperFollowsReplyUserResultData) *NullableSuperFollowsReplyUserResultData {
	return &NullableSuperFollowsReplyUserResultData{value: val, isSet: true}
}

func (v NullableSuperFollowsReplyUserResultData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuperFollowsReplyUserResultData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
