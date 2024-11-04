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

// checks if the MediaSize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaSize{}

// MediaSize struct for MediaSize
type MediaSize struct {
	H      int32  `json:"h"`
	Resize string `json:"resize"`
	W      int32  `json:"w"`
}

type _MediaSize MediaSize

// NewMediaSize instantiates a new MediaSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaSize(h int32, resize string, w int32) *MediaSize {
	this := MediaSize{}
	this.H = h
	this.Resize = resize
	this.W = w
	return &this
}

// NewMediaSizeWithDefaults instantiates a new MediaSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaSizeWithDefaults() *MediaSize {
	this := MediaSize{}
	return &this
}

// GetH returns the H field value
func (o *MediaSize) GetH() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.H
}

// GetHOk returns a tuple with the H field value
// and a boolean to check if the value has been set.
func (o *MediaSize) GetHOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.H, true
}

// SetH sets field value
func (o *MediaSize) SetH(v int32) {
	o.H = v
}

// GetResize returns the Resize field value
func (o *MediaSize) GetResize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resize
}

// GetResizeOk returns a tuple with the Resize field value
// and a boolean to check if the value has been set.
func (o *MediaSize) GetResizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resize, true
}

// SetResize sets field value
func (o *MediaSize) SetResize(v string) {
	o.Resize = v
}

// GetW returns the W field value
func (o *MediaSize) GetW() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.W
}

// GetWOk returns a tuple with the W field value
// and a boolean to check if the value has been set.
func (o *MediaSize) GetWOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.W, true
}

// SetW sets field value
func (o *MediaSize) SetW(v int32) {
	o.W = v
}

func (o MediaSize) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaSize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["h"] = o.H
	toSerialize["resize"] = o.Resize
	toSerialize["w"] = o.W
	return toSerialize, nil
}

func (o *MediaSize) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"h",
		"resize",
		"w",
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

	varMediaSize := _MediaSize{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMediaSize)

	if err != nil {
		return err
	}

	*o = MediaSize(varMediaSize)

	return err
}

type NullableMediaSize struct {
	value *MediaSize
	isSet bool
}

func (v NullableMediaSize) Get() *MediaSize {
	return v.value
}

func (v *NullableMediaSize) Set(val *MediaSize) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaSize) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaSize(val *MediaSize) *NullableMediaSize {
	return &NullableMediaSize{value: val, isSet: true}
}

func (v NullableMediaSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
