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

// checks if the CommunityUnpinActionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunityUnpinActionResult{}

// CommunityUnpinActionResult struct for CommunityUnpinActionResult
type CommunityUnpinActionResult struct {
	Typename TypeName `json:"__typename"`
}

type _CommunityUnpinActionResult CommunityUnpinActionResult

// NewCommunityUnpinActionResult instantiates a new CommunityUnpinActionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityUnpinActionResult(typename TypeName) *CommunityUnpinActionResult {
	this := CommunityUnpinActionResult{}
	this.Typename = typename
	return &this
}

// NewCommunityUnpinActionResultWithDefaults instantiates a new CommunityUnpinActionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityUnpinActionResultWithDefaults() *CommunityUnpinActionResult {
	this := CommunityUnpinActionResult{}
	return &this
}

// GetTypename returns the Typename field value
func (o *CommunityUnpinActionResult) GetTypename() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value
// and a boolean to check if the value has been set.
func (o *CommunityUnpinActionResult) GetTypenameOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typename, true
}

// SetTypename sets field value
func (o *CommunityUnpinActionResult) SetTypename(v TypeName) {
	o.Typename = v
}

func (o CommunityUnpinActionResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunityUnpinActionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["__typename"] = o.Typename
	return toSerialize, nil
}

func (o *CommunityUnpinActionResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"__typename",
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

	varCommunityUnpinActionResult := _CommunityUnpinActionResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommunityUnpinActionResult)

	if err != nil {
		return err
	}

	*o = CommunityUnpinActionResult(varCommunityUnpinActionResult)

	return err
}

type NullableCommunityUnpinActionResult struct {
	value *CommunityUnpinActionResult
	isSet bool
}

func (v NullableCommunityUnpinActionResult) Get() *CommunityUnpinActionResult {
	return v.value
}

func (v *NullableCommunityUnpinActionResult) Set(val *CommunityUnpinActionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityUnpinActionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityUnpinActionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityUnpinActionResult(val *CommunityUnpinActionResult) *NullableCommunityUnpinActionResult {
	return &NullableCommunityUnpinActionResult{value: val, isSet: true}
}

func (v NullableCommunityUnpinActionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityUnpinActionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


