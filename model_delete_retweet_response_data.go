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

// checks if the DeleteRetweetResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRetweetResponseData{}

// DeleteRetweetResponseData struct for DeleteRetweetResponseData
type DeleteRetweetResponseData struct {
	CreateRetweet *CreateRetweetResponseResult `json:"create_retweet,omitempty"`
}

// NewDeleteRetweetResponseData instantiates a new DeleteRetweetResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRetweetResponseData() *DeleteRetweetResponseData {
	this := DeleteRetweetResponseData{}
	return &this
}

// NewDeleteRetweetResponseDataWithDefaults instantiates a new DeleteRetweetResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRetweetResponseDataWithDefaults() *DeleteRetweetResponseData {
	this := DeleteRetweetResponseData{}
	return &this
}

// GetCreateRetweet returns the CreateRetweet field value if set, zero value otherwise.
func (o *DeleteRetweetResponseData) GetCreateRetweet() CreateRetweetResponseResult {
	if o == nil || IsNil(o.CreateRetweet) {
		var ret CreateRetweetResponseResult
		return ret
	}
	return *o.CreateRetweet
}

// GetCreateRetweetOk returns a tuple with the CreateRetweet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRetweetResponseData) GetCreateRetweetOk() (*CreateRetweetResponseResult, bool) {
	if o == nil || IsNil(o.CreateRetweet) {
		return nil, false
	}
	return o.CreateRetweet, true
}

// HasCreateRetweet returns a boolean if a field has been set.
func (o *DeleteRetweetResponseData) HasCreateRetweet() bool {
	if o != nil && !IsNil(o.CreateRetweet) {
		return true
	}

	return false
}

// SetCreateRetweet gets a reference to the given CreateRetweetResponseResult and assigns it to the CreateRetweet field.
func (o *DeleteRetweetResponseData) SetCreateRetweet(v CreateRetweetResponseResult) {
	o.CreateRetweet = &v
}

func (o DeleteRetweetResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRetweetResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateRetweet) {
		toSerialize["create_retweet"] = o.CreateRetweet
	}
	return toSerialize, nil
}

type NullableDeleteRetweetResponseData struct {
	value *DeleteRetweetResponseData
	isSet bool
}

func (v NullableDeleteRetweetResponseData) Get() *DeleteRetweetResponseData {
	return v.value
}

func (v *NullableDeleteRetweetResponseData) Set(val *DeleteRetweetResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRetweetResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRetweetResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRetweetResponseData(val *DeleteRetweetResponseData) *NullableDeleteRetweetResponseData {
	return &NullableDeleteRetweetResponseData{value: val, isSet: true}
}

func (v NullableDeleteRetweetResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRetweetResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
