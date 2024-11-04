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

// checks if the CreateTweet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTweet{}

// CreateTweet struct for CreateTweet
type CreateTweet struct {
	Result Tweet `json:"result"`
}

type _CreateTweet CreateTweet

// NewCreateTweet instantiates a new CreateTweet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTweet(result Tweet) *CreateTweet {
	this := CreateTweet{}
	this.Result = result
	return &this
}

// NewCreateTweetWithDefaults instantiates a new CreateTweet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTweetWithDefaults() *CreateTweet {
	this := CreateTweet{}
	return &this
}

// GetResult returns the Result field value
func (o *CreateTweet) GetResult() Tweet {
	if o == nil {
		var ret Tweet
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CreateTweet) GetResultOk() (*Tweet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CreateTweet) SetResult(v Tweet) {
	o.Result = v
}

func (o CreateTweet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTweet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *CreateTweet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varCreateTweet := _CreateTweet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTweet)

	if err != nil {
		return err
	}

	*o = CreateTweet(varCreateTweet)

	return err
}

type NullableCreateTweet struct {
	value *CreateTweet
	isSet bool
}

func (v NullableCreateTweet) Get() *CreateTweet {
	return v.value
}

func (v *NullableCreateTweet) Set(val *CreateTweet) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTweet) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTweet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTweet(val *CreateTweet) *NullableCreateTweet {
	return &NullableCreateTweet{value: val, isSet: true}
}

func (v NullableCreateTweet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTweet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
