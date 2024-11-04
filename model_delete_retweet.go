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

// checks if the DeleteRetweet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRetweet{}

// DeleteRetweet struct for DeleteRetweet
type DeleteRetweet struct {
	Result []Retweet `json:"result"`
}

type _DeleteRetweet DeleteRetweet

// NewDeleteRetweet instantiates a new DeleteRetweet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRetweet(result []Retweet) *DeleteRetweet {
	this := DeleteRetweet{}
	this.Result = result
	return &this
}

// NewDeleteRetweetWithDefaults instantiates a new DeleteRetweet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRetweetWithDefaults() *DeleteRetweet {
	this := DeleteRetweet{}
	return &this
}

// GetResult returns the Result field value
func (o *DeleteRetweet) GetResult() []Retweet {
	if o == nil {
		var ret []Retweet
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *DeleteRetweet) GetResultOk() ([]Retweet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *DeleteRetweet) SetResult(v []Retweet) {
	o.Result = v
}

func (o DeleteRetweet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRetweet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *DeleteRetweet) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteRetweet := _DeleteRetweet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteRetweet)

	if err != nil {
		return err
	}

	*o = DeleteRetweet(varDeleteRetweet)

	return err
}

type NullableDeleteRetweet struct {
	value *DeleteRetweet
	isSet bool
}

func (v NullableDeleteRetweet) Get() *DeleteRetweet {
	return v.value
}

func (v *NullableDeleteRetweet) Set(val *DeleteRetweet) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRetweet) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRetweet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRetweet(val *DeleteRetweet) *NullableDeleteRetweet {
	return &NullableDeleteRetweet{value: val, isSet: true}
}

func (v NullableDeleteRetweet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRetweet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
