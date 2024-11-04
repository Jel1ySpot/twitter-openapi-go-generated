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

// checks if the TweetRetweetersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetRetweetersResponse{}

// TweetRetweetersResponse struct for TweetRetweetersResponse
type TweetRetweetersResponse struct {
	Data TweetRetweetersResponseData `json:"data"`
}

type _TweetRetweetersResponse TweetRetweetersResponse

// NewTweetRetweetersResponse instantiates a new TweetRetweetersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetRetweetersResponse(data TweetRetweetersResponseData) *TweetRetweetersResponse {
	this := TweetRetweetersResponse{}
	this.Data = data
	return &this
}

// NewTweetRetweetersResponseWithDefaults instantiates a new TweetRetweetersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetRetweetersResponseWithDefaults() *TweetRetweetersResponse {
	this := TweetRetweetersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *TweetRetweetersResponse) GetData() TweetRetweetersResponseData {
	if o == nil {
		var ret TweetRetweetersResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TweetRetweetersResponse) GetDataOk() (*TweetRetweetersResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TweetRetweetersResponse) SetData(v TweetRetweetersResponseData) {
	o.Data = v
}

func (o TweetRetweetersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetRetweetersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *TweetRetweetersResponse) UnmarshalJSON(data []byte) (err error) {
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

	varTweetRetweetersResponse := _TweetRetweetersResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetRetweetersResponse)

	if err != nil {
		return err
	}

	*o = TweetRetweetersResponse(varTweetRetweetersResponse)

	return err
}

type NullableTweetRetweetersResponse struct {
	value *TweetRetweetersResponse
	isSet bool
}

func (v NullableTweetRetweetersResponse) Get() *TweetRetweetersResponse {
	return v.value
}

func (v *NullableTweetRetweetersResponse) Set(val *TweetRetweetersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetRetweetersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetRetweetersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetRetweetersResponse(val *TweetRetweetersResponse) *NullableTweetRetweetersResponse {
	return &NullableTweetRetweetersResponse{value: val, isSet: true}
}

func (v NullableTweetRetweetersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetRetweetersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

