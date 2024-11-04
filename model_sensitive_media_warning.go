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

// checks if the SensitiveMediaWarning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensitiveMediaWarning{}

// SensitiveMediaWarning struct for SensitiveMediaWarning
type SensitiveMediaWarning struct {
	AdultContent bool `json:"adult_content"`
	GraphicViolence bool `json:"graphic_violence"`
	Other bool `json:"other"`
}

type _SensitiveMediaWarning SensitiveMediaWarning

// NewSensitiveMediaWarning instantiates a new SensitiveMediaWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensitiveMediaWarning(adultContent bool, graphicViolence bool, other bool) *SensitiveMediaWarning {
	this := SensitiveMediaWarning{}
	this.AdultContent = adultContent
	this.GraphicViolence = graphicViolence
	this.Other = other
	return &this
}

// NewSensitiveMediaWarningWithDefaults instantiates a new SensitiveMediaWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensitiveMediaWarningWithDefaults() *SensitiveMediaWarning {
	this := SensitiveMediaWarning{}
	return &this
}

// GetAdultContent returns the AdultContent field value
func (o *SensitiveMediaWarning) GetAdultContent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AdultContent
}

// GetAdultContentOk returns a tuple with the AdultContent field value
// and a boolean to check if the value has been set.
func (o *SensitiveMediaWarning) GetAdultContentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdultContent, true
}

// SetAdultContent sets field value
func (o *SensitiveMediaWarning) SetAdultContent(v bool) {
	o.AdultContent = v
}

// GetGraphicViolence returns the GraphicViolence field value
func (o *SensitiveMediaWarning) GetGraphicViolence() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GraphicViolence
}

// GetGraphicViolenceOk returns a tuple with the GraphicViolence field value
// and a boolean to check if the value has been set.
func (o *SensitiveMediaWarning) GetGraphicViolenceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GraphicViolence, true
}

// SetGraphicViolence sets field value
func (o *SensitiveMediaWarning) SetGraphicViolence(v bool) {
	o.GraphicViolence = v
}

// GetOther returns the Other field value
func (o *SensitiveMediaWarning) GetOther() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Other
}

// GetOtherOk returns a tuple with the Other field value
// and a boolean to check if the value has been set.
func (o *SensitiveMediaWarning) GetOtherOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Other, true
}

// SetOther sets field value
func (o *SensitiveMediaWarning) SetOther(v bool) {
	o.Other = v
}

func (o SensitiveMediaWarning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensitiveMediaWarning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adult_content"] = o.AdultContent
	toSerialize["graphic_violence"] = o.GraphicViolence
	toSerialize["other"] = o.Other
	return toSerialize, nil
}

func (o *SensitiveMediaWarning) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"adult_content",
		"graphic_violence",
		"other",
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

	varSensitiveMediaWarning := _SensitiveMediaWarning{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSensitiveMediaWarning)

	if err != nil {
		return err
	}

	*o = SensitiveMediaWarning(varSensitiveMediaWarning)

	return err
}

type NullableSensitiveMediaWarning struct {
	value *SensitiveMediaWarning
	isSet bool
}

func (v NullableSensitiveMediaWarning) Get() *SensitiveMediaWarning {
	return v.value
}

func (v *NullableSensitiveMediaWarning) Set(val *SensitiveMediaWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableSensitiveMediaWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableSensitiveMediaWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensitiveMediaWarning(val *SensitiveMediaWarning) *NullableSensitiveMediaWarning {
	return &NullableSensitiveMediaWarning{value: val, isSet: true}
}

func (v NullableSensitiveMediaWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensitiveMediaWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


