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

// checks if the TweetCardPlatform type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetCardPlatform{}

// TweetCardPlatform struct for TweetCardPlatform
type TweetCardPlatform struct {
	Audience TweetCardPlatformAudience `json:"audience"`
	Device TweetCardPlatformDevice `json:"device"`
}

type _TweetCardPlatform TweetCardPlatform

// NewTweetCardPlatform instantiates a new TweetCardPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetCardPlatform(audience TweetCardPlatformAudience, device TweetCardPlatformDevice) *TweetCardPlatform {
	this := TweetCardPlatform{}
	this.Audience = audience
	this.Device = device
	return &this
}

// NewTweetCardPlatformWithDefaults instantiates a new TweetCardPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetCardPlatformWithDefaults() *TweetCardPlatform {
	this := TweetCardPlatform{}
	return &this
}

// GetAudience returns the Audience field value
func (o *TweetCardPlatform) GetAudience() TweetCardPlatformAudience {
	if o == nil {
		var ret TweetCardPlatformAudience
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *TweetCardPlatform) GetAudienceOk() (*TweetCardPlatformAudience, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *TweetCardPlatform) SetAudience(v TweetCardPlatformAudience) {
	o.Audience = v
}

// GetDevice returns the Device field value
func (o *TweetCardPlatform) GetDevice() TweetCardPlatformDevice {
	if o == nil {
		var ret TweetCardPlatformDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *TweetCardPlatform) GetDeviceOk() (*TweetCardPlatformDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *TweetCardPlatform) SetDevice(v TweetCardPlatformDevice) {
	o.Device = v
}

func (o TweetCardPlatform) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetCardPlatform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audience"] = o.Audience
	toSerialize["device"] = o.Device
	return toSerialize, nil
}

func (o *TweetCardPlatform) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"audience",
		"device",
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

	varTweetCardPlatform := _TweetCardPlatform{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetCardPlatform)

	if err != nil {
		return err
	}

	*o = TweetCardPlatform(varTweetCardPlatform)

	return err
}

type NullableTweetCardPlatform struct {
	value *TweetCardPlatform
	isSet bool
}

func (v NullableTweetCardPlatform) Get() *TweetCardPlatform {
	return v.value
}

func (v *NullableTweetCardPlatform) Set(val *TweetCardPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetCardPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetCardPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetCardPlatform(val *TweetCardPlatform) *NullableTweetCardPlatform {
	return &NullableTweetCardPlatform{value: val, isSet: true}
}

func (v NullableTweetCardPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetCardPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

