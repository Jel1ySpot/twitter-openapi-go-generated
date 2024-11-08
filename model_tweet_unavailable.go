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

// checks if the TweetUnavailable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetUnavailable{}

// TweetUnavailable struct for TweetUnavailable
type TweetUnavailable struct {
	Typename *TypeName `json:"__typename,omitempty"`
	Reason   *string   `json:"reason,omitempty"`
}

// NewTweetUnavailable instantiates a new TweetUnavailable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetUnavailable() *TweetUnavailable {
	this := TweetUnavailable{}
	return &this
}

// NewTweetUnavailableWithDefaults instantiates a new TweetUnavailable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetUnavailableWithDefaults() *TweetUnavailable {
	this := TweetUnavailable{}
	return &this
}

// GetTypename returns the Typename field value if set, zero value otherwise.
func (o *TweetUnavailable) GetTypename() TypeName {
	if o == nil || IsNil(o.Typename) {
		var ret TypeName
		return ret
	}
	return *o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetUnavailable) GetTypenameOk() (*TypeName, bool) {
	if o == nil || IsNil(o.Typename) {
		return nil, false
	}
	return o.Typename, true
}

// HasTypename returns a boolean if a field has been set.
func (o *TweetUnavailable) HasTypename() bool {
	if o != nil && !IsNil(o.Typename) {
		return true
	}

	return false
}

// SetTypename gets a reference to the given TypeName and assigns it to the Typename field.
func (o *TweetUnavailable) SetTypename(v TypeName) {
	o.Typename = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TweetUnavailable) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetUnavailable) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TweetUnavailable) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TweetUnavailable) SetReason(v string) {
	o.Reason = &v
}

func (o TweetUnavailable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetUnavailable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Typename) {
		toSerialize["__typename"] = o.Typename
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableTweetUnavailable struct {
	value *TweetUnavailable
	isSet bool
}

func (v NullableTweetUnavailable) Get() *TweetUnavailable {
	return v.value
}

func (v *NullableTweetUnavailable) Set(val *TweetUnavailable) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetUnavailable) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetUnavailable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetUnavailable(val *TweetUnavailable) *NullableTweetUnavailable {
	return &NullableTweetUnavailable{value: val, isSet: true}
}

func (v NullableTweetUnavailable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetUnavailable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
