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

// checks if the NoteTweetResultMedia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoteTweetResultMedia{}

// NoteTweetResultMedia struct for NoteTweetResultMedia
type NoteTweetResultMedia struct {
	InlineMedia []NoteTweetResultMediaInlineMedia `json:"inline_media"`
}

type _NoteTweetResultMedia NoteTweetResultMedia

// NewNoteTweetResultMedia instantiates a new NoteTweetResultMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteTweetResultMedia(inlineMedia []NoteTweetResultMediaInlineMedia) *NoteTweetResultMedia {
	this := NoteTweetResultMedia{}
	this.InlineMedia = inlineMedia
	return &this
}

// NewNoteTweetResultMediaWithDefaults instantiates a new NoteTweetResultMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteTweetResultMediaWithDefaults() *NoteTweetResultMedia {
	this := NoteTweetResultMedia{}
	return &this
}

// GetInlineMedia returns the InlineMedia field value
func (o *NoteTweetResultMedia) GetInlineMedia() []NoteTweetResultMediaInlineMedia {
	if o == nil {
		var ret []NoteTweetResultMediaInlineMedia
		return ret
	}

	return o.InlineMedia
}

// GetInlineMediaOk returns a tuple with the InlineMedia field value
// and a boolean to check if the value has been set.
func (o *NoteTweetResultMedia) GetInlineMediaOk() ([]NoteTweetResultMediaInlineMedia, bool) {
	if o == nil {
		return nil, false
	}
	return o.InlineMedia, true
}

// SetInlineMedia sets field value
func (o *NoteTweetResultMedia) SetInlineMedia(v []NoteTweetResultMediaInlineMedia) {
	o.InlineMedia = v
}

func (o NoteTweetResultMedia) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoteTweetResultMedia) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inline_media"] = o.InlineMedia
	return toSerialize, nil
}

func (o *NoteTweetResultMedia) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inline_media",
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

	varNoteTweetResultMedia := _NoteTweetResultMedia{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNoteTweetResultMedia)

	if err != nil {
		return err
	}

	*o = NoteTweetResultMedia(varNoteTweetResultMedia)

	return err
}

type NullableNoteTweetResultMedia struct {
	value *NoteTweetResultMedia
	isSet bool
}

func (v NullableNoteTweetResultMedia) Get() *NoteTweetResultMedia {
	return v.value
}

func (v *NullableNoteTweetResultMedia) Set(val *NoteTweetResultMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteTweetResultMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteTweetResultMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteTweetResultMedia(val *NoteTweetResultMedia) *NullableNoteTweetResultMedia {
	return &NullableNoteTweetResultMedia{value: val, isSet: true}
}

func (v NullableNoteTweetResultMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteTweetResultMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

