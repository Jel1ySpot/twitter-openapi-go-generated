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

// checks if the UserProfessional type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProfessional{}

// UserProfessional struct for UserProfessional
type UserProfessional struct {
	Category []UserProfessionalCategory `json:"category"`
	ProfessionalType string `json:"professional_type"`
	RestId string `json:"rest_id" validate:"regexp=^[0-9]+$"`
}

type _UserProfessional UserProfessional

// NewUserProfessional instantiates a new UserProfessional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfessional(category []UserProfessionalCategory, professionalType string, restId string) *UserProfessional {
	this := UserProfessional{}
	this.Category = category
	this.ProfessionalType = professionalType
	this.RestId = restId
	return &this
}

// NewUserProfessionalWithDefaults instantiates a new UserProfessional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfessionalWithDefaults() *UserProfessional {
	this := UserProfessional{}
	return &this
}

// GetCategory returns the Category field value
func (o *UserProfessional) GetCategory() []UserProfessionalCategory {
	if o == nil {
		var ret []UserProfessionalCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *UserProfessional) GetCategoryOk() ([]UserProfessionalCategory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category, true
}

// SetCategory sets field value
func (o *UserProfessional) SetCategory(v []UserProfessionalCategory) {
	o.Category = v
}

// GetProfessionalType returns the ProfessionalType field value
func (o *UserProfessional) GetProfessionalType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfessionalType
}

// GetProfessionalTypeOk returns a tuple with the ProfessionalType field value
// and a boolean to check if the value has been set.
func (o *UserProfessional) GetProfessionalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfessionalType, true
}

// SetProfessionalType sets field value
func (o *UserProfessional) SetProfessionalType(v string) {
	o.ProfessionalType = v
}

// GetRestId returns the RestId field value
func (o *UserProfessional) GetRestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RestId
}

// GetRestIdOk returns a tuple with the RestId field value
// and a boolean to check if the value has been set.
func (o *UserProfessional) GetRestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestId, true
}

// SetRestId sets field value
func (o *UserProfessional) SetRestId(v string) {
	o.RestId = v
}

func (o UserProfessional) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProfessional) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category
	toSerialize["professional_type"] = o.ProfessionalType
	toSerialize["rest_id"] = o.RestId
	return toSerialize, nil
}

func (o *UserProfessional) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"category",
		"professional_type",
		"rest_id",
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

	varUserProfessional := _UserProfessional{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserProfessional)

	if err != nil {
		return err
	}

	*o = UserProfessional(varUserProfessional)

	return err
}

type NullableUserProfessional struct {
	value *UserProfessional
	isSet bool
}

func (v NullableUserProfessional) Get() *UserProfessional {
	return v.value
}

func (v *NullableUserProfessional) Set(val *UserProfessional) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfessional) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfessional) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfessional(val *UserProfessional) *NullableUserProfessional {
	return &NullableUserProfessional{value: val, isSet: true}
}

func (v NullableUserProfessional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfessional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

