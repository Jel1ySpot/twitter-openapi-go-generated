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

// checks if the Retweet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Retweet{}

// Retweet struct for Retweet
type Retweet struct {
	Legacy RetweetLegacy `json:"legacy"`
	RestId string        `json:"rest_id" validate:"regexp=^[0-9]+$"`
}

type _Retweet Retweet

// NewRetweet instantiates a new Retweet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetweet(legacy RetweetLegacy, restId string) *Retweet {
	this := Retweet{}
	this.Legacy = legacy
	this.RestId = restId
	return &this
}

// NewRetweetWithDefaults instantiates a new Retweet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetweetWithDefaults() *Retweet {
	this := Retweet{}
	return &this
}

// GetLegacy returns the Legacy field value
func (o *Retweet) GetLegacy() RetweetLegacy {
	if o == nil {
		var ret RetweetLegacy
		return ret
	}

	return o.Legacy
}

// GetLegacyOk returns a tuple with the Legacy field value
// and a boolean to check if the value has been set.
func (o *Retweet) GetLegacyOk() (*RetweetLegacy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Legacy, true
}

// SetLegacy sets field value
func (o *Retweet) SetLegacy(v RetweetLegacy) {
	o.Legacy = v
}

// GetRestId returns the RestId field value
func (o *Retweet) GetRestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RestId
}

// GetRestIdOk returns a tuple with the RestId field value
// and a boolean to check if the value has been set.
func (o *Retweet) GetRestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestId, true
}

// SetRestId sets field value
func (o *Retweet) SetRestId(v string) {
	o.RestId = v
}

func (o Retweet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Retweet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["legacy"] = o.Legacy
	toSerialize["rest_id"] = o.RestId
	return toSerialize, nil
}

func (o *Retweet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"legacy",
		"rest_id",
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

	varRetweet := _Retweet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRetweet)

	if err != nil {
		return err
	}

	*o = Retweet(varRetweet)

	return err
}

type NullableRetweet struct {
	value *Retweet
	isSet bool
}

func (v NullableRetweet) Get() *Retweet {
	return v.value
}

func (v *NullableRetweet) Set(val *Retweet) {
	v.value = val
	v.isSet = true
}

func (v NullableRetweet) IsSet() bool {
	return v.isSet
}

func (v *NullableRetweet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetweet(val *Retweet) *NullableRetweet {
	return &NullableRetweet{value: val, isSet: true}
}

func (v NullableRetweet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetweet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
