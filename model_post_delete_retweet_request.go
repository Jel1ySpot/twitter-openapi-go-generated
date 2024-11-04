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

// checks if the PostDeleteRetweetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostDeleteRetweetRequest{}

// PostDeleteRetweetRequest struct for PostDeleteRetweetRequest
type PostDeleteRetweetRequest struct {
	QueryId string `json:"queryId"`
	Variables PostDeleteRetweetRequestVariables `json:"variables"`
}

type _PostDeleteRetweetRequest PostDeleteRetweetRequest

// NewPostDeleteRetweetRequest instantiates a new PostDeleteRetweetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostDeleteRetweetRequest(queryId string, variables PostDeleteRetweetRequestVariables) *PostDeleteRetweetRequest {
	this := PostDeleteRetweetRequest{}
	this.QueryId = queryId
	this.Variables = variables
	return &this
}

// NewPostDeleteRetweetRequestWithDefaults instantiates a new PostDeleteRetweetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostDeleteRetweetRequestWithDefaults() *PostDeleteRetweetRequest {
	this := PostDeleteRetweetRequest{}
	var queryId string = "iQtK4dl5hBmXewYZuEOKVw"
	this.QueryId = queryId
	return &this
}

// GetQueryId returns the QueryId field value
func (o *PostDeleteRetweetRequest) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *PostDeleteRetweetRequest) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *PostDeleteRetweetRequest) SetQueryId(v string) {
	o.QueryId = v
}

// GetVariables returns the Variables field value
func (o *PostDeleteRetweetRequest) GetVariables() PostDeleteRetweetRequestVariables {
	if o == nil {
		var ret PostDeleteRetweetRequestVariables
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *PostDeleteRetweetRequest) GetVariablesOk() (*PostDeleteRetweetRequestVariables, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variables, true
}

// SetVariables sets field value
func (o *PostDeleteRetweetRequest) SetVariables(v PostDeleteRetweetRequestVariables) {
	o.Variables = v
}

func (o PostDeleteRetweetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostDeleteRetweetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryId"] = o.QueryId
	toSerialize["variables"] = o.Variables
	return toSerialize, nil
}

func (o *PostDeleteRetweetRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPostDeleteRetweetRequest := _PostDeleteRetweetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostDeleteRetweetRequest)

	if err != nil {
		return err
	}

	*o = PostDeleteRetweetRequest(varPostDeleteRetweetRequest)

	return err
}

type NullablePostDeleteRetweetRequest struct {
	value *PostDeleteRetweetRequest
	isSet bool
}

func (v NullablePostDeleteRetweetRequest) Get() *PostDeleteRetweetRequest {
	return v.value
}

func (v *NullablePostDeleteRetweetRequest) Set(val *PostDeleteRetweetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostDeleteRetweetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostDeleteRetweetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostDeleteRetweetRequest(val *PostDeleteRetweetRequest) *NullablePostDeleteRetweetRequest {
	return &NullablePostDeleteRetweetRequest{value: val, isSet: true}
}

func (v NullablePostDeleteRetweetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostDeleteRetweetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

