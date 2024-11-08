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

// checks if the PostCreateBookmarkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCreateBookmarkRequest{}

// PostCreateBookmarkRequest struct for PostCreateBookmarkRequest
type PostCreateBookmarkRequest struct {
	QueryId   string                             `json:"queryId"`
	Variables PostCreateBookmarkRequestVariables `json:"variables"`
}

type _PostCreateBookmarkRequest PostCreateBookmarkRequest

// NewPostCreateBookmarkRequest instantiates a new PostCreateBookmarkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCreateBookmarkRequest(queryId string, variables PostCreateBookmarkRequestVariables) *PostCreateBookmarkRequest {
	this := PostCreateBookmarkRequest{}
	this.QueryId = queryId
	this.Variables = variables
	return &this
}

// NewPostCreateBookmarkRequestWithDefaults instantiates a new PostCreateBookmarkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCreateBookmarkRequestWithDefaults() *PostCreateBookmarkRequest {
	this := PostCreateBookmarkRequest{}
	var queryId string = "aoDbu3RHznuiSkQ9aNM67Q"
	this.QueryId = queryId
	return &this
}

// GetQueryId returns the QueryId field value
func (o *PostCreateBookmarkRequest) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *PostCreateBookmarkRequest) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *PostCreateBookmarkRequest) SetQueryId(v string) {
	o.QueryId = v
}

// GetVariables returns the Variables field value
func (o *PostCreateBookmarkRequest) GetVariables() PostCreateBookmarkRequestVariables {
	if o == nil {
		var ret PostCreateBookmarkRequestVariables
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *PostCreateBookmarkRequest) GetVariablesOk() (*PostCreateBookmarkRequestVariables, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variables, true
}

// SetVariables sets field value
func (o *PostCreateBookmarkRequest) SetVariables(v PostCreateBookmarkRequestVariables) {
	o.Variables = v
}

func (o PostCreateBookmarkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCreateBookmarkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryId"] = o.QueryId
	toSerialize["variables"] = o.Variables
	return toSerialize, nil
}

func (o *PostCreateBookmarkRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPostCreateBookmarkRequest := _PostCreateBookmarkRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCreateBookmarkRequest)

	if err != nil {
		return err
	}

	*o = PostCreateBookmarkRequest(varPostCreateBookmarkRequest)

	return err
}

type NullablePostCreateBookmarkRequest struct {
	value *PostCreateBookmarkRequest
	isSet bool
}

func (v NullablePostCreateBookmarkRequest) Get() *PostCreateBookmarkRequest {
	return v.value
}

func (v *NullablePostCreateBookmarkRequest) Set(val *PostCreateBookmarkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCreateBookmarkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCreateBookmarkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCreateBookmarkRequest(val *PostCreateBookmarkRequest) *NullablePostCreateBookmarkRequest {
	return &NullablePostCreateBookmarkRequest{value: val, isSet: true}
}

func (v NullablePostCreateBookmarkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCreateBookmarkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
