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

// checks if the TweetCardLegacyBindingValueDataImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetCardLegacyBindingValueDataImage{}

// TweetCardLegacyBindingValueDataImage struct for TweetCardLegacyBindingValueDataImage
type TweetCardLegacyBindingValueDataImage struct {
	Alt *string `json:"alt,omitempty"`
	Height int32 `json:"height"`
	Url string `json:"url"`
	Width int32 `json:"width"`
}

type _TweetCardLegacyBindingValueDataImage TweetCardLegacyBindingValueDataImage

// NewTweetCardLegacyBindingValueDataImage instantiates a new TweetCardLegacyBindingValueDataImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetCardLegacyBindingValueDataImage(height int32, url string, width int32) *TweetCardLegacyBindingValueDataImage {
	this := TweetCardLegacyBindingValueDataImage{}
	this.Height = height
	this.Url = url
	this.Width = width
	return &this
}

// NewTweetCardLegacyBindingValueDataImageWithDefaults instantiates a new TweetCardLegacyBindingValueDataImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetCardLegacyBindingValueDataImageWithDefaults() *TweetCardLegacyBindingValueDataImage {
	this := TweetCardLegacyBindingValueDataImage{}
	return &this
}

// GetAlt returns the Alt field value if set, zero value otherwise.
func (o *TweetCardLegacyBindingValueDataImage) GetAlt() string {
	if o == nil || IsNil(o.Alt) {
		var ret string
		return ret
	}
	return *o.Alt
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetCardLegacyBindingValueDataImage) GetAltOk() (*string, bool) {
	if o == nil || IsNil(o.Alt) {
		return nil, false
	}
	return o.Alt, true
}

// HasAlt returns a boolean if a field has been set.
func (o *TweetCardLegacyBindingValueDataImage) HasAlt() bool {
	if o != nil && !IsNil(o.Alt) {
		return true
	}

	return false
}

// SetAlt gets a reference to the given string and assigns it to the Alt field.
func (o *TweetCardLegacyBindingValueDataImage) SetAlt(v string) {
	o.Alt = &v
}

// GetHeight returns the Height field value
func (o *TweetCardLegacyBindingValueDataImage) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *TweetCardLegacyBindingValueDataImage) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *TweetCardLegacyBindingValueDataImage) SetHeight(v int32) {
	o.Height = v
}

// GetUrl returns the Url field value
func (o *TweetCardLegacyBindingValueDataImage) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TweetCardLegacyBindingValueDataImage) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TweetCardLegacyBindingValueDataImage) SetUrl(v string) {
	o.Url = v
}

// GetWidth returns the Width field value
func (o *TweetCardLegacyBindingValueDataImage) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *TweetCardLegacyBindingValueDataImage) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *TweetCardLegacyBindingValueDataImage) SetWidth(v int32) {
	o.Width = v
}

func (o TweetCardLegacyBindingValueDataImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetCardLegacyBindingValueDataImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alt) {
		toSerialize["alt"] = o.Alt
	}
	toSerialize["height"] = o.Height
	toSerialize["url"] = o.Url
	toSerialize["width"] = o.Width
	return toSerialize, nil
}

func (o *TweetCardLegacyBindingValueDataImage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"height",
		"url",
		"width",
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

	varTweetCardLegacyBindingValueDataImage := _TweetCardLegacyBindingValueDataImage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetCardLegacyBindingValueDataImage)

	if err != nil {
		return err
	}

	*o = TweetCardLegacyBindingValueDataImage(varTweetCardLegacyBindingValueDataImage)

	return err
}

type NullableTweetCardLegacyBindingValueDataImage struct {
	value *TweetCardLegacyBindingValueDataImage
	isSet bool
}

func (v NullableTweetCardLegacyBindingValueDataImage) Get() *TweetCardLegacyBindingValueDataImage {
	return v.value
}

func (v *NullableTweetCardLegacyBindingValueDataImage) Set(val *TweetCardLegacyBindingValueDataImage) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetCardLegacyBindingValueDataImage) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetCardLegacyBindingValueDataImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetCardLegacyBindingValueDataImage(val *TweetCardLegacyBindingValueDataImage) *NullableTweetCardLegacyBindingValueDataImage {
	return &NullableTweetCardLegacyBindingValueDataImage{value: val, isSet: true}
}

func (v NullableTweetCardLegacyBindingValueDataImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetCardLegacyBindingValueDataImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


