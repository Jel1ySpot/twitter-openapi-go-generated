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

// checks if the TweetResultByRestIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetResultByRestIdResponse{}

// TweetResultByRestIdResponse struct for TweetResultByRestIdResponse
type TweetResultByRestIdResponse struct {
	Data TweetResultByRestIdData `json:"data"`
}

type _TweetResultByRestIdResponse TweetResultByRestIdResponse

// NewTweetResultByRestIdResponse instantiates a new TweetResultByRestIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetResultByRestIdResponse(data TweetResultByRestIdData) *TweetResultByRestIdResponse {
	this := TweetResultByRestIdResponse{}
	this.Data = data
	return &this
}

// NewTweetResultByRestIdResponseWithDefaults instantiates a new TweetResultByRestIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetResultByRestIdResponseWithDefaults() *TweetResultByRestIdResponse {
	this := TweetResultByRestIdResponse{}
	return &this
}

// GetData returns the Data field value
func (o *TweetResultByRestIdResponse) GetData() TweetResultByRestIdData {
	if o == nil {
		var ret TweetResultByRestIdData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TweetResultByRestIdResponse) GetDataOk() (*TweetResultByRestIdData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TweetResultByRestIdResponse) SetData(v TweetResultByRestIdData) {
	o.Data = v
}

func (o TweetResultByRestIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetResultByRestIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *TweetResultByRestIdResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varTweetResultByRestIdResponse := _TweetResultByRestIdResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetResultByRestIdResponse)

	if err != nil {
		return err
	}

	*o = TweetResultByRestIdResponse(varTweetResultByRestIdResponse)

	return err
}

type NullableTweetResultByRestIdResponse struct {
	value *TweetResultByRestIdResponse
	isSet bool
}

func (v NullableTweetResultByRestIdResponse) Get() *TweetResultByRestIdResponse {
	return v.value
}

func (v *NullableTweetResultByRestIdResponse) Set(val *TweetResultByRestIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetResultByRestIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetResultByRestIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetResultByRestIdResponse(val *TweetResultByRestIdResponse) *NullableTweetResultByRestIdResponse {
	return &NullableTweetResultByRestIdResponse{value: val, isSet: true}
}

func (v NullableTweetResultByRestIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetResultByRestIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

