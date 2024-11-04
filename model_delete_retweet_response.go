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

// checks if the DeleteRetweetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRetweetResponse{}

// DeleteRetweetResponse struct for DeleteRetweetResponse
type DeleteRetweetResponse struct {
	Data DeleteRetweetResponseData `json:"data"`
}

type _DeleteRetweetResponse DeleteRetweetResponse

// NewDeleteRetweetResponse instantiates a new DeleteRetweetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRetweetResponse(data DeleteRetweetResponseData) *DeleteRetweetResponse {
	this := DeleteRetweetResponse{}
	this.Data = data
	return &this
}

// NewDeleteRetweetResponseWithDefaults instantiates a new DeleteRetweetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRetweetResponseWithDefaults() *DeleteRetweetResponse {
	this := DeleteRetweetResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteRetweetResponse) GetData() DeleteRetweetResponseData {
	if o == nil {
		var ret DeleteRetweetResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteRetweetResponse) GetDataOk() (*DeleteRetweetResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteRetweetResponse) SetData(v DeleteRetweetResponseData) {
	o.Data = v
}

func (o DeleteRetweetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRetweetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DeleteRetweetResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteRetweetResponse := _DeleteRetweetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteRetweetResponse)

	if err != nil {
		return err
	}

	*o = DeleteRetweetResponse(varDeleteRetweetResponse)

	return err
}

type NullableDeleteRetweetResponse struct {
	value *DeleteRetweetResponse
	isSet bool
}

func (v NullableDeleteRetweetResponse) Get() *DeleteRetweetResponse {
	return v.value
}

func (v *NullableDeleteRetweetResponse) Set(val *DeleteRetweetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRetweetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRetweetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRetweetResponse(val *DeleteRetweetResponse) *NullableDeleteRetweetResponse {
	return &NullableDeleteRetweetResponse{value: val, isSet: true}
}

func (v NullableDeleteRetweetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRetweetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
