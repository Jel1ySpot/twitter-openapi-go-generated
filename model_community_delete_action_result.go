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

// checks if the CommunityDeleteActionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunityDeleteActionResult{}

// CommunityDeleteActionResult struct for CommunityDeleteActionResult
type CommunityDeleteActionResult struct {
	Typename TypeName `json:"__typename"`
	Reason   string   `json:"reason"`
}

type _CommunityDeleteActionResult CommunityDeleteActionResult

// NewCommunityDeleteActionResult instantiates a new CommunityDeleteActionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityDeleteActionResult(typename TypeName, reason string) *CommunityDeleteActionResult {
	this := CommunityDeleteActionResult{}
	this.Typename = typename
	this.Reason = reason
	return &this
}

// NewCommunityDeleteActionResultWithDefaults instantiates a new CommunityDeleteActionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityDeleteActionResultWithDefaults() *CommunityDeleteActionResult {
	this := CommunityDeleteActionResult{}
	return &this
}

// GetTypename returns the Typename field value
func (o *CommunityDeleteActionResult) GetTypename() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value
// and a boolean to check if the value has been set.
func (o *CommunityDeleteActionResult) GetTypenameOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typename, true
}

// SetTypename sets field value
func (o *CommunityDeleteActionResult) SetTypename(v TypeName) {
	o.Typename = v
}

// GetReason returns the Reason field value
func (o *CommunityDeleteActionResult) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *CommunityDeleteActionResult) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *CommunityDeleteActionResult) SetReason(v string) {
	o.Reason = v
}

func (o CommunityDeleteActionResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunityDeleteActionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["__typename"] = o.Typename
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

func (o *CommunityDeleteActionResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"__typename",
		"reason",
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

	varCommunityDeleteActionResult := _CommunityDeleteActionResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommunityDeleteActionResult)

	if err != nil {
		return err
	}

	*o = CommunityDeleteActionResult(varCommunityDeleteActionResult)

	return err
}

type NullableCommunityDeleteActionResult struct {
	value *CommunityDeleteActionResult
	isSet bool
}

func (v NullableCommunityDeleteActionResult) Get() *CommunityDeleteActionResult {
	return v.value
}

func (v *NullableCommunityDeleteActionResult) Set(val *CommunityDeleteActionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityDeleteActionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityDeleteActionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityDeleteActionResult(val *CommunityDeleteActionResult) *NullableCommunityDeleteActionResult {
	return &NullableCommunityDeleteActionResult{value: val, isSet: true}
}

func (v NullableCommunityDeleteActionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityDeleteActionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
