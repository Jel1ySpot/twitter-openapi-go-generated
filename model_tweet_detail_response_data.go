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

// checks if the TweetDetailResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetDetailResponseData{}

// TweetDetailResponseData struct for TweetDetailResponseData
type TweetDetailResponseData struct {
	ThreadedConversationWithInjectionsV2 Timeline `json:"threaded_conversation_with_injections_v2"`
}

type _TweetDetailResponseData TweetDetailResponseData

// NewTweetDetailResponseData instantiates a new TweetDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetDetailResponseData(threadedConversationWithInjectionsV2 Timeline) *TweetDetailResponseData {
	this := TweetDetailResponseData{}
	this.ThreadedConversationWithInjectionsV2 = threadedConversationWithInjectionsV2
	return &this
}

// NewTweetDetailResponseDataWithDefaults instantiates a new TweetDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetDetailResponseDataWithDefaults() *TweetDetailResponseData {
	this := TweetDetailResponseData{}
	return &this
}

// GetThreadedConversationWithInjectionsV2 returns the ThreadedConversationWithInjectionsV2 field value
func (o *TweetDetailResponseData) GetThreadedConversationWithInjectionsV2() Timeline {
	if o == nil {
		var ret Timeline
		return ret
	}

	return o.ThreadedConversationWithInjectionsV2
}

// GetThreadedConversationWithInjectionsV2Ok returns a tuple with the ThreadedConversationWithInjectionsV2 field value
// and a boolean to check if the value has been set.
func (o *TweetDetailResponseData) GetThreadedConversationWithInjectionsV2Ok() (*Timeline, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadedConversationWithInjectionsV2, true
}

// SetThreadedConversationWithInjectionsV2 sets field value
func (o *TweetDetailResponseData) SetThreadedConversationWithInjectionsV2(v Timeline) {
	o.ThreadedConversationWithInjectionsV2 = v
}

func (o TweetDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threaded_conversation_with_injections_v2"] = o.ThreadedConversationWithInjectionsV2
	return toSerialize, nil
}

func (o *TweetDetailResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threaded_conversation_with_injections_v2",
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

	varTweetDetailResponseData := _TweetDetailResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetDetailResponseData)

	if err != nil {
		return err
	}

	*o = TweetDetailResponseData(varTweetDetailResponseData)

	return err
}

type NullableTweetDetailResponseData struct {
	value *TweetDetailResponseData
	isSet bool
}

func (v NullableTweetDetailResponseData) Get() *TweetDetailResponseData {
	return v.value
}

func (v *NullableTweetDetailResponseData) Set(val *TweetDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetDetailResponseData(val *TweetDetailResponseData) *NullableTweetDetailResponseData {
	return &NullableTweetDetailResponseData{value: val, isSet: true}
}

func (v NullableTweetDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


