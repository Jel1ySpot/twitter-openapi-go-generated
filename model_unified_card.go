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

// checks if the UnifiedCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnifiedCard{}

// UnifiedCard struct for UnifiedCard
type UnifiedCard struct {
	CardFetchState string `json:"card_fetch_state"`
}

type _UnifiedCard UnifiedCard

// NewUnifiedCard instantiates a new UnifiedCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnifiedCard(cardFetchState string) *UnifiedCard {
	this := UnifiedCard{}
	this.CardFetchState = cardFetchState
	return &this
}

// NewUnifiedCardWithDefaults instantiates a new UnifiedCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnifiedCardWithDefaults() *UnifiedCard {
	this := UnifiedCard{}
	return &this
}

// GetCardFetchState returns the CardFetchState field value
func (o *UnifiedCard) GetCardFetchState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardFetchState
}

// GetCardFetchStateOk returns a tuple with the CardFetchState field value
// and a boolean to check if the value has been set.
func (o *UnifiedCard) GetCardFetchStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardFetchState, true
}

// SetCardFetchState sets field value
func (o *UnifiedCard) SetCardFetchState(v string) {
	o.CardFetchState = v
}

func (o UnifiedCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnifiedCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["card_fetch_state"] = o.CardFetchState
	return toSerialize, nil
}

func (o *UnifiedCard) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"card_fetch_state",
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

	varUnifiedCard := _UnifiedCard{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnifiedCard)

	if err != nil {
		return err
	}

	*o = UnifiedCard(varUnifiedCard)

	return err
}

type NullableUnifiedCard struct {
	value *UnifiedCard
	isSet bool
}

func (v NullableUnifiedCard) Get() *UnifiedCard {
	return v.value
}

func (v *NullableUnifiedCard) Set(val *UnifiedCard) {
	v.value = val
	v.isSet = true
}

func (v NullableUnifiedCard) IsSet() bool {
	return v.isSet
}

func (v *NullableUnifiedCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnifiedCard(val *UnifiedCard) *NullableUnifiedCard {
	return &NullableUnifiedCard{value: val, isSet: true}
}

func (v NullableUnifiedCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnifiedCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
