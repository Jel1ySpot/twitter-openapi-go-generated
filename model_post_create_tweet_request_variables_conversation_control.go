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

// checks if the PostCreateTweetRequestVariablesConversationControl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCreateTweetRequestVariablesConversationControl{}

// PostCreateTweetRequestVariablesConversationControl struct for PostCreateTweetRequestVariablesConversationControl
type PostCreateTweetRequestVariablesConversationControl struct {
	Mode string `json:"mode"`
}

type _PostCreateTweetRequestVariablesConversationControl PostCreateTweetRequestVariablesConversationControl

// NewPostCreateTweetRequestVariablesConversationControl instantiates a new PostCreateTweetRequestVariablesConversationControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCreateTweetRequestVariablesConversationControl(mode string) *PostCreateTweetRequestVariablesConversationControl {
	this := PostCreateTweetRequestVariablesConversationControl{}
	this.Mode = mode
	return &this
}

// NewPostCreateTweetRequestVariablesConversationControlWithDefaults instantiates a new PostCreateTweetRequestVariablesConversationControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCreateTweetRequestVariablesConversationControlWithDefaults() *PostCreateTweetRequestVariablesConversationControl {
	this := PostCreateTweetRequestVariablesConversationControl{}
	var mode string = "Community"
	this.Mode = mode
	return &this
}

// GetMode returns the Mode field value
func (o *PostCreateTweetRequestVariablesConversationControl) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariablesConversationControl) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *PostCreateTweetRequestVariablesConversationControl) SetMode(v string) {
	o.Mode = v
}

func (o PostCreateTweetRequestVariablesConversationControl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCreateTweetRequestVariablesConversationControl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mode"] = o.Mode
	return toSerialize, nil
}

func (o *PostCreateTweetRequestVariablesConversationControl) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mode",
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

	varPostCreateTweetRequestVariablesConversationControl := _PostCreateTweetRequestVariablesConversationControl{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCreateTweetRequestVariablesConversationControl)

	if err != nil {
		return err
	}

	*o = PostCreateTweetRequestVariablesConversationControl(varPostCreateTweetRequestVariablesConversationControl)

	return err
}

type NullablePostCreateTweetRequestVariablesConversationControl struct {
	value *PostCreateTweetRequestVariablesConversationControl
	isSet bool
}

func (v NullablePostCreateTweetRequestVariablesConversationControl) Get() *PostCreateTweetRequestVariablesConversationControl {
	return v.value
}

func (v *NullablePostCreateTweetRequestVariablesConversationControl) Set(val *PostCreateTweetRequestVariablesConversationControl) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCreateTweetRequestVariablesConversationControl) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCreateTweetRequestVariablesConversationControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCreateTweetRequestVariablesConversationControl(val *PostCreateTweetRequestVariablesConversationControl) *NullablePostCreateTweetRequestVariablesConversationControl {
	return &NullablePostCreateTweetRequestVariablesConversationControl{value: val, isSet: true}
}

func (v NullablePostCreateTweetRequestVariablesConversationControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCreateTweetRequestVariablesConversationControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
