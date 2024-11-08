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

// checks if the PostDeleteTweetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostDeleteTweetRequest{}

// PostDeleteTweetRequest struct for PostDeleteTweetRequest
type PostDeleteTweetRequest struct {
	QueryId   string                            `json:"queryId"`
	Variables PostCreateRetweetRequestVariables `json:"variables"`
}

type _PostDeleteTweetRequest PostDeleteTweetRequest

// NewPostDeleteTweetRequest instantiates a new PostDeleteTweetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostDeleteTweetRequest(queryId string, variables PostCreateRetweetRequestVariables) *PostDeleteTweetRequest {
	this := PostDeleteTweetRequest{}
	this.QueryId = queryId
	this.Variables = variables
	return &this
}

// NewPostDeleteTweetRequestWithDefaults instantiates a new PostDeleteTweetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostDeleteTweetRequestWithDefaults() *PostDeleteTweetRequest {
	this := PostDeleteTweetRequest{}
	var queryId string = "VaenaVgh5q5ih7kvyVjgtg"
	this.QueryId = queryId
	return &this
}

// GetQueryId returns the QueryId field value
func (o *PostDeleteTweetRequest) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *PostDeleteTweetRequest) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *PostDeleteTweetRequest) SetQueryId(v string) {
	o.QueryId = v
}

// GetVariables returns the Variables field value
func (o *PostDeleteTweetRequest) GetVariables() PostCreateRetweetRequestVariables {
	if o == nil {
		var ret PostCreateRetweetRequestVariables
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *PostDeleteTweetRequest) GetVariablesOk() (*PostCreateRetweetRequestVariables, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variables, true
}

// SetVariables sets field value
func (o *PostDeleteTweetRequest) SetVariables(v PostCreateRetweetRequestVariables) {
	o.Variables = v
}

func (o PostDeleteTweetRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostDeleteTweetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryId"] = o.QueryId
	toSerialize["variables"] = o.Variables
	return toSerialize, nil
}

func (o *PostDeleteTweetRequest) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostDeleteTweetRequest := _PostDeleteTweetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostDeleteTweetRequest)

	if err != nil {
		return err
	}

	*o = PostDeleteTweetRequest(varPostDeleteTweetRequest)

	return err
}

type NullablePostDeleteTweetRequest struct {
	value *PostDeleteTweetRequest
	isSet bool
}

func (v NullablePostDeleteTweetRequest) Get() *PostDeleteTweetRequest {
	return v.value
}

func (v *NullablePostDeleteTweetRequest) Set(val *PostDeleteTweetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostDeleteTweetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostDeleteTweetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostDeleteTweetRequest(val *PostDeleteTweetRequest) *NullablePostDeleteTweetRequest {
	return &NullablePostDeleteTweetRequest{value: val, isSet: true}
}

func (v NullablePostDeleteTweetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostDeleteTweetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
