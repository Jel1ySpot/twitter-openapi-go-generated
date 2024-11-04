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

// checks if the TweetDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetDetailResponse{}

// TweetDetailResponse struct for TweetDetailResponse
type TweetDetailResponse struct {
	Data TweetDetailResponseData `json:"data"`
}

type _TweetDetailResponse TweetDetailResponse

// NewTweetDetailResponse instantiates a new TweetDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetDetailResponse(data TweetDetailResponseData) *TweetDetailResponse {
	this := TweetDetailResponse{}
	this.Data = data
	return &this
}

// NewTweetDetailResponseWithDefaults instantiates a new TweetDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetDetailResponseWithDefaults() *TweetDetailResponse {
	this := TweetDetailResponse{}
	return &this
}

// GetData returns the Data field value
func (o *TweetDetailResponse) GetData() TweetDetailResponseData {
	if o == nil {
		var ret TweetDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TweetDetailResponse) GetDataOk() (*TweetDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TweetDetailResponse) SetData(v TweetDetailResponseData) {
	o.Data = v
}

func (o TweetDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *TweetDetailResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varTweetDetailResponse := _TweetDetailResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetDetailResponse)

	if err != nil {
		return err
	}

	*o = TweetDetailResponse(varTweetDetailResponse)

	return err
}

type NullableTweetDetailResponse struct {
	value *TweetDetailResponse
	isSet bool
}

func (v NullableTweetDetailResponse) Get() *TweetDetailResponse {
	return v.value
}

func (v *NullableTweetDetailResponse) Set(val *TweetDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetDetailResponse(val *TweetDetailResponse) *NullableTweetDetailResponse {
	return &NullableTweetDetailResponse{value: val, isSet: true}
}

func (v NullableTweetDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
