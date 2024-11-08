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

// checks if the AdditionalMediaInfoCallToActionsUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalMediaInfoCallToActionsUrl{}

// AdditionalMediaInfoCallToActionsUrl struct for AdditionalMediaInfoCallToActionsUrl
type AdditionalMediaInfoCallToActionsUrl struct {
	Url string `json:"url"`
}

type _AdditionalMediaInfoCallToActionsUrl AdditionalMediaInfoCallToActionsUrl

// NewAdditionalMediaInfoCallToActionsUrl instantiates a new AdditionalMediaInfoCallToActionsUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalMediaInfoCallToActionsUrl(url string) *AdditionalMediaInfoCallToActionsUrl {
	this := AdditionalMediaInfoCallToActionsUrl{}
	this.Url = url
	return &this
}

// NewAdditionalMediaInfoCallToActionsUrlWithDefaults instantiates a new AdditionalMediaInfoCallToActionsUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalMediaInfoCallToActionsUrlWithDefaults() *AdditionalMediaInfoCallToActionsUrl {
	this := AdditionalMediaInfoCallToActionsUrl{}
	return &this
}

// GetUrl returns the Url field value
func (o *AdditionalMediaInfoCallToActionsUrl) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AdditionalMediaInfoCallToActionsUrl) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AdditionalMediaInfoCallToActionsUrl) SetUrl(v string) {
	o.Url = v
}

func (o AdditionalMediaInfoCallToActionsUrl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalMediaInfoCallToActionsUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *AdditionalMediaInfoCallToActionsUrl) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varAdditionalMediaInfoCallToActionsUrl := _AdditionalMediaInfoCallToActionsUrl{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdditionalMediaInfoCallToActionsUrl)

	if err != nil {
		return err
	}

	*o = AdditionalMediaInfoCallToActionsUrl(varAdditionalMediaInfoCallToActionsUrl)

	return err
}

type NullableAdditionalMediaInfoCallToActionsUrl struct {
	value *AdditionalMediaInfoCallToActionsUrl
	isSet bool
}

func (v NullableAdditionalMediaInfoCallToActionsUrl) Get() *AdditionalMediaInfoCallToActionsUrl {
	return v.value
}

func (v *NullableAdditionalMediaInfoCallToActionsUrl) Set(val *AdditionalMediaInfoCallToActionsUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalMediaInfoCallToActionsUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalMediaInfoCallToActionsUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalMediaInfoCallToActionsUrl(val *AdditionalMediaInfoCallToActionsUrl) *NullableAdditionalMediaInfoCallToActionsUrl {
	return &NullableAdditionalMediaInfoCallToActionsUrl{value: val, isSet: true}
}

func (v NullableAdditionalMediaInfoCallToActionsUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalMediaInfoCallToActionsUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
