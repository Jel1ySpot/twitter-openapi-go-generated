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

// checks if the CreateBookmarkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBookmarkResponse{}

// CreateBookmarkResponse struct for CreateBookmarkResponse
type CreateBookmarkResponse struct {
	Data CreateBookmarkResponseData `json:"data"`
}

type _CreateBookmarkResponse CreateBookmarkResponse

// NewCreateBookmarkResponse instantiates a new CreateBookmarkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBookmarkResponse(data CreateBookmarkResponseData) *CreateBookmarkResponse {
	this := CreateBookmarkResponse{}
	this.Data = data
	return &this
}

// NewCreateBookmarkResponseWithDefaults instantiates a new CreateBookmarkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBookmarkResponseWithDefaults() *CreateBookmarkResponse {
	this := CreateBookmarkResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateBookmarkResponse) GetData() CreateBookmarkResponseData {
	if o == nil {
		var ret CreateBookmarkResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateBookmarkResponse) GetDataOk() (*CreateBookmarkResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateBookmarkResponse) SetData(v CreateBookmarkResponseData) {
	o.Data = v
}

func (o CreateBookmarkResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBookmarkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateBookmarkResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varCreateBookmarkResponse := _CreateBookmarkResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateBookmarkResponse)

	if err != nil {
		return err
	}

	*o = CreateBookmarkResponse(varCreateBookmarkResponse)

	return err
}

type NullableCreateBookmarkResponse struct {
	value *CreateBookmarkResponse
	isSet bool
}

func (v NullableCreateBookmarkResponse) Get() *CreateBookmarkResponse {
	return v.value
}

func (v *NullableCreateBookmarkResponse) Set(val *CreateBookmarkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBookmarkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBookmarkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBookmarkResponse(val *CreateBookmarkResponse) *NullableCreateBookmarkResponse {
	return &NullableCreateBookmarkResponse{value: val, isSet: true}
}

func (v NullableCreateBookmarkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBookmarkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
