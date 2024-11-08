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

// checks if the CreateTweetResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTweetResponseData{}

// CreateTweetResponseData struct for CreateTweetResponseData
type CreateTweetResponseData struct {
	CreateTweet CreateTweetResponseResult `json:"create_tweet"`
}

type _CreateTweetResponseData CreateTweetResponseData

// NewCreateTweetResponseData instantiates a new CreateTweetResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTweetResponseData(createTweet CreateTweetResponseResult) *CreateTweetResponseData {
	this := CreateTweetResponseData{}
	this.CreateTweet = createTweet
	return &this
}

// NewCreateTweetResponseDataWithDefaults instantiates a new CreateTweetResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTweetResponseDataWithDefaults() *CreateTweetResponseData {
	this := CreateTweetResponseData{}
	return &this
}

// GetCreateTweet returns the CreateTweet field value
func (o *CreateTweetResponseData) GetCreateTweet() CreateTweetResponseResult {
	if o == nil {
		var ret CreateTweetResponseResult
		return ret
	}

	return o.CreateTweet
}

// GetCreateTweetOk returns a tuple with the CreateTweet field value
// and a boolean to check if the value has been set.
func (o *CreateTweetResponseData) GetCreateTweetOk() (*CreateTweetResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateTweet, true
}

// SetCreateTweet sets field value
func (o *CreateTweetResponseData) SetCreateTweet(v CreateTweetResponseResult) {
	o.CreateTweet = v
}

func (o CreateTweetResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTweetResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["create_tweet"] = o.CreateTweet
	return toSerialize, nil
}

func (o *CreateTweetResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"create_tweet",
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

	varCreateTweetResponseData := _CreateTweetResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTweetResponseData)

	if err != nil {
		return err
	}

	*o = CreateTweetResponseData(varCreateTweetResponseData)

	return err
}

type NullableCreateTweetResponseData struct {
	value *CreateTweetResponseData
	isSet bool
}

func (v NullableCreateTweetResponseData) Get() *CreateTweetResponseData {
	return v.value
}

func (v *NullableCreateTweetResponseData) Set(val *CreateTweetResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTweetResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTweetResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTweetResponseData(val *CreateTweetResponseData) *NullableCreateTweetResponseData {
	return &NullableCreateTweetResponseData{value: val, isSet: true}
}

func (v NullableCreateTweetResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTweetResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
