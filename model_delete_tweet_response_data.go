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

// checks if the DeleteTweetResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTweetResponseData{}

// DeleteTweetResponseData struct for DeleteTweetResponseData
type DeleteTweetResponseData struct {
	DeleteRetweet *DeleteTweetResponseResult `json:"delete_retweet,omitempty"`
}

// NewDeleteTweetResponseData instantiates a new DeleteTweetResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTweetResponseData() *DeleteTweetResponseData {
	this := DeleteTweetResponseData{}
	return &this
}

// NewDeleteTweetResponseDataWithDefaults instantiates a new DeleteTweetResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTweetResponseDataWithDefaults() *DeleteTweetResponseData {
	this := DeleteTweetResponseData{}
	return &this
}

// GetDeleteRetweet returns the DeleteRetweet field value if set, zero value otherwise.
func (o *DeleteTweetResponseData) GetDeleteRetweet() DeleteTweetResponseResult {
	if o == nil || IsNil(o.DeleteRetweet) {
		var ret DeleteTweetResponseResult
		return ret
	}
	return *o.DeleteRetweet
}

// GetDeleteRetweetOk returns a tuple with the DeleteRetweet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTweetResponseData) GetDeleteRetweetOk() (*DeleteTweetResponseResult, bool) {
	if o == nil || IsNil(o.DeleteRetweet) {
		return nil, false
	}
	return o.DeleteRetweet, true
}

// HasDeleteRetweet returns a boolean if a field has been set.
func (o *DeleteTweetResponseData) HasDeleteRetweet() bool {
	if o != nil && !IsNil(o.DeleteRetweet) {
		return true
	}

	return false
}

// SetDeleteRetweet gets a reference to the given DeleteTweetResponseResult and assigns it to the DeleteRetweet field.
func (o *DeleteTweetResponseData) SetDeleteRetweet(v DeleteTweetResponseResult) {
	o.DeleteRetweet = &v
}

func (o DeleteTweetResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTweetResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteRetweet) {
		toSerialize["delete_retweet"] = o.DeleteRetweet
	}
	return toSerialize, nil
}

type NullableDeleteTweetResponseData struct {
	value *DeleteTweetResponseData
	isSet bool
}

func (v NullableDeleteTweetResponseData) Get() *DeleteTweetResponseData {
	return v.value
}

func (v *NullableDeleteTweetResponseData) Set(val *DeleteTweetResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTweetResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTweetResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTweetResponseData(val *DeleteTweetResponseData) *NullableDeleteTweetResponseData {
	return &NullableDeleteTweetResponseData{value: val, isSet: true}
}

func (v NullableDeleteTweetResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTweetResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


