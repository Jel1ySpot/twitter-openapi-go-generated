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

// checks if the CreateRetweetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRetweetResponse{}

// CreateRetweetResponse struct for CreateRetweetResponse
type CreateRetweetResponse struct {
	Data CreateRetweetResponseData `json:"data"`
}

type _CreateRetweetResponse CreateRetweetResponse

// NewCreateRetweetResponse instantiates a new CreateRetweetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRetweetResponse(data CreateRetweetResponseData) *CreateRetweetResponse {
	this := CreateRetweetResponse{}
	this.Data = data
	return &this
}

// NewCreateRetweetResponseWithDefaults instantiates a new CreateRetweetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRetweetResponseWithDefaults() *CreateRetweetResponse {
	this := CreateRetweetResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateRetweetResponse) GetData() CreateRetweetResponseData {
	if o == nil {
		var ret CreateRetweetResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateRetweetResponse) GetDataOk() (*CreateRetweetResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateRetweetResponse) SetData(v CreateRetweetResponseData) {
	o.Data = v
}

func (o CreateRetweetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRetweetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateRetweetResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varCreateRetweetResponse := _CreateRetweetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRetweetResponse)

	if err != nil {
		return err
	}

	*o = CreateRetweetResponse(varCreateRetweetResponse)

	return err
}

type NullableCreateRetweetResponse struct {
	value *CreateRetweetResponse
	isSet bool
}

func (v NullableCreateRetweetResponse) Get() *CreateRetweetResponse {
	return v.value
}

func (v *NullableCreateRetweetResponse) Set(val *CreateRetweetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRetweetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRetweetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRetweetResponse(val *CreateRetweetResponse) *NullableCreateRetweetResponse {
	return &NullableCreateRetweetResponse{value: val, isSet: true}
}

func (v NullableCreateRetweetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRetweetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

