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
)

// checks if the TweetCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetCard{}

// TweetCard struct for TweetCard
type TweetCard struct {
	Legacy *TweetCardLegacy `json:"legacy,omitempty"`
	RestId *string `json:"rest_id,omitempty"`
}

// NewTweetCard instantiates a new TweetCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetCard() *TweetCard {
	this := TweetCard{}
	return &this
}

// NewTweetCardWithDefaults instantiates a new TweetCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetCardWithDefaults() *TweetCard {
	this := TweetCard{}
	return &this
}

// GetLegacy returns the Legacy field value if set, zero value otherwise.
func (o *TweetCard) GetLegacy() TweetCardLegacy {
	if o == nil || IsNil(o.Legacy) {
		var ret TweetCardLegacy
		return ret
	}
	return *o.Legacy
}

// GetLegacyOk returns a tuple with the Legacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetCard) GetLegacyOk() (*TweetCardLegacy, bool) {
	if o == nil || IsNil(o.Legacy) {
		return nil, false
	}
	return o.Legacy, true
}

// HasLegacy returns a boolean if a field has been set.
func (o *TweetCard) HasLegacy() bool {
	if o != nil && !IsNil(o.Legacy) {
		return true
	}

	return false
}

// SetLegacy gets a reference to the given TweetCardLegacy and assigns it to the Legacy field.
func (o *TweetCard) SetLegacy(v TweetCardLegacy) {
	o.Legacy = &v
}

// GetRestId returns the RestId field value if set, zero value otherwise.
func (o *TweetCard) GetRestId() string {
	if o == nil || IsNil(o.RestId) {
		var ret string
		return ret
	}
	return *o.RestId
}

// GetRestIdOk returns a tuple with the RestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetCard) GetRestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RestId) {
		return nil, false
	}
	return o.RestId, true
}

// HasRestId returns a boolean if a field has been set.
func (o *TweetCard) HasRestId() bool {
	if o != nil && !IsNil(o.RestId) {
		return true
	}

	return false
}

// SetRestId gets a reference to the given string and assigns it to the RestId field.
func (o *TweetCard) SetRestId(v string) {
	o.RestId = &v
}

func (o TweetCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Legacy) {
		toSerialize["legacy"] = o.Legacy
	}
	if !IsNil(o.RestId) {
		toSerialize["rest_id"] = o.RestId
	}
	return toSerialize, nil
}

type NullableTweetCard struct {
	value *TweetCard
	isSet bool
}

func (v NullableTweetCard) Get() *TweetCard {
	return v.value
}

func (v *NullableTweetCard) Set(val *TweetCard) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetCard) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetCard(val *TweetCard) *NullableTweetCard {
	return &NullableTweetCard{value: val, isSet: true}
}

func (v NullableTweetCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


