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

// checks if the TweetInterstitialTextEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetInterstitialTextEntity{}

// TweetInterstitialTextEntity struct for TweetInterstitialTextEntity
type TweetInterstitialTextEntity struct {
	FromIndex int32 `json:"fromIndex"`
	Ref TweetInterstitialTextEntityRef `json:"ref"`
	ToIndex int32 `json:"toIndex"`
}

type _TweetInterstitialTextEntity TweetInterstitialTextEntity

// NewTweetInterstitialTextEntity instantiates a new TweetInterstitialTextEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetInterstitialTextEntity(fromIndex int32, ref TweetInterstitialTextEntityRef, toIndex int32) *TweetInterstitialTextEntity {
	this := TweetInterstitialTextEntity{}
	this.FromIndex = fromIndex
	this.Ref = ref
	this.ToIndex = toIndex
	return &this
}

// NewTweetInterstitialTextEntityWithDefaults instantiates a new TweetInterstitialTextEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetInterstitialTextEntityWithDefaults() *TweetInterstitialTextEntity {
	this := TweetInterstitialTextEntity{}
	return &this
}

// GetFromIndex returns the FromIndex field value
func (o *TweetInterstitialTextEntity) GetFromIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromIndex
}

// GetFromIndexOk returns a tuple with the FromIndex field value
// and a boolean to check if the value has been set.
func (o *TweetInterstitialTextEntity) GetFromIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromIndex, true
}

// SetFromIndex sets field value
func (o *TweetInterstitialTextEntity) SetFromIndex(v int32) {
	o.FromIndex = v
}

// GetRef returns the Ref field value
func (o *TweetInterstitialTextEntity) GetRef() TweetInterstitialTextEntityRef {
	if o == nil {
		var ret TweetInterstitialTextEntityRef
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *TweetInterstitialTextEntity) GetRefOk() (*TweetInterstitialTextEntityRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *TweetInterstitialTextEntity) SetRef(v TweetInterstitialTextEntityRef) {
	o.Ref = v
}

// GetToIndex returns the ToIndex field value
func (o *TweetInterstitialTextEntity) GetToIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToIndex
}

// GetToIndexOk returns a tuple with the ToIndex field value
// and a boolean to check if the value has been set.
func (o *TweetInterstitialTextEntity) GetToIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToIndex, true
}

// SetToIndex sets field value
func (o *TweetInterstitialTextEntity) SetToIndex(v int32) {
	o.ToIndex = v
}

func (o TweetInterstitialTextEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetInterstitialTextEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fromIndex"] = o.FromIndex
	toSerialize["ref"] = o.Ref
	toSerialize["toIndex"] = o.ToIndex
	return toSerialize, nil
}

func (o *TweetInterstitialTextEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fromIndex",
		"ref",
		"toIndex",
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

	varTweetInterstitialTextEntity := _TweetInterstitialTextEntity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetInterstitialTextEntity)

	if err != nil {
		return err
	}

	*o = TweetInterstitialTextEntity(varTweetInterstitialTextEntity)

	return err
}

type NullableTweetInterstitialTextEntity struct {
	value *TweetInterstitialTextEntity
	isSet bool
}

func (v NullableTweetInterstitialTextEntity) Get() *TweetInterstitialTextEntity {
	return v.value
}

func (v *NullableTweetInterstitialTextEntity) Set(val *TweetInterstitialTextEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetInterstitialTextEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetInterstitialTextEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetInterstitialTextEntity(val *TweetInterstitialTextEntity) *NullableTweetInterstitialTextEntity {
	return &NullableTweetInterstitialTextEntity{value: val, isSet: true}
}

func (v NullableTweetInterstitialTextEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetInterstitialTextEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

