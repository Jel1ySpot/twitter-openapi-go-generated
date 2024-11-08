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

// checks if the TextEntityRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextEntityRef{}

// TextEntityRef struct for TextEntityRef
type TextEntityRef struct {
	Type    string `json:"type"`
	Url     string `json:"url"`
	UrlType string `json:"urlType"`
}

type _TextEntityRef TextEntityRef

// NewTextEntityRef instantiates a new TextEntityRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextEntityRef(type_ string, url string, urlType string) *TextEntityRef {
	this := TextEntityRef{}
	this.Type = type_
	this.Url = url
	this.UrlType = urlType
	return &this
}

// NewTextEntityRefWithDefaults instantiates a new TextEntityRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextEntityRefWithDefaults() *TextEntityRef {
	this := TextEntityRef{}
	return &this
}

// GetType returns the Type field value
func (o *TextEntityRef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextEntityRef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TextEntityRef) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *TextEntityRef) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TextEntityRef) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TextEntityRef) SetUrl(v string) {
	o.Url = v
}

// GetUrlType returns the UrlType field value
func (o *TextEntityRef) GetUrlType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlType
}

// GetUrlTypeOk returns a tuple with the UrlType field value
// and a boolean to check if the value has been set.
func (o *TextEntityRef) GetUrlTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlType, true
}

// SetUrlType sets field value
func (o *TextEntityRef) SetUrlType(v string) {
	o.UrlType = v
}

func (o TextEntityRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextEntityRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	toSerialize["urlType"] = o.UrlType
	return toSerialize, nil
}

func (o *TextEntityRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"url",
		"urlType",
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

	varTextEntityRef := _TextEntityRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTextEntityRef)

	if err != nil {
		return err
	}

	*o = TextEntityRef(varTextEntityRef)

	return err
}

type NullableTextEntityRef struct {
	value *TextEntityRef
	isSet bool
}

func (v NullableTextEntityRef) Get() *TextEntityRef {
	return v.value
}

func (v *NullableTextEntityRef) Set(val *TextEntityRef) {
	v.value = val
	v.isSet = true
}

func (v NullableTextEntityRef) IsSet() bool {
	return v.isSet
}

func (v *NullableTextEntityRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextEntityRef(val *TextEntityRef) *NullableTextEntityRef {
	return &NullableTextEntityRef{value: val, isSet: true}
}

func (v NullableTextEntityRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextEntityRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
