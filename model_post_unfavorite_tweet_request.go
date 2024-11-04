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

// checks if the PostUnfavoriteTweetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostUnfavoriteTweetRequest{}

// PostUnfavoriteTweetRequest struct for PostUnfavoriteTweetRequest
type PostUnfavoriteTweetRequest struct {
	QueryId string `json:"queryId"`
	Variables PostCreateRetweetRequestVariables `json:"variables"`
}

type _PostUnfavoriteTweetRequest PostUnfavoriteTweetRequest

// NewPostUnfavoriteTweetRequest instantiates a new PostUnfavoriteTweetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostUnfavoriteTweetRequest(queryId string, variables PostCreateRetweetRequestVariables) *PostUnfavoriteTweetRequest {
	this := PostUnfavoriteTweetRequest{}
	this.QueryId = queryId
	this.Variables = variables
	return &this
}

// NewPostUnfavoriteTweetRequestWithDefaults instantiates a new PostUnfavoriteTweetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostUnfavoriteTweetRequestWithDefaults() *PostUnfavoriteTweetRequest {
	this := PostUnfavoriteTweetRequest{}
	var queryId string = "ZYKSe-w7KEslx3JhSIk5LA"
	this.QueryId = queryId
	return &this
}

// GetQueryId returns the QueryId field value
func (o *PostUnfavoriteTweetRequest) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *PostUnfavoriteTweetRequest) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *PostUnfavoriteTweetRequest) SetQueryId(v string) {
	o.QueryId = v
}

// GetVariables returns the Variables field value
func (o *PostUnfavoriteTweetRequest) GetVariables() PostCreateRetweetRequestVariables {
	if o == nil {
		var ret PostCreateRetweetRequestVariables
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *PostUnfavoriteTweetRequest) GetVariablesOk() (*PostCreateRetweetRequestVariables, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variables, true
}

// SetVariables sets field value
func (o *PostUnfavoriteTweetRequest) SetVariables(v PostCreateRetweetRequestVariables) {
	o.Variables = v
}

func (o PostUnfavoriteTweetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostUnfavoriteTweetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryId"] = o.QueryId
	toSerialize["variables"] = o.Variables
	return toSerialize, nil
}

func (o *PostUnfavoriteTweetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryId",
		"variables",
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

	varPostUnfavoriteTweetRequest := _PostUnfavoriteTweetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostUnfavoriteTweetRequest)

	if err != nil {
		return err
	}

	*o = PostUnfavoriteTweetRequest(varPostUnfavoriteTweetRequest)

	return err
}

type NullablePostUnfavoriteTweetRequest struct {
	value *PostUnfavoriteTweetRequest
	isSet bool
}

func (v NullablePostUnfavoriteTweetRequest) Get() *PostUnfavoriteTweetRequest {
	return v.value
}

func (v *NullablePostUnfavoriteTweetRequest) Set(val *PostUnfavoriteTweetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostUnfavoriteTweetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostUnfavoriteTweetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostUnfavoriteTweetRequest(val *PostUnfavoriteTweetRequest) *NullablePostUnfavoriteTweetRequest {
	return &NullablePostUnfavoriteTweetRequest{value: val, isSet: true}
}

func (v NullablePostUnfavoriteTweetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostUnfavoriteTweetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

