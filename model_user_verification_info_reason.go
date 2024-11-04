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

// checks if the UserVerificationInfoReason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserVerificationInfoReason{}

// UserVerificationInfoReason struct for UserVerificationInfoReason
type UserVerificationInfoReason struct {
	Description UserVerificationInfoReasonDescription `json:"description"`
	OverrideVerifiedYear int32 `json:"override_verified_year"`
	VerifiedSinceMsec string `json:"verified_since_msec" validate:"regexp=^-?[0-9]+$"`
}

type _UserVerificationInfoReason UserVerificationInfoReason

// NewUserVerificationInfoReason instantiates a new UserVerificationInfoReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserVerificationInfoReason(description UserVerificationInfoReasonDescription, overrideVerifiedYear int32, verifiedSinceMsec string) *UserVerificationInfoReason {
	this := UserVerificationInfoReason{}
	this.Description = description
	this.OverrideVerifiedYear = overrideVerifiedYear
	this.VerifiedSinceMsec = verifiedSinceMsec
	return &this
}

// NewUserVerificationInfoReasonWithDefaults instantiates a new UserVerificationInfoReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserVerificationInfoReasonWithDefaults() *UserVerificationInfoReason {
	this := UserVerificationInfoReason{}
	return &this
}

// GetDescription returns the Description field value
func (o *UserVerificationInfoReason) GetDescription() UserVerificationInfoReasonDescription {
	if o == nil {
		var ret UserVerificationInfoReasonDescription
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UserVerificationInfoReason) GetDescriptionOk() (*UserVerificationInfoReasonDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UserVerificationInfoReason) SetDescription(v UserVerificationInfoReasonDescription) {
	o.Description = v
}

// GetOverrideVerifiedYear returns the OverrideVerifiedYear field value
func (o *UserVerificationInfoReason) GetOverrideVerifiedYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OverrideVerifiedYear
}

// GetOverrideVerifiedYearOk returns a tuple with the OverrideVerifiedYear field value
// and a boolean to check if the value has been set.
func (o *UserVerificationInfoReason) GetOverrideVerifiedYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverrideVerifiedYear, true
}

// SetOverrideVerifiedYear sets field value
func (o *UserVerificationInfoReason) SetOverrideVerifiedYear(v int32) {
	o.OverrideVerifiedYear = v
}

// GetVerifiedSinceMsec returns the VerifiedSinceMsec field value
func (o *UserVerificationInfoReason) GetVerifiedSinceMsec() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerifiedSinceMsec
}

// GetVerifiedSinceMsecOk returns a tuple with the VerifiedSinceMsec field value
// and a boolean to check if the value has been set.
func (o *UserVerificationInfoReason) GetVerifiedSinceMsecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifiedSinceMsec, true
}

// SetVerifiedSinceMsec sets field value
func (o *UserVerificationInfoReason) SetVerifiedSinceMsec(v string) {
	o.VerifiedSinceMsec = v
}

func (o UserVerificationInfoReason) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserVerificationInfoReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["override_verified_year"] = o.OverrideVerifiedYear
	toSerialize["verified_since_msec"] = o.VerifiedSinceMsec
	return toSerialize, nil
}

func (o *UserVerificationInfoReason) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"override_verified_year",
		"verified_since_msec",
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

	varUserVerificationInfoReason := _UserVerificationInfoReason{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserVerificationInfoReason)

	if err != nil {
		return err
	}

	*o = UserVerificationInfoReason(varUserVerificationInfoReason)

	return err
}

type NullableUserVerificationInfoReason struct {
	value *UserVerificationInfoReason
	isSet bool
}

func (v NullableUserVerificationInfoReason) Get() *UserVerificationInfoReason {
	return v.value
}

func (v *NullableUserVerificationInfoReason) Set(val *UserVerificationInfoReason) {
	v.value = val
	v.isSet = true
}

func (v NullableUserVerificationInfoReason) IsSet() bool {
	return v.isSet
}

func (v *NullableUserVerificationInfoReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserVerificationInfoReason(val *UserVerificationInfoReason) *NullableUserVerificationInfoReason {
	return &NullableUserVerificationInfoReason{value: val, isSet: true}
}

func (v NullableUserVerificationInfoReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserVerificationInfoReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

