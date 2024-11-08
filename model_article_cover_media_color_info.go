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

// checks if the ArticleCoverMediaColorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticleCoverMediaColorInfo{}

// ArticleCoverMediaColorInfo struct for ArticleCoverMediaColorInfo
type ArticleCoverMediaColorInfo struct {
	Palette []ArticleCoverMediaColorInfoPalette `json:"palette"`
}

type _ArticleCoverMediaColorInfo ArticleCoverMediaColorInfo

// NewArticleCoverMediaColorInfo instantiates a new ArticleCoverMediaColorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleCoverMediaColorInfo(palette []ArticleCoverMediaColorInfoPalette) *ArticleCoverMediaColorInfo {
	this := ArticleCoverMediaColorInfo{}
	this.Palette = palette
	return &this
}

// NewArticleCoverMediaColorInfoWithDefaults instantiates a new ArticleCoverMediaColorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleCoverMediaColorInfoWithDefaults() *ArticleCoverMediaColorInfo {
	this := ArticleCoverMediaColorInfo{}
	return &this
}

// GetPalette returns the Palette field value
func (o *ArticleCoverMediaColorInfo) GetPalette() []ArticleCoverMediaColorInfoPalette {
	if o == nil {
		var ret []ArticleCoverMediaColorInfoPalette
		return ret
	}

	return o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value
// and a boolean to check if the value has been set.
func (o *ArticleCoverMediaColorInfo) GetPaletteOk() ([]ArticleCoverMediaColorInfoPalette, bool) {
	if o == nil {
		return nil, false
	}
	return o.Palette, true
}

// SetPalette sets field value
func (o *ArticleCoverMediaColorInfo) SetPalette(v []ArticleCoverMediaColorInfoPalette) {
	o.Palette = v
}

func (o ArticleCoverMediaColorInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticleCoverMediaColorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["palette"] = o.Palette
	return toSerialize, nil
}

func (o *ArticleCoverMediaColorInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"palette",
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

	varArticleCoverMediaColorInfo := _ArticleCoverMediaColorInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArticleCoverMediaColorInfo)

	if err != nil {
		return err
	}

	*o = ArticleCoverMediaColorInfo(varArticleCoverMediaColorInfo)

	return err
}

type NullableArticleCoverMediaColorInfo struct {
	value *ArticleCoverMediaColorInfo
	isSet bool
}

func (v NullableArticleCoverMediaColorInfo) Get() *ArticleCoverMediaColorInfo {
	return v.value
}

func (v *NullableArticleCoverMediaColorInfo) Set(val *ArticleCoverMediaColorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleCoverMediaColorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleCoverMediaColorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleCoverMediaColorInfo(val *ArticleCoverMediaColorInfo) *NullableArticleCoverMediaColorInfo {
	return &NullableArticleCoverMediaColorInfo{value: val, isSet: true}
}

func (v NullableArticleCoverMediaColorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleCoverMediaColorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
