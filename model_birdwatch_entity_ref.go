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

// checks if the BirdwatchEntityRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BirdwatchEntityRef{}

// BirdwatchEntityRef struct for BirdwatchEntityRef
type BirdwatchEntityRef struct {
	Text    *string `json:"text,omitempty"`
	Type    string  `json:"type"`
	Url     *string `json:"url,omitempty"`
	UrlType *string `json:"urlType,omitempty"`
}

type _BirdwatchEntityRef BirdwatchEntityRef

// NewBirdwatchEntityRef instantiates a new BirdwatchEntityRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBirdwatchEntityRef(type_ string) *BirdwatchEntityRef {
	this := BirdwatchEntityRef{}
	this.Type = type_
	return &this
}

// NewBirdwatchEntityRefWithDefaults instantiates a new BirdwatchEntityRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBirdwatchEntityRefWithDefaults() *BirdwatchEntityRef {
	this := BirdwatchEntityRef{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *BirdwatchEntityRef) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BirdwatchEntityRef) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *BirdwatchEntityRef) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *BirdwatchEntityRef) SetText(v string) {
	o.Text = &v
}

// GetType returns the Type field value
func (o *BirdwatchEntityRef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BirdwatchEntityRef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BirdwatchEntityRef) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BirdwatchEntityRef) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BirdwatchEntityRef) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BirdwatchEntityRef) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BirdwatchEntityRef) SetUrl(v string) {
	o.Url = &v
}

// GetUrlType returns the UrlType field value if set, zero value otherwise.
func (o *BirdwatchEntityRef) GetUrlType() string {
	if o == nil || IsNil(o.UrlType) {
		var ret string
		return ret
	}
	return *o.UrlType
}

// GetUrlTypeOk returns a tuple with the UrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BirdwatchEntityRef) GetUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UrlType) {
		return nil, false
	}
	return o.UrlType, true
}

// HasUrlType returns a boolean if a field has been set.
func (o *BirdwatchEntityRef) HasUrlType() bool {
	if o != nil && !IsNil(o.UrlType) {
		return true
	}

	return false
}

// SetUrlType gets a reference to the given string and assigns it to the UrlType field.
func (o *BirdwatchEntityRef) SetUrlType(v string) {
	o.UrlType = &v
}

func (o BirdwatchEntityRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BirdwatchEntityRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.UrlType) {
		toSerialize["urlType"] = o.UrlType
	}
	return toSerialize, nil
}

func (o *BirdwatchEntityRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varBirdwatchEntityRef := _BirdwatchEntityRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBirdwatchEntityRef)

	if err != nil {
		return err
	}

	*o = BirdwatchEntityRef(varBirdwatchEntityRef)

	return err
}

type NullableBirdwatchEntityRef struct {
	value *BirdwatchEntityRef
	isSet bool
}

func (v NullableBirdwatchEntityRef) Get() *BirdwatchEntityRef {
	return v.value
}

func (v *NullableBirdwatchEntityRef) Set(val *BirdwatchEntityRef) {
	v.value = val
	v.isSet = true
}

func (v NullableBirdwatchEntityRef) IsSet() bool {
	return v.isSet
}

func (v *NullableBirdwatchEntityRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBirdwatchEntityRef(val *BirdwatchEntityRef) *NullableBirdwatchEntityRef {
	return &NullableBirdwatchEntityRef{value: val, isSet: true}
}

func (v NullableBirdwatchEntityRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBirdwatchEntityRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
