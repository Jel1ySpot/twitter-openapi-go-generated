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

// checks if the UserFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFeatures{}

// UserFeatures struct for UserFeatures
type UserFeatures struct {
	MediatoolStudioLibrary bool `json:"mediatool_studio_library"`
}

type _UserFeatures UserFeatures

// NewUserFeatures instantiates a new UserFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFeatures(mediatoolStudioLibrary bool) *UserFeatures {
	this := UserFeatures{}
	this.MediatoolStudioLibrary = mediatoolStudioLibrary
	return &this
}

// NewUserFeaturesWithDefaults instantiates a new UserFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFeaturesWithDefaults() *UserFeatures {
	this := UserFeatures{}
	return &this
}

// GetMediatoolStudioLibrary returns the MediatoolStudioLibrary field value
func (o *UserFeatures) GetMediatoolStudioLibrary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MediatoolStudioLibrary
}

// GetMediatoolStudioLibraryOk returns a tuple with the MediatoolStudioLibrary field value
// and a boolean to check if the value has been set.
func (o *UserFeatures) GetMediatoolStudioLibraryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediatoolStudioLibrary, true
}

// SetMediatoolStudioLibrary sets field value
func (o *UserFeatures) SetMediatoolStudioLibrary(v bool) {
	o.MediatoolStudioLibrary = v
}

func (o UserFeatures) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mediatool_studio_library"] = o.MediatoolStudioLibrary
	return toSerialize, nil
}

func (o *UserFeatures) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mediatool_studio_library",
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

	varUserFeatures := _UserFeatures{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserFeatures)

	if err != nil {
		return err
	}

	*o = UserFeatures(varUserFeatures)

	return err
}

type NullableUserFeatures struct {
	value *UserFeatures
	isSet bool
}

func (v NullableUserFeatures) Get() *UserFeatures {
	return v.value
}

func (v *NullableUserFeatures) Set(val *UserFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFeatures(val *UserFeatures) *NullableUserFeatures {
	return &NullableUserFeatures{value: val, isSet: true}
}

func (v NullableUserFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
