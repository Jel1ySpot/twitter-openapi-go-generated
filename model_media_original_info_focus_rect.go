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

// checks if the MediaOriginalInfoFocusRect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaOriginalInfoFocusRect{}

// MediaOriginalInfoFocusRect struct for MediaOriginalInfoFocusRect
type MediaOriginalInfoFocusRect struct {
	H int32 `json:"h"`
	W int32 `json:"w"`
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type _MediaOriginalInfoFocusRect MediaOriginalInfoFocusRect

// NewMediaOriginalInfoFocusRect instantiates a new MediaOriginalInfoFocusRect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaOriginalInfoFocusRect(h int32, w int32, x int32, y int32) *MediaOriginalInfoFocusRect {
	this := MediaOriginalInfoFocusRect{}
	this.H = h
	this.W = w
	this.X = x
	this.Y = y
	return &this
}

// NewMediaOriginalInfoFocusRectWithDefaults instantiates a new MediaOriginalInfoFocusRect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaOriginalInfoFocusRectWithDefaults() *MediaOriginalInfoFocusRect {
	this := MediaOriginalInfoFocusRect{}
	return &this
}

// GetH returns the H field value
func (o *MediaOriginalInfoFocusRect) GetH() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.H
}

// GetHOk returns a tuple with the H field value
// and a boolean to check if the value has been set.
func (o *MediaOriginalInfoFocusRect) GetHOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.H, true
}

// SetH sets field value
func (o *MediaOriginalInfoFocusRect) SetH(v int32) {
	o.H = v
}

// GetW returns the W field value
func (o *MediaOriginalInfoFocusRect) GetW() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.W
}

// GetWOk returns a tuple with the W field value
// and a boolean to check if the value has been set.
func (o *MediaOriginalInfoFocusRect) GetWOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.W, true
}

// SetW sets field value
func (o *MediaOriginalInfoFocusRect) SetW(v int32) {
	o.W = v
}

// GetX returns the X field value
func (o *MediaOriginalInfoFocusRect) GetX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *MediaOriginalInfoFocusRect) GetXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *MediaOriginalInfoFocusRect) SetX(v int32) {
	o.X = v
}

// GetY returns the Y field value
func (o *MediaOriginalInfoFocusRect) GetY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *MediaOriginalInfoFocusRect) GetYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *MediaOriginalInfoFocusRect) SetY(v int32) {
	o.Y = v
}

func (o MediaOriginalInfoFocusRect) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaOriginalInfoFocusRect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["h"] = o.H
	toSerialize["w"] = o.W
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return toSerialize, nil
}

func (o *MediaOriginalInfoFocusRect) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"h",
		"w",
		"x",
		"y",
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

	varMediaOriginalInfoFocusRect := _MediaOriginalInfoFocusRect{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMediaOriginalInfoFocusRect)

	if err != nil {
		return err
	}

	*o = MediaOriginalInfoFocusRect(varMediaOriginalInfoFocusRect)

	return err
}

type NullableMediaOriginalInfoFocusRect struct {
	value *MediaOriginalInfoFocusRect
	isSet bool
}

func (v NullableMediaOriginalInfoFocusRect) Get() *MediaOriginalInfoFocusRect {
	return v.value
}

func (v *NullableMediaOriginalInfoFocusRect) Set(val *MediaOriginalInfoFocusRect) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaOriginalInfoFocusRect) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaOriginalInfoFocusRect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaOriginalInfoFocusRect(val *MediaOriginalInfoFocusRect) *NullableMediaOriginalInfoFocusRect {
	return &NullableMediaOriginalInfoFocusRect{value: val, isSet: true}
}

func (v NullableMediaOriginalInfoFocusRect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaOriginalInfoFocusRect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
