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

// checks if the CommunitiesActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunitiesActions{}

// CommunitiesActions struct for CommunitiesActions
type CommunitiesActions struct {
	Create bool `json:"create"`
}

type _CommunitiesActions CommunitiesActions

// NewCommunitiesActions instantiates a new CommunitiesActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunitiesActions(create bool) *CommunitiesActions {
	this := CommunitiesActions{}
	this.Create = create
	return &this
}

// NewCommunitiesActionsWithDefaults instantiates a new CommunitiesActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunitiesActionsWithDefaults() *CommunitiesActions {
	this := CommunitiesActions{}
	return &this
}

// GetCreate returns the Create field value
func (o *CommunitiesActions) GetCreate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Create
}

// GetCreateOk returns a tuple with the Create field value
// and a boolean to check if the value has been set.
func (o *CommunitiesActions) GetCreateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Create, true
}

// SetCreate sets field value
func (o *CommunitiesActions) SetCreate(v bool) {
	o.Create = v
}

func (o CommunitiesActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunitiesActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["create"] = o.Create
	return toSerialize, nil
}

func (o *CommunitiesActions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"create",
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

	varCommunitiesActions := _CommunitiesActions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommunitiesActions)

	if err != nil {
		return err
	}

	*o = CommunitiesActions(varCommunitiesActions)

	return err
}

type NullableCommunitiesActions struct {
	value *CommunitiesActions
	isSet bool
}

func (v NullableCommunitiesActions) Get() *CommunitiesActions {
	return v.value
}

func (v *NullableCommunitiesActions) Set(val *CommunitiesActions) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunitiesActions) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunitiesActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunitiesActions(val *CommunitiesActions) *NullableCommunitiesActions {
	return &NullableCommunitiesActions{value: val, isSet: true}
}

func (v NullableCommunitiesActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunitiesActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
